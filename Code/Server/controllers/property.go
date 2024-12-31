package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"immotep/backend/models"
	imageservice "immotep/backend/services/image"
	propertyservice "immotep/backend/services/property"
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
//	@Router			/owner/properties [get]
func GetAllProperties(c *gin.Context) {
	claims := utils.GetClaims(c)
	allProperties := propertyservice.GetAllByOwnerId(claims["id"])
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
//	@Router			/owner/properties/{property_id} [get]
func GetPropertyById(c *gin.Context) {
	property := propertyservice.GetByID(c.Param("property_id"))
	c.JSON(http.StatusOK, models.DbPropertyToResponse(*property))
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
//	@Router			/owner/properties [post]
func CreateProperty(c *gin.Context) {
	claims := utils.GetClaims(c)

	var req models.PropertyRequest
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, utils.MissingFields, err)
		return
	}

	property := propertyservice.Create(req.ToDbProperty(), claims["id"])
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
//	@Router			/owner/properties/{property_id}/picture [get]
func GetPropertyPicture(c *gin.Context) {
	property := propertyservice.GetByID(c.Param("property_id"))
	pictureId, ok := property.PictureID()
	if !ok {
		c.Status(http.StatusNoContent)
		return
	}
	image := imageservice.GetByID(pictureId)
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
//	@Router			/owner/properties/{property_id}/picture [put]
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
	newImage := imageservice.Create(*image)

	property := propertyservice.GetByID(c.Param("property_id"))
	newProperty := propertyservice.UpdatePicture(*property, newImage)
	if newProperty == nil {
		utils.SendError(c, http.StatusInternalServerError, utils.FailedLinkImage, nil)
		return
	}
	c.JSON(http.StatusOK, models.DbPropertyToResponse(*newProperty))
}
