package router_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"immotep/backend/prisma/db"
	"immotep/backend/router"
	"immotep/backend/services"
	"immotep/backend/services/database"
)

func BuildTestUser(id string) db.UserModel {
	return db.UserModel{
		InnerUser: db.InnerUser{
			ID:        id,
			Email:     "test@example.com",
			Firstname: "Test",
			Lastname:  "User",
			Password:  "$2a$14$BBhItuuxFbqV0rr0.r/reODEI78NEBnFIIK5W19qdybIYBvqNyyw.",
			Role:      db.RoleOwner,
		},
	}
}

func TestValidateUser(t *testing.T) {
	testOauth := router.TestUserVerifier{}

	t.Run("Valid user", func(t *testing.T) {
		c, m, ensure := services.ConnectDBTest()
		defer ensure(t)

		m.User.Expect(database.MockGetUserByEmail(c)).Returns(BuildTestUser("1"))

		err := testOauth.ValidateUser("test@example.com", "Password123", "", nil)
		require.NoError(t, err)
	})

	t.Run("Invalid password", func(t *testing.T) {
		c, m, ensure := services.ConnectDBTest()
		defer ensure(t)

		m.User.Expect(database.MockGetUserByEmail(c)).Returns(BuildTestUser("1"))

		err := testOauth.ValidateUser("test@example.com", "azerty", "", nil)
		require.Error(t, err)
	})

	t.Run("Not found user", func(t *testing.T) {
		c, m, ensure := services.ConnectDBTest()
		defer ensure(t)

		m.User.Expect(database.MockGetUserByEmail(c)).Errors(db.ErrNotFound)

		err := testOauth.ValidateUser("test@example.com", "Password123", "", nil)
		require.Error(t, err)
	})
}

func TestAddClaims(t *testing.T) {
	testOauth := router.TestUserVerifier{}

	t.Run("Valid user", func(t *testing.T) {
		c, m, ensure := services.ConnectDBTest()
		defer ensure(t)

		m.User.Expect(database.MockGetUserByEmail(c)).Returns(BuildTestUser("1"))

		claims, err := testOauth.AddClaims("test@example.com", "", "", "")
		require.NoError(t, err)
		assert.NotNil(t, claims)
		assert.Equal(t, "1", claims["id"])
		assert.Equal(t, "owner", claims["role"])
	})

	t.Run("Not found user", func(t *testing.T) {
		c, m, ensure := services.ConnectDBTest()
		defer ensure(t)

		m.User.Expect(database.MockGetUserByEmail(c)).Errors(db.ErrNotFound)

		claims, err := testOauth.AddClaims("test@example.com", "", "", "")
		require.Error(t, err)
		assert.Nil(t, claims)
	})
}

func TestAddProperties(t *testing.T) {
	testOauth := router.TestUserVerifier{}

	t.Run("Add properties", func(t *testing.T) {
		c, m, ensure := services.ConnectDBTest()
		defer ensure(t)

		m.User.Expect(database.MockGetUserByEmail(c)).Returns(BuildTestUser("1"))

		props, err := testOauth.AddProperties("test@example.com", "", "", "")
		require.NoError(t, err)
		assert.NotNil(t, props)
		assert.Equal(t, "1", props["id"])
		assert.Equal(t, "owner", props["role"])
	})

	t.Run("Not found user", func(t *testing.T) {
		c, m, ensure := services.ConnectDBTest()
		defer ensure(t)

		m.User.Expect(database.MockGetUserByEmail(c)).Errors(db.ErrNotFound)

		props, err := testOauth.AddProperties("test@example.com", "", "", "")
		require.Error(t, err)
		assert.Nil(t, props)
	})
}

// Unused methods ----------------------------------------------------------------------------------------
func TestValidateClient(t *testing.T) {
	testOauth := router.TestUserVerifier{}

	t.Run("Validate c", func(t *testing.T) {
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
