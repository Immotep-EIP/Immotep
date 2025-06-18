package models_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"keyz/backend/models"
)

func TestContactRequest(t *testing.T) {
	t.Run("ToDbContact", func(t *testing.T) {
		req := models.ContactMessageRequest{
			Firstname: "John",
			Lastname:  "Doe",
			Email:     "john.doe@example.com",
			Subject:   "Test Subject",
			Message:   "This is a test message.",
		}

		dbContact := req.ToDbContact()

		assert.Equal(t, req.Firstname, dbContact.Firstname)
		assert.Equal(t, req.Lastname, dbContact.Lastname)
		assert.Equal(t, req.Email, dbContact.Email)
		assert.Equal(t, req.Subject, dbContact.Subject)
		assert.Equal(t, req.Message, dbContact.Message)
	})
}
