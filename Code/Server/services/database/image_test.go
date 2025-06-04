package database_test

import (
	"errors"
	"testing"

	"github.com/steebchen/prisma-client-go/engine/protocol"
	"github.com/stretchr/testify/assert"
	"immotep/backend/models"
	"immotep/backend/prisma/db"
	"immotep/backend/services"
	"immotep/backend/services/database"
)

func BuildTestImage(id string, base64data string) db.ImageModel {
	ret := models.StringToDbImage(base64data)
	if ret == nil {
		panic("Invalid base64 string")
	}
	ret.ID = id
	return *ret
}

// #############################################################################

func TestGetImageByID(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	image := BuildTestImage("1", "data:image/jpeg;base64,b3Vp")
	m.Image.Expect(database.MockGetImageByID(c)).Returns(image)

	foundImage := database.GetImageByID("1")
	assert.NotNil(t, foundImage)
	assert.Equal(t, image.ID, foundImage.ID)
}

func TestGetImageByID_NotFound(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.Image.Expect(database.MockGetImageByID(c)).Errors(db.ErrNotFound)

	foundImage := database.GetImageByID("1")
	assert.Nil(t, foundImage)
}

func TestGetImageByID_NoConnection(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.Image.Expect(database.MockGetImageByID(c)).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.GetImageByID("1")
	})
}

// #############################################################################

func TestCreateImage(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	image := BuildTestImage("1", "data:image/jpeg;base64,b3Vp")
	m.Image.Expect(database.MockCreateImage(c, image)).Returns(image)

	newImage := database.CreateImage(image)
	assert.NotNil(t, newImage)
	assert.Equal(t, image.ID, newImage.ID)
}

func TestCreateImage_AlreadyExists(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	image := BuildTestImage("1", "data:image/jpeg;base64,b3Vp")
	m.Image.Expect(database.MockCreateImage(c, image)).Errors(&protocol.UserFacingError{
		IsPanic:   false,
		ErrorCode: "P2002", // https://www.prisma.io/docs/orm/reference/error-reference
		Meta: protocol.Meta{
			Target: []any{"id"},
		},
		Message: "Unique constraint failed",
	})

	assert.Panics(t, func() {
		database.CreateImage(image)
	})
}

func TestCreateImage_NoConnection(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	image := BuildTestImage("1", "data:image/jpeg;base64,b3Vp")
	m.Image.Expect(database.MockCreateImage(c, image)).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.CreateImage(image)
	})
}
