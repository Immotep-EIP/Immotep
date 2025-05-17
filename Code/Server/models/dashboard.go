package models

import (
	"os"
	"strconv"
	"strings"

	"immotep/backend/prisma/db"
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
	Link     string      `json:"link"`
}

type reminderModel map[string]Reminder

func (model reminderModel) Get(lang string) Reminder {
	var res Reminder
	if r, ok := model[lang]; ok {
		res = r
	} else {
		res = model["en"]
	}
	res.Link = os.Getenv("WEB_PUBLIC_URL") + res.Link
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
	return r
}

// 1
var ReminderLeaseEnding = reminderModel{
	"en": {
		Id:       "1",
		Priority: db.PriorityHigh,
		Title:    "Lease of property {property} is ending in {days} days.",
		Advice:   "Plan an inventory appointment with the tenant to fill the inventory report.",
		Link:     "/TODO",
	},
	"fr": {
		Id:       "1",
		Priority: db.PriorityHigh,
		Title:    "Le bail de la propriété {property} se termine dans {days} jours.",
		Advice:   "Planifiez un rendez-vous avec le locataire pour remplir l'état des lieux.",
		Link:     "/TODO",
	},
}

func GetReminderLeaseEnding(lang string, property string, days int) Reminder {
	return ReminderLeaseEnding.Get(lang).WithPlaceholders(map[string]string{
		"property": property,
		"days":     strconv.Itoa(days),
	})
}

// 2
var ReminderNoInventoryReport = reminderModel{
	"en": {
		Id:       "2",
		Priority: db.PriorityHigh,
		Title:    "Lease of property {property} started recently and does not have an inventory report.",
		Advice:   "Plan an appointment with the tenant to fill the inventory report.",
		Link:     "/TODO",
	},
	"fr": {
		Id:       "2",
		Priority: db.PriorityHigh,
		Title:    "Le bail de la propriété {property} a récemment commencé et n'a pas d'état des lieux.",
		Advice:   "Planifiez un rendez-vous avec le locataire pour remplir l'état des lieux.",
		Link:     "/TODO",
	},
}

func GetReminderNoInventoryReport(lang string, property string) Reminder {
	return ReminderNoInventoryReport.Get(lang).WithPlaceholders(map[string]string{
		"property": property,
	})
}

// 3
var ReminderPropertyAvailable = reminderModel{
	"en": {
		Id:       "3",
		Priority: db.PriorityLow,
		Title:    "Property {property} is available for rent.",
		Advice:   "Consider inviting a new tenant.",
		Link:     "/TODO",
	},
	"fr": {
		Id:       "3",
		Priority: db.PriorityLow,
		Title:    "La propriété {property} est disponible à la location.",
		Advice:   "Envisagez d'inviter un nouveau locataire.",
		Link:     "/TODO",
	},
}

func GetReminderPropertyAvailable(lang string, property string) Reminder {
	return ReminderPropertyAvailable.Get(lang).WithPlaceholders(map[string]string{
		"property": property,
	})
}

// 4
var ReminderEmptyInventory = reminderModel{
	"en": {
		Id:       "4",
		Priority: db.PriorityMedium,
		Title:    "Inventory of property {property} is empty.",
		Advice:   "Please add rooms and furnitures to be able to fill the inventory report quickly during your first visit of the property.",
		Link:     "/TODO",
	},
	"fr": {
		Id:       "4",
		Priority: db.PriorityMedium,
		Title:    "L'inventaire de la propriété {property} est vide.",
		Advice:   "Veuillez ajouter des chambers et des meubles pour pouvoir remplir rapidement l'état des lieux lors de votre première visite de la propriété.",
		Link:     "/TODO",
	},
}

func GetReminderEmptyInventory(lang string, property string) Reminder {
	return ReminderEmptyInventory.Get(lang).WithPlaceholders(map[string]string{
		"property": property,
	})
}

// 5
var ReminderPendingLeaseInvitation = reminderModel{
	"en": {
		Id:       "5",
		Priority: db.PriorityMedium,
		Title:    "Lease invitation for property {property} has been pending for {days} days.",
		Advice:   "Follow up with the prospective tenant.",
		Link:     "/TODO",
	},
	"fr": {
		Id:       "5",
		Priority: db.PriorityMedium,
		Title:    "L'invitation à la location de la propriété {property} est en attente depuis {days} jours.",
		Advice:   "Relancez le locataire potentiel.",
		Link:     "/TODO",
	},
}

func GetReminderPendingLeaseInvitation(lang string, property string, days int) Reminder {
	return ReminderPendingLeaseInvitation.Get(lang).WithPlaceholders(map[string]string{
		"property": property,
		"days":     strconv.Itoa(days),
	})
}

// 6
var ReminderNewDamageReported = reminderModel{
	"en": {
		Id:       "6",
		Priority: db.PriorityHigh,
		Title:    "New {priority} damage reported in room {room} of property {property}.",
		Advice:   "Please review and plan a fix date.",
		Link:     "/TODO",
	},
	"fr": {
		Id:       "6",
		Priority: db.PriorityHigh,
		Title:    "Nouveau dommage {priority} signalé dans la pièce {room} de la propriété {property}.",
		Advice:   "Veuillez examiner et planifier une date de réparation.",
		Link:     "/TODO",
	},
}

func GetReminderNewDamageReported(lang string, property string, room string, priority db.Priority) Reminder {
	return ReminderNewDamageReported.Get(lang).WithPlaceholders(map[string]string{
		"property": property,
		"room":     room,
		"priority": string(priority),
	})
}

// 7
var ReminderDamageFixPlanned = reminderModel{
	"en": {
		Id:       "7",
		Priority: db.PriorityHigh,
		Title:    "Damage in room {room} of property {property} is planned to be fixed in {days} days.",
		Advice:   "Please remember to check the progress.",
		Link:     "/TODO",
	},
	"fr": {
		Id:       "7",
		Priority: db.PriorityHigh,
		Title:    "Le dommage dans la pièce {room} de la propriété {property} doit être réparé dans {days} jours.",
		Advice:   "Veuillez vous souvenir de vérifier l'avancement.",
		Link:     "/TODO",
	},
}

func GetReminderDamageFixPlanned(lang string, property string, room string, days int) Reminder {
	return ReminderDamageFixPlanned.Get(lang).WithPlaceholders(map[string]string{
		"property": property,
		"room":     room,
		"days":     strconv.Itoa(days),
	})
}

// 8
var ReminderUrgentDamageNotPlanned = reminderModel{
	"en": {
		Id:       "8",
		Priority: db.PriorityUrgent,
		Title:    "Urgent damage in room {room} of property {property} is not planned to be fixed.",
		Advice:   "Please plan a fix date.",
		Link:     "/TODO",
	},
	"fr": {
		Id:       "8",
		Priority: db.PriorityUrgent,
		Title:    "Le dommage urgent dans la pièce {room} de la propriété {property} n'est pas prévu pour être réparé.",
		Advice:   "Veuillez planifier une date de réparation.",
		Link:     "/TODO",
	},
}

func GetReminderUrgentDamageNotPlanned(lang string, property string, room string) Reminder {
	return ReminderUrgentDamageNotPlanned.Get(lang).WithPlaceholders(map[string]string{
		"property": property,
		"room":     room,
	})
}

// 9
var ReminderDamageOlderThan7Days = reminderModel{
	"en": {
		Id:       "9",
		Priority: db.PriorityUrgent,
		Title:    "Damage in room {room} of property {property} was created more than 7 days ago.",
		Advice:   "Please plan a fix date.",
		Link:     "/TODO",
	},
	"fr": {
		Id:       "9",
		Priority: db.PriorityUrgent,
		Title:    "Le dommage dans la pièce {room} de la propriété {property} a été créé il y a plus de 7 jours.",
		Advice:   "Veuillez planifier une date de réparation.",
		Link:     "/TODO",
	},
}

func GetReminderDamageOlderThan7Days(lang string, property string, room string) Reminder {
	return ReminderDamageOlderThan7Days.Get(lang).WithPlaceholders(map[string]string{
		"property": property,
		"room":     room,
	})
}

// 10
var ReminderDamageFixedByTenant = reminderModel{
	"en": {
		Id:       "10",
		Priority: db.PriorityMedium,
		Title:    "Damage in room {room} of property {property} was marked as 'fixed by tenant'.",
		Advice:   "Please review and confirm.",
		Link:     "/TODO",
	},
	"fr": {
		Id:       "10",
		Priority: db.PriorityMedium,
		Title:    "Le dommage dans la pièce {room} de la propriété {property} a été marqué comme 'réparé par le locataire'.",
		Advice:   "Veuillez examiner et confirmer.",
		Link:     "/TODO",
	},
}

func GetReminderDamageFixedByTenant(lang string, property string, room string) Reminder {
	return ReminderDamageFixedByTenant.Get(lang).WithPlaceholders(map[string]string{
		"property": property,
		"room":     room,
	})
}

// 11
var ReminderFixDateOverdue = reminderModel{
	"en": {
		Id:       "11",
		Priority: db.PriorityHigh,
		Title:    "Damage in room {room} of property {property} was planned to be fixed {days} days ago.",
		Advice:   "Please mark it as 'fixed' or modify the planned fix date.",
		Link:     "/TODO",
	},
	"fr": {
		Id:       "11",
		Priority: db.PriorityHigh,
		Title:    "Le dommage dans la pièce {room} de la propriété {property} devait être réparé il y a {days} jours.",
		Advice:   "Veuillez le marquer comme 'réparé' ou modifier la date de réparation prévue.",
		Link:     "/TODO",
	},
}

func GetReminderFixDateOverdue(lang string, property string, room string, days int) Reminder {
	return ReminderFixDateOverdue.Get(lang).WithPlaceholders(map[string]string{
		"property": property,
		"room":     room,
		"days":     strconv.Itoa(days),
	})
}

// 12
var ReminderUnreadMessages = reminderModel{
	"en": {
		Id:       "12",
		Priority: db.PriorityLow,
		Title:    "You have {messages} unread messages.",
		Advice:   "Please check your inbox.",
		Link:     "/TODO",
	},
	"fr": {
		Id:       "12",
		Priority: db.PriorityLow,
		Title:    "Vous avez {messages} messages non lus.",
		Advice:   "Veuillez vérifier votre boîte de réception.",
		Link:     "/TODO",
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
		Link:     "/TODO",
	},
	"fr": {
		Id:       "13",
		Priority: db.PriorityLow,
		Title:    "Bonne nouvelle !",
		Advice:   "Toutes vos propriétés sont en bon état et n'ont aucun problème en attente.",
		Link:     "/TODO",
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

type DashboardOpenDamages struct {
	NbrTotal                int              `json:"nbr_total"`
	NbrUrgent               int              `json:"nbr_urgent"`
	NbrHigh                 int              `json:"nbr_high"`
	NbrMedium               int              `json:"nbr_medium"`
	NbrLow                  int              `json:"nbr_low"`
	NbrPlannedToFixThisWeek int              `json:"nbr_planned_to_fix_this_week"`
	ListToFix               []db.InnerDamage `json:"list_to_fix"`
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
