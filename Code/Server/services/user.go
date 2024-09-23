package userservice

import (
	"immotep/backend/database"
	"immotep/backend/prisma/db"
	"strings"
)

func GetAll() []db.UserModel {
	pdb := database.DBclient
	allUsers, err := pdb.Client.User.FindMany().Exec(pdb.Context)
	if err != nil {
		panic(err)
	}
	return allUsers
}

func GetByID(id string) *db.UserModel {
	pdb := database.DBclient
	user, err := pdb.Client.User.FindUnique(db.User.ID.Equals(id)).Exec(pdb.Context)
	if err != nil {
		if strings.Contains(err.Error(), "ErrNotFound") {
			return nil
		} else {
			panic(err)
		}
	}
	return user
}

func Create(user db.UserModel) *db.UserModel {
	pdb := database.DBclient
	newUser, err := pdb.Client.User.CreateOne(
		db.User.Email.Set(user.Email),
		db.User.Password.Set(user.Password),
		db.User.Firstname.Set(user.Firstname),
		db.User.Lastname.Set(user.Lastname),
	).Exec(pdb.Context)
	if err != nil {
		if strings.Contains(err.Error(), "Unique constraint failed") && strings.Contains(err.Error(), "email") {
			return nil
		} else {
			panic(err)
		}
	}
	return newUser
}
