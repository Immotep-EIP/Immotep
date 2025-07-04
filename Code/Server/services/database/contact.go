package database

import (
	"keyz/backend/prisma/db"
	"keyz/backend/services"
	"keyz/backend/utils"
)

func CreateContactMessage(contact db.ContactMessageModel) db.ContactMessageModel {
	pdb := services.DBclient
	newContact, err := pdb.Client.ContactMessage.CreateOne(
		db.ContactMessage.Firstname.Set(contact.Firstname),
		db.ContactMessage.Lastname.Set(contact.Lastname),
		db.ContactMessage.Email.Set(utils.SanitizeEmail(contact.Email)),
		db.ContactMessage.Subject.Set(contact.Subject),
		db.ContactMessage.Message.Set(contact.Message),
	).Exec(pdb.Context)
	if err != nil {
		panic(err)
	}
	return *newContact
}

func MockCreateContactMessage(c *services.PrismaDB, contact db.ContactMessageModel) db.ContactMessageMockExpectParam {
	return c.Client.ContactMessage.CreateOne(
		db.ContactMessage.Firstname.Set(contact.Firstname),
		db.ContactMessage.Lastname.Set(contact.Lastname),
		db.ContactMessage.Email.Set(utils.SanitizeEmail(contact.Email)),
		db.ContactMessage.Subject.Set(contact.Subject),
		db.ContactMessage.Message.Set(contact.Message),
	)
}
