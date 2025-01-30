package propertyservice_test

import (
	"errors"
	"testing"
	"time"

	"github.com/steebchen/prisma-client-go/engine/protocol"
	"github.com/stretchr/testify/assert"
	"immotep/backend/database"
	"immotep/backend/models"
	"immotep/backend/prisma/db"
	propertyservice "immotep/backend/services/property"
	"immotep/backend/utils"
)

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
			PictureID:           utils.Ptr("1"),
		},
		RelationsProperty: db.RelationsProperty{
			Damages:   []db.DamageModel{{}},
			Contracts: []db.ContractModel{{}},
		},
	}
}

func BuildTestImage(id string, base64data string) db.ImageModel {
	ret := models.StringToDbImage(base64data)
	if ret == nil {
		panic("Invalid base64 string")
	}
	ret.ID = id
	return *ret
}

func TestGetAllProperties(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")

	mock.Property.Expect(
		client.Client.Property.FindMany(
			db.Property.OwnerID.Equals("1"),
		).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
		),
	).ReturnsMany([]db.PropertyModel{property})

	allProperties := propertyservice.GetAllByOwnerId("1")
	assert.Len(t, allProperties, 1)
	assert.Equal(t, property.ID, allProperties[0].ID)
}

func TestGetAllProperties_MultipleProperties(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	user1 := BuildTestProperty("1")
	user2 := BuildTestProperty("2")

	mock.Property.Expect(
		client.Client.Property.FindMany(
			db.Property.OwnerID.Equals("1"),
		).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
		),
	).ReturnsMany([]db.PropertyModel{user1, user2})

	allProperties := propertyservice.GetAllByOwnerId("1")
	assert.Len(t, allProperties, 2)
	assert.Equal(t, user1.ID, allProperties[0].ID)
	assert.Equal(t, user2.ID, allProperties[1].ID)
}

func TestGetAllProperties_NoProperties(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	mock.Property.Expect(
		client.Client.Property.FindMany(
			db.Property.OwnerID.Equals("1"),
		).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
		),
	).ReturnsMany([]db.PropertyModel{})

	allProperties := propertyservice.GetAllByOwnerId("1")
	assert.Empty(t, allProperties)
}

func TestGetAllProperties_NoConnection(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	mock.Property.Expect(
		client.Client.Property.FindMany(
			db.Property.OwnerID.Equals("1"),
		).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
		),
	).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		propertyservice.GetAllByOwnerId("1")
	})
}

func TestGetPropertyByID(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")

	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals("1")).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
		),
	).Returns(property)

	foundProperty := propertyservice.GetByID("1")
	assert.NotNil(t, foundProperty)
	assert.Equal(t, property.ID, foundProperty.ID)
}

func TestGetPropertyByID_NotFound(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals("1")).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
		),
	).Errors(db.ErrNotFound)

	foundProperty := propertyservice.GetByID("1")
	assert.Nil(t, foundProperty)
}

func TestGetPropertyByID_NoConnection(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals("1")).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
		),
	).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		propertyservice.GetByID("1")
	})
}

func TestCreateProperty(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
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
		).With(
			db.Property.Contracts.Fetch(),
			db.Property.Damages.Fetch(),
		),
	).Returns(property)

	newProperty := propertyservice.Create(property, "1")
	assert.NotNil(t, newProperty)
	assert.Equal(t, property.ID, newProperty.ID)
}

func TestCreateProperty_AlreadyExists(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
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
		).With(
			db.Property.Contracts.Fetch(),
			db.Property.Damages.Fetch(),
		),
	).Errors(&protocol.UserFacingError{
		IsPanic:   false,
		ErrorCode: "P2002", // https://www.prisma.io/docs/orm/reference/error-reference
		Meta: protocol.Meta{
			Target: []any{"email"},
		},
		Message: "Unique constraint failed",
	})

	newProperty := propertyservice.Create(property, "1")
	assert.Nil(t, newProperty)
}

func TestCreateProperty_NoConnection(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
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
		).With(
			db.Property.Contracts.Fetch(),
			db.Property.Damages.Fetch(),
		),
	).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		propertyservice.Create(property, "1")
	})
}

func TestUpdatePicture(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	image := BuildTestImage("1", "b3Vp")

	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals(property.ID)).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
		).Update(
			db.Property.Picture.Link(db.Image.ID.Equals(image.ID)),
		),
	).Returns(property)

	updatedProperty := propertyservice.UpdatePicture(property, image)
	assert.NotNil(t, updatedProperty)
	assert.Equal(t, property.ID, updatedProperty.ID)
}

func TestUpdatePicture_NotFound(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	image := BuildTestImage("1", "b3Vp")

	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals(property.ID)).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
		).Update(
			db.Property.Picture.Link(db.Image.ID.Equals(image.ID)),
		),
	).Errors(db.ErrNotFound)

	updatedProperty := propertyservice.UpdatePicture(property, image)
	assert.Nil(t, updatedProperty)
}

func TestUpdatePicture_NoConnection(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	image := BuildTestImage("1", "b3Vp")

	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals(property.ID)).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
		).Update(
			db.Property.Picture.Link(db.Image.ID.Equals(image.ID)),
		),
	).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		propertyservice.UpdatePicture(property, image)
	})
}
