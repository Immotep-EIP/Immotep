package models

import "immotep/backend/prisma/db"

type RoomRequest struct {
	Name string `binding:"required" json:"name"`
}

func (r *RoomRequest) ToDbRoom() db.RoomModel {
	return db.RoomModel{
		InnerRoom: db.InnerRoom{
			Name: r.Name,
		},
	}
}

type RoomResponse struct {
	ID         string `json:"id"`
	PropertyID string `json:"property_id"`
	Name       string `json:"name"`
}

func (i *RoomResponse) FromDbRoom(model db.RoomModel) {
	i.ID = model.ID
	i.PropertyID = model.PropertyID
	i.Name = model.Name
}

func DbRoomToResponse(pc db.RoomModel) RoomResponse {
	var resp RoomResponse
	resp.FromDbRoom(pc)
	return resp
}
