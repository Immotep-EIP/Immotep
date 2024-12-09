package userservice

import (
	"immotep/backend/database"
	"immotep/backend/prisma/db"
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
		if db.IsErrNotFound(err) {
			return nil
		}
		panic(err)
	}
	return user
}

func Create(user db.UserModel, role db.Role) *db.UserModel {
	pdb := database.DBclient
	newUser, err := pdb.Client.User.CreateOne(
		db.User.Email.Set(user.Email),
		db.User.Password.Set(user.Password),
		db.User.Firstname.Set(user.Firstname),
		db.User.Lastname.Set(user.Lastname),
		db.User.Role.Set(role),
	).Exec(pdb.Context)
	if err != nil {
		if info, is := db.IsErrUniqueConstraint(err); is && info.Fields[0] == db.User.Email.Field() {
			return nil
		}
		panic(err)
	}
	return newUser
}
