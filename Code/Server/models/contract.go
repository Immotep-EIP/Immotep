package models

import (
	"immotep/backend/prisma/db"
)

type InviteRequest struct {
	TenantEmail string       `binding:"required,email" json:"tenant_email"`
	StartDate   db.DateTime  `binding:"required"       json:"start_date"`
	EndDate     *db.DateTime `binding:"-"              json:"end_date,omitempty"`
}

func (i *InviteRequest) ToDbLeaseInvite() db.LeaseInviteModel {
	return db.LeaseInviteModel{
		InnerLeaseInvite: db.InnerLeaseInvite{
			TenantEmail: i.TenantEmail,
			StartDate:   i.StartDate,
			EndDate:     i.EndDate,
		},
	}
}

type InviteResponse struct {
	ID          string       `json:"id"`
	TenantEmail string       `json:"tenant_email"`
	StartDate   db.DateTime  `json:"start_date"`
	EndDate     *db.DateTime `json:"end_date"`
	PropertyID  string       `json:"property_id"`
	CreatedAt   db.DateTime  `json:"created_at"`
}

func (i *InviteResponse) FromDbLeaseInvite(model db.LeaseInviteModel) {
	i.ID = model.ID
	i.TenantEmail = model.TenantEmail
	i.StartDate = model.StartDate
	i.EndDate = model.InnerLeaseInvite.EndDate
	i.PropertyID = model.PropertyID
	i.CreatedAt = model.CreatedAt
}

func DbLeaseInviteToResponse(pc db.LeaseInviteModel) InviteResponse {
	var resp InviteResponse
	resp.FromDbLeaseInvite(pc)
	return resp
}
