package database

import (
	"immotep/backend/prisma/db"
	"immotep/backend/services"
)

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
		db.Document.Contract.Link(db.Contract.TenantIDPropertyID(
			db.Contract.TenantID.Equals(doc.ContractTenantID),
			db.Contract.PropertyID.Equals(doc.ContractPropertyID),
		)),
	).Exec(pdb.Context)
	if err != nil || newDocument == nil {
		panic(err)
	}
	return *newDocument
}
