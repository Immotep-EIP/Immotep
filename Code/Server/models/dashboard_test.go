package models_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"immotep/backend/models"
	"immotep/backend/prisma/db"
)

func TestReminderModel(t *testing.T) {
	t.Run("GetEn", func(t *testing.T) {
		rm := models.ReminderAllGood
		r := rm.Get("en")
		assert.Equal(t, "Good news!", r.Title)
	})

	t.Run("GetFr", func(t *testing.T) {
		rm := models.ReminderAllGood
		r := rm.Get("fr")
		assert.Equal(t, "Bonne nouvelle !", r.Title)
	})
}

func TestReminder(t *testing.T) {
	t.Run("Placeholder", func(t *testing.T) {
		r := models.ReminderLeaseEnding.Get("en")
		r2 := r.WithPlaceholders(map[string]string{
			"property": "Test Property",
			"days":     "30",
		})
		assert.Equal(t, "Lease of property Test Property is ending in 30 days.", r2.Title)
	})
}

func TestGetReminders(t *testing.T) {
	t.Run("LeaseEnding", func(t *testing.T) {
		r := models.GetReminderLeaseEnding("en", BuildTestProperty("1"), 30)
		assert.Equal(t, "Lease of property Test is ending in 30 days.", r.Title)
	})

	t.Run("NoInventoryReport", func(t *testing.T) {
		r := models.GetReminderNoInventoryReport("en", BuildTestProperty("1"))
		assert.Equal(t, "Lease of property Test started recently and does not have an inventory report.", r.Title)
	})

	t.Run("PropertyAvailable", func(t *testing.T) {
		r := models.GetReminderPropertyAvailable("en", BuildTestProperty("1"))
		assert.Equal(t, "Property Test is available for rent.", r.Title)
	})

	t.Run("EmptyInventory", func(t *testing.T) {
		r := models.GetReminderEmptyInventory("en", BuildTestProperty("1"))
		assert.Equal(t, "Inventory of property Test is empty.", r.Title)
	})

	t.Run("PendingLeaseInvitation", func(t *testing.T) {
		r := models.GetReminderPendingLeaseInvitation("en", BuildTestProperty("1"), 5)
		assert.Equal(t, "Lease invitation for property Test has been pending for 5 days.", r.Title)
	})

	t.Run("NewDamageReported", func(t *testing.T) {
		r := models.GetReminderNewDamageReported("en", BuildTestProperty("1"), BuildTestDamage("1"))
		assert.Equal(t, "New high damage reported in room Living Room of property Test.", r.Title)
	})

	t.Run("DamageFixPlanned", func(t *testing.T) {
		r := models.GetReminderDamageFixPlanned("en", BuildTestProperty("1"), BuildTestDamage("1"), 3)
		assert.Equal(t, "Damage in room Living Room of property Test is planned to be fixed in 3 days.", r.Title)
	})

	t.Run("UrgentDamageNotPlanned", func(t *testing.T) {
		r := models.GetReminderUrgentDamageNotPlanned("en", BuildTestProperty("1"), BuildTestDamage("1"))
		assert.Equal(t, "Urgent damage in room Living Room of property Test is not planned to be fixed.", r.Title)
	})

	t.Run("DamageOlderThan7Days", func(t *testing.T) {
		r := models.GetReminderDamageOlderThan7Days("en", BuildTestProperty("1"), BuildTestDamage("1"))
		assert.Equal(t, "Damage in room Living Room of property Test was created more than 7 days ago.", r.Title)
	})

	t.Run("DamageFixedByTenant", func(t *testing.T) {
		r := models.GetReminderDamageFixedByTenant("en", BuildTestProperty("1"), BuildTestDamage("1"))
		assert.Equal(t, "Damage in room Living Room of property Test was marked as 'fixed by tenant'.", r.Title)
	})

	t.Run("FixDateOverdue", func(t *testing.T) {
		r := models.GetReminderFixDateOverdue("en", BuildTestProperty("1"), BuildTestDamage("1"), 2)
		assert.Equal(t, "Damage in room Living Room of property Test was planned to be fixed 2 days ago.", r.Title)
	})

	t.Run("UnreadMessages", func(t *testing.T) {
		r := models.GetReminderUnreadMessages("en", 7)
		assert.Equal(t, "You have 7 unread messages.", r.Title)
	})

	t.Run("AllGood", func(t *testing.T) {
		r := models.GetReminderAllGood("en")
		assert.Equal(t, "Good news!", r.Title)
	})
}

func TestOpenDamageResponse(t *testing.T) {
	t.Run("FromDbDamage", func(t *testing.T) {
		mockDamageModel := BuildTestDamage("1")

		resp := models.OpenDamageResponse{}
		resp.FromDbDamage(mockDamageModel, db.UserModel{
			InnerUser: db.InnerUser{
				Firstname: "John",
				Lastname:  "Doe",
			},
		}, BuildTestProperty("1"))

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
	})
}
