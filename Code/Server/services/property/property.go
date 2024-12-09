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

func Create(property db.PropertyModel, ownerId string) *db.PropertyModel {
	pdb := database.DBclient
	newProperty, err := pdb.Client.Property.CreateOne(
		db.Property.Name.Set(property.Name),
		db.Property.Address.Set(property.Address),
		db.Property.City.Set(property.City),
		db.Property.PostalCode.Set(property.PostalCode),
		db.Property.Country.Set(property.Country),
		db.Property.AreaSqm.Set(property.AreaSqm),
		db.Property.RentalPricePerMonth.Set(property.RentalPricePerMonth),
		db.Property.DepositPrice.Set(property.DepositPrice),
		db.Property.Owner.Link(db.User.ID.Equals(ownerId)),
		db.Property.Picture.SetIfPresent(property.InnerProperty.Picture),
	).Exec(pdb.Context)
	if err != nil {
		if _, is := db.IsErrUniqueConstraint(err); is {
			return nil
		}
		panic(err)
	}
	return newProperty
}
