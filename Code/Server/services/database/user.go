package database

import (
	"immotep/backend/models"
	"immotep/backend/prisma/db"
	"immotep/backend/services"
)

func GetAllUsers() []db.UserModel {
	pdb := services.DBclient
	allUsers, err := pdb.Client.User.FindMany().Exec(pdb.Context)
	if err != nil {
		panic(err)
	}
	return allUsers
}

func GetUserByID(id string) *db.UserModel {
	pdb := services.DBclient
	user, err := pdb.Client.User.FindUnique(db.User.ID.Equals(id)).Exec(pdb.Context)
	if err != nil {
		if db.IsErrNotFound(err) {
			return nil
		}
		panic(err)
	}
	return user
}

func GetUserByEmail(email string) *db.UserModel {
	pdb := services.DBclient
	user, err := pdb.Client.User.FindUnique(db.User.Email.Equals(email)).Exec(pdb.Context)
	if err != nil {
		if db.IsErrNotFound(err) {
			return nil
		}
		panic(err)
	}
	return user
}

func CreateUser(user db.UserModel, role db.Role) *db.UserModel {
	pdb := services.DBclient
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

func UpdateUser(id string, user models.UserUpdateRequest) *db.UserModel {
	pdb := services.DBclient
	newUser, err := pdb.Client.User.FindUnique(db.User.ID.Equals(id)).Update(
		db.User.Email.SetIfPresent(user.Email),
		db.User.Firstname.SetIfPresent(user.Firstname),
		db.User.Lastname.SetIfPresent(user.Lastname),
	).Exec(pdb.Context)
	if err != nil {
		if db.IsErrNotFound(err) {
			return nil
		}
		if info, is := db.IsErrUniqueConstraint(err); is && info.Fields[0] == db.User.Email.Field() {
			return nil
		}
		panic(err)
	}
	return newUser
}

func UpdateUserPicture(user db.UserModel, image db.ImageModel) *db.UserModel {
	pdb := services.DBclient
	newUser, err := pdb.Client.User.FindUnique(
		db.User.ID.Equals(user.ID),
	).Update(
		db.User.ProfilePicture.Link(db.Image.ID.Equals(image.ID)),
	).Exec(pdb.Context)
	if err != nil {
		if db.IsErrNotFound(err) {
			return nil
		}
		panic(err)
	}
	return newUser
}
