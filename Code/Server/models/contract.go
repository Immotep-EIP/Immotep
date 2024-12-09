package models

import (
	"immotep/backend/prisma/db"
)

type InviteRequest struct {
	TenantEmail string       `binding:"required,email" json:"tenant_email"`
	StartDate   db.DateTime  `binding:"required"       json:"start_date"`
	EndDate     *db.DateTime `binding:"-"              json:"end_date,omitempty"`
}

func (i *InviteRequest) ToDbPendingContract() db.PendingContractModel {
	return db.PendingContractModel{
		InnerPendingContract: db.InnerPendingContract{
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

func (i *InviteResponse) FromDbPendingContract(model db.PendingContractModel) {
	i.ID = model.ID
	i.TenantEmail = model.TenantEmail
	i.StartDate = model.StartDate
	i.EndDate = model.InnerPendingContract.EndDate
	i.PropertyID = model.PropertyID
	i.CreatedAt = model.CreatedAt
}

func DbPendingContractToResponse(pc db.PendingContractModel) InviteResponse {
	var resp InviteResponse
	resp.FromDbPendingContract(pc)
	return resp
}
