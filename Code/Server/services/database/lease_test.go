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

func BuildTestLeaseInvite() db.LeaseInviteModel {
	end := time.Now().Add(time.Hour)
	return db.LeaseInviteModel{
		InnerLeaseInvite: db.InnerLeaseInvite{
			ID:          "1",
			TenantEmail: "test@example.com",
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

// #############################################################################

func TestCreateLease(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	tenant := BuildTestTenant("1")
	leaseInvite := BuildTestLeaseInvite()
	lease := BuildTestLease()
	m.Lease.Expect(database.MockCreateLease(c, leaseInvite)).Returns(lease)
	m.LeaseInvite.Expect(database.MockDeleteLeaseInviteById(c)).Returns(leaseInvite)

	newLease := database.CreateLease(leaseInvite, tenant)
	assert.NotNil(t, newLease)
	assert.Equal(t, lease.TenantID, newLease.TenantID)
	assert.Equal(t, lease.PropertyID, newLease.PropertyID)
}

func TestCreateLease_NoConnection1(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	tenant := BuildTestTenant("1")
	leaseInvite := BuildTestLeaseInvite()
	m.Lease.Expect(database.MockCreateLease(c, leaseInvite)).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.CreateLease(leaseInvite, tenant)
	})
}

func TestCreateLease_NoConnection2(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	tenant := BuildTestTenant("1")
	leaseInvite := BuildTestLeaseInvite()
	lease := BuildTestLease()
	m.Lease.Expect(database.MockCreateLease(c, leaseInvite)).Returns(lease)
	m.LeaseInvite.Expect(database.MockDeleteLeaseInviteById(c)).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.CreateLease(leaseInvite, tenant)
	})
}

// #############################################################################

func TestGetLeaseInviteByID(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	user := BuildTestLeaseInvite()
	m.LeaseInvite.Expect(database.MockGetLeaseInviteByID(c)).Returns(user)

	foundPending := database.GetLeaseInviteById("1")
	assert.NotNil(t, foundPending)
	assert.Equal(t, user.ID, foundPending.ID)
}

func TestGetLeaseInviteByID_NotFound(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.LeaseInvite.Expect(database.MockGetLeaseInviteByID(c)).Errors(db.ErrNotFound)

	foundPending := database.GetLeaseInviteById("1")
	assert.Nil(t, foundPending)
}

func TestGetLeaseInviteByID_NoConnection(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.LeaseInvite.Expect(database.MockGetLeaseInviteByID(c)).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.GetLeaseInviteById("1")
	})
}

// #############################################################################

func TestCreateLeaseInvite(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	leaseInvite := BuildTestLeaseInvite()
	m.LeaseInvite.Expect(database.MockCreateLeaseInvite(c, leaseInvite)).Returns(leaseInvite)

	newLease := database.CreateLeaseInvite(leaseInvite, "1")
	assert.NotNil(t, newLease)
	assert.Equal(t, leaseInvite.ID, newLease.ID)
}

func TestCreateLeaseInvite_AlreadyExists1(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	leaseInvite := BuildTestLeaseInvite()
	m.LeaseInvite.Expect(database.MockCreateLeaseInvite(c, leaseInvite)).Errors(&protocol.UserFacingError{
		IsPanic:   false,
		ErrorCode: "P2002", // https://www.prisma.io/docs/orm/reference/error-reference
		Meta: protocol.Meta{
			Target: []any{"tenant_email"},
		},
		Message: "Unique constraint failed",
	})

	newLease := database.CreateLeaseInvite(leaseInvite, "1")
	assert.Nil(t, newLease)
}

func TestCreateLeaseInvite_AlreadyExists2(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	leaseInvite := BuildTestLeaseInvite()
	m.LeaseInvite.Expect(database.MockCreateLeaseInvite(c, leaseInvite)).Errors(&protocol.UserFacingError{
		IsPanic:   false,
		ErrorCode: "P2014", // https://www.prisma.io/docs/orm/reference/error-reference
		Meta: protocol.Meta{
			Target: []any{"property_id"},
		},
		Message: "Unique constraint failed",
	})

	newLease := database.CreateLeaseInvite(leaseInvite, "1")
	assert.Nil(t, newLease)
}

func TestCreateLeaseInvite_NoConnection(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	leaseInvite := BuildTestLeaseInvite()
	m.LeaseInvite.Expect(database.MockCreateLeaseInvite(c, leaseInvite)).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.CreateLeaseInvite(leaseInvite, "1")
	})
}

// #############################################################################

func TestGetCurrentActiveLeaseByProperty(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	lease := BuildTestLease()
	m.Lease.Expect(database.MockGetCurrentActiveLeaseByProperty(c)).ReturnsMany([]db.LeaseModel{lease})

	activeLease := database.GetCurrentActiveLeaseByProperty("1")
	assert.NotNil(t, activeLease)
	assert.Equal(t, lease.PropertyID, activeLease.PropertyID)
	assert.Equal(t, lease.TenantID, activeLease.TenantID)
}

func TestGetCurrentActiveLeaseByProperty_NotFound(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.Lease.Expect(database.MockGetCurrentActiveLeaseByProperty(c)).Errors(db.ErrNotFound)

	activeLease := database.GetCurrentActiveLeaseByProperty("1")
	assert.Nil(t, activeLease)
}

func TestGetCurrentActiveLeaseByProperty_NoActiveLease(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.Lease.Expect(database.MockGetCurrentActiveLeaseByProperty(c)).ReturnsMany([]db.LeaseModel{})

	activeLease := database.GetCurrentActiveLeaseByProperty("1")
	assert.Nil(t, activeLease)
}

func TestGetCurrentActiveLeaseByProperty_MultipleActiveLeases(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	lease1 := BuildTestLease()
	lease2 := BuildTestLease()
	m.Lease.Expect(database.MockGetCurrentActiveLeaseByProperty(c)).ReturnsMany([]db.LeaseModel{lease1, lease2})

	assert.Panics(t, func() {
		database.GetCurrentActiveLeaseByProperty("1")
	})
}

func TestGetCurrentActiveLeaseByProperty_NoConnection(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.Lease.Expect(database.MockGetCurrentActiveLeaseByProperty(c)).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.GetCurrentActiveLeaseByProperty("1")
	})
}

// #############################################################################

func TestEndLease(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	lease := BuildTestLease()
	lease.Active = false
	endDate := time.Now()
	m.Lease.Expect(database.MockEndLease(c, &endDate)).Returns(lease)

	endedLease := database.EndLease(lease.ID, &endDate)
	assert.NotNil(t, endedLease)
	assert.Equal(t, "1", endedLease.TenantID)
	assert.Equal(t, "1", endedLease.PropertyID)
	assert.False(t, endedLease.Active)
}

func TestEndLease_NotFound(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	lease := BuildTestLease()
	endDate := time.Now()
	m.Lease.Expect(database.MockEndLease(c, &endDate)).Errors(db.ErrNotFound)

	endedLease := database.EndLease(lease.ID, &endDate)
	assert.Nil(t, endedLease)
}

func TestEndLease_NoConnection(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	lease := BuildTestLease()
	endDate := time.Now()
	m.Lease.Expect(database.MockEndLease(c, &endDate)).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.EndLease(lease.ID, &endDate)
	})
}

// #############################################################################

func TestGetCurrentActiveLeaseByTenant(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	tenant := BuildTestTenant("1")
	lease := BuildTestLease()
	m.Lease.Expect(database.MockGetCurrentActiveLeaseByTenant(c)).ReturnsMany([]db.LeaseModel{lease})

	activeLease := database.GetCurrentActiveLeaseByTenant(tenant.ID)
	assert.NotNil(t, activeLease)
	assert.Equal(t, lease.TenantID, activeLease.TenantID)
	assert.Equal(t, lease.PropertyID, activeLease.PropertyID)
}

func TestGetCurrentActiveLeaseByTenant_NotFound(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	tenant := BuildTestTenant("1")
	m.Lease.Expect(database.MockGetCurrentActiveLeaseByTenant(c)).Errors(db.ErrNotFound)

	activeLease := database.GetCurrentActiveLeaseByTenant(tenant.ID)
	assert.Nil(t, activeLease)
}

func TestGetCurrentActiveLeaseByTenant_NoActiveLease(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	tenant := BuildTestTenant("1")

	m.Lease.Expect(database.MockGetCurrentActiveLeaseByTenant(c)).ReturnsMany([]db.LeaseModel{})

	activeLease := database.GetCurrentActiveLeaseByTenant(tenant.ID)
	assert.Nil(t, activeLease)
}

func TestGetCurrentActiveLeaseByTenant_MultipleActiveLeases(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	tenant := BuildTestTenant("1")
	lease1 := BuildTestLease()
	lease2 := BuildTestLease()

	m.Lease.Expect(database.MockGetCurrentActiveLeaseByTenant(c)).ReturnsMany([]db.LeaseModel{lease1, lease2})

	assert.Panics(t, func() {
		database.GetCurrentActiveLeaseByTenant(tenant.ID)
	})
}

func TestGetCurrentActiveLeaseByTenant_NoConnection(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	tenant := BuildTestTenant("1")

	m.Lease.Expect(database.MockGetCurrentActiveLeaseByTenant(c)).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.GetCurrentActiveLeaseByTenant(tenant.ID)
	})
}

// #############################################################################

func TestGetCurrentLeaseInvite(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	leaseInvite := BuildTestLeaseInvite()
	m.LeaseInvite.Expect(database.MockGetCurrentLeaseInvite(c)).Returns(leaseInvite)

	foundPending := database.GetCurrentLeaseInvite("1")
	assert.NotNil(t, foundPending)
	assert.Equal(t, leaseInvite.ID, foundPending.ID)
}

func TestGetCurrentLeaseInvite_NotFound(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.LeaseInvite.Expect(database.MockGetCurrentLeaseInvite(c)).Errors(db.ErrNotFound)

	foundPending := database.GetCurrentLeaseInvite("1")
	assert.Nil(t, foundPending)
}

func TestGetCurrentLeaseInvite_NoConnection(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.LeaseInvite.Expect(database.MockGetCurrentLeaseInvite(c)).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.GetCurrentLeaseInvite("1")
	})
}

// #############################################################################

func TestDeleteCurrentLeaseInvite(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	leaseInvite := BuildTestLeaseInvite()
	m.LeaseInvite.Expect(database.MockDeleteCurrentLeaseInvite(c)).Returns(leaseInvite)

	assert.NotPanics(t, func() {
		database.DeleteCurrentLeaseInvite("1")
	})
}

func TestDeleteCurrentLeaseInvite_NotFound(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.LeaseInvite.Expect(database.MockDeleteCurrentLeaseInvite(c)).Errors(db.ErrNotFound)

	assert.Panics(t, func() {
		database.DeleteCurrentLeaseInvite("1")
	})
}

func TestDeleteCurrentLeaseInvite_NoConnection(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.LeaseInvite.Expect(database.MockDeleteCurrentLeaseInvite(c)).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.DeleteCurrentLeaseInvite("1")
	})
}

// #############################################################################

func TestGetLeaseByID(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	lease := BuildTestLease()
	m.Lease.Expect(database.MockGetLeaseByID(c)).Returns(lease)

	foundLease := database.GetLeaseByID(lease.ID)
	assert.NotNil(t, foundLease)
	assert.Equal(t, lease.ID, foundLease.ID)
	assert.Equal(t, lease.TenantID, foundLease.TenantID)
	assert.Equal(t, lease.PropertyID, foundLease.PropertyID)
}

func TestGetLeaseByID_NotFound(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.Lease.Expect(database.MockGetLeaseByID(c)).Errors(db.ErrNotFound)

	foundLease := database.GetLeaseByID("1")
	assert.Nil(t, foundLease)
}

func TestGetLeaseByID_NoConnection(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.Lease.Expect(database.MockGetLeaseByID(c)).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.GetLeaseByID("1")
	})
}

// #############################################################################

func TestGetLeasesByProperty(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	lease1 := BuildTestLease()
	lease2 := BuildTestLease()
	m.Lease.Expect(database.MockGetLeasesByProperty(c)).ReturnsMany([]db.LeaseModel{lease1, lease2})

	leases := database.GetLeasesByProperty("1")
	assert.NotNil(t, leases)
	assert.Len(t, leases, 2)
	assert.Equal(t, lease1.ID, leases[0].ID)
	assert.Equal(t, lease2.ID, leases[1].ID)
}

func TestGetLeasesByProperty_NotFound(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.Lease.Expect(database.MockGetLeasesByProperty(c)).Errors(db.ErrNotFound)

	leases := database.GetLeasesByProperty("1")
	assert.Nil(t, leases)
}

func TestGetLeasesByProperty_NoConnection(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.Lease.Expect(database.MockGetLeasesByProperty(c)).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.GetLeasesByProperty("1")
	})
}

// #############################################################################

func TestGetLeasesByTenant(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	tenant := BuildTestTenant("1")
	lease1 := BuildTestLease()
	lease2 := BuildTestLease()
	m.Lease.Expect(database.MockGetLeasesByTenant(c)).ReturnsMany([]db.LeaseModel{lease1, lease2})

	leases := database.GetLeasesByTenant(tenant.ID)
	assert.NotNil(t, leases)
	assert.Len(t, leases, 2)
	assert.Equal(t, lease1.ID, leases[0].ID)
	assert.Equal(t, lease2.ID, leases[1].ID)
}

func TestGetLeasesByTenant_NotFound(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	tenant := BuildTestTenant("1")

	m.Lease.Expect(database.MockGetLeasesByTenant(c)).Errors(db.ErrNotFound)

	leases := database.GetLeasesByTenant(tenant.ID)
	assert.Nil(t, leases)
}

func TestGetLeasesByTenant_NoConnection(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	tenant := BuildTestTenant("1")

	m.Lease.Expect(database.MockGetLeasesByTenant(c)).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.GetLeasesByTenant(tenant.ID)
	})
}
