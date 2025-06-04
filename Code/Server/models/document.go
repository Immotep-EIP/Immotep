package models

import (
	"encoding/base64"
	"strings"

	"immotep/backend/prisma/db"
)

type DocumentRequest struct {
	Name string `binding:"required"         json:"name"`
	Data string `binding:"required,datauri" json:"data"`
}

func (i *DocumentRequest) ToDbDocument() *db.DocumentModel {
	var docType db.DocType

	switch {
	case strings.HasPrefix(i.Data, "data:application/pdf;base64,"):
		i.Data = strings.TrimPrefix(i.Data, "data:application/pdf;base64,")
		docType = db.DocTypePdf
	case strings.HasPrefix(i.Data, "data:application/vnd.openxmlformats-officedocument.wordprocessingml.document;base64,"):
		i.Data = strings.TrimPrefix(i.Data, "data:application/vnd.openxmlformats-officedocument.wordprocessingml.document;base64,")
		docType = db.DocTypeDocx
	case strings.HasPrefix(i.Data, "data:application/vnd.openxmlformats-officedocument.spreadsheetml.sheet;base64,"):
		i.Data = strings.TrimPrefix(i.Data, "data:application/vnd.openxmlformats-officedocument.spreadsheetml.sheet;base64,")
		docType = db.DocTypeXlsx
	default:
		return nil
	}

	decoded, err := base64.StdEncoding.DecodeString(i.Data)
	if err != nil {
		return nil
	}

	return &db.DocumentModel{
		InnerDocument: db.InnerDocument{
			Name: i.Name,
			Data: decoded,
			Type: docType,
		},
	}
}

type DocumentResponse struct {
	ID        string      `json:"id"`
	Name      string      `json:"name"`
	Data      string      `json:"data"`
	CreatedAt db.DateTime `json:"created_at"`
}

func (i *DocumentResponse) FromDbDocument(model db.DocumentModel) {
	i.ID = model.ID
	i.Name = model.Name
	switch model.Type {
	case db.DocTypePdf:
		i.Data = "data:application/pdf;base64,"
	case db.DocTypeDocx:
		i.Data = "data:application/vnd.openxmlformats-officedocument.wordprocessingml.document;base64,"
	case db.DocTypeXlsx:
		i.Data = "data:application/vnd.openxmlformats-officedocument.spreadsheetml.sheet;base64,"
	default:
		panic("unknown document type")
	}
	i.Data += base64.StdEncoding.EncodeToString(model.Data)
	i.CreatedAt = model.CreatedAt
}

func DbDocumentToResponse(pc db.DocumentModel) DocumentResponse {
	var resp DocumentResponse
	resp.FromDbDocument(pc)
	return resp
}
