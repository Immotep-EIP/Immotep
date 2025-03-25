package models

import "immotep/backend/prisma/db"

type FurnitureRequest struct {
	Name     string `binding:"required" json:"name"`
	Quantity *int   `binding:"-"        json:"quantity,omitempty"`
}

func (r *FurnitureRequest) ToDbFurniture() db.FurnitureModel {
	if r.Quantity != nil {
		return db.FurnitureModel{
			InnerFurniture: db.InnerFurniture{
				Name:     r.Name,
				Quantity: *r.Quantity,
			},
		}
	}
	return db.FurnitureModel{
		InnerFurniture: db.InnerFurniture{
			Name:     r.Name,
			Quantity: 1,
		},
	}
}

type FurnitureResponse struct {
	ID         string `json:"id"`
	PropertyID string `json:"property_id"`
	RoomID     string `json:"room_id"`
	Name       string `json:"name"`
	Quantity   int    `json:"quantity"`
	Archived   bool   `json:"archived"`
}

func (i *FurnitureResponse) FromDbFurniture(model db.FurnitureModel) {
	i.ID = model.ID
	i.PropertyID = model.Room().PropertyID
	i.RoomID = model.RoomID
	i.Name = model.Name
	i.Quantity = model.Quantity
	i.Archived = model.Archived
}

func DbFurnitureToResponse(pc db.FurnitureModel) FurnitureResponse {
	var resp FurnitureResponse
	resp.FromDbFurniture(pc)
	return resp
}
