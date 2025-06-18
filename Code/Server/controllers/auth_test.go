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
	"keyz/backend/controllers"
	"keyz/backend/prisma/db"
	"keyz/backend/router"
	"keyz/backend/services"
	"keyz/backend/services/database"
	"keyz/backend/utils"
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
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.LeaseInvite.Expect(database.MockGetLeaseInviteByID(c)).Errors(db.ErrNotFound)

	user := BuildTestUser("1")
	b, err := json.Marshal(user)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/v1/auth/invite/1/", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	var errorResponse utils.Error
	err = json.Unmarshal(w.Body.Bytes(), &errorResponse)
	require.NoError(t, err)
	assert.Equal(t, utils.InviteNotFound, errorResponse.Code)
}

func TestRegisterTenantWrongEmail(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	leaseInvite := BuildTestLeaseInvite()
	leaseInvite.TenantEmail = TENANT_EMAIL
	m.LeaseInvite.Expect(database.MockGetLeaseInviteByID(c)).Returns(leaseInvite)

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
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	leaseInvite := BuildTestLeaseInvite()
	leaseInvite.TenantEmail = TENANT_EMAIL
	m.LeaseInvite.Expect(database.MockGetLeaseInviteByID(c)).Returns(leaseInvite)
	m.Lease.Expect(database.MockGetCurrentActiveLeaseByProperty(c)).ReturnsMany([]db.LeaseModel{BuildTestLease("1")})

	user := BuildTestUser("1")
	user.Email = leaseInvite.TenantEmail
	b, err := json.Marshal(user)
	require.NoError(t, err)

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
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	user := BuildTestUser("1")
	user.Role = db.RoleTenant
	leaseInvite := BuildTestLeaseInvite()
	leaseInvite.TenantEmail = user.Email
	m.User.Expect(database.MockGetUserByID(c)).Returns(user)
	m.LeaseInvite.Expect(database.MockGetLeaseInviteByID(c)).Returns(leaseInvite)
	m.Lease.Expect(database.MockGetCurrentActiveLeaseByProperty(c)).ReturnsMany([]db.LeaseModel{})
	m.Lease.Expect(database.MockGetCurrentActiveLeaseByTenant(c)).ReturnsMany([]db.LeaseModel{})
	m.Lease.Expect(database.MockCreateLease(c, leaseInvite)).Returns(BuildTestLease("1"))
	m.LeaseInvite.Expect(database.MockDeleteLeaseInviteById(c)).Returns(db.LeaseInviteModel{})

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/v1/tenant/invite/1/", nil)
	req.Header.Set("Oauth.claims.id", user.ID)
	req.Header.Set("Oauth.claims.role", string(user.Role))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestAcceptInviteNotATenant(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	user := BuildTestUser("1")
	m.User.Expect(database.MockGetUserByID(c)).Returns(user)

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
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	user := BuildTestUser("1")
	user.Role = db.RoleTenant
	m.User.Expect(database.MockGetUserByID(c)).Returns(user)
	m.LeaseInvite.Expect(database.MockGetLeaseInviteByID(c)).Errors(db.ErrNotFound)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/v1/tenant/invite/1/", nil)
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
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	user := BuildTestUser("1")
	user.Role = db.RoleTenant
	user.Email = TENANT_EMAIL
	leaseInvite := BuildTestLeaseInvite()
	leaseInvite.TenantEmail = "test2@example.com"
	m.User.Expect(database.MockGetUserByID(c)).Returns(user)
	m.LeaseInvite.Expect(database.MockGetLeaseInviteByID(c)).Returns(leaseInvite)

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
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	user := BuildTestUser("1")
	user.Role = db.RoleTenant
	leaseInvite := BuildTestLeaseInvite()
	leaseInvite.TenantEmail = user.Email
	activeLease := BuildTestLease("1")
	m.User.Expect(database.MockGetUserByID(c)).Returns(user)
	m.LeaseInvite.Expect(database.MockGetLeaseInviteByID(c)).Returns(leaseInvite)
	m.Lease.Expect(database.MockGetCurrentActiveLeaseByProperty(c)).ReturnsMany([]db.LeaseModel{activeLease})

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
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	user := BuildTestUser("1")
	user.Role = db.RoleTenant
	leaseInvite := BuildTestLeaseInvite()
	leaseInvite.TenantEmail = user.Email
	activeLease := BuildTestLease("1")
	m.User.Expect(database.MockGetUserByID(c)).Returns(user)
	m.LeaseInvite.Expect(database.MockGetLeaseInviteByID(c)).Returns(leaseInvite)
	m.Lease.Expect(database.MockGetCurrentActiveLeaseByProperty(c)).Errors(db.ErrNotFound)
	m.Lease.Expect(database.MockGetCurrentActiveLeaseByTenant(c)).ReturnsMany([]db.LeaseModel{activeLease})

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
