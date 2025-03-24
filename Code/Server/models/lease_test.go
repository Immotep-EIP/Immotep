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
