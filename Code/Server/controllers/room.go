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
//	@Success		201			{object}	models.IdResponse	"Created room ID"
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
	c.JSON(http.StatusCreated, models.IdResponse{ID: room.ID})
}

// GetRoomsByProperty godoc
//
//	@Summary		Get rooms by property ID
//	@Description	Get all rooms for a specific property, optionally filtered by archive status
//	@Tags			inventory
//	@Accept			json
//	@Produce		json
//	@Param			property_id	path		string				true	"Property ID"
//	@Param			archive		query		boolean				false	"Archive status (default: false)"
//	@Success		200			{array}		models.RoomResponse	"List of rooms"
//	@Failure		403			{object}	utils.Error			"Property not yours"
//	@Failure		404			{object}	utils.Error			"Property not found"
//	@Failure		500
//	@Security		Bearer
//	@Router			/owner/properties/{property_id}/rooms/ [get]
//	@Router			/tenant/leases/{lease_id}/property/rooms/ [get]
func GetRoomsByProperty(c *gin.Context) {
	property, _ := c.MustGet("property").(db.PropertyModel)
	archive := c.DefaultQuery("archive", "false") == utils.Strue
	rooms := database.GetRoomsByPropertyID(property.ID, archive)
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
//	@Router			/tenant/leases/{lease_id}/property/rooms/{room_id}/ [get]
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
//	@Success		200			{object}	models.IdResponse		"Updated room ID"
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
	c.JSON(http.StatusOK, models.IdResponse{ID: room.ID})
}

// DeleteRoom godoc
//
//	@Summary		Delete room by ID
//	@Description	Delete a room by its ID
//	@Tags			inventory
//	@Accept			json
//	@Produce		json
//	@Param			property_id	path		string				true	"Property ID"
//	@Param			room_id		path		string				true	"Room ID"
//	@Success		204			{object}	models.IdResponse	"Deleted room ID"
//	@Failure		403			{object}	utils.Error			"Property not yours"
//	@Failure		404			{object}	utils.Error			"Room not found"
//	@Failure		500
//	@Security		Bearer
//	@Router			/owner/properties/{property_id}/rooms/{room_id}/ [delete]
func DeleteRoom(c *gin.Context) {
	room, _ := c.MustGet("room").(db.RoomModel)
	database.DeleteRoom(room.ID)
	c.JSON(http.StatusNoContent, nil)
}
