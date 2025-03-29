package models

import (
	"immotep/backend/prisma/db"
)

type LeaseResponse struct {
	ID           string       `json:"id"`
	PropertyID   string       `json:"property_id"`
	PropertyName string       `json:"property_name"`
	OwnerID      string       `json:"owner_id"`
	OwnerName    string       `json:"owner_name"`
	OwnerEmail   string       `json:"owner_email"`
	TenantID     string       `json:"tenant_id"`
	TenantName   string       `json:"tenant_name"`
	TenantEmail  string       `json:"tenant_email"`
	Active       bool         `json:"active"`
	StartDate    db.DateTime  `json:"start_date"`
	EndDate      *db.DateTime `json:"end_date"`
	CreatedAt    db.DateTime  `json:"created_at"`
}

func (l *LeaseResponse) FromDbLease(model db.LeaseModel) {
	l.ID = model.ID
	l.PropertyID = model.PropertyID
	l.PropertyName = model.Property().Name
	l.OwnerID = model.Property().OwnerID
	l.OwnerName = model.Property().Owner().Name()
	l.OwnerEmail = model.Property().Owner().Email
	l.TenantID = model.TenantID
	l.TenantName = model.Tenant().Name()
	l.TenantEmail = model.Tenant().Email
	l.Active = model.Active
	l.StartDate = model.StartDate
	l.EndDate = model.InnerLease.EndDate
	l.CreatedAt = model.CreatedAt
}

func DbLeaseToResponse(model db.LeaseModel) LeaseResponse {
	var resp LeaseResponse
	resp.FromDbLease(model)
	return resp
}

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
