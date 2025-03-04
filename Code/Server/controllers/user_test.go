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
			Role:      db.RoleOwner,
		},
	}
}

func TestGetAllUsers(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
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
	client, mock, ensure := services.ConnectDBTest()
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

func TestGetUserByID_NotFound(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
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
	client, mock, ensure := services.ConnectDBTest()
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

func TestGetProfile_UserNotFound(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
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

func TestGetUserProfilePicture(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	user := BuildTestUser("1")
	image := BuildTestImage("1", "b3Vp")
	user.InnerUser.ProfilePictureID = &image.ID

	mock.User.Expect(
		client.Client.User.FindUnique(db.User.ID.Equals(user.ID)),
	).Returns(user)

	mock.Image.Expect(
		client.Client.Image.FindUnique(db.Image.ID.Equals(image.ID)),
	).Returns(image)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/user/"+user.ID+"/picture/", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var imageResponse models.ImageResponse
	err := json.Unmarshal(w.Body.Bytes(), &imageResponse)
	require.NoError(t, err)
	assert.Equal(t, image.ID, imageResponse.ID)
}

func TestGetUserProfilePicture_NoContent(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	user := BuildTestUser("1")

	mock.User.Expect(
		client.Client.User.FindUnique(db.User.ID.Equals(user.ID)),
	).Returns(user)

	r := router.TestRoutes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/user/"+user.ID+"/picture/", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestGetUserProfilePicture_UserNotFound(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	mock.User.Expect(
		client.Client.User.FindUnique(db.User.ID.Equals("nonexistent")),
	).Errors(db.ErrNotFound)

	r := router.TestRoutes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/user/nonexistent/picture/", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	var errorResponse utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
	require.NoError(t, err)
	assert.Equal(t, utils.UserNotFound, errorResponse.Code)
}

func TestGetUserProfilePicture_ImageNotFound(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	user := BuildTestUser("1")
	user.InnerUser.ProfilePictureID = utils.Ptr("nonexistent")

	mock.User.Expect(
		client.Client.User.FindUnique(db.User.ID.Equals(user.ID)),
	).Returns(user)

	mock.Image.Expect(
		client.Client.Image.FindUnique(db.Image.ID.Equals("nonexistent")),
	).Errors(db.ErrNotFound)

	r := router.TestRoutes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/user/"+user.ID+"/picture/", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	var errorResponse utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
	require.NoError(t, err)
	assert.Equal(t, utils.UserProfilePictureNotFound, errorResponse.Code)
}

func TestUpdateCurrentUserProfile(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	user := BuildTestUser("1")
	mock.User.Expect(
		client.Client.User.FindUnique(db.User.ID.Equals(user.ID)),
	).Returns(user)

	updatedUser := user
	updatedUser.InnerUser.Firstname = "Updated"
	mock.User.Expect(
		client.Client.User.FindUnique(db.User.ID.Equals(user.ID)).Update(
			db.User.Email.SetIfPresent(nil),
			db.User.Firstname.SetIfPresent(&updatedUser.Firstname),
			db.User.Lastname.SetIfPresent(nil),
		),
	).Returns(updatedUser)

	reqBody := models.UserUpdateRequest{
		Firstname: utils.Ptr("Updated"),
	}
	b, err := json.Marshal(reqBody)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/profile/", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var userResponse models.UserResponse
	err = json.Unmarshal(w.Body.Bytes(), &userResponse)
	require.NoError(t, err)
	assert.Equal(t, updatedUser.ID, userResponse.ID)
	assert.Equal(t, "Updated", userResponse.Firstname)
}

func TestUpdateCurrentUserProfile_UserNotFound(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	mock.User.Expect(
		client.Client.User.FindUnique(db.User.ID.Equals("nonexistent")),
	).Errors(db.ErrNotFound)

	reqBody := models.UserUpdateRequest{
		Firstname: utils.Ptr("Updated"),
	}
	b, err := json.Marshal(reqBody)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/profile/", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "nonexistent")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	var errorResponse utils.Error
	err = json.Unmarshal(w.Body.Bytes(), &errorResponse)
	require.NoError(t, err)
	assert.Equal(t, utils.UserNotFound, errorResponse.Code)
}

func TestUpdateCurrentUserProfile_MissingFields(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	user := BuildTestUser("1")
	mock.User.Expect(
		client.Client.User.FindUnique(db.User.ID.Equals(user.ID)),
	).Returns(user)

	reqBody := models.UserUpdateRequest{
		Email: utils.Ptr(""),
	}
	b, err := json.Marshal(reqBody)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/profile/", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	var errorResponse utils.Error
	err = json.Unmarshal(w.Body.Bytes(), &errorResponse)
	require.NoError(t, err)
	assert.Equal(t, utils.MissingFields, errorResponse.Code)
}

func TestUpdateCurrentUserProfile_EmailAlreadyExists(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	user := BuildTestUser("1")
	mock.User.Expect(
		client.Client.User.FindUnique(db.User.ID.Equals(user.ID)),
	).Returns(user)

	mock.User.Expect(
		client.Client.User.FindUnique(db.User.ID.Equals(user.ID)).Update(
			db.User.Email.SetIfPresent(utils.Ptr("existing@example.com")),
			db.User.Firstname.SetIfPresent(nil),
			db.User.Lastname.SetIfPresent(nil),
		),
	).Errors(&protocol.UserFacingError{
		IsPanic:   false,
		ErrorCode: "P2002", // https://www.prisma.io/docs/orm/reference/error-reference#p2002
		Meta: protocol.Meta{
			Target: []any{"email"},
		},
		Message: "Unique constraint failed",
	})

	reqBody := models.UserUpdateRequest{
		Email: utils.Ptr("existing@example.com"),
	}
	b, err := json.Marshal(reqBody)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/profile/", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusConflict, w.Code)
	var errorResponse utils.Error
	err = json.Unmarshal(w.Body.Bytes(), &errorResponse)
	require.NoError(t, err)
	assert.Equal(t, utils.EmailAlreadyExists, errorResponse.Code)
}

func TestUpdateCurrentUserProfilePicture(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	user := BuildTestUser("1")
	mock.User.Expect(
		client.Client.User.FindUnique(db.User.ID.Equals(user.ID)),
	).Returns(user)

	image := BuildTestImage("1", "b3Vp")
	mock.Image.Expect(
		client.Client.Image.CreateOne(
			db.Image.Data.Set(image.Data),
		),
	).Returns(image)

	updatedUser := user
	updatedUser.InnerUser.ProfilePictureID = &image.ID
	mock.User.Expect(
		client.Client.User.FindUnique(db.User.ID.Equals(user.ID)).Update(
			db.User.ProfilePicture.Link(db.Image.ID.Equals(image.ID)),
		),
	).Returns(updatedUser)

	reqBody := models.ImageRequest{
		Data: "b3Vp",
	}
	b, err := json.Marshal(reqBody)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/profile/picture/", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var userResponse models.UserResponse
	err = json.Unmarshal(w.Body.Bytes(), &userResponse)
	require.NoError(t, err)
	assert.Equal(t, updatedUser.ID, userResponse.ID)
	assert.Equal(t, *updatedUser.InnerUser.ProfilePictureID, *userResponse.ProfilePictureID)
}

func TestUpdateCurrentUserProfilePicture_UserNotFound(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	mock.User.Expect(
		client.Client.User.FindUnique(db.User.ID.Equals("nonexistent")),
	).Errors(db.ErrNotFound)

	reqBody := models.ImageRequest{
		Data: "b3Vp",
	}
	b, err := json.Marshal(reqBody)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/profile/picture/", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "nonexistent")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	var errorResponse utils.Error
	err = json.Unmarshal(w.Body.Bytes(), &errorResponse)
	require.NoError(t, err)
	assert.Equal(t, utils.UserNotFound, errorResponse.Code)
}

func TestUpdateCurrentUserProfilePicture_MissingFields(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	user := BuildTestUser("1")
	mock.User.Expect(
		client.Client.User.FindUnique(db.User.ID.Equals(user.ID)),
	).Returns(user)

	reqBody := models.ImageRequest{}
	b, err := json.Marshal(reqBody)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/profile/picture/", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	var errorResponse utils.Error
	err = json.Unmarshal(w.Body.Bytes(), &errorResponse)
	require.NoError(t, err)
	assert.Equal(t, utils.MissingFields, errorResponse.Code)
}

func TestUpdateCurrentUserProfilePicture_BadBase64String(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	user := BuildTestUser("1")
	mock.User.Expect(
		client.Client.User.FindUnique(db.User.ID.Equals(user.ID)),
	).Returns(user)

	reqBody := models.ImageRequest{
		Data: "invalid_base64",
	}
	b, err := json.Marshal(reqBody)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/profile/picture/", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	var errorResponse utils.Error
	err = json.Unmarshal(w.Body.Bytes(), &errorResponse)
	require.NoError(t, err)
	assert.Equal(t, utils.MissingFields, errorResponse.Code)
}

func TestGetCurrentUserProfilePicture(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	user := BuildTestUser("1")
	image := BuildTestImage("1", "b3Vp")
	user.InnerUser.ProfilePictureID = &image.ID

	mock.User.Expect(
		client.Client.User.FindUnique(db.User.ID.Equals(user.ID)),
	).Returns(user)

	mock.Image.Expect(
		client.Client.Image.FindUnique(db.Image.ID.Equals(image.ID)),
	).Returns(image)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/profile/picture/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var imageResponse models.ImageResponse
	err := json.Unmarshal(w.Body.Bytes(), &imageResponse)
	require.NoError(t, err)
	assert.Equal(t, image.ID, imageResponse.ID)
}

func TestGetCurrentUserProfilePicture_NoContent(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	user := BuildTestUser("1")

	mock.User.Expect(
		client.Client.User.FindUnique(db.User.ID.Equals(user.ID)),
	).Returns(user)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/profile/picture/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestGetCurrentUserProfilePicture_UserNotFound(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	mock.User.Expect(
		client.Client.User.FindUnique(db.User.ID.Equals("nonexistent")),
	).Errors(db.ErrNotFound)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/profile/picture/", nil)
	req.Header.Set("Oauth.claims.id", "nonexistent")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	var errorResponse utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
	require.NoError(t, err)
	assert.Equal(t, utils.UserNotFound, errorResponse.Code)
}

func TestGetCurrentUserProfilePicture_ImageNotFound(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	user := BuildTestUser("1")
	user.InnerUser.ProfilePictureID = utils.Ptr("nonexistent")

	mock.User.Expect(
		client.Client.User.FindUnique(db.User.ID.Equals(user.ID)),
	).Returns(user)

	mock.Image.Expect(
		client.Client.Image.FindUnique(db.Image.ID.Equals("nonexistent")),
	).Errors(db.ErrNotFound)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/profile/picture/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	var errorResponse utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
	require.NoError(t, err)
	assert.Equal(t, utils.UserProfilePictureNotFound, errorResponse.Code)
}
