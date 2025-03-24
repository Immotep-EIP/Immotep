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
	leaseInvite := BuildTestLeaseInvite()
	lease := BuildTestLease()

	mock.Lease.Expect(
		client.Client.Lease.CreateOne(
			db.Lease.StartDate.Set(leaseInvite.StartDate),
			db.Lease.Tenant.Link(db.User.ID.Equals(tenant.ID)),
			db.Lease.Property.Link(db.Property.ID.Equals(leaseInvite.PropertyID)),
			db.Lease.EndDate.SetIfPresent(leaseInvite.InnerLeaseInvite.EndDate),
		),
	).Returns(lease)

	mock.LeaseInvite.Expect(
		client.Client.LeaseInvite.FindUnique(
			db.LeaseInvite.ID.Equals(leaseInvite.ID),
		).Delete(),
	).Returns(leaseInvite)

	newLease := database.CreateLease(leaseInvite, tenant)
	assert.NotNil(t, newLease)
	assert.Equal(t, lease.TenantID, newLease.TenantID)
	assert.Equal(t, lease.PropertyID, newLease.PropertyID)
}

func TestCreateLease_NoConnection1(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	tenant := BuildTestTenant("1")
	leaseInvite := BuildTestLeaseInvite()

	mock.Lease.Expect(
		client.Client.Lease.CreateOne(
			db.Lease.StartDate.Set(leaseInvite.StartDate),
			db.Lease.Tenant.Link(db.User.ID.Equals(tenant.ID)),
			db.Lease.Property.Link(db.Property.ID.Equals(leaseInvite.PropertyID)),
			db.Lease.EndDate.SetIfPresent(leaseInvite.InnerLeaseInvite.EndDate),
		),
	).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.CreateLease(leaseInvite, tenant)
	})
}

func TestCreateLease_NoConnection2(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	tenant := BuildTestTenant("1")
	leaseInvite := BuildTestLeaseInvite()
	lease := BuildTestLease()

	mock.Lease.Expect(
		client.Client.Lease.CreateOne(
			db.Lease.StartDate.Set(leaseInvite.StartDate),
			db.Lease.Tenant.Link(db.User.ID.Equals(tenant.ID)),
			db.Lease.Property.Link(db.Property.ID.Equals(leaseInvite.PropertyID)),
			db.Lease.EndDate.SetIfPresent(leaseInvite.InnerLeaseInvite.EndDate),
		),
	).Returns(lease)

	mock.LeaseInvite.Expect(
		client.Client.LeaseInvite.FindUnique(
			db.LeaseInvite.ID.Equals(leaseInvite.ID),
		).Delete(),
	).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.CreateLease(leaseInvite, tenant)
	})
}

func TestGetLeaseInviteByID(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	user := BuildTestLeaseInvite()

	mock.LeaseInvite.Expect(
		client.Client.LeaseInvite.FindUnique(db.LeaseInvite.ID.Equals("1")),
	).Returns(user)

	foundPending := database.GetLeaseInviteById("1")
	assert.NotNil(t, foundPending)
	assert.Equal(t, user.ID, foundPending.ID)
}

func TestGetLeaseInviteByID_NotFound(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	mock.LeaseInvite.Expect(
		client.Client.LeaseInvite.FindUnique(db.LeaseInvite.ID.Equals("1")),
	).Errors(db.ErrNotFound)

	foundPending := database.GetLeaseInviteById("1")
	assert.Nil(t, foundPending)
}

func TestGetLeaseInviteByID_NoConnection(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	mock.LeaseInvite.Expect(
		client.Client.LeaseInvite.FindUnique(db.LeaseInvite.ID.Equals("1")),
	).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.GetLeaseInviteById("1")
	})
}

func TestCreateLeaseInvite(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	leaseInvite := BuildTestLeaseInvite()
	mock.LeaseInvite.Expect(
		client.Client.LeaseInvite.CreateOne(
			db.LeaseInvite.TenantEmail.Set(leaseInvite.TenantEmail),
			db.LeaseInvite.StartDate.Set(leaseInvite.StartDate),
			db.LeaseInvite.Property.Link(db.Property.ID.Equals("1")),
			db.LeaseInvite.EndDate.SetIfPresent(leaseInvite.InnerLeaseInvite.EndDate),
		).With(
			db.LeaseInvite.Property.Fetch().With(db.Property.Owner.Fetch()),
		),
	).Returns(leaseInvite)

	newLease := database.CreateLeaseInvite(leaseInvite, property.ID)
	assert.NotNil(t, newLease)
	assert.Equal(t, leaseInvite.ID, newLease.ID)
}

func TestCreateLeaseInvite_AlreadyExists1(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	leaseInvite := BuildTestLeaseInvite()
	mock.LeaseInvite.Expect(
		client.Client.LeaseInvite.CreateOne(
			db.LeaseInvite.TenantEmail.Set(leaseInvite.TenantEmail),
			db.LeaseInvite.StartDate.Set(leaseInvite.StartDate),
			db.LeaseInvite.Property.Link(db.Property.ID.Equals("1")),
			db.LeaseInvite.EndDate.SetIfPresent(leaseInvite.InnerLeaseInvite.EndDate),
		).With(
			db.LeaseInvite.Property.Fetch().With(db.Property.Owner.Fetch()),
		),
	).Errors(&protocol.UserFacingError{
		IsPanic:   false,
		ErrorCode: "P2002", // https://www.prisma.io/docs/orm/reference/error-reference
		Meta: protocol.Meta{
			Target: []any{"tenant_email"},
		},
		Message: "Unique constraint failed",
	})

	newLease := database.CreateLeaseInvite(leaseInvite, property.ID)
	assert.Nil(t, newLease)
}

func TestCreateLeaseInvite_AlreadyExists2(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	leaseInvite := BuildTestLeaseInvite()
	mock.LeaseInvite.Expect(
		client.Client.LeaseInvite.CreateOne(
			db.LeaseInvite.TenantEmail.Set(leaseInvite.TenantEmail),
			db.LeaseInvite.StartDate.Set(leaseInvite.StartDate),
			db.LeaseInvite.Property.Link(db.Property.ID.Equals("1")),
			db.LeaseInvite.EndDate.SetIfPresent(leaseInvite.InnerLeaseInvite.EndDate),
		).With(
			db.LeaseInvite.Property.Fetch().With(db.Property.Owner.Fetch()),
		),
	).Errors(&protocol.UserFacingError{
		IsPanic:   false,
		ErrorCode: "P2014", // https://www.prisma.io/docs/orm/reference/error-reference
		Meta: protocol.Meta{
			Target: []any{"property_id"},
		},
		Message: "Unique constraint failed",
	})

	newLease := database.CreateLeaseInvite(leaseInvite, property.ID)
	assert.Nil(t, newLease)
}

func TestCreateLeaseInvite_NoConnection(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	leaseInvite := BuildTestLeaseInvite()
	mock.LeaseInvite.Expect(
		client.Client.LeaseInvite.CreateOne(
			db.LeaseInvite.TenantEmail.Set(leaseInvite.TenantEmail),
			db.LeaseInvite.StartDate.Set(leaseInvite.StartDate),
			db.LeaseInvite.Property.Link(db.Property.ID.Equals("1")),
			db.LeaseInvite.EndDate.SetIfPresent(leaseInvite.InnerLeaseInvite.EndDate),
		).With(
			db.LeaseInvite.Property.Fetch().With(db.Property.Owner.Fetch()),
		),
	).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.CreateLeaseInvite(leaseInvite, property.ID)
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

func TestGetCurrentLeaseInvite(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	leaseInvite := BuildTestLeaseInvite()

	mock.LeaseInvite.Expect(
		client.Client.LeaseInvite.FindUnique(
			db.LeaseInvite.PropertyID.Equals(property.ID),
		),
	).Returns(leaseInvite)

	foundPending := database.GetCurrentLeaseInvite(property.ID)
	assert.NotNil(t, foundPending)
	assert.Equal(t, leaseInvite.ID, foundPending.ID)
}

func TestGetCurrentLeaseInvite_NotFound(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")

	mock.LeaseInvite.Expect(
		client.Client.LeaseInvite.FindUnique(
			db.LeaseInvite.PropertyID.Equals(property.ID),
		),
	).Errors(db.ErrNotFound)

	foundPending := database.GetCurrentLeaseInvite(property.ID)
	assert.Nil(t, foundPending)
}

func TestGetCurrentLeaseInvite_NoConnection(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")

	mock.LeaseInvite.Expect(
		client.Client.LeaseInvite.FindUnique(
			db.LeaseInvite.PropertyID.Equals(property.ID),
		),
	).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.GetCurrentLeaseInvite(property.ID)
	})
}

func TestDeleteCurrentLeaseInvite(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	leaseInvite := BuildTestLeaseInvite()

	mock.LeaseInvite.Expect(
		client.Client.LeaseInvite.FindUnique(
			db.LeaseInvite.PropertyID.Equals(property.ID),
		).Delete(),
	).Returns(leaseInvite)

	assert.NotPanics(t, func() {
		database.DeleteCurrentLeaseInvite(property.ID)
	})
}

func TestDeleteCurrentLeaseInvite_NotFound(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")

	mock.LeaseInvite.Expect(
		client.Client.LeaseInvite.FindUnique(
			db.LeaseInvite.PropertyID.Equals(property.ID),
		).Delete(),
	).Errors(db.ErrNotFound)

	assert.Panics(t, func() {
		database.DeleteCurrentLeaseInvite(property.ID)
	})
}

func TestDeleteCurrentLeaseInvite_NoConnection(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")

	mock.LeaseInvite.Expect(
		client.Client.LeaseInvite.FindUnique(
			db.LeaseInvite.PropertyID.Equals(property.ID),
		).Delete(),
	).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.DeleteCurrentLeaseInvite(property.ID)
	})
}
