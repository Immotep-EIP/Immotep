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

func BuildTestLease() db.LeaseModel {
	end := time.Now().Add(time.Hour)
	return db.LeaseModel{
		InnerLease: db.InnerLease{
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

func TestCreateLease(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	tenant := BuildTestTenant("1")
	pendingContract := BuildTestPendingContract()
	lease := BuildTestLease()

	mock.Lease.Expect(
		client.Client.Lease.CreateOne(
			db.Lease.StartDate.Set(pendingContract.StartDate),
			db.Lease.Tenant.Link(db.User.ID.Equals(tenant.ID)),
			db.Lease.Property.Link(db.Property.ID.Equals(pendingContract.PropertyID)),
			db.Lease.EndDate.SetIfPresent(pendingContract.InnerPendingContract.EndDate),
		),
	).Returns(lease)

	mock.PendingContract.Expect(
		client.Client.PendingContract.FindUnique(
			db.PendingContract.ID.Equals(pendingContract.ID),
		).Delete(),
	).Returns(pendingContract)

	newLease := database.CreateLease(pendingContract, tenant)
	assert.NotNil(t, newLease)
	assert.Equal(t, lease.TenantID, newLease.TenantID)
	assert.Equal(t, lease.PropertyID, newLease.PropertyID)
}

func TestCreateLease_NoConnection1(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	tenant := BuildTestTenant("1")
	pendingContract := BuildTestPendingContract()

	mock.Lease.Expect(
		client.Client.Lease.CreateOne(
			db.Lease.StartDate.Set(pendingContract.StartDate),
			db.Lease.Tenant.Link(db.User.ID.Equals(tenant.ID)),
			db.Lease.Property.Link(db.Property.ID.Equals(pendingContract.PropertyID)),
			db.Lease.EndDate.SetIfPresent(pendingContract.InnerPendingContract.EndDate),
		),
	).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.CreateLease(pendingContract, tenant)
	})
}

func TestCreateLease_NoConnection2(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	tenant := BuildTestTenant("1")
	pendingContract := BuildTestPendingContract()
	lease := BuildTestLease()

	mock.Lease.Expect(
		client.Client.Lease.CreateOne(
			db.Lease.StartDate.Set(pendingContract.StartDate),
			db.Lease.Tenant.Link(db.User.ID.Equals(tenant.ID)),
			db.Lease.Property.Link(db.Property.ID.Equals(pendingContract.PropertyID)),
			db.Lease.EndDate.SetIfPresent(pendingContract.InnerPendingContract.EndDate),
		),
	).Returns(lease)

	mock.PendingContract.Expect(
		client.Client.PendingContract.FindUnique(
			db.PendingContract.ID.Equals(pendingContract.ID),
		).Delete(),
	).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.CreateLease(pendingContract, tenant)
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

	newLease := database.CreatePendingContract(pendingContract, property.ID)
	assert.NotNil(t, newLease)
	assert.Equal(t, pendingContract.ID, newLease.ID)
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

	newLease := database.CreatePendingContract(pendingContract, property.ID)
	assert.Nil(t, newLease)
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

	newLease := database.CreatePendingContract(pendingContract, property.ID)
	assert.Nil(t, newLease)
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

func TestGetCurrentActiveLease(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	lease := BuildTestLease()

	mock.Lease.Expect(
		client.Client.Lease.FindMany(
			db.Lease.PropertyID.Equals(property.ID),
			db.Lease.Active.Equals(true),
		),
	).ReturnsMany([]db.LeaseModel{lease})

	activeLease := database.GetCurrentActiveLease(property.ID)
	assert.NotNil(t, activeLease)
	assert.Equal(t, lease.PropertyID, activeLease.PropertyID)
	assert.Equal(t, lease.TenantID, activeLease.TenantID)
}

func TestGetCurrentActiveLease_NotFound(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")

	mock.Lease.Expect(
		client.Client.Lease.FindMany(
			db.Lease.PropertyID.Equals(property.ID),
			db.Lease.Active.Equals(true),
		),
	).Errors(db.ErrNotFound)

	activeLease := database.GetCurrentActiveLease(property.ID)
	assert.Nil(t, activeLease)
}

func TestGetCurrentActiveLease_NoActiveLease(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")

	mock.Lease.Expect(
		client.Client.Lease.FindMany(
			db.Lease.PropertyID.Equals(property.ID),
			db.Lease.Active.Equals(true),
		),
	).ReturnsMany([]db.LeaseModel{})

	activeLease := database.GetCurrentActiveLease(property.ID)
	assert.Nil(t, activeLease)
}

func TestGetCurrentActiveLease_MultipleActiveLeases(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	lease1 := BuildTestLease()
	lease2 := BuildTestLease()

	mock.Lease.Expect(
		client.Client.Lease.FindMany(
			db.Lease.PropertyID.Equals(property.ID),
			db.Lease.Active.Equals(true),
		),
	).ReturnsMany([]db.LeaseModel{lease1, lease2})

	assert.Panics(t, func() {
		database.GetCurrentActiveLease(property.ID)
	})
}

func TestGetCurrentActiveLease_NoConnection(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")

	mock.Lease.Expect(
		client.Client.Lease.FindMany(
			db.Lease.PropertyID.Equals(property.ID),
			db.Lease.Active.Equals(true),
		),
	).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.GetCurrentActiveLease(property.ID)
	})
}

func TestEndLease(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	lease := BuildTestLease()
	lease.Active = false
	endDate := time.Now()

	mock.Lease.Expect(
		client.Client.Lease.FindUnique(
			db.Lease.ID.Equals(lease.ID),
		).Update(
			db.Lease.Active.Set(false),
			db.Lease.EndDate.SetIfPresent(&endDate),
		),
	).Returns(lease)

	endedLease := database.EndLease(lease.ID, &endDate)
	assert.NotNil(t, endedLease)
	assert.Equal(t, "1", endedLease.TenantID)
	assert.Equal(t, "1", endedLease.PropertyID)
	assert.False(t, endedLease.Active)
}

func TestEndLease_NotFound(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	lease := BuildTestLease()
	endDate := time.Now()

	mock.Lease.Expect(
		client.Client.Lease.FindUnique(
			db.Lease.ID.Equals(lease.ID),
		).Update(
			db.Lease.Active.Set(false),
			db.Lease.EndDate.SetIfPresent(&endDate),
		),
	).Errors(db.ErrNotFound)

	endedLease := database.EndLease(lease.ID, &endDate)
	assert.Nil(t, endedLease)
}

func TestEndLease_NoConnection(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	lease := BuildTestLease()
	endDate := time.Now()

	mock.Lease.Expect(
		client.Client.Lease.FindUnique(
			db.Lease.ID.Equals(lease.ID),
		).Update(
			db.Lease.Active.Set(false),
			db.Lease.EndDate.SetIfPresent(&endDate),
		),
	).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.EndLease(lease.ID, &endDate)
	})
}

func TestGetCurrentActiveLeaseWithInfos(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	lease := BuildTestLease()

	mock.Lease.Expect(
		client.Client.Lease.FindMany(
			db.Lease.PropertyID.Equals(property.ID),
			db.Lease.Active.Equals(true),
		).With(
			db.Lease.Tenant.Fetch(),
			db.Lease.Property.Fetch().With(db.Property.Owner.Fetch()),
		),
	).ReturnsMany([]db.LeaseModel{lease})

	activeLease := database.GetCurrentActiveLeaseWithInfos(property.ID)
	assert.NotNil(t, activeLease)
	assert.Equal(t, lease.PropertyID, activeLease.PropertyID)
	assert.Equal(t, lease.TenantID, activeLease.TenantID)
}

func TestGetCurrentActiveLeaseWithInfos_NotFound(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")

	mock.Lease.Expect(
		client.Client.Lease.FindMany(
			db.Lease.PropertyID.Equals(property.ID),
			db.Lease.Active.Equals(true),
		).With(
			db.Lease.Tenant.Fetch(),
			db.Lease.Property.Fetch().With(db.Property.Owner.Fetch()),
		),
	).Errors(db.ErrNotFound)

	activeLease := database.GetCurrentActiveLeaseWithInfos(property.ID)
	assert.Nil(t, activeLease)
}

func TestGetCurrentActiveLeaseWithInfos_NoActiveLease(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")

	mock.Lease.Expect(
		client.Client.Lease.FindMany(
			db.Lease.PropertyID.Equals(property.ID),
			db.Lease.Active.Equals(true),
		).With(
			db.Lease.Tenant.Fetch(),
			db.Lease.Property.Fetch().With(db.Property.Owner.Fetch()),
		),
	).ReturnsMany([]db.LeaseModel{})

	activeLease := database.GetCurrentActiveLeaseWithInfos(property.ID)
	assert.Nil(t, activeLease)
}

func TestGetCurrentActiveLeaseWithInfos_MultipleActiveLeases(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	lease1 := BuildTestLease()
	lease2 := BuildTestLease()

	mock.Lease.Expect(
		client.Client.Lease.FindMany(
			db.Lease.PropertyID.Equals(property.ID),
			db.Lease.Active.Equals(true),
		).With(
			db.Lease.Tenant.Fetch(),
			db.Lease.Property.Fetch().With(db.Property.Owner.Fetch()),
		),
	).ReturnsMany([]db.LeaseModel{lease1, lease2})

	assert.Panics(t, func() {
		database.GetCurrentActiveLeaseWithInfos(property.ID)
	})
}

func TestGetCurrentActiveLeaseWithInfos_NoConnection(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")

	mock.Lease.Expect(
		client.Client.Lease.FindMany(
			db.Lease.PropertyID.Equals(property.ID),
			db.Lease.Active.Equals(true),
		).With(
			db.Lease.Tenant.Fetch(),
			db.Lease.Property.Fetch().With(db.Property.Owner.Fetch()),
		),
	).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.GetCurrentActiveLeaseWithInfos(property.ID)
	})
}

func TestGetTenantCurrentActiveLease(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	tenant := BuildTestTenant("1")
	lease := BuildTestLease()

	mock.Lease.Expect(
		client.Client.Lease.FindMany(
			db.Lease.TenantID.Equals(tenant.ID),
			db.Lease.Active.Equals(true),
		),
	).ReturnsMany([]db.LeaseModel{lease})

	activeLease := database.GetTenantCurrentActiveLease(tenant.ID)
	assert.NotNil(t, activeLease)
	assert.Equal(t, lease.TenantID, activeLease.TenantID)
	assert.Equal(t, lease.PropertyID, activeLease.PropertyID)
}

func TestGetTenantCurrentActiveLease_NotFound(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	tenant := BuildTestTenant("1")

	mock.Lease.Expect(
		client.Client.Lease.FindMany(
			db.Lease.TenantID.Equals(tenant.ID),
			db.Lease.Active.Equals(true),
		),
	).Errors(db.ErrNotFound)

	activeLease := database.GetTenantCurrentActiveLease(tenant.ID)
	assert.Nil(t, activeLease)
}

func TestGetTenantCurrentActiveLease_NoActiveLease(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	tenant := BuildTestTenant("1")

	mock.Lease.Expect(
		client.Client.Lease.FindMany(
			db.Lease.TenantID.Equals(tenant.ID),
			db.Lease.Active.Equals(true),
		),
	).ReturnsMany([]db.LeaseModel{})

	activeLease := database.GetTenantCurrentActiveLease(tenant.ID)
	assert.Nil(t, activeLease)
}

func TestGetTenantCurrentActiveLease_MultipleActiveLeases(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	tenant := BuildTestTenant("1")
	lease1 := BuildTestLease()
	lease2 := BuildTestLease()

	mock.Lease.Expect(
		client.Client.Lease.FindMany(
			db.Lease.TenantID.Equals(tenant.ID),
			db.Lease.Active.Equals(true),
		),
	).ReturnsMany([]db.LeaseModel{lease1, lease2})

	assert.Panics(t, func() {
		database.GetTenantCurrentActiveLease(tenant.ID)
	})
}

func TestGetTenantCurrentActiveLease_NoConnection(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	tenant := BuildTestTenant("1")

	mock.Lease.Expect(
		client.Client.Lease.FindMany(
			db.Lease.TenantID.Equals(tenant.ID),
			db.Lease.Active.Equals(true),
		),
	).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.GetTenantCurrentActiveLease(tenant.ID)
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
