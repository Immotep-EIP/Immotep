package imageservice_test

import (
	"errors"
	"testing"

	"github.com/steebchen/prisma-client-go/engine/protocol"
	"github.com/stretchr/testify/assert"
	"immotep/backend/database"
	"immotep/backend/models"
	"immotep/backend/prisma/db"
	imageservice "immotep/backend/services/image"
)

func BuildTestImage(id string, base64data string) db.ImageModel {
	ret := models.StringToDbImage(base64data)
	if ret == nil {
		panic("Invalid base64 string")
	}
	ret.ID = id
	return *ret
}

func TestGetImageByID(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	image := BuildTestImage("1", "b3Vp")

	mock.Image.Expect(
		client.Client.Image.FindUnique(db.Image.ID.Equals("1")),
	).Returns(image)

	foundImage := imageservice.GetByID("1")
	assert.NotNil(t, foundImage)
	assert.Equal(t, image.ID, foundImage.ID)
}

func TestGetImageByID_NotFound(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	mock.Image.Expect(
		client.Client.Image.FindUnique(db.Image.ID.Equals("1")),
	).Errors(db.ErrNotFound)

	foundImage := imageservice.GetByID("1")
	assert.Nil(t, foundImage)
}

func TestGetImageByID_NoConnection(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	mock.Image.Expect(
		client.Client.Image.FindUnique(db.Image.ID.Equals("1")),
	).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		imageservice.GetByID("1")
	})
}

func TestCreateImage(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	image := BuildTestImage("1", "b3Vp")

	mock.Image.Expect(
		client.Client.Image.CreateOne(
			db.Image.Data.Set(image.Data),
		),
	).Returns(image)

	newImage := imageservice.Create(image)
	assert.NotNil(t, newImage)
	assert.Equal(t, image.ID, newImage.ID)
}

func TestCreateImage_AlreadyExists(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	image := BuildTestImage("1", "b3Vp")

	mock.Image.Expect(
		client.Client.Image.CreateOne(
			db.Image.Data.Set(image.Data),
		),
	).Errors(&protocol.UserFacingError{
		IsPanic:   false,
		ErrorCode: "P2002", // https://www.prisma.io/docs/orm/reference/error-reference
		Meta: protocol.Meta{
			Target: []any{"id"},
		},
		Message: "Unique constraint failed",
	})

	assert.Panics(t, func() {
		imageservice.Create(image)
	})
}

func TestCreateImage_NoConnection(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	image := BuildTestImage("1", "b3Vp")

	mock.Image.Expect(
		client.Client.Image.CreateOne(
			db.Image.Data.Set(image.Data),
		),
	).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		imageservice.Create(image)
	})
}
