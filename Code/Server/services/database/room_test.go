package database_test

import (
	"errors"
	"testing"

	"github.com/steebchen/prisma-client-go/engine/protocol"
	"github.com/stretchr/testify/assert"
	"immotep/backend/prisma/db"
	"immotep/backend/services"
	"immotep/backend/services/database"
)

func BuildTestRoom(id string) db.RoomModel {
	return db.RoomModel{
		InnerRoom: db.InnerRoom{
			ID:   id,
			Name: "Test Room",
			Type: db.RoomTypeOther,
		},
	}
}

// #############################################################################

func TestCreateRoom(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	room := BuildTestRoom("1")
	m.Room.Expect(database.MockCreateRoom(c, room)).Returns(room)

	newRoom := database.CreateRoom(room, "1")
	assert.NotNil(t, newRoom)
	assert.Equal(t, room.ID, newRoom.ID)
}

func TestCreateRoom_AlreadyExists(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	room := BuildTestRoom("1")
	m.Room.Expect(database.MockCreateRoom(c, room)).Errors(&protocol.UserFacingError{
		IsPanic:   false,
		ErrorCode: "P2002", // https://www.prisma.io/docs/orm/reference/error-reference
		Meta: protocol.Meta{
			Target: []any{"name"},
		},
		Message: "Unique constraint failed",
	})

	newRoom := database.CreateRoom(room, "1")
	assert.Nil(t, newRoom)
}

func TestCreateRoom_NoConnection(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	room := BuildTestRoom("1")
	m.Room.Expect(database.MockCreateRoom(c, room)).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.CreateRoom(room, "1")
	})
}

// #############################################################################

func TestGetRoomByPropertyID(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	room1 := BuildTestRoom("1")
	room2 := BuildTestRoom("2")
	m.Room.Expect(database.MockGetRoomsByPropertyID(c, false)).ReturnsMany([]db.RoomModel{room1, room2})

	rooms := database.GetRoomsByPropertyID("1", false)
	assert.Len(t, rooms, 2)
	assert.Equal(t, room1.ID, rooms[0].ID)
	assert.Equal(t, room2.ID, rooms[1].ID)
}

func TestGetRoomByPropertyID_NoRooms(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.Room.Expect(database.MockGetRoomsByPropertyID(c, false)).ReturnsMany([]db.RoomModel{})

	rooms := database.GetRoomsByPropertyID("1", false)
	assert.Empty(t, rooms)
}

func TestGetRoomByPropertyID_NoConnection(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.Room.Expect(database.MockGetRoomsByPropertyID(c, false)).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.GetRoomsByPropertyID("1", false)
	})
}

// #############################################################################

func TestGetRoomByID(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	room := BuildTestRoom("1")
	m.Room.Expect(database.MockGetRoomByID(c)).Returns(room)

	foundRoom := database.GetRoomByID("1")
	assert.NotNil(t, foundRoom)
	assert.Equal(t, room.ID, foundRoom.ID)
}

func TestGetRoomByID_NotFound(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.Room.Expect(database.MockGetRoomByID(c)).Errors(db.ErrNotFound)

	foundRoom := database.GetRoomByID("1")
	assert.Nil(t, foundRoom)
}

func TestGetRoomByID_NoConnection(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.Room.Expect(database.MockGetRoomByID(c)).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.GetRoomByID("1")
	})
}

// #############################################################################

func TestArchiveRoom(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	room := BuildTestRoom("1")
	room.Archived = true
	m.Room.Expect(database.MockArchiveRoom(c)).Returns(room)

	archivedRoom := database.ToggleArchiveRoom("1", true)
	assert.NotNil(t, archivedRoom)
	assert.Equal(t, room.ID, archivedRoom.ID)
	assert.True(t, archivedRoom.Archived)
}

func TestArchiveRoom_NotFound(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.Room.Expect(database.MockArchiveRoom(c)).Errors(db.ErrNotFound)

	archivedRoom := database.ToggleArchiveRoom("1", true)
	assert.Nil(t, archivedRoom)
}

func TestArchiveRoom_NoConnection(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.Room.Expect(database.MockArchiveRoom(c)).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.ToggleArchiveRoom("1", true)
	})
}
