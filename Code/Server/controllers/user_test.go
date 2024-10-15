package controllers_test

import (
	"encoding/json"
	"immotep/backend/controllers"
	"immotep/backend/database"
	"immotep/backend/models"
	"immotep/backend/prisma/db"
	"immotep/backend/utils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func BuildTestUser(id string) db.UserModel {
	return db.UserModel{
		InnerUser: db.InnerUser{
			ID:        id,
			Email:     "test" + id + "@example.com",
			Firstname: "Test",
			Lastname:  "User",
			Password:  "Password123",
		},
	}
}

func TestGetAllUsers(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	user := BuildTestUser("1")
	mock.User.Expect(
		client.Client.User.FindMany(),
	).ReturnsMany([]db.UserModel{user})

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	controllers.GetAllUsers(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var users []models.UserResponse
	err := json.Unmarshal(w.Body.Bytes(), &users)
	assert.NoError(t, err)
	assert.JSONEq(t, users[0].ID, user.ID)
}

func TestGetUserByID(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	user := BuildTestUser("1")
	mock.User.Expect(
		client.Client.User.FindUnique(db.User.ID.Equals(user.ID)),
	).Returns(user)

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{gin.Param{Key: "id", Value: user.ID}}
	controllers.GetUserByID(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var userResponse models.UserResponse
	err := json.Unmarshal(w.Body.Bytes(), &userResponse)
	assert.NoError(t, err)
	assert.Equal(t, user.ID, userResponse.ID)
}

func TestGetUserByIDNotFound(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	mock.User.Expect(
		client.Client.User.FindUnique(db.User.ID.Equals("nonexistent")),
	).Errors(db.ErrNotFound)

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{gin.Param{Key: "id", Value: "nonexistent"}}
	controllers.GetUserByID(c)

	assert.Equal(t, http.StatusNotFound, w.Code)
	var errorResponse utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
	assert.NoError(t, err)
	assert.Equal(t, utils.CannotFindUser, errorResponse.Code)
}

func TestGetProfile(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	user := BuildTestUser("1")
	mock.User.Expect(
		client.Client.User.FindUnique(db.User.ID.Equals(user.ID)),
	).Returns(user)

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("oauth.claims", map[string]string{"id": user.ID})
	controllers.GetProfile(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var userResponse models.UserResponse
	err := json.Unmarshal(w.Body.Bytes(), &userResponse)
	assert.NoError(t, err)
	assert.Equal(t, user.ID, userResponse.ID)
}

func TestGetProfileNoClaims(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	controllers.GetProfile(c)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	var errorResponse utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
	assert.NoError(t, err)
	assert.Equal(t, utils.NoClaims, errorResponse.Code)
}

func TestGetProfileUserNotFound(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	mock.User.Expect(
		client.Client.User.FindUnique(db.User.ID.Equals("nonexistent")),
	).Errors(db.ErrNotFound)

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("oauth.claims", map[string]string{"id": "nonexistent"})
	controllers.GetProfile(c)

	assert.Equal(t, http.StatusNotFound, w.Code)
	var errorResponse utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
	assert.NoError(t, err)
	assert.Equal(t, utils.CannotFindUser, errorResponse.Code)
}
