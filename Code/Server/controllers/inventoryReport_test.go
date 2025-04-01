package controllers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"immotep/backend/models"
	"immotep/backend/prisma/db"
	"immotep/backend/router"
	"immotep/backend/services"
	"immotep/backend/services/database"
	"immotep/backend/utils"
)

func BuildInvReportRequest() models.InventoryReportRequest {
	return models.InventoryReportRequest{
		Type: "start",
		Rooms: []models.RoomStateRequest{
			{
				ID:          "1",
				State:       "good",
				Cleanliness: "clean",
				Note:        "Test note",
				Pictures:    []string{"/9j/4AAQSkZJRgABAQAAAQABAAD/2wCEAAEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQH/2wBDAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQH/wAARCAAFAAUDASIAAhEBAxEB/8QAFQABAQAAAAAAAAAAAAAAAAAAAAf/xAAUEAEAAAAAAAAAAAAAAAAAAAAA/9oADAMBAAIQAxAAAABP/8QAFBABAAAAAAAAAAAAAAAAAAAAAP/aAAgBAQABBQJ//8QAFBEBAAAAAAAAAAAAAAAAAAAAAP/aAAgBAwEBPwH/xAAUEQEAAAAAAAAAAAAAAAAAAAAA/9oACAECAQE/Ad//xAAUEAEAAAAAAAAAAAAAAAAAAAAA/9oACAEBAAY/Ah//xAAVEAEBAAAAAAAAAAAAAAAAAAAAEf/aAAgBAQABPyH/2gAMAwEAAgADAAAAEB//xAAVEQEBAAAAAAAAAAAAAAAAAAAAEf/aAAgBAwEBPxD/xAAVEQEBAAAAAAAAAAAAAAAAAAAAEf/aAAgBAgEBPxD/xAAVEQEBAAAAAAAAAAAAAAAAAAAAEf/aAAgBAQABPxD/2Q=="},
				Furnitures: []models.FurnitureStateRequest{
					{
						ID:          "1",
						State:       "good",
						Cleanliness: "clean",
						Note:        "Test note",
						Pictures:    []string{"/9j/4AAQSkZJRgABAQAAAQABAAD/2wCEAAEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQH/2wBDAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQH/wAARCAAFAAUDASIAAhEBAxEB/8QAFQABAQAAAAAAAAAAAAAAAAAAAAf/xAAUEAEAAAAAAAAAAAAAAAAAAAAA/9oADAMBAAIQAxAAAABP/8QAFBABAAAAAAAAAAAAAAAAAAAAAP/aAAgBAQABBQJ//8QAFBEBAAAAAAAAAAAAAAAAAAAAAP/aAAgBAwEBPwH/xAAUEQEAAAAAAAAAAAAAAAAAAAAA/9oACAECAQE/Ad//xAAUEAEAAAAAAAAAAAAAAAAAAAAA/9oACAEBAAY/Ah//xAAVEAEBAAAAAAAAAAAAAAAAAAAAEf/aAAgBAQABPyH/2gAMAwEAAgADAAAAEB//xAAVEQEBAAAAAAAAAAAAAAAAAAAAEf/aAAgBAwEBPxD/xAAVEQEBAAAAAAAAAAAAAAAAAAAAEf/aAAgBAgEBPxD/xAAVEQEBAAAAAAAAAAAAAAAAAAAAEf/aAAgBAQABPxD/2Q=="},
					},
				},
			},
		},
	}
}

func BuildTestLeaseWithInfo() db.LeaseModel {
	return db.LeaseModel{
		InnerLease: db.InnerLease{
			ID:         "1",
			PropertyID: "1",
			TenantID:   "1",
			Active:     true,
			CreatedAt:  time.Now(),
			StartDate:  time.Now(),
			EndDate:    utils.Ptr(time.Now().Add(time.Hour)),
		},
		RelationsLease: db.RelationsLease{
			Tenant: &db.UserModel{
				InnerUser: db.InnerUser{
					ID:        "1",
					Firstname: "Test",
					Lastname:  "Tenant",
					Email:     "tenant@example.com",
					Role:      db.RoleTenant,
				},
			},
			Property: &db.PropertyModel{
				InnerProperty: db.InnerProperty{
					ID:                  "1",
					Name:                "Test property",
					RentalPricePerMonth: 1000,
				},
				RelationsProperty: db.RelationsProperty{
					Owner: &db.UserModel{
						InnerUser: db.InnerUser{
							ID:        "1",
							Firstname: "Test",
							Lastname:  "Owner",
							Email:     "owner@example.com",
							Role:      db.RoleOwner,
						},
					},
				},
			},
		},
	}
}

func BuildTestInvReport(id string, propertyId string, withRelations bool) db.InventoryReportModel {
	if withRelations {
		testRoom := BuildTestRoom("1", "1")
		testFurniture := BuildTestFurniture("1", "1")
		return db.InventoryReportModel{
			InnerInventoryReport: db.InnerInventoryReport{
				ID:         id,
				Type:       db.ReportTypeStart,
				PropertyID: propertyId,
			},
			RelationsInventoryReport: db.RelationsInventoryReport{
				RoomStates: []db.RoomStateModel{
					{
						InnerRoomState: db.InnerRoomState{
							ID:          "1",
							RoomID:      "1",
							ReportID:    id,
							Cleanliness: db.CleanlinessClean,
							State:       db.StateGood,
							Note:        "Test note",
						},
						RelationsRoomState: db.RelationsRoomState{
							Pictures: []db.ImageModel{
								BuildTestImage("1", "/9j/4AAQSkZJRgABAQAAAQABAAD/2wCEAAEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQH/2wBDAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQH/wAARCAAFAAUDASIAAhEBAxEB/8QAFQABAQAAAAAAAAAAAAAAAAAAAAf/xAAUEAEAAAAAAAAAAAAAAAAAAAAA/9oADAMBAAIQAxAAAABP/8QAFBABAAAAAAAAAAAAAAAAAAAAAP/aAAgBAQABBQJ//8QAFBEBAAAAAAAAAAAAAAAAAAAAAP/aAAgBAwEBPwH/xAAUEQEAAAAAAAAAAAAAAAAAAAAA/9oACAECAQE/Ad//xAAUEAEAAAAAAAAAAAAAAAAAAAAA/9oACAEBAAY/Ah//xAAVEAEBAAAAAAAAAAAAAAAAAAAAEf/aAAgBAQABPyH/2gAMAwEAAgADAAAAEB//xAAVEQEBAAAAAAAAAAAAAAAAAAAAEf/aAAgBAwEBPxD/xAAVEQEBAAAAAAAAAAAAAAAAAAAAEf/aAAgBAgEBPxD/xAAVEQEBAAAAAAAAAAAAAAAAAAAAEf/aAAgBAQABPxD/2Q=="),
							},
							Room: &testRoom,
						},
					},
				},
				FurnitureStates: []db.FurnitureStateModel{
					{
						InnerFurnitureState: db.InnerFurnitureState{
							ID:          "1",
							FurnitureID: "1",
							ReportID:    id,
							Cleanliness: db.CleanlinessClean,
							State:       db.StateGood,
							Note:        "Test note",
						},
						RelationsFurnitureState: db.RelationsFurnitureState{
							Pictures: []db.ImageModel{
								BuildTestImage("2", "/9j/4AAQSkZJRgABAQAAAQABAAD/2wCEAAEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQH/2wBDAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQH/wAARCAAFAAUDASIAAhEBAxEB/8QAFQABAQAAAAAAAAAAAAAAAAAAAAf/xAAUEAEAAAAAAAAAAAAAAAAAAAAA/9oADAMBAAIQAxAAAABP/8QAFBABAAAAAAAAAAAAAAAAAAAAAP/aAAgBAQABBQJ//8QAFBEBAAAAAAAAAAAAAAAAAAAAAP/aAAgBAwEBPwH/xAAUEQEAAAAAAAAAAAAAAAAAAAAA/9oACAECAQE/Ad//xAAUEAEAAAAAAAAAAAAAAAAAAAAA/9oACAEBAAY/Ah//xAAVEAEBAAAAAAAAAAAAAAAAAAAAEf/aAAgBAQABPyH/2gAMAwEAAgADAAAAEB//xAAVEQEBAAAAAAAAAAAAAAAAAAAAEf/aAAgBAwEBPxD/xAAVEQEBAAAAAAAAAAAAAAAAAAAAEf/aAAgBAgEBPxD/xAAVEQEBAAAAAAAAAAAAAAAAAAAAEf/aAAgBAQABPxD/2Q=="),
							},
							Furniture: &testFurniture,
						},
					},
				},
			},
		}
	}
	return db.InventoryReportModel{
		InnerInventoryReport: db.InnerInventoryReport{
			ID:         id,
			Type:       db.ReportTypeStart,
			PropertyID: propertyId,
		},
	}
}

func BuildTestRoomState(id string, roomId string, reportId string) db.RoomStateModel {
	return db.RoomStateModel{
		InnerRoomState: db.InnerRoomState{
			ID:          id,
			RoomID:      roomId,
			ReportID:    reportId,
			Cleanliness: db.CleanlinessClean,
			State:       db.StateGood,
			Note:        "Test note",
		},
	}
}

func BuildTestFurnitureState(id string, furnitureId string, reportId string) db.FurnitureStateModel {
	return db.FurnitureStateModel{
		InnerFurnitureState: db.InnerFurnitureState{
			ID:          id,
			FurnitureID: furnitureId,
			ReportID:    reportId,
			Cleanliness: db.CleanlinessClean,
			State:       db.StateGood,
			Note:        "Test note",
		},
	}
}

func BuildTestImage(id string, base64data string) db.ImageModel {
	ret := models.StringToDbImage(base64data)
	if ret == nil {
		panic("Invalid base64 string")
	}
	ret.ID = id
	return *ret
}

// func TestCreateInventoryReport(t *testing.T) {
// 	c, m, ensure := services.ConnectDBTest()
// 	defer ensure(t)

// 	property := BuildTestProperty("1")
// 	m.Property.Expect(
// 		c.Client.Property.FindUnique(
// 			db.Property.ID.Equals(property.ID),
// 		).With(
// 			db.Property.Leases.Fetch().With(db.Lease.Tenant.Fetch()),
// 			db.Property.LeaseInvite.Fetch(),
// 		),
// 	).Returns(property)

// 	lease := BuildTestLeaseWithInfo()
// 	m.Lease.Expect(
// 		c.Client.Lease.FindMany(
// 			db.Lease.PropertyID.Equals(property.ID),
// 			db.Lease.Active.Equals(true),
// 		).With(
// 			db.Lease.Tenant.Fetch(),
// 			db.Lease.Property.Fetch().With(db.Property.Owner.Fetch()),
// 		),
// 	).ReturnsMany([]db.LeaseModel{lease})

// 	invReport1 := BuildTestInvReport("1", "1", false)
// 	m.InventoryReport.Expect(
// 		c.Client.InventoryReport.CreateOne(
// 			db.InventoryReport.Type.Set(invReport1.Type),
// 			db.InventoryReport.Property.Link(db.Property.ID.Equals(invReport1.PropertyID)),
// 		),
// 	).Returns(invReport1)

// 	room := BuildTestRoom("1", "1")
// 	m.Room.Expect(
// 		c.Client.Room.FindUnique(
// 			db.Room.ID.Equals(room.ID),
// 		),
// 	).Returns(room)

// 	roomPicture := BuildTestImage("1", "/9j/4AAQSkZJRgABAQAAAQABAAD/2wCEAAEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQH/2wBDAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQH/wAARCAAFAAUDASIAAhEBAxEB/8QAFQABAQAAAAAAAAAAAAAAAAAAAAf/xAAUEAEAAAAAAAAAAAAAAAAAAAAA/9oADAMBAAIQAxAAAABP/8QAFBABAAAAAAAAAAAAAAAAAAAAAP/aAAgBAQABBQJ//8QAFBEBAAAAAAAAAAAAAAAAAAAAAP/aAAgBAwEBPwH/xAAUEQEAAAAAAAAAAAAAAAAAAAAA/9oACAECAQE/Ad//xAAUEAEAAAAAAAAAAAAAAAAAAAAA/9oACAEBAAY/Ah//xAAVEAEBAAAAAAAAAAAAAAAAAAAAEf/aAAgBAQABPyH/2gAMAwEAAgADAAAAEB//xAAVEQEBAAAAAAAAAAAAAAAAAAAAEf/aAAgBAwEBPxD/xAAVEQEBAAAAAAAAAAAAAAAAAAAAEf/aAAgBAgEBPxD/xAAVEQEBAAAAAAAAAAAAAAAAAAAAEf/aAAgBAQABPxD/2Q==")
// 	m.Image.Expect(
// 		c.Client.Image.CreateOne(
// 			db.Image.Data.Set(roomPicture.Data),
// 		),
// 	).Returns(roomPicture)

// 	picturesId := []string{roomPicture.ID}
// 	roomParams := make([]db.RoomStateSetParam, 0, len(picturesId))
// 	for _, id := range picturesId {
// 		roomParams = append(roomParams, db.RoomState.Pictures.Link(db.Image.ID.Equals(id)))
// 	}
// 	roomState := BuildTestRoomState("1", "1", "1")
// 	m.RoomState.Expect(
// 		c.Client.RoomState.CreateOne(
// 			db.RoomState.Cleanliness.Set(roomState.Cleanliness),
// 			db.RoomState.State.Set(roomState.State),
// 			db.RoomState.Note.Set(roomState.Note),
// 			db.RoomState.Report.Link(db.InventoryReport.ID.Equals(invReport1.ID)),
// 			db.RoomState.Room.Link(db.Room.ID.Equals(roomState.RoomID)),
// 			roomParams...,
// 		),
// 	).Returns(roomState)

// 	furniture := BuildTestFurniture("1", "1")
// 	m.Furniture.Expect(
// 		c.Client.Furniture.FindUnique(
// 			db.Furniture.ID.Equals(furniture.ID),
// 		).With(
// 			db.Furniture.Room.Fetch(),
// 		),
// 	).Returns(furniture)

// 	// furniturePicture := BuildTestImage("2", "/9j/4AAQSkZJRgABAQAAAQABAAD/2wCEAAEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQH/2wBDAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQH/wAARCAAFAAUDASIAAhEBAxEB/8QAFQABAQAAAAAAAAAAAAAAAAAAAAf/xAAUEAEAAAAAAAAAAAAAAAAAAAAA/9oADAMBAAIQAxAAAABP/8QAFBABAAAAAAAAAAAAAAAAAAAAAP/aAAgBAQABBQJ//8QAFBEBAAAAAAAAAAAAAAAAAAAAAP/aAAgBAwEBPwH/xAAUEQEAAAAAAAAAAAAAAAAAAAAA/9oACAECAQE/Ad//xAAUEAEAAAAAAAAAAAAAAAAAAAAA/9oACAEBAAY/Ah//xAAVEAEBAAAAAAAAAAAAAAAAAAAAEf/aAAgBAQABPyH/2gAMAwEAAgADAAAAEB//xAAVEQEBAAAAAAAAAAAAAAAAAAAAEf/aAAgBAwEBPxD/xAAVEQEBAAAAAAAAAAAAAAAAAAAAEf/aAAgBAgEBPxD/xAAVEQEBAAAAAAAAAAAAAAAAAAAAEf/aAAgBAQABPxD/2Q==")
// 	// m.Image.Expect(
// 	// 	c.Client.Image.CreateOne(
// 	// 		db.Image.Data.Set(furniturePicture.Data),
// 	// 	),
// 	// ).Returns(furniturePicture)

// 	// picturesId = []string{furniturePicture.ID}
// 	picturesId = []string{}
// 	furnitureParams := make([]db.FurnitureStateSetParam, 0, len(picturesId))
// 	for _, id := range picturesId {
// 		furnitureParams = append(furnitureParams, db.FurnitureState.Pictures.Link(db.Image.ID.Equals(id)))
// 	}
// 	furnitureState := BuildTestFurnitureState("1", "1", "1")
// 	m.FurnitureState.Expect(
// 		c.Client.FurnitureState.CreateOne(
// 			db.FurnitureState.Cleanliness.Set(furnitureState.Cleanliness),
// 			db.FurnitureState.State.Set(furnitureState.State),
// 			db.FurnitureState.Note.Set(furnitureState.Note),
// 			db.FurnitureState.Report.Link(db.InventoryReport.ID.Equals(invReport1.ID)),
// 			db.FurnitureState.Furniture.Link(db.Furniture.ID.Equals(furnitureState.FurnitureID)),
// 			furnitureParams...,
// 		),
// 	).Returns(furnitureState)

// 	// invReport2 := BuildTestInvReport("1", "1", true)
// 	// m.InventoryReport.Expect(
// 	// 	c.Client.InventoryReport.FindUnique(
// 	// 		db.InventoryReport.ID.Equals(invReport2.ID),
// 	// 	).With(
// 	// 		db.InventoryReport.Property.Fetch(),
// 	// 		db.InventoryReport.RoomStates.Fetch().With(db.RoomState.Room.Fetch()).With(db.RoomState.Pictures.Fetch()),
// 	// 		db.InventoryReport.FurnitureStates.Fetch().With(db.FurnitureState.Furniture.Fetch()).With(db.FurnitureState.Pictures.Fetch()),
// 	// 	),
// 	// ).Returns(invReport2)

// 	// doc, err := pdf.NewInventoryReportPDF(invReport, lease)
// 	// require.NoError(t, err)

// 	// m.Document.Expect(
// 	// 	c.Client.Document.CreateOne(
// 	// 		db.Document.Name.Set("inventory_report_"+time.Now().Format("2006-01-02")+"_"+invReport.ID+".pdf"),
// 	// 		db.Document.Data.Set(doc),
// 	// 		db.Document.Lease.Link(db.Lease.ID.Equals(lease.ID)),
// 	// 	),
// 	// ).Returns(db.DocumentModel{
// 	// 	InnerDocument: db.InnerDocument{
// 	// 		ID:         "1",
// 	// 		Name:       "inventory_report_" + time.Now().Format("2006-01-02") + "_" + invReport.ID + ".pdf",
// 	// 		Data:       doc,
// 	// 		LeaseID: lease.ID,
// 	// 	},
// 	// })

// 	reqBody := BuildInvReportRequest()
// 	b, err := json.Marshal(reqBody)
// 	require.NoError(t, err)

// 	r := router.TestRoutes()
// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest(http.MethodPost, "/v1/owner/properties/1/inventory-reports/", bytes.NewReader(b))
// 	req.Header.Set("Content-Type", "application/json")
// 	req.Header.Set("Oauth.claims.id", "1")
// 	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
// 	r.ServeHTTP(w, req)

// 	require.Equal(t, http.StatusCreated, w.Code)
// 	var resp []string
// 	err = json.Unmarshal(w.Body.Bytes(), &resp)
// 	require.NoError(t, err)
// 	assert.Empty(t, resp)
// }

// func TestCreateInventoryReport_MissingFields(t *testing.T) {
// 	c, m, ensure := services.ConnectDBTest()
// 	defer ensure(t)

// 	property := BuildTestProperty("1")
// 	m.Property.Expect(
// 		c.Client.Property.FindUnique(
// 			db.Property.ID.Equals(property.ID),
// 		).With(
// 			db.Property.Leases.Fetch().With(db.Lease.Tenant.Fetch()),
// 			db.Property.LeaseInvite.Fetch(),
// 		),
// 	).Returns(property)

// 	reqBody := BuildInvReportRequest()
// 	reqBody.Rooms[0].Furnitures[0].Pictures = []string{}
// 	b, err := json.Marshal(reqBody)
// 	require.NoError(t, err)

// 	r := router.TestRoutes()
// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest(http.MethodPost, "/v1/owner/properties/1/inventory-reports/", bytes.NewReader(b))
// 	req.Header.Set("Content-Type", "application/json")
// 	req.Header.Set("Oauth.claims.id", "1")
// 	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
// 	r.ServeHTTP(w, req)

// 	require.Equal(t, http.StatusBadRequest, w.Code)
// 	var resp utils.Error
// 	err = json.Unmarshal(w.Body.Bytes(), &resp)
// 	require.NoError(t, err)
// 	assert.Equal(t, utils.MissingFields, resp.Code)
// }

// func TestCreateInventoryReport_AlreadyExists(t *testing.T) {
// 	c, m, ensure := services.ConnectDBTest()
// 	defer ensure(t)

// 	property := BuildTestProperty("1")
// 	m.Property.Expect(
// 		c.Client.Property.FindUnique(
// 			db.Property.ID.Equals(property.ID),
// 		).With(
// 			db.Property.Leases.Fetch().With(db.Lease.Tenant.Fetch()),
// 			db.Property.LeaseInvite.Fetch(),
// 		),
// 	).Returns(property)

// 	invReport := db.InventoryReportModel{
// 		InnerInventoryReport: db.InnerInventoryReport{
// 			ID:         "1",
// 			Type:       db.ReportTypeStart,
// 			PropertyID: "1",
// 		},
// 	}
// 	m.InventoryReport.Expect(
// 		c.Client.InventoryReport.CreateOne(
// 			db.InventoryReport.Type.Set(invReport.Type),
// 			db.InventoryReport.Property.Link(db.Property.ID.Equals(invReport.PropertyID)),
// 		),
// 	).Errors(&protocol.UserFacingError{
// 		IsPanic:   false,
// 		ErrorCode: "P2002", // https://www.prisma.io/docs/orm/reference/error-reference#p2002
// 		Meta: protocol.Meta{
// 			Target: []any{"type", "property_id"},
// 		},
// 		Message: "Unique constraint failed",
// 	})

// 	reqBody := BuildInvReportRequest()
// 	b, err := json.Marshal(reqBody)
// 	require.NoError(t, err)

// 	r := router.TestRoutes()
// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest(http.MethodPost, "/v1/owner/properties/1/inventory-reports/", bytes.NewReader(b))
// 	req.Header.Set("Content-Type", "application/json")
// 	req.Header.Set("Oauth.claims.id", "1")
// 	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
// 	r.ServeHTTP(w, req)

// 	require.Equal(t, http.StatusConflict, w.Code)
// 	var resp utils.Error
// 	err = json.Unmarshal(w.Body.Bytes(), &resp)
// 	require.NoError(t, err)
// 	assert.Equal(t, utils.InventoryReportAlreadyExists, resp.Code)
// }

// func TestCreateInventoryReport_PropertyNotFound(t *testing.T) {
// 	c, m, ensure := services.ConnectDBTest()
// 	defer ensure(t)

// 	m.Property.Expect(
// 		c.Client.Property.FindUnique(
// 			db.Property.ID.Equals("1"),
// 		).With(
// 			db.Property.Leases.Fetch().With(db.Lease.Tenant.Fetch()),
// 			db.Property.LeaseInvite.Fetch(),
// 		),
// 	).Errors(db.ErrNotFound)

// 	reqBody := BuildInvReportRequest()
// 	b, err := json.Marshal(reqBody)
// 	require.NoError(t, err)

// 	r := router.TestRoutes()
// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest(http.MethodPost, "/v1/owner/properties/1/inventory-reports/", bytes.NewReader(b))
// 	req.Header.Set("Content-Type", "application/json")
// 	req.Header.Set("Oauth.claims.id", "1")
// 	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
// 	r.ServeHTTP(w, req)

// 	require.Equal(t, http.StatusNotFound, w.Code)
// 	var resp utils.Error
// 	err = json.Unmarshal(w.Body.Bytes(), &resp)
// 	require.NoError(t, err)
// 	assert.Equal(t, utils.PropertyNotFound, resp.Code)
// }

// func TestCreateInventoryReport_RoomNotFound(t *testing.T) {
// 	c, m, ensure := services.ConnectDBTest()
// 	defer ensure(t)

// 	property := BuildTestProperty("1")
// 	m.Property.Expect(
// 		c.Client.Property.FindUnique(
// 			db.Property.ID.Equals(property.ID),
// 		).With(
// 			db.Property.Leases.Fetch().With(db.Lease.Tenant.Fetch()),
// 			db.Property.LeaseInvite.Fetch(),
// 		),
// 	).Returns(property)

// 	invReport := BuildTestInvReport("1", "1", false)
// 	m.InventoryReport.Expect(
// 		c.Client.InventoryReport.CreateOne(
// 			db.InventoryReport.Type.Set(invReport.Type),
// 			db.InventoryReport.Property.Link(db.Property.ID.Equals(invReport.PropertyID)),
// 		),
// 	).Returns(invReport)

// 	room := BuildTestRoom("1", "1")
// 	m.Room.Expect(
// 		c.Client.Room.FindUnique(
// 			db.Room.ID.Equals(room.ID),
// 		),
// 	).Errors(db.ErrNotFound)

// 	reqBody := BuildInvReportRequest()
// 	b, err := json.Marshal(reqBody)
// 	require.NoError(t, err)

// 	r := router.TestRoutes()
// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest(http.MethodPost, "/v1/owner/properties/1/inventory-reports/", bytes.NewReader(b))
// 	req.Header.Set("Content-Type", "application/json")
// 	req.Header.Set("Oauth.claims.id", "1")
// 	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
// 	r.ServeHTTP(w, req)

// 	require.Equal(t, http.StatusCreated, w.Code)
// 	var resp []string
// 	err = json.Unmarshal(w.Body.Bytes(), &resp)
// 	require.NoError(t, err)
// 	require.Len(t, resp, 1)
// 	assert.Equal(t, string(utils.RoomNotFound), resp[0])
// }

// func TestCreateInventoryReport_FurnitureNotFound(t *testing.T) {
// 	c, m, ensure := services.ConnectDBTest()
// 	defer ensure(t)

// 	property := BuildTestProperty("1")
// 	m.Property.Expect(
// 		c.Client.Property.FindUnique(
// 			db.Property.ID.Equals(property.ID),
// 		).With(
// 			db.Property.Leases.Fetch().With(db.Lease.Tenant.Fetch()),
// 			db.Property.LeaseInvite.Fetch(),
// 		),
// 	).Returns(property)

// 	invReport := BuildTestInvReport("1", "1", false)
// 	m.InventoryReport.Expect(
// 		c.Client.InventoryReport.CreateOne(
// 			db.InventoryReport.Type.Set(invReport.Type),
// 			db.InventoryReport.Property.Link(db.Property.ID.Equals(invReport.PropertyID)),
// 		),
// 	).Returns(invReport)

// 	room := BuildTestRoom("1", "1")
// 	m.Room.Expect(
// 		c.Client.Room.FindUnique(
// 			db.Room.ID.Equals(room.ID),
// 		),
// 	).Returns(room)

// 	roomPicture := BuildTestImage("1", "b3Vp")
// 	m.Image.Expect(
// 		c.Client.Image.CreateOne(
// 			db.Image.Data.Set(roomPicture.Data),
// 		),
// 	).Returns(roomPicture)

// 	picturesId := []string{roomPicture.ID}
// 	roomParams := make([]db.RoomStateSetParam, 0, len(picturesId))
// 	for _, id := range picturesId {
// 		roomParams = append(roomParams, db.RoomState.Pictures.Link(db.Image.ID.Equals(id)))
// 	}
// 	roomState := BuildTestRoomState("1", "1", "1")
// 	m.RoomState.Expect(
// 		c.Client.RoomState.CreateOne(
// 			db.RoomState.Cleanliness.Set(roomState.Cleanliness),
// 			db.RoomState.State.Set(roomState.State),
// 			db.RoomState.Note.Set(roomState.Note),
// 			db.RoomState.Report.Link(db.InventoryReport.ID.Equals(invReport.ID)),
// 			db.RoomState.Room.Link(db.Room.ID.Equals(roomState.RoomID)),
// 			roomParams...,
// 		),
// 	).Returns(roomState)

// 	furniture := BuildTestFurniture("1", "1")
// 	m.Furniture.Expect(
// 		c.Client.Furniture.FindUnique(
// 			db.Furniture.ID.Equals(furniture.ID),
// 		).With(
// 			db.Furniture.Room.Fetch(),
// 		),
// 	).Errors(db.ErrNotFound)

// 	reqBody := BuildInvReportRequest()
// 	b, err := json.Marshal(reqBody)
// 	require.NoError(t, err)

// 	r := router.TestRoutes()
// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest(http.MethodPost, "/v1/owner/properties/1/inventory-reports/", bytes.NewReader(b))
// 	req.Header.Set("Content-Type", "application/json")
// 	req.Header.Set("Oauth.claims.id", "1")
// 	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
// 	r.ServeHTTP(w, req)

// 	require.Equal(t, http.StatusCreated, w.Code)
// 	var resp []string
// 	err = json.Unmarshal(w.Body.Bytes(), &resp)
// 	require.NoError(t, err)
// 	require.Len(t, resp, 1)
// 	assert.Equal(t, string(utils.FurnitureNotFound), resp[0])
// }

func TestGetInventoryReportsByProperty(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	invReports := []db.InventoryReportModel{
		BuildTestInvReport("1", "1", true),
		BuildTestInvReport("2", "1", true),
	}
	m.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)
	m.InventoryReport.Expect(database.MockGetInvReportByPropertyID(c)).ReturnsMany(invReports)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/owner/properties/1/inventory-reports/", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	var resp []models.InventoryReportResponse
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Len(t, resp, 2)
	assert.Equal(t, invReports[0].ID, resp[0].ID)
}

func TestGetInventoryReportsByProperty_PropertyNotFound(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.Property.Expect(database.MockGetPropertyByID(c)).Errors(db.ErrNotFound)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/owner/properties/1/inventory-reports/", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)
	var resp utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.PropertyNotFound, resp.Code)
}

func TestGetInventoryReportByID(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	invReport := BuildTestInvReport("1", "1", true)
	m.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)
	m.InventoryReport.Expect(database.MockGetInvReportByID(c)).Returns(invReport)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/owner/properties/1/inventory-reports/1/", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	var resp models.InventoryReportResponse
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, invReport.ID, resp.ID)
}

func TestGetInventoryReportByID_Latest(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	invReport := BuildTestInvReport("1", "1", true)
	m.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)
	m.InventoryReport.Expect(database.MockGetLatestInvReport(c)).Returns(invReport)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/owner/properties/1/inventory-reports/latest/", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	var resp models.InventoryReportResponse
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, invReport.ID, resp.ID)
}

func TestGetInventoryReportByID_NotFound(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	m.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)
	m.InventoryReport.Expect(database.MockGetInvReportByID(c)).Errors(db.ErrNotFound)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/owner/properties/1/inventory-reports/1/", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)
	var resp utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.InventoryReportNotFound, resp.Code)
}
