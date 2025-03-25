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
	"immotep/backend/models"
	"immotep/backend/prisma/db"
	"immotep/backend/router"
	"immotep/backend/services"
	"immotep/backend/utils"
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
			Damages:     []db.DamageModel{{}},
			Leases:      []db.LeaseModel{{}},
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
			Damages: []db.DamageModel{{}},
			Leases:  []db.LeaseModel{{}},
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
			TenantEmail: "test.test@example.com",
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
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(
		client.Client.Property.FindMany(
			db.Property.OwnerID.Equals("1"),
			db.Property.Archived.Equals(false),
		).With(
			db.Property.Damages.Fetch(),
			db.Property.Leases.Fetch().With(db.Lease.Tenant.Fetch()),
			db.Property.LeaseInvite.Fetch(),
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
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals(property.ID)).With(
			db.Property.Damages.Fetch(),
			db.Property.Leases.Fetch().With(db.Lease.Tenant.Fetch()),
			db.Property.LeaseInvite.Fetch(),
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
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals(property.ID)).With(
			db.Property.Damages.Fetch(),
			db.Property.Leases.Fetch().With(db.Lease.Tenant.Fetch()),
			db.Property.LeaseInvite.Fetch(),
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
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals(property.ID)).With(
			db.Property.Damages.Fetch(),
			db.Property.Leases.Fetch().With(db.Lease.Tenant.Fetch()),
			db.Property.LeaseInvite.Fetch(),
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

func TestGetPropertyInventory(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(
		client.Client.Property.FindUnique(
			db.Property.ID.Equals(property.ID),
		).With(
			db.Property.Damages.Fetch(),
			db.Property.Leases.Fetch().With(db.Lease.Tenant.Fetch()),
			db.Property.LeaseInvite.Fetch(),
		),
	).Returns(property)

	propertyInv := BuildTestPropertyWithInventory("1")
	mock.Property.Expect(
		client.Client.Property.FindUnique(
			db.Property.ID.Equals(propertyInv.ID),
		).With(
			db.Property.Damages.Fetch(),
			db.Property.Leases.Fetch().With(db.Lease.Tenant.Fetch()),
			db.Property.LeaseInvite.Fetch(),
			db.Property.Rooms.Fetch().With(db.Room.Furnitures.Fetch()),
		),
	).Returns(propertyInv)

	r := router.TestRoutes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/owner/properties/1/inventory/", nil)
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
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	mock.Property.Expect(
		client.Client.Property.FindUnique(
			db.Property.ID.Equals("wrong"),
		).With(
			db.Property.Damages.Fetch(),
			db.Property.Leases.Fetch().With(db.Lease.Tenant.Fetch()),
			db.Property.LeaseInvite.Fetch(),
		),
	).Errors(db.ErrNotFound)

	r := router.TestRoutes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/owner/properties/wrong/inventory/", nil)
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
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(
		client.Client.Property.FindUnique(
			db.Property.ID.Equals(property.ID),
		).With(
			db.Property.Damages.Fetch(),
			db.Property.Leases.Fetch().With(db.Lease.Tenant.Fetch()),
			db.Property.LeaseInvite.Fetch(),
		),
	).Returns(property)

	r := router.TestRoutes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/owner/properties/1/inventory/", nil)
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
	client, mock, ensure := services.ConnectDBTest()
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
			db.Property.ApartmentNumber.SetIfPresent(property.InnerProperty.ApartmentNumber),
		).With(
			db.Property.Damages.Fetch(),
			db.Property.Leases.Fetch(),
			db.Property.LeaseInvite.Fetch(),
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
	client, mock, ensure := services.ConnectDBTest()
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
			db.Property.ApartmentNumber.SetIfPresent(property.InnerProperty.ApartmentNumber),
		).With(
			db.Property.Damages.Fetch(),
			db.Property.Leases.Fetch(),
			db.Property.LeaseInvite.Fetch(),
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
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals(property.ID)).With(
			db.Property.Damages.Fetch(),
			db.Property.Leases.Fetch().With(db.Lease.Tenant.Fetch()),
			db.Property.LeaseInvite.Fetch(),
		),
	).Returns(property)

	mock.Lease.Expect(
		client.Client.Lease.FindMany(
			db.Lease.PropertyID.Equals(property.ID),
			db.Lease.Active.Equals(true),
		).With(
			db.Lease.Tenant.Fetch(),
		),
	).Errors(db.ErrNotFound)

	leaseInvite := BuildTestLeaseInvite()
	mock.User.Expect(
		client.Client.User.FindUnique(db.User.Email.Equals(leaseInvite.TenantEmail)),
	).Errors(db.ErrNotFound)

	mock.LeaseInvite.Expect(
		client.Client.LeaseInvite.CreateOne(
			db.LeaseInvite.TenantEmail.Set(leaseInvite.TenantEmail),
			db.LeaseInvite.StartDate.Set(leaseInvite.StartDate),
			db.LeaseInvite.Property.Link(db.Property.ID.Equals(property.ID)),
			db.LeaseInvite.EndDate.SetIfPresent(leaseInvite.InnerLeaseInvite.EndDate),
		).With(
			db.LeaseInvite.Property.Fetch().With(db.Property.Owner.Fetch()),
		),
	).Returns(leaseInvite)

	reqBody := models.InviteRequest{
		TenantEmail: leaseInvite.TenantEmail,
		StartDate:   leaseInvite.StartDate,
		EndDate:     leaseInvite.InnerLeaseInvite.EndDate,
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
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals("wrong")).With(
			db.Property.Damages.Fetch(),
			db.Property.Leases.Fetch().With(db.Lease.Tenant.Fetch()),
			db.Property.LeaseInvite.Fetch(),
		),
	).Errors(db.ErrNotFound)

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
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals(property.ID)).With(
			db.Property.Damages.Fetch(),
			db.Property.Leases.Fetch().With(db.Lease.Tenant.Fetch()),
			db.Property.LeaseInvite.Fetch(),
		),
	).Returns(property)

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
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals(property.ID)).With(
			db.Property.Damages.Fetch(),
			db.Property.Leases.Fetch().With(db.Lease.Tenant.Fetch()),
			db.Property.LeaseInvite.Fetch(),
		),
	).Returns(property)

	lease := BuildTestLease("1")
	mock.Lease.Expect(
		client.Client.Lease.FindMany(
			db.Lease.PropertyID.Equals(property.ID),
			db.Lease.Active.Equals(true),
		).With(
			db.Lease.Tenant.Fetch(),
		),
	).ReturnsMany([]db.LeaseModel{lease})

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
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals(property.ID)).With(
			db.Property.Damages.Fetch(),
			db.Property.Leases.Fetch().With(db.Lease.Tenant.Fetch()),
			db.Property.LeaseInvite.Fetch(),
		),
	).Returns(property)

	mock.Lease.Expect(
		client.Client.Lease.FindMany(
			db.Lease.PropertyID.Equals(property.ID),
			db.Lease.Active.Equals(true),
		).With(
			db.Lease.Tenant.Fetch(),
		),
	).Errors(db.ErrNotFound)

	leaseInvite := BuildTestLeaseInvite()
	mock.User.Expect(
		client.Client.User.FindUnique(db.User.Email.Equals(leaseInvite.TenantEmail)),
	).Errors(db.ErrNotFound)

	mock.LeaseInvite.Expect(
		client.Client.LeaseInvite.CreateOne(
			db.LeaseInvite.TenantEmail.Set(leaseInvite.TenantEmail),
			db.LeaseInvite.StartDate.Set(leaseInvite.StartDate),
			db.LeaseInvite.Property.Link(db.Property.ID.Equals(property.ID)),
			db.LeaseInvite.EndDate.SetIfPresent(leaseInvite.InnerLeaseInvite.EndDate),
		).With(
			db.LeaseInvite.Property.Fetch().With(db.Property.Owner.Fetch()),
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
		TenantEmail: leaseInvite.TenantEmail,
		StartDate:   leaseInvite.StartDate,
		EndDate:     leaseInvite.InnerLeaseInvite.EndDate,
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

func TestInviteTenant_AlreadyExistsAsOwner(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals(property.ID)).With(
			db.Property.Damages.Fetch(),
			db.Property.Leases.Fetch().With(db.Lease.Tenant.Fetch()),
			db.Property.LeaseInvite.Fetch(),
		),
	).Returns(property)

	mock.Lease.Expect(
		client.Client.Lease.FindMany(
			db.Lease.PropertyID.Equals(property.ID),
			db.Lease.Active.Equals(true),
		).With(
			db.Lease.Tenant.Fetch(),
		),
	).Errors(db.ErrNotFound)

	leaseInvite := BuildTestLeaseInvite()
	owner := BuildTestUser("1")
	mock.User.Expect(
		client.Client.User.FindUnique(db.User.Email.Equals(leaseInvite.TenantEmail)),
	).Returns(owner)

	reqBody := models.InviteRequest{
		TenantEmail: leaseInvite.TenantEmail,
		StartDate:   leaseInvite.StartDate,
		EndDate:     leaseInvite.InnerLeaseInvite.EndDate,
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
	assert.Equal(t, utils.UserAlreadyExistsAsOwner, resp.Code)
}

func TestInviteTenant_AlreadyHasLease(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals(property.ID)).With(
			db.Property.Damages.Fetch(),
			db.Property.Leases.Fetch().With(db.Lease.Tenant.Fetch()),
			db.Property.LeaseInvite.Fetch(),
		),
	).Returns(property)

	lease := BuildTestLease("1")
	mock.Lease.Expect(
		client.Client.Lease.FindMany(
			db.Lease.PropertyID.Equals(property.ID),
			db.Lease.Active.Equals(true),
		).With(
			db.Lease.Tenant.Fetch(),
		),
	).Errors(db.ErrNotFound)

	leaseInvite := BuildTestLeaseInvite()
	user := BuildTestUser("1")
	user.Role = db.RoleTenant
	mock.User.Expect(
		client.Client.User.FindUnique(db.User.Email.Equals(leaseInvite.TenantEmail)),
	).Returns(user)

	mock.Lease.Expect(
		client.Client.Lease.FindMany(
			db.Lease.TenantID.Equals("1"),
			db.Lease.Active.Equals(true),
		).With(
			db.Lease.Tenant.Fetch(),
		),
	).ReturnsMany([]db.LeaseModel{lease})

	reqBody := models.InviteRequest{
		TenantEmail: leaseInvite.TenantEmail,
		StartDate:   leaseInvite.StartDate,
		EndDate:     leaseInvite.InnerLeaseInvite.EndDate,
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
	assert.Equal(t, utils.TenantAlreadyHasLease, resp.Code)
}

func TestGetPropertyPicture(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals(property.ID)).With(
			db.Property.Damages.Fetch(),
			db.Property.Leases.Fetch().With(db.Lease.Tenant.Fetch()),
			db.Property.LeaseInvite.Fetch(),
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
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	property.InnerProperty.PictureID = nil
	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals(property.ID)).With(
			db.Property.Damages.Fetch(),
			db.Property.Leases.Fetch().With(db.Lease.Tenant.Fetch()),
			db.Property.LeaseInvite.Fetch(),
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
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals(property.ID)).With(
			db.Property.Damages.Fetch(),
			db.Property.Leases.Fetch().With(db.Lease.Tenant.Fetch()),
			db.Property.LeaseInvite.Fetch(),
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
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals("wrong")).With(
			db.Property.Damages.Fetch(),
			db.Property.Leases.Fetch().With(db.Lease.Tenant.Fetch()),
			db.Property.LeaseInvite.Fetch(),
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
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals(property.ID)).With(
			db.Property.Damages.Fetch(),
			db.Property.Leases.Fetch().With(db.Lease.Tenant.Fetch()),
			db.Property.LeaseInvite.Fetch(),
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
			db.Property.Leases.Fetch().With(db.Lease.Tenant.Fetch()),
			db.Property.LeaseInvite.Fetch(),
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
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals(property.ID)).With(
			db.Property.Damages.Fetch(),
			db.Property.Leases.Fetch().With(db.Lease.Tenant.Fetch()),
			db.Property.LeaseInvite.Fetch(),
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
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals(property.ID)).With(
			db.Property.Damages.Fetch(),
			db.Property.Leases.Fetch().With(db.Lease.Tenant.Fetch()),
			db.Property.LeaseInvite.Fetch(),
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
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals("wrong")).With(
			db.Property.Damages.Fetch(),
			db.Property.Leases.Fetch().With(db.Lease.Tenant.Fetch()),
			db.Property.LeaseInvite.Fetch(),
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
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals(property.ID)).With(
			db.Property.Damages.Fetch(),
			db.Property.Leases.Fetch().With(db.Lease.Tenant.Fetch()),
			db.Property.LeaseInvite.Fetch(),
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
			db.Property.Leases.Fetch().With(db.Lease.Tenant.Fetch()),
			db.Property.LeaseInvite.Fetch(),
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

func TestArchiveProperty(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(
		client.Client.Property.FindUnique(
			db.Property.ID.Equals(property.ID),
		).With(
			db.Property.Damages.Fetch(),
			db.Property.Leases.Fetch().With(db.Lease.Tenant.Fetch()),
			db.Property.LeaseInvite.Fetch(),
		),
	).Returns(property)

	updatedProperty := property
	updatedProperty.Archived = true
	mock.Property.Expect(
		client.Client.Property.FindUnique(
			db.Property.ID.Equals(property.ID),
		).With(
			db.Property.Damages.Fetch(),
			db.Property.Leases.Fetch().With(db.Lease.Tenant.Fetch()),
			db.Property.LeaseInvite.Fetch(),
		).Update(
			db.Property.Archived.Set(true),
		),
	).Returns(updatedProperty)

	b, err := json.Marshal(map[string]bool{"archive": true})
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/owner/properties/1/archive/", bytes.NewReader(b))
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	var resp models.PropertyResponse
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.JSONEq(t, resp.ID, property.ID)
	assert.True(t, resp.Archived)
}

func TestArchiveProperty_NotFound(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(
		client.Client.Property.FindUnique(
			db.Property.ID.Equals(property.ID),
		).With(
			db.Property.Damages.Fetch(),
			db.Property.Leases.Fetch().With(db.Lease.Tenant.Fetch()),
			db.Property.LeaseInvite.Fetch(),
		),
	).Errors(db.ErrNotFound)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/owner/properties/1/archive/", nil)
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
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	property.Archived = true
	mock.Property.Expect(
		client.Client.Property.FindMany(
			db.Property.OwnerID.Equals("1"),
			db.Property.Archived.Equals(true),
		).With(
			db.Property.Damages.Fetch(),
			db.Property.Leases.Fetch().With(db.Lease.Tenant.Fetch()),
			db.Property.LeaseInvite.Fetch(),
		),
	).ReturnsMany([]db.PropertyModel{property})

	r := router.TestRoutes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/owner/properties/archived/", nil)
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
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals(property.ID)).With(
			db.Property.Damages.Fetch(),
			db.Property.Leases.Fetch().With(db.Lease.Tenant.Fetch()),
			db.Property.LeaseInvite.Fetch(),
		),
	).Returns(property)

	updatedProperty := property
	updatedProperty.Name = "Updated Test"
	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals(property.ID)).With(
			db.Property.Damages.Fetch(),
			db.Property.Leases.Fetch().With(db.Lease.Tenant.Fetch()),
			db.Property.LeaseInvite.Fetch(),
		).Update(
			db.Property.Name.SetIfPresent(&updatedProperty.Name),
			db.Property.Address.SetIfPresent(&updatedProperty.Address),
			db.Property.ApartmentNumber.SetIfPresent(updatedProperty.InnerProperty.ApartmentNumber),
			db.Property.City.SetIfPresent(&updatedProperty.City),
			db.Property.PostalCode.SetIfPresent(&updatedProperty.PostalCode),
			db.Property.Country.SetIfPresent(&updatedProperty.Country),
			db.Property.AreaSqm.SetIfPresent(&updatedProperty.AreaSqm),
			db.Property.RentalPricePerMonth.SetIfPresent(&updatedProperty.RentalPricePerMonth),
			db.Property.DepositPrice.SetIfPresent(&updatedProperty.DepositPrice),
		),
	).Returns(updatedProperty)

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
	b, err := json.Marshal(reqBody)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/owner/properties/1/", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	var resp models.PropertyResponse
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, updatedProperty.ID, resp.ID)
	assert.Equal(t, updatedProperty.Name, resp.Name)
}

func TestUpdateProperty_NotFound(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals("wrong")).With(
			db.Property.Damages.Fetch(),
			db.Property.Leases.Fetch().With(db.Lease.Tenant.Fetch()),
			db.Property.LeaseInvite.Fetch(),
		),
	).Errors(db.ErrNotFound)

	reqBody := models.PropertyUpdateRequest{
		Name: utils.Ptr("Updated Test"),
	}
	b, err := json.Marshal(reqBody)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/owner/properties/wrong/", bytes.NewReader(b))
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
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals(property.ID)).With(
			db.Property.Damages.Fetch(),
			db.Property.Leases.Fetch().With(db.Lease.Tenant.Fetch()),
			db.Property.LeaseInvite.Fetch(),
		),
	).Returns(property)

	reqBody := models.PropertyUpdateRequest{
		Name: utils.Ptr("Updated Test"),
	}
	b, err := json.Marshal(reqBody)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/owner/properties/1/", bytes.NewReader(b))
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
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals(property.ID)).With(
			db.Property.Damages.Fetch(),
			db.Property.Leases.Fetch().With(db.Lease.Tenant.Fetch()),
			db.Property.LeaseInvite.Fetch(),
		),
	).Returns(property)

	mock.LeaseInvite.Expect(
		client.Client.LeaseInvite.FindUnique(db.LeaseInvite.PropertyID.Equals(property.ID)),
	).Returns(BuildTestLeaseInvite())

	mock.LeaseInvite.Expect(
		client.Client.LeaseInvite.FindUnique(db.LeaseInvite.PropertyID.Equals(property.ID)).Delete(),
	).Returns(BuildTestLeaseInvite())

	r := router.TestRoutes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/owner/properties/1/cancel-invite/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNoContent, w.Code)
}

func TestCancelInvite_PropertyNotYours(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals(property.ID)).With(
			db.Property.Damages.Fetch(),
			db.Property.Leases.Fetch().With(db.Lease.Tenant.Fetch()),
			db.Property.LeaseInvite.Fetch(),
		),
	).Returns(property)

	r := router.TestRoutes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/owner/properties/1/cancel-invite/", nil)
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
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals(property.ID)).With(
			db.Property.Damages.Fetch(),
			db.Property.Leases.Fetch().With(db.Lease.Tenant.Fetch()),
			db.Property.LeaseInvite.Fetch(),
		),
	).Returns(property)

	mock.LeaseInvite.Expect(
		client.Client.LeaseInvite.FindUnique(db.LeaseInvite.PropertyID.Equals(property.ID)),
	).Errors(db.ErrNotFound)

	r := router.TestRoutes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/owner/properties/1/cancel-invite/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)
	var resp utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.NoLeaseInvite, resp.Code)
}
