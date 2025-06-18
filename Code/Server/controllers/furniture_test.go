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

func BuildTestFurniture(id string, roomId string) db.FurnitureModel {
	return db.FurnitureModel{
		InnerFurniture: db.InnerFurniture{
			ID:       id,
			Name:     "Test",
			Quantity: 1,
			RoomID:   roomId,
			Archived: false,
		},
		RelationsFurniture: db.RelationsFurniture{
			Room: &db.RoomModel{},
		},
	}
}

func TestCreateFurniture(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	room := BuildTestRoom("1", "1")
	furniture := BuildTestFurniture("1", "1")
	m.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)
	m.Room.Expect(database.MockGetRoomByID(c)).Returns(room)
	m.Furniture.Expect(database.MockCreateFurniture(c, furniture)).Returns(furniture)

	b, err := json.Marshal(furniture)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/v1/owner/properties/1/rooms/1/furnitures/", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusCreated, w.Code)
	var resp models.IdResponse
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.JSONEq(t, resp.ID, furniture.ID)
}

func TestCreateFurniture_MissingFields(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	room := BuildTestRoom("1", "1")
	m.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)
	m.Room.Expect(database.MockGetRoomByID(c)).Returns(room)

	furniture := BuildTestFurniture("1", "1")
	furniture.Name = ""
	b, err := json.Marshal(furniture)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/v1/owner/properties/1/rooms/1/furnitures/", bytes.NewReader(b))
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

func TestCreateFurniture_AlreadyExists(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	room := BuildTestRoom("1", "1")
	furniture := BuildTestFurniture("1", "1")
	m.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)
	m.Room.Expect(database.MockGetRoomByID(c)).Returns(room)
	m.Furniture.Expect(database.MockCreateFurniture(c, furniture)).Errors(&protocol.UserFacingError{
		IsPanic:   false,
		ErrorCode: "P2002", // https://www.prisma.io/docs/orm/reference/error-reference#p2002
		Meta: protocol.Meta{
			Target: []any{"name"},
		},
		Message: "Unique constraint failed",
	})

	b, err := json.Marshal(furniture)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/v1/owner/properties/1/rooms/1/furnitures/", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusConflict, w.Code)
	var resp utils.Error
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.FurnitureAlreadyExists, resp.Code)
}

func TestGetFurnituresByRoom(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	room := BuildTestRoom("1", "1")
	furnitures := []db.FurnitureModel{
		BuildTestFurniture("1", "1"),
		BuildTestFurniture("2", "1"),
	}
	m.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)
	m.Room.Expect(database.MockGetRoomByID(c)).Returns(room)
	m.Furniture.Expect(database.MockGetFurnituresByRoomID(c, false)).ReturnsMany(furnitures)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/owner/properties/1/rooms/1/furnitures/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	var resp []models.FurnitureResponse
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Len(t, resp, 2)
	assert.Equal(t, resp[0].ID, furnitures[0].ID)
}

func TestGetFurnituresByRoom_RoomNotFound(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	m.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)
	m.Room.Expect(database.MockGetRoomByID(c)).Errors(db.ErrNotFound)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/owner/properties/1/rooms/1/furnitures/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)
	var resp utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.RoomNotFound, resp.Code)
}

func TestGetFurnituresByRoom_WrongProperty(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	room := BuildTestRoom("1", "2")
	m.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)
	m.Room.Expect(database.MockGetRoomByID(c)).Returns(room)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/owner/properties/1/rooms/1/furnitures/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)
	var resp utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.RoomNotFound, resp.Code)
}

func TestGetFurnituresByRoom_PropertyNotFound(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.Property.Expect(database.MockGetPropertyByID(c)).Errors(db.ErrNotFound)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/owner/properties/1/rooms/1/furnitures/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)
	var resp utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.PropertyNotFound, resp.Code)
}

func TestGetFurnitureById(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	room := BuildTestRoom("1", "1")
	furniture := BuildTestFurniture("1", "1")
	m.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)
	m.Room.Expect(database.MockGetRoomByID(c)).Returns(room)
	m.Furniture.Expect(database.MockGetFurnitureByID(c)).Returns(furniture)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/owner/properties/1/rooms/1/furnitures/1/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	var resp models.FurnitureResponse
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.JSONEq(t, resp.ID, furniture.ID)
}

func TestGetFurnitureById_NotFound(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	room := BuildTestRoom("1", "1")
	m.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)
	m.Room.Expect(database.MockGetRoomByID(c)).Returns(room)
	m.Furniture.Expect(database.MockGetFurnitureByID(c)).Errors(db.ErrNotFound)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/owner/properties/1/rooms/1/furnitures/1/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)
	var resp utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.FurnitureNotFound, resp.Code)
}

func TestGetFurnitureById_WrongRoom(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	room := BuildTestRoom("1", "1")
	furniture := BuildTestFurniture("1", "2")
	m.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)
	m.Room.Expect(database.MockGetRoomByID(c)).Returns(room)
	m.Furniture.Expect(database.MockGetFurnitureByID(c)).Returns(furniture)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/owner/properties/1/rooms/1/furnitures/1/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)
	var resp utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.FurnitureNotFound, resp.Code)
}

func TestGetFurnitureById_RoomNotFound(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	m.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)
	m.Room.Expect(database.MockGetRoomByID(c)).Errors(db.ErrNotFound)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/owner/properties/1/rooms/1/furnitures/1/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)
	var resp utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.RoomNotFound, resp.Code)
}

func TestGetFurnitureById_WrongProperty(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	room := BuildTestRoom("1", "2")
	m.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)
	m.Room.Expect(database.MockGetRoomByID(c)).Returns(room)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/owner/properties/1/rooms/1/furnitures/1/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)
	var resp utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.RoomNotFound, resp.Code)
}

func TestGetFurnitureById_PropertyNotFound(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.Property.Expect(database.MockGetPropertyByID(c)).Errors(db.ErrNotFound)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/owner/properties/1/rooms/1/furnitures/1/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)
	var resp utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.PropertyNotFound, resp.Code)
}

func TestArchiveFurniture(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	room := BuildTestRoom("1", "1")
	furniture := BuildTestFurniture("1", "1")
	m.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)
	m.Room.Expect(database.MockGetRoomByID(c)).Returns(room)
	m.Furniture.Expect(database.MockGetFurnitureByID(c)).Returns(furniture)

	updatedFurniture := furniture
	updatedFurniture.Archived = true
	m.Furniture.Expect(database.MockArchiveFurniture(c)).Returns(updatedFurniture)

	b, err := json.Marshal(map[string]bool{"archive": true})
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/v1/owner/properties/1/rooms/1/furnitures/1/archive/", bytes.NewReader(b))
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	var resp models.IdResponse
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.JSONEq(t, resp.ID, furniture.ID)
}

func TestArchiveFurniture_MissingFields(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	room := BuildTestRoom("1", "1")
	furniture := BuildTestFurniture("1", "1")
	m.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)
	m.Room.Expect(database.MockGetRoomByID(c)).Returns(room)
	m.Furniture.Expect(database.MockGetFurnitureByID(c)).Returns(furniture)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/v1/owner/properties/1/rooms/1/furnitures/1/archive/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusBadRequest, w.Code)
	var resp utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.MissingFields, resp.Code)
}

func TestArchiveFurniture_NotFound(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	room := BuildTestRoom("1", "1")
	m.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)
	m.Room.Expect(database.MockGetRoomByID(c)).Returns(room)
	m.Furniture.Expect(database.MockGetFurnitureByID(c)).Errors(db.ErrNotFound)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/v1/owner/properties/1/rooms/1/furnitures/1/archive/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)
	var resp utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.FurnitureNotFound, resp.Code)
}

func TestArchiveFurniture_WrongRoom(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	room := BuildTestRoom("1", "1")
	furniture := BuildTestFurniture("1", "2")
	m.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)
	m.Room.Expect(database.MockGetRoomByID(c)).Returns(room)
	m.Furniture.Expect(database.MockGetFurnitureByID(c)).Returns(furniture)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/v1/owner/properties/1/rooms/1/furnitures/1/archive/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)
	var resp utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.FurnitureNotFound, resp.Code)
}

func TestArchiveFurniture_PropertyNotFound(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.Property.Expect(database.MockGetPropertyByID(c)).Errors(db.ErrNotFound)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/v1/owner/properties/1/rooms/1/furnitures/1/archive/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)
	var resp utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.PropertyNotFound, resp.Code)
}

func TestGetArchivedFurnituresByRoom(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	room := BuildTestRoom("1", "1")
	furnitures := []db.FurnitureModel{
		BuildTestFurniture("1", "1"),
		BuildTestFurniture("2", "1"),
	}
	for i := range furnitures {
		furnitures[i].Archived = true
	}
	m.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)
	m.Room.Expect(database.MockGetRoomByID(c)).Returns(room)
	m.Furniture.Expect(database.MockGetFurnituresByRoomID(c, true)).ReturnsMany(furnitures)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/owner/properties/1/rooms/1/furnitures/?archive=true", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	var resp []models.FurnitureResponse
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Len(t, resp, 2)
	assert.Equal(t, resp[0].ID, furnitures[0].ID)
	assert.True(t, resp[0].Archived)
}

func TestGetArchivedFurnituresByRoom_RoomNotFound(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	m.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)
	m.Room.Expect(database.MockGetRoomByID(c)).Errors(db.ErrNotFound)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/owner/properties/1/rooms/1/furnitures/?archive=true", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)
	var resp utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.RoomNotFound, resp.Code)
}

func TestGetArchivedFurnituresByRoom_WrongProperty(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	room := BuildTestRoom("1", "2")
	m.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)
	m.Room.Expect(database.MockGetRoomByID(c)).Returns(room)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/owner/properties/1/rooms/1/furnitures/?archive=true", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)
	var resp utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.RoomNotFound, resp.Code)
}

func TestGetArchivedFurnituresByRoom_PropertyNotFound(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.Property.Expect(database.MockGetPropertyByID(c)).Errors(db.ErrNotFound)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/owner/properties/1/rooms/1/furnitures/?archive=true", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)
	var resp utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.PropertyNotFound, resp.Code)
}

func TestDeleteFurniture(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	room := BuildTestRoom("1", "1")
	furniture := BuildTestFurniture("1", "1")
	m.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)
	m.Room.Expect(database.MockGetRoomByID(c)).Returns(room)
	m.Furniture.Expect(database.MockGetFurnitureByID(c)).Returns(furniture)
	m.Furniture.Expect(database.MockDeleteFurniture(c)).Returns(furniture)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/v1/owner/properties/1/rooms/1/furnitures/1/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNoContent, w.Code)
}

func TestDeleteFurniture_NotFound(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	room := BuildTestRoom("1", "1")
	m.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)
	m.Room.Expect(database.MockGetRoomByID(c)).Returns(room)
	m.Furniture.Expect(database.MockGetFurnitureByID(c)).Errors(db.ErrNotFound)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/v1/owner/properties/1/rooms/1/furnitures/1/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)
	var resp utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.FurnitureNotFound, resp.Code)
}

func TestDeleteFurniture_RoomNotFound(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	m.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)
	m.Room.Expect(database.MockGetRoomByID(c)).Errors(db.ErrNotFound)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/v1/owner/properties/1/rooms/1/furnitures/1/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)
	var resp utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.RoomNotFound, resp.Code)
}

func TestDeleteFurniture_PropertyNotFound(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.Property.Expect(database.MockGetPropertyByID(c)).Errors(db.ErrNotFound)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/v1/owner/properties/1/rooms/1/furnitures/1/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)
	var resp utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.PropertyNotFound, resp.Code)
}
