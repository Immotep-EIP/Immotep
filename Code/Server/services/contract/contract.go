package contractservice

import (
	"immotep/backend/database"
	"immotep/backend/prisma/db"
)

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
		if info, is := db.IsErrUniqueConstraint(err); is && info.Fields[0] == db.PendingContract.TenantEmail.Field() {
			return nil
		}
		panic(err)
	}
	return newContract
}
