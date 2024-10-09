package controllers

import (
	"immotep/backend/models"
	"immotep/backend/prisma/db"
	userservice "immotep/backend/services"
	"immotep/backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	allUsers := userservice.GetAll()
	c.JSON(http.StatusOK, utils.Map(allUsers, models.UserToResponse))
}

func GetUserByID(c *gin.Context) {
	user := userservice.GetByID(c.Params.ByName("id"))
	if user == nil {
		utils.SendError(c, http.StatusNotFound, utils.CannotFindUser, nil)
		return
	}
	c.JSON(http.StatusOK, models.UserToResponse(*user))
}

func GetProfile(c *gin.Context) {
	claims := utils.GetClaims(c)
	if claims == nil {
		utils.SendError(c, http.StatusUnauthorized, utils.NoClaims, nil)
		return
	}

	user := userservice.GetByID(claims["id"])
	if user == nil {
		utils.SendError(c, http.StatusNotFound, utils.CannotFindUser, nil)
		return
	}
	c.JSON(http.StatusOK, models.UserToResponse(*user))
}

func CreateUser(c *gin.Context) {
	var userResp db.UserModel
	err := c.ShouldBindBodyWithJSON(&userResp)
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, utils.CannotDecodeUser, err)
		return
	}

	if userResp.Email == "" || userResp.Firstname == "" || userResp.Lastname == "" || userResp.Password == "" {
		utils.SendError(c, http.StatusBadRequest, utils.MissingFields, nil)
		return
	}

	userResp.Password, err = utils.HashPassword(userResp.Password)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, utils.CannotHashPassword, err)
		return
	}

	user := userservice.Create(userResp)
	if user == nil {
		utils.SendError(c, http.StatusConflict, utils.EmailAlreadyExists, err)
		return
	}
	c.JSON(http.StatusCreated, models.UserToResponse(*user))
}
