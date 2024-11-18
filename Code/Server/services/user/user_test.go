package userservice_test

import (
	"errors"
	"testing"

	"github.com/steebchen/prisma-client-go/engine/protocol"
	"github.com/stretchr/testify/assert"
	"immotep/backend/database"
	"immotep/backend/prisma/db"
	userservice "immotep/backend/services/user"
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
