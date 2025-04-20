package database_test

import (
	"errors"
	"testing"
	"time"

	"github.com/steebchen/prisma-client-go/engine/protocol"
	"github.com/stretchr/testify/assert"
	"immotep/backend/models"
	"immotep/backend/prisma/db"
	"immotep/backend/services"
	"immotep/backend/services/database"
	"immotep/backend/utils"
)

func BuildTestProperty(id string) db.PropertyModel {
	return db.PropertyModel{
		InnerProperty: db.InnerProperty{
			ID:                  id,
			Name:                "Test",
			Address:             "Test",
			ApartmentNumber:     utils.Ptr("Test"),
			City:                "Test",
			PostalCode:          "Test",
			Country:             "Test",
			AreaSqm:             20.0,
			RentalPricePerMonth: 500,
			DepositPrice:        1000,
			CreatedAt:           time.Now(),
			OwnerID:             "1",
			PictureID:           utils.Ptr("1"),
			Archived:            false,
		},
		RelationsProperty: db.RelationsProperty{
			Leases: []db.LeaseModel{{
				InnerLease: db.InnerLease{
					ID:     "1",
					Active: true,
				},
				RelationsLease: db.RelationsLease{
					Tenant: &db.UserModel{
						InnerUser: db.InnerUser{
							Firstname: "Test",
							Lastname:  "Test",
						},
					},
					Damages: []db.DamageModel{{
						InnerDamage: db.InnerDamage{
							ID:      "1",
							FixedAt: nil,
						}},
					},
				},
			}},
			LeaseInvite: &db.LeaseInviteModel{},
		},
	}
}

// #############################################################################

func TestGetAllProperties(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	m.Property.Expect(database.MockGetAllPropertyByOwnerId(c, false)).ReturnsMany([]db.PropertyModel{property})

	allProperties := database.GetPropertiesByOwnerId("1", false)
	assert.Len(t, allProperties, 1)
	assert.Equal(t, property.ID, allProperties[0].ID)
}

func TestGetAllProperties_MultipleProperties(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	p1 := BuildTestProperty("1")
	p2 := BuildTestProperty("2")
	m.Property.Expect(database.MockGetAllPropertyByOwnerId(c, false)).ReturnsMany([]db.PropertyModel{p1, p2})

	allProperties := database.GetPropertiesByOwnerId("1", false)
	assert.Len(t, allProperties, 2)
	assert.Equal(t, p1.ID, allProperties[0].ID)
	assert.Equal(t, p2.ID, allProperties[1].ID)
}

func TestGetAllProperties_NoProperties(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.Property.Expect(database.MockGetAllPropertyByOwnerId(c, false)).ReturnsMany([]db.PropertyModel{})

	allProperties := database.GetPropertiesByOwnerId("1", false)
	assert.Empty(t, allProperties)
}

func TestGetAllProperties_NoConnection(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.Property.Expect(database.MockGetAllPropertyByOwnerId(c, false)).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.GetPropertiesByOwnerId("1", false)
	})
}

// #############################################################################

func TestGetPropertyByID(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	m.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)

	foundProperty := database.GetPropertyByID("1")
	assert.NotNil(t, foundProperty)
	assert.Equal(t, property.ID, foundProperty.ID)
}

func TestGetPropertyByID_NotFound(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.Property.Expect(database.MockGetPropertyByID(c)).Errors(db.ErrNotFound)

	foundProperty := database.GetPropertyByID("1")
	assert.Nil(t, foundProperty)
}

func TestGetPropertyByID_NoConnection(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.Property.Expect(database.MockGetPropertyByID(c)).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.GetPropertyByID("1")
	})
}

// #############################################################################

func TestGetPropertyInventory(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	m.Property.Expect(database.MockGetPropertyInventory(c)).Returns(property)

	foundProperty := database.GetPropertyInventory("1")
	assert.NotNil(t, foundProperty)
	assert.Equal(t, property.ID, foundProperty.ID)
}

func TestGetPropertyInventory_NotFound(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.Property.Expect(database.MockGetPropertyInventory(c)).Errors(db.ErrNotFound)

	foundProperty := database.GetPropertyInventory("1")
	assert.Nil(t, foundProperty)
}

func TestGetPropertyInventory_NoConnection(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.Property.Expect(database.MockGetPropertyInventory(c)).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.GetPropertyInventory("1")
	})
}

// #############################################################################

func TestCreateProperty(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	m.Property.Expect(database.MockCreateProperty(c, property)).Returns(property)

	newProperty := database.CreateProperty(property, "1")
	assert.NotNil(t, newProperty)
	assert.Equal(t, property.ID, newProperty.ID)
}

func TestCreateProperty_AlreadyExists(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	m.Property.Expect(database.MockCreateProperty(c, property)).Errors(&protocol.UserFacingError{
		IsPanic:   false,
		ErrorCode: "P2002", // https://www.prisma.io/docs/orm/reference/error-reference
		Meta: protocol.Meta{
			Target: []any{"email"},
		},
		Message: "Unique constraint failed",
	})

	newProperty := database.CreateProperty(property, "1")
	assert.Nil(t, newProperty)
}

func TestCreateProperty_NoConnection(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	m.Property.Expect(database.MockCreateProperty(c, property)).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.CreateProperty(property, "1")
	})
}

// #############################################################################

func TestUpdatePropertyPicture(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	image := BuildTestImage("1", "b3Vp")
	m.Property.Expect(database.MockUpdatePropertyPicture(c)).Returns(property)

	updatedProperty := database.UpdatePropertyPicture(property, image)
	assert.NotNil(t, updatedProperty)
	assert.Equal(t, property.ID, updatedProperty.ID)
}

func TestUpdatePropertyPicture_NotFound(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	image := BuildTestImage("1", "b3Vp")
	m.Property.Expect(database.MockUpdatePropertyPicture(c)).Errors(db.ErrNotFound)

	updatedProperty := database.UpdatePropertyPicture(property, image)
	assert.Nil(t, updatedProperty)
}

func TestUpdatePropertyPicture_NoConnection(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	image := BuildTestImage("1", "b3Vp")
	m.Property.Expect(database.MockUpdatePropertyPicture(c)).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.UpdatePropertyPicture(property, image)
	})
}

// #############################################################################

func TestArchiveProperty(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	property.Archived = true
	m.Property.Expect(database.MockArchiveProperty(c)).Returns(property)

	archivedProperty := database.ArchiveProperty(property.ID, true)
	assert.NotNil(t, archivedProperty)
	assert.Equal(t, property.ID, archivedProperty.ID)
	assert.True(t, archivedProperty.Archived)
}

func TestArchiveProperty_NotFound(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	m.Property.Expect(database.MockArchiveProperty(c)).Errors(db.ErrNotFound)

	archivedProperty := database.ArchiveProperty(property.ID, true)
	assert.Nil(t, archivedProperty)
}

func TestArchiveProperty_NoConnection(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	m.Property.Expect(database.MockArchiveProperty(c)).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.ArchiveProperty(property.ID, true)
	})
}

// #############################################################################

func TestUpdateProperty(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	updateRequest := models.PropertyUpdateRequest{
		Name:                utils.Ptr("Updated Name"),
		Address:             utils.Ptr("Updated Address"),
		ApartmentNumber:     utils.Ptr("Updated Apartment Number"),
		City:                utils.Ptr("Updated City"),
		PostalCode:          utils.Ptr("Updated Postal Code"),
		Country:             utils.Ptr("Updated Country"),
		AreaSqm:             utils.Ptr(30.0),
		RentalPricePerMonth: utils.Ptr(600.0),
		DepositPrice:        utils.Ptr(1200.0),
	}
	m.Property.Expect(database.MockUpdateProperty(c, updateRequest)).Returns(property)

	updatedProperty := database.UpdateProperty(property, updateRequest)
	assert.NotNil(t, updatedProperty)
	assert.Equal(t, property.ID, updatedProperty.ID)
}

func TestUpdateProperty_AlreadyExists(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	updateRequest := models.PropertyUpdateRequest{
		Name:                utils.Ptr("Updated Name"),
		Address:             utils.Ptr("Updated Address"),
		ApartmentNumber:     utils.Ptr("Updated Apartment Number"),
		City:                utils.Ptr("Updated City"),
		PostalCode:          utils.Ptr("Updated Postal Code"),
		Country:             utils.Ptr("Updated Country"),
		AreaSqm:             utils.Ptr(30.0),
		RentalPricePerMonth: utils.Ptr(600.0),
		DepositPrice:        utils.Ptr(1200.0),
	}
	m.Property.Expect(database.MockUpdateProperty(c, updateRequest)).Errors(&protocol.UserFacingError{
		IsPanic:   false,
		ErrorCode: "P2002", // https://www.prisma.io/docs/orm/reference/error-reference#p2002
		Meta: protocol.Meta{
			Target: []any{"owner_id", "name"},
		},
		Message: "Unique constraint failed",
	})

	updatedProperty := database.UpdateProperty(BuildTestProperty("1"), updateRequest)
	assert.Nil(t, updatedProperty)
}

func TestUpdateProperty_NotFound(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	updateRequest := models.PropertyUpdateRequest{
		Name:                utils.Ptr("Updated Name"),
		Address:             utils.Ptr("Updated Address"),
		ApartmentNumber:     utils.Ptr("Updated Apartment Number"),
		City:                utils.Ptr("Updated City"),
		PostalCode:          utils.Ptr("Updated Postal Code"),
		Country:             utils.Ptr("Updated Country"),
		AreaSqm:             utils.Ptr(30.0),
		RentalPricePerMonth: utils.Ptr(600.0),
		DepositPrice:        utils.Ptr(1200.0),
	}
	m.Property.Expect(database.MockUpdateProperty(c, updateRequest)).Errors(db.ErrNotFound)

	assert.Panics(t, func() {
		database.UpdateProperty(BuildTestProperty("1"), updateRequest)
	})
}

func TestUpdateProperty_NoConnection(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	updateRequest := models.PropertyUpdateRequest{
		Name:                utils.Ptr("Updated Name"),
		Address:             utils.Ptr("Updated Address"),
		ApartmentNumber:     utils.Ptr("Updated Apartment Number"),
		City:                utils.Ptr("Updated City"),
		PostalCode:          utils.Ptr("Updated Postal Code"),
		Country:             utils.Ptr("Updated Country"),
		AreaSqm:             utils.Ptr(30.0),
		RentalPricePerMonth: utils.Ptr(600.0),
		DepositPrice:        utils.Ptr(1200.0),
	}
	m.Property.Expect(database.MockUpdateProperty(c, updateRequest)).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.UpdateProperty(BuildTestProperty("1"), updateRequest)
	})
}
