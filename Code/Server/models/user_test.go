package models_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"keyz/backend/models"
	"keyz/backend/prisma/db"
)

func TestUserRequest(t *testing.T) {
	req := models.UserRequest{
		Email:     "test1@example.com",
		Firstname: "Test",
		Lastname:  "User",
		Password:  "Password123",
	}

	t.Run("ToUser", func(t *testing.T) {
		user := req.ToDbUser()

		assert.Equal(t, req.Email, user.Email)
		assert.Equal(t, req.Firstname, user.Firstname)
		assert.Equal(t, req.Lastname, user.Lastname)
		assert.Equal(t, req.Password, user.Password)
	})
}

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
		resp := models.UserResponse{}
		resp.FromDbUser(user)

		assert.Equal(t, user.ID, resp.ID)
		assert.Equal(t, user.Email, resp.Email)
		assert.Equal(t, user.Firstname, resp.Firstname)
		assert.Equal(t, user.Lastname, resp.Lastname)
		assert.Equal(t, user.Role, resp.Role)
		assert.Equal(t, user.CreatedAt, resp.CreatedAt)
		assert.Equal(t, user.UpdatedAt, resp.UpdatedAt)
	})

	t.Run("UserToResponse", func(t *testing.T) {
		resp := models.DbUserToResponse(user)

		assert.Equal(t, user.ID, resp.ID)
		assert.Equal(t, user.Email, resp.Email)
		assert.Equal(t, user.Firstname, resp.Firstname)
		assert.Equal(t, user.Lastname, resp.Lastname)
		assert.Equal(t, user.Role, resp.Role)
		assert.Equal(t, user.CreatedAt, resp.CreatedAt)
		assert.Equal(t, user.UpdatedAt, resp.UpdatedAt)
	})
}
