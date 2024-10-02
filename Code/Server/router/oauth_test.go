package router_test

import (
	"immotep/backend/database"
	"immotep/backend/prisma/db"
	"immotep/backend/router"
	"testing"

	"github.com/stretchr/testify/assert"
)

func BuildTestUser(id string) db.UserModel {
	return db.UserModel{
		InnerUser: db.InnerUser{
			ID:        id,
			Email:     "test" + id + "@example.com",
			Firstname: "Test",
			Lastname:  "User",
			Password:  "$2a$14$BBhItuuxFbqV0rr0.r/reODEI78NEBnFIIK5W19qdybIYBvqNyyw.",
			Role:      db.RoleMember,
		},
	}
}

func TestValidateUser(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	testOauth := router.TestUserVerifier{}

	mock.User.Expect(
		client.Client.User.FindUnique(db.User.Email.Equals("test1@example.com")),
	).Returns(BuildTestUser("1"))

	t.Run("Valid user", func(t *testing.T) {
		err := testOauth.ValidateUser("test1@example.com", "Password123", "", nil)
		assert.Nil(t, err)
	})

	t.Run("Invalid password", func(t *testing.T) {
		err := testOauth.ValidateUser("test1@example.com", "azerty", "", nil)
		assert.NotNil(t, err)
	})

	mock.User.Expect(
		client.Client.User.FindUnique(db.User.Email.Equals("test2@example.com")),
	).Errors(db.ErrNotFound)

	t.Run("Not found user", func(t *testing.T) {
		err := testOauth.ValidateUser("test2@example.com", "Password123", "", nil)
		assert.NotNil(t, err)
	})
}
func TestAddClaims(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	testOauth := router.TestUserVerifier{}

	mock.User.Expect(
		client.Client.User.FindUnique(db.User.Email.Equals("test1@example.com")),
	).Returns(BuildTestUser("1"))

	t.Run("Valid user", func(t *testing.T) {
		claims, err := testOauth.AddClaims("test1@example.com", "", "", "")
		assert.Nil(t, err)
		assert.NotNil(t, claims)
		assert.Equal(t, "1", claims["id"])
		assert.Equal(t, "member", claims["role"])
	})

	mock.User.Expect(
		client.Client.User.FindUnique(db.User.Email.Equals("test2@example.com")),
	).Errors(db.ErrNotFound)

	t.Run("Not found user", func(t *testing.T) {
		claims, err := testOauth.AddClaims("test2@example.com", "", "", "")
		assert.NotNil(t, err)
		assert.Nil(t, claims)
	})
}
