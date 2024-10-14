package controllers

import (
	"immotep/backend/models"
	userservice "immotep/backend/services"
	"immotep/backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllUsers godoc
//
//	@Summary		Get all users
//	@Description	Get all users information
//	@Produce		json
//	@Success		200	{array}	models.UserResponse	"List of users"
//	@Failure		500
//	@Security		Bearer
//	@Router			/users [get]
func GetAllUsers(c *gin.Context) {
	allUsers := userservice.GetAll()
	c.JSON(http.StatusOK, utils.Map(allUsers, models.UserToResponse))
}

// GetUserByID godoc
//
//	@Summary		Get user by ID
//	@Description	Get user information by its ID
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string				true	"User ID"
//	@Success		200	{object}	models.UserResponse	"User data"
//	@Failure		401	{object}	utils.Error			"Unauthorized"
//	@Failure		404	{object}	utils.Error			"Cannot find user"
//	@Failure		500
//	@Security		Bearer
//	@Router			/users/{id} [get]
func GetUserByID(c *gin.Context) {
	user := userservice.GetByID(c.Param("id"))
	if user == nil {
		utils.SendError(c, http.StatusNotFound, utils.CannotFindUser, nil)
		return
	}
	c.JSON(http.StatusOK, models.UserToResponse(*user))
}

// GetProfile godoc
//
//	@Summary		Get user profile
//	@Description	Get user profile information
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	models.UserResponse	"User data"
//	@Failure		401	{object}	utils.Error			"Unauthorized"
//	@Failure		404	{object}	utils.Error			"Cannot find user"
//	@Failure		500
//	@Security		Bearer
//	@Router			/profile [get]
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

// CreateUser godoc
//
//	@Summary		Create a new user
//	@Description	Create a new user
//	@Accept			json
//	@Produce		json
//	@Param			user	body		models.UserRequest	true	"User data"
//	@Success		201		{object}	models.UserResponse	"Created user data"
//	@Failure		400		{object}	utils.Error			"Cannot decode user"
//	@Failure		400		{object}	utils.Error			"Missing fields"
//	@Failure		409		{object}	utils.Error			"Email already exists"
//	@Failure		500
//	@Router			/auth/register [post]
func CreateUser(c *gin.Context) {
	var userReq models.UserRequest
	err := c.ShouldBindBodyWithJSON(&userReq)
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, utils.MissingFields, err)
		return
	}

	userReq.Password, err = utils.HashPassword(userReq.Password)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, utils.CannotHashPassword, err)
		return
	}

	user := userservice.Create(userReq.ToUser())
	if user == nil {
		utils.SendError(c, http.StatusConflict, utils.EmailAlreadyExists, err)
		return
	}
	c.JSON(http.StatusCreated, models.UserToResponse(*user))
}
