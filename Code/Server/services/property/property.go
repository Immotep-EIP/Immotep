package propertyservice

import (
	"immotep/backend/database"
	"immotep/backend/prisma/db"
)

func GetAllByOwnerId(ownerId string) []db.PropertyModel {
	pdb := database.DBclient
	allProperties, err := pdb.Client.Property.FindMany(
		db.Property.OwnerID.Equals(ownerId),
	).Exec(pdb.Context)
	if err != nil {
		panic(err)
	}
	return allProperties
}

func GetByID(id string) *db.PropertyModel {
	pdb := database.DBclient
	property, err := pdb.Client.Property.FindUnique(db.Property.ID.Equals(id)).Exec(pdb.Context)
	if err != nil {
		if db.IsErrNotFound(err) {
			return nil
		}
		panic(err)
	}
	return property
}