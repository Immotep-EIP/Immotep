package database_test

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"immotep/backend/prisma/db"
	"immotep/backend/services"
	"immotep/backend/services/database"
)

func BuildTestDocument(id string) db.DocumentModel {
	return db.DocumentModel{
		InnerDocument: db.InnerDocument{
			ID:        id,
			Name:      "Document",
			Data:      []byte("data"),
			LeaseID:   "1",
			CreatedAt: time.Now(),
		},
	}
}

// #############################################################################

func TestGetLeaseDocuments(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	documents := []db.DocumentModel{
		BuildTestDocument("1"),
		BuildTestDocument("2"),
	}
	m.Document.Expect(database.MockGetDocumentsByLease(c)).ReturnsMany(documents)

	foundDocuments := database.GetDocumentsByLease("1")
	assert.NotNil(t, foundDocuments)
	assert.Len(t, foundDocuments, len(documents))
	assert.Equal(t, documents[0].ID, foundDocuments[0].ID)
	assert.Equal(t, documents[1].ID, foundDocuments[1].ID)
}

func TestGetLeaseDocuments_NoDocuments(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.Document.Expect(database.MockGetDocumentsByLease(c)).ReturnsMany([]db.DocumentModel{})

	foundDocuments := database.GetDocumentsByLease("1")
	assert.NotNil(t, foundDocuments)
	assert.Empty(t, foundDocuments)
}

func TestGetLeaseDocuments_NoConnection(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.Document.Expect(database.MockGetDocumentsByLease(c)).Errors(errors.New("connection error"))

	assert.Panics(t, func() {
		database.GetDocumentsByLease("1")
	})
}

// #############################################################################

func TestGetDocumentByID(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	document := BuildTestDocument("1")
	m.Document.Expect(database.MockGetDocumentByID(c)).Returns(document)

	foundDocument := database.GetDocumentByID("1")
	assert.NotNil(t, foundDocument)
	assert.Equal(t, document.ID, foundDocument.ID)
}

func TestGetDocumentByID_NotFound(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.Document.Expect(database.MockGetDocumentByID(c)).Errors(db.ErrNotFound)

	foundDocument := database.GetDocumentByID("1")
	assert.Nil(t, foundDocument)
}

func TestGetDocumentByID_NoConnection(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.Document.Expect(database.MockGetDocumentByID(c)).Errors(errors.New("connection error"))

	assert.Panics(t, func() {
		database.GetDocumentByID("1")
	})
}

// #############################################################################

func TestCreateDocument(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	document := BuildTestDocument("1")
	m.Document.Expect(database.MockCreateDocument(c, document)).Returns(document)

	newDocument := database.CreateDocument(document, document.LeaseID)
	assert.NotNil(t, newDocument)
	assert.Equal(t, document.ID, newDocument.ID)
}

func TestCreateDocument_NoConnection(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	document := BuildTestDocument("1")

	m.Document.Expect(database.MockCreateDocument(c, document)).Errors(errors.New("connection error"))

	assert.Panics(t, func() {
		database.CreateDocument(document, document.LeaseID)
	})
}

// #############################################################################

func TestDeleteDocument(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.Document.Expect(database.MockDeleteDocument(c)).Returns(BuildTestDocument("1"))

	database.DeleteDocument("1")
}

func TestDeleteDocument_NotFound(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.Document.Expect(database.MockDeleteDocument(c)).Errors(db.ErrNotFound)

	assert.Panics(t, func() {
		database.DeleteDocument("1")
	})
}
