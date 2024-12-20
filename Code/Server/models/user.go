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

func (u *UserRequest) ToUser() db.UserModel {
	return db.UserModel{
		InnerUser: db.InnerUser{
			Email:     u.Email,
			Firstname: u.Firstname,
			Lastname:  u.Lastname,
			Password:  u.Password,
		},
	}
}

type UserResponse struct {
	ID        string      `json:"id"`
	Email     string      `json:"email"`
	Firstname string      `json:"firstname"`
	Lastname  string      `json:"lastname"`
	Role      db.Role     `json:"role"`
	CreatedAt db.DateTime `json:"created_at"`
	UpdatedAt db.DateTime `json:"updated_at"`
}

func (u *UserResponse) FromUser(user db.UserModel) {
	u.ID = user.ID
	u.Email = user.Email
	u.Firstname = user.Firstname
	u.Lastname = user.Lastname
	u.Role = user.Role
	u.CreatedAt = user.CreatedAt
	u.UpdatedAt = user.UpdatedAt
}

func UserToResponse(user db.UserModel) UserResponse {
	var userResp UserResponse
	userResp.FromUser(user)
	return userResp
}
