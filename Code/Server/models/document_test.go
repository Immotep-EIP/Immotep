package models_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"immotep/backend/models"
	"immotep/backend/prisma/db"
)

func TestDocumentResponse(t *testing.T) {
	document := db.DocumentModel{
		InnerDocument: db.InnerDocument{
			ID:        "1",
			Data:      []byte("test data"),
			Type:      db.DocTypePdf,
			CreatedAt: time.Now(),
		},
	}

	t.Run("FromDbDocument_PDF", func(t *testing.T) {
		nDocument := document
		documentResponse := models.DocumentResponse{}
		documentResponse.FromDbDocument(nDocument)

		assert.Equal(t, nDocument.ID, documentResponse.ID)
		assert.Equal(t, "data:application/pdf;base64,"+"dGVzdCBkYXRh", documentResponse.Data) // base64 for "test data"
		assert.Equal(t, nDocument.CreatedAt, documentResponse.CreatedAt)
	})

	t.Run("FromDbDocument_DOCX", func(t *testing.T) {
		nDocument := document
		nDocument.Type = db.DocTypeDocx
		documentResponse := models.DocumentResponse{}
		documentResponse.FromDbDocument(nDocument)

		assert.Equal(t, nDocument.ID, documentResponse.ID)
		assert.Equal(t, "data:application/vnd.openxmlformats-officedocument.wordprocessingml.document;base64,"+"dGVzdCBkYXRh", documentResponse.Data) // base64 for "test data"
		assert.Equal(t, nDocument.CreatedAt, documentResponse.CreatedAt)
	})

	t.Run("FromDbDocument_XLSX", func(t *testing.T) {
		nDocument := document
		nDocument.Type = db.DocTypeXlsx
		documentResponse := models.DocumentResponse{}
		documentResponse.FromDbDocument(nDocument)

		assert.Equal(t, nDocument.ID, documentResponse.ID)
		assert.Equal(t, "data:application/vnd.openxmlformats-officedocument.spreadsheetml.sheet;base64,"+"dGVzdCBkYXRh", documentResponse.Data) // base64 for "test data"
		assert.Equal(t, nDocument.CreatedAt, documentResponse.CreatedAt)
	})

	t.Run("FromDbDocument_InvalidType", func(t *testing.T) {
		nDocument := document
		nDocument.Type = "wrong"
		documentResponse := models.DocumentResponse{}

		assert.Panics(t, func() {
			documentResponse.FromDbDocument(nDocument)
		})
	})

	t.Run("DbDocumentToResponse", func(t *testing.T) {
		documentResponse := models.DbDocumentToResponse(document)

		assert.Equal(t, document.ID, documentResponse.ID)
		assert.Equal(t, "data:application/pdf;base64,"+"dGVzdCBkYXRh", documentResponse.Data) // base64 for "test data"
		assert.Equal(t, document.CreatedAt, documentResponse.CreatedAt)
	})
}

func TestDocumentRequest(t *testing.T) {
	t.Run("ValidPDF", func(t *testing.T) {
		documentRequest := models.DocumentRequest{
			Name: "Test Document",
			Data: "data:application/pdf;base64,dGVzdCBkYXRh", // base64 for "test data"
		}

		dbDocument := documentRequest.ToDbDocument()

		require.NotNil(t, dbDocument)
		assert.Equal(t, documentRequest.Name, dbDocument.Name)
		assert.Equal(t, []byte("test data"), dbDocument.Data)
		assert.Equal(t, db.DocTypePdf, dbDocument.Type)
	})

	t.Run("ValidDOCX", func(t *testing.T) {
		documentRequest := models.DocumentRequest{
			Name: "Test Document",
			Data: "data:application/vnd.openxmlformats-officedocument.wordprocessingml.document;base64,dGVzdCBkYXRh", // base64 for "test data"
		}

		dbDocument := documentRequest.ToDbDocument()

		require.NotNil(t, dbDocument)
		assert.Equal(t, documentRequest.Name, dbDocument.Name)
		assert.Equal(t, []byte("test data"), dbDocument.Data)
		assert.Equal(t, db.DocTypeDocx, dbDocument.Type)
	})

	t.Run("ValidXLSX", func(t *testing.T) {
		documentRequest := models.DocumentRequest{
			Name: "Test Document",
			Data: "data:application/vnd.openxmlformats-officedocument.spreadsheetml.sheet;base64,dGVzdCBkYXRh", // base64 for "test data"
		}

		dbDocument := documentRequest.ToDbDocument()

		require.NotNil(t, dbDocument)
		assert.Equal(t, documentRequest.Name, dbDocument.Name)
		assert.Equal(t, []byte("test data"), dbDocument.Data)
		assert.Equal(t, db.DocTypeXlsx, dbDocument.Type)
	})

	t.Run("InvalidType", func(t *testing.T) {
		documentRequest := models.DocumentRequest{
			Name: "Test Document",
			Data: "data:text/plain;base64,dGVzdCBkYXRh", // base64 for "test data"
		}

		dbDocument := documentRequest.ToDbDocument()

		assert.Nil(t, dbDocument)
	})

	t.Run("InvalidBase64Data", func(t *testing.T) {
		documentRequest := models.DocumentRequest{
			Name: "Invalid Document",
			Data: "data:application/pdf;base64,invalid_base64_data",
		}

		dbDocument := documentRequest.ToDbDocument()

		assert.Nil(t, dbDocument)
	})
}
