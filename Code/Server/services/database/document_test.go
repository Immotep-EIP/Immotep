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
			ID:         id,
			Name:       "Document",
			Data:       []byte("data"),
			ContractID: "1",
			CreatedAt:  time.Now(),
		},
	}
}

func TestGetCurrentActiveContractDocuments(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	contract := BuildTestContract()
	documents := []db.DocumentModel{
		BuildTestDocument("1"),
		BuildTestDocument("2"),
	}

	mock.Contract.Expect(
		client.Client.Contract.FindMany(
			db.Contract.PropertyID.Equals("1"),
			db.Contract.Active.Equals(true),
		),
	).ReturnsMany([]db.ContractModel{contract})

	mock.Document.Expect(
		client.Client.Document.FindMany(
			db.Document.ContractID.Equals(contract.ID),
		),
	).ReturnsMany(documents)

	foundDocuments := database.GetCurrentActiveContractDocuments("1")
	assert.NotNil(t, foundDocuments)
	assert.Equal(t, len(documents), len(foundDocuments))
	assert.Equal(t, documents[0].ID, foundDocuments[0].ID)
	assert.Equal(t, documents[1].ID, foundDocuments[1].ID)
}

func TestGetCurrentActiveContractDocuments_NoActiveContract(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	mock.Contract.Expect(
		client.Client.Contract.FindMany(
			db.Contract.PropertyID.Equals("1"),
			db.Contract.Active.Equals(true),
		),
	).ReturnsMany([]db.ContractModel{})

	assert.Panics(t, func() {
		database.GetCurrentActiveContractDocuments("1")
	})
}

func TestGetCurrentActiveContractDocuments_NoConnection(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	contract := BuildTestContract()

	mock.Contract.Expect(
		client.Client.Contract.FindMany(
			db.Contract.PropertyID.Equals("1"),
			db.Contract.Active.Equals(true),
		),
	).ReturnsMany([]db.ContractModel{contract})

	mock.Document.Expect(
		client.Client.Document.FindMany(
			db.Document.ContractID.Equals(contract.ID),
		),
	).Errors(errors.New("connection error"))

	assert.Panics(t, func() {
		database.GetCurrentActiveContractDocuments("1")
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
			db.Document.Contract.Link(db.Contract.ID.Equals(document.ContractID)),
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
			db.Document.Contract.Link(db.Contract.ID.Equals(document.ContractID)),
		),
	).Errors(errors.New("connection error"))

	assert.Panics(t, func() {
		database.CreateDocument(document)
	})
}
