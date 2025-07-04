package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"keyz/backend/models"
	"keyz/backend/services/database"
	"keyz/backend/utils"
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
//	@Router			/users/ [get]
func GetAllUsers(c *gin.Context) {
	allUsers := database.GetAllUsers()
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
//	@Router			/users/{id}/ [get]
func GetUserByID(c *gin.Context) {
	user := database.GetUserByID(c.Param("id"))
	if user == nil {
		utils.SendError(c, http.StatusNotFound, utils.UserNotFound, nil)
		return
	}
	c.JSON(http.StatusOK, models.DbUserToResponse(*user))
}

// GetUserProfilePicture godoc
//
//	@Summary		Get user's picture
//	@Description	Get user's picture
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string					true	"User ID"
//	@Success		200	{object}	models.ImageResponse	"Image data"
//	@Success		204	"No picture associated"
//	@Failure		401	{object}	utils.Error	"Unauthorized"
//	@Failure		404	{object}	utils.Error	"User not found"
//	@Failure		500
//	@Security		Bearer
//	@Router			/users/{id}/picture/ [get]
func GetUserProfilePicture(c *gin.Context) {
	user := database.GetUserByID(c.Param("id"))
	if user == nil {
		utils.SendError(c, http.StatusNotFound, utils.UserNotFound, nil)
		return
	}

	pictureId, ok := user.ProfilePictureID()
	if !ok {
		c.Status(http.StatusNoContent)
		return
	}
	image := database.GetImageByID(pictureId)
	if image == nil {
		utils.SendError(c, http.StatusNotFound, utils.UserProfilePictureNotFound, nil)
		return
	}
	c.JSON(http.StatusOK, models.DbImageToResponse(*image))
}

// GetCurrentUserProfile godoc
//
//	@Summary		Get current user profile
//	@Description	Get current user profile
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	models.UserResponse	"User data"
//	@Failure		401	{object}	utils.Error			"Unauthorized"
//	@Failure		404	{object}	utils.Error			"User not found"
//	@Failure		500
//	@Security		Bearer
//	@Router			/profile/ [get]
func GetCurrentUserProfile(c *gin.Context) {
	claims := utils.GetClaims(c)
	user := database.GetUserByID(claims["id"])
	if user == nil {
		utils.SendError(c, http.StatusNotFound, utils.UserNotFound, nil)
		return
	}
	c.JSON(http.StatusOK, models.DbUserToResponse(*user))
}

// UpdateCurrentUserProfile godoc
//
//	@Summary		Update current user profile
//	@Description	Update current user profile
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			user	body		models.UserUpdateRequest	true	"User update info"
//	@Success		200		{object}	models.UserResponse			"User data"
//	@Failure		400		{object}	utils.Error					"Missing fields"
//	@Failure		404		{object}	utils.Error					"User not found"
//	@Failure		409		{object}	utils.Error					"Email already exists"
//	@Failure		500
//	@Security		Bearer
//	@Router			/profile/ [put]
func UpdateCurrentUserProfile(c *gin.Context) {
	claims := utils.GetClaims(c)
	user := database.GetUserByID(claims["id"])
	if user == nil {
		utils.SendError(c, http.StatusNotFound, utils.UserNotFound, nil)
		return
	}

	var req models.UserUpdateRequest
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, utils.MissingFields, err)
		return
	}
	if req.Email != nil {
		req.Email = utils.Ptr(utils.SanitizeEmail(*req.Email))
	}

	newUser := database.UpdateUser(*user, req)
	if newUser == nil {
		utils.SendError(c, http.StatusConflict, utils.EmailAlreadyExists, nil)
		return
	}
	c.JSON(http.StatusOK, models.DbUserToResponse(*newUser))
}

// GetCurrentUserProfilePicture godoc
//
//	@Summary		Get current user's profile picture
//	@Description	Get current user's profile picture
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string					true	"User ID"
//	@Success		200	{object}	models.ImageResponse	"Image data"
//	@Success		204	"No picture associated"
//	@Failure		401	{object}	utils.Error	"Unauthorized"
//	@Failure		404	{object}	utils.Error	"User not found"
//	@Failure		500
//	@Security		Bearer
//	@Router			/profile/picture/ [get]
func GetCurrentUserProfilePicture(c *gin.Context) {
	claims := utils.GetClaims(c)
	user := database.GetUserByID(claims["id"])
	if user == nil {
		utils.SendError(c, http.StatusNotFound, utils.UserNotFound, nil)
		return
	}

	pictureId, ok := user.ProfilePictureID()
	if !ok {
		c.Status(http.StatusNoContent)
		return
	}
	image := database.GetImageByID(pictureId)
	if image == nil {
		utils.SendError(c, http.StatusNotFound, utils.UserProfilePictureNotFound, nil)
		return
	}
	c.JSON(http.StatusOK, models.DbImageToResponse(*image))
}

// UpdateCurrentUserProfilePicture godoc
//
//	@Summary		Update current user's profile picture
//	@Description	Update current user's profile picture
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string				true	"User ID"
//	@Param			picture	body		models.ImageRequest	true	"Picture data as a Base64 string"
//	@Success		201		{object}	models.UserResponse	"Updated user data"
//	@Failure		400		{object}	utils.Error			"Missing fields or bad base64 string"
//	@Failure		401		{object}	utils.Error			"Unauthorized"
//	@Failure		404		{object}	utils.Error			"User not found"
//	@Failure		500
//	@Security		Bearer
//	@Router			/profile/picture/ [put]
func UpdateCurrentUserProfilePicture(c *gin.Context) {
	claims := utils.GetClaims(c)
	user := database.GetUserByID(claims["id"])
	if user == nil {
		utils.SendError(c, http.StatusNotFound, utils.UserNotFound, nil)
		return
	}

	var req models.ImageRequest
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, utils.MissingFields, err)
		return
	}

	image := req.ToDbImage()
	if image == nil {
		utils.SendError(c, http.StatusBadRequest, utils.BadBase64OrUnsupportedType, nil)
		return
	}
	newImage := database.CreateImage(*image)

	newUser := database.UpdateUserPicture(*user, newImage)
	if newUser == nil {
		utils.SendError(c, http.StatusInternalServerError, utils.FailedLinkImage, nil)
		return
	}
	c.JSON(http.StatusOK, models.DbUserToResponse(*newUser))
}
