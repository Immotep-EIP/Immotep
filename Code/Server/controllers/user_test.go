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
	"keyz/backend/models"
	"keyz/backend/prisma/db"
	"keyz/backend/router"
	"keyz/backend/services"
	"keyz/backend/services/database"
	"keyz/backend/utils"
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
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	user := BuildTestUser("1")
	m.User.Expect(database.MockGetAllUsers(c)).ReturnsMany([]db.UserModel{user})

	r := router.TestRoutes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/users/", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var users []models.UserResponse
	err := json.Unmarshal(w.Body.Bytes(), &users)
	require.NoError(t, err)
	assert.JSONEq(t, users[0].ID, user.ID)
}

func TestGetUserByID(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	user := BuildTestUser("1")
	m.User.Expect(database.MockGetUserByID(c)).Returns(user)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/user/1/", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var userResponse models.UserResponse
	err := json.Unmarshal(w.Body.Bytes(), &userResponse)
	require.NoError(t, err)
	assert.Equal(t, user.ID, userResponse.ID)
}

func TestGetUserByID_NotFound(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.User.Expect(database.MockGetUserByID(c)).Errors(db.ErrNotFound)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/user/1/", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	var errorResponse utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
	require.NoError(t, err)
	assert.Equal(t, utils.UserNotFound, errorResponse.Code)
}

func TestGetProfile(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	user := BuildTestUser("1")
	m.User.Expect(database.MockGetUserByID(c)).Returns(user)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/profile/", nil)
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
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.User.Expect(database.MockGetUserByID(c)).Errors(db.ErrNotFound)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/profile/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	var errorResponse utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
	require.NoError(t, err)
	assert.Equal(t, utils.UserNotFound, errorResponse.Code)
}

func TestGetUserProfilePicture(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	user := BuildTestUser("1")
	image := BuildTestImage("1", "data:image/jpeg;base64,b3Vp")
	user.InnerUser.ProfilePictureID = &image.ID
	m.User.Expect(database.MockGetUserByID(c)).Returns(user)
	m.Image.Expect(database.MockGetImageByID(c)).Returns(image)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/user/1/picture/", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var imageResponse models.ImageResponse
	err := json.Unmarshal(w.Body.Bytes(), &imageResponse)
	require.NoError(t, err)
	assert.Equal(t, image.ID, imageResponse.ID)
}

func TestGetUserProfilePicture_NoContent(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	user := BuildTestUser("1")
	m.User.Expect(database.MockGetUserByID(c)).Returns(user)

	r := router.TestRoutes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/user/1/picture/", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestGetUserProfilePicture_UserNotFound(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.User.Expect(database.MockGetUserByID(c)).Errors(db.ErrNotFound)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/user/1/picture/", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	var errorResponse utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
	require.NoError(t, err)
	assert.Equal(t, utils.UserNotFound, errorResponse.Code)
}

func TestGetUserProfilePicture_ImageNotFound(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	user := BuildTestUser("1")
	user.InnerUser.ProfilePictureID = utils.Ptr("1")
	m.User.Expect(database.MockGetUserByID(c)).Returns(user)
	m.Image.Expect(database.MockGetImageByID(c)).Errors(db.ErrNotFound)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/user/1/picture/", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	var errorResponse utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
	require.NoError(t, err)
	assert.Equal(t, utils.UserProfilePictureNotFound, errorResponse.Code)
}

func TestUpdateCurrentUserProfile(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	user := BuildTestUser("1")
	updatedUser := user
	updatedUser.Firstname = "Updated"
	reqBody := models.UserUpdateRequest{
		Firstname: utils.Ptr("Updated"),
	}
	m.User.Expect(database.MockGetUserByID(c)).Returns(user)
	m.User.Expect(database.MockUpdateUser(c, reqBody)).Returns(updatedUser)

	b, err := json.Marshal(reqBody)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/v1/profile/", bytes.NewReader(b))
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
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.User.Expect(database.MockGetUserByID(c)).Errors(db.ErrNotFound)
	reqBody := models.UserUpdateRequest{
		Firstname: utils.Ptr("Updated"),
	}
	b, err := json.Marshal(reqBody)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/v1/profile/", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	var errorResponse utils.Error
	err = json.Unmarshal(w.Body.Bytes(), &errorResponse)
	require.NoError(t, err)
	assert.Equal(t, utils.UserNotFound, errorResponse.Code)
}

func TestUpdateCurrentUserProfile_MissingFields(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	user := BuildTestUser("1")
	m.User.Expect(database.MockGetUserByID(c)).Returns(user)

	reqBody := models.UserUpdateRequest{
		Email: utils.Ptr(""),
	}
	b, err := json.Marshal(reqBody)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/v1/profile/", bytes.NewReader(b))
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
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	user := BuildTestUser("1")
	reqBody := models.UserUpdateRequest{
		Email: utils.Ptr("existing@example.com"),
	}
	m.User.Expect(database.MockGetUserByID(c)).Returns(user)
	m.User.Expect(database.MockUpdateUser(c, reqBody)).Errors(&protocol.UserFacingError{
		IsPanic:   false,
		ErrorCode: "P2002", // https://www.prisma.io/docs/orm/reference/error-reference#p2002
		Meta: protocol.Meta{
			Target: []any{"email"},
		},
		Message: "Unique constraint failed",
	})

	b, err := json.Marshal(reqBody)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/v1/profile/", bytes.NewReader(b))
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
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	user := BuildTestUser("1")
	image := BuildTestImage("1", "data:image/jpeg;base64,b3Vp")
	updatedUser := user
	updatedUser.InnerUser.ProfilePictureID = &image.ID
	m.User.Expect(database.MockGetUserByID(c)).Returns(user)
	m.Image.Expect(database.MockCreateImage(c, image)).Returns(image)
	m.User.Expect(database.MockUpdateUserPicture(c)).Returns(updatedUser)

	reqBody := models.ImageRequest{
		Data: "data:image/jpeg;base64,b3Vp",
	}
	b, err := json.Marshal(reqBody)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/v1/profile/picture/", bytes.NewReader(b))
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
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.User.Expect(database.MockGetUserByID(c)).Errors(db.ErrNotFound)

	reqBody := models.ImageRequest{
		Data: "data:image/jpeg;base64,b3Vp",
	}
	b, err := json.Marshal(reqBody)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/v1/profile/picture/", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	var errorResponse utils.Error
	err = json.Unmarshal(w.Body.Bytes(), &errorResponse)
	require.NoError(t, err)
	assert.Equal(t, utils.UserNotFound, errorResponse.Code)
}

func TestUpdateCurrentUserProfilePicture_MissingFields(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	user := BuildTestUser("1")
	m.User.Expect(database.MockGetUserByID(c)).Returns(user)

	reqBody := models.ImageRequest{}
	b, err := json.Marshal(reqBody)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/v1/profile/picture/", bytes.NewReader(b))
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
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	user := BuildTestUser("1")
	m.User.Expect(database.MockGetUserByID(c)).Returns(user)

	reqBody := models.ImageRequest{
		Data: "invalid_base64",
	}
	b, err := json.Marshal(reqBody)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/v1/profile/picture/", bytes.NewReader(b))
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
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	user := BuildTestUser("1")
	image := BuildTestImage("1", "data:image/jpeg;base64,b3Vp")
	user.InnerUser.ProfilePictureID = &image.ID
	m.User.Expect(database.MockGetUserByID(c)).Returns(user)
	m.Image.Expect(database.MockGetImageByID(c)).Returns(image)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/profile/picture/", nil)
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
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	user := BuildTestUser("1")
	m.User.Expect(database.MockGetUserByID(c)).Returns(user)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/profile/picture/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestGetCurrentUserProfilePicture_UserNotFound(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.User.Expect(database.MockGetUserByID(c)).Errors(db.ErrNotFound)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/profile/picture/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	var errorResponse utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
	require.NoError(t, err)
	assert.Equal(t, utils.UserNotFound, errorResponse.Code)
}

func TestGetCurrentUserProfilePicture_ImageNotFound(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	user := BuildTestUser("1")
	user.InnerUser.ProfilePictureID = utils.Ptr("1")
	m.User.Expect(database.MockGetUserByID(c)).Returns(user)
	m.Image.Expect(database.MockGetImageByID(c)).Errors(db.ErrNotFound)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/profile/picture/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	var errorResponse utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
	require.NoError(t, err)
	assert.Equal(t, utils.UserProfilePictureNotFound, errorResponse.Code)
}
