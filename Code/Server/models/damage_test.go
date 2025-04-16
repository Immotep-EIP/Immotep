package models_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"immotep/backend/models"
	"immotep/backend/prisma/db"
)

func TestDamageRequest(t *testing.T) {
	t.Run("ToDbDamage", func(t *testing.T) {
		req := models.DamageRequest{
			Comment:  "Test Comment",
			Priority: db.PriorityHigh,
			RoomID:   "room123",
			Pictures: []string{"base64image1", "base64image2"},
		}

		dbDamage := req.ToDbDamage()

		assert.Equal(t, req.Comment, dbDamage.Comment)
		assert.Equal(t, req.Priority, dbDamage.Priority)
		assert.Equal(t, req.RoomID, dbDamage.RoomID)
	})
}

func BuildTestDamage(id string) db.DamageModel {
	return db.DamageModel{
		InnerDamage: db.InnerDamage{
			ID:           id,
			LeaseID:      "1",
			RoomID:       "1",
			Comment:      "Test Comment",
			Priority:     db.PriorityHigh,
			Read:         true,
			FixedOwner:   false,
			FixedTenant:  false,
			FixPlannedAt: nil,
			FixedAt:      nil,
		},
		RelationsDamage: db.RelationsDamage{
			Lease: &db.LeaseModel{
				RelationsLease: db.RelationsLease{
					Tenant: &db.UserModel{
						InnerUser: db.InnerUser{
							Firstname: "John",
							Lastname:  "Doe",
						},
					},
				},
			},
			Room: &db.RoomModel{
				InnerRoom: db.InnerRoom{
					Name: "Living Room",
				},
			},
			Pictures: []db.ImageModel{
				{
					InnerImage: db.InnerImage{
						ID:   "1",
						Data: db.Bytes("base64image1"),
					},
				},
			},
		},
	}
}

func TestDamageResponse(t *testing.T) {
	t.Run("FromDbDamage", func(t *testing.T) {
		mockDamageModel := BuildTestDamage("1")

		resp := models.DamageResponse{}
		resp.FromDbDamage(mockDamageModel)

		assert.Equal(t, mockDamageModel.ID, resp.ID)
		assert.Equal(t, mockDamageModel.LeaseID, resp.LeaseID)
		assert.Equal(t, "John Doe", resp.TenantName)
		assert.Equal(t, mockDamageModel.RoomID, resp.RoomID)
		assert.Equal(t, "Living Room", resp.RoomName)
		assert.Equal(t, mockDamageModel.Comment, resp.Comment)
		assert.Equal(t, mockDamageModel.Priority, resp.Priority)
		assert.Equal(t, mockDamageModel.Read, resp.Read)
		assert.Equal(t, mockDamageModel.CreatedAt, resp.CreatedAt)
		assert.Equal(t, mockDamageModel.UpdatedAt, resp.UpdatedAt)
		assert.Nil(t, resp.FixPlannedAt)
		assert.Nil(t, resp.FixedAt)
		assert.Len(t, resp.Pictures, 1)
	})

	t.Run("DbDamageToResponse", func(t *testing.T) {
		mockDamageModel := BuildTestDamage("1")

		resp := models.DbDamageToResponse(mockDamageModel)

		assert.Equal(t, mockDamageModel.ID, resp.ID)
		assert.Equal(t, mockDamageModel.LeaseID, resp.LeaseID)
		assert.Equal(t, "John Doe", resp.TenantName)
		assert.Equal(t, mockDamageModel.RoomID, resp.RoomID)
		assert.Equal(t, "Living Room", resp.RoomName)
		assert.Equal(t, mockDamageModel.Comment, resp.Comment)
		assert.Equal(t, mockDamageModel.Priority, resp.Priority)
		assert.Equal(t, mockDamageModel.Read, resp.Read)
		assert.Equal(t, mockDamageModel.CreatedAt, resp.CreatedAt)
		assert.Equal(t, mockDamageModel.UpdatedAt, resp.UpdatedAt)
		assert.Nil(t, resp.FixPlannedAt)
		assert.Nil(t, resp.FixedAt)
		assert.Len(t, resp.Pictures, 1)
	})
}
