package database

import (
	"immotep/backend/models"
	"immotep/backend/prisma/db"
	"immotep/backend/services"
)

func GetPropertiesByOwnerId(ownerId string, archived bool) []db.PropertyModel {
	pdb := services.DBclient
	allProperties, err := pdb.Client.Property.FindMany(
		db.Property.OwnerID.Equals(ownerId),
		db.Property.Archived.Equals(archived),
	).With(
		db.Property.Leases.Fetch().With(
			db.Lease.Tenant.Fetch(),
			db.Lease.Damages.Fetch(db.Damage.FixedAt.IsNull()),
		),
		db.Property.LeaseInvite.Fetch(),
	).Exec(pdb.Context)
	if err != nil {
		panic(err)
	}
	return allProperties
}

func MockGetAllPropertyByOwnerId(c *services.PrismaDB, archived bool) db.PropertyMockExpectParam {
	return c.Client.Property.FindMany(
		db.Property.OwnerID.Equals("1"),
		db.Property.Archived.Equals(archived),
	).With(
		db.Property.Leases.Fetch().With(
			db.Lease.Tenant.Fetch(),
			db.Lease.Damages.Fetch(db.Damage.FixedAt.IsNull()),
		),
		db.Property.LeaseInvite.Fetch(),
	)
}

func GetPropertyByID(id string) *db.PropertyModel {
	pdb := services.DBclient
	property, err := pdb.Client.Property.FindUnique(
		db.Property.ID.Equals(id),
	).With(
		db.Property.Leases.Fetch().With(
			db.Lease.Tenant.Fetch(),
			db.Lease.Damages.Fetch(db.Damage.FixedAt.IsNull()),
		),
		db.Property.LeaseInvite.Fetch(),
	).Exec(pdb.Context)
	if err != nil {
		if db.IsErrNotFound(err) {
			return nil
		}
		panic(err)
	}
	return property
}

func MockGetPropertyByID(c *services.PrismaDB) db.PropertyMockExpectParam {
	return c.Client.Property.FindUnique(
		db.Property.ID.Equals("1"),
	).With(
		db.Property.Leases.Fetch().With(
			db.Lease.Tenant.Fetch(),
			db.Lease.Damages.Fetch(db.Damage.FixedAt.IsNull()),
		),
		db.Property.LeaseInvite.Fetch(),
	)
}

func GetPropertyInventory(id string) *db.PropertyModel {
	pdb := services.DBclient
	property, err := pdb.Client.Property.FindUnique(
		db.Property.ID.Equals(id),
	).With(
		db.Property.Leases.Fetch().With(
			db.Lease.Tenant.Fetch(),
			db.Lease.Damages.Fetch(db.Damage.FixedAt.IsNull()),
		),
		db.Property.LeaseInvite.Fetch(),
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

func MockGetPropertyInventory(c *services.PrismaDB) db.PropertyMockExpectParam {
	return c.Client.Property.FindUnique(
		db.Property.ID.Equals("1"),
	).With(
		db.Property.Leases.Fetch().With(
			db.Lease.Tenant.Fetch(),
			db.Lease.Damages.Fetch(db.Damage.FixedAt.IsNull()),
		),
		db.Property.LeaseInvite.Fetch(),
		db.Property.Rooms.Fetch().With(db.Room.Furnitures.Fetch()),
	)
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
	).Exec(pdb.Context)
	if err != nil {
		if _, is := db.IsErrUniqueConstraint(err); is {
			return nil
		}
		panic(err)
	}
	return newProperty
}

func MockCreateProperty(c *services.PrismaDB, property db.PropertyModel) db.PropertyMockExpectParam {
	return c.Client.Property.CreateOne(
		db.Property.Name.Set(property.Name),
		db.Property.Address.Set(property.Address),
		db.Property.City.Set(property.City),
		db.Property.PostalCode.Set(property.PostalCode),
		db.Property.Country.Set(property.Country),
		db.Property.AreaSqm.Set(property.AreaSqm),
		db.Property.RentalPricePerMonth.Set(property.RentalPricePerMonth),
		db.Property.DepositPrice.Set(property.DepositPrice),
		db.Property.Owner.Link(db.User.ID.Equals("1")),
		db.Property.ApartmentNumber.SetIfPresent(property.InnerProperty.ApartmentNumber),
	)
}

func UpdateProperty(property db.PropertyModel, req models.PropertyUpdateRequest) *db.PropertyModel {
	pdb := services.DBclient
	newProperty, err := pdb.Client.Property.FindUnique(
		db.Property.ID.Equals(property.ID),
	).Update(
		db.Property.Name.SetIfPresent(req.Name),
		db.Property.Address.SetIfPresent(req.Address),
		db.Property.ApartmentNumber.SetIfPresent(req.ApartmentNumber),
		db.Property.City.SetIfPresent(req.City),
		db.Property.PostalCode.SetIfPresent(req.PostalCode),
		db.Property.Country.SetIfPresent(req.Country),
		db.Property.AreaSqm.SetIfPresent(req.AreaSqm),
		db.Property.RentalPricePerMonth.SetIfPresent(req.RentalPricePerMonth),
		db.Property.DepositPrice.SetIfPresent(req.DepositPrice),
	).Exec(pdb.Context)
	if err != nil {
		if _, is := db.IsErrUniqueConstraint(err); is {
			return nil
		}
		panic(err)
	}
	return newProperty
}

func MockUpdateProperty(c *services.PrismaDB, uProperty models.PropertyUpdateRequest) db.PropertyMockExpectParam {
	return c.Client.Property.FindUnique(
		db.Property.ID.Equals("1"),
	).Update(
		db.Property.Name.SetIfPresent(uProperty.Name),
		db.Property.Address.SetIfPresent(uProperty.Address),
		db.Property.ApartmentNumber.SetIfPresent(uProperty.ApartmentNumber),
		db.Property.City.SetIfPresent(uProperty.City),
		db.Property.PostalCode.SetIfPresent(uProperty.PostalCode),
		db.Property.Country.SetIfPresent(uProperty.Country),
		db.Property.AreaSqm.SetIfPresent(uProperty.AreaSqm),
		db.Property.RentalPricePerMonth.SetIfPresent(uProperty.RentalPricePerMonth),
		db.Property.DepositPrice.SetIfPresent(uProperty.DepositPrice),
	)
}

func UpdatePropertyPicture(property db.PropertyModel, image db.ImageModel) *db.PropertyModel {
	pdb := services.DBclient
	newProperty, err := pdb.Client.Property.FindUnique(
		db.Property.ID.Equals(property.ID),
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

func MockUpdatePropertyPicture(c *services.PrismaDB) db.PropertyMockExpectParam {
	return c.Client.Property.FindUnique(
		db.Property.ID.Equals("1"),
	).Update(
		db.Property.Picture.Link(db.Image.ID.Equals("1")),
	)
}

func ArchiveProperty(propertyId string, archive bool) *db.PropertyModel {
	pdb := services.DBclient
	archivedProperty, err := pdb.Client.Property.FindUnique(
		db.Property.ID.Equals(propertyId),
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

func MockArchiveProperty(c *services.PrismaDB) db.PropertyMockExpectParam {
	return c.Client.Property.FindUnique(
		db.Property.ID.Equals("1"),
	).Update(
		db.Property.Archived.Set(true),
	)
}
