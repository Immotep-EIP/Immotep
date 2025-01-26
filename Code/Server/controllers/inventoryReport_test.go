package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/steebchen/prisma-client-go/engine/protocol"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"immotep/backend/database"
	"immotep/backend/models"
	"immotep/backend/prisma/db"
	"immotep/backend/router"
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
				Pictures:    []string{"b3Vp"},
				Furnitures: []models.FurnitureStateRequest{
					{
						ID:          "1",
						State:       "good",
						Cleanliness: "clean",
						Note:        "Test note",
						Pictures:    []string{"bm9u"},
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
								BuildTestImage("1", "b3Vp"),
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
								BuildTestImage("2", "bm9u"),
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

func TestCreateInventoryReport(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(
		client.Client.Property.FindUnique(
			db.Property.ID.Equals(property.ID),
		).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
		),
	).Returns(property)

	invReport := BuildTestInvReport("1", "1", false)
	mock.InventoryReport.Expect(
		client.Client.InventoryReport.CreateOne(
			db.InventoryReport.Type.Set(invReport.Type),
			db.InventoryReport.Property.Link(db.Property.ID.Equals(invReport.PropertyID)),
		),
	).Returns(invReport)

	room := BuildTestRoom("1", "1")
	mock.Room.Expect(
		client.Client.Room.FindUnique(
			db.Room.ID.Equals(room.ID),
		),
	).Returns(room)

	roomPicture := BuildTestImage("1", "b3Vp")
	mock.Image.Expect(
		client.Client.Image.CreateOne(
			db.Image.Data.Set(roomPicture.Data),
		),
	).Returns(roomPicture)

	picturesId := []string{roomPicture.ID}
	roomParams := make([]db.RoomStateSetParam, 0, len(picturesId))
	for _, id := range picturesId {
		roomParams = append(roomParams, db.RoomState.Pictures.Link(db.Image.ID.Equals(id)))
	}
	roomState := BuildTestRoomState("1", "1", "1")
	mock.RoomState.Expect(
		client.Client.RoomState.CreateOne(
			db.RoomState.Cleanliness.Set(roomState.Cleanliness),
			db.RoomState.State.Set(roomState.State),
			db.RoomState.Note.Set(roomState.Note),
			db.RoomState.Report.Link(db.InventoryReport.ID.Equals(invReport.ID)),
			db.RoomState.Room.Link(db.Room.ID.Equals(roomState.RoomID)),
			roomParams...,
		),
	).Returns(roomState)

	furniture := BuildTestFurniture("1", "1")
	mock.Furniture.Expect(
		client.Client.Furniture.FindUnique(
			db.Furniture.ID.Equals(furniture.ID),
		).With(
			db.Furniture.Room.Fetch(),
		),
	).Returns(furniture)

	furniturePicture := BuildTestImage("2", "bm9u")
	mock.Image.Expect(
		client.Client.Image.CreateOne(
			db.Image.Data.Set(furniturePicture.Data),
		),
	).Returns(furniturePicture)

	picturesId = []string{furniturePicture.ID}
	furnitureParams := make([]db.FurnitureStateSetParam, 0, len(picturesId))
	for _, id := range picturesId {
		furnitureParams = append(furnitureParams, db.FurnitureState.Pictures.Link(db.Image.ID.Equals(id)))
	}
	furnitureState := BuildTestFurnitureState("1", "1", "1")
	mock.FurnitureState.Expect(
		client.Client.FurnitureState.CreateOne(
			db.FurnitureState.Cleanliness.Set(furnitureState.Cleanliness),
			db.FurnitureState.State.Set(furnitureState.State),
			db.FurnitureState.Note.Set(furnitureState.Note),
			db.FurnitureState.Report.Link(db.InventoryReport.ID.Equals(invReport.ID)),
			db.FurnitureState.Furniture.Link(db.Furniture.ID.Equals(furnitureState.FurnitureID)),
			furnitureParams...,
		),
	).Returns(furnitureState)

	reqBody := BuildInvReportRequest()
	b, err := json.Marshal(reqBody)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/owner/properties/1/inventory-reports/", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusCreated, w.Code)
	var resp []string
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Empty(t, resp)
}

func TestCreateInventoryReport_MissingFields(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(
		client.Client.Property.FindUnique(
			db.Property.ID.Equals(property.ID),
		).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
		),
	).Returns(property)

	reqBody := BuildInvReportRequest()
	reqBody.Rooms[0].Furnitures[0].Pictures = []string{}
	b, err := json.Marshal(reqBody)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/owner/properties/1/inventory-reports/", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusBadRequest, w.Code)
	var resp utils.Error
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.MissingFields, resp.Code)
}

func TestCreateInventoryReport_AlreadyExists(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(
		client.Client.Property.FindUnique(
			db.Property.ID.Equals(property.ID),
		).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
		),
	).Returns(property)

	invReport := db.InventoryReportModel{
		InnerInventoryReport: db.InnerInventoryReport{
			ID:         "1",
			Type:       db.ReportTypeStart,
			PropertyID: "1",
		},
	}
	mock.InventoryReport.Expect(
		client.Client.InventoryReport.CreateOne(
			db.InventoryReport.Type.Set(invReport.Type),
			db.InventoryReport.Property.Link(db.Property.ID.Equals(invReport.PropertyID)),
		),
	).Errors(&protocol.UserFacingError{
		IsPanic:   false,
		ErrorCode: "P2002", // https://www.prisma.io/docs/orm/reference/error-reference#p2002
		Meta: protocol.Meta{
			Target: []any{"type", "property_id"},
		},
		Message: "Unique constraint failed",
	})

	reqBody := BuildInvReportRequest()
	b, err := json.Marshal(reqBody)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/owner/properties/1/inventory-reports/", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusConflict, w.Code)
	var resp utils.Error
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.InventoryReportAlreadyExists, resp.Code)
}

func TestCreateInventoryReport_PropertyNotFound(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	mock.Property.Expect(
		client.Client.Property.FindUnique(
			db.Property.ID.Equals("1"),
		).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
		),
	).Errors(db.ErrNotFound)

	reqBody := BuildInvReportRequest()
	b, err := json.Marshal(reqBody)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/owner/properties/1/inventory-reports/", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)
	var resp utils.Error
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.PropertyNotFound, resp.Code)
}

func TestGetInventoryReportsByProperty(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(
		client.Client.Property.FindUnique(
			db.Property.ID.Equals(property.ID),
		).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
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
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/owner/properties/1/inventory-reports/", nil)
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
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	mock.Property.Expect(
		client.Client.Property.FindUnique(
			db.Property.ID.Equals("1"),
		).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
		),
	).Errors(db.ErrNotFound)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/owner/properties/1/inventory-reports/", nil)
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
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(
		client.Client.Property.FindUnique(
			db.Property.ID.Equals(property.ID),
		).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
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
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/owner/properties/1/inventory-reports/1/", nil)
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
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(
		client.Client.Property.FindUnique(
			db.Property.ID.Equals(property.ID),
		).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
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
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/owner/properties/1/inventory-reports/latest/", nil)
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
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(
		client.Client.Property.FindUnique(
			db.Property.ID.Equals(property.ID),
		).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
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
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/owner/properties/1/inventory-reports/1/", nil)
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

func TestCreateInventoryReport_RoomNotFound(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(
		client.Client.Property.FindUnique(
			db.Property.ID.Equals(property.ID),
		).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
		),
	).Returns(property)

	invReport := BuildTestInvReport("1", "1", false)
	mock.InventoryReport.Expect(
		client.Client.InventoryReport.CreateOne(
			db.InventoryReport.Type.Set(invReport.Type),
			db.InventoryReport.Property.Link(db.Property.ID.Equals(invReport.PropertyID)),
		),
	).Returns(invReport)

	room := BuildTestRoom("1", "1")
	mock.Room.Expect(
		client.Client.Room.FindUnique(
			db.Room.ID.Equals(room.ID),
		),
	).Errors(db.ErrNotFound)

	reqBody := BuildInvReportRequest()
	b, err := json.Marshal(reqBody)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/owner/properties/1/inventory-reports/", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusCreated, w.Code)
	var resp []string
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	require.Len(t, resp, 1)
	assert.Equal(t, string(utils.RoomNotFound), resp[0])
}

func TestCreateInventoryReport_FurnitureNotFound(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(
		client.Client.Property.FindUnique(
			db.Property.ID.Equals(property.ID),
		).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
		),
	).Returns(property)

	invReport := BuildTestInvReport("1", "1", false)
	mock.InventoryReport.Expect(
		client.Client.InventoryReport.CreateOne(
			db.InventoryReport.Type.Set(invReport.Type),
			db.InventoryReport.Property.Link(db.Property.ID.Equals(invReport.PropertyID)),
		),
	).Returns(invReport)

	room := BuildTestRoom("1", "1")
	mock.Room.Expect(
		client.Client.Room.FindUnique(
			db.Room.ID.Equals(room.ID),
		),
	).Returns(room)

	roomPicture := BuildTestImage("1", "b3Vp")
	mock.Image.Expect(
		client.Client.Image.CreateOne(
			db.Image.Data.Set(roomPicture.Data),
		),
	).Returns(roomPicture)

	picturesId := []string{roomPicture.ID}
	roomParams := make([]db.RoomStateSetParam, 0, len(picturesId))
	for _, id := range picturesId {
		roomParams = append(roomParams, db.RoomState.Pictures.Link(db.Image.ID.Equals(id)))
	}
	roomState := BuildTestRoomState("1", "1", "1")
	mock.RoomState.Expect(
		client.Client.RoomState.CreateOne(
			db.RoomState.Cleanliness.Set(roomState.Cleanliness),
			db.RoomState.State.Set(roomState.State),
			db.RoomState.Note.Set(roomState.Note),
			db.RoomState.Report.Link(db.InventoryReport.ID.Equals(invReport.ID)),
			db.RoomState.Room.Link(db.Room.ID.Equals(roomState.RoomID)),
			roomParams...,
		),
	).Returns(roomState)

	furniture := BuildTestFurniture("1", "1")
	mock.Furniture.Expect(
		client.Client.Furniture.FindUnique(
			db.Furniture.ID.Equals(furniture.ID),
		).With(
			db.Furniture.Room.Fetch(),
		),
	).Errors(db.ErrNotFound)

	reqBody := BuildInvReportRequest()
	b, err := json.Marshal(reqBody)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/owner/properties/1/inventory-reports/", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusCreated, w.Code)
	var resp []string
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	require.Len(t, resp, 1)
	assert.Equal(t, string(utils.FurnitureNotFound), resp[0])
}
