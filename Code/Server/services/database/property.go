package database

import (
	"immotep/backend/prisma/db"
	"immotep/backend/services"
)

func GetAllPropertyByOwnerId(ownerId string) []db.PropertyModel {
	pdb := services.DBclient
	allProperties, err := pdb.Client.Property.FindMany(
		db.Property.OwnerID.Equals(ownerId),
	).With(
		db.Property.Damages.Fetch(),
		db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
	).Exec(pdb.Context)
	if err != nil {
		panic(err)
	}
	return allProperties
}

func GetPropertyByID(id string) *db.PropertyModel {
	pdb := services.DBclient
	property, err := pdb.Client.Property.FindUnique(
		db.Property.ID.Equals(id),
	).With(
		db.Property.Damages.Fetch(),
		db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
	).Exec(pdb.Context)
	if err != nil {
		if db.IsErrNotFound(err) {
			return nil
		}
		panic(err)
	}
	return property
}

func CreateProperty(property db.PropertyModel, ownerId string) *db.PropertyModel {
	pdb := services.DBclient
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
	).With(
		db.Property.Contracts.Fetch(),
		db.Property.Damages.Fetch(),
	).Exec(pdb.Context)
	if err != nil {
		if _, is := db.IsErrUniqueConstraint(err); is {
			return nil
		}
		panic(err)
	}
	return newProperty
}

func UpdatePropertyPicture(property db.PropertyModel, image db.ImageModel) *db.PropertyModel {
	pdb := services.DBclient
	newProperty, err := pdb.Client.Property.FindUnique(
		db.Property.ID.Equals(property.ID),
	).With(
		db.Property.Damages.Fetch(),
		db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
	).Update(
		db.Property.Picture.Link(db.Image.ID.Equals(image.ID)),
	).Exec(pdb.Context)
	if err != nil {
		if db.IsErrNotFound(err) {
			return nil
		}
		panic(err)
	}
	return newProperty
}
