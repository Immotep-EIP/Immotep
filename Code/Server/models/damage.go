package models

import (
	"immotep/backend/prisma/db"
)

type DamageRequest struct {
	Comment  string      `binding:"required"             json:"comment"`
	Priority db.Priority `binding:"required,priority"    json:"priority"`
	RoomID   string      `binding:"required"             json:"room_id"`
	Pictures []string    `binding:"dive,required,base64" json:"pictures"`
}

func (r *DamageRequest) ToDbDamage() db.DamageModel {
	return db.DamageModel{
		InnerDamage: db.InnerDamage{
			Comment:  r.Comment,
			Priority: r.Priority,
			RoomID:   r.RoomID,
		},
	}
}

type DamageOwnerUpdateRequest struct {
	Read         *bool        `json:"read,omitempty"`
	FixPlannedAt *db.DateTime `json:"fix_planned_at,omitempty"`
}

type DamageTenantUpdateRequest struct {
	Comment     *string      `json:"comment,omitempty"`
	Priority    *db.Priority `json:"priority,omitempty"`
	AddPictures []string     `binding:"dive,required,base64" json:"add_pictures,omitempty"`
}

type DamageResponse struct {
	ID         string `json:"id"`
	LeaseID    string `json:"lease_id"`
	TenantName string `json:"tenant_name"`
	RoomID     string `json:"room_id"`
	RoomName   string `json:"room_name"`

	Comment      string       `json:"comment"`
	Priority     db.Priority  `json:"priority"`
	Read         bool         `json:"read"`
	CreatedAt    db.DateTime  `json:"created_at"`
	UpdatedAt    db.DateTime  `json:"updated_at"`
	FixStatus    db.FixStatus `json:"fix_status"`
	FixPlannedAt *db.DateTime `json:"fix_planned_at"`
	FixedAt      *db.DateTime `json:"fixed_at,omitempty"`

	Pictures []string `json:"pictures"`
}

func (i *DamageResponse) FromDbDamage(model db.DamageModel) {
	i.ID = model.ID
	i.LeaseID = model.LeaseID
	i.TenantName = model.Lease().Tenant().Name()
	i.RoomID = model.RoomID
	i.RoomName = model.Room().Name

	i.Comment = model.Comment
	i.Priority = model.Priority
	i.Read = model.Read
	i.CreatedAt = model.CreatedAt
	i.UpdatedAt = model.UpdatedAt

	i.FixStatus = model.FixStatus()
	i.FixPlannedAt = model.InnerDamage.FixPlannedAt
	i.FixedAt = model.InnerDamage.FixedAt

	for _, picture := range model.Pictures() {
		i.Pictures = append(i.Pictures, DbImageToResponse(picture).Data)
	}
}

func DbDamageToResponse(pc db.DamageModel) DamageResponse {
	var resp DamageResponse
	resp.FromDbDamage(pc)
	return resp
}
