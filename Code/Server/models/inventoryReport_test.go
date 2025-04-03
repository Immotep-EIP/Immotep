package models_test

import (
	"encoding/base64"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"immotep/backend/models"
	"immotep/backend/prisma/db"
)

func TestInventoryReport(t *testing.T) {
	model := db.InventoryReportModel{
		InnerInventoryReport: db.InnerInventoryReport{
			ID:      "1",
			LeaseID: "1",
			Date:    time.Now(),
			Type:    db.ReportTypeStart,
		},
		RelationsInventoryReport: db.RelationsInventoryReport{
			Lease: &db.LeaseModel{
				InnerLease: db.InnerLease{
					ID:         "1",
					PropertyID: "1",
				},
			},
			RoomStates: []db.RoomStateModel{
				{
					InnerRoomState: db.InnerRoomState{
						RoomID:      "1",
						State:       db.StateGood,
						Cleanliness: db.CleanlinessClean,
						Note:        "Room is in good condition",
					},
					RelationsRoomState: db.RelationsRoomState{
						Pictures: []db.ImageModel{
							{
								InnerImage: db.InnerImage{
									Data: []byte("base64image1"),
								},
							},
						},
						Room: &db.RoomModel{
							InnerRoom: db.InnerRoom{
								ID:   "1",
								Name: "Living Room",
							},
						},
					},
				},
				{
					InnerRoomState: db.InnerRoomState{
						RoomID:      "2",
						State:       db.StateGood,
						Cleanliness: db.CleanlinessClean,
						Note:        "Room is in good condition",
					},
					RelationsRoomState: db.RelationsRoomState{
						Pictures: []db.ImageModel{
							{
								InnerImage: db.InnerImage{
									Data: []byte("base64image1"),
								},
							},
						},
						Room: &db.RoomModel{
							InnerRoom: db.InnerRoom{
								ID:   "2",
								Name: "Kitchen",
							},
						},
					},
				},
			},
			FurnitureStates: []db.FurnitureStateModel{
				{
					InnerFurnitureState: db.InnerFurnitureState{
						FurnitureID: "1",
						State:       db.StateGood,
						Cleanliness: db.CleanlinessClean,
						Note:        "Furniture is in good condition",
					},
					RelationsFurnitureState: db.RelationsFurnitureState{
						Pictures: []db.ImageModel{
							{
								InnerImage: db.InnerImage{
									Data: []byte("base64image2"),
								},
							},
						},
						Furniture: &db.FurnitureModel{
							InnerFurniture: db.InnerFurniture{
								ID:       "1",
								RoomID:   "2",
								Name:     "Sofa",
								Quantity: 1,
							},
						},
					},
				},
			},
		},
	}

	t.Run("FromDbInventoryReport", func(t *testing.T) {
		var resp models.InventoryReportResponse
		resp.FromDbInventoryReport(model)

		assert.Equal(t, model.ID, resp.ID)
		assert.Equal(t, model.LeaseID, resp.LeaseID)
		assert.Equal(t, model.Date, resp.Date)
		assert.Equal(t, model.Type, resp.Type)
		assert.Len(t, resp.Rooms, 2)

		assert.Empty(t, resp.Rooms[0].Furnitures)
		assert.Len(t, resp.Rooms[1].Furnitures, 1)

		room := resp.Rooms[1]
		assert.Equal(t, "2", room.ID)
		assert.Equal(t, "Kitchen", room.Name)
		assert.Equal(t, db.StateGood, room.State)
		assert.Equal(t, db.CleanlinessClean, room.Cleanliness)
		assert.Equal(t, "Room is in good condition", room.Note)
		assert.Len(t, room.Pictures, 1)
		assert.Equal(t, "YmFzZTY0aW1hZ2Ux", room.Pictures[0])
		assert.Len(t, room.Furnitures, 1)

		furniture := room.Furnitures[0]
		assert.Equal(t, "1", furniture.ID)
		assert.Equal(t, "Sofa", furniture.Name)
		assert.Equal(t, 1, furniture.Quantity)
		assert.Equal(t, db.StateGood, furniture.State)
		assert.Equal(t, db.CleanlinessClean, furniture.Cleanliness)
		assert.Equal(t, "Furniture is in good condition", furniture.Note)
		assert.Len(t, furniture.Pictures, 1)
		assert.Equal(t, "YmFzZTY0aW1hZ2Uy", furniture.Pictures[0])
	})

	t.Run("DbInventoryReportToResponse", func(t *testing.T) {
		resp := models.DbInventoryReportToResponse(model)

		assert.Equal(t, model.ID, resp.ID)
		assert.Equal(t, model.LeaseID, resp.LeaseID)
		assert.Equal(t, model.Date, resp.Date)
		assert.Equal(t, model.Type, resp.Type)
		assert.Len(t, resp.Rooms, 2)

		assert.Empty(t, resp.Rooms[0].Furnitures)
		assert.Len(t, resp.Rooms[1].Furnitures, 1)

		room := resp.Rooms[1]
		assert.Equal(t, "2", room.ID)
		assert.Equal(t, "Kitchen", room.Name)
		assert.Equal(t, db.StateGood, room.State)
		assert.Equal(t, db.CleanlinessClean, room.Cleanliness)
		assert.Equal(t, "Room is in good condition", room.Note)
		assert.Len(t, room.Pictures, 1)
		assert.Equal(t, "YmFzZTY0aW1hZ2Ux", room.Pictures[0])
		assert.Len(t, room.Furnitures, 1)

		furniture := room.Furnitures[0]
		assert.Equal(t, "1", furniture.ID)
		assert.Equal(t, "Sofa", furniture.Name)
		assert.Equal(t, 1, furniture.Quantity)
		assert.Equal(t, db.StateGood, furniture.State)
		assert.Equal(t, db.CleanlinessClean, furniture.Cleanliness)
		assert.Equal(t, "Furniture is in good condition", furniture.Note)
		assert.Len(t, furniture.Pictures, 1)
		assert.Equal(t, "YmFzZTY0aW1hZ2Uy", furniture.Pictures[0])
	})
}

func TestCreateInventoryReportResponse(t *testing.T) {
	model := db.InventoryReportModel{
		InnerInventoryReport: db.InnerInventoryReport{
			ID:      "1",
			LeaseID: "1",
			Date:    time.Now(),
			Type:    db.ReportTypeStart,
		},
		RelationsInventoryReport: db.RelationsInventoryReport{
			Lease: &db.LeaseModel{
				InnerLease: db.InnerLease{
					ID:         "1",
					PropertyID: "1",
				},
			},
		},
	}

	pdf := &db.DocumentModel{
		InnerDocument: db.InnerDocument{
			Name: "report.pdf",
			Data: []byte("pdfdata"),
		},
	}

	errors := []string{"error1", "error2"}

	t.Run("FromDbInventoryReport", func(t *testing.T) {
		var resp models.CreateInventoryReportResponse
		resp.FromDbInventoryReport(model, pdf, errors)

		assert.Equal(t, model.ID, resp.ID)
		assert.Equal(t, model.LeaseID, resp.LeaseID)
		assert.Equal(t, model.Date, resp.Date)
		assert.Equal(t, model.Type, resp.Type)
		assert.Equal(t, pdf.Name, resp.PdfName)
		assert.Equal(t, "data:application/pdf;base64,"+base64.StdEncoding.EncodeToString(pdf.Data), resp.PdfData)
		assert.Equal(t, errors, resp.Errors)
	})

	t.Run("FromDbInventoryReport with nil pdf", func(t *testing.T) {
		var resp models.CreateInventoryReportResponse
		resp.FromDbInventoryReport(model, nil, errors)

		assert.Equal(t, model.ID, resp.ID)
		assert.Equal(t, model.LeaseID, resp.LeaseID)
		assert.Equal(t, model.Date, resp.Date)
		assert.Equal(t, model.Type, resp.Type)
		assert.Empty(t, resp.PdfName)
		assert.Empty(t, resp.PdfData)
		assert.Equal(t, errors, resp.Errors)
	})

	t.Run("DbInventoryReportToCreateResponse", func(t *testing.T) {
		resp := models.DbInventoryReportToCreateResponse(model, pdf, errors)

		assert.Equal(t, model.ID, resp.ID)
		assert.Equal(t, model.LeaseID, resp.LeaseID)
		assert.Equal(t, model.Date, resp.Date)
		assert.Equal(t, model.Type, resp.Type)
		assert.Equal(t, pdf.Name, resp.PdfName)
		assert.Equal(t, "data:application/pdf;base64,"+base64.StdEncoding.EncodeToString(pdf.Data), resp.PdfData)
		assert.Equal(t, errors, resp.Errors)
	})

	t.Run("DbInventoryReportToCreateResponse with nil pdf", func(t *testing.T) {
		resp := models.DbInventoryReportToCreateResponse(model, nil, errors)

		assert.Equal(t, model.ID, resp.ID)
		assert.Equal(t, model.LeaseID, resp.LeaseID)
		assert.Equal(t, model.Date, resp.Date)
		assert.Equal(t, model.Type, resp.Type)
		assert.Empty(t, resp.PdfName)
		assert.Empty(t, resp.PdfData)
		assert.Equal(t, errors, resp.Errors)
	})
}
