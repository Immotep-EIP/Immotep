package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"immotep/backend/models"
	roomservice "immotep/backend/services/room"
	"immotep/backend/utils"
)

// CreateRoom godoc
//
//	@Summary		Create a new room
//	@Description	Create a new room for a property
//	@Tags			owner
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
//	@Router			/owner/properties/{property_id}/rooms [post]
func CreateRoom(c *gin.Context) {
	var req models.RoomRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendError(c, http.StatusBadRequest, utils.MissingFields, err)
		return
	}

	room := roomservice.Create(req.ToDbRoom(), c.Param("property_id"))
	if room == nil {
		utils.SendError(c, http.StatusNotFound, utils.RoomAlreadyExists, nil)
		return
	}
	c.JSON(http.StatusCreated, models.DbRoomToResponse(*room))
}

// GetRoomsByProperty godoc
//
//	@Summary		Get rooms by property ID
//	@Description	Get all rooms for a specific property
//	@Tags			owner
//	@Accept			json
//	@Produce		json
//	@Param			property_id	path		string				true	"Property ID"
//	@Success		200			{array}		models.RoomResponse	"List of rooms"
//	@Failure		403			{object}	utils.Error			"Property not yours"
//	@Failure		404			{object}	utils.Error			"Property not found"
//	@Failure		500
//	@Security		Bearer
//	@Router			/owner/properties/{property_id}/rooms [get]
func GetRoomsByProperty(c *gin.Context) {
	rooms := roomservice.GetByPropertyID(c.Param("property_id"))
	c.JSON(http.StatusOK, utils.Map(rooms, models.DbRoomToResponse))
}

// GetRoomByID godoc
//
//	@Summary		Get room by ID
//	@Description	Get room information by its ID
//	@Tags			owner
//	@Accept			json
//	@Produce		json
//	@Param			property_id	path		string				true	"Property ID"
//	@Param			room_id		path		string				true	"Room ID"
//	@Success		200			{object}	models.RoomResponse	"Room data"
//	@Failure		403			{object}	utils.Error			"Property not yours"
//	@Failure		404			{object}	utils.Error			"Room not found"
//	@Failure		500
//	@Security		Bearer
//	@Router			/owner/properties/{property_id}/rooms/{room_id} [get]
func GetRoomByID(c *gin.Context) {
	room := roomservice.GetByID(c.Param("room_id"))
	if room == nil || room.PropertyID != c.Param("property_id") {
		utils.SendError(c, http.StatusNotFound, utils.RoomNotFound, nil)
		return
	}
	c.JSON(http.StatusOK, models.DbRoomToResponse(*room))
}

// DeleteRoom godoc
//
//	@Summary		Delete room by ID
//	@Description	Delete a room by its ID
//	@Tags			owner
//	@Accept			json
//	@Produce		json
//	@Param			property_id	path	string	true	"Property ID"
//	@Param			room_id		path	string	true	"Room ID"
//	@Success		204
//	@Failure		403	{object}	utils.Error	"Property not yours"
//	@Failure		404	{object}	utils.Error	"Room not found"
//	@Failure		500
//	@Security		Bearer
//	@Router			/owner/properties/{property_id}/rooms/{room_id} [delete]
func DeleteRoom(c *gin.Context) {
	room := roomservice.GetByID(c.Param("room_id"))
	if room == nil || room.PropertyID != c.Param("property_id") {
		utils.SendError(c, http.StatusNotFound, utils.RoomNotFound, nil)
		return
	}
	ok := roomservice.Delete(c.Param("room_id"))
	if !ok {
		utils.SendError(c, http.StatusNotFound, utils.RoomNotFound, nil)
		return
	}
	c.Status(http.StatusNoContent)
}
