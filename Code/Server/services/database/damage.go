package database

import (
	"time"

	"immotep/backend/models"
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
	).OrderBy(
		db.Damage.FixedAt.Order(db.SortOrderDesc),
		db.Damage.CreatedAt.Order(db.SortOrderDesc),
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
	).OrderBy(
		db.Damage.FixedAt.Order(db.SortOrderDesc),
		db.Damage.CreatedAt.Order(db.SortOrderDesc),
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

func GetDamageByID(damageID string) *db.DamageModel {
	pdb := services.DBclient
	damage, err := pdb.Client.Damage.FindUnique(
		db.Damage.ID.Equals(damageID),
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
	return damage
}

func UpdateDamageTenant(damage db.DamageModel, req models.DamageTenantUpdateRequest, picturesId []string) *db.DamageModel {
	pdb := services.DBclient

	updates := []db.DamageSetParam{
		db.Damage.Comment.SetIfPresent(req.Comment),
		db.Damage.Priority.SetIfPresent(req.Priority),
	}
	for _, picId := range picturesId {
		updates = append(updates, db.Damage.Pictures.Link(db.Image.ID.Equals(picId)))
	}

	dmg, err := pdb.Client.Damage.FindUnique(
		db.Damage.ID.Equals(damage.ID),
	).With(
		db.Damage.Lease.Fetch().With(db.Lease.Tenant.Fetch()),
		db.Damage.Room.Fetch(),
	).Update(
		updates...,
	).Exec(pdb.Context)
	if err != nil {
		if _, is := db.IsErrUniqueConstraint(err); is {
			return nil
		}
		panic(err)
	}
	return dmg
}

func UpdateDamageOwner(damage db.DamageModel, req models.DamageOwnerUpdateRequest) *db.DamageModel {
	pdb := services.DBclient

	updates := []db.DamageSetParam{
		db.Damage.Read.SetIfPresent(req.Read),
		db.Damage.FixPlannedAt.SetIfPresent(req.FixPlannedAt),
	}

	dmg, err := pdb.Client.Damage.FindUnique(
		db.Damage.ID.Equals(damage.ID),
	).With(
		db.Damage.Lease.Fetch().With(db.Lease.Tenant.Fetch()),
		db.Damage.Room.Fetch(),
	).Update(
		updates...,
	).Exec(pdb.Context)
	if err != nil {
		if _, is := db.IsErrUniqueConstraint(err); is {
			return nil
		}
		panic(err)
	}
	return dmg
}

func MarkDamageAsFixed(damage db.DamageModel, role db.Role) db.DamageModel {
	var params []db.DamageSetParam
	if role == db.RoleTenant {
		params = append(params, db.Damage.FixedTenant.Set(true))
		damage.FixedTenant = true
	}
	if role == db.RoleOwner {
		params = append(params, db.Damage.FixedOwner.Set(true))
		damage.FixedOwner = true
	}
	if damage.IsFixed() {
		params = append(params, db.Damage.FixedAt.Set(time.Now()))
	}

	pdb := services.DBclient
	newDamage, err := pdb.Client.Damage.FindUnique(
		db.Damage.ID.Equals(damage.ID),
	).With(
		db.Damage.Lease.Fetch().With(db.Lease.Tenant.Fetch()),
		db.Damage.Room.Fetch(),
		db.Damage.Pictures.Fetch(),
	).Update(
		params...,
	).Exec(pdb.Context)
	if err != nil {
		panic(err)
	}
	return *newDamage
}
