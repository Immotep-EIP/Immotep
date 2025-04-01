package models

import (
	"immotep/backend/prisma/db"
	"immotep/backend/utils"
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

type DamageUpdateEvent string

const (
	DamageUpdateEventFixPlanned DamageUpdateEvent = "fix_planned"
	DamageUpdateEventFixed      DamageUpdateEvent = "fixed"
	DamageUpdateEventRead       DamageUpdateEvent = "read"
)

type DamageOwnerUpdateRequest struct {
	Event        DamageUpdateEvent `binding:"required,damageUpdateEvent" json:"event"`
	FixPlannedAt *db.DateTime      `json:"fix_planned_at,omitempty"`
}

type DamageTenantUpdateRequest struct {
	Comment     *string      `json:"comment,omitempty"`
	Priority    *db.Priority `json:"priority,omitempty"`
	AddPictures []string     `json:"add_pictures,omitempty"`
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
	FixPlannedAt *db.DateTime `json:"fix_planned_at"`
	Fixed        bool         `json:"fixed"`
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

	fixedAt, fixed := model.FixedAt()
	i.Fixed = fixed
	i.FixedAt = utils.Ternary(fixed, &fixedAt, nil)
	i.FixPlannedAt = model.InnerDamage.FixPlannedAt

	for _, picture := range model.Pictures() {
		i.Pictures = append(i.Pictures, DbImageToResponse(picture).Data)
	}
}

func DbDamageToResponse(pc db.DamageModel) DamageResponse {
	var resp DamageResponse
	resp.FromDbDamage(pc)
	return resp
}

type DamageCreateResponse struct {
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
	FixPlannedAt *db.DateTime `json:"fix_planned_at"`
	Fixed        bool         `json:"fixed"`
	FixedAt      *db.DateTime `json:"fixed_at,omitempty"`

	Error string `json:"error,omitempty"`
}

func (i *DamageCreateResponse) FromDbDamage(model db.DamageModel, err error) {
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

	fixedAt, fixed := model.FixedAt()
	i.Fixed = fixed
	i.FixedAt = utils.Ternary(fixed, &fixedAt, nil)
	i.FixPlannedAt = model.InnerDamage.FixPlannedAt

	if err != nil {
		i.Error = err.Error()
	} else {
		i.Error = ""
	}
}

func DbDamageToCreateResponse(pc db.DamageModel, err error) DamageCreateResponse {
	var resp DamageCreateResponse
	resp.FromDbDamage(pc, err)
	return resp
}
