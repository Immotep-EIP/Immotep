package controllers

import (
	"log"
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
	"keyz/backend/models"
	"keyz/backend/prisma/db"
	"keyz/backend/services/brevo"
	"keyz/backend/services/database"
	"keyz/backend/utils"
)

// GetPropertiesByOwner godoc
//
//	@Summary		Get properties of an owner
//	@Description	Get properties information of an owner, optionally filtered by archive status
//	@Tags			property
//	@Accept			json
//	@Produce		json
//	@Param			archive	query		bool					false	"Filter by archive status (default: false)"
//	@Success		200		{array}		models.PropertyResponse	"List of properties"
//	@Failure		401		{object}	utils.Error				"Unauthorized"
//	@Failure		500
//	@Security		Bearer
//	@Router			/owner/properties/ [get]
func GetPropertiesByOwner(c *gin.Context) {
	claims := utils.GetClaims(c)
	archive := c.DefaultQuery("archive", "false") == utils.Strue
	allProperties := database.GetPropertiesByOwnerId(claims["id"], archive)
	c.JSON(http.StatusOK, utils.Map(allProperties, func(property db.PropertyModel) models.PropertyResponse {
		return models.DbPropertyToResponse(property, "current")
	}))
}

// GetProperty godoc
//
//	@Summary		Get property by ID
//	@Description	Get property information by its ID
//	@Tags			property
//	@Accept			json
//	@Produce		json
//	@Param			property_id	path		string					true	"Property ID"
//	@Param			lease_id	path		string					true	"Lease ID or `current`"
//	@Param			lease_id	query		string					false	"ONLY FOR OWNER: optional Lease ID (default: current)"
//	@Success		200			{object}	models.PropertyResponse	"Property data"
//	@Failure		401			{object}	utils.Error				"Unauthorized"
//	@Failure		403			{object}	utils.Error				"Property not yours"
//	@Failure		404			{object}	utils.Error				"Property not found"
//	@Failure		500
//	@Security		Bearer
//	@Router			/owner/properties/{property_id}/ [get]
//	@Router			/tenant/leases/{lease_id}/property/ [get]
func GetProperty(c *gin.Context) {
	claims := utils.GetClaims(c)

	var leaseId string
	if claims["role"] == string(db.RoleTenant) {
		leaseId = c.Param("lease_id")
	} else {
		leaseId = c.DefaultQuery("lease_id", "current")
	}

	property, _ := c.MustGet("property").(db.PropertyModel)
	c.JSON(http.StatusOK, models.DbPropertyToResponse(property, leaseId))
}

// GetPropertyInventory godoc
//
//	@Summary		Get property inventory by ID
//	@Description	Get property information by its ID with inventory including rooms and furnitures
//	@Tags			property,inventory
//	@Accept			json
//	@Produce		json
//	@Param			property_id	path		string								true	"Property ID"
//	@Param			lease_id	path		string								true	"Lease ID or `current`"
//	@Param			lease_id	query		string								false	"ONLY FOR OWNER: optional Lease ID (default: current)"
//	@Success		200			{object}	models.PropertyInventoryResponse	"Property inventory data"
//	@Failure		401			{object}	utils.Error							"Unauthorized"
//	@Failure		403			{object}	utils.Error							"Property not yours"
//	@Failure		404			{object}	utils.Error							"Property not found"
//	@Failure		500
//	@Security		Bearer
//	@Router			/owner/properties/{property_id}/inventory/ [get]
//	@Router			/tenant/leases/{lease_id}/property/inventory/ [get]
func GetPropertyInventory(c *gin.Context) {
	claims := utils.GetClaims(c)

	var leaseId string
	if claims["role"] == string(db.RoleTenant) {
		leaseId = c.Param("lease_id")
	} else {
		leaseId = c.DefaultQuery("lease_id", "current")
	}

	property, _ := c.MustGet("property").(db.PropertyModel)
	propertyInv := database.GetPropertyInventory(property.ID)
	c.JSON(http.StatusOK, models.DbPropertyInventoryToResponse(*propertyInv, leaseId))
}

// CreateProperty godoc
//
//	@Summary		Create a new property
//	@Description	Create a new property for an owner
//	@Tags			property
//	@Accept			json
//	@Produce		json
//	@Param			property	body		models.PropertyRequest	true	"Property data"
//	@Success		201			{object}	models.IdResponse		"Created property ID"
//	@Failure		400			{object}	utils.Error				"Missing fields"
//	@Failure		409			{object}	utils.Error				"Property already exists"
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
	c.JSON(http.StatusCreated, models.IdResponse{ID: property.ID})
}

// UpdateProperty godoc
//
//	@Summary		Update property by ID
//	@Description	Update property information by its ID
//	@Tags			property
//	@Accept			json
//	@Produce		json
//	@Param			property_id	path		string							true	"Property ID"
//	@Param			property	body		models.PropertyUpdateRequest	true	"Property data"
//	@Success		200			{object}	models.IdResponse				"Updated property ID"
//	@Failure		401			{object}	utils.Error						"Unauthorized"
//	@Failure		403			{object}	utils.Error						"Property not yours"
//	@Failure		404			{object}	utils.Error						"Property not found"
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

	property, _ := c.MustGet("property").(db.PropertyModel)
	newProperty := database.UpdateProperty(property, req)
	if newProperty == nil {
		utils.SendError(c, http.StatusConflict, utils.PropertyAlreadyExists, nil)
		return
	}
	c.JSON(http.StatusOK, models.IdResponse{ID: newProperty.ID})
}

// GetPropertyPicture godoc
//
//	@Summary		Get property's picture
//	@Description	Get property's picture
//	@Tags			property
//	@Accept			json
//	@Produce		json
//	@Param			property_id	path		string					true	"Property ID"
//	@Param			lease_id	path		string					true	"Lease ID or `current`"
//	@Success		201			{object}	models.ImageResponse	"Image data"
//	@Success		204			"No picture associated"
//	@Failure		401			{object}	utils.Error	"Unauthorized"
//	@Failure		403			{object}	utils.Error	"Property not yours"
//	@Failure		404			{object}	utils.Error	"Property not found"
//	@Failure		500
//	@Security		Bearer
//	@Router			/owner/properties/{property_id}/picture/ [get]
//	@Router			/tenant/leases/{lease_id}/property/picture/ [get]
func GetPropertyPicture(c *gin.Context) {
	property, _ := c.MustGet("property").(db.PropertyModel)
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
//	@Tags			property
//	@Accept			json
//	@Produce		json
//	@Param			property_id	path		string				true	"Property ID"
//	@Param			picture		body		models.ImageRequest	true	"Picture data as a Base64 string"
//	@Success		201			{object}	models.IdResponse	"Updated property ID"
//	@Failure		400			{object}	utils.Error			"Missing fields or bad base64 string"
//	@Failure		401			{object}	utils.Error			"Unauthorized"
//	@Failure		403			{object}	utils.Error			"Property not yours"
//	@Failure		404			{object}	utils.Error			"Property not found"
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
		utils.SendError(c, http.StatusBadRequest, utils.BadBase64OrUnsupportedType, nil)
		return
	}
	newImage := database.CreateImage(*image)

	property, _ := c.MustGet("property").(db.PropertyModel)
	newProperty := database.UpdatePropertyPicture(property, newImage)
	if newProperty == nil {
		utils.SendError(c, http.StatusInternalServerError, utils.FailedLinkImage, nil)
		return
	}
	c.JSON(http.StatusOK, models.IdResponse{ID: newProperty.ID})
}

// InviteTenant godoc
//
//	@Summary		Invite tenant to owner's property
//	@Description	Invite tenant to owner's property
//	@Tags			property
//	@Accept			json
//	@Produce		json
//	@Param			property_id	path		string					true	"Property ID"
//	@Param			user		body		models.InviteRequest	true	"Invite params"
//	@Success		201			{object}	models.IdResponse		"Created invite ID"
//	@Failure		400			{object}	utils.Error				"Missing fields"
//	@Failure		403			{object}	utils.Error				"Property is not yours"
//	@Failure		404			{object}	utils.Error				"Property not found"
//	@Failure		409			{object}	utils.Error				"Invite already exists for this email"
//	@Failure		500
//	@Security		Bearer
//	@Router			/owner/properties/{property_id}/send-invite/ [post]
func InviteTenant(c *gin.Context) {
	var inviteReq models.InviteRequest
	err := c.ShouldBindBodyWithJSON(&inviteReq)
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, utils.MissingFields, err)
		return
	}
	inviteReq.TenantEmail = utils.SanitizeEmail(inviteReq.TenantEmail)

	if database.GetCurrentActiveLeaseByProperty(c.Param("property_id")) != nil {
		utils.SendError(c, http.StatusConflict, utils.PropertyNotAvailable, nil)
		return
	}

	user := database.GetUserByEmail(inviteReq.TenantEmail)
	if !checkInvitedTenant(c, user) {
		return
	}

	leaseInvite := database.CreateLeaseInvite(inviteReq.ToDbLeaseInvite(), c.Param("property_id"))
	if leaseInvite == nil {
		utils.SendError(c, http.StatusConflict, utils.InviteAlreadyExists, nil)
		return
	}

	res, err := brevo.SendEmailInvite(*leaseInvite, user != nil)
	if err != nil {
		log.Println(res, err.Error())
		utils.SendError(c, http.StatusInternalServerError, utils.FailedSendEmail, err)
		return
	}

	c.JSON(http.StatusCreated, models.IdResponse{ID: leaseInvite.ID})
}

func checkInvitedTenant(c *gin.Context, user *db.UserModel) bool {
	if user != nil {
		if user.Role != db.RoleTenant {
			utils.SendError(c, http.StatusConflict, utils.UserAlreadyExistsAsOwner, nil)
			return false
		}
		if database.GetCurrentActiveLeaseByTenant(user.ID) != nil {
			utils.SendError(c, http.StatusConflict, utils.TenantAlreadyHasLease, nil)
			return false
		}
	}
	return true
}

// CancelInvite godoc
//
//	@Summary		Cancel invite
//	@Description	Cancel pending lease invite
//	@Tags			property
//	@Accept			json
//	@Produce		json
//	@Param			property_id	path	string	true	"Property ID"
//	@Success		204			"Invite canceled"
//	@Failure		403			{object}	utils.Error	"Property is not yours"
//	@Failure		404			{object}	utils.Error	"No pending lease"
//	@Failure		500
//	@Security		Bearer
//	@Router			/owner/properties/{property_id}/cancel-invite/ [delete]
func CancelInvite(c *gin.Context) {
	database.DeleteCurrentLeaseInvite(c.Param("property_id"))
	c.Status(http.StatusNoContent)
}

// ArchiveProperty godoc
//
//	@Summary		Toggle archive property by ID
//	@Description	Toggle archive status of a property by its ID
//	@Tags			property
//	@Accept			json
//	@Produce		json
//	@Param			property_id	path		string					true	"Property ID"
//	@Param			archive		body		models.ArchiveRequest	true	"Archive status"
//	@Success		200			{object}	models.IdResponse		"Property archive status"
//	@Failure		400			{object}	utils.Error				"Mising fields"
//	@Failure		403			{object}	utils.Error				"Property not yours"
//	@Failure		404			{object}	utils.Error				"Property not found"
//	@Failure		500
//	@Security		Bearer
//	@Router			/owner/properties/{property_id}/archive/ [put]
func ArchiveProperty(c *gin.Context) {
	var req models.ArchiveRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendError(c, http.StatusBadRequest, utils.MissingFields, err)
		return
	}

	property, _ := c.MustGet("property").(db.PropertyModel)
	if slices.ContainsFunc(property.Leases(), func(x db.LeaseModel) bool { return x.Active }) && req.Archive {
		utils.SendError(c, http.StatusConflict, utils.CannotArchiveNonFreeProperty, nil)
		return
	}

	res := database.ArchiveProperty(property.ID, req.Archive)
	c.JSON(http.StatusOK, models.IdResponse{ID: res.ID})
}
