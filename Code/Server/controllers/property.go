package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"immotep/backend/models"
	"immotep/backend/prisma/db"
	"immotep/backend/services/brevo"
	"immotep/backend/services/database"
	"immotep/backend/utils"
)

// GetAllProperties godoc
//
//	@Summary		Get all properties of an owner
//	@Description	Get all properties information of an owner
//	@Tags			owner
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		models.PropertyResponse	"List of properties"
//	@Failure		401	{object}	utils.Error				"Unauthorized"
//	@Failure		500
//	@Security		Bearer
//	@Router			/owner/properties/ [get]
func GetAllProperties(c *gin.Context) {
	claims := utils.GetClaims(c)
	allProperties := database.GetAllPropertyByOwnerId(claims["id"], false)
	c.JSON(http.StatusOK, utils.Map(allProperties, models.DbPropertyToResponse))
}

// GetAllArchivedProperties godoc
//
//	@Summary		Get all archived properties of an owner
//	@Description	Get all archived properties information of an owner
//	@Tags			owner
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		models.PropertyResponse	"List of archived properties"
//	@Failure		401	{object}	utils.Error				"Unauthorized"
//	@Failure		500
//	@Security		Bearer
//	@Router			/owner/properties/archived/ [get]
func GetAllArchivedProperties(c *gin.Context) {
	claims := utils.GetClaims(c)
	allProperties := database.GetAllPropertyByOwnerId(claims["id"], true)
	c.JSON(http.StatusOK, utils.Map(allProperties, models.DbPropertyToResponse))
}

// GetPropertyById godoc
//
//	@Summary		Get property by ID
//	@Description	Get property information by its ID
//	@Tags			owner
//	@Accept			json
//	@Produce		json
//	@Param			property_id	path		string					true	"Property ID"
//	@Success		200			{object}	models.PropertyResponse	"Property data"
//	@Failure		401			{object}	utils.Error				"Unauthorized"
//	@Failure		403			{object}	utils.Error				"Property not yours"
//	@Failure		404			{object}	utils.Error				"Property not found"
//	@Failure		500
//	@Security		Bearer
//	@Router			/owner/properties/{property_id}/ [get]
func GetPropertyById(c *gin.Context) {
	property := database.GetPropertyByID(c.Param("property_id"))
	c.JSON(http.StatusOK, models.DbPropertyToResponse(*property))
}

// GetPropertyInventory godoc
//
//	@Summary		Get property inventory by ID
//	@Description	Get property information by its ID with inventory including rooms and furnitures
//	@Tags			owner
//	@Accept			json
//	@Produce		json
//	@Param			property_id	path		string								true	"Property ID"
//	@Success		200			{object}	models.PropertyInventoryResponse	"Property inventory data"
//	@Failure		401			{object}	utils.Error							"Unauthorized"
//	@Failure		403			{object}	utils.Error							"Property not yours"
//	@Failure		404			{object}	utils.Error							"Property not found"
//	@Failure		500
//	@Security		Bearer
//	@Router			/owner/properties/{property_id}/inventory/ [get]
func GetPropertyInventory(c *gin.Context) {
	property := database.GetPropertyInventory(c.Param("property_id"))
	c.JSON(http.StatusOK, models.DbPropertyInventoryToResponse(*property))
}

// CreateProperty godoc
//
//	@Summary		Create a new property
//	@Description	Create a new property for an owner
//	@Tags			owner
//	@Accept			json
//	@Produce		json
//	@Param			user	body		models.PropertyRequest	true	"Property data"
//	@Success		201		{object}	models.PropertyResponse	"Created property data"
//	@Failure		400		{object}	utils.Error				"Missing fields"
//	@Failure		409		{object}	utils.Error				"Property already exists"
//	@Failure		500
//	@Security		Bearer
//	@Router			/owner/properties/ [post]
func CreateProperty(c *gin.Context) {
	claims := utils.GetClaims(c)

	var req models.PropertyRequest
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, utils.MissingFields, err)
		return
	}

	property := database.CreateProperty(req.ToDbProperty(), claims["id"])
	if property == nil {
		utils.SendError(c, http.StatusConflict, utils.PropertyAlreadyExists, nil)
		return
	}
	c.JSON(http.StatusCreated, models.DbPropertyToResponse(*property))
}

// UpdateProperty godoc
//
//	@Summary		Update property by ID
//	@Description	Update property information by its ID
//	@Tags			owner
//	@Accept			json
//	@Produce		json
//	@Param			property_id	path		string					true	"Property ID"
//	@Success		200			{object}	models.PropertyResponse	"Property data"
//	@Failure		401			{object}	utils.Error				"Unauthorized"
//	@Failure		403			{object}	utils.Error				"Property not yours"
//	@Failure		404			{object}	utils.Error				"Property not found"
//	@Failure		500
//	@Security		Bearer
//	@Router			/owner/properties/{property_id}/ [put]
func UpdateProperty(c *gin.Context) {
	var req models.PropertyUpdateRequest
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, utils.MissingFields, err)
		return
	}

	property := database.UpdateProperty(c.Param("property_id"), req)
	if property == nil {
		utils.SendError(c, http.StatusConflict, utils.PropertyAlreadyExists, nil)
		return
	}
	c.JSON(http.StatusOK, models.DbPropertyToResponse(*property))
}

// GetPropertyPicture godoc
//
//	@Summary		Get property's picture
//	@Description	Get property's picture
//	@Tags			owner
//	@Accept			json
//	@Produce		json
//	@Param			property_id	path		string					true	"Property ID"
//	@Success		201			{object}	models.ImageResponse	"Image data"
//	@Success		204			"No picture associated"
//	@Failure		401			{object}	utils.Error	"Unauthorized"
//	@Failure		403			{object}	utils.Error	"Property not yours"
//	@Failure		404			{object}	utils.Error	"Property not found"
//	@Failure		500
//	@Security		Bearer
//	@Router			/owner/properties/{property_id}/picture/ [get]
func GetPropertyPicture(c *gin.Context) {
	property := database.GetPropertyByID(c.Param("property_id"))
	pictureId, ok := property.PictureID()
	if !ok {
		c.Status(http.StatusNoContent)
		return
	}
	image := database.GetImageByID(pictureId)
	if image == nil {
		utils.SendError(c, http.StatusNotFound, utils.PropertyPictureNotFound, nil)
		return
	}
	c.JSON(http.StatusOK, models.DbImageToResponse(*image))
}

// UpdatePropertyPicture godoc
//
//	@Summary		Update property's picture
//	@Description	Update property's picture
//	@Tags			owner
//	@Accept			json
//	@Produce		json
//	@Param			property_id	path		string					true	"Property ID"
//	@Param			picture		body		models.ImageRequest		true	"Picture data as a Base64 string"
//	@Success		201			{object}	models.PropertyResponse	"Updated property data"
//	@Failure		400			{object}	utils.Error				"Missing fields or bad base64 string"
//	@Failure		401			{object}	utils.Error				"Unauthorized"
//	@Failure		403			{object}	utils.Error				"Property not yours"
//	@Failure		404			{object}	utils.Error				"Property not found"
//	@Failure		500
//	@Security		Bearer
//	@Router			/owner/properties/{property_id}/picture/ [put]
func UpdatePropertyPicture(c *gin.Context) {
	var req models.ImageRequest
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, utils.MissingFields, err)
		return
	}

	image := req.ToDbImage()
	if image == nil {
		utils.SendError(c, http.StatusBadRequest, utils.BadBase64String, nil)
		return
	}
	newImage := database.CreateImage(*image)

	property := database.GetPropertyByID(c.Param("property_id"))
	newProperty := database.UpdatePropertyPicture(*property, newImage)
	if newProperty == nil {
		utils.SendError(c, http.StatusInternalServerError, utils.FailedLinkImage, nil)
		return
	}
	c.JSON(http.StatusOK, models.DbPropertyToResponse(*newProperty))
}

// InviteTenant godoc
//
//	@Summary		Invite tenant to owner's property
//	@Description	Invite tenant to owner's property
//	@Tags			owner
//	@Accept			json
//	@Produce		json
//	@Param			property_id	path		string					true	"Property ID"
//	@Param			user		body		models.InviteRequest	true	"Invite params"
//	@Success		200			{object}	models.InviteResponse	"Created invite"
//	@Failure		400			{object}	utils.Error				"Missing fields"
//	@Failure		403			{object}	utils.Error				"Property is not yours"
//	@Failure		404			{object}	utils.Error				"Property not found"
//	@Failure		409			{object}	utils.Error				"Invite already exists for this email"
//	@Failure		500
//	@Security		Bearer
//	@Router			/owner/properties/{property_id}/send-invite [post]
func InviteTenant(c *gin.Context) {
	var inviteReq models.InviteRequest
	err := c.ShouldBindBodyWithJSON(&inviteReq)
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, utils.MissingFields, err)
		return
	}

	if database.GetCurrentActiveLease(c.Param("property_id")) != nil {
		utils.SendError(c, http.StatusConflict, utils.PropertyNotAvailable, nil)
		return
	}

	user := database.GetUserByEmail(inviteReq.TenantEmail)
	if !checkInvitedTenant(c, user) {
		return
	}

	pendingContract := database.CreatePendingContract(inviteReq.ToDbPendingContract(), c.Param("property_id"))
	if pendingContract == nil {
		utils.SendError(c, http.StatusConflict, utils.InviteAlreadyExists, nil)
		return
	}

	res, err := brevo.SendEmailInvite(*pendingContract, user != nil)
	if err != nil {
		log.Println(res, err.Error())
		utils.SendError(c, http.StatusInternalServerError, utils.FailedSendEmail, err)
		return
	}

	c.JSON(http.StatusOK, models.DbPendingContractToResponse(*pendingContract))
}

func checkInvitedTenant(c *gin.Context, user *db.UserModel) bool {
	if user != nil {
		if user.Role != db.RoleTenant {
			utils.SendError(c, http.StatusConflict, utils.UserAlreadyExistsAsOwner, nil)
			return false
		}
		if database.GetTenantCurrentActiveLease(user.ID) != nil {
			utils.SendError(c, http.StatusConflict, utils.TenantAlreadyHasLease, nil)
			return false
		}
	}
	return true
}

// EndLease godoc
//
//	@Summary		End lease
//	@Description	End active lease for a property
//	@Tags			owner
//	@Accept			json
//	@Produce		json
//	@Param			property_id	path	string	true	"Property ID"
//	@Success		204			"Lease ended"
//	@Failure		403			{object}	utils.Error	"Property is not yours"
//	@Failure		404			{object}	utils.Error	"No active lease"
//	@Failure		500
//	@Security		Bearer
//	@Router			/owner/properties/{property_id}/end-lease [put]
func EndLease(c *gin.Context) {
	currentActive := database.GetCurrentActiveLease(c.Param("property_id"))
	_, ok := currentActive.EndDate()
	if !ok {
		now := time.Now().Truncate(time.Minute)
		database.EndLease(currentActive.ID, &now)
		c.Status(http.StatusNoContent)
	} else {
		database.EndLease(currentActive.PropertyID, nil)
		c.Status(http.StatusNoContent)
	}
}

// CancelInvite godoc
//
//	@Summary		Cancel invite
//	@Description	Cancel pending lease invite
//	@Tags			owner
//	@Accept			json
//	@Produce		json
//	@Param			property_id	path	string	true	"Property ID"
//	@Success		204			"Invite canceled"
//	@Failure		403			{object}	utils.Error	"Property is not yours"
//	@Failure		404			{object}	utils.Error	"No pending lease"
//	@Failure		500
//	@Security		Bearer
//	@Router			/owner/properties/{property_id}/cancel-invite [delete]
func CancelInvite(c *gin.Context) {
	database.DeleteCurrentPendingContract(c.Param("property_id"))
	c.Status(http.StatusNoContent)
}

// ArchiveProperty godoc
//
//	@Summary		Toggle archive property by ID
//	@Description	Toggle archive status of a property by its ID
//	@Tags			owner
//	@Accept			json
//	@Produce		json
//	@Param			property_id	path		string					true	"Property ID"
//	@Param			archive		body		models.ArchiveRequest	true	"Archive status"
//	@Success		200			{object}	models.PropertyResponse	"Toggled archive property data"
//	@Failure		400			{object}	utils.Error				"Mising fields"
//	@Failure		403			{object}	utils.Error				"Property not yours"
//	@Failure		404			{object}	utils.Error				"Property not found"
//	@Failure		500
//	@Security		Bearer
//	@Router			/owner/properties/{property_id}/archive [put]
func ArchiveProperty(c *gin.Context) {
	var req models.ArchiveRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendError(c, http.StatusBadRequest, utils.MissingFields, err)
		return
	}

	property := database.ToggleArchiveProperty(c.Param("property_id"), req.Archive)
	c.JSON(http.StatusOK, models.DbPropertyToResponse(*property))
}
