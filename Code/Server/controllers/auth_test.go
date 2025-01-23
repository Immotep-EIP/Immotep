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
	"immotep/backend/database"
	"immotep/backend/prisma/db"
	"immotep/backend/router"
	"immotep/backend/utils"
)

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

	gin.SetMode(gin.TestMode)
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

	gin.SetMode(gin.TestMode)
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

	gin.SetMode(gin.TestMode)
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

	gin.SetMode(gin.TestMode)
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
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	mock.PendingContract.Expect(
		client.Client.PendingContract.FindUnique(db.PendingContract.ID.Equals("wrong")),
	).Errors(db.ErrNotFound)

	user := BuildTestUser("1")
	b, err := json.Marshal(user)
	require.NoError(t, err)

	gin.SetMode(gin.TestMode)
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
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	pendingContract := BuildTestPendingContract()
	pendingContract.TenantEmail = "test1@example.com"
	mock.PendingContract.Expect(
		client.Client.PendingContract.FindUnique(db.PendingContract.ID.Equals(pendingContract.ID)),
	).Returns(pendingContract)

	user := BuildTestUser("1")
	user.Email = "test2@example.com"
	b, err := json.Marshal(user)
	require.NoError(t, err)

	gin.SetMode(gin.TestMode)
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
