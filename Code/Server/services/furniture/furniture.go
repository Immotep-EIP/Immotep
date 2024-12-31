package furnitureservice

import (
	"immotep/backend/database"
	"immotep/backend/prisma/db"
)

func Create(furniture db.FurnitureModel, roomId string) *db.FurnitureModel {
	pdb := database.DBclient
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

func GetByRoomID(roomID string) []db.FurnitureModel {
	pdb := database.DBclient
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

func GetByID(id string) *db.FurnitureModel {
	pdb := database.DBclient
	furniture, err := pdb.Client.Furniture.FindUnique(
		db.Furniture.ID.Equals(id),
	).With(
		db.Furniture.Room.Fetch(),
	).Exec(pdb.Context)
	if err != nil {
		return nil
	}
	return furniture
}

func Delete(id string) bool {
	pdb := database.DBclient
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
