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
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/auth/register/", bytes.NewReader(b))
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
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/auth/register/", bytes.NewReader(b))
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
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/auth/invite/1/", bytes.NewReader(b))
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
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/auth/invite/1/", bytes.NewReader(b))
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

	mock.PendingContract.Expect(
		client.Client.PendingContract.FindUnique(db.PendingContract.ID.Equals("wrong")),
	).Errors(db.ErrNotFound)

	user := BuildTestUser("1")
	b, err := json.Marshal(user)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/auth/invite/wrong/", bytes.NewReader(b))
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

	pendingContract := BuildTestPendingContract()
	pendingContract.TenantEmail = TENANT_EMAIL
	mock.PendingContract.Expect(
		client.Client.PendingContract.FindUnique(db.PendingContract.ID.Equals(pendingContract.ID)),
	).Returns(pendingContract)

	user := BuildTestUser("1")
	user.Email = "test2@example.com"
	b, err := json.Marshal(user)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/auth/invite/1/", bytes.NewReader(b))
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

	pendingContract := BuildTestPendingContract()
	pendingContract.TenantEmail = TENANT_EMAIL
	mock.PendingContract.Expect(
		client.Client.PendingContract.FindUnique(db.PendingContract.ID.Equals(pendingContract.ID)),
	).Returns(pendingContract)

	user := BuildTestUser("1")
	user.Email = pendingContract.TenantEmail
	b, err := json.Marshal(user)
	require.NoError(t, err)

	mock.Contract.Expect(
		client.Client.Contract.FindMany(
			db.Contract.PropertyID.Equals("1"),
			db.Contract.Active.Equals(true),
		),
	).ReturnsMany([]db.ContractModel{BuildTestContract()})

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/auth/invite/1/", bytes.NewReader(b))
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

	pendingContract := BuildTestPendingContract()
	pendingContract.TenantEmail = user.Email
	mock.PendingContract.Expect(
		client.Client.PendingContract.FindUnique(db.PendingContract.ID.Equals(pendingContract.ID)),
	).Returns(pendingContract)

	mock.Contract.Expect(
		client.Client.Contract.FindMany(
			db.Contract.PropertyID.Equals(pendingContract.PropertyID),
			db.Contract.Active.Equals(true),
		),
	).ReturnsMany([]db.ContractModel{})

	mock.Contract.Expect(
		client.Client.Contract.FindMany(
			db.Contract.TenantID.Equals(user.ID),
			db.Contract.Active.Equals(true),
		),
	).ReturnsMany([]db.ContractModel{})

	mock.Contract.Expect(
		client.Client.Contract.CreateOne(
			db.Contract.StartDate.Set(pendingContract.StartDate),
			db.Contract.Tenant.Link(db.User.ID.Equals(user.ID)),
			db.Contract.Property.Link(db.Property.ID.Equals(pendingContract.PropertyID)),
			db.Contract.EndDate.SetIfPresent(pendingContract.InnerPendingContract.EndDate),
		),
	).Returns(BuildTestContract())

	mock.PendingContract.Expect(
		client.Client.PendingContract.FindUnique(
			db.PendingContract.ID.Equals(pendingContract.ID),
		).Delete(),
	).Returns(db.PendingContractModel{})

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/tenant/invite/1/", nil)
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
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/tenant/invite/1/", nil)
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

	mock.PendingContract.Expect(
		client.Client.PendingContract.FindUnique(db.PendingContract.ID.Equals("wrong")),
	).Errors(db.ErrNotFound)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/tenant/invite/wrong/", nil)
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

	pendingContract := BuildTestPendingContract()
	pendingContract.TenantEmail = "test2@example.com"
	mock.PendingContract.Expect(
		client.Client.PendingContract.FindUnique(db.PendingContract.ID.Equals(pendingContract.ID)),
	).Returns(pendingContract)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/tenant/invite/1/", nil)
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

	pendingContract := BuildTestPendingContract()
	pendingContract.TenantEmail = user.Email
	mock.PendingContract.Expect(
		client.Client.PendingContract.FindUnique(db.PendingContract.ID.Equals(pendingContract.ID)),
	).Returns(pendingContract)

	activeContract := BuildTestContract()
	mock.Contract.Expect(
		client.Client.Contract.FindMany(
			db.Contract.PropertyID.Equals(pendingContract.PropertyID),
			db.Contract.Active.Equals(true),
		),
	).ReturnsMany([]db.ContractModel{activeContract})

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/tenant/invite/1/", nil)
	req.Header.Set("Oauth.claims.id", user.ID)
	req.Header.Set("Oauth.claims.role", string(user.Role))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusConflict, w.Code)
	var errorResponse utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
	require.NoError(t, err)
	assert.Equal(t, utils.PropertyNotAvailable, errorResponse.Code)
}

func TestAcceptInviteTenantAlreadyHasContract(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	user := BuildTestUser("1")
	user.Role = db.RoleTenant
	mock.User.Expect(
		client.Client.User.FindUnique(db.User.ID.Equals(user.ID)),
	).Returns(user)

	pendingContract := BuildTestPendingContract()
	pendingContract.TenantEmail = user.Email
	mock.PendingContract.Expect(
		client.Client.PendingContract.FindUnique(db.PendingContract.ID.Equals(pendingContract.ID)),
	).Returns(pendingContract)

	activeContract := BuildTestContract()
	mock.Contract.Expect(
		client.Client.Contract.FindMany(
			db.Contract.PropertyID.Equals(pendingContract.PropertyID),
			db.Contract.Active.Equals(true),
		),
	).Errors(db.ErrNotFound)

	mock.Contract.Expect(
		client.Client.Contract.FindMany(
			db.Contract.TenantID.Equals(user.ID),
			db.Contract.Active.Equals(true),
		),
	).ReturnsMany([]db.ContractModel{activeContract})

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/tenant/invite/1/", nil)
	req.Header.Set("Oauth.claims.id", user.ID)
	req.Header.Set("Oauth.claims.role", string(user.Role))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusConflict, w.Code)
	var errorResponse utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
	require.NoError(t, err)
	assert.Equal(t, utils.TenantAlreadyHasContract, errorResponse.Code)
}
