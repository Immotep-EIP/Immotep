package models_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"immotep/backend/models"
	"immotep/backend/prisma/db"
)

func TestPropertyRequest(t *testing.T) {
	req := models.PropertyRequest{
		Name:       "Test",
		Address:    "Test",
		City:       "Test",
		PostalCode: "Test",
		Country:    "Test",
	}

	t.Run("ToProperty", func(t *testing.T) {
		pc := req.ToDbProperty()

		assert.Equal(t, req.Name, pc.Name)
		assert.Equal(t, req.Address, pc.Address)
		assert.Equal(t, req.City, pc.City)
		assert.Equal(t, req.PostalCode, pc.PostalCode)
		assert.Equal(t, req.Country, pc.Country)
	})
}

func TestPropertyResponse(t *testing.T) {
	pc := db.PropertyModel{
		InnerProperty: db.InnerProperty{
			ID:         "1",
			Name:       "Test",
			Address:    "Test",
			City:       "Test",
			PostalCode: "Test",
			Country:    "Test",
			OwnerID:    "1",
		},
	}

	t.Run("FromProperty", func(t *testing.T) {
		inviteResponse := models.PropertyResponse{}
		inviteResponse.FromDbProperty(pc)

		assert.Equal(t, pc.ID, inviteResponse.ID)
		assert.Equal(t, pc.Name, inviteResponse.Name)
		assert.Equal(t, pc.Address, inviteResponse.Address)
		assert.Equal(t, pc.City, inviteResponse.City)
		assert.Equal(t, pc.PostalCode, inviteResponse.PostalCode)
		assert.Equal(t, pc.Country, inviteResponse.Country)
		assert.Equal(t, pc.OwnerID, inviteResponse.OwnerID)
	})

	t.Run("PropertyToResponse", func(t *testing.T) {
		inviteResponse := models.DbPropertyToResponse(pc)

		assert.Equal(t, pc.ID, inviteResponse.ID)
		assert.Equal(t, pc.Name, inviteResponse.Name)
		assert.Equal(t, pc.Address, inviteResponse.Address)
		assert.Equal(t, pc.City, inviteResponse.City)
		assert.Equal(t, pc.PostalCode, inviteResponse.PostalCode)
		assert.Equal(t, pc.Country, inviteResponse.Country)
		assert.Equal(t, pc.OwnerID, inviteResponse.OwnerID)
	})
}
