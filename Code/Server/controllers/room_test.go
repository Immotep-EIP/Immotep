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
	"keyz/backend/models"
	"keyz/backend/prisma/db"
	"keyz/backend/router"
	"keyz/backend/services"
	"keyz/backend/services/database"
	"keyz/backend/utils"
)

func BuildTestRoom(id string, propertyId string) db.RoomModel {
	return db.RoomModel{
		InnerRoom: db.InnerRoom{
			ID:         id,
			Name:       "Test",
			Type:       db.RoomTypeOther,
			PropertyID: propertyId,
			Archived:   false,
		},
	}
}

func TestCreateRoom(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	room := BuildTestRoom("1", "1")
	m.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)
	m.Room.Expect(database.MockCreateRoom(c, room)).Returns(room)

	b, err := json.Marshal(room)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/v1/owner/properties/1/rooms/", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusCreated, w.Code)
	var resp models.IdResponse
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.JSONEq(t, resp.ID, room.ID)
}

func TestCreateRoom_MissingFields(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	m.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)

	room := BuildTestRoom("1", "1")
	room.Name = ""
	b, err := json.Marshal(room)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/v1/owner/properties/1/rooms/", bytes.NewReader(b))
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

func TestCreateRoom_AlreadyExists(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	room := BuildTestRoom("1", "1")
	m.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)
	m.Room.Expect(database.MockCreateRoom(c, room)).Errors(&protocol.UserFacingError{
		IsPanic:   false,
		ErrorCode: "P2002", // https://www.prisma.io/docs/orm/reference/error-reference#p2002
		Meta: protocol.Meta{
			Target: []any{"name"},
		},
		Message: "Unique constraint failed",
	})

	b, err := json.Marshal(room)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/v1/owner/properties/1/rooms/", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusConflict, w.Code)
	var resp utils.Error
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.RoomAlreadyExists, resp.Code)
}

func TestGetRoomsByProperty(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	rooms := []db.RoomModel{
		BuildTestRoom("1", "1"),
		BuildTestRoom("2", "1"),
	}
	m.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)
	m.Room.Expect(database.MockGetRoomsByPropertyID(c, false)).ReturnsMany(rooms)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/owner/properties/1/rooms/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	var resp []models.RoomResponse
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Len(t, resp, 2)
	assert.Equal(t, resp[0].ID, rooms[0].ID)
}

func TestGetRoomsByProperty_PropertyNotFound(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.Property.Expect(database.MockGetPropertyByID(c)).Errors(db.ErrNotFound)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/owner/properties/1/rooms/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	var errorResponse utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
	require.NoError(t, err)
	assert.Equal(t, utils.PropertyNotFound, errorResponse.Code)
}

func TestGetRoomByID(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	room := BuildTestRoom("1", "1")
	m.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)
	m.Room.Expect(database.MockGetRoomByID(c)).Returns(room)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/owner/properties/1/rooms/1/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	var resp models.RoomResponse
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.JSONEq(t, resp.ID, room.ID)
}

func TestGetRoomByID_RoomNotFound(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	m.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)
	m.Room.Expect(database.MockGetRoomByID(c)).Errors(db.ErrNotFound)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/owner/properties/1/rooms/1/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	var errorResponse utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
	require.NoError(t, err)
	assert.Equal(t, utils.RoomNotFound, errorResponse.Code)
}

func TestGetRoomByID_WrongProperty(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	room := BuildTestRoom("1", "2")
	m.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)
	m.Room.Expect(database.MockGetRoomByID(c)).Returns(room)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/owner/properties/1/rooms/1/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	var errorResponse utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
	require.NoError(t, err)
	assert.Equal(t, utils.RoomNotFound, errorResponse.Code)
}

func TestGetRoomByID_PropertyNotFound(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.Property.Expect(database.MockGetPropertyByID(c)).Errors(db.ErrNotFound)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/owner/properties/1/rooms/1/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	var errorResponse utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
	require.NoError(t, err)
	assert.Equal(t, utils.PropertyNotFound, errorResponse.Code)
}

func TestArchiveRoom(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	room := BuildTestRoom("1", "1")
	updatedRoom := room
	updatedRoom.Archived = true
	m.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)
	m.Room.Expect(database.MockGetRoomByID(c)).Returns(room)
	m.Room.Expect(database.MockArchiveRoom(c)).Returns(updatedRoom)

	b, err := json.Marshal(map[string]bool{"archive": true})
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/v1/owner/properties/1/rooms/1/archive/", bytes.NewReader(b))
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	var resp models.IdResponse
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.JSONEq(t, resp.ID, room.ID)
}

func TestArchiveRoom_NotFound(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	m.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)
	m.Room.Expect(database.MockGetRoomByID(c)).Errors(db.ErrNotFound)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/v1/owner/properties/1/rooms/1/archive/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)
	var resp utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.RoomNotFound, resp.Code)
}

func TestGetArchivedRoomsByProperty(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	rooms := []db.RoomModel{
		BuildTestRoom("1", "1"),
		BuildTestRoom("2", "1"),
	}
	m.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)
	m.Room.Expect(database.MockGetRoomsByPropertyID(c, true)).ReturnsMany(rooms)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/owner/properties/1/rooms/?archive=true", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	var resp []models.RoomResponse
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Len(t, resp, 2)
	assert.Equal(t, resp[0].ID, rooms[0].ID)
}

func TestGetArchivedRoomsByProperty_PropertyNotFound(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.Property.Expect(database.MockGetPropertyByID(c)).Errors(db.ErrNotFound)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/owner/properties/1/rooms/?archive=true", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	var errorResponse utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
	require.NoError(t, err)
	assert.Equal(t, utils.PropertyNotFound, errorResponse.Code)
}

func TestDeleteRoom(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	room := BuildTestRoom("1", "1")
	m.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)
	m.Room.Expect(database.MockGetRoomByID(c)).Returns(room)
	m.Room.Expect(database.MockDeleteRoom(c)).Returns(room)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/v1/owner/properties/1/rooms/1/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNoContent, w.Code)
}

func TestDeleteRoom_NotFound(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	m.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)
	m.Room.Expect(database.MockGetRoomByID(c)).Errors(db.ErrNotFound)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/v1/owner/properties/1/rooms/1/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	var errorResponse utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
	require.NoError(t, err)
	assert.Equal(t, utils.RoomNotFound, errorResponse.Code)
}

func TestDeleteRoom_PropertyNotFound(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.Property.Expect(database.MockGetPropertyByID(c)).Errors(db.ErrNotFound)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/v1/owner/properties/1/rooms/1/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	var errorResponse utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
	require.NoError(t, err)
	assert.Equal(t, utils.PropertyNotFound, errorResponse.Code)
}
