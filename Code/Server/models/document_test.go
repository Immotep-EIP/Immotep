package models_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"immotep/backend/models"
	"immotep/backend/prisma/db"
)

func TestDocumentResponse(t *testing.T) {
	document := db.DocumentModel{
		InnerDocument: db.InnerDocument{
			ID:        "1",
			Data:      []byte("test data"),
			CreatedAt: time.Now(),
		},
	}

	t.Run("FromDbDocument", func(t *testing.T) {
		documentResponse := models.DocumentResponse{}
		documentResponse.FromDbDocument(document)

		assert.Equal(t, document.ID, documentResponse.ID)
		assert.Equal(t, "data:application/pdf;base64,"+"dGVzdCBkYXRh", documentResponse.Data) // base64 for "test data"
		assert.Equal(t, document.CreatedAt, documentResponse.CreatedAt)
	})

	t.Run("DbDocumentToResponse", func(t *testing.T) {
		documentResponse := models.DbDocumentToResponse(document)

		assert.Equal(t, document.ID, documentResponse.ID)
		assert.Equal(t, "data:application/pdf;base64,"+"dGVzdCBkYXRh", documentResponse.Data) // base64 for "test data"
		assert.Equal(t, document.CreatedAt, documentResponse.CreatedAt)
	})
}

func TestDocumentRequest_ToDbDocument(t *testing.T) {
	t.Run("ValidBase64Data", func(t *testing.T) {
		documentRequest := models.DocumentRequest{
			Name: "Test Document",
			Data: "dGVzdCBkYXRh", // base64 for "test data"
		}

		dbDocument := documentRequest.ToDbDocument()

		assert.NotNil(t, dbDocument)
		assert.Equal(t, documentRequest.Name, dbDocument.Name)
		assert.Equal(t, []byte("test data"), dbDocument.Data)
	})

	t.Run("InvalidBase64Data", func(t *testing.T) {
		documentRequest := models.DocumentRequest{
			Name: "Invalid Document",
			Data: "invalid_base64_data",
		}

		dbDocument := documentRequest.ToDbDocument()

		assert.Nil(t, dbDocument)
	})
}
