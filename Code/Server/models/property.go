package models

import (
	"slices"

	"immotep/backend/prisma/db"
	"immotep/backend/utils"
)

type PropertyRequest struct {
	Name                string  `binding:"required"       json:"name"`
	Address             string  `binding:"required"       json:"address"`
	City                string  `binding:"required"       json:"city"`
	PostalCode          string  `binding:"required"       json:"postal_code"`
	Country             string  `binding:"required"       json:"country"`
	AreaSqm             float64 `binding:"required"       json:"area_sqm"`
	RentalPricePerMonth int     `binding:"required"       json:"rental_price_per_month"`
	DepositPrice        int     `binding:"required"       json:"deposit_price"`
	Picture             *string `json:"picture,omitempty"`
}

func (p *PropertyRequest) ToDbProperty() db.PropertyModel {
	return db.PropertyModel{
		InnerProperty: db.InnerProperty{
			Name:                p.Name,
			Address:             p.Address,
			City:                p.City,
			PostalCode:          p.PostalCode,
			Country:             p.Country,
			AreaSqm:             p.AreaSqm,
			RentalPricePerMonth: p.RentalPricePerMonth,
			DepositPrice:        p.DepositPrice,
			Picture:             p.Picture,
		},
	}
}

type PropertyResponse struct {
	ID                  string      `json:"id"`
	OwnerID             string      `json:"owner_id"`
	Name                string      `json:"name"`
	Address             string      `json:"address"`
	City                string      `json:"city"`
	PostalCode          string      `json:"postal_code"`
	Country             string      `json:"country"`
	AreaSqm             float64     `json:"area_sqm"`
	RentalPricePerMonth int         `json:"rental_price_per_month"`
	DepositPrice        int         `json:"deposit_price"`
	Picture             *string     `json:"picture"`
	CreatedAt           db.DateTime `json:"created_at"`

	// calculated fields
	Status    string       `json:"status"`
	NbDamage  int          `json:"nb_damage"`
	Tenant    string       `json:"tenant"`
	StartDate *db.DateTime `json:"start_date"`
	EndDate   *db.DateTime `json:"end_date"`
}

func (p *PropertyResponse) FromDbProperty(model db.PropertyModel) {
	p.ID = model.ID
	p.OwnerID = model.OwnerID
	p.Name = model.Name
	p.Address = model.Address
	p.City = model.City
	p.PostalCode = model.PostalCode
	p.Country = model.Country
	p.AreaSqm = model.AreaSqm
	p.RentalPricePerMonth = model.RentalPricePerMonth
	p.DepositPrice = model.DepositPrice
	p.Picture = model.InnerProperty.Picture
	p.CreatedAt = model.CreatedAt

	p.NbDamage = utils.CountIf(model.Damages(), func(x db.DamageModel) bool { return x.InnerDamage.FixedAt == nil })

	activeIndex := slices.IndexFunc(model.Contracts(), func(x db.ContractModel) bool { return x.Active })
	if activeIndex != -1 {
		active := model.Contracts()[activeIndex]

		p.Status = "unavailable"
		p.Tenant = active.Tenant().Firstname + " " + active.Tenant().Lastname
		p.StartDate = &active.StartDate
		p.EndDate = active.InnerContract.EndDate
	} else {
		p.Status = "available"
		p.Tenant = ""
		p.StartDate = nil
		p.EndDate = nil
	}
}

func DbPropertyToResponse(pc db.PropertyModel) PropertyResponse {
	var resp PropertyResponse
	resp.FromDbProperty(pc)
	return resp
}
