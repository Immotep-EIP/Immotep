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

func TestCreateFurniture(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	furniture := db.FurnitureModel{
		InnerFurniture: db.InnerFurniture{
			ID:       "1",
			Name:     "Test Furniture",
			Quantity: 1,
		},
	}

	mock.Furniture.Expect(
		client.Client.Furniture.CreateOne(
			db.Furniture.Name.Set(furniture.Name),
			db.Furniture.Room.Link(db.Room.ID.Equals("1")),
			db.Furniture.Quantity.Set(furniture.Quantity),
		).With(
			db.Furniture.Room.Fetch(),
		),
	).Returns(furniture)

	newFurniture := database.CreateFurniture(furniture, "1")
	assert.NotNil(t, newFurniture)
	assert.Equal(t, furniture.ID, newFurniture.ID)
}

func TestCreateFurniture_AlreadyExists(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	furniture := db.FurnitureModel{
		InnerFurniture: db.InnerFurniture{
			ID:       "1",
			Name:     "Test Furniture",
			Quantity: 1,
		},
	}

	mock.Furniture.Expect(
		client.Client.Furniture.CreateOne(
			db.Furniture.Name.Set(furniture.Name),
			db.Furniture.Room.Link(db.Room.ID.Equals("1")),
			db.Furniture.Quantity.Set(furniture.Quantity),
		).With(
			db.Furniture.Room.Fetch(),
		),
	).Errors(&protocol.UserFacingError{
		IsPanic:   false,
		ErrorCode: "P2002", // https://www.prisma.io/docs/orm/reference/error-reference
		Meta: protocol.Meta{
			Target: []any{"name"},
		},
		Message: "Unique constraint failed",
	})

	newFurniture := database.CreateFurniture(furniture, "1")
	assert.Nil(t, newFurniture)
}

func TestCreateFurniture_NoConnection(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	furniture := db.FurnitureModel{
		InnerFurniture: db.InnerFurniture{
			ID:       "1",
			Name:     "Test Furniture",
			Quantity: 1,
		},
	}

	mock.Furniture.Expect(
		client.Client.Furniture.CreateOne(
			db.Furniture.Name.Set(furniture.Name),
			db.Furniture.Room.Link(db.Room.ID.Equals("1")),
			db.Furniture.Quantity.Set(furniture.Quantity),
		).With(
			db.Furniture.Room.Fetch(),
		),
	).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.CreateFurniture(furniture, "1")
	})
}

func TestGetFurnitureByRoomID(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	furniture1 := db.FurnitureModel{
		InnerFurniture: db.InnerFurniture{
			ID:       "1",
			Name:     "Test Furniture 1",
			Quantity: 1,
		},
	}
	furniture2 := db.FurnitureModel{
		InnerFurniture: db.InnerFurniture{
			ID:       "2",
			Name:     "Test Furniture 2",
			Quantity: 2,
		},
	}

	mock.Furniture.Expect(
		client.Client.Furniture.FindMany(
			db.Furniture.RoomID.Equals("1"),
			db.Furniture.Archived.Equals(false),
		).With(
			db.Furniture.Room.Fetch(),
		),
	).ReturnsMany([]db.FurnitureModel{furniture1, furniture2})

	furnitures := database.GetFurnitureByRoomID("1", false)
	assert.Len(t, furnitures, 2)
	assert.Equal(t, furniture1.ID, furnitures[0].ID)
	assert.Equal(t, furniture2.ID, furnitures[1].ID)
}

func TestGetFurnitureByRoomID_NoFurnitures(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	mock.Furniture.Expect(
		client.Client.Furniture.FindMany(
			db.Furniture.RoomID.Equals("1"),
			db.Furniture.Archived.Equals(false),
		).With(
			db.Furniture.Room.Fetch(),
		),
	).ReturnsMany([]db.FurnitureModel{})

	furnitures := database.GetFurnitureByRoomID("1", false)
	assert.Empty(t, furnitures)
}

func TestGetFurnitureByRoomID_NoConnection(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	mock.Furniture.Expect(
		client.Client.Furniture.FindMany(
			db.Furniture.RoomID.Equals("1"),
			db.Furniture.Archived.Equals(false),
		).With(
			db.Furniture.Room.Fetch(),
		),
	).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.GetFurnitureByRoomID("1", false)
	})
}

func TestGetFurnitureByID(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	furniture := db.FurnitureModel{
		InnerFurniture: db.InnerFurniture{
			ID:       "1",
			Name:     "Test Furniture",
			Quantity: 1,
		},
	}

	mock.Furniture.Expect(
		client.Client.Furniture.FindUnique(
			db.Furniture.ID.Equals("1"),
		).With(
			db.Furniture.Room.Fetch(),
		),
	).Returns(furniture)

	foundFurniture := database.GetFurnitureByID("1")
	assert.NotNil(t, foundFurniture)
	assert.Equal(t, furniture.ID, foundFurniture.ID)
}

func TestGetFurnitureByID_NotFound(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	mock.Furniture.Expect(
		client.Client.Furniture.FindUnique(
			db.Furniture.ID.Equals("1"),
		).With(
			db.Furniture.Room.Fetch(),
		),
	).Errors(db.ErrNotFound)

	foundFurniture := database.GetFurnitureByID("1")
	assert.Nil(t, foundFurniture)
}

func TestGetFurnitureByID_NoConnection(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	mock.Furniture.Expect(
		client.Client.Furniture.FindUnique(
			db.Furniture.ID.Equals("1"),
		).With(
			db.Furniture.Room.Fetch(),
		),
	).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.GetFurnitureByID("1")
	})
}

func TestArchiveFurniture(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	furniture := db.FurnitureModel{
		InnerFurniture: db.InnerFurniture{
			ID:       "1",
			Name:     "Test Furniture",
			Quantity: 1,
			Archived: true,
		},
	}

	mock.Furniture.Expect(
		client.Client.Furniture.FindUnique(
			db.Furniture.ID.Equals(furniture.ID),
		).With(
			db.Furniture.Room.Fetch(),
		).Update(
			db.Furniture.Archived.Set(true),
		),
	).Returns(furniture)

	archivedFurniture := database.ToggleArchiveFurniture("1", true)
	assert.NotNil(t, archivedFurniture)
	assert.Equal(t, furniture.ID, archivedFurniture.ID)
	assert.True(t, archivedFurniture.Archived)
}

func TestArchiveFurniture_NotFound(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	mock.Furniture.Expect(
		client.Client.Furniture.FindUnique(
			db.Furniture.ID.Equals("1"),
		).With(
			db.Furniture.Room.Fetch(),
		).Update(
			db.Furniture.Archived.Set(true),
		),
	).Errors(db.ErrNotFound)

	archivedFurniture := database.ToggleArchiveFurniture("1", true)
	assert.Nil(t, archivedFurniture)
}

func TestArchiveFurniture_NoConnection(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	mock.Furniture.Expect(
		client.Client.Furniture.FindUnique(
			db.Furniture.ID.Equals("1"),
		).With(
			db.Furniture.Room.Fetch(),
		).Update(
			db.Furniture.Archived.Set(true),
		),
	).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.ToggleArchiveFurniture("1", true)
	})
}
