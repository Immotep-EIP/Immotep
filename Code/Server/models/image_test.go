package models_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"keyz/backend/models"
	"keyz/backend/prisma/db"
)

func TestImageRequest(t *testing.T) {
	t.Run("ValidJPEG", func(t *testing.T) {
		req := models.ImageRequest{
			Data: "data:image/jpeg;base64,dGVzdCBkYXRh", // base64 for "test data"
		}
		image := req.ToDbImage()

		require.NotNil(t, image)
		assert.Equal(t, "test data", string(image.Data))
		assert.Equal(t, db.ImageTypeJpeg, image.Type)
	})

	t.Run("ValidJPG", func(t *testing.T) {
		req := models.ImageRequest{
			Data: "data:image/jpg;base64,dGVzdCBkYXRh", // base64 for "test data"
		}
		image := req.ToDbImage()

		require.NotNil(t, image)
		assert.Equal(t, "test data", string(image.Data))
		assert.Equal(t, db.ImageTypeJpeg, image.Type)
	})

	t.Run("ValidPNG", func(t *testing.T) {
		req := models.ImageRequest{
			Data: "data:image/png;base64,dGVzdCBkYXRh", // base64 for "test data"
		}
		image := req.ToDbImage()

		require.NotNil(t, image)
		assert.Equal(t, "test data", string(image.Data))
		assert.Equal(t, db.ImageTypePng, image.Type)
	})

	t.Run("InvalidType", func(t *testing.T) {
		req := models.ImageRequest{
			Data: "data:image/gif;base64,dGVzdCBkYXRh", // base64 for "test data"
		}
		image := req.ToDbImage()

		assert.Nil(t, image)
	})

	t.Run("InvalidBase64Data", func(t *testing.T) {
		invalidReq := models.ImageRequest{
			Data: "data:image/jpeg;base64,invalid base64",
		}
		image := invalidReq.ToDbImage()

		assert.Nil(t, image)
	})
}

func TestStringToDbImage(t *testing.T) {
	t.Run("ValidBase64", func(t *testing.T) {
		data := "data:image/jpeg;base64,dGVzdCBkYXRh" // base64 for "test data"
		image := models.StringToDbImage(data)

		require.NotNil(t, image)
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
			Type:      db.ImageTypeJpeg,
			CreatedAt: time.Now(),
		},
	}

	t.Run("FromDbImage_JPEG", func(t *testing.T) {
		nImage := image
		imageResponse := models.ImageResponse{}
		imageResponse.FromDbImage(nImage)

		assert.Equal(t, nImage.ID, imageResponse.ID)
		assert.Equal(t, "data:image/jpeg;base64,dGVzdCBkYXRh", imageResponse.Data) // base64 for "test data"
		assert.Equal(t, nImage.CreatedAt, imageResponse.CreatedAt)
	})

	t.Run("FromDbImage_PNG", func(t *testing.T) {
		nImage := image
		nImage.Type = db.ImageTypePng
		imageResponse := models.ImageResponse{}
		imageResponse.FromDbImage(nImage)

		assert.Equal(t, nImage.ID, imageResponse.ID)
		assert.Equal(t, "data:image/png;base64,dGVzdCBkYXRh", imageResponse.Data) // base64 for "test data"
		assert.Equal(t, nImage.CreatedAt, imageResponse.CreatedAt)
	})

	t.Run("FromDbImage_InvalidType", func(t *testing.T) {
		nImage := image
		nImage.Type = "wrong"
		imageResponse := models.ImageResponse{}

		assert.Panics(t, func() {
			imageResponse.FromDbImage(nImage)
		})
	})

	t.Run("DbImageToResponse", func(t *testing.T) {
		imageResponse := models.DbImageToResponse(image)

		assert.Equal(t, image.ID, imageResponse.ID)
		assert.Equal(t, "data:image/jpeg;base64,dGVzdCBkYXRh", imageResponse.Data) // base64 for "test data"
		assert.Equal(t, image.CreatedAt, imageResponse.CreatedAt)
	})
}
