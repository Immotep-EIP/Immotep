package controllers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"immotep/backend/database"
	"immotep/backend/models"
	"immotep/backend/prisma/db"
	"immotep/backend/router"
	"immotep/backend/utils"
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

	r := router.TestRoutes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/users/", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var users []models.UserResponse
	err := json.Unmarshal(w.Body.Bytes(), &users)
	require.NoError(t, err)
	assert.JSONEq(t, users[0].ID, user.ID)
}

func TestGetUserByID(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	user := BuildTestUser("1")
	mock.User.Expect(
		client.Client.User.FindUnique(db.User.ID.Equals(user.ID)),
	).Returns(user)

	r := router.TestRoutes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/user/"+user.ID+"/", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var userResponse models.UserResponse
	err := json.Unmarshal(w.Body.Bytes(), &userResponse)
	require.NoError(t, err)
	assert.Equal(t, user.ID, userResponse.ID)
}

func TestGetUserByIDNotFound(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	mock.User.Expect(
		client.Client.User.FindUnique(db.User.ID.Equals("nonexistent")),
	).Errors(db.ErrNotFound)

	r := router.TestRoutes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/user/nonexistent/", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	var errorResponse utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
	require.NoError(t, err)
	assert.Equal(t, utils.UserNotFound, errorResponse.Code)
}

func TestGetProfile(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	user := BuildTestUser("1")
	mock.User.Expect(
		client.Client.User.FindUnique(db.User.ID.Equals(user.ID)),
	).Returns(user)

	r := router.TestRoutes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/profile/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var userResponse models.UserResponse
	err := json.Unmarshal(w.Body.Bytes(), &userResponse)
	require.NoError(t, err)
	assert.Equal(t, user.ID, userResponse.ID)
}

func TestGetProfileUserNotFound(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	mock.User.Expect(
		client.Client.User.FindUnique(db.User.ID.Equals("nonexistent")),
	).Errors(db.ErrNotFound)

	r := router.TestRoutes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/profile/", nil)
	req.Header.Set("Oauth.claims.id", "nonexistent")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	var errorResponse utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
	require.NoError(t, err)
	assert.Equal(t, utils.UserNotFound, errorResponse.Code)
}
