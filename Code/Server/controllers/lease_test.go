package controllers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"immotep/backend/models"
	"immotep/backend/prisma/db"
	"immotep/backend/router"
	"immotep/backend/services"
	"immotep/backend/utils"
)

func BuildTestLease(id string) db.LeaseModel {
	return db.LeaseModel{
		InnerLease: db.InnerLease{
			ID:         id,
			PropertyID: "1",
			TenantID:   "1",
			Active:     true,
			CreatedAt:  time.Now(),
			StartDate:  time.Now(),
			EndDate:    utils.Ptr(time.Now().Add(time.Hour)),
		},
		RelationsLease: db.RelationsLease{
			Tenant: &db.UserModel{
				InnerUser: db.InnerUser{
					ID:        "1",
					Firstname: "John",
					Lastname:  "Doe",
					Email:     "johndoe@example.com",
				},
			},
			Property: &db.PropertyModel{
				InnerProperty: db.InnerProperty{
					ID:   "1",
					Name: "Test Property",
				},
				RelationsProperty: db.RelationsProperty{
					Owner: &db.UserModel{
						InnerUser: db.InnerUser{
							ID:        "1",
							Firstname: "John",
							Lastname:  "Doe",
							Email:     "johndoe@example.com",
						},
					},
				},
			},
		},
	}
}

func TestGetLeasesByProperty(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	lease1 := BuildTestLease("1")
	lease2 := BuildTestLease("2")

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

	mock.Lease.Expect(
		client.Client.Lease.FindMany(
			db.Lease.PropertyID.Equals("1"),
		).With(
			db.Lease.Tenant.Fetch(),
			db.Lease.Property.Fetch().With(db.Property.Owner.Fetch()),
		),
	).ReturnsMany([]db.LeaseModel{lease1, lease2})

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/owner/properties/1/leases/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	var resp []models.LeaseResponse
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Len(t, resp, 2)
	assert.Equal(t, lease1.ID, resp[0].ID)
	assert.Equal(t, lease2.ID, resp[1].ID)
}

func TestGetLeasesByProperty_NotFound(t *testing.T) {
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

	mock.Lease.Expect(
		client.Client.Lease.FindMany(
			db.Lease.PropertyID.Equals("1"),
		).With(
			db.Lease.Tenant.Fetch(),
			db.Lease.Property.Fetch().With(db.Property.Owner.Fetch()),
		),
	).ReturnsMany([]db.LeaseModel{})

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/owner/properties/1/leases/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	var resp []models.LeaseResponse
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Empty(t, resp)
}

func TestGetLeasesByProperty_PropertyNotFound(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	mock.Property.Expect(
		client.Client.Property.FindUnique(
			db.Property.ID.Equals("1"),
		).With(
			db.Property.Damages.Fetch(),
			db.Property.Leases.Fetch().With(db.Lease.Tenant.Fetch()),
			db.Property.LeaseInvite.Fetch(),
		),
	).Errors(db.ErrNotFound)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/owner/properties/1/leases/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)
}

func TestGetLeaseByID(t *testing.T) {
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

	lease := BuildTestLease("1")
	mock.Lease.Expect(
		client.Client.Lease.FindUnique(
			db.Lease.ID.Equals(lease.ID),
		).With(
			db.Lease.Tenant.Fetch(),
			db.Lease.Property.Fetch().With(db.Property.Owner.Fetch()),
		),
	).Returns(lease)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/owner/properties/1/leases/1/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	var resp models.LeaseResponse
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, lease.ID, resp.ID)
	assert.Equal(t, lease.Tenant().Name(), resp.TenantName)
}

func TestGetLeaseByID_NotFound(t *testing.T) {
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

	mock.Lease.Expect(
		client.Client.Lease.FindUnique(
			db.Lease.ID.Equals("1"),
		).With(
			db.Lease.Tenant.Fetch(),
			db.Lease.Property.Fetch().With(db.Property.Owner.Fetch()),
		),
	).Errors(db.ErrNotFound)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/owner/properties/1/leases/1/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)
}

func TestGetLeaseByID_CurrentActive(t *testing.T) {
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

	lease := BuildTestLease("1")
	mock.Lease.Expect(
		client.Client.Lease.FindMany(
			db.Lease.PropertyID.Equals("1"),
			db.Lease.Active.Equals(true),
		).With(
			db.Lease.Tenant.Fetch(),
			db.Lease.Property.Fetch().With(db.Property.Owner.Fetch()),
		),
	).ReturnsMany([]db.LeaseModel{lease})

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/owner/properties/1/leases/current/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	var resp models.LeaseResponse
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, lease.ID, resp.ID)
	assert.Equal(t, lease.Tenant().Name(), resp.TenantName)
}

func TestGetLeaseByID_CurrentActive_NotFound(t *testing.T) {
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

	mock.Lease.Expect(
		client.Client.Lease.FindMany(
			db.Lease.PropertyID.Equals("1"),
			db.Lease.Active.Equals(true),
		).With(
			db.Lease.Tenant.Fetch(),
			db.Lease.Property.Fetch().With(db.Property.Owner.Fetch()),
		),
	).ReturnsMany([]db.LeaseModel{})

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/owner/properties/1/leases/current/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)
}

func TestEndLease1(t *testing.T) {
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
			db.Lease.Property.Fetch().With(db.Property.Owner.Fetch()),
		),
	).ReturnsMany([]db.LeaseModel{lease})

	mock.Lease.Expect(
		client.Client.Lease.FindUnique(
			db.Lease.ID.Equals(lease.ID),
		).Update(
			db.Lease.Active.Set(false),
			db.Lease.EndDate.SetIfPresent(nil),
		),
	).Returns(lease)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/v1/owner/properties/1/leases/current/end/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNoContent, w.Code)
}

func TestEndLease2(t *testing.T) {
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
	lease.InnerLease.EndDate = nil
	mock.Lease.Expect(
		client.Client.Lease.FindMany(
			db.Lease.PropertyID.Equals(property.ID),
			db.Lease.Active.Equals(true),
		).With(
			db.Lease.Tenant.Fetch(),
			db.Lease.Property.Fetch().With(db.Property.Owner.Fetch()),
		),
	).ReturnsMany([]db.LeaseModel{lease})

	mock.Lease.Expect(
		client.Client.Lease.FindUnique(
			db.Lease.ID.Equals(lease.ID),
		).Update(
			db.Lease.Active.Set(false),
			db.Lease.EndDate.SetIfPresent(utils.Ptr(time.Now().Truncate(time.Minute))),
		),
	).Returns(lease)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/v1/owner/properties/1/leases/current/end/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNoContent, w.Code)
}

func TestEndLease_NotCurrent(t *testing.T) {
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
		client.Client.Lease.FindUnique(
			db.Lease.ID.Equals(lease.ID),
		).With(
			db.Lease.Tenant.Fetch(),
			db.Lease.Property.Fetch().With(db.Property.Owner.Fetch()),
		),
	).Returns(lease)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/v1/owner/properties/1/leases/1/end/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusBadRequest, w.Code)
	var resp utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.CannotEndNonCurrentLease, resp.Code)
}

func TestEndLease_NoActiveLease(t *testing.T) {
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
			db.Lease.Property.Fetch().With(db.Property.Owner.Fetch()),
		),
	).Errors(db.ErrNotFound)

	r := router.TestRoutes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/v1/owner/properties/1/leases/current/end/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)
	var resp utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.NoActiveLease, resp.Code)
}

func TestEndLease_PropertyNotFound(t *testing.T) {
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
	req, _ := http.NewRequest(http.MethodPut, "/v1/owner/properties/wrong/leases/current/end/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)
	var resp utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.PropertyNotFound, resp.Code)
}

func TestGetAllLeasesByTenant(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	lease1 := BuildTestLease("1")
	lease2 := BuildTestLease("2")

	mock.Lease.Expect(
		client.Client.Lease.FindMany(
			db.Lease.TenantID.Equals("1"),
		).With(
			db.Lease.Tenant.Fetch(),
			db.Lease.Property.Fetch().With(db.Property.Owner.Fetch()),
		),
	).ReturnsMany([]db.LeaseModel{lease1, lease2})

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/tenant/leases/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleTenant))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	var resp []models.LeaseResponse
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Len(t, resp, 2)
	assert.Equal(t, lease1.ID, resp[0].ID)
	assert.Equal(t, lease2.ID, resp[1].ID)
}

func TestGetAllLeasesByTenant_NotFound(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	mock.Lease.Expect(
		client.Client.Lease.FindMany(
			db.Lease.TenantID.Equals("1"),
		).With(
			db.Lease.Tenant.Fetch(),
			db.Lease.Property.Fetch().With(db.Property.Owner.Fetch()),
		),
	).ReturnsMany([]db.LeaseModel{})

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/tenant/leases/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleTenant))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	var resp []models.LeaseResponse
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Empty(t, resp)
}

func TestGetAllLeasesByTenant_Unauthorized(t *testing.T) {
	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/tenant/leases/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner)) // Unauthorized role
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusForbidden, w.Code)
	var resp utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.NotATenant, resp.Code)
}
