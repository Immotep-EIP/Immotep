package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"immotep/backend/models"
	"immotep/backend/prisma/db"
	"immotep/backend/services/database"
	"immotep/backend/utils"
)

// CreateRoom godoc
//
//	@Summary		Create a new room
//	@Description	Create a new room for a property
//	@Tags			inventory
//	@Accept			json
//	@Produce		json
//	@Param			property_id	path		string				true	"Property ID"
//	@Param			room		body		models.RoomRequest	true	"Room data"
//	@Success		201			{object}	models.RoomResponse	"Created room data"
//	@Failure		400			{object}	utils.Error			"Missing fields"
//	@Failure		403			{object}	utils.Error			"Property not yours"
//	@Failure		404			{object}	utils.Error			"Property not found"
//	@Failure		500
//	@Security		Bearer
//	@Router			/owner/properties/{property_id}/rooms/ [post]
func CreateRoom(c *gin.Context) {
	var req models.RoomRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendError(c, http.StatusBadRequest, utils.MissingFields, err)
		return
	}

	room := database.CreateRoom(req.ToDbRoom(), c.Param("property_id"))
	if room == nil {
		utils.SendError(c, http.StatusConflict, utils.RoomAlreadyExists, nil)
		return
	}
	c.JSON(http.StatusCreated, models.DbRoomToResponse(*room))
}

// GetAllRoomsByProperty godoc
//
//	@Summary		Get rooms by property ID
//	@Description	Get all rooms for a specific property
//	@Tags			inventory
//	@Accept			json
//	@Produce		json
//	@Param			property_id	path		string				true	"Property ID"
//	@Success		200			{array}		models.RoomResponse	"List of rooms"
//	@Failure		403			{object}	utils.Error			"Property not yours"
//	@Failure		404			{object}	utils.Error			"Property not found"
//	@Failure		500
//	@Security		Bearer
//	@Router			/owner/properties/{property_id}/rooms/ [get]
func GetAllRoomsByProperty(c *gin.Context) {
	rooms := database.GetRoomsByPropertyID(c.Param("property_id"), false)
	c.JSON(http.StatusOK, utils.Map(rooms, models.DbRoomToResponse))
}

// GetArchivedRoomsByProperty godoc
//
//	@Summary		Get archived rooms by property ID
//	@Description	Get all archived rooms for a specific property
//	@Tags			inventory
//	@Accept			json
//	@Produce		json
//	@Param			property_id	path		string				true	"Property ID"
//	@Success		200			{array}		models.RoomResponse	"List of archived rooms"
//	@Failure		403			{object}	utils.Error			"Property not yours"
//	@Failure		404			{object}	utils.Error			"Property not found"
//	@Failure		500
//	@Security		Bearer
//	@Router			/owner/properties/{property_id}/rooms/archived/ [get]
func GetArchivedRoomsByProperty(c *gin.Context) {
	rooms := database.GetRoomsByPropertyID(c.Param("property_id"), true)
	c.JSON(http.StatusOK, utils.Map(rooms, models.DbRoomToResponse))
}

// GetRoom godoc
//
//	@Summary		Get room by ID
//	@Description	Get room information by its ID
//	@Tags			inventory
//	@Accept			json
//	@Produce		json
//	@Param			property_id	path		string				true	"Property ID"
//	@Param			room_id		path		string				true	"Room ID"
//	@Success		200			{object}	models.RoomResponse	"Room data"
//	@Failure		403			{object}	utils.Error			"Property not yours"
//	@Failure		404			{object}	utils.Error			"Room not found"
//	@Failure		500
//	@Security		Bearer
//	@Router			/owner/properties/{property_id}/rooms/{room_id}/ [get]
func GetRoom(c *gin.Context) {
	room, _ := c.MustGet("room").(db.RoomModel)
	c.JSON(http.StatusOK, models.DbRoomToResponse(room))
}

// ArchiveRoom godoc
//
//	@Summary		Toggle archive room by ID
//	@Description	Toggle archive status of a room by its ID
//	@Tags			inventory
//	@Accept			json
//	@Produce		json
//	@Param			property_id	path		string					true	"Property ID"
//	@Param			room_id		path		string					true	"Room ID"
//	@Param			archive		body		models.ArchiveRequest	true	"Archive status"
//	@Success		200			{object}	models.PropertyResponse	"Toggled archive room data"
//	@Failure		400			{object}	utils.Error				"Mising fields"
//	@Failure		403			{object}	utils.Error				"Property not yours"
//	@Failure		404			{object}	utils.Error				"Room not found"
//	@Failure		500
//	@Security		Bearer
//	@Router			/owner/properties/{property_id}/rooms/{room_id}/archive [put]
func ArchiveRoom(c *gin.Context) {
	var req models.ArchiveRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendError(c, http.StatusBadRequest, utils.MissingFields, err)
		return
	}

	room := database.ToggleArchiveRoom(c.Param("room_id"), req.Archive)
	c.JSON(http.StatusOK, models.DbRoomToResponse(*room))
}
