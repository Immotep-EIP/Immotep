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

func MockCreateFurniture(c *services.PrismaDB, furniture db.FurnitureModel) db.FurnitureMockExpectParam {
	return c.Client.Furniture.CreateOne(
		db.Furniture.Name.Set(furniture.Name),
		db.Furniture.Room.Link(db.Room.ID.Equals("1")),
		db.Furniture.Quantity.Set(furniture.Quantity),
	).With(
		db.Furniture.Room.Fetch(),
	)
}

func GetFurnituresByRoomID(roomID string, archived bool) []db.FurnitureModel {
	pdb := services.DBclient
	furnitures, err := pdb.Client.Furniture.FindMany(
		db.Furniture.RoomID.Equals(roomID),
		db.Furniture.Archived.Equals(archived),
	).With(
		db.Furniture.Room.Fetch(),
	).Exec(pdb.Context)
	if err != nil {
		panic(err)
	}
	return furnitures
}

func MockGetFurnituresByRoomID(c *services.PrismaDB, archived bool) db.FurnitureMockExpectParam {
	return c.Client.Furniture.FindMany(
		db.Furniture.RoomID.Equals("1"),
		db.Furniture.Archived.Equals(archived),
	).With(
		db.Furniture.Room.Fetch(),
	)
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

func MockGetFurnitureByID(c *services.PrismaDB) db.FurnitureMockExpectParam {
	return c.Client.Furniture.FindUnique(
		db.Furniture.ID.Equals("1"),
	).With(
		db.Furniture.Room.Fetch(),
	)
}

func ToggleArchiveFurniture(furnitureId string, archive bool) *db.FurnitureModel {
	pdb := services.DBclient
	archivedFurniture, err := pdb.Client.Furniture.FindUnique(
		db.Furniture.ID.Equals(furnitureId),
	).With(
		db.Furniture.Room.Fetch(),
	).Update(
		db.Furniture.Archived.Set(archive),
	).Exec(pdb.Context)
	if err != nil {
		if db.IsErrNotFound(err) {
			return nil
		}
		panic(err)
	}
	return archivedFurniture
}

func MockArchiveFurniture(c *services.PrismaDB) db.FurnitureMockExpectParam {
	return c.Client.Furniture.FindUnique(
		db.Furniture.ID.Equals("1"),
	).With(
		db.Furniture.Room.Fetch(),
	).Update(
		db.Furniture.Archived.Set(true),
	)
}
