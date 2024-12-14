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

func TestCreateAlreadyExists(t *testing.T) {
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
		),
	).Returns(pendingContract)

	newContract := contractservice.CreatePending(pendingContract, property)
	assert.NotNil(t, newContract)
	assert.Equal(t, pendingContract.ID, newContract.ID)
}

func TestCreatePendingAlreadyExists1(t *testing.T) {
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
		),
	).Errors(&protocol.UserFacingError{
		IsPanic:   false,
		ErrorCode: "P2002", // https://www.prisma.io/docs/orm/reference/error-reference
		Meta: protocol.Meta{
			Target: []any{"tenant_email"},
		},
		Message: "Unique constraint failed",
	})

	newContract := contractservice.CreatePending(pendingContract, property)
	assert.Nil(t, newContract)
}

func TestCreatePendingAlreadyExists2(t *testing.T) {
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
		),
	).Errors(&protocol.UserFacingError{
		IsPanic:   false,
		ErrorCode: "P2014", // https://www.prisma.io/docs/orm/reference/error-reference
		Meta: protocol.Meta{
			Target: []any{"property_id"},
		},
		Message: "Unique constraint failed",
	})

	newContract := contractservice.CreatePending(pendingContract, property)
	assert.Nil(t, newContract)
}
