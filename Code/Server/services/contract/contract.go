package contractservice

import (
	"errors"

	"github.com/steebchen/prisma-client-go/engine/protocol"
	"immotep/backend/database"
	"immotep/backend/prisma/db"
)

func GetCurrentActive(propertyId string) *db.ContractModel {
	pdb := database.DBclient
	c, err := pdb.Client.Contract.FindMany(
		db.Contract.PropertyID.Equals(propertyId),
		db.Contract.Active.Equals(true),
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
		panic("Only one active contract must exist for a property")
	}
	return &c[0]
}

func Create(pendingContract db.PendingContractModel, tenant db.UserModel) *db.ContractModel {
	pdb := database.DBclient
	newContract, err := pdb.Client.Contract.CreateOne(
		db.Contract.Tenant.Link(db.User.ID.Equals(tenant.ID)),
		db.Contract.Property.Link(db.Property.ID.Equals(pendingContract.PropertyID)),
		db.Contract.StartDate.Set(pendingContract.StartDate),
		db.Contract.EndDate.SetIfPresent(pendingContract.InnerPendingContract.EndDate),
	).Exec(pdb.Context)
	if err != nil {
		if _, is := db.IsErrUniqueConstraint(err); is {
			return nil
		}
		panic(err)
	}
	_, err = pdb.Client.PendingContract.FindUnique(
		db.PendingContract.ID.Equals(pendingContract.ID),
	).Delete().Exec(pdb.Context)
	if err != nil {
		panic(err)
	}
	return newContract
}

func GetPendingById(id string) *db.PendingContractModel {
	pdb := database.DBclient
	pc, err := pdb.Client.PendingContract.FindUnique(db.PendingContract.ID.Equals(id)).Exec(pdb.Context)
	if err != nil {
		if db.IsErrNotFound(err) {
			return nil
		}
		panic(err)
	}
	return pc
}

func CreatePending(pendingContract db.PendingContractModel, property db.PropertyModel) *db.PendingContractModel {
	pdb := database.DBclient
	newContract, err := pdb.Client.PendingContract.CreateOne(
		db.PendingContract.TenantEmail.Set(pendingContract.TenantEmail),
		db.PendingContract.StartDate.Set(pendingContract.StartDate),
		db.PendingContract.Property.Link(db.Property.ID.Equals(property.ID)),
		db.PendingContract.EndDate.SetIfPresent(pendingContract.InnerPendingContract.EndDate),
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
	return newContract
}
