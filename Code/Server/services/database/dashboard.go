package database

import (
	"keyz/backend/prisma/db"
	"keyz/backend/services"
)

func GetAllDatasFromProperties(ownerId string) []db.PropertyModel {
	pdb := services.DBclient
	allProperties, err := pdb.Client.Property.FindMany(
		db.Property.OwnerID.Equals(ownerId),
		db.Property.Archived.Equals(false),
	).With(
		db.Property.Owner.Fetch(),
		db.Property.Leases.Fetch().With(
			db.Lease.Tenant.Fetch(),
			db.Lease.Damages.Fetch().With(db.Damage.Room.Fetch()),
			db.Lease.Reports.Fetch().With(
				db.InventoryReport.RoomStates.Fetch().With(db.RoomState.Room.Fetch()),
				db.InventoryReport.FurnitureStates.Fetch().With(db.FurnitureState.Furniture.Fetch()),
			),
		),
		db.Property.LeaseInvite.Fetch(),
		db.Property.Rooms.Fetch().With(db.Room.Furnitures.Fetch()),
	).Exec(pdb.Context)
	if err != nil {
		panic(err)
	}
	return allProperties
}

func MockGetAllDatasFromProperties(c *services.PrismaDB) db.PropertyMockExpectParam {
	return c.Client.Property.FindMany(
		db.Property.OwnerID.Equals("1"),
		db.Property.Archived.Equals(false),
	).With(
		db.Property.Owner.Fetch(),
		db.Property.Leases.Fetch().With(
			db.Lease.Tenant.Fetch(),
			db.Lease.Damages.Fetch().With(db.Damage.Room.Fetch()),
			db.Lease.Reports.Fetch().With(
				db.InventoryReport.RoomStates.Fetch().With(db.RoomState.Room.Fetch()),
				db.InventoryReport.FurnitureStates.Fetch().With(db.FurnitureState.Furniture.Fetch()),
			),
		),
		db.Property.LeaseInvite.Fetch(),
		db.Property.Rooms.Fetch().With(db.Room.Furnitures.Fetch()),
	)
}
