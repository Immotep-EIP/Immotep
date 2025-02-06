package database

import (
	"errors"

	"github.com/steebchen/prisma-client-go/engine/protocol"
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

func GetRoomByPropertyID(propertyID string) []db.RoomModel {
	pdb := services.DBclient
	rooms, err := pdb.Client.Room.FindMany(
		db.Room.PropertyID.Equals(propertyID),
	).Exec(pdb.Context)
	if err != nil {
		panic(err)
	}
	return rooms
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

func DeleteRoom(id string) bool {
	pdb := services.DBclient
	_, err := pdb.Client.Room.FindUnique(
		db.Room.ID.Equals(id),
	).Delete().Exec(pdb.Context)
	if err != nil {
		// https://www.prisma.io/docs/orm/reference/error-reference#p2025
		var ufr *protocol.UserFacingError
		if ok := errors.As(err, &ufr); ok && ufr.ErrorCode == "P2025" {
			return false
		}
		panic(err)
	}
	return true
}
