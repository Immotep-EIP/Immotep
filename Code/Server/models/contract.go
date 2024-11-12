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
	EndDate     *db.DateTime `json:"end_date,omitempty"`
	PropertyID  string       `json:"property_id"`
	CreatedAt   db.DateTime  `json:"created_at"`
}

func (u *InviteResponse) FromDbPendingContract(pc db.PendingContractModel) {
	u.ID = pc.ID
	u.TenantEmail = pc.TenantEmail
	u.StartDate = pc.StartDate
	u.EndDate = pc.InnerPendingContract.EndDate
	u.PropertyID = pc.PropertyID
	u.CreatedAt = pc.CreatedAt
}

func DbPendingContractToResponse(pc db.PendingContractModel) InviteResponse {
	var userResp InviteResponse
	userResp.FromDbPendingContract(pc)
	return userResp
}
