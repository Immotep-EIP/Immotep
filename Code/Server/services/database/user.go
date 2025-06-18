package database

import (
	"strings"

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

func MockGetAllUsers(c *services.PrismaDB) db.UserMockExpectParam {
	return c.Client.User.FindMany()
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

func MockGetUserByID(c *services.PrismaDB) db.UserMockExpectParam {
	return c.Client.User.FindUnique(
		db.User.ID.Equals("1"),
	)
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

func MockGetUserByEmail(c *services.PrismaDB) db.UserMockExpectParam {
	return c.Client.User.FindUnique(
		db.User.Email.Equals("test@example.com"),
	)
}

func CreateUser(user db.UserModel, role db.Role) *db.UserModel {
	pdb := services.DBclient
	newUser, err := pdb.Client.User.CreateOne(
		db.User.Email.Set(strings.ToLower(user.Email)),
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

func MockCreateUser(c *services.PrismaDB, user db.UserModel) db.UserMockExpectParam {
	return c.Client.User.CreateOne(
		db.User.Email.Set(strings.ToLower(user.Email)),
		db.User.Password.Set(user.Password),
		db.User.Firstname.Set(user.Firstname),
		db.User.Lastname.Set(user.Lastname),
		db.User.Role.Set(db.RoleOwner),
	)
}

func UpdateUser(user db.UserModel, req models.UserUpdateRequest) *db.UserModel {
	if req.Email != nil {
		*req.Email = strings.ToLower(*req.Email)
	}

	pdb := services.DBclient
	newUser, err := pdb.Client.User.FindUnique(db.User.ID.Equals(user.ID)).Update(
		db.User.Email.SetIfPresent(req.Email),
		db.User.Firstname.SetIfPresent(req.Firstname),
		db.User.Lastname.SetIfPresent(req.Lastname),
	).Exec(pdb.Context)
	if err != nil {
		if info, is := db.IsErrUniqueConstraint(err); is && info.Fields[0] == db.User.Email.Field() {
			return nil
		}
		panic(err)
	}
	return newUser
}

func MockUpdateUser(c *services.PrismaDB, uUser models.UserUpdateRequest) db.UserMockExpectParam {
	if uUser.Email != nil {
		*uUser.Email = strings.ToLower(*uUser.Email)
	}

	return c.Client.User.FindUnique(
		db.User.ID.Equals("1"),
	).Update(
		db.User.Email.SetIfPresent(uUser.Email),
		db.User.Firstname.SetIfPresent(uUser.Firstname),
		db.User.Lastname.SetIfPresent(uUser.Lastname),
	)
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

func MockUpdateUserPicture(c *services.PrismaDB) db.UserMockExpectParam {
	return c.Client.User.FindUnique(
		db.User.ID.Equals("1"),
	).Update(
		db.User.ProfilePicture.Link(db.Image.ID.Equals("1")),
	)
}
