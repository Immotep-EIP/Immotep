package models

import (
	"encoding/base64"

	"immotep/backend/prisma/db"
)

type DocumentRequest struct {
	Name string `binding:"required"        json:"name"`
	Data string `binding:"required,base64" json:"data"`
}

func (i *DocumentRequest) ToDbDocument() *db.DocumentModel {
	decoded, err := base64.StdEncoding.DecodeString(i.Data)
	if err != nil {
		return nil
	}

	return &db.DocumentModel{
		InnerDocument: db.InnerDocument{
			Name: i.Name,
			Data: decoded,
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
	i.Data = "data:application/pdf;base64," + base64.StdEncoding.EncodeToString(model.Data)
	i.CreatedAt = model.CreatedAt
}

func DbDocumentToResponse(pc db.DocumentModel) DocumentResponse {
	var resp DocumentResponse
	resp.FromDbDocument(pc)
	return resp
}
