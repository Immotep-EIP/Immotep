package models_test

import (
	"immotep/backend/models"
	"immotep/backend/prisma/db"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserResponse(t *testing.T) {
	user := db.UserModel{
		InnerUser: db.InnerUser{
			ID:        "1",
			Email:     "test1@example.com",
			Firstname: "Test",
			Lastname:  "User",
			Password:  "Password123",
		},
	}

	t.Run("FromUser", func(t *testing.T) {
		userResponse := models.UserResponse{}
		userResponse.FromUser(user)

		assert.Equal(t, user.ID, userResponse.ID)
		assert.Equal(t, user.Email, userResponse.Email)
		assert.Equal(t, user.Firstname, userResponse.Firstname)
		assert.Equal(t, user.Lastname, userResponse.Lastname)
		assert.Equal(t, user.Role, userResponse.Role)
		assert.Equal(t, user.CreatedAt, userResponse.CreatedAt)
		assert.Equal(t, user.UpdatedAt, userResponse.UpdatedAt)
	})

	t.Run("UserToResponse", func(t *testing.T) {
		userResponse := models.UserToResponse(user)	

		assert.Equal(t, user.ID, userResponse.ID)
		assert.Equal(t, user.Email, userResponse.Email)
		assert.Equal(t, user.Firstname, userResponse.Firstname)
		assert.Equal(t, user.Lastname, userResponse.Lastname)
		assert.Equal(t, user.Role, userResponse.Role)
		assert.Equal(t, user.CreatedAt, userResponse.CreatedAt)
		assert.Equal(t, user.UpdatedAt, userResponse.UpdatedAt)
	})
}
