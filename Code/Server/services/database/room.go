package database

import (
	"immotep/backend/prisma/db"
	"immotep/backend/services"
)

func CreateRoom(room db.RoomModel, proppertyId string) *db.RoomModel {
	pdb := services.DBclient
	newRoom, err := pdb.Client.Room.CreateOne(
		db.Room.Name.Set(room.Name),
		db.Room.Property.Link(db.Property.ID.Equals(proppertyId)),
	).Exec(pdb.Context)
	if err != nil {
		if _, is := db.IsErrUniqueConstraint(err); is {
			return nil
		}
		panic(err)
	}
	return newRoom
}

func MockCreateRoom(c *services.PrismaDB, room db.RoomModel) db.RoomMockExpectParam {
	return c.Client.Room.CreateOne(
		db.Room.Name.Set(room.Name),
		db.Room.Property.Link(db.Property.ID.Equals("1")),
	)
}

func GetRoomByPropertyID(propertyID string, archived bool) []db.RoomModel {
	pdb := services.DBclient
	rooms, err := pdb.Client.Room.FindMany(
		db.Room.PropertyID.Equals(propertyID),
		db.Room.Archived.Equals(archived),
	).Exec(pdb.Context)
	if err != nil {
		panic(err)
	}
	return rooms
}

func MockGetRoomsByPropertyID(c *services.PrismaDB, archived bool) db.RoomMockExpectParam {
	return c.Client.Room.FindMany(
		db.Room.PropertyID.Equals("1"),
		db.Room.Archived.Equals(archived),
	)
}

func GetRoomByID(id string) *db.RoomModel {
	pdb := services.DBclient
	room, err := pdb.Client.Room.FindUnique(
		db.Room.ID.Equals(id),
	).Exec(pdb.Context)
	if err != nil {
		if db.IsErrNotFound(err) {
			return nil
		}
		panic(err)
	}
	return room
}

func MockGetRoomByID(c *services.PrismaDB) db.RoomMockExpectParam {
	return c.Client.Room.FindUnique(
		db.Room.ID.Equals("1"),
	)
}

func ToggleArchiveRoom(roomId string, archive bool) *db.RoomModel {
	pdb := services.DBclient
	archivedRoom, err := pdb.Client.Room.FindUnique(
		db.Room.ID.Equals(roomId),
	).Update(
		db.Room.Archived.Set(archive),
	).Exec(pdb.Context)
	if err != nil {
		if db.IsErrNotFound(err) {
			return nil
		}
		panic(err)
	}
	return archivedRoom
}

func MockArchiveRoom(c *services.PrismaDB) db.RoomMockExpectParam {
	return c.Client.Room.FindUnique(
		db.Room.ID.Equals("1"),
	).Update(
		db.Room.Archived.Set(true),
	)
}
