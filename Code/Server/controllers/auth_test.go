package controllers_test

import (
	"bytes"
	"encoding/json"
	"immotep/backend/controllers"
	"immotep/backend/router"
	"immotep/backend/utils"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/maxzerbini/oauth"
	"github.com/stretchr/testify/assert"
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

func TestCreateUserInvalidBody(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("POST", "/users", nil)
	c.Request.Header.Set("Content-Type", "application/json")

	controllers.CreateUser(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	var errorResponse utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
	assert.NoError(t, err)
	assert.Equal(t, utils.MissingFields, errorResponse.Code)
}

func TestCreateUserMissingFields(t *testing.T) {
	user := BuildTestUser("1")
	user.Email = ""
	b, err := json.Marshal(user)
	assert.NoError(t, err)

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("POST", "/users", bytes.NewReader(b))
	c.Request.Header.Set("Content-Type", "application/json")

	controllers.CreateUser(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	var errorResponse utils.Error
	err = json.Unmarshal(w.Body.Bytes(), &errorResponse)
	assert.NoError(t, err)
	assert.Equal(t, utils.MissingFields, errorResponse.Code)
}

func TestCreateUserWrongPassword(t *testing.T) {
	user := BuildTestUser("1")
	user.Password = "1234"
	b, err := json.Marshal(user)
	assert.NoError(t, err)

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("POST", "/users", bytes.NewReader(b))
	c.Request.Header.Set("Content-Type", "application/json")

	controllers.CreateUser(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	var errorResponse utils.Error
	err = json.Unmarshal(w.Body.Bytes(), &errorResponse)
	assert.NoError(t, err)
	assert.Equal(t, utils.MissingFields, errorResponse.Code)
}