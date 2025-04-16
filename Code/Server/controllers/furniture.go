package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"immotep/backend/models"
	"immotep/backend/prisma/db"
	"immotep/backend/services/database"
	"immotep/backend/utils"
)

// CreateFurniture godoc
//
//	@Summary		Create a new furniture
//	@Description	Create a new furniture for a room
//	@Tags			inventory
//	@Accept			json
//	@Produce		json
//	@Param			property_id	path		string					true	"Property ID"
//	@Param			room_id		path		string					true	"Room ID"
//	@Param			furniture	body		models.FurnitureRequest	true	"Furniture data"
//	@Success		201			{object}	models.IdResponse		"Created furniture ID"
//	@Failure		400			{object}	utils.Error				"Missing fields"
//	@Failure		403			{object}	utils.Error				"Property not yours"
//	@Failure		409			{object}	utils.Error				"Furniture already exists"
//	@Failure		500
//	@Security		Bearer
//	@Router			/owner/properties/{property_id}/rooms/{room_id}/furnitures/ [post]
func CreateFurniture(c *gin.Context) {
	var req models.FurnitureRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendError(c, http.StatusBadRequest, utils.MissingFields, err)
		return
	}

	furniture := database.CreateFurniture(req.ToDbFurniture(), c.Param("room_id"))
	if furniture == nil {
		utils.SendError(c, http.StatusConflict, utils.FurnitureAlreadyExists, nil)
		return
	}
	c.JSON(http.StatusCreated, models.IdResponse{ID: furniture.ID})
}

// GetFurnituresByRoom godoc
//
//	@Summary		Get furnitures by room ID
//	@Description	Get all furnitures for a specific room, optionally filtered by archive status
//	@Tags			inventory
//	@Accept			json
//	@Produce		json
//	@Param			property_id	path		string						true	"Property ID"
//	@Param			room_id		path		string						true	"Room ID"
//	@Param			archive		query		bool						false	"Archive status filter"
//	@Success		200			{array}		models.FurnitureResponse	"List of furnitures"
//	@Failure		403			{object}	utils.Error					"Property not yours"
//	@Failure		404			{object}	utils.Error					"Room not found"
//	@Failure		500
//	@Security		Bearer
//	@Router			/owner/properties/{property_id}/rooms/{room_id}/furnitures/ [get]
func GetFurnituresByRoom(c *gin.Context) {
	archive := c.DefaultQuery("archive", "false") == utils.Strue
	furnitures := database.GetFurnituresByRoomID(c.Param("room_id"), archive)
	c.JSON(http.StatusOK, utils.Map(furnitures, models.DbFurnitureToResponse))
}

// GetFurniture godoc
//
//	@Summary		Get furniture by ID
//	@Description	Get furniture information by its ID
//	@Tags			inventory
//	@Accept			json
//	@Produce		json
//	@Param			property_id		path		string						true	"Property ID"
//	@Param			room_id			path		string						true	"Room ID"
//	@Param			furniture_id	path		string						true	"Furniture ID"
//	@Success		200				{object}	models.FurnitureResponse	"Furniture data"
//	@Failure		403				{object}	utils.Error					"Property not yours"
//	@Failure		404				{object}	utils.Error					"Furniture not found"
//	@Failure		500
//	@Security		Bearer
//	@Router			/owner/properties/{property_id}/rooms/{room_id}/furnitures/{furniture_id}/ [get]
func GetFurniture(c *gin.Context) {
	furniture, _ := c.MustGet("furniture").(db.FurnitureModel)
	c.JSON(http.StatusOK, models.DbFurnitureToResponse(furniture))
}

// ArchiveFurniture godoc
//
//	@Summary		Toggle archive furniture by ID
//	@Description	Toggle archive status of a furniture by its ID
//	@Tags			inventory
//	@Accept			json
//	@Produce		json
//	@Param			property_id		path		string					true	"Property ID"
//	@Param			room_id			path		string					true	"Room ID"
//	@Param			furniture_id	path		string					true	"Furniture ID"
//	@Param			archive			body		models.ArchiveRequest	true	"Archive status"
//	@Success		200				{object}	models.IdResponse		"Updated furniture ID"
//	@Failure		400				{object}	utils.Error				"Mising fields"
//	@Failure		403				{object}	utils.Error				"Property not yours"
//	@Failure		404				{object}	utils.Error				"Furniture not found"
//	@Failure		500
//	@Security		Bearer
//	@Router			/owner/properties/{property_id}/rooms/{room_id}/furnitures/{furniture_id}/archive [put]
func ArchiveFurniture(c *gin.Context) {
	var req models.ArchiveRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendError(c, http.StatusBadRequest, utils.MissingFields, err)
		return
	}

	furniture := database.ToggleArchiveFurniture(c.Param("furniture_id"), req.Archive)
	c.JSON(http.StatusOK, models.IdResponse{ID: furniture.ID})
}
