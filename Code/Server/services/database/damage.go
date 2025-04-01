package database

import (
	"immotep/backend/prisma/db"
	"immotep/backend/services"
	"immotep/backend/utils"
)

func CreateDamage(damage db.DamageModel, leaseId string, picturesId []string) db.DamageModel {
	params := make([]db.DamageSetParam, 0, len(picturesId))
	for _, id := range picturesId {
		params = append(params, db.Damage.Pictures.Link(db.Image.ID.Equals(id)))
	}

	pdb := services.DBclient
	newDamage, err := pdb.Client.Damage.CreateOne(
		db.Damage.Comment.Set(damage.Comment),
		db.Damage.Priority.Set(damage.Priority),
		db.Damage.Lease.Link(db.Lease.ID.Equals(leaseId)),
		db.Damage.Room.Link(db.Room.ID.Equals(damage.RoomID)),
		params...,
	).With(
		db.Damage.Lease.Fetch().With(db.Lease.Tenant.Fetch()),
		db.Damage.Room.Fetch(),
	).Exec(pdb.Context)
	if err != nil {
		panic(err)
	}
	return *newDamage
}

func GetDamagesByPropertyID(propertyID string, fixed bool) []db.DamageModel {
	fixedParam := utils.Ternary(fixed, db.Damage.FixedAt.Gt(db.DateTime{}), db.Damage.FixedAt.IsNull())

	pdb := services.DBclient
	damages, err := pdb.Client.Damage.FindMany(
		db.Damage.Lease.Where(db.Lease.PropertyID.Equals(propertyID)),
		fixedParam,
	).With(
		db.Damage.Lease.Fetch().With(db.Lease.Tenant.Fetch()),
		db.Damage.Room.Fetch(),
		db.Damage.Pictures.Fetch(),
	).Exec(pdb.Context)
	if err != nil {
		if db.IsErrNotFound(err) {
			return nil
		}
		panic(err)
	}
	return damages
}

func GetDamagesByLeaseID(leaseID string, fixed bool) []db.DamageModel {
	fixedParam := utils.Ternary(fixed, db.Damage.FixedAt.Gt(db.DateTime{}), db.Damage.FixedAt.IsNull())

	pdb := services.DBclient
	damages, err := pdb.Client.Damage.FindMany(
		db.Damage.LeaseID.Equals(leaseID),
		fixedParam,
	).With(
		db.Damage.Lease.Fetch().With(db.Lease.Tenant.Fetch()),
		db.Damage.Room.Fetch(),
		db.Damage.Pictures.Fetch(),
	).Exec(pdb.Context)
	if err != nil {
		if db.IsErrNotFound(err) {
			return nil
		}
		panic(err)
	}
	return damages
}
