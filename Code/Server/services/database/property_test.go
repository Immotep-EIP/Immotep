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
			Damages:   []db.DamageModel{{}},
			Contracts: []db.ContractModel{{}},
		},
	}
}

func TestGetAllProperties(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")

	mock.Property.Expect(
		client.Client.Property.FindMany(
			db.Property.OwnerID.Equals("1"),
			db.Property.Archived.Equals(false),
		).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
			db.Property.PendingContract.Fetch(),
		),
	).ReturnsMany([]db.PropertyModel{property})

	allProperties := database.GetAllPropertyByOwnerId("1", false)
	assert.Len(t, allProperties, 1)
	assert.Equal(t, property.ID, allProperties[0].ID)
}

func TestGetAllProperties_MultipleProperties(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	p1 := BuildTestProperty("1")
	p2 := BuildTestProperty("2")

	mock.Property.Expect(
		client.Client.Property.FindMany(
			db.Property.OwnerID.Equals("1"),
			db.Property.Archived.Equals(false),
		).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
			db.Property.PendingContract.Fetch(),
		),
	).ReturnsMany([]db.PropertyModel{p1, p2})

	allProperties := database.GetAllPropertyByOwnerId("1", false)
	assert.Len(t, allProperties, 2)
	assert.Equal(t, p1.ID, allProperties[0].ID)
	assert.Equal(t, p2.ID, allProperties[1].ID)
}

func TestGetAllProperties_NoProperties(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	mock.Property.Expect(
		client.Client.Property.FindMany(
			db.Property.OwnerID.Equals("1"),
			db.Property.Archived.Equals(false),
		).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
			db.Property.PendingContract.Fetch(),
		),
	).ReturnsMany([]db.PropertyModel{})

	allProperties := database.GetAllPropertyByOwnerId("1", false)
	assert.Empty(t, allProperties)
}

func TestGetAllProperties_NoConnection(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	mock.Property.Expect(
		client.Client.Property.FindMany(
			db.Property.OwnerID.Equals("1"),
			db.Property.Archived.Equals(false),
		).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
			db.Property.PendingContract.Fetch(),
		),
	).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.GetAllPropertyByOwnerId("1", false)
	})
}

func TestGetPropertyByID(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")

	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals("1")).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
			db.Property.PendingContract.Fetch(),
		),
	).Returns(property)

	foundProperty := database.GetPropertyByID("1")
	assert.NotNil(t, foundProperty)
	assert.Equal(t, property.ID, foundProperty.ID)
}

func TestGetPropertyByID_NotFound(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals("1")).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
			db.Property.PendingContract.Fetch(),
		),
	).Errors(db.ErrNotFound)

	foundProperty := database.GetPropertyByID("1")
	assert.Nil(t, foundProperty)
}

func TestGetPropertyByID_NoConnection(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals("1")).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
			db.Property.PendingContract.Fetch(),
		),
	).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.GetPropertyByID("1")
	})
}

func TestGetPropertyInventory(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")

	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals("1")).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
			db.Property.PendingContract.Fetch(),
			db.Property.Rooms.Fetch().With(db.Room.Furnitures.Fetch()),
		),
	).Returns(property)

	foundProperty := database.GetPropertyInventory("1")
	assert.NotNil(t, foundProperty)
	assert.Equal(t, property.ID, foundProperty.ID)
}

func TestGetPropertyInventory_NotFound(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals("1")).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
			db.Property.PendingContract.Fetch(),
			db.Property.Rooms.Fetch().With(db.Room.Furnitures.Fetch()),
		),
	).Errors(db.ErrNotFound)

	foundProperty := database.GetPropertyInventory("1")
	assert.Nil(t, foundProperty)
}

func TestGetPropertyInventory_NoConnection(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals("1")).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
			db.Property.PendingContract.Fetch(),
			db.Property.Rooms.Fetch().With(db.Room.Furnitures.Fetch()),
		),
	).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.GetPropertyInventory("1")
	})
}

func TestCreateProperty(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")

	mock.Property.Expect(
		client.Client.Property.CreateOne(
			db.Property.Name.Set(property.Name),
			db.Property.Address.Set(property.Address),
			db.Property.City.Set(property.City),
			db.Property.PostalCode.Set(property.PostalCode),
			db.Property.Country.Set(property.Country),
			db.Property.AreaSqm.Set(property.AreaSqm),
			db.Property.RentalPricePerMonth.Set(property.RentalPricePerMonth),
			db.Property.DepositPrice.Set(property.DepositPrice),
			db.Property.Owner.Link(db.User.ID.Equals("1")),
			db.Property.ApartmentNumber.SetIfPresent(property.InnerProperty.ApartmentNumber),
		).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch(),
			db.Property.PendingContract.Fetch(),
		),
	).Returns(property)

	newProperty := database.CreateProperty(property, "1")
	assert.NotNil(t, newProperty)
	assert.Equal(t, property.ID, newProperty.ID)
}

func TestCreateProperty_AlreadyExists(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")

	mock.Property.Expect(
		client.Client.Property.CreateOne(
			db.Property.Name.Set(property.Name),
			db.Property.Address.Set(property.Address),
			db.Property.City.Set(property.City),
			db.Property.PostalCode.Set(property.PostalCode),
			db.Property.Country.Set(property.Country),
			db.Property.AreaSqm.Set(property.AreaSqm),
			db.Property.RentalPricePerMonth.Set(property.RentalPricePerMonth),
			db.Property.DepositPrice.Set(property.DepositPrice),
			db.Property.Owner.Link(db.User.ID.Equals("1")),
			db.Property.ApartmentNumber.SetIfPresent(property.InnerProperty.ApartmentNumber),
		).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch(),
			db.Property.PendingContract.Fetch(),
		),
	).Errors(&protocol.UserFacingError{
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
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")

	mock.Property.Expect(
		client.Client.Property.CreateOne(
			db.Property.Name.Set(property.Name),
			db.Property.Address.Set(property.Address),
			db.Property.City.Set(property.City),
			db.Property.PostalCode.Set(property.PostalCode),
			db.Property.Country.Set(property.Country),
			db.Property.AreaSqm.Set(property.AreaSqm),
			db.Property.RentalPricePerMonth.Set(property.RentalPricePerMonth),
			db.Property.DepositPrice.Set(property.DepositPrice),
			db.Property.Owner.Link(db.User.ID.Equals("1")),
			db.Property.ApartmentNumber.SetIfPresent(property.InnerProperty.ApartmentNumber),
		).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch(),
			db.Property.PendingContract.Fetch(),
		),
	).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.CreateProperty(property, "1")
	})
}

func TestUpdatePropertyPicture(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	image := BuildTestImage("1", "b3Vp")

	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals(property.ID)).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
			db.Property.PendingContract.Fetch(),
		).Update(
			db.Property.Picture.Link(db.Image.ID.Equals(image.ID)),
		),
	).Returns(property)

	updatedProperty := database.UpdatePropertyPicture(property, image)
	assert.NotNil(t, updatedProperty)
	assert.Equal(t, property.ID, updatedProperty.ID)
}

func TestUpdatePropertyPicture_NotFound(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	image := BuildTestImage("1", "b3Vp")

	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals(property.ID)).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
			db.Property.PendingContract.Fetch(),
		).Update(
			db.Property.Picture.Link(db.Image.ID.Equals(image.ID)),
		),
	).Errors(db.ErrNotFound)

	updatedProperty := database.UpdatePropertyPicture(property, image)
	assert.Nil(t, updatedProperty)
}

func TestUpdatePropertyPicture_NoConnection(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	image := BuildTestImage("1", "b3Vp")

	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals(property.ID)).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
			db.Property.PendingContract.Fetch(),
		).Update(
			db.Property.Picture.Link(db.Image.ID.Equals(image.ID)),
		),
	).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.UpdatePropertyPicture(property, image)
	})
}

func TestArchiveProperty(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	property.Archived = true

	mock.Property.Expect(
		client.Client.Property.FindUnique(
			db.Property.ID.Equals(property.ID),
		).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
			db.Property.PendingContract.Fetch(),
		).Update(
			db.Property.Archived.Set(true),
		),
	).Returns(property)

	archivedProperty := database.ToggleArchiveProperty(property.ID, true)
	assert.NotNil(t, archivedProperty)
	assert.Equal(t, property.ID, archivedProperty.ID)
	assert.True(t, archivedProperty.Archived)
}

func TestArchiveProperty_NotFound(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")

	mock.Property.Expect(
		client.Client.Property.FindUnique(
			db.Property.ID.Equals(property.ID),
		).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
			db.Property.PendingContract.Fetch(),
		).Update(
			db.Property.Archived.Set(true),
		),
	).Errors(db.ErrNotFound)

	archivedProperty := database.ToggleArchiveProperty(property.ID, true)
	assert.Nil(t, archivedProperty)
}

func TestArchiveProperty_NoConnection(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")

	mock.Property.Expect(
		client.Client.Property.FindUnique(
			db.Property.ID.Equals(property.ID),
		).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
			db.Property.PendingContract.Fetch(),
		).Update(
			db.Property.Archived.Set(true),
		),
	).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.ToggleArchiveProperty(property.ID, true)
	})
}

func TestUpdateProperty(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
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

	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals(property.ID)).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
			db.Property.PendingContract.Fetch(),
		).Update(
			db.Property.Name.SetIfPresent(updateRequest.Name),
			db.Property.Address.SetIfPresent(updateRequest.Address),
			db.Property.ApartmentNumber.SetIfPresent(updateRequest.ApartmentNumber),
			db.Property.City.SetIfPresent(updateRequest.City),
			db.Property.PostalCode.SetIfPresent(updateRequest.PostalCode),
			db.Property.Country.SetIfPresent(updateRequest.Country),
			db.Property.AreaSqm.SetIfPresent(updateRequest.AreaSqm),
			db.Property.RentalPricePerMonth.SetIfPresent(updateRequest.RentalPricePerMonth),
			db.Property.DepositPrice.SetIfPresent(updateRequest.DepositPrice),
		),
	).Returns(property)

	updatedProperty := database.UpdateProperty(property.ID, updateRequest)
	assert.NotNil(t, updatedProperty)
	assert.Equal(t, property.ID, updatedProperty.ID)
}

func TestUpdateProperty_NotFound(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
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

	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals("1")).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
			db.Property.PendingContract.Fetch(),
		).Update(
			db.Property.Name.SetIfPresent(updateRequest.Name),
			db.Property.Address.SetIfPresent(updateRequest.Address),
			db.Property.ApartmentNumber.SetIfPresent(updateRequest.ApartmentNumber),
			db.Property.City.SetIfPresent(updateRequest.City),
			db.Property.PostalCode.SetIfPresent(updateRequest.PostalCode),
			db.Property.Country.SetIfPresent(updateRequest.Country),
			db.Property.AreaSqm.SetIfPresent(updateRequest.AreaSqm),
			db.Property.RentalPricePerMonth.SetIfPresent(updateRequest.RentalPricePerMonth),
			db.Property.DepositPrice.SetIfPresent(updateRequest.DepositPrice),
		),
	).Errors(db.ErrNotFound)

	updatedProperty := database.UpdateProperty("1", updateRequest)
	assert.Nil(t, updatedProperty)
}

func TestUpdateProperty_NoConnection(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
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

	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals("1")).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
			db.Property.PendingContract.Fetch(),
		).Update(
			db.Property.Name.SetIfPresent(updateRequest.Name),
			db.Property.Address.SetIfPresent(updateRequest.Address),
			db.Property.ApartmentNumber.SetIfPresent(updateRequest.ApartmentNumber),
			db.Property.City.SetIfPresent(updateRequest.City),
			db.Property.PostalCode.SetIfPresent(updateRequest.PostalCode),
			db.Property.Country.SetIfPresent(updateRequest.Country),
			db.Property.AreaSqm.SetIfPresent(updateRequest.AreaSqm),
			db.Property.RentalPricePerMonth.SetIfPresent(updateRequest.RentalPricePerMonth),
			db.Property.DepositPrice.SetIfPresent(updateRequest.DepositPrice),
		),
	).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.UpdateProperty("1", updateRequest)
	})
}
