package models_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"immotep/backend/models"
	"immotep/backend/prisma/db"
)

func TestRoomRequest(t *testing.T) {
	req := models.RoomRequest{
		Name: "Test Room",
	}

	t.Run("ToRoom", func(t *testing.T) {
		room := req.ToDbRoom()

		assert.Equal(t, req.Name, room.Name)
	})
}

func TestRoomResponse(t *testing.T) {
	room := db.RoomModel{
		InnerRoom: db.InnerRoom{
			ID:         "1",
			PropertyID: "1",
			Name:       "Test Room",
		},
	}

	t.Run("FromRoom", func(t *testing.T) {
		roomResponse := models.RoomResponse{}
		roomResponse.FromDbRoom(room)

		assert.Equal(t, room.ID, roomResponse.ID)
		assert.Equal(t, room.PropertyID, roomResponse.PropertyID)
		assert.Equal(t, room.Name, roomResponse.Name)
	})

	t.Run("RoomToResponse", func(t *testing.T) {
		roomResponse := models.DbRoomToResponse(room)

		assert.Equal(t, room.ID, roomResponse.ID)
		assert.Equal(t, room.PropertyID, roomResponse.PropertyID)
		assert.Equal(t, room.Name, roomResponse.Name)
	})
}
