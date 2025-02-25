package database

import (
	"immotep/backend/models"
	"immotep/backend/prisma/db"
	"immotep/backend/services"
)

func GetAllPropertyByOwnerId(ownerId string) []db.PropertyModel {
	pdb := services.DBclient
	allProperties, err := pdb.Client.Property.FindMany(
		db.Property.OwnerID.Equals(ownerId),
		db.Property.Archived.Equals(false),
	).With(
		db.Property.Damages.Fetch(),
		db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
		db.Property.PendingContract.Fetch(),
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
		db.Property.PendingContract.Fetch(),
	).Exec(pdb.Context)
	if err != nil {
		if db.IsErrNotFound(err) {
			return nil
		}
		panic(err)
	}
	return property
}

func GetPropertyInventory(id string) *db.PropertyModel {
	pdb := services.DBclient
	property, err := pdb.Client.Property.FindUnique(
		db.Property.ID.Equals(id),
	).With(
		db.Property.Damages.Fetch(),
		db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
		db.Property.PendingContract.Fetch(),
		db.Property.Rooms.Fetch().With(db.Room.Furnitures.Fetch()),
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
		db.Property.ApartmentNumber.SetIfPresent(property.InnerProperty.ApartmentNumber),
	).With(
		db.Property.Damages.Fetch(),
		db.Property.Contracts.Fetch(),
		db.Property.PendingContract.Fetch(),
	).Exec(pdb.Context)
	if err != nil {
		if _, is := db.IsErrUniqueConstraint(err); is {
			return nil
		}
		panic(err)
	}
	return newProperty
}

func UpdateProperty(id string, property models.PropertyUpdateRequest) *db.PropertyModel {
	pdb := services.DBclient
	newProperty, err := pdb.Client.Property.FindUnique(db.Property.ID.Equals(id)).With(
		db.Property.Damages.Fetch(),
		db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
		db.Property.PendingContract.Fetch(),
	).Update(
		db.Property.Name.SetIfPresent(property.Name),
		db.Property.Address.SetIfPresent(property.Address),
		db.Property.ApartmentNumber.SetIfPresent(property.ApartmentNumber),
		db.Property.City.SetIfPresent(property.City),
		db.Property.PostalCode.SetIfPresent(property.PostalCode),
		db.Property.Country.SetIfPresent(property.Country),
		db.Property.AreaSqm.SetIfPresent(property.AreaSqm),
		db.Property.RentalPricePerMonth.SetIfPresent(property.RentalPricePerMonth),
		db.Property.DepositPrice.SetIfPresent(property.DepositPrice),
	).Exec(pdb.Context)
	if err != nil {
		if db.IsErrNotFound(err) {
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
		db.Property.PendingContract.Fetch(),
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

func ToggleArchiveProperty(propertyId string, archive bool) *db.PropertyModel {
	pdb := services.DBclient
	archivedProperty, err := pdb.Client.Property.FindUnique(
		db.Property.ID.Equals(propertyId),
	).With(
		db.Property.Damages.Fetch(),
		db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
		db.Property.PendingContract.Fetch(),
	).Update(
		db.Property.Archived.Set(archive),
	).Exec(pdb.Context)
	if err != nil {
		if db.IsErrNotFound(err) {
			return nil
		}
		panic(err)
	}
	return archivedProperty
}
