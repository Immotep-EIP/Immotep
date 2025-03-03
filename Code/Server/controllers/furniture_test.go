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
	"immotep/backend/models"
	"immotep/backend/prisma/db"
	"immotep/backend/router"
	"immotep/backend/services"
	"immotep/backend/utils"
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

	room := BuildTestRoom("1", "1")
	mock.Room.Expect(
		client.Client.Room.FindUnique(
			db.Room.ID.Equals(room.ID),
		),
	).Returns(room)

	furniture := BuildTestFurniture("1", "1")
	mock.Furniture.Expect(
		client.Client.Furniture.CreateOne(
			db.Furniture.Name.Set(furniture.Name),
			db.Furniture.Room.Link(db.Room.ID.Equals(furniture.RoomID)),
			db.Furniture.Quantity.Set(furniture.Quantity),
		).With(
			db.Furniture.Room.Fetch(),
		),
	).Returns(furniture)

	b, err := json.Marshal(furniture)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/owner/properties/1/rooms/1/furnitures/", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusCreated, w.Code)
	var resp models.FurnitureResponse
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.JSONEq(t, resp.ID, furniture.ID)
}

func TestCreateFurniture_MissingFields(t *testing.T) {
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

	room := BuildTestRoom("1", "1")
	mock.Room.Expect(
		client.Client.Room.FindUnique(
			db.Room.ID.Equals(room.ID),
		),
	).Returns(room)

	furniture := BuildTestFurniture("1", "1")
	furniture.Name = ""
	b, err := json.Marshal(furniture)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/owner/properties/1/rooms/1/furnitures/", bytes.NewReader(b))
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

	room := BuildTestRoom("1", "1")
	mock.Room.Expect(
		client.Client.Room.FindUnique(
			db.Room.ID.Equals(room.ID),
		),
	).Returns(room)

	furniture := BuildTestFurniture("1", "1")
	mock.Furniture.Expect(
		client.Client.Furniture.CreateOne(
			db.Furniture.Name.Set(furniture.Name),
			db.Furniture.Room.Link(db.Room.ID.Equals(furniture.RoomID)),
			db.Furniture.Quantity.Set(furniture.Quantity),
		).With(
			db.Furniture.Room.Fetch(),
		),
	).Errors(&protocol.UserFacingError{
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
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/owner/properties/1/rooms/1/furnitures/", bytes.NewReader(b))
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

	room := BuildTestRoom("1", "1")
	mock.Room.Expect(
		client.Client.Room.FindUnique(
			db.Room.ID.Equals(room.ID),
		),
	).Returns(room)

	furnitures := []db.FurnitureModel{
		BuildTestFurniture("1", "1"),
		BuildTestFurniture("2", "1"),
	}
	mock.Furniture.Expect(
		client.Client.Furniture.FindMany(
			db.Furniture.RoomID.Equals(room.ID),
			db.Furniture.Archived.Equals(false),
		).With(
			db.Furniture.Room.Fetch(),
		),
	).ReturnsMany(furnitures)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/owner/properties/1/rooms/1/furnitures/", nil)
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

	mock.Room.Expect(
		client.Client.Room.FindUnique(
			db.Room.ID.Equals("1"),
		),
	).Errors(db.ErrNotFound)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/owner/properties/1/rooms/1/furnitures/", nil)
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

	room := BuildTestRoom("1", "2")
	mock.Room.Expect(
		client.Client.Room.FindUnique(
			db.Room.ID.Equals(room.ID),
		),
	).Returns(room)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/owner/properties/1/rooms/1/furnitures/", nil)
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
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/owner/properties/1/rooms/1/furnitures/", nil)
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

	room := BuildTestRoom("1", "1")
	mock.Room.Expect(
		client.Client.Room.FindUnique(
			db.Room.ID.Equals(room.ID),
		),
	).Returns(room)

	furniture := BuildTestFurniture("1", "1")
	mock.Furniture.Expect(
		client.Client.Furniture.FindUnique(
			db.Furniture.ID.Equals(furniture.ID),
		).With(
			db.Furniture.Room.Fetch(),
		),
	).Returns(furniture)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/owner/properties/1/rooms/1/furnitures/1/", nil)
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

	room := BuildTestRoom("1", "1")
	mock.Room.Expect(
		client.Client.Room.FindUnique(
			db.Room.ID.Equals(room.ID),
		),
	).Returns(room)

	mock.Furniture.Expect(
		client.Client.Furniture.FindUnique(
			db.Furniture.ID.Equals("1"),
		).With(
			db.Furniture.Room.Fetch(),
		),
	).Errors(db.ErrNotFound)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/owner/properties/1/rooms/1/furnitures/1/", nil)
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

	room := BuildTestRoom("1", "1")
	mock.Room.Expect(
		client.Client.Room.FindUnique(
			db.Room.ID.Equals(room.ID),
		),
	).Returns(room)

	furniture := BuildTestFurniture("1", "2")
	mock.Furniture.Expect(
		client.Client.Furniture.FindUnique(
			db.Furniture.ID.Equals(furniture.ID),
		).With(
			db.Furniture.Room.Fetch(),
		),
	).Returns(furniture)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/owner/properties/1/rooms/1/furnitures/1/", nil)
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

	mock.Room.Expect(
		client.Client.Room.FindUnique(
			db.Room.ID.Equals("1"),
		),
	).Errors(db.ErrNotFound)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/owner/properties/1/rooms/1/furnitures/1/", nil)
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

	room := BuildTestRoom("1", "2")
	mock.Room.Expect(
		client.Client.Room.FindUnique(
			db.Room.ID.Equals(room.ID),
		),
	).Returns(room)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/owner/properties/1/rooms/1/furnitures/1/", nil)
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
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/owner/properties/1/rooms/1/furnitures/1/", nil)
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

	room := BuildTestRoom("1", "1")
	mock.Room.Expect(
		client.Client.Room.FindUnique(
			db.Room.ID.Equals(room.ID),
		),
	).Returns(room)

	furniture := BuildTestFurniture("1", "1")
	mock.Furniture.Expect(
		client.Client.Furniture.FindUnique(
			db.Furniture.ID.Equals(furniture.ID),
		).With(
			db.Furniture.Room.Fetch(),
		),
	).Returns(furniture)

	updatedFurniture := furniture
	updatedFurniture.Archived = true
	mock.Furniture.Expect(
		client.Client.Furniture.FindUnique(
			db.Furniture.ID.Equals(furniture.ID),
		).With(
			db.Furniture.Room.Fetch(),
		).Update(
			db.Furniture.Archived.Set(true),
		),
	).Returns(updatedFurniture)

	b, err := json.Marshal(map[string]bool{"archive": true})
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/owner/properties/1/rooms/1/furnitures/1/archive/", bytes.NewReader(b))
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	var resp models.FurnitureResponse
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.JSONEq(t, resp.ID, furniture.ID)
	assert.True(t, resp.Archived)
}

func TestArchiveFurniture_MissingFields(t *testing.T) {
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

	room := BuildTestRoom("1", "1")
	mock.Room.Expect(
		client.Client.Room.FindUnique(
			db.Room.ID.Equals(room.ID),
		),
	).Returns(room)

	furniture := BuildTestFurniture("1", "1")
	mock.Furniture.Expect(
		client.Client.Furniture.FindUnique(
			db.Furniture.ID.Equals(furniture.ID),
		).With(
			db.Furniture.Room.Fetch(),
		),
	).Returns(furniture)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/owner/properties/1/rooms/1/furnitures/1/archive/", nil)
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

	room := BuildTestRoom("1", "1")
	mock.Room.Expect(
		client.Client.Room.FindUnique(
			db.Room.ID.Equals(room.ID),
		),
	).Returns(room)

	mock.Furniture.Expect(
		client.Client.Furniture.FindUnique(
			db.Furniture.ID.Equals("1"),
		).With(
			db.Furniture.Room.Fetch(),
		),
	).Errors(db.ErrNotFound)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/owner/properties/1/rooms/1/furnitures/1/archive/", nil)
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

	room := BuildTestRoom("1", "1")
	mock.Room.Expect(
		client.Client.Room.FindUnique(
			db.Room.ID.Equals(room.ID),
		),
	).Returns(room)

	furniture := BuildTestFurniture("1", "2")
	mock.Furniture.Expect(
		client.Client.Furniture.FindUnique(
			db.Furniture.ID.Equals(furniture.ID),
		).With(
			db.Furniture.Room.Fetch(),
		),
	).Returns(furniture)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/owner/properties/1/rooms/1/furnitures/1/archive/", nil)
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
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/owner/properties/1/rooms/1/furnitures/1/archive/", nil)
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

	room := BuildTestRoom("1", "1")
	mock.Room.Expect(
		client.Client.Room.FindUnique(
			db.Room.ID.Equals(room.ID),
		),
	).Returns(room)

	furnitures := []db.FurnitureModel{
		BuildTestFurniture("1", "1"),
		BuildTestFurniture("2", "1"),
	}
	for i := range furnitures {
		furnitures[i].Archived = true
	}
	mock.Furniture.Expect(
		client.Client.Furniture.FindMany(
			db.Furniture.RoomID.Equals(room.ID),
			db.Furniture.Archived.Equals(true),
		).With(
			db.Furniture.Room.Fetch(),
		),
	).ReturnsMany(furnitures)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/owner/properties/1/rooms/1/furnitures/archived/", nil)
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

	mock.Room.Expect(
		client.Client.Room.FindUnique(
			db.Room.ID.Equals("1"),
		),
	).Errors(db.ErrNotFound)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/owner/properties/1/rooms/1/furnitures/archived/", nil)
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

	room := BuildTestRoom("1", "2")
	mock.Room.Expect(
		client.Client.Room.FindUnique(
			db.Room.ID.Equals(room.ID),
		),
	).Returns(room)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/owner/properties/1/rooms/1/furnitures/archived/", nil)
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
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/owner/properties/1/rooms/1/furnitures/archived/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)
	var resp utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.PropertyNotFound, resp.Code)
}
