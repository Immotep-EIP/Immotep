package models_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"immotep/backend/models"
	"immotep/backend/prisma/db"
	"immotep/backend/utils"
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
			CreatedAt:           db.DateTime{},
			OwnerID:             "1",
		},
		RelationsProperty: db.RelationsProperty{
			Leases: []db.LeaseModel{{
				InnerLease: db.InnerLease{
					ID:        "1",
					Active:    true,
					StartDate: db.DateTime{},
					EndDate:   nil,
				},
				RelationsLease: db.RelationsLease{
					Tenant: &db.UserModel{
						InnerUser: db.InnerUser{
							Firstname: "Test",
							Lastname:  "Name",
						},
					},
					Damages: []db.DamageModel{{
						InnerDamage: db.InnerDamage{
							ID:      "1",
							FixedAt: nil,
						}},
					},
				},
			}},
		},
	}
}

func BuildTestPropertyWithInventory(id string) db.PropertyModel {
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
			CreatedAt:           db.DateTime{},
			OwnerID:             "1",
			PictureID:           utils.Ptr("1"),
			Archived:            false,
		},
		RelationsProperty: db.RelationsProperty{
			Leases: []db.LeaseModel{{
				InnerLease: db.InnerLease{
					ID:        "1",
					Active:    true,
					StartDate: db.DateTime{},
					EndDate:   nil,
				},
				RelationsLease: db.RelationsLease{
					Tenant: &db.UserModel{
						InnerUser: db.InnerUser{
							Firstname: "Test",
							Lastname:  "Name",
						},
					},
					Damages: []db.DamageModel{{
						InnerDamage: db.InnerDamage{
							ID:      "1",
							FixedAt: nil,
						}},
					},
				},
			}},
			Rooms: []db.RoomModel{
				{
					InnerRoom: db.InnerRoom{
						ID:         "1",
						Name:       "Test",
						Archived:   false,
						PropertyID: id,
					},
					RelationsRoom: db.RelationsRoom{
						Furnitures: []db.FurnitureModel{{}},
					},
				},
			},
		},
	}
}

func TestPropertyResponse(t *testing.T) {
	pc := BuildTestProperty("1")

	t.Run("FromProperty1", func(t *testing.T) {
		propertyResponse := models.PropertyResponse{}
		propertyResponse.FromDbProperty(pc)

		assert.Equal(t, pc.ID, propertyResponse.ID)
		assert.Equal(t, pc.Name, propertyResponse.Name)
		assert.Equal(t, pc.Address, propertyResponse.Address)
		assert.Equal(t, pc.City, propertyResponse.City)
		assert.Equal(t, pc.PostalCode, propertyResponse.PostalCode)
		assert.Equal(t, pc.Country, propertyResponse.Country)
		assert.Equal(t, pc.OwnerID, propertyResponse.OwnerID)

		assert.Equal(t, models.StatusUnavailable, propertyResponse.Status)
		assert.Equal(t, 1, propertyResponse.NbDamage)
		assert.Equal(t, "Test Name", propertyResponse.Tenant)
		assert.Equal(t, propertyResponse.StartDate, utils.Ptr(db.DateTime{}))
		assert.Nil(t, propertyResponse.EndDate)
	})

	t.Run("FromProperty2", func(t *testing.T) {
		newPc := BuildTestProperty("2")
		newPc.RelationsProperty.Leases[0].Active = false

		propertyResponse := models.PropertyResponse{}
		propertyResponse.FromDbProperty(newPc)

		assert.Equal(t, newPc.ID, propertyResponse.ID)
		assert.Equal(t, newPc.Name, propertyResponse.Name)
		assert.Equal(t, newPc.Address, propertyResponse.Address)
		assert.Equal(t, newPc.City, propertyResponse.City)
		assert.Equal(t, newPc.PostalCode, propertyResponse.PostalCode)
		assert.Equal(t, newPc.Country, propertyResponse.Country)
		assert.Equal(t, newPc.OwnerID, propertyResponse.OwnerID)

		assert.Equal(t, models.StatusAvailable, propertyResponse.Status)
		assert.Equal(t, 1, propertyResponse.NbDamage)
		assert.Empty(t, propertyResponse.Tenant)
		assert.Nil(t, propertyResponse.StartDate)
		assert.Nil(t, propertyResponse.EndDate)
	})

	t.Run("FromProperty3", func(t *testing.T) {
		newPc := BuildTestProperty("3")
		newPc.RelationsProperty.Leases[0].Active = false
		newPc.RelationsProperty.LeaseInvite = &db.LeaseInviteModel{
			InnerLeaseInvite: db.InnerLeaseInvite{
				TenantEmail: "test@example.com",
				StartDate:   db.DateTime{},
				EndDate:     nil,
			},
		}

		propertyResponse := models.PropertyResponse{}
		propertyResponse.FromDbProperty(newPc)

		assert.Equal(t, newPc.ID, propertyResponse.ID)
		assert.Equal(t, newPc.Name, propertyResponse.Name)
		assert.Equal(t, newPc.Address, propertyResponse.Address)
		assert.Equal(t, newPc.City, propertyResponse.City)
		assert.Equal(t, newPc.PostalCode, propertyResponse.PostalCode)
		assert.Equal(t, newPc.Country, propertyResponse.Country)
		assert.Equal(t, newPc.OwnerID, propertyResponse.OwnerID)

		assert.Equal(t, models.StatusInviteSent, propertyResponse.Status)
		assert.Equal(t, 1, propertyResponse.NbDamage)
		assert.Equal(t, "test@example.com", propertyResponse.Tenant)
		assert.NotNil(t, propertyResponse.StartDate)
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

		assert.Equal(t, models.StatusUnavailable, propertyResponse.Status)
		assert.Equal(t, 1, propertyResponse.NbDamage)
		assert.Equal(t, "Test Name", propertyResponse.Tenant)
		assert.Equal(t, propertyResponse.StartDate, utils.Ptr(db.DateTime{}))
		assert.Nil(t, propertyResponse.EndDate)
	})
}

func TestPropertyInventoryResponse(t *testing.T) {
	pc := BuildTestPropertyWithInventory("1")

	t.Run("FromProperty1", func(t *testing.T) {
		propertyResponse := models.PropertyInventoryResponse{}
		propertyResponse.FromDbProperty(pc)

		assert.Equal(t, pc.ID, propertyResponse.ID)
		assert.Equal(t, pc.Name, propertyResponse.Name)
		assert.Equal(t, pc.Address, propertyResponse.Address)
		assert.Equal(t, pc.City, propertyResponse.City)
		assert.Equal(t, pc.PostalCode, propertyResponse.PostalCode)
		assert.Equal(t, pc.Country, propertyResponse.Country)
		assert.Equal(t, pc.OwnerID, propertyResponse.OwnerID)

		assert.Equal(t, models.StatusUnavailable, propertyResponse.Status)
		assert.Equal(t, 1, propertyResponse.NbDamage)
		assert.Equal(t, "Test Name", propertyResponse.Tenant)
		assert.Equal(t, propertyResponse.StartDate, utils.Ptr(db.DateTime{}))
		assert.Nil(t, propertyResponse.EndDate)

		assert.Equal(t, pc.InnerProperty.PictureID, propertyResponse.PictureID)
		assert.Equal(t, pc.Archived, propertyResponse.Archived)
		assert.Len(t, propertyResponse.Rooms, 1)
		assert.Len(t, propertyResponse.Rooms[0].Furnitures, 1)
	})

	t.Run("FromProperty2", func(t *testing.T) {
		newPc := BuildTestPropertyWithInventory("2")
		newPc.RelationsProperty.Leases[0].Active = false

		propertyResponse := models.PropertyInventoryResponse{}
		propertyResponse.FromDbProperty(newPc)

		assert.Equal(t, newPc.ID, propertyResponse.ID)
		assert.Equal(t, newPc.Name, propertyResponse.Name)
		assert.Equal(t, newPc.Address, propertyResponse.Address)
		assert.Equal(t, newPc.City, propertyResponse.City)
		assert.Equal(t, newPc.PostalCode, propertyResponse.PostalCode)
		assert.Equal(t, newPc.Country, propertyResponse.Country)
		assert.Equal(t, newPc.OwnerID, propertyResponse.OwnerID)

		assert.Equal(t, models.StatusAvailable, propertyResponse.Status)
		assert.Equal(t, 1, propertyResponse.NbDamage)
		assert.Empty(t, propertyResponse.Tenant)
		assert.Nil(t, propertyResponse.StartDate)
		assert.Nil(t, propertyResponse.EndDate)

		assert.Equal(t, newPc.InnerProperty.PictureID, propertyResponse.PictureID)
		assert.Equal(t, newPc.Archived, propertyResponse.Archived)
		assert.Len(t, propertyResponse.Rooms, 1)
		assert.Len(t, propertyResponse.Rooms[0].Furnitures, 1)
	})

	t.Run("FromProperty3", func(t *testing.T) {
		newPc := BuildTestPropertyWithInventory("3")
		newPc.RelationsProperty.Leases[0].Active = false
		newPc.RelationsProperty.LeaseInvite = &db.LeaseInviteModel{
			InnerLeaseInvite: db.InnerLeaseInvite{
				TenantEmail: "test@example.com",
				StartDate:   db.DateTime{},
				EndDate:     nil,
			},
		}

		propertyResponse := models.PropertyInventoryResponse{}
		propertyResponse.FromDbProperty(newPc)

		assert.Equal(t, newPc.ID, propertyResponse.ID)
		assert.Equal(t, newPc.Name, propertyResponse.Name)
		assert.Equal(t, newPc.Address, propertyResponse.Address)
		assert.Equal(t, newPc.City, propertyResponse.City)
		assert.Equal(t, newPc.PostalCode, propertyResponse.PostalCode)
		assert.Equal(t, newPc.Country, propertyResponse.Country)
		assert.Equal(t, newPc.OwnerID, propertyResponse.OwnerID)

		assert.Equal(t, models.StatusInviteSent, propertyResponse.Status)
		assert.Equal(t, 1, propertyResponse.NbDamage)
		assert.Equal(t, "test@example.com", propertyResponse.Tenant)
		assert.NotNil(t, propertyResponse.StartDate)
		assert.Nil(t, propertyResponse.EndDate)

		assert.Equal(t, newPc.InnerProperty.PictureID, propertyResponse.PictureID)
		assert.Equal(t, newPc.Archived, propertyResponse.Archived)
		assert.Len(t, propertyResponse.Rooms, 1)
		assert.Len(t, propertyResponse.Rooms[0].Furnitures, 1)
	})

	t.Run("PropertyToResponse", func(t *testing.T) {
		propertyResponse := models.DbPropertyInventoryToResponse(pc)

		assert.Equal(t, pc.ID, propertyResponse.ID)
		assert.Equal(t, pc.Name, propertyResponse.Name)
		assert.Equal(t, pc.Address, propertyResponse.Address)
		assert.Equal(t, pc.City, propertyResponse.City)
		assert.Equal(t, pc.PostalCode, propertyResponse.PostalCode)
		assert.Equal(t, pc.Country, propertyResponse.Country)
		assert.Equal(t, pc.OwnerID, propertyResponse.OwnerID)

		assert.Equal(t, models.StatusUnavailable, propertyResponse.Status)
		assert.Equal(t, 1, propertyResponse.NbDamage)
		assert.Equal(t, "Test Name", propertyResponse.Tenant)
		assert.Equal(t, propertyResponse.StartDate, utils.Ptr(db.DateTime{}))
		assert.Nil(t, propertyResponse.EndDate)

		assert.Equal(t, pc.InnerProperty.PictureID, propertyResponse.PictureID)
		assert.Equal(t, pc.Archived, propertyResponse.Archived)
		assert.Len(t, propertyResponse.Rooms, 1)
		assert.Len(t, propertyResponse.Rooms[0].Furnitures, 1)
	})
}
