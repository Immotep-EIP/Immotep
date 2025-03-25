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

func TestGetLeaseDocuments(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	documents := []db.DocumentModel{
		BuildTestDocument("1"),
		BuildTestDocument("2"),
	}

	mock.Document.Expect(
		client.Client.Document.FindMany(
			db.Document.LeaseID.Equals("1"),
		),
	).ReturnsMany(documents)

	foundDocuments := database.GetLeaseDocuments("1")
	assert.NotNil(t, foundDocuments)
	assert.Len(t, foundDocuments, len(documents))
	assert.Equal(t, documents[0].ID, foundDocuments[0].ID)
	assert.Equal(t, documents[1].ID, foundDocuments[1].ID)
}

func TestGetLeaseDocuments_NoDocuments(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	mock.Document.Expect(
		client.Client.Document.FindMany(
			db.Document.LeaseID.Equals("1"),
		),
	).ReturnsMany([]db.DocumentModel{})

	foundDocuments := database.GetLeaseDocuments("1")
	assert.NotNil(t, foundDocuments)
	assert.Empty(t, foundDocuments)
}

func TestGetLeaseDocuments_NoConnection(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	mock.Document.Expect(
		client.Client.Document.FindMany(
			db.Document.LeaseID.Equals("1"),
		),
	).Errors(errors.New("connection error"))

	assert.Panics(t, func() {
		database.GetLeaseDocuments("1")
	})
}

func TestGetDocumentByID(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	document := BuildTestDocument("1")

	mock.Document.Expect(
		client.Client.Document.FindUnique(db.Document.ID.Equals("1")),
	).Returns(document)

	foundDocument := database.GetDocumentByID("1")
	assert.NotNil(t, foundDocument)
	assert.Equal(t, document.ID, foundDocument.ID)
}

func TestGetDocumentByID_NotFound(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	mock.Document.Expect(
		client.Client.Document.FindUnique(db.Document.ID.Equals("1")),
	).Errors(db.ErrNotFound)

	foundDocument := database.GetDocumentByID("1")
	assert.Nil(t, foundDocument)
}

func TestGetDocumentByID_NoConnection(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	mock.Document.Expect(
		client.Client.Document.FindUnique(db.Document.ID.Equals("1")),
	).Errors(errors.New("connection error"))

	assert.Panics(t, func() {
		database.GetDocumentByID("1")
	})
}

func TestCreateDocument(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	document := BuildTestDocument("1")

	mock.Document.Expect(
		client.Client.Document.CreateOne(
			db.Document.Name.Set(document.Name),
			db.Document.Data.Set(document.Data),
			db.Document.Lease.Link(db.Lease.ID.Equals(document.LeaseID)),
		),
	).Returns(document)

	newDocument := database.CreateDocument(document)
	assert.NotNil(t, newDocument)
	assert.Equal(t, document.ID, newDocument.ID)
}

func TestCreateDocument_NoConnection(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	document := BuildTestDocument("1")

	mock.Document.Expect(
		client.Client.Document.CreateOne(
			db.Document.Name.Set(document.Name),
			db.Document.Data.Set(document.Data),
			db.Document.Lease.Link(db.Lease.ID.Equals(document.LeaseID)),
		),
	).Errors(errors.New("connection error"))

	assert.Panics(t, func() {
		database.CreateDocument(document)
	})
}
