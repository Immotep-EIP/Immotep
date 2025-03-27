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

func BuildTestContractWithInfo() db.ContractModel {
	return db.ContractModel{
		InnerContract: db.InnerContract{
			ID:         "1",
			PropertyID: "1",
			TenantID:   "1",
			Active:     true,
			CreatedAt:  time.Now(),
			StartDate:  time.Now(),
			EndDate:    utils.Ptr(time.Now().Add(time.Hour)),
		},
		RelationsContract: db.RelationsContract{
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

func TestGetInventoryReportsByProperty(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(
		client.Client.Property.FindUnique(
			db.Property.ID.Equals(property.ID),
		).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
			db.Property.PendingContract.Fetch(),
		),
	).Returns(property)

	invReports := []db.InventoryReportModel{
		BuildTestInvReport("1", "1", true),
		BuildTestInvReport("2", "1", true),
	}
	mock.InventoryReport.Expect(
		client.Client.InventoryReport.FindMany(
			db.InventoryReport.PropertyID.Equals("1"),
		).OrderBy(
			db.InventoryReport.Date.Order(db.SortOrderDesc),
		).With(
			db.InventoryReport.Property.Fetch(),
			db.InventoryReport.RoomStates.Fetch().With(db.RoomState.Room.Fetch()).With(db.RoomState.Pictures.Fetch()),
			db.InventoryReport.FurnitureStates.Fetch().With(db.FurnitureState.Furniture.Fetch()).With(db.FurnitureState.Pictures.Fetch()),
		),
	).ReturnsMany(invReports)

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
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	mock.Property.Expect(
		client.Client.Property.FindUnique(
			db.Property.ID.Equals("1"),
		).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
			db.Property.PendingContract.Fetch(),
		),
	).Errors(db.ErrNotFound)

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
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(
		client.Client.Property.FindUnique(
			db.Property.ID.Equals(property.ID),
		).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
			db.Property.PendingContract.Fetch(),
		),
	).Returns(property)

	invReport := BuildTestInvReport("1", "1", true)
	mock.InventoryReport.Expect(
		client.Client.InventoryReport.FindUnique(
			db.InventoryReport.ID.Equals(invReport.ID),
		).With(
			db.InventoryReport.Property.Fetch(),
			db.InventoryReport.RoomStates.Fetch().With(db.RoomState.Room.Fetch()).With(db.RoomState.Pictures.Fetch()),
			db.InventoryReport.FurnitureStates.Fetch().With(db.FurnitureState.Furniture.Fetch()).With(db.FurnitureState.Pictures.Fetch()),
		),
	).Returns(invReport)

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
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(
		client.Client.Property.FindUnique(
			db.Property.ID.Equals(property.ID),
		).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
			db.Property.PendingContract.Fetch(),
		),
	).Returns(property)

	invReport := BuildTestInvReport("1", "1", true)
	mock.InventoryReport.Expect(
		client.Client.InventoryReport.FindFirst(
			db.InventoryReport.PropertyID.Equals(property.ID),
		).OrderBy(
			db.InventoryReport.Date.Order(db.SortOrderDesc),
		).With(
			db.InventoryReport.Property.Fetch(),
			db.InventoryReport.RoomStates.Fetch().With(db.RoomState.Room.Fetch()).With(db.RoomState.Pictures.Fetch()),
			db.InventoryReport.FurnitureStates.Fetch().With(db.FurnitureState.Furniture.Fetch()).With(db.FurnitureState.Pictures.Fetch()),
		),
	).Returns(invReport)

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
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(
		client.Client.Property.FindUnique(
			db.Property.ID.Equals(property.ID),
		).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
			db.Property.PendingContract.Fetch(),
		),
	).Returns(property)

	mock.InventoryReport.Expect(
		client.Client.InventoryReport.FindUnique(
			db.InventoryReport.ID.Equals("1"),
		).With(
			db.InventoryReport.Property.Fetch(),
			db.InventoryReport.RoomStates.Fetch().With(db.RoomState.Room.Fetch()).With(db.RoomState.Pictures.Fetch()),
			db.InventoryReport.FurnitureStates.Fetch().With(db.FurnitureState.Furniture.Fetch()).With(db.FurnitureState.Pictures.Fetch()),
		),
	).Errors(db.ErrNotFound)

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
