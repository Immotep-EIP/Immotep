package models

import (
	"strconv"
	"strings"

	"keyz/backend/prisma/db"
)

// Example of a Reminder:
// 1. Lease of property X is ending in X days. Plan an inventory appointment with the tenant to fill the inventory report.
// 2. Lease of property X started recently and does not have an inventory report. Plan an appointment with the tenant to fill the inventory report.
// 3. Property X is available for rent. Consider inviting a new tenant.
// 4. Property X does not have an inventory with rooms and furnitures. Please fill to fill the inventory report quickly the first time you visit the property.
// 5. Lease invitation for property X has been pending for more than X days. Follow up with the prospective tenant.
// 6. New P damage reported in room Y of property X. Please review and plan a fix date. {if !d.read}
// 7. Damage in room Y of property X is planned to be fixed in X days. Please remember to check the progress.
// 8. Urgent damage in room Y of property X is not planned to be fixed. Please plan a fix date. {if d.read && d.planned_fix_date == nil}
// 9. Damage in room Y of property X was created more than 7 days ago. Please plan a fix date. {if d.read && d.planned_fix_date == nil && d.created_at < 7 days}
// 10. Damage in room Y of property X was marked as 'fixed by tenant'. Review and confirm. {if d.fixed_tenant}
// 11. Damage in room Y of property X was planned to be fixed X days ago. Please mark it as 'fixed' or modify the planned fix date.
// 12. You have X unread messages. Please check your inbox.
//
// If no reminders:
// 13. Good news! All your properties are in good condition and have no pending issues.
//
// TODO later:
// - Room Y in property X was rated 'broken' or 'needsRepair' in the latest report. Schedule a check and mark it as 'fixed' in the inventory report.
// - Furniture 'Z' in property X was rated 'broken' 'broken' or 'needsRepair' in the latest report. Schedule a check and mark it as 'fixed' in the inventory report.
type Reminder struct {
	Id       string      `json:"id"`
	Priority db.Priority `json:"priority"`
	Title    string      `json:"title"`
	Advice   string      `json:"advice"`
	Link     string      `json:"link,omitempty"`
}

type reminderModel map[string]Reminder

func (model reminderModel) Get(lang string) Reminder {
	var res Reminder
	if r, ok := model[lang]; ok {
		res = r
	} else {
		res = model["en"]
	}
	return res
}

func replacePlaceholders(text string, placeholders map[string]string) string {
	for key, value := range placeholders {
		text = strings.ReplaceAll(text, "{"+key+"}", value)
	}
	return text
}

func (r Reminder) WithPlaceholders(values map[string]string) Reminder {
	r.Title = replacePlaceholders(r.Title, values)
	r.Advice = replacePlaceholders(r.Advice, values)
	r.Link = replacePlaceholders(r.Link, values)
	return r
}

// 1
var ReminderLeaseEnding = reminderModel{
	"en": {
		Id:       "1",
		Priority: db.PriorityHigh,
		Title:    "Lease of property {property} is ending in {days} days.",
		Advice:   "Plan an inventory appointment with the tenant to fill the inventory report.",
		Link:     "/real-property/details/{property_id}",
	},
	"fr": {
		Id:       "1",
		Priority: db.PriorityHigh,
		Title:    "Le bail de la propriété {property} se termine dans {days} jours.",
		Advice:   "Planifiez un rendez-vous avec le locataire pour remplir l'état des lieux.",
		Link:     "/real-property/details/{property_id}",
	},
}

func GetReminderLeaseEnding(lang string, property db.PropertyModel, days int) Reminder {
	return ReminderLeaseEnding.Get(lang).WithPlaceholders(map[string]string{
		"property":    property.Name,
		"days":        strconv.Itoa(days),
		"property_id": property.ID,
	})
}

// 2
var ReminderNoInventoryReport = reminderModel{
	"en": {
		Id:       "2",
		Priority: db.PriorityHigh,
		Title:    "Lease of property {property} started recently and does not have an inventory report.",
		Advice:   "Plan an appointment with the tenant to fill the inventory report.",
		Link:     "/real-property/details/{property_id}",
	},
	"fr": {
		Id:       "2",
		Priority: db.PriorityHigh,
		Title:    "Le bail de la propriété {property} a récemment commencé et n'a pas d'état des lieux.",
		Advice:   "Planifiez un rendez-vous avec le locataire pour remplir l'état des lieux.",
		Link:     "/real-property/details/{property_id}",
	},
}

func GetReminderNoInventoryReport(lang string, property db.PropertyModel) Reminder {
	return ReminderNoInventoryReport.Get(lang).WithPlaceholders(map[string]string{
		"property":    property.Name,
		"property_id": property.ID,
	})
}

// 3
var ReminderPropertyAvailable = reminderModel{
	"en": {
		Id:       "3",
		Priority: db.PriorityLow,
		Title:    "Property {property} is available for rent.",
		Advice:   "Consider inviting a new tenant.",
		Link:     "/real-property/details/{property_id}",
	},
	"fr": {
		Id:       "3",
		Priority: db.PriorityLow,
		Title:    "La propriété {property} est disponible à la location.",
		Advice:   "Envisagez d'inviter un nouveau locataire.",
		Link:     "/real-property/details/{property_id}",
	},
}

func GetReminderPropertyAvailable(lang string, property db.PropertyModel) Reminder {
	return ReminderPropertyAvailable.Get(lang).WithPlaceholders(map[string]string{
		"property":    property.Name,
		"property_id": property.ID,
	})
}

// 4
var ReminderEmptyInventory = reminderModel{
	"en": {
		Id:       "4",
		Priority: db.PriorityMedium,
		Title:    "Inventory of property {property} is empty.",
		Advice:   "Please add rooms and furnitures to be able to fill the inventory report quickly during your first visit of the property.",
		Link:     "/real-property/details/{property_id}",
	},
	"fr": {
		Id:       "4",
		Priority: db.PriorityMedium,
		Title:    "L'inventaire de la propriété {property} est vide.",
		Advice:   "Veuillez ajouter des chambers et des meubles pour pouvoir remplir rapidement l'état des lieux lors de votre première visite de la propriété.",
		Link:     "/real-property/details/{property_id}",
	},
}

func GetReminderEmptyInventory(lang string, property db.PropertyModel) Reminder {
	return ReminderEmptyInventory.Get(lang).WithPlaceholders(map[string]string{
		"property":    property.Name,
		"property_id": property.ID,
	})
}

// 5
var ReminderPendingLeaseInvitation = reminderModel{
	"en": {
		Id:       "5",
		Priority: db.PriorityMedium,
		Title:    "Lease invitation for property {property} has been pending for {days} days.",
		Advice:   "Follow up with the prospective tenant.",
		Link:     "/real-property/details/{property_id}",
	},
	"fr": {
		Id:       "5",
		Priority: db.PriorityMedium,
		Title:    "L'invitation à la location de la propriété {property} est en attente depuis {days} jours.",
		Advice:   "Relancez le locataire potentiel.",
		Link:     "/real-property/details/{property_id}",
	},
}

func GetReminderPendingLeaseInvitation(lang string, property db.PropertyModel, days int) Reminder {
	return ReminderPendingLeaseInvitation.Get(lang).WithPlaceholders(map[string]string{
		"property":    property.Name,
		"days":        strconv.Itoa(days),
		"property_id": property.ID,
	})
}

// 6
var ReminderNewDamageReported = reminderModel{
	"en": {
		Id:       "6",
		Priority: db.PriorityHigh,
		Title:    "New {priority} damage reported in room {room} of property {property}.",
		Advice:   "Please review and plan a fix date.",
		Link:     "/real-property/details/{property_id}/damage/{damage_id}",
	},
	"fr": {
		Id:       "6",
		Priority: db.PriorityHigh,
		Title:    "Nouveau dommage {priority} signalé dans la pièce {room} de la propriété {property}.",
		Advice:   "Veuillez examiner et planifier une date de réparation.",
		Link:     "/real-property/details/{property_id}/damage/{damage_id}",
	},
}

func GetReminderNewDamageReported(lang string, property db.PropertyModel, damage db.DamageModel) Reminder {
	return ReminderNewDamageReported.Get(lang).WithPlaceholders(map[string]string{
		"property":    property.Name,
		"room":        damage.Room().Name,
		"priority":    string(damage.Priority),
		"property_id": property.ID,
		"damage_id":   damage.ID,
	})
}

// 7
var ReminderDamageFixPlanned = reminderModel{
	"en": {
		Id:       "7",
		Priority: db.PriorityHigh,
		Title:    "Damage in room {room} of property {property} is planned to be fixed in {days} days.",
		Advice:   "Please remember to check the progress.",
		Link:     "/real-property/details/{property_id}/damage/{damage_id}",
	},
	"fr": {
		Id:       "7",
		Priority: db.PriorityHigh,
		Title:    "Le dommage dans la pièce {room} de la propriété {property} doit être réparé dans {days} jours.",
		Advice:   "Veuillez vous souvenir de vérifier l'avancement.",
		Link:     "/real-property/details/{property_id}/damage/{damage_id}",
	},
}

func GetReminderDamageFixPlanned(lang string, property db.PropertyModel, damage db.DamageModel, days int) Reminder {
	return ReminderDamageFixPlanned.Get(lang).WithPlaceholders(map[string]string{
		"property":    property.Name,
		"room":        damage.Room().Name,
		"days":        strconv.Itoa(days),
		"property_id": property.ID,
		"damage_id":   damage.ID,
	})
}

// 8
var ReminderUrgentDamageNotPlanned = reminderModel{
	"en": {
		Id:       "8",
		Priority: db.PriorityUrgent,
		Title:    "Urgent damage in room {room} of property {property} is not planned to be fixed.",
		Advice:   "Please plan a fix date.",
		Link:     "/real-property/details/{property_id}/damage/{damage_id}",
	},
	"fr": {
		Id:       "8",
		Priority: db.PriorityUrgent,
		Title:    "Le dommage urgent dans la pièce {room} de la propriété {property} n'est pas prévu pour être réparé.",
		Advice:   "Veuillez planifier une date de réparation.",
		Link:     "/real-property/details/{property_id}/damage/{damage_id}",
	},
}

func GetReminderUrgentDamageNotPlanned(lang string, property db.PropertyModel, damage db.DamageModel) Reminder {
	return ReminderUrgentDamageNotPlanned.Get(lang).WithPlaceholders(map[string]string{
		"property":    property.Name,
		"room":        damage.Room().Name,
		"property_id": property.ID,
		"damage_id":   damage.ID,
	})
}

// 9
var ReminderDamageOlderThan7Days = reminderModel{
	"en": {
		Id:       "9",
		Priority: db.PriorityUrgent,
		Title:    "Damage in room {room} of property {property} was created more than 7 days ago.",
		Advice:   "Please plan a fix date.",
		Link:     "/real-property/details/{property_id}/damage/{damage_id}",
	},
	"fr": {
		Id:       "9",
		Priority: db.PriorityUrgent,
		Title:    "Le dommage dans la pièce {room} de la propriété {property} a été créé il y a plus de 7 jours.",
		Advice:   "Veuillez planifier une date de réparation.",
		Link:     "/real-property/details/{property_id}/damage/{damage_id}",
	},
}

func GetReminderDamageOlderThan7Days(lang string, property db.PropertyModel, damage db.DamageModel) Reminder {
	return ReminderDamageOlderThan7Days.Get(lang).WithPlaceholders(map[string]string{
		"property":    property.Name,
		"room":        damage.Room().Name,
		"property_id": property.ID,
		"damage_id":   damage.ID,
	})
}

// 10
var ReminderDamageFixedByTenant = reminderModel{
	"en": {
		Id:       "10",
		Priority: db.PriorityMedium,
		Title:    "Damage in room {room} of property {property} was marked as 'fixed by tenant'.",
		Advice:   "Please review and confirm.",
		Link:     "/real-property/details/{property_id}/damage/{damage_id}",
	},
	"fr": {
		Id:       "10",
		Priority: db.PriorityMedium,
		Title:    "Le dommage dans la pièce {room} de la propriété {property} a été marqué comme 'réparé par le locataire'.",
		Advice:   "Veuillez examiner et confirmer.",
		Link:     "/real-property/details/{property_id}/damage/{damage_id}",
	},
}

func GetReminderDamageFixedByTenant(lang string, property db.PropertyModel, damage db.DamageModel) Reminder {
	return ReminderDamageFixedByTenant.Get(lang).WithPlaceholders(map[string]string{
		"property":    property.Name,
		"room":        damage.Room().Name,
		"property_id": property.ID,
		"damage_id":   damage.ID,
	})
}

// 11
var ReminderFixDateOverdue = reminderModel{
	"en": {
		Id:       "11",
		Priority: db.PriorityHigh,
		Title:    "Damage in room {room} of property {property} was planned to be fixed {days} days ago.",
		Advice:   "Please mark it as 'fixed' or modify the planned fix date.",
		Link:     "/real-property/details/{property_id}/damage/{damage_id}",
	},
	"fr": {
		Id:       "11",
		Priority: db.PriorityHigh,
		Title:    "Le dommage dans la pièce {room} de la propriété {property} devait être réparé il y a {days} jours.",
		Advice:   "Veuillez le marquer comme 'réparé' ou modifier la date de réparation prévue.",
		Link:     "/real-property/details/{property_id}/damage/{damage_id}",
	},
}

func GetReminderFixDateOverdue(lang string, property db.PropertyModel, damage db.DamageModel, days int) Reminder {
	return ReminderFixDateOverdue.Get(lang).WithPlaceholders(map[string]string{
		"property":    property.Name,
		"room":        damage.Room().Name,
		"days":        strconv.Itoa(days),
		"property_id": property.ID,
		"damage_id":   damage.ID,
	})
}

// 12
var ReminderUnreadMessages = reminderModel{
	"en": {
		Id:       "12",
		Priority: db.PriorityLow,
		Title:    "You have {messages} unread messages.",
		Advice:   "Please check your inbox.",
		Link:     "/messages",
	},
	"fr": {
		Id:       "12",
		Priority: db.PriorityLow,
		Title:    "Vous avez {messages} messages non lus.",
		Advice:   "Veuillez vérifier votre boîte de réception.",
		Link:     "/messages",
	},
}

func GetReminderUnreadMessages(lang string, messages int) Reminder {
	return ReminderUnreadMessages.Get(lang).WithPlaceholders(map[string]string{
		"messages": strconv.Itoa(messages),
	})
}

// 13
var ReminderAllGood = reminderModel{
	"en": {
		Id:       "13",
		Priority: db.PriorityLow,
		Title:    "Good news!",
		Advice:   "All your properties are in good condition and have no pending issues.",
		Link:     "",
	},
	"fr": {
		Id:       "13",
		Priority: db.PriorityLow,
		Title:    "Bonne nouvelle !",
		Advice:   "Toutes vos propriétés sont en bon état et n'ont aucun problème en attente.",
		Link:     "",
	},
}

func GetReminderAllGood(lang string) Reminder {
	return ReminderAllGood.Get(lang)
}

type DashboardProperties struct {
	NbrTotal          int                `json:"nbr_total"`
	NbrArchived       int                `json:"nbr_archived"`
	NbrOccupied       int                `json:"nbr_occupied"`
	NbrAvailable      int                `json:"nbr_available"`
	NbrPendingInvites int                `json:"nbr_pending_invites"`
	ListRecentlyAdded []db.InnerProperty `json:"list_recently_added"`
}

type OpenDamageResponse struct {
	ID           string `json:"id"`
	LeaseID      string `json:"lease_id"`
	TenantName   string `json:"tenant_name"`
	PropertyID   string `json:"property_id"`
	PropertyName string `json:"property_name"`
	RoomID       string `json:"room_id"`
	RoomName     string `json:"room_name"`

	Comment      string       `json:"comment"`
	Priority     db.Priority  `json:"priority"`
	Read         bool         `json:"read"`
	CreatedAt    db.DateTime  `json:"created_at"`
	UpdatedAt    db.DateTime  `json:"updated_at"`
	FixStatus    db.FixStatus `json:"fix_status"`
	FixPlannedAt *db.DateTime `json:"fix_planned_at"`
	FixedAt      *db.DateTime `json:"fixed_at,omitempty"`
}

func (i *OpenDamageResponse) FromDbDamage(model db.DamageModel, tenant db.UserModel, property db.PropertyModel) {
	i.ID = model.ID
	i.LeaseID = model.LeaseID
	i.TenantName = tenant.Name()
	i.PropertyID = property.ID
	i.PropertyName = property.Name
	i.RoomID = model.RoomID
	i.RoomName = model.Room().Name

	i.Comment = model.Comment
	i.Priority = model.Priority
	i.Read = model.Read
	i.CreatedAt = model.CreatedAt
	i.UpdatedAt = model.UpdatedAt

	i.FixStatus = model.FixStatus()
	i.FixPlannedAt = model.InnerDamage.FixPlannedAt
	i.FixedAt = model.InnerDamage.FixedAt
}

type DashboardOpenDamages struct {
	NbrTotal                int                  `json:"nbr_total"`
	NbrUrgent               int                  `json:"nbr_urgent"`
	NbrHigh                 int                  `json:"nbr_high"`
	NbrMedium               int                  `json:"nbr_medium"`
	NbrLow                  int                  `json:"nbr_low"`
	NbrPlannedToFixThisWeek int                  `json:"nbr_planned_to_fix_this_week"`
	ListToFix               []OpenDamageResponse `json:"list_to_fix"`
}

// type messages struct {
// 	NbrUnread  int               `json:"nbr_unread"`
// 	ListUnread []db.MessageModel `json:"list_unread"`
// }

type DashboardResponse struct {
	Reminders   []Reminder           `json:"reminders"`
	Properties  DashboardProperties  `json:"properties"`
	OpenDamages DashboardOpenDamages `json:"open_damages"`
	// Messages    messages    `json:"messages"`
}
