package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/maxzerbini/oauth"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"immotep/backend/controllers"
	"immotep/backend/prisma/db"
	"immotep/backend/router"
	"immotep/backend/services"
	"immotep/backend/utils"
)

const TENANT_EMAIL = "test1@example.com"

func TestTokenAuth(t *testing.T) {
	bServer := oauth.NewOAuthBearerServer(
		"1234567890",
		time.Hour*24,
		&router.TestUserVerifier{},
		nil)

	f := controllers.TokenAuth(bServer)
	assert.NotNil(t, f)
	var expected func(*gin.Context)
	assert.IsType(t, expected, f)
}

func TestRegisterOwnerMissingFields(t *testing.T) {
	user := BuildTestUser("1")
	user.Email = ""
	b, err := json.Marshal(user)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/v1/auth/register/", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	var errorResponse utils.Error
	err = json.Unmarshal(w.Body.Bytes(), &errorResponse)
	require.NoError(t, err)
	assert.Equal(t, utils.MissingFields, errorResponse.Code)
}

func TestRegisterOwnerWrongPassword(t *testing.T) {
	user := BuildTestUser("1")
	user.Password = "1234"
	b, err := json.Marshal(user)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/v1/auth/register/", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	var errorResponse utils.Error
	err = json.Unmarshal(w.Body.Bytes(), &errorResponse)
	require.NoError(t, err)
	assert.Equal(t, utils.MissingFields, errorResponse.Code)
}

func TestRegisterTenantMissingFields(t *testing.T) {
	user := BuildTestUser("1")
	user.Email = ""
	b, err := json.Marshal(user)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/v1/auth/invite/1/", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	var errorResponse utils.Error
	err = json.Unmarshal(w.Body.Bytes(), &errorResponse)
	require.NoError(t, err)
	assert.Equal(t, utils.MissingFields, errorResponse.Code)
}

func TestRegisterTenantWrongPassword(t *testing.T) {
	user := BuildTestUser("1")
	user.Password = "1234"
	b, err := json.Marshal(user)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/v1/auth/invite/1/", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	var errorResponse utils.Error
	err = json.Unmarshal(w.Body.Bytes(), &errorResponse)
	require.NoError(t, err)
	assert.Equal(t, utils.MissingFields, errorResponse.Code)
}

func TestRegisterTenantInviteNotFound(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	mock.LeaseInvite.Expect(
		client.Client.LeaseInvite.FindUnique(db.LeaseInvite.ID.Equals("wrong")),
	).Errors(db.ErrNotFound)

	user := BuildTestUser("1")
	b, err := json.Marshal(user)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/v1/auth/invite/wrong/", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	var errorResponse utils.Error
	err = json.Unmarshal(w.Body.Bytes(), &errorResponse)
	require.NoError(t, err)
	assert.Equal(t, utils.InviteNotFound, errorResponse.Code)
}

func TestRegisterTenantWrongEmail(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	leaseInvite := BuildTestLeaseInvite()
	leaseInvite.TenantEmail = TENANT_EMAIL
	mock.LeaseInvite.Expect(
		client.Client.LeaseInvite.FindUnique(db.LeaseInvite.ID.Equals(leaseInvite.ID)),
	).Returns(leaseInvite)

	user := BuildTestUser("1")
	user.Email = "test2@example.com"
	b, err := json.Marshal(user)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/v1/auth/invite/1/", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	var errorResponse utils.Error
	err = json.Unmarshal(w.Body.Bytes(), &errorResponse)
	require.NoError(t, err)
	assert.Equal(t, utils.UserSameEmailAsInvite, errorResponse.Code)
}

func TestRegisterTenantPropertyNotAvailable(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	leaseInvite := BuildTestLeaseInvite()
	leaseInvite.TenantEmail = TENANT_EMAIL
	mock.LeaseInvite.Expect(
		client.Client.LeaseInvite.FindUnique(db.LeaseInvite.ID.Equals(leaseInvite.ID)),
	).Returns(leaseInvite)

	user := BuildTestUser("1")
	user.Email = leaseInvite.TenantEmail
	b, err := json.Marshal(user)
	require.NoError(t, err)

	mock.Lease.Expect(
		client.Client.Lease.FindMany(
			db.Lease.PropertyID.Equals("1"),
			db.Lease.Active.Equals(true),
		).With(
			db.Lease.Tenant.Fetch(),
			db.Lease.Property.Fetch().With(db.Property.Owner.Fetch()),
		),
	).ReturnsMany([]db.LeaseModel{BuildTestLease("1")})

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/v1/auth/invite/1/", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusConflict, w.Code)
	var errorResponse utils.Error
	err = json.Unmarshal(w.Body.Bytes(), &errorResponse)
	require.NoError(t, err)
	assert.Equal(t, utils.PropertyNotAvailable, errorResponse.Code)
}

func TestAcceptInvite(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	user := BuildTestUser("1")
	user.Role = db.RoleTenant
	mock.User.Expect(
		client.Client.User.FindUnique(db.User.ID.Equals(user.ID)),
	).Returns(user)

	leaseInvite := BuildTestLeaseInvite()
	leaseInvite.TenantEmail = user.Email
	mock.LeaseInvite.Expect(
		client.Client.LeaseInvite.FindUnique(db.LeaseInvite.ID.Equals(leaseInvite.ID)),
	).Returns(leaseInvite)

	mock.Lease.Expect(
		client.Client.Lease.FindMany(
			db.Lease.PropertyID.Equals(leaseInvite.PropertyID),
			db.Lease.Active.Equals(true),
		).With(
			db.Lease.Tenant.Fetch(),
			db.Lease.Property.Fetch().With(db.Property.Owner.Fetch()),
		),
	).ReturnsMany([]db.LeaseModel{})

	mock.Lease.Expect(
		client.Client.Lease.FindMany(
			db.Lease.TenantID.Equals(user.ID),
			db.Lease.Active.Equals(true),
		).With(
			db.Lease.Tenant.Fetch(),
			db.Lease.Property.Fetch().With(db.Property.Owner.Fetch()),
		),
	).ReturnsMany([]db.LeaseModel{})

	mock.Lease.Expect(
		client.Client.Lease.CreateOne(
			db.Lease.StartDate.Set(leaseInvite.StartDate),
			db.Lease.Tenant.Link(db.User.ID.Equals(user.ID)),
			db.Lease.Property.Link(db.Property.ID.Equals(leaseInvite.PropertyID)),
			db.Lease.EndDate.SetIfPresent(leaseInvite.InnerLeaseInvite.EndDate),
		),
	).Returns(BuildTestLease("1"))

	mock.LeaseInvite.Expect(
		client.Client.LeaseInvite.FindUnique(
			db.LeaseInvite.ID.Equals(leaseInvite.ID),
		).Delete(),
	).Returns(db.LeaseInviteModel{})

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/v1/tenant/invite/1/", nil)
	req.Header.Set("Oauth.claims.id", user.ID)
	req.Header.Set("Oauth.claims.role", string(user.Role))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestAcceptInviteNotATenant(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	user := BuildTestUser("1")
	mock.User.Expect(
		client.Client.User.FindUnique(db.User.ID.Equals(user.ID)),
	).Returns(user)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/v1/tenant/invite/1/", nil)
	req.Header.Set("Oauth.claims.id", user.ID)
	req.Header.Set("Oauth.claims.role", string(db.RoleTenant))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusForbidden, w.Code)
	var errorResponse utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
	require.NoError(t, err)
	assert.Equal(t, utils.NotATenant, errorResponse.Code)
}

func TestAcceptInviteInviteNotFound(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	user := BuildTestUser("1")
	user.Role = db.RoleTenant
	mock.User.Expect(
		client.Client.User.FindUnique(db.User.ID.Equals(user.ID)),
	).Returns(user)

	mock.LeaseInvite.Expect(
		client.Client.LeaseInvite.FindUnique(db.LeaseInvite.ID.Equals("wrong")),
	).Errors(db.ErrNotFound)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/v1/tenant/invite/wrong/", nil)
	req.Header.Set("Oauth.claims.id", user.ID)
	req.Header.Set("Oauth.claims.role", string(user.Role))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	var errorResponse utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
	require.NoError(t, err)
	assert.Equal(t, utils.InviteNotFound, errorResponse.Code)
}

func TestAcceptInviteWrongEmail(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	user := BuildTestUser("1")
	user.Role = db.RoleTenant
	user.Email = TENANT_EMAIL
	mock.User.Expect(
		client.Client.User.FindUnique(db.User.ID.Equals(user.ID)),
	).Returns(user)

	leaseInvite := BuildTestLeaseInvite()
	leaseInvite.TenantEmail = "test2@example.com"
	mock.LeaseInvite.Expect(
		client.Client.LeaseInvite.FindUnique(db.LeaseInvite.ID.Equals(leaseInvite.ID)),
	).Returns(leaseInvite)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/v1/tenant/invite/1/", nil)
	req.Header.Set("Oauth.claims.id", user.ID)
	req.Header.Set("Oauth.claims.role", string(user.Role))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusForbidden, w.Code)
	var errorResponse utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
	require.NoError(t, err)
	assert.Equal(t, utils.UserSameEmailAsInvite, errorResponse.Code)
}

func TestAcceptInvitePropertyNotAvailable(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	user := BuildTestUser("1")
	user.Role = db.RoleTenant
	mock.User.Expect(
		client.Client.User.FindUnique(db.User.ID.Equals(user.ID)),
	).Returns(user)

	leaseInvite := BuildTestLeaseInvite()
	leaseInvite.TenantEmail = user.Email
	mock.LeaseInvite.Expect(
		client.Client.LeaseInvite.FindUnique(db.LeaseInvite.ID.Equals(leaseInvite.ID)),
	).Returns(leaseInvite)

	activeLease := BuildTestLease("1")
	mock.Lease.Expect(
		client.Client.Lease.FindMany(
			db.Lease.PropertyID.Equals(leaseInvite.PropertyID),
			db.Lease.Active.Equals(true),
		).With(
			db.Lease.Tenant.Fetch(),
			db.Lease.Property.Fetch().With(db.Property.Owner.Fetch()),
		),
	).ReturnsMany([]db.LeaseModel{activeLease})

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/v1/tenant/invite/1/", nil)
	req.Header.Set("Oauth.claims.id", user.ID)
	req.Header.Set("Oauth.claims.role", string(user.Role))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusConflict, w.Code)
	var errorResponse utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
	require.NoError(t, err)
	assert.Equal(t, utils.PropertyNotAvailable, errorResponse.Code)
}

func TestAcceptInviteTenantAlreadyHasLease(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	user := BuildTestUser("1")
	user.Role = db.RoleTenant
	mock.User.Expect(
		client.Client.User.FindUnique(db.User.ID.Equals(user.ID)),
	).Returns(user)

	leaseInvite := BuildTestLeaseInvite()
	leaseInvite.TenantEmail = user.Email
	mock.LeaseInvite.Expect(
		client.Client.LeaseInvite.FindUnique(db.LeaseInvite.ID.Equals(leaseInvite.ID)),
	).Returns(leaseInvite)

	activeLease := BuildTestLease("1")
	mock.Lease.Expect(
		client.Client.Lease.FindMany(
			db.Lease.PropertyID.Equals(leaseInvite.PropertyID),
			db.Lease.Active.Equals(true),
		).With(
			db.Lease.Tenant.Fetch(),
			db.Lease.Property.Fetch().With(db.Property.Owner.Fetch()),
		),
	).Errors(db.ErrNotFound)

	mock.Lease.Expect(
		client.Client.Lease.FindMany(
			db.Lease.TenantID.Equals(user.ID),
			db.Lease.Active.Equals(true),
		).With(
			db.Lease.Tenant.Fetch(),
			db.Lease.Property.Fetch().With(db.Property.Owner.Fetch()),
		),
	).ReturnsMany([]db.LeaseModel{activeLease})

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/v1/tenant/invite/1/", nil)
	req.Header.Set("Oauth.claims.id", user.ID)
	req.Header.Set("Oauth.claims.role", string(user.Role))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusConflict, w.Code)
	var errorResponse utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
	require.NoError(t, err)
	assert.Equal(t, utils.TenantAlreadyHasLease, errorResponse.Code)
}
