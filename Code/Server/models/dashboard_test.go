package models_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"immotep/backend/models"
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
		r := models.GetReminderLeaseEnding("en", "Test Property", 30)
		assert.Equal(t, "Lease of property Test Property is ending in 30 days.", r.Title)
	})

	t.Run("NoInventoryReport", func(t *testing.T) {
		r := models.GetReminderNoInventoryReport("en", "Test Property")
		assert.Equal(t, "Lease of property Test Property started recently and does not have an inventory report.", r.Title)
	})

	t.Run("PropertyAvailable", func(t *testing.T) {
		r := models.GetReminderPropertyAvailable("en", "Test Property")
		assert.Equal(t, "Property Test Property is available for rent.", r.Title)
	})

	t.Run("EmptyInventory", func(t *testing.T) {
		r := models.GetReminderEmptyInventory("en", "Test Property")
		assert.Equal(t, "Inventory of property Test Property is empty.", r.Title)
	})

	t.Run("PendingLeaseInvitation", func(t *testing.T) {
		r := models.GetReminderPendingLeaseInvitation("en", "Test Property", 5)
		assert.Equal(t, "Lease invitation for property Test Property has been pending for 5 days.", r.Title)
	})

	t.Run("NewDamageReported", func(t *testing.T) {
		r := models.GetReminderNewDamageReported("en", "Test Property", "Living Room", "high")
		assert.Equal(t, "New high damage reported in room Living Room of property Test Property.", r.Title)
	})

	t.Run("DamageFixPlanned", func(t *testing.T) {
		r := models.GetReminderDamageFixPlanned("en", "Test Property", "Living Room", 3)
		assert.Equal(t, "Damage in room Living Room of property Test Property is planned to be fixed in 3 days.", r.Title)
	})

	t.Run("UrgentDamageNotPlanned", func(t *testing.T) {
		r := models.GetReminderUrgentDamageNotPlanned("en", "Test Property", "Living Room")
		assert.Equal(t, "Urgent damage in room Living Room of property Test Property is not planned to be fixed.", r.Title)
	})

	t.Run("DamageOlderThan7Days", func(t *testing.T) {
		r := models.GetReminderDamageOlderThan7Days("en", "Test Property", "Living Room")
		assert.Equal(t, "Damage in room Living Room of property Test Property was created more than 7 days ago.", r.Title)
	})

	t.Run("DamageFixedByTenant", func(t *testing.T) {
		r := models.GetReminderDamageFixedByTenant("en", "Test Property", "Living Room")
		assert.Equal(t, "Damage in room Living Room of property Test Property was marked as 'fixed by tenant'.", r.Title)
	})

	t.Run("FixDateOverdue", func(t *testing.T) {
		r := models.GetReminderFixDateOverdue("en", "Test Property", "Living Room", 2)
		assert.Equal(t, "Damage in room Living Room of property Test Property was planned to be fixed 2 days ago.", r.Title)
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
