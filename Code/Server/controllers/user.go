package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"immotep/backend/models"
	"immotep/backend/prisma/db"
	"immotep/backend/services/database"
	"immotep/backend/services/filesystem"
	"immotep/backend/utils"
)

func getProfilePicture(user db.UserModel) string {
	ppURL := ""
	ppPath, ok := user.ProfilePicture()
	if ok {
		ppURL = filesystem.GetImageURL(ppPath)
	}
	return ppURL
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
	c.JSON(http.StatusOK, models.DbUserToResponse(*user, getProfilePicture(*user)))
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
	c.JSON(http.StatusOK, models.DbUserToResponse(*user, getProfilePicture(*user)))
}

// UpdateCurrentUserProfile godoc
//
//	@Summary		Update current user profile
//	@Description	Update current user profile
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			user	body		models.UserUpdateRequest	true	"User update info"
//	@Success		200		{object}	models.IdResponse			"Updated user ID"
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

	newUser := database.UpdateUser(*user, req)
	if newUser == nil {
		utils.SendError(c, http.StatusConflict, utils.EmailAlreadyExists, nil)
		return
	}
	c.JSON(http.StatusOK, models.IdResponse{ID: newUser.ID})
}

// UpdateCurrentUserProfilePicture godoc
//
//	@Summary		Update current user's profile picture
//	@Description	Update current user's profile picture
//	@Tags			user
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			id		path		string				true	"User ID"
//	@Param			picture	formData	file				true	"Profile picture"
//	@Success		201		{object}	models.IdResponse	"Updated user ID"
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

	file, err := c.FormFile("picture")
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, utils.MissingFile, err)
		return
	}

	fileInfo := filesystem.UploadUserProfileImage(user.ID, file)
	newUser := database.UpdateUserPicture(*user, fileInfo.Key)
	c.JSON(http.StatusOK, models.IdResponse{ID: newUser.ID})
}
