package models

import (
	"immotep/backend/prisma/db"
)

type PropertyRequest struct {
	Name       string `binding:"required" json:"name"`
	Address    string `binding:"required" json:"address"`
	City       string `binding:"required" json:"city"`
	PostalCode string `binding:"required" json:"postal_code"`
	Country    string `binding:"required" json:"country"`
}

func (p *PropertyRequest) ToDbPendingContract() db.PropertyModel {
	return db.PropertyModel{
		InnerProperty: db.InnerProperty{
			Name:       p.Name,
			Address:    p.Address,
			City:       p.City,
			PostalCode: p.PostalCode,
			Country:    p.Country,
		},
	}
}

type PropertyResponse struct {
	ID         string      `json:"id"`
	OwnerID    string      `json:"owner_id"`
	Name       string      `json:"name"`
	Address    string      `json:"address"`
	City       string      `json:"city"`
	PostalCode string      `json:"postal_code"`
	Country    string      `json:"country"`
	CreatedAt  db.DateTime `json:"created_at"`
}

func (p *PropertyResponse) FromDbProperty(model db.PropertyModel) {
	p.ID = model.ID
	p.OwnerID = model.OwnerID
	p.Name = model.Name
	p.Address = model.Address
	p.City = model.City
	p.PostalCode = model.PostalCode
	p.Country = model.Country
	p.CreatedAt = model.CreatedAt
}

func DbPropertyToResponse(pc db.PropertyModel) PropertyResponse {
	var resp PropertyResponse
	resp.FromDbProperty(pc)
	return resp
}
