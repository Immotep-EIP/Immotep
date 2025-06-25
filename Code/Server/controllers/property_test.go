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
	"keyz/backend/models"
	"keyz/backend/prisma/db"
	"keyz/backend/router"
	"keyz/backend/services"
	"keyz/backend/services/database"
	"keyz/backend/utils"
)

func BuildTestProperty(id string) db.PropertyModel {
	return db.PropertyModel{
		InnerProperty: db.InnerProperty{
			ID:                  id,
			Name:                "Test",
			Address:             "Test",
			ApartmentNumber:     utils.Ptr("Test"),
			City:                "Test",
			PostalCode:          "Test",
			Country:             "Test",
			AreaSqm:             20.0,
			RentalPricePerMonth: 500,
			DepositPrice:        1000,
			CreatedAt:           time.Now(),
			OwnerID:             "1",
			PictureID:           utils.Ptr("1"),
			Archived:            false,
		},
		RelationsProperty: db.RelationsProperty{
			Leases: []db.LeaseModel{{
				InnerLease: db.InnerLease{
					ID:     "1",
					Active: true,
				},
				RelationsLease: db.RelationsLease{
					Tenant: &db.UserModel{
						InnerUser: db.InnerUser{
							Firstname: "Test",
							Lastname:  "Test",
						},
					},
					Damages: []db.DamageModel{{
						InnerDamage: db.InnerDamage{
							ID:      "1",
							FixedAt: nil,
						}},
					},
				},
			}},
			LeaseInvite: &db.LeaseInviteModel{},
		},
	}
}

func BuildTestPropertyWithInventory(id string) db.PropertyModel {
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
			Archived:            false,
		},
		RelationsProperty: db.RelationsProperty{
			Leases: []db.LeaseModel{{
				InnerLease: db.InnerLease{
					ID:     "1",
					Active: true,
				},
				RelationsLease: db.RelationsLease{
					Tenant: &db.UserModel{
						InnerUser: db.InnerUser{
							Firstname: "Test",
							Lastname:  "Test",
						},
					},
					Damages: []db.DamageModel{{
						InnerDamage: db.InnerDamage{
							ID:      "1",
							FixedAt: nil,
						}},
					},
				},
			}},
			LeaseInvite: &db.LeaseInviteModel{},
			Rooms: []db.RoomModel{
				{
					InnerRoom: db.InnerRoom{
						ID:         "1",
						Name:       "Test",
						Archived:   false,
						PropertyID: id,
					},
					RelationsRoom: db.RelationsRoom{
						Furnitures: []db.FurnitureModel{{}},
					},
				},
			},
		},
	}
}

func BuildTestLeaseInvite() db.LeaseInviteModel {
	return db.LeaseInviteModel{
		InnerLeaseInvite: db.InnerLeaseInvite{
			ID:          "1",
			PropertyID:  "1",
			TenantEmail: "test@example.com",
			CreatedAt:   time.Now(),
			StartDate:   time.Now(),
			EndDate:     utils.Ptr(time.Now().Add(time.Hour)),
		},
		RelationsLeaseInvite: db.RelationsLeaseInvite{
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

func TestGetAllProperties(t *testing.T) {
	c, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(database.MockGetAllPropertyByOwnerId(c, false)).ReturnsMany([]db.PropertyModel{property})

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/owner/properties/", nil)
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
	c, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/owner/properties/1/", nil)
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
	c, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	mock.Property.Expect(database.MockGetPropertyByID(c)).Errors(db.ErrNotFound)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/owner/properties/1/", nil)
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
	c, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/owner/properties/1/", nil)
	req.Header.Set("Oauth.claims.id", "2")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusForbidden, w.Code)
	var errorResponse utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
	require.NoError(t, err)
	assert.Equal(t, utils.PropertyNotYours, errorResponse.Code)
}

func TestGetPropertyInventory(t *testing.T) {
	c, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	propertyInv := BuildTestPropertyWithInventory("1")
	mock.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)
	mock.Property.Expect(database.MockGetPropertyInventory(c)).Returns(propertyInv)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/owner/properties/1/inventory/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	var propertyInventoryResponse models.PropertyInventoryResponse
	err := json.Unmarshal(w.Body.Bytes(), &propertyInventoryResponse)
	require.NoError(t, err)
	assert.Equal(t, propertyInv.ID, propertyInventoryResponse.ID)
}

func TestGetPropertyInventory_NotFound(t *testing.T) {
	c, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	mock.Property.Expect(database.MockGetPropertyByID(c)).Errors(db.ErrNotFound)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/owner/properties/1/inventory/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)
	var errorResponse utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
	require.NoError(t, err)
	assert.Equal(t, utils.PropertyNotFound, errorResponse.Code)
}

func TestGetPropertyInventory_NotYours(t *testing.T) {
	c, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/owner/properties/1/inventory/", nil)
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
	c, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(database.MockCreateProperty(c, property)).Returns(property)

	b, err := json.Marshal(property)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/v1/owner/properties/", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusCreated, w.Code)
	var resp models.IdResponse
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
	req, _ := http.NewRequest(http.MethodPost, "/v1/owner/properties/", bytes.NewReader(b))
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
	c, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(database.MockCreateProperty(c, property)).Errors(&protocol.UserFacingError{
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
	req, _ := http.NewRequest(http.MethodPost, "/v1/owner/properties/", bytes.NewReader(b))
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
	c, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	leaseInvite := BuildTestLeaseInvite()
	mock.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)
	mock.Lease.Expect(database.MockGetCurrentActiveLeaseByProperty(c)).Errors(db.ErrNotFound)
	mock.User.Expect(database.MockGetUserByEmail(c)).Errors(db.ErrNotFound)
	mock.LeaseInvite.Expect(database.MockCreateLeaseInvite(c, leaseInvite)).Returns(leaseInvite)

	reqBody := models.InviteRequest{
		TenantEmail: leaseInvite.TenantEmail,
		StartDate:   leaseInvite.StartDate,
		EndDate:     leaseInvite.InnerLeaseInvite.EndDate,
	}
	b, err := json.Marshal(reqBody)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/v1/owner/properties/1/send-invite/", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusCreated, w.Code)
	var resp db.InnerLeaseInvite
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.JSONEq(t, resp.ID, leaseInvite.ID)
}

func TestInviteTenant_MissingField(t *testing.T) {
	leaseInvite := BuildTestLeaseInvite()
	reqBody := models.InviteRequest{
		StartDate: leaseInvite.StartDate,
		EndDate:   leaseInvite.InnerLeaseInvite.EndDate,
	}
	b, err := json.Marshal(reqBody)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/v1/owner/properties/1/send-invite/", bytes.NewReader(b))
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
	c, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	mock.Property.Expect(database.MockGetPropertyByID(c)).Errors(db.ErrNotFound)

	leaseInvite := BuildTestLeaseInvite()
	reqBody := models.InviteRequest{
		TenantEmail: leaseInvite.TenantEmail,
		StartDate:   leaseInvite.StartDate,
		EndDate:     leaseInvite.InnerLeaseInvite.EndDate,
	}
	b, err := json.Marshal(reqBody)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/v1/owner/properties/1/send-invite/", bytes.NewReader(b))
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
	c, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)

	leaseInvite := BuildTestLeaseInvite()

	reqBody := models.InviteRequest{
		TenantEmail: leaseInvite.TenantEmail,
		StartDate:   leaseInvite.StartDate,
		EndDate:     leaseInvite.InnerLeaseInvite.EndDate,
	}
	b, err := json.Marshal(reqBody)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/v1/owner/properties/1/send-invite/", bytes.NewReader(b))
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
	c, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	lease := BuildTestLease("1")
	mock.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)
	mock.Lease.Expect(database.MockGetCurrentActiveLeaseByProperty(c)).ReturnsMany([]db.LeaseModel{lease})

	leaseInvite := BuildTestLeaseInvite()
	reqBody := models.InviteRequest{
		TenantEmail: leaseInvite.TenantEmail,
		StartDate:   leaseInvite.StartDate,
		EndDate:     leaseInvite.InnerLeaseInvite.EndDate,
	}
	b, err := json.Marshal(reqBody)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/v1/owner/properties/1/send-invite/", bytes.NewReader(b))
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
	c, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	leaseInvite := BuildTestLeaseInvite()
	mock.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)
	mock.Lease.Expect(database.MockGetCurrentActiveLeaseByProperty(c)).Errors(db.ErrNotFound)
	mock.User.Expect(database.MockGetUserByEmail(c)).Errors(db.ErrNotFound)
	mock.LeaseInvite.Expect(database.MockCreateLeaseInvite(c, leaseInvite)).Errors(&protocol.UserFacingError{
		IsPanic:   false,
		ErrorCode: "P2002", // https://www.prisma.io/docs/orm/reference/error-reference
		Meta: protocol.Meta{
			Target: []any{"tenant_email"},
		},
		Message: "Unique constraint failed",
	})

	reqBody := models.InviteRequest{
		TenantEmail: leaseInvite.TenantEmail,
		StartDate:   leaseInvite.StartDate,
		EndDate:     leaseInvite.InnerLeaseInvite.EndDate,
	}
	b, err := json.Marshal(reqBody)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/v1/owner/properties/1/send-invite/", bytes.NewReader(b))
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

func TestInviteTenant_AlreadyExistsAsOwner(t *testing.T) {
	c, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	leaseInvite := BuildTestLeaseInvite()
	owner := BuildTestUser("1")
	mock.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)
	mock.Lease.Expect(database.MockGetCurrentActiveLeaseByProperty(c)).Errors(db.ErrNotFound)
	mock.User.Expect(database.MockGetUserByEmail(c)).Returns(owner)

	reqBody := models.InviteRequest{
		TenantEmail: leaseInvite.TenantEmail,
		StartDate:   leaseInvite.StartDate,
		EndDate:     leaseInvite.InnerLeaseInvite.EndDate,
	}
	b, err := json.Marshal(reqBody)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/v1/owner/properties/1/send-invite/", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusConflict, w.Code)
	var resp utils.Error
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.UserAlreadyExistsAsOwner, resp.Code)
}

func TestInviteTenant_AlreadyHasLease(t *testing.T) {
	c, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	lease := BuildTestLease("1")
	leaseInvite := BuildTestLeaseInvite()
	user := BuildTestUser("1")
	user.Role = db.RoleTenant
	mock.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)
	mock.Lease.Expect(database.MockGetCurrentActiveLeaseByProperty(c)).Errors(db.ErrNotFound)
	mock.User.Expect(database.MockGetUserByEmail(c)).Returns(user)
	mock.Lease.Expect(database.MockGetCurrentActiveLeaseByTenant(c)).ReturnsMany([]db.LeaseModel{lease})

	reqBody := models.InviteRequest{
		TenantEmail: leaseInvite.TenantEmail,
		StartDate:   leaseInvite.StartDate,
		EndDate:     leaseInvite.InnerLeaseInvite.EndDate,
	}
	b, err := json.Marshal(reqBody)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/v1/owner/properties/1/send-invite/", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusConflict, w.Code)
	var resp utils.Error
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.TenantAlreadyHasLease, resp.Code)
}

func TestGetPropertyPicture(t *testing.T) {
	c, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	image := BuildTestImage("1", "data:image/jpeg;base64,b3Vp")
	mock.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)
	mock.Image.Expect(database.MockGetImageByID(c)).Returns(image)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/owner/properties/1/picture/", nil)
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
	c, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	property.InnerProperty.PictureID = nil
	mock.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/owner/properties/1/picture/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNoContent, w.Code)
}

func TestGetPropertyPicture_NotFound(t *testing.T) {
	c, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)
	mock.Image.Expect(database.MockGetImageByID(c)).Errors(db.ErrNotFound)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/owner/properties/1/picture/", nil)
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
	c, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	mock.Property.Expect(database.MockGetPropertyByID(c)).Errors(db.ErrNotFound)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/owner/properties/1/picture/", nil)
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
	c, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	image := BuildTestImage("1", "data:image/jpeg;base64,b3Vp")
	mock.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)
	mock.Image.Expect(database.MockCreateImage(c, image)).Returns(image)
	mock.Property.Expect(database.MockUpdatePropertyPicture(c)).Returns(property)

	reqBody := models.ImageRequest{
		Data: "data:image/jpeg;base64,b3Vp",
	}
	b, err := json.Marshal(reqBody)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/v1/owner/properties/1/picture/", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	var resp models.IdResponse
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, property.ID, resp.ID)
}

func TestUpdatePropertyPicture_MissingFields(t *testing.T) {
	c, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)

	reqBody := models.ImageRequest{}
	b, err := json.Marshal(reqBody)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/v1/owner/properties/1/picture/", bytes.NewReader(b))
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
	c, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)

	reqBody := models.ImageRequest{
		Data: "invalid_base64",
	}
	b, err := json.Marshal(reqBody)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/v1/owner/properties/1/picture/", bytes.NewReader(b))
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
	c, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	mock.Property.Expect(database.MockGetPropertyByID(c)).Errors(db.ErrNotFound)

	reqBody := models.ImageRequest{
		Data: "data:image/jpeg;base64,b3Vp",
	}
	b, err := json.Marshal(reqBody)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/v1/owner/properties/1/picture/", bytes.NewReader(b))
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
	c, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	image := BuildTestImage("1", "data:image/jpeg;base64,b3Vp")
	mock.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)
	mock.Image.Expect(database.MockCreateImage(c, image)).Returns(image)
	mock.Property.Expect(database.MockUpdatePropertyPicture(c)).Errors(db.ErrNotFound)

	reqBody := models.ImageRequest{
		Data: "data:image/jpeg;base64,b3Vp",
	}
	b, err := json.Marshal(reqBody)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/v1/owner/properties/1/picture/", bytes.NewReader(b))
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

func TestArchiveProperty(t *testing.T) {
	c, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	property.Leases()[0].Active = false
	updatedProperty := property
	updatedProperty.Archived = true
	mock.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)
	mock.Property.Expect(database.MockArchiveProperty(c)).Returns(updatedProperty)

	b, err := json.Marshal(map[string]bool{"archive": true})
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/v1/owner/properties/1/archive/", bytes.NewReader(b))
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	var resp models.IdResponse
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.JSONEq(t, resp.ID, property.ID)
}

func TestArchiveProperty_NonFree(t *testing.T) {
	c, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	updatedProperty := property
	updatedProperty.Archived = true
	mock.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)

	b, err := json.Marshal(map[string]bool{"archive": true})
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/v1/owner/properties/1/archive/", bytes.NewReader(b))
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusConflict, w.Code)
	var resp utils.Error
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.CannotArchiveNonFreeProperty, resp.Code)
}

func TestArchiveProperty_NotFound(t *testing.T) {
	c, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	mock.Property.Expect(database.MockGetPropertyByID(c)).Errors(db.ErrNotFound)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/v1/owner/properties/1/archive/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)
	var resp utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.PropertyNotFound, resp.Code)
}

func TestGetAllArchivedProperties(t *testing.T) {
	c, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	property.Archived = true
	mock.Property.Expect(database.MockGetAllPropertyByOwnerId(c, true)).ReturnsMany([]db.PropertyModel{property})

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/owner/properties/?archive=true", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	var resp []models.PropertyResponse
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.JSONEq(t, resp[0].ID, property.ID)
}

func TestUpdateProperty(t *testing.T) {
	c, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	updatedProperty := property
	updatedProperty.Name = "Updated Test"
	reqBody := models.PropertyUpdateRequest{
		Name:                &updatedProperty.Name,
		Address:             &updatedProperty.Address,
		ApartmentNumber:     updatedProperty.InnerProperty.ApartmentNumber,
		City:                &updatedProperty.City,
		PostalCode:          &updatedProperty.PostalCode,
		Country:             &updatedProperty.Country,
		AreaSqm:             &updatedProperty.AreaSqm,
		RentalPricePerMonth: &updatedProperty.RentalPricePerMonth,
		DepositPrice:        &updatedProperty.DepositPrice,
	}
	mock.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)
	mock.Property.Expect(database.MockUpdateProperty(c, reqBody)).Returns(updatedProperty)

	b, err := json.Marshal(reqBody)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/v1/owner/properties/1/", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	var resp models.IdResponse
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, updatedProperty.ID, resp.ID)
}

func TestUpdateProperty_NotFound(t *testing.T) {
	c, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	mock.Property.Expect(database.MockGetPropertyByID(c)).Errors(db.ErrNotFound)

	reqBody := models.PropertyUpdateRequest{
		Name: utils.Ptr("Updated Test"),
	}
	b, err := json.Marshal(reqBody)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/v1/owner/properties/1/", bytes.NewReader(b))
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

func TestUpdateProperty_NotYours(t *testing.T) {
	c, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)

	reqBody := models.PropertyUpdateRequest{
		Name: utils.Ptr("Updated Test"),
	}
	b, err := json.Marshal(reqBody)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/v1/owner/properties/1/", bytes.NewReader(b))
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

func TestCancelInvite(t *testing.T) {
	c, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)
	mock.LeaseInvite.Expect(database.MockGetCurrentLeaseInvite(c)).Returns(BuildTestLeaseInvite())
	mock.LeaseInvite.Expect(database.MockDeleteCurrentLeaseInvite(c)).Returns(BuildTestLeaseInvite())

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/v1/owner/properties/1/cancel-invite/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNoContent, w.Code)
}

func TestCancelInvite_PropertyNotYours(t *testing.T) {
	c, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/v1/owner/properties/1/cancel-invite/", nil)
	req.Header.Set("Oauth.claims.id", "wrong")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusForbidden, w.Code)
	var resp utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.PropertyNotYours, resp.Code)
}

func TestCancelInvite_NoLeaseInvite(t *testing.T) {
	c, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)
	mock.LeaseInvite.Expect(database.MockGetCurrentLeaseInvite(c)).Errors(db.ErrNotFound)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/v1/owner/properties/1/cancel-invite/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)
	var resp utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.NoLeaseInvite, resp.Code)
}
