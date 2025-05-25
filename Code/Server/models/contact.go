package models

import "immotep/backend/prisma/db"

type ContactMessageRequest struct {
	Firstname string `binding:"required"       json:"firstname"`
	Lastname  string `binding:"required"       json:"lastname"`
	Email     string `binding:"required,email" json:"email"`
	Subject   string `binding:"required"       json:"subject"`
	Message   string `binding:"required"       json:"message"`
}

func (r *ContactMessageRequest) ToDbContact() db.ContactMessageModel {
	return db.ContactMessageModel{
		InnerContactMessage: db.InnerContactMessage{
			Firstname: r.Firstname,
			Lastname:  r.Lastname,
			Email:     r.Email,
			Subject:   r.Subject,
			Message:   r.Message,
		},
	}
}
