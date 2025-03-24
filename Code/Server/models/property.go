package models

import (
	"slices"

	"immotep/backend/prisma/db"
	"immotep/backend/utils"
)

type PropertyRequest struct {
	Name                string  `binding:"required"                json:"name"`
	Address             string  `binding:"required"                json:"address"`
	ApartmentNumber     *string `json:"apartment_number,omitempty"`
	City                string  `binding:"required"                json:"city"`
	PostalCode          string  `binding:"required"                json:"postal_code"`
	Country             string  `binding:"required"                json:"country"`
	AreaSqm             float64 `binding:"required"                json:"area_sqm"`
	RentalPricePerMonth float64 `binding:"required"                json:"rental_price_per_month"`
	DepositPrice        float64 `binding:"required"                json:"deposit_price"`
}

func (p *PropertyRequest) ToDbProperty() db.PropertyModel {
	return db.PropertyModel{
		InnerProperty: db.InnerProperty{
			Name:                p.Name,
			Address:             p.Address,
			ApartmentNumber:     p.ApartmentNumber,
			City:                p.City,
			PostalCode:          p.PostalCode,
			Country:             p.Country,
			AreaSqm:             p.AreaSqm,
			RentalPricePerMonth: p.RentalPricePerMonth,
			DepositPrice:        p.DepositPrice,
		},
	}
}

type PropertyUpdateRequest struct {
	Name                *string  `json:"name,omitempty"`
	Address             *string  `json:"address,omitempty"`
	ApartmentNumber     *string  `json:"apartment_number,omitempty"`
	City                *string  `json:"city,omitempty"`
	PostalCode          *string  `json:"postal_code,omitempty"`
	Country             *string  `json:"country,omitempty"`
	AreaSqm             *float64 `json:"area_sqm,omitempty"`
	RentalPricePerMonth *float64 `json:"rental_price_per_month,omitempty"`
	DepositPrice        *float64 `json:"deposit_price,omitempty"`
}

type PropertyResponse struct {
	ID                  string      `json:"id"`
	OwnerID             string      `json:"owner_id"`
	PictureID           *string     `json:"picture_id,omitempty"`
	Name                string      `json:"name"`
	Address             string      `json:"address"`
	ApartmentNumber     *string     `json:"apartment_number,omitempty"`
	City                string      `json:"city"`
	PostalCode          string      `json:"postal_code"`
	Country             string      `json:"country"`
	AreaSqm             float64     `json:"area_sqm"`
	RentalPricePerMonth float64     `json:"rental_price_per_month"`
	DepositPrice        float64     `json:"deposit_price"`
	CreatedAt           db.DateTime `json:"created_at"`
	Archived            bool        `json:"archived"`

	// calculated fields

	NbDamage  int          `json:"nb_damage"`
	Status    string       `json:"status"`
	Tenant    string       `json:"tenant,omitempty"`
	StartDate *db.DateTime `json:"start_date,omitempty"`
	EndDate   *db.DateTime `json:"end_date,omitempty"`
}

func (p *PropertyResponse) FromDbProperty(model db.PropertyModel) {
	p.ID = model.ID
	p.OwnerID = model.OwnerID
	p.PictureID = model.InnerProperty.PictureID
	p.Name = model.Name
	p.Address = model.Address
	p.ApartmentNumber = model.InnerProperty.ApartmentNumber
	p.City = model.City
	p.PostalCode = model.PostalCode
	p.Country = model.Country
	p.AreaSqm = model.AreaSqm
	p.RentalPricePerMonth = model.RentalPricePerMonth
	p.DepositPrice = model.DepositPrice
	p.CreatedAt = model.CreatedAt
	p.Archived = model.Archived

	p.NbDamage = utils.CountIf(model.Damages(), func(x db.DamageModel) bool { return x.InnerDamage.FixedAt == nil })

	activeIndex := slices.IndexFunc(model.Leases(), func(x db.LeaseModel) bool { return x.Active })
	invite, inviteOk := model.PendingContract()
	switch {
	case activeIndex != -1:
		active := model.Leases()[activeIndex]
		p.Status = "unavailable"
		p.Tenant = active.Tenant().Firstname + " " + active.Tenant().Lastname
		p.StartDate = &active.StartDate
		p.EndDate = active.InnerLease.EndDate
	case inviteOk:
		p.Status = "invite sent"
		p.Tenant = invite.TenantEmail
		p.StartDate = &invite.StartDate
		p.EndDate = invite.InnerPendingContract.EndDate
	default:
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

type PropertyInventoryResponse struct {
	ID                  string      `json:"id"`
	OwnerID             string      `json:"owner_id"`
	PictureID           *string     `json:"picture_id,omitempty"`
	Name                string      `json:"name"`
	Address             string      `json:"address"`
	ApartmentNumber     *string     `json:"apartment_number,omitempty"`
	City                string      `json:"city"`
	PostalCode          string      `json:"postal_code"`
	Country             string      `json:"country"`
	AreaSqm             float64     `json:"area_sqm"`
	RentalPricePerMonth float64     `json:"rental_price_per_month"`
	DepositPrice        float64     `json:"deposit_price"`
	CreatedAt           db.DateTime `json:"created_at"`
	Archived            bool        `json:"archived"`

	// calculated fields

	NbDamage  int          `json:"nb_damage"`
	Status    string       `json:"status"`
	Tenant    string       `json:"tenant,omitempty"`
	StartDate *db.DateTime `json:"start_date,omitempty"`
	EndDate   *db.DateTime `json:"end_date,omitempty"`

	Rooms []roomResponse `json:"rooms"`
}

type roomResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Archived bool   `json:"archived"`

	Furnitures []furnitureResponse `json:"furnitures"`
}

type furnitureResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Archived bool   `json:"archived"`
}

func (p *PropertyInventoryResponse) FromDbProperty(model db.PropertyModel) {
	p.ID = model.ID
	p.OwnerID = model.OwnerID
	p.PictureID = model.InnerProperty.PictureID
	p.Name = model.Name
	p.Address = model.Address
	p.ApartmentNumber = model.InnerProperty.ApartmentNumber
	p.City = model.City
	p.PostalCode = model.PostalCode
	p.Country = model.Country
	p.AreaSqm = model.AreaSqm
	p.RentalPricePerMonth = model.RentalPricePerMonth
	p.DepositPrice = model.DepositPrice
	p.CreatedAt = model.CreatedAt
	p.Archived = model.Archived

	p.NbDamage = utils.CountIf(model.Damages(), func(x db.DamageModel) bool { return x.InnerDamage.FixedAt == nil })

	activeIndex := slices.IndexFunc(model.Leases(), func(x db.LeaseModel) bool { return x.Active })
	invite, inviteOk := model.PendingContract()
	switch {
	case activeIndex != -1:
		active := model.Leases()[activeIndex]
		p.Status = "unavailable"
		p.Tenant = active.Tenant().Firstname + " " + active.Tenant().Lastname
		p.StartDate = &active.StartDate
		p.EndDate = active.InnerLease.EndDate
	case inviteOk:
		p.Status = "invite sent"
		p.Tenant = invite.TenantEmail
		p.StartDate = &invite.StartDate
		p.EndDate = invite.InnerPendingContract.EndDate
	default:
		p.Status = "available"
		p.Tenant = ""
		p.StartDate = nil
		p.EndDate = nil
	}

	p.Rooms = make([]roomResponse, len(model.Rooms()))
	for i, room := range model.Rooms() {
		p.Rooms[i].ID = room.ID
		p.Rooms[i].Name = room.Name
		p.Rooms[i].Archived = room.Archived

		p.Rooms[i].Furnitures = make([]furnitureResponse, len(room.Furnitures()))
		for j, furniture := range room.Furnitures() {
			p.Rooms[i].Furnitures[j].ID = furniture.ID
			p.Rooms[i].Furnitures[j].Name = furniture.Name
			p.Rooms[i].Furnitures[j].Quantity = furniture.Quantity
			p.Rooms[i].Furnitures[j].Archived = furniture.Archived
		}
	}
}

func DbPropertyInventoryToResponse(pc db.PropertyModel) PropertyInventoryResponse {
	var resp PropertyInventoryResponse
	resp.FromDbProperty(pc)
	return resp
}
