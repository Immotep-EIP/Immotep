package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"immotep/backend/models"
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
//	@Param			id	path		string					true	"Property ID"
//	@Success		200	{object}	models.PropertyResponse	"Property data"
//	@Failure		401	{object}	utils.Error				"Unauthorized"
//	@Failure		403	{object}	utils.Error				"Property not yours"
//	@Failure		404	{object}	utils.Error				"Property not found"
//	@Failure		500
//	@Security		Bearer
//	@Router			/owner/properties/{id} [get]
func GetPropertyById(c *gin.Context) {
	claims := utils.GetClaims(c)
	property := propertyservice.GetByID(c.Param("id"))
	if property == nil {
		utils.SendError(c, http.StatusNotFound, utils.PropertyNotFound, nil)
		return
	}
	if property.OwnerID != claims["id"] {
		utils.SendError(c, http.StatusForbidden, utils.PropertyNotYours, nil)
		return
	}
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
