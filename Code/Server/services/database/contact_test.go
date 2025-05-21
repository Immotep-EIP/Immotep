package database_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"immotep/backend/prisma/db"
	"immotep/backend/services"
	"immotep/backend/services/database"
)

func BuildTestContactMessage(id string) db.ContactMessageModel {
	return db.ContactMessageModel{
		InnerContactMessage: db.InnerContactMessage{
			ID:        id,
			Firstname: "John",
			Lastname:  "Doe",
			Email:     "john.doe@example.com",
			Subject:   "Test Subject",
			Message:   "This is a test message.",
		},
	}
}

func TestCreateContactMessage(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	contact := BuildTestContactMessage("1")
	m.ContactMessage.Expect(database.MockCreateContactMessage(c, contact)).Returns(contact)

	newContact := database.CreateContactMessage(contact)
	assert.NotNil(t, newContact)
	assert.Equal(t, contact.ID, newContact.ID)
	assert.Equal(t, contact.Firstname, newContact.Firstname)
	assert.Equal(t, contact.Lastname, newContact.Lastname)
	assert.Equal(t, contact.Email, newContact.Email)
	assert.Equal(t, contact.Subject, newContact.Subject)
	assert.Equal(t, contact.Message, newContact.Message)
}

func TestCreateContactMessage_PanicOnError(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	contact := BuildTestContactMessage("2")
	m.ContactMessage.Expect(database.MockCreateContactMessage(c, contact)).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.CreateContactMessage(contact)
	})
}
