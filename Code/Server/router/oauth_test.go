package router_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"immotep/backend/database"
	"immotep/backend/prisma/db"
	"immotep/backend/router"
)

func BuildTestUser(id string) db.UserModel {
	return db.UserModel{
		InnerUser: db.InnerUser{
			ID:        id,
			Email:     "test" + id + "@example.com",
			Firstname: "Test",
			Lastname:  "User",
			Password:  "$2a$14$BBhItuuxFbqV0rr0.r/reODEI78NEBnFIIK5W19qdybIYBvqNyyw.",
			Role:      db.RoleOwner,
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
		require.NoError(t, err)
	})

	t.Run("Invalid password", func(t *testing.T) {
		err := testOauth.ValidateUser("test1@example.com", "azerty", "", nil)
		require.Error(t, err)
	})

	mock.User.Expect(
		client.Client.User.FindUnique(db.User.Email.Equals("test2@example.com")),
	).Errors(db.ErrNotFound)

	t.Run("Not found user", func(t *testing.T) {
		err := testOauth.ValidateUser("test2@example.com", "Password123", "", nil)
		require.Error(t, err)
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
		require.NoError(t, err)
		assert.NotNil(t, claims)
		assert.Equal(t, "1", claims["id"])
		assert.Equal(t, "owner", claims["role"])
	})

	mock.User.Expect(
		client.Client.User.FindUnique(db.User.Email.Equals("test2@example.com")),
	).Errors(db.ErrNotFound)

	t.Run("Not found user", func(t *testing.T) {
		claims, err := testOauth.AddClaims("test2@example.com", "", "", "")
		require.Error(t, err)
		assert.Nil(t, claims)
	})
}

func TestAddProperties(t *testing.T) {
	testOauth := router.TestUserVerifier{}

	t.Run("Add properties", func(t *testing.T) {
		props, err := testOauth.AddProperties("test1@example.com", "", "", "")
		require.NoError(t, err)
		assert.NotNil(t, props)
	})
}

// Unused methods ----------------------------------------------------------------------------------------
func TestValidateClient(t *testing.T) {
	testOauth := router.TestUserVerifier{}

	t.Run("Validate client", func(t *testing.T) {
		err := testOauth.ValidateClient("", "", "", nil)
		require.Error(t, err)
	})
}

func TestValidateTokenId(t *testing.T) {
	testOauth := router.TestUserVerifier{}

	t.Run("Validate token id", func(t *testing.T) {
		err := testOauth.ValidateTokenId("", "", "", "")
		require.NoError(t, err)
	})
}

func TestStoreTokenId(t *testing.T) {
	testOauth := router.TestUserVerifier{}

	t.Run("Validate token id", func(t *testing.T) {
		err := testOauth.StoreTokenId("", "", "", "")
		require.NoError(t, err)
	})
}
