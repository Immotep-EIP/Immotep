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
