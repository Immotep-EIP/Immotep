package userservice_test

import (
	"errors"
	"testing"

	"github.com/steebchen/prisma-client-go/engine/protocol"
	"github.com/stretchr/testify/assert"
	"immotep/backend/database"
	"immotep/backend/models"
	"immotep/backend/prisma/db"
	userservice "immotep/backend/services/user"
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

func BuildTestImage(id string, base64data string) db.ImageModel {
	ret := models.StringToDbImage(base64data)
	if ret == nil {
		panic("Invalid base64 string")
	}
	ret.ID = id
	return *ret
}

func TestGetAllUsers(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	user := BuildTestUser("1")

	mock.User.Expect(
		client.Client.User.FindMany(),
	).ReturnsMany([]db.UserModel{user})

	allUsers := userservice.GetAll()
	assert.Len(t, allUsers, 1)
	assert.Equal(t, user.ID, allUsers[0].ID)
}

func TestGetAllUsers_MultipleUsers(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	user1 := BuildTestUser("1")
	user2 := BuildTestUser("2")

	mock.User.Expect(
		client.Client.User.FindMany(),
	).ReturnsMany([]db.UserModel{user1, user2})

	allUsers := userservice.GetAll()
	assert.Len(t, allUsers, 2)
	assert.Equal(t, user1.ID, allUsers[0].ID)
	assert.Equal(t, user2.ID, allUsers[1].ID)
}

func TestGetAllUsers_NoUsers(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	mock.User.Expect(
		client.Client.User.FindMany(),
	).ReturnsMany([]db.UserModel{})

	allUsers := userservice.GetAll()
	assert.Empty(t, allUsers)
}

func TestGetAllUsers_NoConnection(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	mock.User.Expect(
		client.Client.User.FindMany(),
	).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		userservice.GetAll()
	})
}

func TestGetUserByID(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	user := BuildTestUser("1")

	mock.User.Expect(
		client.Client.User.FindUnique(db.User.ID.Equals("1")),
	).Returns(user)

	foundUser := userservice.GetByID("1")
	assert.NotNil(t, foundUser)
	assert.Equal(t, user.ID, foundUser.ID)
}

func TestGetUserByID_NotFound(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	mock.User.Expect(
		client.Client.User.FindUnique(db.User.ID.Equals("1")),
	).Errors(db.ErrNotFound)

	foundUser := userservice.GetByID("1")
	assert.Nil(t, foundUser)
}

func TestGetUserByID_NoConnection(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	mock.User.Expect(
		client.Client.User.FindUnique(db.User.ID.Equals("1")),
	).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		userservice.GetByID("1")
	})
}

func TestCreateUser(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	user := BuildTestUser("1")

	mock.User.Expect(
		client.Client.User.CreateOne(
			db.User.Email.Set(user.Email),
			db.User.Password.Set(user.Password),
			db.User.Firstname.Set(user.Firstname),
			db.User.Lastname.Set(user.Lastname),
			db.User.Role.Set(db.RoleOwner),
		),
	).Returns(user)

	newUser := userservice.Create(user, db.RoleOwner)
	assert.NotNil(t, newUser)
	assert.Equal(t, user.ID, newUser.ID)
}

func TestCreateUser_DuplicateEmail(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	user := db.UserModel{
		InnerUser: db.InnerUser{
			Email:     "test@example.com",
			Firstname: "Test",
			Lastname:  "User",
			Password:  "Password123",
		},
	}

	mock.User.Expect(
		client.Client.User.CreateOne(
			db.User.Email.Set(user.Email),
			db.User.Password.Set(user.Password),
			db.User.Firstname.Set(user.Firstname),
			db.User.Lastname.Set(user.Lastname),
			db.User.Role.Set(db.RoleOwner),
		),
	).Errors(&protocol.UserFacingError{
		IsPanic:   false,
		ErrorCode: "P2002", // https://www.prisma.io/docs/orm/reference/error-reference
		Meta: protocol.Meta{
			Target: []any{"email"},
		},
		Message: "Unique constraint failed",
	})

	newUser := userservice.Create(user, db.RoleOwner)
	assert.Nil(t, newUser)
}

func TestCreateUser_NoConnection(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	user := BuildTestUser("1")

	mock.User.Expect(
		client.Client.User.CreateOne(
			db.User.Email.Set(user.Email),
			db.User.Password.Set(user.Password),
			db.User.Firstname.Set(user.Firstname),
			db.User.Lastname.Set(user.Lastname),
			db.User.Role.Set(db.RoleOwner),
		),
	).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		userservice.Create(user, db.RoleOwner)
	})
}

func TestUpdateUser(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	user := BuildTestUser("1")
	updateRequest := models.UserUpdateRequest{
		Email:     utils.Ptr("updated@example.com"),
		Firstname: utils.Ptr("Updated"),
		Lastname:  utils.Ptr("User"),
	}

	mock.User.Expect(
		client.Client.User.FindUnique(db.User.ID.Equals(user.ID)).Update(
			db.User.Email.SetIfPresent(updateRequest.Email),
			db.User.Firstname.SetIfPresent(updateRequest.Firstname),
			db.User.Lastname.SetIfPresent(updateRequest.Lastname),
		),
	).Returns(user)

	updatedUser := userservice.Update(user.ID, updateRequest)
	assert.NotNil(t, updatedUser)
	assert.Equal(t, user.ID, updatedUser.ID)
}

func TestUpdateUser_NotFound(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	updateRequest := models.UserUpdateRequest{
		Email:     utils.Ptr("updated@example.com"),
		Firstname: utils.Ptr("Updated"),
		Lastname:  utils.Ptr("User"),
	}

	mock.User.Expect(
		client.Client.User.FindUnique(db.User.ID.Equals("1")).Update(
			db.User.Email.SetIfPresent(updateRequest.Email),
			db.User.Firstname.SetIfPresent(updateRequest.Firstname),
			db.User.Lastname.SetIfPresent(updateRequest.Lastname),
		),
	).Errors(db.ErrNotFound)

	updatedUser := userservice.Update("1", updateRequest)
	assert.Nil(t, updatedUser)
}

func TestUpdateUser_DuplicateEmail(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	user := BuildTestUser("1")
	updateRequest := models.UserUpdateRequest{
		Email:     utils.Ptr("duplicate@example.com"),
		Firstname: utils.Ptr("Updated"),
		Lastname:  utils.Ptr("User"),
	}

	mock.User.Expect(
		client.Client.User.FindUnique(db.User.ID.Equals(user.ID)).Update(
			db.User.Email.SetIfPresent(updateRequest.Email),
			db.User.Firstname.SetIfPresent(updateRequest.Firstname),
			db.User.Lastname.SetIfPresent(updateRequest.Lastname),
		),
	).Errors(&protocol.UserFacingError{
		IsPanic:   false,
		ErrorCode: "P2002",
		Meta: protocol.Meta{
			Target: []any{"email"},
		},
		Message: "Unique constraint failed",
	})

	updatedUser := userservice.Update(user.ID, updateRequest)
	assert.Nil(t, updatedUser)
}

func TestUpdateUser_NoConnection(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	user := BuildTestUser("1")
	updateRequest := models.UserUpdateRequest{
		Email:     utils.Ptr("updated@example.com"),
		Firstname: utils.Ptr("Updated"),
		Lastname:  utils.Ptr("User"),
	}

	mock.User.Expect(
		client.Client.User.FindUnique(db.User.ID.Equals(user.ID)).Update(
			db.User.Email.SetIfPresent(updateRequest.Email),
			db.User.Firstname.SetIfPresent(updateRequest.Firstname),
			db.User.Lastname.SetIfPresent(updateRequest.Lastname),
		),
	).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		userservice.Update(user.ID, updateRequest)
	})
}

func TestUpdatePicture(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	user := BuildTestUser("1")
	image := BuildTestImage("1", "b3Vp")

	mock.User.Expect(
		client.Client.User.FindUnique(db.User.ID.Equals(user.ID)).Update(
			db.User.ProfilePicture.Link(db.Image.ID.Equals(image.ID)),
		),
	).Returns(user)

	updatedUser := userservice.UpdatePicture(user, image)
	assert.NotNil(t, updatedUser)
	assert.Equal(t, user.ID, updatedUser.ID)
}

func TestUpdatePicture_NotFound(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	user := BuildTestUser("1")
	image := BuildTestImage("1", "b3Vp")

	mock.User.Expect(
		client.Client.User.FindUnique(db.User.ID.Equals(user.ID)).Update(
			db.User.ProfilePicture.Link(db.Image.ID.Equals(image.ID)),
		),
	).Errors(db.ErrNotFound)

	updatedUser := userservice.UpdatePicture(user, image)
	assert.Nil(t, updatedUser)
}

func TestUpdatePicture_NoConnection(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	user := BuildTestUser("1")
	image := BuildTestImage("1", "b3Vp")

	mock.User.Expect(
		client.Client.User.FindUnique(db.User.ID.Equals(user.ID)).Update(
			db.User.ProfilePicture.Link(db.Image.ID.Equals(image.ID)),
		),
	).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		userservice.UpdatePicture(user, image)
	})
}
