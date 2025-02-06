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

func TestCreateRoom(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	room := db.RoomModel{
		InnerRoom: db.InnerRoom{
			ID:   "1",
			Name: "Test Room",
		},
	}

	mock.Room.Expect(
		client.Client.Room.CreateOne(
			db.Room.Name.Set(room.Name),
			db.Room.Property.Link(db.Property.ID.Equals("1")),
		),
	).Returns(room)

	newRoom := database.CreateRoom(room, "1")
	assert.NotNil(t, newRoom)
	assert.Equal(t, room.ID, newRoom.ID)
}

func TestCreateRoom_AlreadyExists(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	room := db.RoomModel{
		InnerRoom: db.InnerRoom{
			ID:   "1",
			Name: "Test Room",
		},
	}

	mock.Room.Expect(
		client.Client.Room.CreateOne(
			db.Room.Name.Set(room.Name),
			db.Room.Property.Link(db.Property.ID.Equals("1")),
		),
	).Errors(&protocol.UserFacingError{
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
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	room := db.RoomModel{
		InnerRoom: db.InnerRoom{
			ID:   "1",
			Name: "Test Room",
		},
	}

	mock.Room.Expect(
		client.Client.Room.CreateOne(
			db.Room.Name.Set(room.Name),
			db.Room.Property.Link(db.Property.ID.Equals("1")),
		),
	).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.CreateRoom(room, "1")
	})
}

func TestGetRoomByPropertyID(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	room1 := db.RoomModel{
		InnerRoom: db.InnerRoom{
			ID:   "1",
			Name: "Test Room 1",
		},
	}
	room2 := db.RoomModel{
		InnerRoom: db.InnerRoom{
			ID:   "2",
			Name: "Test Room 2",
		},
	}

	mock.Room.Expect(
		client.Client.Room.FindMany(
			db.Room.PropertyID.Equals("1"),
		),
	).ReturnsMany([]db.RoomModel{room1, room2})

	rooms := database.GetRoomByPropertyID("1")
	assert.Len(t, rooms, 2)
	assert.Equal(t, room1.ID, rooms[0].ID)
	assert.Equal(t, room2.ID, rooms[1].ID)
}

func TestGetRoomByPropertyID_NoRooms(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	mock.Room.Expect(
		client.Client.Room.FindMany(
			db.Room.PropertyID.Equals("1"),
		),
	).ReturnsMany([]db.RoomModel{})

	rooms := database.GetRoomByPropertyID("1")
	assert.Empty(t, rooms)
}

func TestGetRoomByPropertyID_NoConnection(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	mock.Room.Expect(
		client.Client.Room.FindMany(
			db.Room.PropertyID.Equals("1"),
		),
	).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.GetRoomByPropertyID("1")
	})
}

func TestGetRoomByID(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	room := db.RoomModel{
		InnerRoom: db.InnerRoom{
			ID:   "1",
			Name: "Test Room",
		},
	}

	mock.Room.Expect(
		client.Client.Room.FindUnique(
			db.Room.ID.Equals("1"),
		),
	).Returns(room)

	foundRoom := database.GetRoomByID("1")
	assert.NotNil(t, foundRoom)
	assert.Equal(t, room.ID, foundRoom.ID)
}

func TestGetRoomByID_NotFound(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	mock.Room.Expect(
		client.Client.Room.FindUnique(
			db.Room.ID.Equals("1"),
		),
	).Errors(db.ErrNotFound)

	foundRoom := database.GetRoomByID("1")
	assert.Nil(t, foundRoom)
}

func TestGetRoomByID_NoConnection(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	mock.Room.Expect(
		client.Client.Room.FindUnique(
			db.Room.ID.Equals("1"),
		),
	).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.GetRoomByID("1")
	})
}
