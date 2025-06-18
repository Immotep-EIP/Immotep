package database_test

import (
	"errors"
	"testing"

	"github.com/steebchen/prisma-client-go/engine/protocol"
	"github.com/stretchr/testify/assert"
	"keyz/backend/prisma/db"
	"keyz/backend/services"
	"keyz/backend/services/database"
)

func BuildTestFurniture(id string) db.FurnitureModel {
	return db.FurnitureModel{
		InnerFurniture: db.InnerFurniture{
			ID:       id,
			Name:     "Test Furniture",
			Quantity: 1,
		},
	}
}

// #############################################################################

func TestCreateFurniture(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	furniture := BuildTestFurniture("1")
	m.Furniture.Expect(database.MockCreateFurniture(c, furniture)).Returns(furniture)

	newFurniture := database.CreateFurniture(furniture, "1")
	assert.NotNil(t, newFurniture)
	assert.Equal(t, furniture.ID, newFurniture.ID)
}

func TestCreateFurniture_AlreadyExists(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	furniture := BuildTestFurniture("1")
	m.Furniture.Expect(database.MockCreateFurniture(c, furniture)).Errors(&protocol.UserFacingError{
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
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	furniture := BuildTestFurniture("1")
	m.Furniture.Expect(database.MockCreateFurniture(c, furniture)).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.CreateFurniture(furniture, "1")
	})
}

// #############################################################################

func TestGetFurnitureByRoomID(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	furniture1 := BuildTestFurniture("1")
	furniture2 := BuildTestFurniture("2")
	m.Furniture.Expect(database.MockGetFurnituresByRoomID(c, false)).ReturnsMany([]db.FurnitureModel{furniture1, furniture2})

	furnitures := database.GetFurnituresByRoomID("1", false)
	assert.Len(t, furnitures, 2)
	assert.Equal(t, furniture1.ID, furnitures[0].ID)
	assert.Equal(t, furniture2.ID, furnitures[1].ID)
}

func TestGetFurnitureByRoomID_NoFurnitures(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.Furniture.Expect(database.MockGetFurnituresByRoomID(c, false)).ReturnsMany([]db.FurnitureModel{})

	furnitures := database.GetFurnituresByRoomID("1", false)
	assert.Empty(t, furnitures)
}

func TestGetFurnitureByRoomID_NoConnection(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.Furniture.Expect(database.MockGetFurnituresByRoomID(c, false)).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.GetFurnituresByRoomID("1", false)
	})
}

// #############################################################################

func TestGetFurnitureByID(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	furniture := BuildTestFurniture("1")
	m.Furniture.Expect(database.MockGetFurnitureByID(c)).Returns(furniture)

	foundFurniture := database.GetFurnitureByID("1")
	assert.NotNil(t, foundFurniture)
	assert.Equal(t, furniture.ID, foundFurniture.ID)
}

func TestGetFurnitureByID_NotFound(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.Furniture.Expect(database.MockGetFurnitureByID(c)).Errors(db.ErrNotFound)

	foundFurniture := database.GetFurnitureByID("1")
	assert.Nil(t, foundFurniture)
}

func TestGetFurnitureByID_NoConnection(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.Furniture.Expect(database.MockGetFurnitureByID(c)).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.GetFurnitureByID("1")
	})
}

// #############################################################################

func TestArchiveFurniture(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	furniture := BuildTestFurniture("1")
	furniture.Archived = true
	m.Furniture.Expect(database.MockArchiveFurniture(c)).Returns(furniture)

	archivedFurniture := database.ToggleArchiveFurniture("1", true)
	assert.NotNil(t, archivedFurniture)
	assert.Equal(t, furniture.ID, archivedFurniture.ID)
	assert.True(t, archivedFurniture.Archived)
}

func TestArchiveFurniture_NotFound(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.Furniture.Expect(database.MockArchiveFurniture(c)).Errors(db.ErrNotFound)

	archivedFurniture := database.ToggleArchiveFurniture("1", true)
	assert.Nil(t, archivedFurniture)
}

func TestArchiveFurniture_NoConnection(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.Furniture.Expect(database.MockArchiveFurniture(c)).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.ToggleArchiveFurniture("1", true)
	})
}

// #############################################################################

func TestDeleteFurniture(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.Furniture.Expect(database.MockDeleteFurniture(c)).Returns(BuildTestFurniture("1"))

	database.DeleteFurniture("1")
}

func TestDeleteFurniture_NotFound(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.Furniture.Expect(database.MockDeleteFurniture(c)).Errors(db.ErrNotFound)

	assert.Panics(t, func() {
		database.DeleteFurniture("1")
	})
}
