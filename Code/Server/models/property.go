package models

import (
	"slices"

	"keyz/backend/prisma/db"
	// "keyz/backend/services/database"
	"keyz/backend/utils"
)

type PropertyStatus string

const (
	StatusUnavailable PropertyStatus = "unavailable"
	StatusInviteSent  PropertyStatus = "invite sent"
	StatusAvailable   PropertyStatus = "available"
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

type propertyLeaseResponse struct {
	ID          string       `json:"id"`
	TenantName  string       `json:"tenant_name"`
	TenantEmail string       `json:"tenant_email"`
	Active      bool         `json:"active"`
	StartDate   db.DateTime  `json:"start_date"`
	EndDate     *db.DateTime `json:"end_date"`
}

type propertyInviteResponse struct {
	TenantEmail string       `json:"tenant_email"`
	StartDate   db.DateTime  `json:"start_date"`
	EndDate     *db.DateTime `json:"end_date"`
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

	NbDamage int                     `json:"nb_damage"`
	Status   PropertyStatus          `json:"status"`
	Lease    *propertyLeaseResponse  `json:"lease"`
	Invite   *propertyInviteResponse `json:"invite,omitempty"`
}

func (p *PropertyResponse) FromDbProperty(model db.PropertyModel, leaseId string) {
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

	p.NbDamage = 0
	for _, lease := range model.Leases() {
		p.NbDamage += utils.CountIf(lease.Damages(), func(x db.DamageModel) bool { return x.InnerDamage.FixedAt == nil })
	}
	p.Lease = nil
	p.Invite = nil

	invite, inviteOk := model.LeaseInvite()
	switch {
	case slices.ContainsFunc(model.Leases(), func(x db.LeaseModel) bool { return x.Active }):
		p.Status = StatusUnavailable
	case inviteOk:
		p.Status = StatusInviteSent
		p.Invite = &propertyInviteResponse{
			TenantEmail: invite.TenantEmail,
			StartDate:   invite.StartDate,
			EndDate:     invite.InnerLeaseInvite.EndDate,
		}
	default:
		p.Status = StatusAvailable
	}

	iLease := slices.IndexFunc(model.Leases(), func(x db.LeaseModel) bool { return utils.Ternary(leaseId == "current", x.Active, x.ID == leaseId) })
	if iLease != -1 {
		active := model.Leases()[iLease]
		p.Lease = &propertyLeaseResponse{
			ID:          active.ID,
			TenantName:  active.Tenant().Name(),
			TenantEmail: active.Tenant().Email,
			Active:      active.Active,
			StartDate:   active.StartDate,
			EndDate:     active.InnerLease.EndDate,
		}
	}
}

func DbPropertyToResponse(pc db.PropertyModel, leaseId string) PropertyResponse {
	var resp PropertyResponse
	resp.FromDbProperty(pc, leaseId)
	return resp
}

type furnitureResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Archived bool   `json:"archived"`
}

type roomResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Archived bool   `json:"archived"`

	Furnitures []furnitureResponse `json:"furnitures"`
}

type PropertyInventoryResponse struct {
	PropertyResponse
	Rooms []roomResponse `json:"rooms"`
}

func (p *PropertyInventoryResponse) FromDbProperty(model db.PropertyModel, leaseId string) {
	p.PropertyResponse.FromDbProperty(model, leaseId)

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

func DbPropertyInventoryToResponse(pc db.PropertyModel, leaseId string) PropertyInventoryResponse {
	var resp PropertyInventoryResponse
	resp.FromDbProperty(pc, leaseId)
	return resp
}
