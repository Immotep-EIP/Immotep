package controllers

import (
	"net/http"
	"slices"
	"time"

	"github.com/gin-gonic/gin"
	"immotep/backend/models"
	"immotep/backend/prisma/db"
	"immotep/backend/services/database"
	"immotep/backend/utils"
)

func sortReminder(a, b models.Reminder) int {
	if a.Priority == b.Priority {
		return 0
	}
	if a.Priority == db.PriorityUrgent {
		return -1
	}
	if b.Priority == db.PriorityUrgent {
		return 1
	}
	if a.Priority == db.PriorityHigh {
		return -1
	}
	if b.Priority == db.PriorityHigh {
		return 1
	}
	if a.Priority == db.PriorityMedium {
		return -1
	}
	if b.Priority == db.PriorityMedium {
		return 1
	}
	return 0
}

func getReminders_Damage_FixPlanned(lang string, now time.Time, property db.PropertyModel, damage db.DamageModel) []models.Reminder {
	var res []models.Reminder
	fixPlanAt, fixPlanned := damage.FixPlannedAt()

	// reminder 11
	if fixPlanned && fixPlanAt.Before(now) {
		res = append(res, models.GetReminderFixDateOverdue(lang, property, damage, int(now.Sub(fixPlanAt).Hours())/24))
		// reminder 7
	} else if fixPlanned && fixPlanAt.Before(now.AddDate(0, 0, 7)) {
		res = append(res, models.GetReminderDamageFixPlanned(lang, property, damage, int(fixPlanAt.Sub(now).Hours())/24))
	}

	// reminder 8
	if !fixPlanned && damage.Priority == db.PriorityUrgent {
		res = append(res, models.GetReminderUrgentDamageNotPlanned(lang, property, damage))
	}

	// reminder 9
	if damage.Read && !fixPlanned && damage.CreatedAt.Before(now.AddDate(0, 0, -7)) {
		res = append(res, models.GetReminderDamageOlderThan7Days(lang, property, damage))
	}

	return res
}

func getReminders_Damage(lang string, now time.Time, property db.PropertyModel, damages []db.DamageModel) []models.Reminder {
	var res []models.Reminder

	for _, damage := range damages {
		if damage.FixedOwner {
			continue
		}

		// reminder 6
		if !damage.Read {
			res = append(res, models.GetReminderNewDamageReported(lang, property, damage))
		}

		res = append(res, getReminders_Damage_FixPlanned(lang, now, property, damage)...)

		// reminder 10
		if damage.FixedTenant {
			res = append(res, models.GetReminderDamageFixedByTenant(lang, property, damage))
		}
	}
	return res
}

func getReminders_Lease(lang string, now time.Time, property db.PropertyModel, currentLease db.LeaseModel) []models.Reminder {
	var res []models.Reminder

	// reminder 1
	end, endOk := currentLease.EndDate()
	if endOk && end.Before(now.AddDate(0, 0, 30)) {
		res = append(res, models.GetReminderLeaseEnding(lang, property, int(end.Sub(now).Hours())/24))
	}

	// reminder 2
	if len(currentLease.Reports()) == 0 {
		res = append(res, models.GetReminderNoInventoryReport(lang, property))
	}

	res = append(res, getReminders_Damage(lang, now, property, currentLease.Damages())...)
	return res
}

func getReminders_Property(lang string, now time.Time, property db.PropertyModel) []models.Reminder {
	var res []models.Reminder

	iCurrentLease := slices.IndexFunc(property.Leases(), func(lease db.LeaseModel) bool { return lease.Active })
	invite, inviteOk := property.LeaseInvite()

	if iCurrentLease != -1 {
		currentLease := property.Leases()[iCurrentLease]
		res = append(res, getReminders_Lease(lang, now, property, currentLease)...)
	} else if !inviteOk {
		// reminder 3
		res = append(res, models.GetReminderPropertyAvailable(lang, property))
	}

	// reminder 4
	if len(property.Rooms()) == 0 {
		res = append(res, models.GetReminderEmptyInventory(lang, property))
	}

	// reminder 5
	if inviteOk && invite.CreatedAt.Before(now.AddDate(0, 0, -7)) {
		res = append(res, models.GetReminderPendingLeaseInvitation(lang, property, int(now.Sub(invite.CreatedAt).Hours())/24))
	}
	return res
}

func getReminders(lang string, properties []db.PropertyModel) []models.Reminder {
	var res []models.Reminder
	now := time.Now()

	for _, property := range properties {
		res = append(res, getReminders_Property(lang, now, property)...)
	}

	if len(res) == 0 {
		res = append(res, models.GetReminderAllGood(lang))
	}

	slices.SortFunc(res, sortReminder)
	return res
}

func getPropertyAndDamageDashboard_Properties(pRes *models.DashboardProperties, property db.PropertyModel, now time.Time) {
	pRes.NbrTotal++

	if property.Archived {
		pRes.NbrArchived++
	}
	if slices.IndexFunc(property.Leases(), func(lease db.LeaseModel) bool { return lease.Active }) != -1 {
		pRes.NbrOccupied++
	} else {
		pRes.NbrAvailable++
	}
	if property.RelationsProperty.LeaseInvite != nil {
		pRes.NbrPendingInvites++
	}
	if property.CreatedAt.After(now.AddDate(0, 0, -7)) {
		pRes.ListRecentlyAdded = append(pRes.ListRecentlyAdded, property.InnerProperty)
	}
}

func getPropertyAndDamageDashboard_Damage(dRes *models.DashboardOpenDamages, damage db.DamageModel, lease db.LeaseModel, property db.PropertyModel, now time.Time) {
	if !damage.IsFixed() {
		dRes.NbrTotal++
		switch damage.Priority {
		case db.PriorityUrgent:
			dRes.NbrUrgent++
		case db.PriorityHigh:
			dRes.NbrHigh++
		case db.PriorityMedium:
			dRes.NbrMedium++
		case db.PriorityLow:
			dRes.NbrLow++
		}

		fixPlanAt, fixPlanned := damage.FixPlannedAt()
		if fixPlanned && fixPlanAt.Before(now.AddDate(0, 0, 7)) && fixPlanAt.After(now) {
			dRes.NbrPlannedToFixThisWeek++
		}
	}
	if !damage.FixedOwner {
		d := models.OpenDamageResponse{}
		d.FromDbDamage(damage, *lease.Tenant(), property)
		dRes.ListToFix = append(dRes.ListToFix, d)
	}
}

func getPropertyAndDamageDashboard_Damages(dRes *models.DashboardOpenDamages, property db.PropertyModel, now time.Time) {
	for _, lease := range property.Leases() {
		for _, damage := range lease.Damages() {
			getPropertyAndDamageDashboard_Damage(dRes, damage, lease, property, now)
		}
	}
}

func getPropertyAndDamageDashboard(properties []db.PropertyModel) (models.DashboardProperties, models.DashboardOpenDamages) {
	now := time.Now()
	pRes := models.DashboardProperties{}
	dRes := models.DashboardOpenDamages{}

	for _, property := range properties {
		getPropertyAndDamageDashboard_Properties(&pRes, property, now)
		getPropertyAndDamageDashboard_Damages(&dRes, property, now)
	}

	return pRes, dRes
}

// GetOwnerDashboard godoc
//
//	@Summary		Get dashboard
//	@Description	Get a summary of every data related to an owner
//	@Tags			dashboard
//	@Accept			json
//	@Produce		json
//	@Param			lang	query		string						true	"Language of the response (default: en)"
//	@Success		200		{array}		models.DashboardResponse	"Dashboard data"
//	@Failure		403		{object}	utils.Error					"Not an owner"
//	@Failure		500
//	@Security		Bearer
//	@Router			/owner/dashboard/ [get]
func GetOwnerDashboard(c *gin.Context) {
	lang := c.DefaultQuery("lang", "en")
	claims := utils.GetClaims(c)

	properties := database.GetAllDatasFromProperties(claims["id"])
	pRes, dRes := getPropertyAndDamageDashboard(properties)

	c.JSON(http.StatusOK, models.DashboardResponse{
		Reminders:   getReminders(lang, properties),
		Properties:  pRes,
		OpenDamages: dRes,
	})
}
