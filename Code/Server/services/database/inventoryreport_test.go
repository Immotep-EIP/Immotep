package database_test

import (
	"errors"
	"testing"

	"github.com/steebchen/prisma-client-go/engine/protocol"
	"github.com/stretchr/testify/assert"
	"immotep/backend/prisma/db"
	"immotep/backend/services"
	"immotep/backend/services/database"
)

func BuildTestInventoryReport(id string) db.InventoryReportModel {
	return db.InventoryReportModel{
		InnerInventoryReport: db.InnerInventoryReport{
			ID:         id,
			Type:       db.ReportTypeStart,
			PropertyID: "1",
		},
	}
}

func BuildTestRoomState(id string) db.RoomStateModel {
	return db.RoomStateModel{
		InnerRoomState: db.InnerRoomState{
			ID:          id,
			Cleanliness: db.CleanlinessClean,
			State:       db.StateGood,
			Note:        "Test note",
			RoomID:      "1",
		},
	}
}

func BuildTestFurnitureState(id string) db.FurnitureStateModel {
	return db.FurnitureStateModel{
		InnerFurnitureState: db.InnerFurnitureState{
			ID:          id,
			Cleanliness: db.CleanlinessClean,
			State:       db.StateGood,
			Note:        "Test note",
			FurnitureID: "1",
		},
	}
}

// #############################################################################

func TestCreateInventoryReport(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	invReport := BuildTestInventoryReport("1")
	m.InventoryReport.Expect(database.MockCreateInventoryReport(c, invReport)).Returns(invReport)

	newInvReport := database.CreateInvReport(invReport.Type, "1")
	assert.NotNil(t, newInvReport)
	assert.Equal(t, invReport.ID, newInvReport.ID)
}

func TestCreateInventoryReport_AlreadyExists(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	invReport := BuildTestInventoryReport("1")
	m.InventoryReport.Expect(database.MockCreateInventoryReport(c, invReport)).Errors(&protocol.UserFacingError{
		IsPanic:   false,
		ErrorCode: "P2002",
		Meta: protocol.Meta{
			Target: []any{"type", "propertyId"},
		},
		Message: "Unique constraint failed",
	})

	newInvReport := database.CreateInvReport(invReport.Type, "1")
	assert.Nil(t, newInvReport)
}

func TestCreateInventoryReport_NoConnection(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	invReport := BuildTestInventoryReport("1")
	m.InventoryReport.Expect(database.MockCreateInventoryReport(c, invReport)).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.CreateInvReport(invReport.Type, "1")
	})
}

// #############################################################################

func TestCreateRoomState(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	roomState := BuildTestRoomState("1")
	m.RoomState.Expect(database.MockCreateRoomState(c, roomState)).Returns(roomState)

	newRoomState := database.CreateRoomState(roomState, []string{"1"}, "1")
	assert.NotNil(t, newRoomState)
	assert.Equal(t, roomState.Cleanliness, newRoomState.Cleanliness)
	assert.Equal(t, roomState.State, newRoomState.State)
	assert.Equal(t, roomState.Note, newRoomState.Note)
}

func TestCreateRoomState_AlreadyExists(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	roomState := BuildTestRoomState("1")
	m.RoomState.Expect(database.MockCreateRoomState(c, roomState)).Errors(&protocol.UserFacingError{
		IsPanic:   false,
		ErrorCode: "P2002",
		Meta: protocol.Meta{
			Target: []any{"roomId", "reportId"},
		},
		Message: "Unique constraint failed",
	})

	newRoomState := database.CreateRoomState(roomState, []string{"1"}, "1")
	assert.Nil(t, newRoomState)
}

func TestCreateRoomState_NoConnection(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	roomState := BuildTestRoomState("1")
	m.RoomState.Expect(database.MockCreateRoomState(c, roomState)).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.CreateRoomState(roomState, []string{"1"}, "1")
	})
}

// #############################################################################

func TestCreateFurnitureState(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	furnitureState := BuildTestFurnitureState("1")
	m.FurnitureState.Expect(database.MockCreateFurnitureState(c, furnitureState)).Returns(furnitureState)

	newFurnitureState := database.CreateFurnitureState(furnitureState, []string{"1"}, "1")
	assert.NotNil(t, newFurnitureState)
	assert.Equal(t, furnitureState.Cleanliness, newFurnitureState.Cleanliness)
	assert.Equal(t, furnitureState.State, newFurnitureState.State)
	assert.Equal(t, furnitureState.Note, newFurnitureState.Note)
}

func TestCreateFurnitureState_AlreadyExists(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	furnitureState := BuildTestFurnitureState("1")
	m.FurnitureState.Expect(database.MockCreateFurnitureState(c, furnitureState)).Errors(&protocol.UserFacingError{
		IsPanic:   false,
		ErrorCode: "P2002",
		Meta: protocol.Meta{
			Target: []any{"furnitureId", "reportId"},
		},
		Message: "Unique constraint failed",
	})

	newFurnitureState := database.CreateFurnitureState(furnitureState, []string{"1"}, "1")
	assert.Nil(t, newFurnitureState)
}

func TestCreateFurnitureState_NoConnection(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	furnitureState := BuildTestFurnitureState("1")
	m.FurnitureState.Expect(database.MockCreateFurnitureState(c, furnitureState)).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.CreateFurnitureState(furnitureState, []string{"1"}, "1")
	})
}

// #############################################################################

func TestGetInvReportByPropertyID(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	invReport := BuildTestInventoryReport("1")
	m.InventoryReport.Expect(database.MockGetInvReportByPropertyID(c)).ReturnsMany([]db.InventoryReportModel{invReport})

	invReports := database.GetInvReportsByPropertyID("1")
	assert.Len(t, invReports, 1)
	assert.Equal(t, invReport.ID, invReports[0].ID)
}

func TestGetInvReportByPropertyID_MultipleReports(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	invReport1 := BuildTestInventoryReport("1")
	invReport2 := BuildTestInventoryReport("2")
	m.InventoryReport.Expect(database.MockGetInvReportByPropertyID(c)).ReturnsMany([]db.InventoryReportModel{invReport1, invReport2})

	invReports := database.GetInvReportsByPropertyID("1")
	assert.Len(t, invReports, 2)
	assert.Equal(t, invReport1.ID, invReports[0].ID)
	assert.Equal(t, invReport2.ID, invReports[1].ID)
}

func TestGetInvReportByPropertyID_NoReports(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.InventoryReport.Expect(database.MockGetInvReportByPropertyID(c)).ReturnsMany([]db.InventoryReportModel{})

	invReports := database.GetInvReportsByPropertyID("1")
	assert.Empty(t, invReports)
}

func TestGetInvReportByPropertyID_NoConnection(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.InventoryReport.Expect(database.MockGetInvReportByPropertyID(c)).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.GetInvReportsByPropertyID("1")
	})
}

// #############################################################################

func TestGetInvReportByID(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	invReport := BuildTestInventoryReport("1")
	m.InventoryReport.Expect(database.MockGetInvReportByID(c)).Returns(invReport)

	foundInvReport := database.GetInvReportByID("1")
	assert.NotNil(t, foundInvReport)
	assert.Equal(t, invReport.ID, foundInvReport.ID)
}

func TestGetInvReportByID_NotFound(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.InventoryReport.Expect(database.MockGetInvReportByID(c)).Errors(db.ErrNotFound)

	foundInvReport := database.GetInvReportByID("1")
	assert.Nil(t, foundInvReport)
}

func TestGetInvReportByID_NoConnection(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.InventoryReport.Expect(database.MockGetInvReportByID(c)).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.GetInvReportByID("1")
	})
}

// #############################################################################

func TestGetLatestInvReport(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	invReport := BuildTestInventoryReport("1")
	m.InventoryReport.Expect(database.MockGetLatestInvReport(c)).Returns(invReport)

	latestInvReport := database.GetLatestInvReport("1")
	assert.NotNil(t, latestInvReport)
	assert.Equal(t, invReport.ID, latestInvReport.ID)
}

func TestGetLatestInvReport_NoReports(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.InventoryReport.Expect(database.MockGetLatestInvReport(c)).Errors(db.ErrNotFound)

	latestInvReport := database.GetLatestInvReport("1")
	assert.Nil(t, latestInvReport)
}

func TestGetLatestInvReport_NoConnection(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.InventoryReport.Expect(database.MockGetLatestInvReport(c)).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.GetLatestInvReport("1")
	})
}
