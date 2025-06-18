package models_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"keyz/backend/models"
	"keyz/backend/prisma/db"
)

func TestInviteRequest(t *testing.T) {
	req := models.InviteRequest{
		TenantEmail: "test1@example.com",
		StartDate:   time.Now(),
	}

	t.Run("ToInvite", func(t *testing.T) {
		pc := req.ToDbLeaseInvite()

		assert.Equal(t, req.TenantEmail, pc.TenantEmail)
		assert.Equal(t, req.StartDate, pc.StartDate)
	})
}

func TestLeaseResponse(t *testing.T) {
	model := db.LeaseModel{
		InnerLease: db.InnerLease{
			ID:         "1",
			PropertyID: "101",
			TenantID:   "201",
			Active:     true,
			StartDate:  time.Now(),
			EndDate:    nil,
			CreatedAt:  time.Now(),
		},
		RelationsLease: db.RelationsLease{
			Tenant: &db.UserModel{
				InnerUser: db.InnerUser{
					Firstname: "John",
					Lastname:  "Doe",
					Email:     "johndoe@example.com",
				},
			},
			Property: &db.PropertyModel{
				InnerProperty: db.InnerProperty{
					ID:   "1",
					Name: "Test Property",
				},
				RelationsProperty: db.RelationsProperty{
					Owner: &db.UserModel{
						InnerUser: db.InnerUser{
							ID:        "1",
							Firstname: "John",
							Lastname:  "Doe",
							Email:     "johndoe@example.com",
						},
					},
				},
			},
		},
	}

	t.Run("FromDbLease", func(t *testing.T) {
		resp := models.LeaseResponse{}
		resp.FromDbLease(model)

		assert.Equal(t, model.ID, resp.ID)
		assert.Equal(t, model.PropertyID, resp.PropertyID)
		assert.Equal(t, model.TenantID, resp.TenantID)
		assert.Equal(t, model.Tenant().Name(), resp.TenantName)
		assert.Equal(t, model.Tenant().Email, resp.TenantEmail)
		assert.Equal(t, model.StartDate, resp.StartDate)
		assert.Equal(t, model.InnerLease.EndDate, resp.EndDate)
		assert.Equal(t, model.Active, resp.Active)
		assert.Equal(t, model.CreatedAt, resp.CreatedAt)
	})

	t.Run("DbLeaseToResponse", func(t *testing.T) {
		resp := models.DbLeaseToResponse(model)

		assert.Equal(t, model.ID, resp.ID)
		assert.Equal(t, model.PropertyID, resp.PropertyID)
		assert.Equal(t, model.TenantID, resp.TenantID)
		assert.Equal(t, model.Tenant().Name(), resp.TenantName)
		assert.Equal(t, model.Tenant().Email, resp.TenantEmail)
		assert.Equal(t, model.StartDate, resp.StartDate)
		assert.Equal(t, model.InnerLease.EndDate, resp.EndDate)
		assert.Equal(t, model.Active, resp.Active)
		assert.Equal(t, model.CreatedAt, resp.CreatedAt)
	})
}
