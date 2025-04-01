package database

import (
	"immotep/backend/prisma/db"
	"immotep/backend/services"
)

func GetDocumentsByLease(leaseID string) []db.DocumentModel {
	pdb := services.DBclient
	documents, err := pdb.Client.Document.FindMany(
		db.Document.LeaseID.Equals(leaseID),
	).Exec(pdb.Context)
	if err != nil {
		panic(err)
	}
	return documents
}

func MockGetDocumentsByLease(c *services.PrismaDB) db.DocumentMockExpectParam {
	return c.Client.Document.FindMany(
		db.Document.LeaseID.Equals("1"),
	)
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

func MockGetDocumentByID(c *services.PrismaDB) db.DocumentMockExpectParam {
	return c.Client.Document.FindUnique(
		db.Document.ID.Equals("1"),
	)
}

func CreateDocument(doc db.DocumentModel, leaseId string) db.DocumentModel {
	pdb := services.DBclient
	newDocument, err := pdb.Client.Document.CreateOne(
		db.Document.Name.Set(doc.Name),
		db.Document.Data.Set(doc.Data),
		db.Document.Lease.Link(db.Lease.ID.Equals(leaseId)),
	).Exec(pdb.Context)
	if err != nil || newDocument == nil {
		panic(err)
	}
	return *newDocument
}

func MockCreateDocument(c *services.PrismaDB, document db.DocumentModel) db.DocumentMockExpectParam {
	return c.Client.Document.CreateOne(
		db.Document.Name.Set(document.Name),
		db.Document.Data.Set(document.Data),
		db.Document.Lease.Link(db.Lease.ID.Equals(document.LeaseID)),
	)
}
