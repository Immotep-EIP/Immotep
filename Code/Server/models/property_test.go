package models_test

import (
	"testing"
	"time"

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

func BuildTestProperty(id string) db.PropertyModel {
	return db.PropertyModel{
		InnerProperty: db.InnerProperty{
			ID:                  id,
			Name:                "Test",
			Address:             "Test",
			City:                "Test",
			PostalCode:          "Test",
			Country:             "Test",
			AreaSqm:             20.0,
			RentalPricePerMonth: 500,
			DepositPrice:        1000,
			CreatedAt:           time.Now(),
			OwnerID:             "1",
		},
		RelationsProperty: db.RelationsProperty{
			Damages:   []db.DamageModel{{}},
			Contracts: []db.ContractModel{{}},
		},
	}
}

func TestPropertyResponse(t *testing.T) {
	pc := BuildTestProperty("1")

	t.Run("FromProperty", func(t *testing.T) {
		propertyResponse := models.PropertyResponse{}
		propertyResponse.FromDbProperty(pc)

		assert.Equal(t, pc.ID, propertyResponse.ID)
		assert.Equal(t, pc.Name, propertyResponse.Name)
		assert.Equal(t, pc.Address, propertyResponse.Address)
		assert.Equal(t, pc.City, propertyResponse.City)
		assert.Equal(t, pc.PostalCode, propertyResponse.PostalCode)
		assert.Equal(t, pc.Country, propertyResponse.Country)
		assert.Equal(t, pc.OwnerID, propertyResponse.OwnerID)

		assert.Equal(t, "available", propertyResponse.Status)
		assert.Equal(t, 1, propertyResponse.NbDamage)
		assert.Equal(t, "", propertyResponse.Tenant)
		assert.Nil(t, propertyResponse.StartDate)
		assert.Nil(t, propertyResponse.EndDate)
	})

	t.Run("PropertyToResponse", func(t *testing.T) {
		propertyResponse := models.DbPropertyToResponse(pc)

		assert.Equal(t, pc.ID, propertyResponse.ID)
		assert.Equal(t, pc.Name, propertyResponse.Name)
		assert.Equal(t, pc.Address, propertyResponse.Address)
		assert.Equal(t, pc.City, propertyResponse.City)
		assert.Equal(t, pc.PostalCode, propertyResponse.PostalCode)
		assert.Equal(t, pc.Country, propertyResponse.Country)
		assert.Equal(t, pc.OwnerID, propertyResponse.OwnerID)

		assert.Equal(t, "available", propertyResponse.Status)
		assert.Equal(t, 1, propertyResponse.NbDamage)
		assert.Equal(t, "", propertyResponse.Tenant)
		assert.Nil(t, propertyResponse.StartDate)
		assert.Nil(t, propertyResponse.EndDate)
	})
}
