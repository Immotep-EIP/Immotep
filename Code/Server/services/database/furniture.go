package database

import (
	"immotep/backend/prisma/db"
	"immotep/backend/services"
)

func CreateFurniture(furniture db.FurnitureModel, roomId string) *db.FurnitureModel {
	pdb := services.DBclient
	newFurniture, err := pdb.Client.Furniture.CreateOne(
		db.Furniture.Name.Set(furniture.Name),
		db.Furniture.Room.Link(db.Room.ID.Equals(roomId)),
		db.Furniture.Quantity.Set(furniture.Quantity),
	).With(
		db.Furniture.Room.Fetch(),
	).Exec(pdb.Context)
	if err != nil {
		if _, is := db.IsErrUniqueConstraint(err); is {
			return nil
		}
		panic(err)
	}
	return newFurniture
}

func GetFurnitureByRoomID(roomID string) []db.FurnitureModel {
	pdb := services.DBclient
	furnitures, err := pdb.Client.Furniture.FindMany(
		db.Furniture.RoomID.Equals(roomID),
	).With(
		db.Furniture.Room.Fetch(),
	).Exec(pdb.Context)
	if err != nil {
		panic(err)
	}
	return furnitures
}

func GetFurnitureByID(id string) *db.FurnitureModel {
	pdb := services.DBclient
	furniture, err := pdb.Client.Furniture.FindUnique(
		db.Furniture.ID.Equals(id),
	).With(
		db.Furniture.Room.Fetch(),
	).Exec(pdb.Context)
	if err != nil {
		if db.IsErrNotFound(err) {
			return nil
		}
		panic(err)
	}
	return furniture
}

func DeleteFurniture(id string) bool {
	pdb := services.DBclient
	_, err := pdb.Client.Furniture.FindUnique(
		db.Furniture.ID.Equals(id),
	).Delete().Exec(pdb.Context)
	if err != nil {
		if db.IsErrNotFound(err) {
			return false
		}
		panic(err)
	}
	return true
}
