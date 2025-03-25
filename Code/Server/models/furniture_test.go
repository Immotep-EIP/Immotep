package models_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"immotep/backend/models"
	"immotep/backend/prisma/db"
	"immotep/backend/utils"
)

func TestFurnitureRequest(t *testing.T) {
	req := models.FurnitureRequest{
		Name:     "Test Furniture",
		Quantity: utils.Ptr(5),
	}

	t.Run("ToFurniture", func(t *testing.T) {
		furniture := req.ToDbFurniture()

		assert.Equal(t, req.Name, furniture.Name)
		assert.Equal(t, *req.Quantity, furniture.Quantity)
	})

	t.Run("ToFurnitureDefaultQuantity", func(t *testing.T) {
		reqWithoutQuantity := models.FurnitureRequest{
			Name: "Test Furniture",
		}
		furniture := reqWithoutQuantity.ToDbFurniture()

		assert.Equal(t, reqWithoutQuantity.Name, furniture.Name)
		assert.Equal(t, 1, furniture.Quantity)
	})
}

func TestFurnitureResponse(t *testing.T) {
	furniture := db.FurnitureModel{
		InnerFurniture: db.InnerFurniture{
			ID:       "1",
			RoomID:   "1",
			Name:     "Test Furniture",
			Quantity: 5,
		},
		RelationsFurniture: db.RelationsFurniture{
			Room: &db.RoomModel{
				InnerRoom: db.InnerRoom{
					ID:         "1",
					PropertyID: "1",
				},
			},
		},
	}

	t.Run("FromFurniture", func(t *testing.T) {
		furnitureResponse := models.FurnitureResponse{}
		furnitureResponse.FromDbFurniture(furniture)

		assert.Equal(t, furniture.ID, furnitureResponse.ID)
		assert.Equal(t, furniture.Room().PropertyID, furnitureResponse.PropertyID)
		assert.Equal(t, furniture.RoomID, furnitureResponse.RoomID)
		assert.Equal(t, furniture.Name, furnitureResponse.Name)
		assert.Equal(t, furniture.Quantity, furnitureResponse.Quantity)
	})

	t.Run("FurnitureToResponse", func(t *testing.T) {
		furnitureResponse := models.DbFurnitureToResponse(furniture)

		assert.Equal(t, furniture.ID, furnitureResponse.ID)
		assert.Equal(t, furniture.Room().PropertyID, furnitureResponse.PropertyID)
		assert.Equal(t, furniture.RoomID, furnitureResponse.RoomID)
		assert.Equal(t, furniture.Name, furnitureResponse.Name)
		assert.Equal(t, furniture.Quantity, furnitureResponse.Quantity)
	})
}
