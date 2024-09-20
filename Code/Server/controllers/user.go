package controllers

import (
	"encoding/json"
	"fmt"
	"immotep/backend/database"
	"immotep/backend/prisma/db"
	"immotep/backend/utils"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	pdb := database.DBclient
	allUsers, err := pdb.Client.User.FindMany().Exec(pdb.Context)
	if err != nil {
		fmt.Println("Cannot fetch users")
		return
	}
	usersMap := make(map[string]interface{})
	usersMap["users"] = allUsers
	utils.WriteJSON(w, http.StatusOK, usersMap)
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	pdb := database.DBclient
	user, err := pdb.Client.User.FindUnique(db.User.ID.Equals(id)).Exec(pdb.Context)
	if err != nil {
		if strings.Contains(err.Error(), "ErrNotFound") {
			utils.WriteError(w, http.StatusNotFound, "Cannot find user", err)
		} else {
			utils.WriteError(w, http.StatusInternalServerError, "Cannot fetch user", err)
		}
		return
	}
	err = utils.WriteJSON(w, http.StatusOK, user)
	if err != nil {
		fmt.Println("Cannot form response")
		return
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var userResp db.UserModel

	err := json.NewDecoder(r.Body).Decode(&userResp)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Cannot decode user", err)
		return
	}

	if userResp.Email == "" || userResp.Firstname == "" || userResp.Lastname == "" || userResp.Password == "" {
		utils.WriteError(w, http.StatusBadRequest, "Bad request, missing fields", nil)
		return
	}

	pdb := database.DBclient
	user, err := pdb.Client.User.CreateOne(
		db.User.Email.Set(userResp.Email),
		db.User.Password.Set(userResp.Password), //! TODO: Hash password
		db.User.Firstname.Set(userResp.Firstname),
		db.User.Lastname.Set(userResp.Lastname),
	).Exec(pdb.Context)

	if err != nil {
		if strings.Contains(err.Error(), "Unique constraint failed") && strings.Contains(err.Error(), "email") {
			utils.WriteError(w, http.StatusConflict, "Email already exists", err)
		} else {
			utils.WriteError(w, http.StatusInternalServerError, "Cannot create user", err)
		}
		return
	}
	utils.WriteJSON(w, http.StatusOK, user)
}
