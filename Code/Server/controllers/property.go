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
//	@Tags			property
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
//	@Tags			property
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
