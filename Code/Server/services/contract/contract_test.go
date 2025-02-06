package contractservice_test

import (
	"errors"
	"testing"
	"time"

	"github.com/steebchen/prisma-client-go/engine/protocol"
	"github.com/stretchr/testify/assert"
	"immotep/backend/database"
	"immotep/backend/prisma/db"
	contractservice "immotep/backend/services/contract"
)

func BuildTestContract() db.ContractModel {
	end := time.Now().Add(time.Hour)
	return db.ContractModel{
		InnerContract: db.InnerContract{
			TenantID:   "1",
			Active:     true,
			StartDate:  time.Now(),
			EndDate:    &end,
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

func BuildTestProperty(id string) db.PropertyModel {
	return db.PropertyModel{
		InnerProperty: db.InnerProperty{
			ID:                  id,
			Name:                "Test",
			Address:             "Test",
			City:                "Test",
			PostalCode:          "Test",
			Country:             "Test",
			AreaSqm:             20.0,
			RentalPricePerMonth: 500,
			DepositPrice:        1000,
			CreatedAt:           time.Now(),
			OwnerID:             "1",
		},
		RelationsProperty: db.RelationsProperty{
			Damages:   []db.DamageModel{},
			Contracts: []db.ContractModel{},
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

func TestCreate(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	tenant := BuildTestTenant("1")
	pendingContract := BuildTestPendingContract()
	contract := BuildTestContract()

	mock.Contract.Expect(
		client.Client.Contract.CreateOne(
			db.Contract.Tenant.Link(db.User.ID.Equals(tenant.ID)),
			db.Contract.Property.Link(db.Property.ID.Equals(pendingContract.PropertyID)),
			db.Contract.StartDate.Set(pendingContract.StartDate),
			db.Contract.EndDate.SetIfPresent(pendingContract.InnerPendingContract.EndDate),
		),
	).Returns(contract)

	mock.PendingContract.Expect(
		client.Client.PendingContract.FindUnique(
			db.PendingContract.ID.Equals(pendingContract.ID),
		).Delete(),
	).Returns(pendingContract)

	newContract := contractservice.Create(pendingContract, tenant)
	assert.NotNil(t, newContract)
	assert.Equal(t, contract.TenantID, newContract.TenantID)
	assert.Equal(t, contract.PropertyID, newContract.PropertyID)
}

func TestCreate_AlreadyExists(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	tenant := BuildTestTenant("1")
	pendingContract := BuildTestPendingContract()

	mock.Contract.Expect(
		client.Client.Contract.CreateOne(
			db.Contract.Tenant.Link(db.User.ID.Equals(tenant.ID)),
			db.Contract.Property.Link(db.Property.ID.Equals(pendingContract.PropertyID)),
			db.Contract.StartDate.Set(pendingContract.StartDate),
			db.Contract.EndDate.SetIfPresent(pendingContract.InnerPendingContract.EndDate),
		),
	).Errors(&protocol.UserFacingError{
		IsPanic:   false,
		ErrorCode: "P2002", // https://www.prisma.io/docs/orm/reference/error-reference
		Meta: protocol.Meta{
			Target: []any{"tenant_id", "property_id"},
		},
		Message: "Unique constraint failed",
	})

	newContract := contractservice.Create(pendingContract, tenant)
	assert.Nil(t, newContract)
}

func TestCreate_NoConnection1(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	tenant := BuildTestTenant("1")
	pendingContract := BuildTestPendingContract()

	mock.Contract.Expect(
		client.Client.Contract.CreateOne(
			db.Contract.Tenant.Link(db.User.ID.Equals(tenant.ID)),
			db.Contract.Property.Link(db.Property.ID.Equals(pendingContract.PropertyID)),
			db.Contract.StartDate.Set(pendingContract.StartDate),
			db.Contract.EndDate.SetIfPresent(pendingContract.InnerPendingContract.EndDate),
		),
	).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		contractservice.Create(pendingContract, tenant)
	})
}

func TestCreate_NoConnection2(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	tenant := BuildTestTenant("1")
	pendingContract := BuildTestPendingContract()
	contract := BuildTestContract()

	mock.Contract.Expect(
		client.Client.Contract.CreateOne(
			db.Contract.Tenant.Link(db.User.ID.Equals(tenant.ID)),
			db.Contract.Property.Link(db.Property.ID.Equals(pendingContract.PropertyID)),
			db.Contract.StartDate.Set(pendingContract.StartDate),
			db.Contract.EndDate.SetIfPresent(pendingContract.InnerPendingContract.EndDate),
		),
	).Returns(contract)

	mock.PendingContract.Expect(
		client.Client.PendingContract.FindUnique(
			db.PendingContract.ID.Equals(pendingContract.ID),
		).Delete(),
	).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		contractservice.Create(pendingContract, tenant)
	})
}

func TestGetPendingByID(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	user := BuildTestPendingContract()

	mock.PendingContract.Expect(
		client.Client.PendingContract.FindUnique(db.PendingContract.ID.Equals("1")),
	).Returns(user)

	foundPending := contractservice.GetPendingById("1")
	assert.NotNil(t, foundPending)
	assert.Equal(t, user.ID, foundPending.ID)
}

func TestGetPendingByID_NotFound(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	mock.PendingContract.Expect(
		client.Client.PendingContract.FindUnique(db.PendingContract.ID.Equals("1")),
	).Errors(db.ErrNotFound)

	foundPending := contractservice.GetPendingById("1")
	assert.Nil(t, foundPending)
}

func TestGetPendingByID_NoConnection(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	mock.PendingContract.Expect(
		client.Client.PendingContract.FindUnique(db.PendingContract.ID.Equals("1")),
	).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		contractservice.GetPendingById("1")
	})
}

func TestCreatePending(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
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

	newContract := contractservice.CreatePending(pendingContract, property.ID)
	assert.NotNil(t, newContract)
	assert.Equal(t, pendingContract.ID, newContract.ID)
}

func TestCreatePending_AlreadyExists1(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
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

	newContract := contractservice.CreatePending(pendingContract, property.ID)
	assert.Nil(t, newContract)
}

func TestCreatePending_AlreadyExists2(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
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

	newContract := contractservice.CreatePending(pendingContract, property.ID)
	assert.Nil(t, newContract)
}

func TestCreatePending_NoConnection(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
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
		contractservice.CreatePending(pendingContract, property.ID)
	})
}

func TestGetCurrentActive(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	contract := BuildTestContract()

	mock.Contract.Expect(
		client.Client.Contract.FindMany(
			db.Contract.PropertyID.Equals(property.ID),
			db.Contract.Active.Equals(true),
		),
	).ReturnsMany([]db.ContractModel{contract})

	activeContract := contractservice.GetCurrentActive(property.ID)
	assert.NotNil(t, activeContract)
	assert.Equal(t, contract.PropertyID, activeContract.PropertyID)
	assert.Equal(t, contract.TenantID, activeContract.TenantID)
}

func TestGetCurrentActive_NotFound(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")

	mock.Contract.Expect(
		client.Client.Contract.FindMany(
			db.Contract.PropertyID.Equals(property.ID),
			db.Contract.Active.Equals(true),
		),
	).Errors(db.ErrNotFound)

	activeContract := contractservice.GetCurrentActive(property.ID)
	assert.Nil(t, activeContract)
}

func TestGetCurrentActive_NoActiveContract(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")

	mock.Contract.Expect(
		client.Client.Contract.FindMany(
			db.Contract.PropertyID.Equals(property.ID),
			db.Contract.Active.Equals(true),
		),
	).ReturnsMany([]db.ContractModel{})

	activeContract := contractservice.GetCurrentActive(property.ID)
	assert.Nil(t, activeContract)
}

func TestGetCurrentActive_MultipleActiveContracts(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
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
		contractservice.GetCurrentActive(property.ID)
	})
}

func TestGetCurrentActive_NoConnection(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")

	mock.Contract.Expect(
		client.Client.Contract.FindMany(
			db.Contract.PropertyID.Equals(property.ID),
			db.Contract.Active.Equals(true),
		),
	).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		contractservice.GetCurrentActive(property.ID)
	})
}

func TestEndContract(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	contract := BuildTestContract()
	contract.Active = false
	endDate := time.Now()

	mock.Contract.Expect(
		client.Client.Contract.FindUnique(
			db.Contract.TenantIDPropertyID(db.Contract.TenantID.Equals("1"), db.Contract.PropertyID.Equals("1")),
		).Update(
			db.Contract.Active.Set(false),
			db.Contract.EndDate.SetIfPresent(&endDate),
		),
	).Returns(contract)

	endedContract := contractservice.EndContract("1", "1", &endDate)
	assert.NotNil(t, endedContract)
	assert.Equal(t, "1", endedContract.TenantID)
	assert.Equal(t, "1", endedContract.PropertyID)
	assert.False(t, endedContract.Active)
}

func TestEndContract_NotFound(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	contract := BuildTestContract()
	endDate := time.Now()

	mock.Contract.Expect(
		client.Client.Contract.FindUnique(
			db.Contract.TenantIDPropertyID(db.Contract.TenantID.Equals(contract.TenantID), db.Contract.PropertyID.Equals(contract.PropertyID)),
		).Update(
			db.Contract.Active.Set(false),
			db.Contract.EndDate.SetIfPresent(&endDate),
		),
	).Errors(db.ErrNotFound)

	endedContract := contractservice.EndContract(contract.PropertyID, contract.TenantID, &endDate)
	assert.Nil(t, endedContract)
}

func TestEndContract_NoConnection(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	contract := BuildTestContract()
	endDate := time.Now()

	mock.Contract.Expect(
		client.Client.Contract.FindUnique(
			db.Contract.TenantIDPropertyID(db.Contract.TenantID.Equals(contract.TenantID), db.Contract.PropertyID.Equals(contract.PropertyID)),
		).Update(
			db.Contract.Active.Set(false),
			db.Contract.EndDate.SetIfPresent(&endDate),
		),
	).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		contractservice.EndContract(contract.PropertyID, contract.TenantID, &endDate)
	})
}
