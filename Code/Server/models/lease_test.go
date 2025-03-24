package models_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"immotep/backend/models"
	"immotep/backend/prisma/db"
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

func TestInviteResponse(t *testing.T) {
	pc := db.LeaseInviteModel{
		InnerLeaseInvite: db.InnerLeaseInvite{
			ID:          "1",
			TenantEmail: "test1@example.com",
			StartDate:   time.Now(),
			PropertyID:  "1",
		},
	}

	t.Run("FromInvite", func(t *testing.T) {
		resp := models.InviteResponse{}
		resp.FromDbLeaseInvite(pc)

		assert.Equal(t, pc.ID, resp.ID)
		assert.Equal(t, pc.TenantEmail, resp.TenantEmail)
		assert.Equal(t, pc.StartDate, resp.StartDate)
		assert.Equal(t, pc.PropertyID, resp.PropertyID)
	})

	t.Run("InviteToResponse", func(t *testing.T) {
		resp := models.DbLeaseInviteToResponse(pc)

		assert.Equal(t, pc.ID, resp.ID)
		assert.Equal(t, pc.StartDate, resp.StartDate)
		assert.Equal(t, pc.PropertyID, resp.PropertyID)
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
