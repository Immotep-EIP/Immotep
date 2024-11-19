package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"immotep/backend/models"
	userservice "immotep/backend/services/user"
	"immotep/backend/utils"
)

// GetAllUsers godoc
//
//	@Summary		Get all users
//	@Description	Get all users information
//	@Tags			user
//	@Produce		json
//	@Success		200	{array}	models.UserResponse	"List of users"
//	@Failure		500
//	@Security		Bearer
//	@Router			/users [get]
func GetAllUsers(c *gin.Context) {
	allUsers := userservice.GetAll()
	c.JSON(http.StatusOK, utils.Map(allUsers, models.DbUserToResponse))
}

// GetUserByID godoc
//
//	@Summary		Get user by ID
//	@Description	Get user information by its ID
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string				true	"User ID"
//	@Success		200	{object}	models.UserResponse	"User data"
//	@Failure		401	{object}	utils.Error			"Unauthorized"
//	@Failure		404	{object}	utils.Error			"User not found"
//	@Failure		500
//	@Security		Bearer
//	@Router			/users/{id} [get]
func GetUserByID(c *gin.Context) {
	user := userservice.GetByID(c.Param("id"))
	if user == nil {
		utils.SendError(c, http.StatusNotFound, utils.UserNotFound, nil)
		return
	}
	c.JSON(http.StatusOK, models.DbUserToResponse(*user))
}

// GetProfile godoc
//
//	@Summary		Get user profile
//	@Description	Get user profile information
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	models.UserResponse	"User data"
//	@Failure		401	{object}	utils.Error			"Unauthorized"
//	@Failure		404	{object}	utils.Error			"User not found"
//	@Failure		500
//	@Security		Bearer
//	@Router			/profile [get]
func GetProfile(c *gin.Context) {
	claims := utils.GetClaims(c)

	user := userservice.GetByID(claims["id"])
	if user == nil {
		utils.SendError(c, http.StatusNotFound, utils.UserNotFound, nil)
		return
	}
	c.JSON(http.StatusOK, models.DbUserToResponse(*user))
}
