package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/steebchen/prisma-client-go/engine/protocol"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"immotep/backend/database"
	"immotep/backend/models"
	"immotep/backend/prisma/db"
	"immotep/backend/router"
	"immotep/backend/utils"
)

func BuildTestProperty(id string) db.PropertyModel {
	return db.PropertyModel{
		InnerProperty: db.InnerProperty{
			ID:                  id,
			Name:                "Test",
			Address:             "Test",
			City:                "Test",
			PostalCode:          "Test",
			Country:             "Test",
			AreaSqm:             20.0,
			RentalPricePerMonth: 500,
			DepositPrice:        1000,
			CreatedAt:           time.Now(),
			OwnerID:             "1",
			PictureID:           utils.Ptr("1"),
		},
		RelationsProperty: db.RelationsProperty{
			Damages:   []db.DamageModel{{}},
			Contracts: []db.ContractModel{{}},
		},
	}
}

func BuildTestPendingContract() db.PendingContractModel {
	return db.PendingContractModel{
		InnerPendingContract: db.InnerPendingContract{
			ID:          "1",
			PropertyID:  "1",
			TenantEmail: "test.test@example.com",
			CreatedAt:   time.Now(),
			StartDate:   time.Now(),
			EndDate:     utils.Ptr(time.Now().Add(time.Hour)),
		},
		RelationsPendingContract: db.RelationsPendingContract{
			Property: &db.PropertyModel{
				RelationsProperty: db.RelationsProperty{
					Owner: &db.UserModel{
						InnerUser: db.InnerUser{
							Firstname: "Test",
							Lastname:  "Test",
						},
					},
				},
			},
		},
	}
}

func BuildTestContract() db.ContractModel {
	return db.ContractModel{
		InnerContract: db.InnerContract{
			PropertyID: "1",
			TenantID:   "1",
			Active:     true,
			CreatedAt:  time.Now(),
			StartDate:  time.Now(),
			EndDate:    utils.Ptr(time.Now().Add(time.Hour)),
		},
	}
}

func TestGetAllProperties(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(
		client.Client.Property.FindMany(
			db.Property.OwnerID.Equals("1"),
		).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
		),
	).ReturnsMany([]db.PropertyModel{property})

	r := router.TestRoutes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/owner/properties/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	var resp []models.PropertyResponse
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.JSONEq(t, resp[0].ID, property.ID)
}

func TestGetPropertyById(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals(property.ID)).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
		),
	).Returns(property)

	r := router.TestRoutes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/owner/properties/1/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	var propertyResponse models.PropertyResponse
	err := json.Unmarshal(w.Body.Bytes(), &propertyResponse)
	require.NoError(t, err)
	assert.Equal(t, property.ID, propertyResponse.ID)
}

func TestGetPropertyById_NotFound(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals(property.ID)).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
		),
	).Errors(db.ErrNotFound)

	r := router.TestRoutes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/owner/properties/1/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)
	var errorResponse utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
	require.NoError(t, err)
	assert.Equal(t, utils.PropertyNotFound, errorResponse.Code)
}

func TestGetPropertyById_NotYours(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals(property.ID)).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
		),
	).Returns(property)

	r := router.TestRoutes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/owner/properties/1/", nil)
	req.Header.Set("Oauth.claims.id", "2")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusForbidden, w.Code)
	var errorResponse utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
	require.NoError(t, err)
	assert.Equal(t, utils.PropertyNotYours, errorResponse.Code)
}

func TestCreateProperty(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(
		client.Client.Property.CreateOne(
			db.Property.Name.Set(property.Name),
			db.Property.Address.Set(property.Address),
			db.Property.City.Set(property.City),
			db.Property.PostalCode.Set(property.PostalCode),
			db.Property.Country.Set(property.Country),
			db.Property.AreaSqm.Set(property.AreaSqm),
			db.Property.RentalPricePerMonth.Set(property.RentalPricePerMonth),
			db.Property.DepositPrice.Set(property.DepositPrice),
			db.Property.Owner.Link(db.User.ID.Equals("1")),
		).With(
			db.Property.Contracts.Fetch(),
			db.Property.Damages.Fetch(),
		),
	).Returns(property)

	b, err := json.Marshal(property)
	require.NoError(t, err)

	r := router.TestRoutes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/owner/properties/", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusCreated, w.Code)
	var resp models.PropertyResponse
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.JSONEq(t, resp.ID, property.ID)
}

func TestCreateProperty_MissingFields(t *testing.T) {
	property := BuildTestProperty("1")
	property.Country = ""
	b, err := json.Marshal(property)
	require.NoError(t, err)

	r := router.TestRoutes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/owner/properties/", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusBadRequest, w.Code)
	var resp utils.Error
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.MissingFields, resp.Code)
}

func TestCreateProperty_AlreadyExists(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(
		client.Client.Property.CreateOne(
			db.Property.Name.Set(property.Name),
			db.Property.Address.Set(property.Address),
			db.Property.City.Set(property.City),
			db.Property.PostalCode.Set(property.PostalCode),
			db.Property.Country.Set(property.Country),
			db.Property.AreaSqm.Set(property.AreaSqm),
			db.Property.RentalPricePerMonth.Set(property.RentalPricePerMonth),
			db.Property.DepositPrice.Set(property.DepositPrice),
			db.Property.Owner.Link(db.User.ID.Equals("1")),
		).With(
			db.Property.Contracts.Fetch(),
			db.Property.Damages.Fetch(),
		),
	).Errors(&protocol.UserFacingError{
		IsPanic:   false,
		ErrorCode: "P2002", // https://www.prisma.io/docs/orm/reference/error-reference#p2002
		Meta: protocol.Meta{
			Target: []any{"name"},
		},
		Message: "Unique constraint failed",
	})

	b, err := json.Marshal(property)
	require.NoError(t, err)

	r := router.TestRoutes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/owner/properties/", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusConflict, w.Code)
	var resp utils.Error
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.PropertyAlreadyExists, resp.Code)
}

func TestInviteTenant(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals(property.ID)).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
		),
	).Returns(property)

	mock.Contract.Expect(
		client.Client.Contract.FindMany(
			db.Contract.PropertyID.Equals(property.ID),
			db.Contract.Active.Equals(true),
		),
	).Errors(db.ErrNotFound)

	pendingContract := BuildTestPendingContract()
	mock.PendingContract.Expect(
		client.Client.PendingContract.CreateOne(
			db.PendingContract.TenantEmail.Set(pendingContract.TenantEmail),
			db.PendingContract.StartDate.Set(pendingContract.StartDate),
			db.PendingContract.Property.Link(db.Property.ID.Equals(property.ID)),
			db.PendingContract.EndDate.SetIfPresent(pendingContract.InnerPendingContract.EndDate),
		).With(
			db.PendingContract.Property.Fetch().With(db.Property.Owner.Fetch()),
		),
	).Returns(pendingContract)

	reqBody := models.InviteRequest{
		TenantEmail: pendingContract.TenantEmail,
		StartDate:   pendingContract.StartDate,
		EndDate:     pendingContract.InnerPendingContract.EndDate,
	}
	b, err := json.Marshal(reqBody)
	require.NoError(t, err)

	r := router.TestRoutes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/owner/properties/1/send-invite/", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	var resp models.InviteResponse
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.JSONEq(t, resp.ID, pendingContract.ID)
}

func TestInviteTenant_MissingField(t *testing.T) {
	pendingContract := BuildTestPendingContract()
	reqBody := models.InviteRequest{
		StartDate: pendingContract.StartDate,
		EndDate:   pendingContract.InnerPendingContract.EndDate,
	}
	b, err := json.Marshal(reqBody)
	require.NoError(t, err)

	r := router.TestRoutes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/owner/properties/1/send-invite/", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusBadRequest, w.Code)
	var resp utils.Error
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.MissingFields, resp.Code)
}

func TestInviteTenant_PropertyNotFound(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals("wrong")).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
		),
	).Errors(db.ErrNotFound)

	pendingContract := BuildTestPendingContract()

	reqBody := models.InviteRequest{
		TenantEmail: pendingContract.TenantEmail,
		StartDate:   pendingContract.StartDate,
		EndDate:     pendingContract.InnerPendingContract.EndDate,
	}
	b, err := json.Marshal(reqBody)
	require.NoError(t, err)

	r := router.TestRoutes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/owner/properties/wrong/send-invite/", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)
	var resp utils.Error
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.PropertyNotFound, resp.Code)
}

func TestInviteTenant_PropertyNotYours(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals(property.ID)).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
		),
	).Returns(property)

	pendingContract := BuildTestPendingContract()

	reqBody := models.InviteRequest{
		TenantEmail: pendingContract.TenantEmail,
		StartDate:   pendingContract.StartDate,
		EndDate:     pendingContract.InnerPendingContract.EndDate,
	}
	b, err := json.Marshal(reqBody)
	require.NoError(t, err)

	r := router.TestRoutes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/owner/properties/1/send-invite/", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "wrong")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusForbidden, w.Code)
	var resp utils.Error
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.PropertyNotYours, resp.Code)
}

func TestInviteTenant_PropertyNotAvailable(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals(property.ID)).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
		),
	).Returns(property)

	contract := BuildTestContract()
	mock.Contract.Expect(
		client.Client.Contract.FindMany(
			db.Contract.PropertyID.Equals(property.ID),
			db.Contract.Active.Equals(true),
		),
	).ReturnsMany([]db.ContractModel{contract})

	pendingContract := BuildTestPendingContract()
	reqBody := models.InviteRequest{
		TenantEmail: pendingContract.TenantEmail,
		StartDate:   pendingContract.StartDate,
		EndDate:     pendingContract.InnerPendingContract.EndDate,
	}
	b, err := json.Marshal(reqBody)
	require.NoError(t, err)

	r := router.TestRoutes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/owner/properties/1/send-invite/", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusConflict, w.Code)
	var resp utils.Error
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.PropertyNotAvailable, resp.Code)
}

func TestInviteTenant_AlreadyExists(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals(property.ID)).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
		),
	).Returns(property)

	mock.Contract.Expect(
		client.Client.Contract.FindMany(
			db.Contract.PropertyID.Equals(property.ID),
			db.Contract.Active.Equals(true),
		),
	).Errors(db.ErrNotFound)

	pendingContract := BuildTestPendingContract()
	mock.PendingContract.Expect(
		client.Client.PendingContract.CreateOne(
			db.PendingContract.TenantEmail.Set(pendingContract.TenantEmail),
			db.PendingContract.StartDate.Set(pendingContract.StartDate),
			db.PendingContract.Property.Link(db.Property.ID.Equals(property.ID)),
			db.PendingContract.EndDate.SetIfPresent(pendingContract.InnerPendingContract.EndDate),
		).With(
			db.PendingContract.Property.Fetch().With(db.Property.Owner.Fetch()),
		),
	).Errors(&protocol.UserFacingError{
		IsPanic:   false,
		ErrorCode: "P2002", // https://www.prisma.io/docs/orm/reference/error-reference
		Meta: protocol.Meta{
			Target: []any{"tenant_email"},
		},
		Message: "Unique constraint failed",
	})

	reqBody := models.InviteRequest{
		TenantEmail: pendingContract.TenantEmail,
		StartDate:   pendingContract.StartDate,
		EndDate:     pendingContract.InnerPendingContract.EndDate,
	}
	b, err := json.Marshal(reqBody)
	require.NoError(t, err)

	r := router.TestRoutes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/owner/properties/1/send-invite/", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusConflict, w.Code)
	var resp utils.Error
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.InviteAlreadyExists, resp.Code)
}

func TestGetPropertyPicture(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals(property.ID)).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
		),
	).Returns(property)

	image := BuildTestImage("1", "b3Vp")
	mock.Image.Expect(
		client.Client.Image.FindUnique(db.Image.ID.Equals(image.ID)),
	).Returns(image)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/owner/properties/1/picture/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	var resp models.ImageResponse
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, image.ID, resp.ID)
}

func TestGetPropertyPicture_NoContent(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	property.InnerProperty.PictureID = nil
	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals(property.ID)).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
		),
	).Returns(property)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/owner/properties/1/picture/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNoContent, w.Code)
}

func TestGetPropertyPicture_NotFound(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals(property.ID)).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
		),
	).Returns(property)

	mock.Image.Expect(
		client.Client.Image.FindUnique(db.Image.ID.Equals("1")),
	).Errors(db.ErrNotFound)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/owner/properties/1/picture/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)
	var resp utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.PropertyPictureNotFound, resp.Code)
}

func TestGetPropertyPicture_PropertyNotFound(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals("wrong")).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
		),
	).Errors(db.ErrNotFound)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/owner/properties/wrong/picture/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)
	var resp utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.PropertyNotFound, resp.Code)
}

func TestUpdatePropertyPicture(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals(property.ID)).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
		),
	).Returns(property)

	image := BuildTestImage("1", "b3Vp")
	mock.Image.Expect(
		client.Client.Image.CreateOne(
			db.Image.Data.Set(image.Data),
		),
	).Returns(image)

	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals(property.ID)).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
		).Update(
			db.Property.Picture.Link(db.Image.ID.Equals(image.ID)),
		),
	).Returns(property)

	reqBody := models.ImageRequest{
		Data: "b3Vp",
	}
	b, err := json.Marshal(reqBody)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/owner/properties/1/picture/", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	var resp models.PropertyResponse
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, property.ID, resp.ID)
}

func TestUpdatePropertyPicture_MissingFields(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals(property.ID)).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
		),
	).Returns(property)

	reqBody := models.ImageRequest{}
	b, err := json.Marshal(reqBody)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/owner/properties/1/picture/", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusBadRequest, w.Code)
	var resp utils.Error
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.MissingFields, resp.Code)
}

func TestUpdatePropertyPicture_BadBase64String(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals(property.ID)).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
		),
	).Returns(property)

	reqBody := models.ImageRequest{
		Data: "invalid_base64",
	}
	b, err := json.Marshal(reqBody)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/owner/properties/1/picture/", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusBadRequest, w.Code)
	var resp utils.Error
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.MissingFields, resp.Code)
}

func TestUpdatePropertyPicture_PropertyNotFound(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals("wrong")).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
		),
	).Errors(db.ErrNotFound)

	reqBody := models.ImageRequest{
		Data: "b3Vp",
	}
	b, err := json.Marshal(reqBody)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/owner/properties/wrong/picture/", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)
	var resp utils.Error
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.PropertyNotFound, resp.Code)
}

func TestUpdatePropertyPicture_FailedLinkImage(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals(property.ID)).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
		),
	).Returns(property)

	image := BuildTestImage("1", "b3Vp")
	mock.Image.Expect(
		client.Client.Image.CreateOne(
			db.Image.Data.Set(image.Data),
		),
	).Returns(image)

	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals(property.ID)).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
		).Update(
			db.Property.Picture.Link(db.Image.ID.Equals(image.ID)),
		),
	).Errors(db.ErrNotFound)

	reqBody := models.ImageRequest{
		Data: "b3Vp",
	}
	b, err := json.Marshal(reqBody)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/owner/properties/1/picture/", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusInternalServerError, w.Code)
	var resp utils.Error
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.FailedLinkImage, resp.Code)
}

func TestEndContract1(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals(property.ID)).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
		),
	).Returns(property)

	contract := BuildTestContract()
	mock.Contract.Expect(
		client.Client.Contract.FindMany(
			db.Contract.PropertyID.Equals(property.ID),
			db.Contract.Active.Equals(true),
		),
	).ReturnsMany([]db.ContractModel{contract})

	mock.Contract.Expect(
		client.Client.Contract.FindUnique(
			db.Contract.TenantIDPropertyID(db.Contract.TenantID.Equals(contract.TenantID), db.Contract.PropertyID.Equals(contract.PropertyID)),
		).Update(
			db.Contract.Active.Set(false),
			db.Contract.EndDate.SetIfPresent(nil),
		),
	).Returns(contract)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/owner/properties/1/end-contract/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNoContent, w.Code)
}

func TestEndContract2(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals(property.ID)).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
		),
	).Returns(property)

	contract := BuildTestContract()
	contract.InnerContract.EndDate = nil
	mock.Contract.Expect(
		client.Client.Contract.FindMany(
			db.Contract.PropertyID.Equals(property.ID),
			db.Contract.Active.Equals(true),
		),
	).ReturnsMany([]db.ContractModel{contract})

	mock.Contract.Expect(
		client.Client.Contract.FindUnique(
			db.Contract.TenantIDPropertyID(db.Contract.TenantID.Equals(contract.TenantID), db.Contract.PropertyID.Equals(contract.PropertyID)),
		).Update(
			db.Contract.Active.Set(false),
			db.Contract.EndDate.SetIfPresent(utils.Ptr(time.Now().Truncate(time.Minute))),
		),
	).Returns(contract)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/owner/properties/1/end-contract/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNoContent, w.Code)
}

func TestEndContract_NoActiveContract(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals(property.ID)).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
		),
	).Returns(property)

	mock.Contract.Expect(
		client.Client.Contract.FindMany(
			db.Contract.PropertyID.Equals(property.ID),
			db.Contract.Active.Equals(true),
		),
	).Errors(db.ErrNotFound)

	r := router.TestRoutes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/owner/properties/1/end-contract/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)
	var resp utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.NoActiveContract, resp.Code)
}

func TestEndContract_PropertyNotFound(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals("wrong")).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
		),
	).Errors(db.ErrNotFound)

	r := router.TestRoutes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/owner/properties/wrong/end-contract/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)
	var resp utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.PropertyNotFound, resp.Code)
}
