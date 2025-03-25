package database_test

import (
	"errors"
	"testing"
	"time"

	"github.com/steebchen/prisma-client-go/engine/protocol"
	"github.com/stretchr/testify/assert"
	"immotep/backend/prisma/db"
	"immotep/backend/services"
	"immotep/backend/services/database"
)

func BuildTestContract() db.ContractModel {
	end := time.Now().Add(time.Hour)
	return db.ContractModel{
		InnerContract: db.InnerContract{
			ID:         "1",
			Active:     true,
			StartDate:  time.Now(),
			EndDate:    &end,
			TenantID:   "1",
			PropertyID: "1",
			CreatedAt:  time.Now(),
		},
	}
}

func BuildTestPendingContract() db.PendingContractModel {
	end := time.Now().Add(time.Hour)
	return db.PendingContractModel{
		InnerPendingContract: db.InnerPendingContract{
			ID:          "1",
			TenantEmail: "test.test@example.com",
			StartDate:   time.Now(),
			EndDate:     &end,
			PropertyID:  "1",
			CreatedAt:   time.Now(),
		},
	}
}

func BuildTestTenant(id string) db.UserModel {
	return db.UserModel{
		InnerUser: db.InnerUser{
			ID:        id,
			Email:     "test" + id + "@example.com",
			Firstname: "Test",
			Lastname:  "User",
			Password:  "Password123",
			Role:      db.RoleTenant,
		},
	}
}

func TestCreateContract(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	tenant := BuildTestTenant("1")
	pendingContract := BuildTestPendingContract()
	contract := BuildTestContract()

	mock.Contract.Expect(
		client.Client.Contract.CreateOne(
			db.Contract.StartDate.Set(pendingContract.StartDate),
			db.Contract.Tenant.Link(db.User.ID.Equals(tenant.ID)),
			db.Contract.Property.Link(db.Property.ID.Equals(pendingContract.PropertyID)),
			db.Contract.EndDate.SetIfPresent(pendingContract.InnerPendingContract.EndDate),
		),
	).Returns(contract)

	mock.PendingContract.Expect(
		client.Client.PendingContract.FindUnique(
			db.PendingContract.ID.Equals(pendingContract.ID),
		).Delete(),
	).Returns(pendingContract)

	newContract := database.CreateContract(pendingContract, tenant)
	assert.NotNil(t, newContract)
	assert.Equal(t, contract.TenantID, newContract.TenantID)
	assert.Equal(t, contract.PropertyID, newContract.PropertyID)
}

func TestCreateContract_NoConnection1(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	tenant := BuildTestTenant("1")
	pendingContract := BuildTestPendingContract()

	mock.Contract.Expect(
		client.Client.Contract.CreateOne(
			db.Contract.StartDate.Set(pendingContract.StartDate),
			db.Contract.Tenant.Link(db.User.ID.Equals(tenant.ID)),
			db.Contract.Property.Link(db.Property.ID.Equals(pendingContract.PropertyID)),
			db.Contract.EndDate.SetIfPresent(pendingContract.InnerPendingContract.EndDate),
		),
	).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.CreateContract(pendingContract, tenant)
	})
}

func TestCreateContract_NoConnection2(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	tenant := BuildTestTenant("1")
	pendingContract := BuildTestPendingContract()
	contract := BuildTestContract()

	mock.Contract.Expect(
		client.Client.Contract.CreateOne(
			db.Contract.StartDate.Set(pendingContract.StartDate),
			db.Contract.Tenant.Link(db.User.ID.Equals(tenant.ID)),
			db.Contract.Property.Link(db.Property.ID.Equals(pendingContract.PropertyID)),
			db.Contract.EndDate.SetIfPresent(pendingContract.InnerPendingContract.EndDate),
		),
	).Returns(contract)

	mock.PendingContract.Expect(
		client.Client.PendingContract.FindUnique(
			db.PendingContract.ID.Equals(pendingContract.ID),
		).Delete(),
	).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.CreateContract(pendingContract, tenant)
	})
}

func TestGetPendingContractByID(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	user := BuildTestPendingContract()

	mock.PendingContract.Expect(
		client.Client.PendingContract.FindUnique(db.PendingContract.ID.Equals("1")),
	).Returns(user)

	foundPending := database.GetPendingContractById("1")
	assert.NotNil(t, foundPending)
	assert.Equal(t, user.ID, foundPending.ID)
}

func TestGetPendingContractByID_NotFound(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	mock.PendingContract.Expect(
		client.Client.PendingContract.FindUnique(db.PendingContract.ID.Equals("1")),
	).Errors(db.ErrNotFound)

	foundPending := database.GetPendingContractById("1")
	assert.Nil(t, foundPending)
}

func TestGetPendingContractByID_NoConnection(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	mock.PendingContract.Expect(
		client.Client.PendingContract.FindUnique(db.PendingContract.ID.Equals("1")),
	).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.GetPendingContractById("1")
	})
}

func TestCreatePendingContract(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	pendingContract := BuildTestPendingContract()
	mock.PendingContract.Expect(
		client.Client.PendingContract.CreateOne(
			db.PendingContract.TenantEmail.Set(pendingContract.TenantEmail),
			db.PendingContract.StartDate.Set(pendingContract.StartDate),
			db.PendingContract.Property.Link(db.Property.ID.Equals("1")),
			db.PendingContract.EndDate.SetIfPresent(pendingContract.InnerPendingContract.EndDate),
		).With(
			db.PendingContract.Property.Fetch().With(db.Property.Owner.Fetch()),
		),
	).Returns(pendingContract)

	newContract := database.CreatePendingContract(pendingContract, property.ID)
	assert.NotNil(t, newContract)
	assert.Equal(t, pendingContract.ID, newContract.ID)
}

func TestCreatePendingContract_AlreadyExists1(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	pendingContract := BuildTestPendingContract()
	mock.PendingContract.Expect(
		client.Client.PendingContract.CreateOne(
			db.PendingContract.TenantEmail.Set(pendingContract.TenantEmail),
			db.PendingContract.StartDate.Set(pendingContract.StartDate),
			db.PendingContract.Property.Link(db.Property.ID.Equals("1")),
			db.PendingContract.EndDate.SetIfPresent(pendingContract.InnerPendingContract.EndDate),
		).With(
			db.PendingContract.Property.Fetch().With(db.Property.Owner.Fetch()),
		),
	).Errors(&protocol.UserFacingError{
		IsPanic:   false,
		ErrorCode: "P2002", // https://www.prisma.io/docs/orm/reference/error-reference
		Meta: protocol.Meta{
			Target: []any{"tenant_email"},
		},
		Message: "Unique constraint failed",
	})

	newContract := database.CreatePendingContract(pendingContract, property.ID)
	assert.Nil(t, newContract)
}

func TestCreatePendingContract_AlreadyExists2(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	pendingContract := BuildTestPendingContract()
	mock.PendingContract.Expect(
		client.Client.PendingContract.CreateOne(
			db.PendingContract.TenantEmail.Set(pendingContract.TenantEmail),
			db.PendingContract.StartDate.Set(pendingContract.StartDate),
			db.PendingContract.Property.Link(db.Property.ID.Equals("1")),
			db.PendingContract.EndDate.SetIfPresent(pendingContract.InnerPendingContract.EndDate),
		).With(
			db.PendingContract.Property.Fetch().With(db.Property.Owner.Fetch()),
		),
	).Errors(&protocol.UserFacingError{
		IsPanic:   false,
		ErrorCode: "P2014", // https://www.prisma.io/docs/orm/reference/error-reference
		Meta: protocol.Meta{
			Target: []any{"property_id"},
		},
		Message: "Unique constraint failed",
	})

	newContract := database.CreatePendingContract(pendingContract, property.ID)
	assert.Nil(t, newContract)
}

func TestCreatePendingContract_NoConnection(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	pendingContract := BuildTestPendingContract()
	mock.PendingContract.Expect(
		client.Client.PendingContract.CreateOne(
			db.PendingContract.TenantEmail.Set(pendingContract.TenantEmail),
			db.PendingContract.StartDate.Set(pendingContract.StartDate),
			db.PendingContract.Property.Link(db.Property.ID.Equals("1")),
			db.PendingContract.EndDate.SetIfPresent(pendingContract.InnerPendingContract.EndDate),
		).With(
			db.PendingContract.Property.Fetch().With(db.Property.Owner.Fetch()),
		),
	).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.CreatePendingContract(pendingContract, property.ID)
	})
}

func TestGetCurrentActiveContract(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	contract := BuildTestContract()

	mock.Contract.Expect(
		client.Client.Contract.FindMany(
			db.Contract.PropertyID.Equals(property.ID),
			db.Contract.Active.Equals(true),
		),
	).ReturnsMany([]db.ContractModel{contract})

	activeContract := database.GetCurrentActiveContract(property.ID)
	assert.NotNil(t, activeContract)
	assert.Equal(t, contract.PropertyID, activeContract.PropertyID)
	assert.Equal(t, contract.TenantID, activeContract.TenantID)
}

func TestGetCurrentActiveContract_NotFound(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")

	mock.Contract.Expect(
		client.Client.Contract.FindMany(
			db.Contract.PropertyID.Equals(property.ID),
			db.Contract.Active.Equals(true),
		),
	).Errors(db.ErrNotFound)

	activeContract := database.GetCurrentActiveContract(property.ID)
	assert.Nil(t, activeContract)
}

func TestGetCurrentActiveContract_NoActiveContract(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")

	mock.Contract.Expect(
		client.Client.Contract.FindMany(
			db.Contract.PropertyID.Equals(property.ID),
			db.Contract.Active.Equals(true),
		),
	).ReturnsMany([]db.ContractModel{})

	activeContract := database.GetCurrentActiveContract(property.ID)
	assert.Nil(t, activeContract)
}

func TestGetCurrentActiveContract_MultipleActiveContracts(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	contract1 := BuildTestContract()
	contract2 := BuildTestContract()

	mock.Contract.Expect(
		client.Client.Contract.FindMany(
			db.Contract.PropertyID.Equals(property.ID),
			db.Contract.Active.Equals(true),
		),
	).ReturnsMany([]db.ContractModel{contract1, contract2})

	assert.Panics(t, func() {
		database.GetCurrentActiveContract(property.ID)
	})
}

func TestGetCurrentActiveContract_NoConnection(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")

	mock.Contract.Expect(
		client.Client.Contract.FindMany(
			db.Contract.PropertyID.Equals(property.ID),
			db.Contract.Active.Equals(true),
		),
	).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.GetCurrentActiveContract(property.ID)
	})
}

func TestEndContract(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	contract := BuildTestContract()
	contract.Active = false
	endDate := time.Now()

	mock.Contract.Expect(
		client.Client.Contract.FindUnique(
			db.Contract.ID.Equals(contract.ID),
		).Update(
			db.Contract.Active.Set(false),
			db.Contract.EndDate.SetIfPresent(&endDate),
		),
	).Returns(contract)

	endedContract := database.EndContract(contract.ID, &endDate)
	assert.NotNil(t, endedContract)
	assert.Equal(t, "1", endedContract.TenantID)
	assert.Equal(t, "1", endedContract.PropertyID)
	assert.False(t, endedContract.Active)
}

func TestEndContract_NotFound(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	contract := BuildTestContract()
	endDate := time.Now()

	mock.Contract.Expect(
		client.Client.Contract.FindUnique(
			db.Contract.ID.Equals(contract.ID),
		).Update(
			db.Contract.Active.Set(false),
			db.Contract.EndDate.SetIfPresent(&endDate),
		),
	).Errors(db.ErrNotFound)

	endedContract := database.EndContract(contract.ID, &endDate)
	assert.Nil(t, endedContract)
}

func TestEndContract_NoConnection(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	contract := BuildTestContract()
	endDate := time.Now()

	mock.Contract.Expect(
		client.Client.Contract.FindUnique(
			db.Contract.ID.Equals(contract.ID),
		).Update(
			db.Contract.Active.Set(false),
			db.Contract.EndDate.SetIfPresent(&endDate),
		),
	).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.EndContract(contract.ID, &endDate)
	})
}

func TestGetCurrentActiveContractWithInfos(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	contract := BuildTestContract()

	mock.Contract.Expect(
		client.Client.Contract.FindMany(
			db.Contract.PropertyID.Equals(property.ID),
			db.Contract.Active.Equals(true),
		).With(
			db.Contract.Tenant.Fetch(),
			db.Contract.Property.Fetch().With(db.Property.Owner.Fetch()),
		),
	).ReturnsMany([]db.ContractModel{contract})

	activeContract := database.GetCurrentActiveContractWithInfos(property.ID)
	assert.NotNil(t, activeContract)
	assert.Equal(t, contract.PropertyID, activeContract.PropertyID)
	assert.Equal(t, contract.TenantID, activeContract.TenantID)
}

func TestGetCurrentActiveContractWithInfos_NotFound(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")

	mock.Contract.Expect(
		client.Client.Contract.FindMany(
			db.Contract.PropertyID.Equals(property.ID),
			db.Contract.Active.Equals(true),
		).With(
			db.Contract.Tenant.Fetch(),
			db.Contract.Property.Fetch().With(db.Property.Owner.Fetch()),
		),
	).Errors(db.ErrNotFound)

	activeContract := database.GetCurrentActiveContractWithInfos(property.ID)
	assert.Nil(t, activeContract)
}

func TestGetCurrentActiveContractWithInfos_NoActiveContract(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")

	mock.Contract.Expect(
		client.Client.Contract.FindMany(
			db.Contract.PropertyID.Equals(property.ID),
			db.Contract.Active.Equals(true),
		).With(
			db.Contract.Tenant.Fetch(),
			db.Contract.Property.Fetch().With(db.Property.Owner.Fetch()),
		),
	).ReturnsMany([]db.ContractModel{})

	activeContract := database.GetCurrentActiveContractWithInfos(property.ID)
	assert.Nil(t, activeContract)
}

func TestGetCurrentActiveContractWithInfos_MultipleActiveContracts(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	contract1 := BuildTestContract()
	contract2 := BuildTestContract()

	mock.Contract.Expect(
		client.Client.Contract.FindMany(
			db.Contract.PropertyID.Equals(property.ID),
			db.Contract.Active.Equals(true),
		).With(
			db.Contract.Tenant.Fetch(),
			db.Contract.Property.Fetch().With(db.Property.Owner.Fetch()),
		),
	).ReturnsMany([]db.ContractModel{contract1, contract2})

	assert.Panics(t, func() {
		database.GetCurrentActiveContractWithInfos(property.ID)
	})
}

func TestGetCurrentActiveContractWithInfos_NoConnection(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")

	mock.Contract.Expect(
		client.Client.Contract.FindMany(
			db.Contract.PropertyID.Equals(property.ID),
			db.Contract.Active.Equals(true),
		).With(
			db.Contract.Tenant.Fetch(),
			db.Contract.Property.Fetch().With(db.Property.Owner.Fetch()),
		),
	).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.GetCurrentActiveContractWithInfos(property.ID)
	})
}

func TestGetTenantCurrentActiveContract(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	tenant := BuildTestTenant("1")
	contract := BuildTestContract()

	mock.Contract.Expect(
		client.Client.Contract.FindMany(
			db.Contract.TenantID.Equals(tenant.ID),
			db.Contract.Active.Equals(true),
		),
	).ReturnsMany([]db.ContractModel{contract})

	activeContract := database.GetTenantCurrentActiveContract(tenant.ID)
	assert.NotNil(t, activeContract)
	assert.Equal(t, contract.TenantID, activeContract.TenantID)
	assert.Equal(t, contract.PropertyID, activeContract.PropertyID)
}

func TestGetTenantCurrentActiveContract_NotFound(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	tenant := BuildTestTenant("1")

	mock.Contract.Expect(
		client.Client.Contract.FindMany(
			db.Contract.TenantID.Equals(tenant.ID),
			db.Contract.Active.Equals(true),
		),
	).Errors(db.ErrNotFound)

	activeContract := database.GetTenantCurrentActiveContract(tenant.ID)
	assert.Nil(t, activeContract)
}

func TestGetTenantCurrentActiveContract_NoActiveContract(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	tenant := BuildTestTenant("1")

	mock.Contract.Expect(
		client.Client.Contract.FindMany(
			db.Contract.TenantID.Equals(tenant.ID),
			db.Contract.Active.Equals(true),
		),
	).ReturnsMany([]db.ContractModel{})

	activeContract := database.GetTenantCurrentActiveContract(tenant.ID)
	assert.Nil(t, activeContract)
}

func TestGetTenantCurrentActiveContract_MultipleActiveContracts(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	tenant := BuildTestTenant("1")
	contract1 := BuildTestContract()
	contract2 := BuildTestContract()

	mock.Contract.Expect(
		client.Client.Contract.FindMany(
			db.Contract.TenantID.Equals(tenant.ID),
			db.Contract.Active.Equals(true),
		),
	).ReturnsMany([]db.ContractModel{contract1, contract2})

	assert.Panics(t, func() {
		database.GetTenantCurrentActiveContract(tenant.ID)
	})
}

func TestGetTenantCurrentActiveContract_NoConnection(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	tenant := BuildTestTenant("1")

	mock.Contract.Expect(
		client.Client.Contract.FindMany(
			db.Contract.TenantID.Equals(tenant.ID),
			db.Contract.Active.Equals(true),
		),
	).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.GetTenantCurrentActiveContract(tenant.ID)
	})
}

func TestGetCurrentPendingContract(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	pendingContract := BuildTestPendingContract()

	mock.PendingContract.Expect(
		client.Client.PendingContract.FindUnique(
			db.PendingContract.PropertyID.Equals(property.ID),
		),
	).Returns(pendingContract)

	foundPending := database.GetCurrentPendingContract(property.ID)
	assert.NotNil(t, foundPending)
	assert.Equal(t, pendingContract.ID, foundPending.ID)
}

func TestGetCurrentPendingContract_NotFound(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")

	mock.PendingContract.Expect(
		client.Client.PendingContract.FindUnique(
			db.PendingContract.PropertyID.Equals(property.ID),
		),
	).Errors(db.ErrNotFound)

	foundPending := database.GetCurrentPendingContract(property.ID)
	assert.Nil(t, foundPending)
}

func TestGetCurrentPendingContract_NoConnection(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")

	mock.PendingContract.Expect(
		client.Client.PendingContract.FindUnique(
			db.PendingContract.PropertyID.Equals(property.ID),
		),
	).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.GetCurrentPendingContract(property.ID)
	})
}

func TestDeleteCurrentPendingContract(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	pendingContract := BuildTestPendingContract()

	mock.PendingContract.Expect(
		client.Client.PendingContract.FindUnique(
			db.PendingContract.PropertyID.Equals(property.ID),
		).Delete(),
	).Returns(pendingContract)

	assert.NotPanics(t, func() {
		database.DeleteCurrentPendingContract(property.ID)
	})
}

func TestDeleteCurrentPendingContract_NotFound(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")

	mock.PendingContract.Expect(
		client.Client.PendingContract.FindUnique(
			db.PendingContract.PropertyID.Equals(property.ID),
		).Delete(),
	).Errors(db.ErrNotFound)

	assert.Panics(t, func() {
		database.DeleteCurrentPendingContract(property.ID)
	})
}

func TestDeleteCurrentPendingContract_NoConnection(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")

	mock.PendingContract.Expect(
		client.Client.PendingContract.FindUnique(
			db.PendingContract.PropertyID.Equals(property.ID),
		).Delete(),
	).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.DeleteCurrentPendingContract(property.ID)
	})
}
