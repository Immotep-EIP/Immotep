package database

import (
	"immotep/backend/prisma/db"
	"immotep/backend/services"
)

func GetCurrentActiveContractDocuments(propertyID string) []db.DocumentModel {
	activeContract := GetCurrentActiveContract(propertyID)
	if activeContract == nil {
		panic("No active contract found for property: " + propertyID)
	}

	pdb := services.DBclient
	documents, err := pdb.Client.Document.FindMany(
		db.Document.ContractID.Equals(activeContract.ID),
	).Exec(pdb.Context)
	if err != nil {
		panic(err)
	}
	return documents
}

func GetDocumentByID(id string) *db.DocumentModel {
	pdb := services.DBclient
	doc, err := pdb.Client.Document.FindUnique(db.Document.ID.Equals(id)).Exec(pdb.Context)
	if err != nil {
		if db.IsErrNotFound(err) {
			return nil
		}
		panic(err)
	}
	return doc
}

func CreateDocument(doc db.DocumentModel) db.DocumentModel {
	pdb := services.DBclient
	newDocument, err := pdb.Client.Document.CreateOne(
		db.Document.Name.Set(doc.Name),
		db.Document.Data.Set(doc.Data),
		db.Document.Contract.Link(db.Contract.ID.Equals(doc.ContractID)),
	).Exec(pdb.Context)
	if err != nil || newDocument == nil {
		panic(err)
	}
	return *newDocument
}
