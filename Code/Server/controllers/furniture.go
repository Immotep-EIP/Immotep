package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"immotep/backend/models"
	"immotep/backend/services/database"
	"immotep/backend/utils"
)

// CreateFurniture godoc
//
//	@Summary		Create a new furniture
//	@Description	Create a new furniture for a room
//	@Tags			owner
//	@Accept			json
//	@Produce		json
//	@Param			property_id	path		string						true	"Property ID"
//	@Param			room_id		path		string						true	"Room ID"
//	@Param			furniture	body		models.FurnitureRequest		true	"Furniture data"
//	@Success		201			{object}	models.FurnitureResponse	"Created furniture data"
//	@Failure		400			{object}	utils.Error					"Missing fields"
//	@Failure		403			{object}	utils.Error					"Property not yours"
//	@Failure		409			{object}	utils.Error					"Furniture already exists"
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
	c.JSON(http.StatusCreated, models.DbFurnitureToResponse(*furniture))
}

// GetFurnituresByRoom godoc
//
//	@Summary		Get furnitures by room ID
//	@Description	Get all furnitures for a specific room
//	@Tags			owner
//	@Accept			json
//	@Produce		json
//	@Param			property_id	path		string						true	"Property ID"
//	@Param			room_id		path		string						true	"Room ID"
//	@Success		200			{array}		models.FurnitureResponse	"List of furnitures"
//	@Failure		403			{object}	utils.Error					"Property not yours"
//	@Failure		404			{object}	utils.Error					"Room not found"
//	@Failure		500
//	@Security		Bearer
//	@Router			/owner/properties/{property_id}/rooms/{room_id}/furnitures/ [get]
func GetFurnituresByRoom(c *gin.Context) {
	furnitures := database.GetFurnitureByRoomID(c.Param("room_id"))
	c.JSON(http.StatusOK, utils.Map(furnitures, models.DbFurnitureToResponse))
}

// GetFurnitureByID godoc
//
//	@Summary		Get furniture by ID
//	@Description	Get furniture information by its ID
//	@Tags			owner
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
func GetFurnitureByID(c *gin.Context) {
	furniture := database.GetFurnitureByID(c.Param("furniture_id"))
	c.JSON(http.StatusOK, models.DbFurnitureToResponse(*furniture))
}

// ArchiveFurniture godoc
//
//	@Summary		Archive furniture by ID
//	@Description	Archive a furniture by its ID
//	@Tags			owner
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string						true	"Furniture ID"
//	@Success		200	{object}	models.FurnitureResponse	"Archived furniture data"
//	@Failure		404	{object}	utils.Error					"Furniture not found"
//	@Failure		500
//	@Security		Bearer
//	@Router			/owner/properties/{property_id}/rooms/{room_id}/furnitures/{furniture_id}/ [delete]
func ArchiveFurniture(c *gin.Context) {
	furniture := database.ArchiveFurniture(c.Param("furniture_id"))
	c.JSON(http.StatusOK, models.DbFurnitureToResponse(*furniture))
}
