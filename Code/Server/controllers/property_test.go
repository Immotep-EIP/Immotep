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
		},
		RelationsProperty: db.RelationsProperty{
			Damages:   []db.DamageModel{{}},
			Contracts: []db.ContractModel{{}},
		},
	}
}

func BuildTestPendingContract() db.PendingContractModel {
	end := time.Now().Add(time.Hour)
	return db.PendingContractModel{
		InnerPendingContract: db.InnerPendingContract{
			ID:          "1",
			TenantEmail: "test.test@example.com",
			StartDate:   time.Now(),
			EndDate:     &end,
			PropertyID:  "1",
			CreatedAt:   time.Now(),
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
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/owner/properties/"+property.ID+"/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	var propertyResponse models.PropertyResponse
	err := json.Unmarshal(w.Body.Bytes(), &propertyResponse)
	require.NoError(t, err)
	assert.Equal(t, property.ID, propertyResponse.ID)
}

func TestGetPropertyByIdNotFound(t *testing.T) {
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
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/owner/properties/"+property.ID+"/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)
	var errorResponse utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
	require.NoError(t, err)
	assert.Equal(t, utils.PropertyNotFound, errorResponse.Code)
}

func TestGetPropertyByIdNotYours(t *testing.T) {
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
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/owner/properties/"+property.ID+"/", nil)
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

func TestCreatePropertyMissingFields(t *testing.T) {
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

func TestCreatePropertyAlreadyExists(t *testing.T) {
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
		ErrorCode: "P2002", // https://www.prisma.io/docs/orm/reference/error-reference
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
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/owner/properties/"+property.ID+"/send-invite/", bytes.NewReader(b))
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

func TestInviteTenantMissingField(t *testing.T) {
	property := BuildTestProperty("1")
	pendingContract := BuildTestPendingContract()
	reqBody := models.InviteRequest{
		StartDate: pendingContract.StartDate,
		EndDate:   pendingContract.InnerPendingContract.EndDate,
	}
	b, err := json.Marshal(reqBody)
	require.NoError(t, err)

	r := router.TestRoutes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/owner/properties/"+property.ID+"/send-invite/", bytes.NewReader(b))
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

func TestInviteTenantPropertyNotFound(t *testing.T) {
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

func TestInviteTenantPropertyNotYours(t *testing.T) {
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
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/owner/properties/"+property.ID+"/send-invite/", bytes.NewReader(b))
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

func TestInviteTenantAlreadyExists(t *testing.T) {
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
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/owner/properties/"+property.ID+"/send-invite/", bytes.NewReader(b))
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
