package models

import (
	"immotep/backend/prisma/db"
	"immotep/backend/utils"
)

type FurnitureStateRequest struct {
	FurnitureID string         `binding:"required"             json:"furniture_id"`
	State       db.State       `binding:"required,state"       json:"state"`
	Cleanliness db.Cleanliness `binding:"required,cleanliness" json:"cleanliness"`
	Note        string         `binding:"required"             json:"note"`
}

func (f *FurnitureStateRequest) ToDbFurnitureState() db.FurnitureStateModel {
	return db.FurnitureStateModel{
		InnerFurnitureState: db.InnerFurnitureState{
			FurnitureID: f.FurnitureID,
			State:       f.State,
			Cleanliness: f.Cleanliness,
			Note:        f.Note,
		},
	}
}


type RoomStateRequest struct {
	RoomID      string         `binding:"required"             json:"room_id"`
	State       db.State       `binding:"required,state"       json:"state"`
	Cleanliness db.Cleanliness `binding:"required,cleanliness" json:"cleanliness"`
	Note        string         `binding:"required"             json:"note"`
}

func (r *RoomStateRequest) ToDbRoomState() db.RoomStateModel {
	return db.RoomStateModel{
		InnerRoomState: db.InnerRoomState{
			RoomID:      r.RoomID,
			State:       r.State,
			Cleanliness: r.Cleanliness,
			Note:        r.Note,
		},
	}
}

type InventoryReportRequest struct {
	Type db.ReportType `binding:"required,reportType" json:"type"`
}

type FurnitureStateResponse struct {
	ID          string         `json:"id"`
	Name        string         `json:"name"`
	Quantity    int            `json:"quantity"`
	State       db.State       `json:"state"`
	Cleanliness db.Cleanliness `json:"cleanliness"`
	Note        string         `json:"note"`
	PictureUrls []string       `json:"picture_urls"`
}

type RoomStateResponse struct {
	ID          string                   `json:"id"`
	Name        string                   `json:"name"`
	State       db.State                 `json:"state"`
	Cleanliness db.Cleanliness           `json:"cleanliness"`
	Note        string                   `json:"note"`
	PictureUrls []string                 `json:"picture_urls"`
	Furnitures  []FurnitureStateResponse `json:"furnitures"`
}

type InventoryReportResponse struct {
	ID         string              `json:"id"`
	PropertyID string              `json:"property_id"`
	LeaseID    string              `json:"lease_id"`
	Date       db.DateTime         `json:"date"`
	Type       db.ReportType       `json:"type"`
	Rooms      []RoomStateResponse `json:"rooms"`
}

func (i *InventoryReportResponse) FromDbInventoryReport(model db.InventoryReportModel, pictureLinks map[string]string) {
	i.ID = model.ID
	i.PropertyID = model.Lease().PropertyID
	i.LeaseID = model.LeaseID
	i.Date = model.Date
	i.Type = model.Type
	for _, rs := range model.RoomStates() {
		var r RoomStateResponse
		r.ID = rs.RoomID
		r.Name = rs.Room().Name
		r.State = rs.State
		r.Cleanliness = rs.Cleanliness
		r.Note = rs.Note
		r.PictureUrls = utils.Map(rs.Pictures, func(path string) string { return pictureLinks[path] })
		addFurnitureStatesToRoomState(model, rs, &r, pictureLinks)
		i.Rooms = append(i.Rooms, r)
	}
}

func addFurnitureStatesToRoomState(model db.InventoryReportModel, room db.RoomStateModel, r *RoomStateResponse, pictureLinks map[string]string) {
	for _, fs := range model.FurnitureStates() {
		if fs.Furniture().RoomID != room.RoomID {
			continue
		}
		var f FurnitureStateResponse
		f.ID = fs.FurnitureID
		f.Name = fs.Furniture().Name
		f.Quantity = fs.Furniture().Quantity
		f.State = fs.State
		f.Cleanliness = fs.Cleanliness
		f.Note = fs.Note
		f.PictureUrls = utils.Map(fs.Pictures, func(path string) string { return pictureLinks[path] })
		r.Furnitures = append(r.Furnitures, f)
	}
}

func DbInventoryReportToResponse(pc db.InventoryReportModel, pictureLinks map[string]string) InventoryReportResponse {
	var resp InventoryReportResponse
	resp.FromDbInventoryReport(pc, pictureLinks)
	return resp
}

type CreateInventoryReportResponse struct {
	ID         string        `json:"id"`
	PropertyID string        `json:"property_id"`
	LeaseID    string        `json:"lease_id"`
	Date       db.DateTime   `json:"date"`
	Type       db.ReportType `json:"type"`
	PdfName    string        `json:"pdf_name"`
	PdfLink    string        `json:"pdf_link"`
}

func (c *CreateInventoryReportResponse) FromDbInventoryReport(model db.InventoryReportModel, pdf DocumentResponse) {
	c.ID = model.ID
	c.PropertyID = model.Lease().PropertyID
	c.LeaseID = model.LeaseID
	c.Date = model.Date
	c.Type = model.Type
	c.PdfName = pdf.Name
	c.PdfLink = pdf.Link
}

func DbInventoryReportToCreateResponse(pc db.InventoryReportModel, pdf DocumentResponse) CreateInventoryReportResponse {
	var resp CreateInventoryReportResponse
	resp.FromDbInventoryReport(pc, pdf)
	return resp
}
