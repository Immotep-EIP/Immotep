package models

import (
	"immotep/backend/prisma/db"
)

type UserRequest struct {
	Email     string `binding:"required,email" json:"email"`
	Firstname string `binding:"required"       json:"firstname"`
	Lastname  string `binding:"required"       json:"lastname"`
	Password  string `binding:"required,min=8" json:"password"`
}

func (u *UserRequest) ToDbUser() db.UserModel {
	return db.UserModel{
		InnerUser: db.InnerUser{
			Email:     u.Email,
			Firstname: u.Firstname,
			Lastname:  u.Lastname,
			Password:  u.Password,
		},
	}
}

type UserUpdateRequest struct {
	Email     *string `binding:"omitempty,email"  json:"email,omitempty"`
	Firstname *string `json:"firstname,omitempty"`
	Lastname  *string `json:"lastname,omitempty"`
	// Password  *string `binding:"min=8" json:"password,omitempty"`
}

type UserResponse struct {
	ID                string      `json:"id"`
	Email             string      `json:"email"`
	Firstname         string      `json:"firstname"`
	Lastname          string      `json:"lastname"`
	Role              db.Role     `json:"role"`
	ProfilePictureURL string      `json:"profile_picture_url"`
	CreatedAt         db.DateTime `json:"created_at"`
	UpdatedAt         db.DateTime `json:"updated_at"`
}

func (u *UserResponse) FromDbUser(model db.UserModel, ppURL string) {
	u.ID = model.ID
	u.Email = model.Email
	u.Firstname = model.Firstname
	u.Lastname = model.Lastname
	u.Role = model.Role
	u.ProfilePictureURL = ppURL
	u.CreatedAt = model.CreatedAt
	u.UpdatedAt = model.UpdatedAt
}

func DbUserToResponse(user db.UserModel, ppURL string) UserResponse {
	var resp UserResponse
	resp.FromDbUser(user, ppURL)
	return resp
}
