package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/steebchen/prisma-client-go/engine/protocol"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"immotep/backend/models"
	"immotep/backend/prisma/db"
	"immotep/backend/router"
	"immotep/backend/services"
	"immotep/backend/services/database"
	"immotep/backend/utils"
)

func BuildTestDamage(id string) db.DamageModel {
	return db.DamageModel{
		InnerDamage: db.InnerDamage{
			ID:           id,
			LeaseID:      "1",
			RoomID:       "1",
			Comment:      "Test Comment",
			Priority:     db.PriorityHigh,
			Read:         true,
			FixedOwner:   false,
			FixedTenant:  false,
			FixPlannedAt: nil,
			FixedAt:      nil,
		},
		RelationsDamage: db.RelationsDamage{
			Lease: &db.LeaseModel{
				RelationsLease: db.RelationsLease{
					Tenant: &db.UserModel{
						InnerUser: db.InnerUser{
							Firstname: "John",
							Lastname:  "Doe",
						},
					},
				},
			},
			Room: &db.RoomModel{
				InnerRoom: db.InnerRoom{
					Name: "Living Room",
				},
			},
			Pictures: []db.ImageModel{
				{
					InnerImage: db.InnerImage{
						ID:   "1",
						Data: db.Bytes("base64image1"),
					},
				},
			},
		},
	}
}

func TestCreateDamage(t *testing.T) {
	c, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	lease := BuildTestLease("1")
	room := BuildTestRoom("1", "1")
	damage := BuildTestDamage("1")
	image := BuildTestImage("1", "b3Vp")
	mock.Lease.Expect(database.MockGetLeaseByID(c)).Returns(lease)
	mock.Room.Expect(database.MockGetRoomByID(c)).Returns(room)
	mock.Image.Expect(database.MockCreateImage(c, image)).Returns(image)
	mock.Damage.Expect(database.MockCreateDamage(c, damage, "1", []string{"1"})).Returns(damage)

	reqBody := models.DamageRequest{
		RoomID:   damage.RoomID,
		Comment:  damage.Comment,
		Priority: damage.Priority,
		Pictures: []string{"b3Vp"},
	}
	b, err := json.Marshal(reqBody)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/v1/tenant/leases/1/damages/", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleTenant))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusCreated, w.Code)
	var resp models.PropertyResponse
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.JSONEq(t, resp.ID, damage.ID)
}

func TestCreateDamage_MissingFields(t *testing.T) {
	c, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	lease := BuildTestLease("1")
	mock.Lease.Expect(database.MockGetLeaseByID(c)).Returns(lease)

	reqBody := models.DamageRequest{
		RoomID:   "1",
		Comment:  "Test Comment",
		Pictures: []string{"b3Vp"},
	}
	b, err := json.Marshal(reqBody)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/v1/tenant/leases/1/damages/", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleTenant))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusBadRequest, w.Code)
	var resp utils.Error
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.MissingFields, resp.Code)
}

func TestCreateDamage_RoomNotFound(t *testing.T) {
	c, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	lease := BuildTestLease("1")
	mock.Lease.Expect(database.MockGetLeaseByID(c)).Returns(lease)
	mock.Room.Expect(database.MockGetRoomByID(c)).Errors(db.ErrNotFound)

	reqBody := models.DamageRequest{
		RoomID:   "1",
		Comment:  "Test Comment",
		Priority: db.PriorityHigh,
		Pictures: []string{"b3Vp"},
	}
	b, err := json.Marshal(reqBody)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/v1/tenant/leases/1/damages/", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleTenant))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)
	var resp utils.Error
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.RoomNotFound, resp.Code)
}

func TestGetDamagesByProperty(t *testing.T) {
	c, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	damages := []db.DamageModel{
		BuildTestDamage("1"),
		BuildTestDamage("2"),
	}
	mock.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)
	mock.Damage.Expect(database.MockGetDamagesByPropertyID(c, false)).ReturnsMany(damages)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/owner/properties/1/damages/", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	var resp []models.DamageResponse
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Len(t, resp, 2)
	assert.Equal(t, damages[0].ID, resp[0].ID)
	assert.Equal(t, damages[1].ID, resp[1].ID)
}

func TestGetDamagesByProperty_PropertyNotYours(t *testing.T) {
	c, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/owner/properties/1/damages/", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "2") // Simulate a different user
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusForbidden, w.Code)
	var resp utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.PropertyNotYours, resp.Code)
}

func TestGetFixedDamagesByProperty(t *testing.T) {
	c, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	damages := []db.DamageModel{
		BuildTestDamage("1"),
		BuildTestDamage("2"),
	}
	mock.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)
	mock.Damage.Expect(database.MockGetDamagesByPropertyID(c, true)).ReturnsMany(damages)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/owner/properties/1/damages/?fixed=true", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	var resp []models.DamageResponse
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Len(t, resp, 2)
	assert.Equal(t, damages[0].ID, resp[0].ID)
	assert.Equal(t, damages[1].ID, resp[1].ID)
}

func TestGetFixedDamagesByProperty_PropertyNotYours(t *testing.T) {
	c, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/owner/properties/1/damages/?fixed=true", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "2") // Simulate a different user
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusForbidden, w.Code)
	var resp utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.PropertyNotYours, resp.Code)
}

func TestGetDamagesByLease(t *testing.T) {
	c, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	lease := BuildTestLease("1")
	damages := []db.DamageModel{
		BuildTestDamage("1"),
		BuildTestDamage("2"),
	}
	mock.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)
	mock.Lease.Expect(database.MockGetLeaseByID(c)).Returns(lease)
	mock.Damage.Expect(database.MockGetDamagesByLeaseID(c, false)).ReturnsMany(damages)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/owner/properties/1/leases/1/damages/", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	var resp []models.DamageResponse
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Len(t, resp, 2)
	assert.Equal(t, damages[0].ID, resp[0].ID)
	assert.Equal(t, damages[1].ID, resp[1].ID)
}

func TestGetDamagesByLease_LeaseNotFound(t *testing.T) {
	c, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)
	mock.Lease.Expect(database.MockGetLeaseByID(c)).Errors(db.ErrNotFound)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/owner/properties/1/leases/1/damages/", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)
	var resp utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.LeaseNotFound, resp.Code)
}

func TestGetDamagesByLease_NoActiveLease(t *testing.T) {
	c, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)
	mock.Lease.Expect(database.MockGetCurrentActiveLeaseByProperty(c)).Errors(db.ErrNotFound)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/owner/properties/1/leases/current/damages/", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)
	var resp utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.NoActiveLease, resp.Code)
}

func TestGetFixedDamagesByLease(t *testing.T) {
	c, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	lease := BuildTestLease("1")
	damages := []db.DamageModel{
		BuildTestDamage("1"),
		BuildTestDamage("2"),
	}
	mock.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)
	mock.Lease.Expect(database.MockGetLeaseByID(c)).Returns(lease)
	mock.Damage.Expect(database.MockGetDamagesByLeaseID(c, true)).ReturnsMany(damages)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/owner/properties/1/leases/1/damages/?fixed=true", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	var resp []models.DamageResponse
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Len(t, resp, 2)
	assert.Equal(t, damages[0].ID, resp[0].ID)
	assert.Equal(t, damages[1].ID, resp[1].ID)
}

func TestGetFixedDamagesByLease_LeaseNotFound(t *testing.T) {
	c, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)
	mock.Lease.Expect(database.MockGetLeaseByID(c)).Errors(db.ErrNotFound)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/owner/properties/1/leases/1/damages/?fixed=true", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)
	var resp utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.LeaseNotFound, resp.Code)
}

func TestGetFixedDamagesByLease_NoActiveLease(t *testing.T) {
	c, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)
	mock.Lease.Expect(database.MockGetCurrentActiveLeaseByProperty(c)).Errors(db.ErrNotFound)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/owner/properties/1/leases/current/damages/?fixed=true", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)
	var resp utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.NoActiveLease, resp.Code)
}

func TestGetDamage(t *testing.T) {
	c, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	lease := BuildTestLease("1")
	damage := BuildTestDamage("1")
	mock.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)
	mock.Lease.Expect(database.MockGetLeaseByID(c)).Returns(lease)
	mock.Damage.Expect(database.MockGetDamageByID(c)).Returns(damage)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/owner/properties/1/leases/1/damages/1/", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	var resp models.DamageResponse
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, damage.ID, resp.ID)
	assert.Equal(t, damage.Comment, resp.Comment)
}

func TestGetDamage_NotFound(t *testing.T) {
	c, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	lease := BuildTestLease("1")
	mock.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)
	mock.Lease.Expect(database.MockGetLeaseByID(c)).Returns(lease)
	mock.Damage.Expect(database.MockGetDamageByID(c)).Errors(db.ErrNotFound)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/owner/properties/1/leases/1/damages/1/", nil) // Non-existent damage ID
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)
	var resp utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.DamageNotFound, resp.Code)
}

func TestUpdateDamageOwner(t *testing.T) {
	c, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	lease := BuildTestLease("1")
	damage := BuildTestDamage("1")
	mock.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)
	mock.Lease.Expect(database.MockGetLeaseByID(c)).Returns(lease)
	mock.Damage.Expect(database.MockGetDamageByID(c)).Returns(damage)
	mock.Damage.Expect(database.MockUpdateDamageOwner(c, models.DamageOwnerUpdateRequest{
		Read:         utils.Ptr(true),
		FixPlannedAt: nil,
	})).Returns(damage)

	reqBody := models.DamageOwnerUpdateRequest{
		Read: utils.Ptr(true),
	}
	b, err := json.Marshal(reqBody)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/v1/owner/properties/1/leases/1/damages/1/", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	var resp models.DamageCreateResponse
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, damage.ID, resp.ID)
}

func TestUpdateDamageOwner_DamageNotFound(t *testing.T) {
	c, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	lease := BuildTestLease("1")
	mock.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)
	mock.Lease.Expect(database.MockGetLeaseByID(c)).Returns(lease)
	mock.Damage.Expect(database.MockGetDamageByID(c)).Errors(db.ErrNotFound)

	reqBody := models.DamageOwnerUpdateRequest{
		Read: utils.Ptr(true),
	}
	b, err := json.Marshal(reqBody)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/v1/owner/properties/1/leases/1/damages/1/", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)
	var resp utils.Error
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.DamageNotFound, resp.Code)
}

func TestUpdateDamageOwner_AlreadyFixed(t *testing.T) {
	c, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	lease := BuildTestLease("1")
	damage := BuildTestDamage("1")
	damage.FixedOwner = true
	damage.FixedTenant = true
	mock.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)
	mock.Lease.Expect(database.MockGetLeaseByID(c)).Returns(lease)
	mock.Damage.Expect(database.MockGetDamageByID(c)).Returns(damage)

	reqBody := models.DamageOwnerUpdateRequest{
		Read: utils.Ptr(true),
	}
	b, err := json.Marshal(reqBody)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/v1/owner/properties/1/leases/1/damages/1/", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusBadRequest, w.Code)
	var resp utils.Error
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.CannotUpdateFixedDamage, resp.Code)
}

func TestUpdateDamageOwner_CannotUpdateFixedDamage(t *testing.T) {
	c, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	lease := BuildTestLease("1")
	damage := BuildTestDamage("1")
	damage.FixedOwner = true
	mock.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)
	mock.Lease.Expect(database.MockGetLeaseByID(c)).Returns(lease)
	mock.Damage.Expect(database.MockGetDamageByID(c)).Returns(damage)
	mock.Damage.Expect(database.MockUpdateDamageOwner(c, models.DamageOwnerUpdateRequest{
		Read:         utils.Ptr(true),
		FixPlannedAt: nil,
	})).Errors(&protocol.UserFacingError{
		IsPanic:   false,
		ErrorCode: "P2002", // https://www.prisma.io/docs/orm/reference/error-reference#p2002
		Meta: protocol.Meta{
			Target: []any{"name"},
		},
		Message: "Unique constraint failed",
	})

	reqBody := models.DamageOwnerUpdateRequest{
		Read: utils.Ptr(true),
	}
	b, err := json.Marshal(reqBody)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/v1/owner/properties/1/leases/1/damages/1/", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusConflict, w.Code)
	var resp utils.Error
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.DamageAlreadyExists, resp.Code)
}

func TestUpdateDamageTenant(t *testing.T) {
	c, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	lease := BuildTestLease("1")
	damage := BuildTestDamage("1")
	image := BuildTestImage("1", "b3Vp")
	mock.Lease.Expect(database.MockGetLeaseByID(c)).Returns(lease)
	mock.Damage.Expect(database.MockGetDamageByID(c)).Returns(damage)
	mock.Image.Expect(database.MockCreateImage(c, image)).Returns(image)
	mock.Damage.Expect(database.MockUpdateDamageTenant(c, models.DamageTenantUpdateRequest{
		Comment:     utils.Ptr("Updated Comment"),
		Priority:    utils.Ptr(db.PriorityLow),
		AddPictures: []string{"1"},
	}, []string{"1"})).Returns(damage)

	reqBody := models.DamageTenantUpdateRequest{
		Comment:     utils.Ptr("Updated Comment"),
		Priority:    utils.Ptr(db.PriorityLow),
		AddPictures: []string{"b3Vp"},
	}
	b, err := json.Marshal(reqBody)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/v1/tenant/leases/1/damages/1/", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleTenant))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	var resp models.DamageCreateResponse
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, damage.ID, resp.ID)
}

func TestUpdateDamageTenant_MissingFields(t *testing.T) {
	c, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	lease := BuildTestLease("1")
	damage := BuildTestDamage("1")
	mock.Lease.Expect(database.MockGetLeaseByID(c)).Returns(lease)
	mock.Damage.Expect(database.MockGetDamageByID(c)).Returns(damage)

	reqBody := models.DamageTenantUpdateRequest{
		AddPictures: []string{""},
	}
	b, err := json.Marshal(reqBody)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/v1/tenant/leases/1/damages/1/", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleTenant))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusBadRequest, w.Code)
	var resp utils.Error
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.MissingFields, resp.Code)
}

func TestUpdateDamageTenant_DamageNotFound(t *testing.T) {
	c, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	lease := BuildTestLease("1")
	mock.Lease.Expect(database.MockGetLeaseByID(c)).Returns(lease)
	mock.Damage.Expect(database.MockGetDamageByID(c)).Errors(db.ErrNotFound)

	reqBody := models.DamageTenantUpdateRequest{
		Comment:     utils.Ptr("Updated Comment"),
		Priority:    utils.Ptr(db.PriorityLow),
		AddPictures: []string{"newPictureBase64"},
	}
	b, err := json.Marshal(reqBody)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/v1/tenant/leases/1/damages/1/", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleTenant))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)
	var resp utils.Error
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.DamageNotFound, resp.Code)
}

func TestUpdateDamageTenant_AlreadyFixed(t *testing.T) {
	c, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	lease := BuildTestLease("1")
	damage := BuildTestDamage("1")
	damage.FixedOwner = true
	damage.FixedTenant = true
	mock.Lease.Expect(database.MockGetLeaseByID(c)).Returns(lease)
	mock.Damage.Expect(database.MockGetDamageByID(c)).Returns(damage)

	reqBody := models.DamageTenantUpdateRequest{
		Comment:  utils.Ptr("Updated Comment"),
		Priority: utils.Ptr(db.PriorityLow),
	}
	b, err := json.Marshal(reqBody)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/v1/tenant/leases/1/damages/1/", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleTenant))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusBadRequest, w.Code)
	var resp utils.Error
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.CannotUpdateFixedDamage, resp.Code)
}

func TestUpdateDamageTenant_AlreadyExists(t *testing.T) {
	c, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	lease := BuildTestLease("1")
	damage := BuildTestDamage("1")
	damage.FixedTenant = true
	image := BuildTestImage("1", "b3Vp")
	mock.Lease.Expect(database.MockGetLeaseByID(c)).Returns(lease)
	mock.Damage.Expect(database.MockGetDamageByID(c)).Returns(damage)
	mock.Image.Expect(database.MockCreateImage(c, image)).Returns(image)
	mock.Damage.Expect(database.MockUpdateDamageTenant(c, models.DamageTenantUpdateRequest{
		Comment:     utils.Ptr("Updated Comment"),
		Priority:    utils.Ptr(db.PriorityLow),
		AddPictures: []string{"1"},
	}, []string{"1"})).Errors(&protocol.UserFacingError{
		IsPanic:   false,
		ErrorCode: "P2002", // https://www.prisma.io/docs/orm/reference/error-reference#p2002
		Meta: protocol.Meta{
			Target: []any{"name"},
		},
		Message: "Unique constraint failed",
	})

	reqBody := models.DamageTenantUpdateRequest{
		Comment:     utils.Ptr("Updated Comment"),
		Priority:    utils.Ptr(db.PriorityLow),
		AddPictures: []string{"b3Vp"},
	}
	b, err := json.Marshal(reqBody)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/v1/tenant/leases/1/damages/1/", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleTenant))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusConflict, w.Code)
	var resp utils.Error
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.DamageAlreadyExists, resp.Code)
}

func TestFixDamageOwner(t *testing.T) {
	c, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	lease := BuildTestLease("1")
	damage := BuildTestDamage("1")
	mock.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)
	mock.Lease.Expect(database.MockGetLeaseByID(c)).Returns(lease)
	mock.Damage.Expect(database.MockGetDamageByID(c)).Returns(damage)
	mock.Damage.Expect(database.MockMarkDamageAsFixed(c, damage, db.RoleOwner)).Returns(damage)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/v1/owner/properties/1/leases/1/damages/1/fix/", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	var resp models.DamageResponse
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, damage.ID, resp.ID)
}

func TestFixDamageTenant(t *testing.T) {
	c, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	lease := BuildTestLease("1")
	damage := BuildTestDamage("1")
	mock.Lease.Expect(database.MockGetLeaseByID(c)).Returns(lease)
	mock.Damage.Expect(database.MockGetDamageByID(c)).Returns(damage)
	mock.Damage.Expect(database.MockMarkDamageAsFixed(c, damage, db.RoleTenant)).Returns(damage)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/v1/tenant/leases/1/damages/1/fix/", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleTenant))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	var resp models.DamageResponse
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, damage.ID, resp.ID)
}

func TestFixDamage_AlreadyFixed(t *testing.T) {
	c, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	lease := BuildTestLease("1")
	damage := BuildTestDamage("1")
	damage.FixedOwner = true
	damage.FixedTenant = true
	mock.Lease.Expect(database.MockGetLeaseByID(c)).Returns(lease)
	mock.Damage.Expect(database.MockGetDamageByID(c)).Returns(damage)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/v1/tenant/leases/1/damages/1/fix/", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleTenant))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusBadRequest, w.Code)
	var resp utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.DamageAlreadyFixed, resp.Code)
}

func TestFixDamage_DamageNotFound(t *testing.T) {
	c, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	lease := BuildTestLease("1")
	mock.Lease.Expect(database.MockGetLeaseByID(c)).Returns(lease)
	mock.Damage.Expect(database.MockGetDamageByID(c)).Errors(db.ErrNotFound)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/v1/tenant/leases/1/damages/1/fix/", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleTenant))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)
	var resp utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.DamageNotFound, resp.Code)
}
