package database_test

import (
	"errors"
	"testing"
	"time"

	"github.com/steebchen/prisma-client-go/engine/protocol"
	"github.com/stretchr/testify/assert"
	"keyz/backend/models"
	"keyz/backend/prisma/db"
	"keyz/backend/services"
	"keyz/backend/services/database"
	"keyz/backend/utils"
)

func BuildTestDamage(id string) db.DamageModel {
	return db.DamageModel{
		InnerDamage: db.InnerDamage{
			ID:           id,
			LeaseID:      "1",
			RoomID:       "1",
			Comment:      "Test Damage",
			Priority:     db.PriorityMedium,
			Read:         false,
			FixedOwner:   false,
			FixedTenant:  false,
			FixedAt:      nil,
			FixPlannedAt: nil,
		},
	}
}

func TestCreateDamage(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	damage := BuildTestDamage("1")
	pictures := []string{"1", "2"}
	m.Damage.Expect(database.MockCreateDamage(c, damage, "1", pictures)).Returns(damage)

	newDamage := database.CreateDamage(damage, "1", pictures)
	assert.Equal(t, damage.ID, newDamage.ID)
}

func TestCreateDamage_NoConnection(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	damage := BuildTestDamage("1")
	pictures := []string{"1", "2"}
	m.Damage.Expect(database.MockCreateDamage(c, damage, "1", pictures)).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.CreateDamage(damage, "1", pictures)
	})
}

func TestGetDamagesByPropertyID(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	damage1 := BuildTestDamage("1")
	damage2 := BuildTestDamage("2")
	m.Damage.Expect(database.MockGetDamagesByPropertyID(c, false)).ReturnsMany([]db.DamageModel{damage1, damage2})

	damages := database.GetDamagesByPropertyID("1", false)
	assert.Len(t, damages, 2)
	assert.Equal(t, damage1.ID, damages[0].ID)
	assert.Equal(t, damage2.ID, damages[1].ID)
}

func TestGetDamagesByPropertyID_NoDamages(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.Damage.Expect(database.MockGetDamagesByPropertyID(c, false)).ReturnsMany([]db.DamageModel{})

	damages := database.GetDamagesByPropertyID("1", false)
	assert.Empty(t, damages)
}

func TestGetDamagesByPropertyID_NotFound(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.Damage.Expect(database.MockGetDamagesByPropertyID(c, false)).Errors(db.ErrNotFound)

	damages := database.GetDamagesByPropertyID("1", false)
	assert.Nil(t, damages)
}

func TestGetDamagesByPropertyID_NoConnection(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.Damage.Expect(database.MockGetDamagesByPropertyID(c, false)).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.GetDamagesByPropertyID("1", false)
	})
}

func TestGetDamagesByLeaseID(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	damage1 := BuildTestDamage("1")
	damage2 := BuildTestDamage("2")
	m.Damage.Expect(database.MockGetDamagesByLeaseID(c, false)).ReturnsMany([]db.DamageModel{damage1, damage2})

	damages := database.GetDamagesByLeaseID("1", false)
	assert.Len(t, damages, 2)
	assert.Equal(t, damage1.ID, damages[0].ID)
	assert.Equal(t, damage2.ID, damages[1].ID)
}

func TestGetDamagesByLeaseID_NoDamages(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.Damage.Expect(database.MockGetDamagesByLeaseID(c, false)).ReturnsMany([]db.DamageModel{})

	damages := database.GetDamagesByLeaseID("1", false)
	assert.Empty(t, damages)
}

func TestGetDamagesByLeaseID_NotFound(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.Damage.Expect(database.MockGetDamagesByLeaseID(c, false)).Errors(db.ErrNotFound)

	damages := database.GetDamagesByLeaseID("1", false)
	assert.Nil(t, damages)
}

func TestGetDamagesByLeaseID_NoConnection(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.Damage.Expect(database.MockGetDamagesByLeaseID(c, false)).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.GetDamagesByLeaseID("1", false)
	})
}

func TestGetDamageByID(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	damage := BuildTestDamage("1")
	m.Damage.Expect(database.MockGetDamageByID(c)).Returns(damage)

	foundDamage := database.GetDamageByID("1")
	assert.NotNil(t, foundDamage)
	assert.Equal(t, damage.ID, foundDamage.ID)
}

func TestGetDamageByID_NotFound(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.Damage.Expect(database.MockGetDamageByID(c)).Errors(db.ErrNotFound)

	foundDamage := database.GetDamageByID("1")
	assert.Nil(t, foundDamage)
}

func TestGetDamageByID_NoConnection(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.Damage.Expect(database.MockGetDamageByID(c)).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.GetDamageByID("1")
	})
}

func TestUpdateDamageTenant(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	damage := BuildTestDamage("1")
	req := models.DamageTenantUpdateRequest{
		Comment:  utils.Ptr("Updated Comment"),
		Priority: utils.Ptr(db.PriorityHigh),
	}
	pictures := []string{"1", "2"}
	updatedDamage := damage
	updatedDamage.Comment = *req.Comment
	updatedDamage.Priority = *req.Priority

	m.Damage.Expect(database.MockUpdateDamageTenant(c, req, pictures)).Returns(updatedDamage)

	result := database.UpdateDamageTenant(damage, req, pictures)
	assert.NotNil(t, result)
	assert.Equal(t, updatedDamage.Comment, result.Comment)
	assert.Equal(t, updatedDamage.Priority, result.Priority)
}

func TestUpdateDamageTenant_NoPictures(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	damage := BuildTestDamage("1")
	req := models.DamageTenantUpdateRequest{
		Comment:  utils.Ptr("Updated Comment"),
		Priority: utils.Ptr(db.PriorityHigh),
	}
	updatedDamage := damage
	updatedDamage.Comment = *req.Comment
	updatedDamage.Priority = *req.Priority

	m.Damage.Expect(database.MockUpdateDamageTenant(c, req, nil)).Returns(updatedDamage)

	result := database.UpdateDamageTenant(damage, req, nil)
	assert.NotNil(t, result)
	assert.Equal(t, updatedDamage.Comment, result.Comment)
	assert.Equal(t, updatedDamage.Priority, result.Priority)
}

func TestUpdateDamageTenant_UniqueConstraintError(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	damage := BuildTestDamage("1")
	req := models.DamageTenantUpdateRequest{
		Comment:  utils.Ptr("Updated Comment"),
		Priority: utils.Ptr(db.PriorityHigh),
	}
	pictures := []string{"1", "2"}

	m.Damage.Expect(database.MockUpdateDamageTenant(c, req, pictures)).Errors(&protocol.UserFacingError{
		IsPanic:   false,
		ErrorCode: "P2002", // https://www.prisma.io/docs/orm/reference/error-reference
		Meta: protocol.Meta{
			Target: []any{},
		},
		Message: "Unique constraint failed",
	})

	result := database.UpdateDamageTenant(damage, req, pictures)
	assert.Nil(t, result)
}

func TestUpdateDamageTenant_NotFound(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	damage := BuildTestDamage("1")
	req := models.DamageTenantUpdateRequest{
		Comment:  utils.Ptr("Updated Comment"),
		Priority: utils.Ptr(db.PriorityHigh),
	}
	pictures := []string{"1", "2"}

	m.Damage.Expect(database.MockUpdateDamageTenant(c, req, pictures)).Errors(db.ErrNotFound)

	assert.Panics(t, func() {
		database.UpdateDamageTenant(damage, req, pictures)
	})
}

func TestUpdateDamageTenant_NoConnection(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	damage := BuildTestDamage("1")
	req := models.DamageTenantUpdateRequest{
		Comment:  utils.Ptr("Updated Comment"),
		Priority: utils.Ptr(db.PriorityHigh),
	}
	pictures := []string{"1", "2"}

	m.Damage.Expect(database.MockUpdateDamageTenant(c, req, pictures)).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.UpdateDamageTenant(damage, req, pictures)
	})
}

func TestUpdateDamageOwner(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	damage := BuildTestDamage("1")
	req := models.DamageOwnerUpdateRequest{
		Read:         utils.Ptr(true),
		FixPlannedAt: utils.Ptr(time.Now()),
	}
	updatedDamage := damage
	updatedDamage.Read = *req.Read
	updatedDamage.InnerDamage.FixPlannedAt = req.FixPlannedAt

	m.Damage.Expect(database.MockUpdateDamageOwner(c, req)).Returns(updatedDamage)

	result := database.UpdateDamageOwner(damage, req)
	assert.NotNil(t, result)
	assert.Equal(t, updatedDamage.Read, result.Read)
}

func TestUpdateDamageOwner_UniqueConstraintError(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	damage := BuildTestDamage("1")
	req := models.DamageOwnerUpdateRequest{
		Read:         utils.Ptr(true),
		FixPlannedAt: utils.Ptr(time.Now()),
	}

	m.Damage.Expect(database.MockUpdateDamageOwner(c, req)).Errors(&protocol.UserFacingError{
		IsPanic:   false,
		ErrorCode: "P2002", // Unique constraint error
		Meta: protocol.Meta{
			Target: []any{},
		},
		Message: "Unique constraint failed",
	})

	result := database.UpdateDamageOwner(damage, req)
	assert.Nil(t, result)
}

func TestUpdateDamageOwner_NotFound(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	damage := BuildTestDamage("1")
	req := models.DamageOwnerUpdateRequest{
		Read:         utils.Ptr(true),
		FixPlannedAt: utils.Ptr(time.Now()),
	}

	m.Damage.Expect(database.MockUpdateDamageOwner(c, req)).Errors(db.ErrNotFound)

	assert.Panics(t, func() {
		database.UpdateDamageOwner(damage, req)
	})
}

func TestUpdateDamageOwner_NoConnection(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	damage := BuildTestDamage("1")
	req := models.DamageOwnerUpdateRequest{
		Read:         utils.Ptr(true),
		FixPlannedAt: utils.Ptr(time.Now()),
	}

	m.Damage.Expect(database.MockUpdateDamageOwner(c, req)).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.UpdateDamageOwner(damage, req)
	})
}

func TestMarkDamageAsFixed_TenantRole(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	damage := BuildTestDamage("1")
	damage.FixedTenant = false
	damage.FixedOwner = false

	updatedDamage := damage
	updatedDamage.FixedTenant = true

	m.Damage.Expect(database.MockMarkDamageAsFixed(c, damage, db.RoleTenant)).Returns(updatedDamage)

	result := database.MarkDamageAsFixed(damage, db.RoleTenant)
	assert.NotNil(t, result)
	assert.True(t, result.FixedTenant)
	assert.False(t, result.FixedOwner)
}

func TestMarkDamageAsFixed_OwnerRole(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	damage := BuildTestDamage("1")
	damage.FixedTenant = false
	damage.FixedOwner = false

	updatedDamage := damage
	updatedDamage.FixedOwner = true

	m.Damage.Expect(database.MockMarkDamageAsFixed(c, damage, db.RoleOwner)).Returns(updatedDamage)

	result := database.MarkDamageAsFixed(damage, db.RoleOwner)
	assert.NotNil(t, result)
	assert.False(t, result.FixedTenant)
	assert.True(t, result.FixedOwner)
}

func TestMarkDamageAsFixed_BothRoles(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	damage := BuildTestDamage("1")
	damage.FixedTenant = true
	damage.FixedOwner = false

	updatedDamage := damage
	updatedDamage.FixedOwner = true
	updatedDamage.InnerDamage.FixedAt = utils.Ptr(time.Now().Truncate(time.Minute))

	m.Damage.Expect(database.MockMarkDamageAsFixed(c, damage, db.RoleOwner)).Returns(updatedDamage)

	result := database.MarkDamageAsFixed(damage, db.RoleOwner)
	assert.NotNil(t, result)
	assert.True(t, result.FixedTenant)
	assert.True(t, result.FixedOwner)
	assert.NotNil(t, result.FixedAt)
}

func TestMarkDamageAsFixed_NoConnection(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	damage := BuildTestDamage("1")

	m.Damage.Expect(database.MockMarkDamageAsFixed(c, damage, db.RoleTenant)).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.MarkDamageAsFixed(damage, db.RoleTenant)
	})
}
