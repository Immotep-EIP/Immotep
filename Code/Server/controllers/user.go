package controllers

import (
	"fmt"
	"immotep/backend/database"
	"immotep/backend/models"
	"immotep/backend/prisma/db"
	"immotep/backend/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	pdb := database.DBclient
	allUsers, err := pdb.Client.User.FindMany().Exec(pdb.Context)
	if err != nil {
		fmt.Println("Cannot fetch users")
		return
	}
	c.JSON(http.StatusOK, utils.Map(allUsers, models.UserToResponse))
}

func GetUserByID(c *gin.Context) {
	id := c.Params.ByName("id")
	pdb := database.DBclient
	user, err := pdb.Client.User.FindUnique(db.User.ID.Equals(id)).Exec(pdb.Context)
	if err != nil {
		if strings.Contains(err.Error(), "ErrNotFound") {
			c.JSON(http.StatusNotFound, gin.H{"error": "Cannot find user"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot fetch user"})
		}
		return
	}
	c.JSON(http.StatusOK, models.UserToResponse(*user))
}

func GetProfile(c *gin.Context) {
	claims := utils.GetClaims(c)
	if claims == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	pdb := database.DBclient
	user, err := pdb.Client.User.FindUnique(db.User.ID.Equals(claims["id"])).Exec(pdb.Context)
	if err != nil {
		if strings.Contains(err.Error(), "ErrNotFound") {
			c.JSON(http.StatusNotFound, gin.H{"error": "Cannot find user"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot fetch user"})
		}
		return
	}
	c.JSON(http.StatusOK, models.UserToResponse(*user))
}

func CreateUser(c *gin.Context) {
	var userResp db.UserModel
	// err := json.NewDecoder(c.Request.Body).Decode(&userResp) // old with chi
	err := c.ShouldBindBodyWithJSON(&userResp)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot decode user"})
		return
	}

	if userResp.Email == "" || userResp.Firstname == "" || userResp.Lastname == "" || userResp.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request, missing fields"})
		return
	}

	hashedPassword, err := utils.HashPassword(userResp.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot hash password"})
		return
	}

	pdb := database.DBclient
	user, err := pdb.Client.User.CreateOne(
		db.User.Email.Set(userResp.Email),
		db.User.Password.Set(hashedPassword),
		db.User.Firstname.Set(userResp.Firstname),
		db.User.Lastname.Set(userResp.Lastname),
	).Exec(pdb.Context)

	if err != nil {
		if strings.Contains(err.Error(), "Unique constraint failed") && strings.Contains(err.Error(), "email") {
			c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot create user"})
		}
		return
	}
	c.JSON(http.StatusOK, models.UserToResponse(*user))
}
