package database_test

import (
	"errors"
	"testing"

	"github.com/steebchen/prisma-client-go/engine/protocol"
	"github.com/stretchr/testify/assert"
	"keyz/backend/models"
	"keyz/backend/prisma/db"
	"keyz/backend/services"
	"keyz/backend/services/database"
	"keyz/backend/utils"
)

func BuildTestUser(id string) db.UserModel {
	return db.UserModel{
		InnerUser: db.InnerUser{
			ID:        id,
			Email:     "test@example.com",
			Firstname: "Test",
			Lastname:  "User",
			Password:  "Password123",
		},
	}
}

// #############################################################################

func TestGetAllUsers(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	user := BuildTestUser("1")
	m.User.Expect(database.MockGetAllUsers(c)).ReturnsMany([]db.UserModel{user})

	allUsers := database.GetAllUsers()
	assert.Len(t, allUsers, 1)
	assert.Equal(t, user.ID, allUsers[0].ID)
}

func TestGetAllUsers_MultipleUsers(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	user1 := BuildTestUser("1")
	user2 := BuildTestUser("2")
	m.User.Expect(database.MockGetAllUsers(c)).ReturnsMany([]db.UserModel{user1, user2})

	allUsers := database.GetAllUsers()
	assert.Len(t, allUsers, 2)
	assert.Equal(t, user1.ID, allUsers[0].ID)
	assert.Equal(t, user2.ID, allUsers[1].ID)
}

func TestGetAllUsers_NoUsers(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.User.Expect(database.MockGetAllUsers(c)).ReturnsMany([]db.UserModel{})

	allUsers := database.GetAllUsers()
	assert.Empty(t, allUsers)
}

func TestGetAllUsers_NoConnection(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.User.Expect(database.MockGetAllUsers(c)).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.GetAllUsers()
	})
}

// #############################################################################

func TestGetUserByID(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	user := BuildTestUser("1")
	m.User.Expect(database.MockGetUserByID(c)).Returns(user)

	foundUser := database.GetUserByID("1")
	assert.NotNil(t, foundUser)
	assert.Equal(t, user.ID, foundUser.ID)
}

func TestGetUserByID_NotFound(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.User.Expect(database.MockGetUserByID(c)).Errors(db.ErrNotFound)

	foundUser := database.GetUserByID("1")
	assert.Nil(t, foundUser)
}

func TestGetUserByID_NoConnection(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.User.Expect(database.MockGetUserByID(c)).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.GetUserByID("1")
	})
}

// #############################################################################

func TestCreateUser(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	user := BuildTestUser("1")
	m.User.Expect(database.MockCreateUser(c, user)).Returns(user)

	newUser := database.CreateUser(user, db.RoleOwner)
	assert.NotNil(t, newUser)
	assert.Equal(t, user.ID, newUser.ID)
}

func TestCreateUser_DuplicateEmail(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	user := BuildTestUser("1")
	m.User.Expect(database.MockCreateUser(c, user)).Errors(&protocol.UserFacingError{
		IsPanic:   false,
		ErrorCode: "P2002", // https://www.prisma.io/docs/orm/reference/error-reference
		Meta: protocol.Meta{
			Target: []any{"email"},
		},
		Message: "Unique constraint failed",
	})

	newUser := database.CreateUser(user, db.RoleOwner)
	assert.Nil(t, newUser)
}

func TestCreateUser_NoConnection(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	user := BuildTestUser("1")
	m.User.Expect(database.MockCreateUser(c, user)).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.CreateUser(user, db.RoleOwner)
	})
}

// #############################################################################

func TestUpdateUser(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	user := BuildTestUser("1")
	updateRequest := models.UserUpdateRequest{
		Email:     utils.Ptr("updated@example.com"),
		Firstname: utils.Ptr("Updated"),
		Lastname:  utils.Ptr("User"),
	}
	m.User.Expect(database.MockUpdateUser(c, updateRequest)).Returns(user)

	updatedUser := database.UpdateUser(user, updateRequest)
	assert.NotNil(t, updatedUser)
	assert.Equal(t, user.ID, updatedUser.ID)
}

func TestUpdateUser_NotFound(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	updateRequest := models.UserUpdateRequest{
		Email:     utils.Ptr("updated@example.com"),
		Firstname: utils.Ptr("Updated"),
		Lastname:  utils.Ptr("User"),
	}
	m.User.Expect(database.MockUpdateUser(c, updateRequest)).Errors(db.ErrNotFound)

	assert.Panics(t, func() {
		database.UpdateUser(BuildTestUser("1"), updateRequest)
	})
}

func TestUpdateUser_DuplicateEmail(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	user := BuildTestUser("1")
	updateRequest := models.UserUpdateRequest{
		Email:     utils.Ptr("duplicate@example.com"),
		Firstname: utils.Ptr("Updated"),
		Lastname:  utils.Ptr("User"),
	}
	m.User.Expect(database.MockUpdateUser(c, updateRequest)).Errors(&protocol.UserFacingError{
		IsPanic:   false,
		ErrorCode: "P2002",
		Meta: protocol.Meta{
			Target: []any{"email"},
		},
		Message: "Unique constraint failed",
	})

	updatedUser := database.UpdateUser(user, updateRequest)
	assert.Nil(t, updatedUser)
}

func TestUpdateUser_NoConnection(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	updateRequest := models.UserUpdateRequest{
		Email:     utils.Ptr("updated@example.com"),
		Firstname: utils.Ptr("Updated"),
		Lastname:  utils.Ptr("User"),
	}
	m.User.Expect(database.MockUpdateUser(c, updateRequest)).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.UpdateUser(BuildTestUser("1"), updateRequest)
	})
}

// #############################################################################

func TestUpdateUserPicture(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	user := BuildTestUser("1")
	image := BuildTestImage("1", "data:image/jpeg;base64,b3Vp")

	m.User.Expect(database.MockUpdateUserPicture(c)).Returns(user)

	updatedUser := database.UpdateUserPicture(user, image)
	assert.NotNil(t, updatedUser)
	assert.Equal(t, user.ID, updatedUser.ID)
}

func TestUpdateUserPicture_NotFound(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	user := BuildTestUser("1")
	image := BuildTestImage("1", "data:image/jpeg;base64,b3Vp")

	m.User.Expect(database.MockUpdateUserPicture(c)).Errors(db.ErrNotFound)

	updatedUser := database.UpdateUserPicture(user, image)
	assert.Nil(t, updatedUser)
}

func TestUpdateUserPicture_NoConnection(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	user := BuildTestUser("1")
	image := BuildTestImage("1", "data:image/jpeg;base64,b3Vp")

	m.User.Expect(database.MockUpdateUserPicture(c)).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.UpdateUserPicture(user, image)
	})
}

// #############################################################################

func TestGetUserByEmail(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	user := BuildTestUser("1")
	m.User.Expect(database.MockGetUserByEmail(c)).Returns(user)

	foundUser := database.GetUserByEmail(user.Email)
	assert.NotNil(t, foundUser)
	assert.Equal(t, user.ID, foundUser.ID)
}

func TestGetUserByEmail_NotFound(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.User.Expect(database.MockGetUserByEmail(c)).Errors(db.ErrNotFound)

	foundUser := database.GetUserByEmail("test@example.com")
	assert.Nil(t, foundUser)
}

func TestGetUserByEmail_NoConnection(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.User.Expect(database.MockGetUserByEmail(c)).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.GetUserByEmail("test@example.com")
	})
}
