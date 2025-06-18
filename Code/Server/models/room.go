package models

import "keyz/backend/prisma/db"

type RoomRequest struct {
	Name string      `binding:"required"          json:"name"`
	Type db.RoomType `binding:"required,roomType" json:"type"`
}

func (r *RoomRequest) ToDbRoom() db.RoomModel {
	return db.RoomModel{
		InnerRoom: db.InnerRoom{
			Name: r.Name,
			Type: r.Type,
		},
	}
}

type RoomResponse struct {
	ID         string      `json:"id"`
	PropertyID string      `json:"property_id"`
	Name       string      `json:"name"`
	Type       db.RoomType `json:"type"`
	Archived   bool        `json:"archived"`
}

func (i *RoomResponse) FromDbRoom(model db.RoomModel) {
	i.ID = model.ID
	i.PropertyID = model.PropertyID
	i.Name = model.Name
	i.Type = model.Type
	i.Archived = model.Archived
}

func DbRoomToResponse(pc db.RoomModel) RoomResponse {
	var resp RoomResponse
	resp.FromDbRoom(pc)
	return resp
}
