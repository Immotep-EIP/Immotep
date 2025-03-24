package database

import (
	"errors"

	"github.com/steebchen/prisma-client-go/engine/protocol"
	"immotep/backend/prisma/db"
	"immotep/backend/services"
)

func GetCurrentActiveLease(propertyId string) *db.LeaseModel {
	pdb := services.DBclient
	c, err := pdb.Client.Lease.FindMany(
		db.Lease.PropertyID.Equals(propertyId),
		db.Lease.Active.Equals(true),
	).Exec(pdb.Context)
	if err != nil {
		if db.IsErrNotFound(err) {
			return nil
		}
		panic(err)
	}
	l := len(c)
	if l == 0 {
		return nil
	}
	if l > 1 {
		panic("Only one active lease must exist for a property")
	}
	return &c[0]
}

func GetCurrentActiveLeaseWithInfos(propertyId string) *db.LeaseModel {
	pdb := services.DBclient
	c, err := pdb.Client.Lease.FindMany(
		db.Lease.PropertyID.Equals(propertyId),
		db.Lease.Active.Equals(true),
	).With(
		db.Lease.Tenant.Fetch(),
		db.Lease.Property.Fetch().With(db.Property.Owner.Fetch()),
	).Exec(pdb.Context)
	if err != nil {
		if db.IsErrNotFound(err) {
			return nil
		}
		panic(err)
	}
	l := len(c)
	if l == 0 {
		return nil
	}
	if l > 1 {
		panic("Only one active lease must exist for a property")
	}
	return &c[0]
}

func GetTenantCurrentActiveLease(tenantId string) *db.LeaseModel {
	pdb := services.DBclient
	c, err := pdb.Client.Lease.FindMany(
		db.Lease.TenantID.Equals(tenantId),
		db.Lease.Active.Equals(true),
	).Exec(pdb.Context)
	if err != nil {
		if db.IsErrNotFound(err) {
			return nil
		}
		panic(err)
	}
	l := len(c)
	if l == 0 {
		return nil
	}
	if l > 1 {
		panic("Only one active lease must exist for a tenant")
	}
	return &c[0]
}

func CreateLease(pendingContract db.PendingContractModel, tenant db.UserModel) db.LeaseModel {
	pdb := services.DBclient
	newLease, err := pdb.Client.Lease.CreateOne(
		db.Lease.StartDate.Set(pendingContract.StartDate),
		db.Lease.Tenant.Link(db.User.ID.Equals(tenant.ID)),
		db.Lease.Property.Link(db.Property.ID.Equals(pendingContract.PropertyID)),
		db.Lease.EndDate.SetIfPresent(pendingContract.InnerPendingContract.EndDate),
	).Exec(pdb.Context)
	if err != nil {
		panic(err)
	}
	_, err = pdb.Client.PendingContract.FindUnique(
		db.PendingContract.ID.Equals(pendingContract.ID),
	).Delete().Exec(pdb.Context)
	if err != nil {
		panic(err)
	}
	return *newLease
}

func EndLease(id string, endDate *db.DateTime) *db.LeaseModel {
	pdb := services.DBclient
	newLease, err := pdb.Client.Lease.FindUnique(
		db.Lease.ID.Equals(id),
	).Update(
		db.Lease.Active.Set(false),
		db.Lease.EndDate.SetIfPresent(endDate),
	).Exec(pdb.Context)
	if err != nil {
		if db.IsErrNotFound(err) {
			return nil
		}
		panic(err)
	}
	return newLease
}

func GetPendingContractById(id string) *db.PendingContractModel {
	pdb := services.DBclient
	pc, err := pdb.Client.PendingContract.FindUnique(db.PendingContract.ID.Equals(id)).Exec(pdb.Context)
	if err != nil {
		if db.IsErrNotFound(err) {
			return nil
		}
		panic(err)
	}
	return pc
}

func GetCurrentPendingContract(propertyId string) *db.PendingContractModel {
	pdb := services.DBclient
	pc, err := pdb.Client.PendingContract.FindUnique(db.PendingContract.PropertyID.Equals(propertyId)).Exec(pdb.Context)
	if err != nil {
		if db.IsErrNotFound(err) {
			return nil
		}
		panic(err)
	}
	return pc
}

func CreatePendingContract(pendingContract db.PendingContractModel, propertyId string) *db.PendingContractModel {
	pdb := services.DBclient
	newLease, err := pdb.Client.PendingContract.CreateOne(
		db.PendingContract.TenantEmail.Set(pendingContract.TenantEmail),
		db.PendingContract.StartDate.Set(pendingContract.StartDate),
		db.PendingContract.Property.Link(db.Property.ID.Equals(propertyId)),
		db.PendingContract.EndDate.SetIfPresent(pendingContract.InnerPendingContract.EndDate),
	).With(
		db.PendingContract.Property.Fetch().With(db.Property.Owner.Fetch()),
	).Exec(pdb.Context)
	if err != nil {
		// https://www.prisma.io/docs/orm/reference/error-reference#p2014
		var ufr *protocol.UserFacingError
		if ok := errors.As(err, &ufr); ok && ufr.ErrorCode == "P2014" {
			return nil
		}
		if _, is := db.IsErrUniqueConstraint(err); is {
			return nil
		}
		panic(err)
	}
	return newLease
}

func DeleteCurrentPendingContract(propertyId string) {
	pdb := services.DBclient
	_, err := pdb.Client.PendingContract.FindUnique(
		db.PendingContract.PropertyID.Equals(propertyId),
	).Delete().Exec(pdb.Context)
	if err != nil {
		panic(err)
	}
}
