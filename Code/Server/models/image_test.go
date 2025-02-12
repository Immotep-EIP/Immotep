package models_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"immotep/backend/models"
	"immotep/backend/prisma/db"
)

func TestImageRequest(t *testing.T) {
	req := models.ImageRequest{
		Data: "dGVzdCBkYXRh", // base64 for "test data"
	}

	t.Run("ToDbImage", func(t *testing.T) {
		image := req.ToDbImage()

		assert.NotNil(t, image)
		assert.Equal(t, "test data", string(image.Data))
	})

	t.Run("ToDbImageInvalidBase64", func(t *testing.T) {
		invalidReq := models.ImageRequest{
			Data: "invalid base64",
		}
		image := invalidReq.ToDbImage()

		assert.Nil(t, image)
	})
}

func TestStringToDbImage(t *testing.T) {
	t.Run("ValidBase64", func(t *testing.T) {
		data := "dGVzdCBkYXRh" // base64 for "test data"
		image := models.StringToDbImage(data)

		assert.NotNil(t, image)
		assert.Equal(t, "test data", string(image.Data))
	})

	t.Run("InvalidBase64", func(t *testing.T) {
		data := "invalid base64"
		image := models.StringToDbImage(data)

		assert.Nil(t, image)
	})
}

func TestImageResponse(t *testing.T) {
	image := db.ImageModel{
		InnerImage: db.InnerImage{
			ID:        "1",
			Data:      []byte("test data"),
			CreatedAt: time.Now(),
		},
	}

	t.Run("FromDbImage", func(t *testing.T) {
		imageResponse := models.ImageResponse{}
		imageResponse.FromDbImage(image)

		assert.Equal(t, image.ID, imageResponse.ID)
		assert.Equal(t, "dGVzdCBkYXRh", imageResponse.Data) // base64 for "test data"
		assert.Equal(t, image.CreatedAt, imageResponse.CreatedAt)
	})

	t.Run("DbImageToResponse", func(t *testing.T) {
		imageResponse := models.DbImageToResponse(image)

		assert.Equal(t, image.ID, imageResponse.ID)
		assert.Equal(t, "dGVzdCBkYXRh", imageResponse.Data) // base64 for "test data"
		assert.Equal(t, image.CreatedAt, imageResponse.CreatedAt)
	})
}
