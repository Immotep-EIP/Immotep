package models

import "immotep/backend/prisma/db"

type FurnitureStateRequest struct {
	ID          string   `binding:"required"                                              json:"id"`
	State       string   `binding:"required,oneof=broken needsRepair bad medium good new" json:"state"`
	Cleanliness string   `binding:"required,oneof=dirty medium clean"                     json:"cleanliness"`
	Note        string   `binding:"required"                                              json:"note"`
	Pictures    []string `binding:"required,min=1,dive,required,base64"                   json:"pictures"`
}

type RoomStateRequest struct {
	ID          string                  `binding:"required"                                              json:"id"`
	State       string                  `binding:"required,oneof=broken needsRepair bad medium good new" json:"state"`
	Cleanliness string                  `binding:"required,oneof=dirty medium clean"                     json:"cleanliness"`
	Note        string                  `binding:"required"                                              json:"note"`
	Pictures    []string                `binding:"required,min=1,dive,required,base64"                   json:"pictures"`
	Furnitures  []FurnitureStateRequest `binding:"required,dive"                                         json:"furnitures"`
}

type InventoryReportRequest struct {
	Type  string             `binding:"required,oneof=start middle end" json:"type"`
	Rooms []RoomStateRequest `binding:"required,dive"                   json:"rooms"`
}

type FurnitureStateResponse struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Quantity    int      `json:"quantity"`
	State       string   `json:"state"`
	Cleanliness string   `json:"cleanliness"`
	Note        string   `json:"note"`
	Pictures    []string `json:"pictures"`
}

type RoomStateResponse struct {
	ID          string                   `json:"id"`
	Name        string                   `json:"name"`
	State       string                   `json:"state"`
	Cleanliness string                   `json:"cleanliness"`
	Note        string                   `json:"note"`
	Pictures    []string                 `json:"pictures"`
	Furnitures  []FurnitureStateResponse `json:"furnitures"`
}

type InventoryReportResponse struct {
	ID         string              `json:"id"`
	PropertyID string              `json:"property_id"`
	Date       db.DateTime         `json:"date"`
	Type       string              `json:"type"`
	Rooms      []RoomStateResponse `json:"rooms"`
}

func (i *InventoryReportResponse) FromDbInventoryReport(model db.InventoryReportModel) {
	i.ID = model.ID
	i.PropertyID = model.PropertyID
	i.Date = model.Date
	i.Type = string(model.Type)
	for _, room := range model.RoomStates() {
		var r RoomStateResponse
		r.ID = room.RoomID
		r.Name = room.Room().Name
		r.State = string(room.State)
		r.Cleanliness = string(room.Cleanliness)
		r.Note = room.Note
		for _, picture := range room.Pictures() {
			r.Pictures = append(r.Pictures, DbImageToResponse(picture).Data)
		}
		addFurnitureStatesToRoomState(model, room, &r)
		i.Rooms = append(i.Rooms, r)
	}
}

func addFurnitureStatesToRoomState(model db.InventoryReportModel, room db.RoomStateModel, r *RoomStateResponse) {
	for _, furniture := range model.FurnitureStates() {
		if furniture.Furniture().RoomID != room.RoomID {
			continue
		}
		var f FurnitureStateResponse
		f.ID = furniture.FurnitureID
		f.Name = furniture.Furniture().Name
		f.Quantity = furniture.Furniture().Quantity
		f.State = string(furniture.State)
		f.Cleanliness = string(furniture.Cleanliness)
		f.Note = furniture.Note
		for _, picture := range furniture.Pictures() {
			f.Pictures = append(f.Pictures, DbImageToResponse(picture).Data)
		}
		r.Furnitures = append(r.Furnitures, f)
	}
}

func DbInventoryReportToResponse(pc db.InventoryReportModel) InventoryReportResponse {
	var resp InventoryReportResponse
	resp.FromDbInventoryReport(pc)
	return resp
}
