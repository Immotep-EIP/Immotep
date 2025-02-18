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
	allProperties := database.GetAllPropertyByOwnerId(claims["id"])
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

	if database.GetCurrentActiveContract(c.Param("property_id")) != nil {
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
		if database.GetTenantCurrentActiveContract(user.ID) != nil {
			utils.SendError(c, http.StatusConflict, utils.TenantAlreadyHasContract, nil)
			return false
		}
	}
	return true
}

// EndContract godoc
//
//	@Summary		End contract
//	@Description	End active contract for a property
//	@Tags			owner
//	@Accept			json
//	@Produce		json
//	@Param			property_id	path	string	true	"Property ID"
//	@Success		204			"Contract ended"
//	@Failure		403			{object}	utils.Error	"Property is not yours"
//	@Failure		404			{object}	utils.Error	"No active contract"
//	@Failure		500
//	@Security		Bearer
//	@Router			/owner/properties/{property_id}/end-contract [put]
func EndContract(c *gin.Context) {
	currentActive := database.GetCurrentActiveContract(c.Param("property_id"))
	_, ok := currentActive.EndDate()
	if !ok {
		now := time.Now().Truncate(time.Minute)
		database.EndContract(currentActive.ID, &now)
		c.Status(http.StatusNoContent)
	} else {
		database.EndContract(currentActive.PropertyID, nil)
		c.Status(http.StatusNoContent)
	}
}

// ArchiveProperty godoc
//
//	@Summary		Archive property by ID
//	@Description	Archive a property by its ID
//	@Tags			owner
//	@Accept			json
//	@Produce		json
//	@Param			property_id	path		string					true	"Property ID"
//	@Success		200			{object}	models.PropertyResponse	"Archived property data"
//	@Failure		403			{object}	utils.Error				"Property not yours"
//	@Failure		404			{object}	utils.Error				"Property not found"
//	@Failure		500
//	@Security		Bearer
//	@Router			/owner/properties/{property_id}/ [delete]
func ArchiveProperty(c *gin.Context) {
	property := database.ArchiveProperty(c.Param("property_id"))
	c.JSON(http.StatusOK, models.DbPropertyToResponse(*property))
}
