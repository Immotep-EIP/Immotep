package furnitureservice_test

import (
	"errors"
	"testing"

	"github.com/steebchen/prisma-client-go/engine/protocol"
	"github.com/stretchr/testify/assert"
	"immotep/backend/database"
	"immotep/backend/prisma/db"
	furnitureservice "immotep/backend/services/furniture"
)

func TestCreateFurniture(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
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

	newFurniture := furnitureservice.Create(furniture, "1")
	assert.NotNil(t, newFurniture)
	assert.Equal(t, furniture.ID, newFurniture.ID)
}

func TestCreateFurniture_AlreadyExists(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
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

	newFurniture := furnitureservice.Create(furniture, "1")
	assert.Nil(t, newFurniture)
}

func TestCreateFurniture_NoConnection(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
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
		furnitureservice.Create(furniture, "1")
	})
}

func TestGetByRoomID(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
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
		).With(
			db.Furniture.Room.Fetch(),
		),
	).ReturnsMany([]db.FurnitureModel{furniture1, furniture2})

	furnitures := furnitureservice.GetByRoomID("1")
	assert.Len(t, furnitures, 2)
	assert.Equal(t, furniture1.ID, furnitures[0].ID)
	assert.Equal(t, furniture2.ID, furnitures[1].ID)
}

func TestGetByRoomID_NoFurnitures(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	mock.Furniture.Expect(
		client.Client.Furniture.FindMany(
			db.Furniture.RoomID.Equals("1"),
		).With(
			db.Furniture.Room.Fetch(),
		),
	).ReturnsMany([]db.FurnitureModel{})

	furnitures := furnitureservice.GetByRoomID("1")
	assert.Empty(t, furnitures)
}

func TestGetByRoomID_NoConnection(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	mock.Furniture.Expect(
		client.Client.Furniture.FindMany(
			db.Furniture.RoomID.Equals("1"),
		).With(
			db.Furniture.Room.Fetch(),
		),
	).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		furnitureservice.GetByRoomID("1")
	})
}

func TestGetByID(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
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

	foundFurniture := furnitureservice.GetByID("1")
	assert.NotNil(t, foundFurniture)
	assert.Equal(t, furniture.ID, foundFurniture.ID)
}

func TestGetByID_NotFound(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	mock.Furniture.Expect(
		client.Client.Furniture.FindUnique(
			db.Furniture.ID.Equals("1"),
		).With(
			db.Furniture.Room.Fetch(),
		),
	).Errors(db.ErrNotFound)

	foundFurniture := furnitureservice.GetByID("1")
	assert.Nil(t, foundFurniture)
}

func TestGetByID_NoConnection(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	mock.Furniture.Expect(
		client.Client.Furniture.FindUnique(
			db.Furniture.ID.Equals("1"),
		).With(
			db.Furniture.Room.Fetch(),
		),
	).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		furnitureservice.GetByID("1")
	})
}
