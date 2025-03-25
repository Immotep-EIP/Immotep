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

func TestCreateInventoryReport(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	invReport := db.InventoryReportModel{
		InnerInventoryReport: db.InnerInventoryReport{
			ID:         "1",
			Type:       db.ReportTypeStart,
			PropertyID: "1",
		},
	}

	mock.InventoryReport.Expect(
		client.Client.InventoryReport.CreateOne(
			db.InventoryReport.Type.Set(db.ReportTypeStart),
			db.InventoryReport.Property.Link(db.Property.ID.Equals("1")),
		),
	).Returns(invReport)

	newInvReport := database.CreateInvReport(db.ReportTypeStart, "1")
	assert.NotNil(t, newInvReport)
	assert.Equal(t, invReport.ID, newInvReport.ID)
}

func TestCreateInventoryReport_AlreadyExists(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	mock.InventoryReport.Expect(
		client.Client.InventoryReport.CreateOne(
			db.InventoryReport.Type.Set(db.ReportTypeStart),
			db.InventoryReport.Property.Link(db.Property.ID.Equals("1")),
		),
	).Errors(&protocol.UserFacingError{
		IsPanic:   false,
		ErrorCode: "P2002",
		Meta: protocol.Meta{
			Target: []any{"type", "propertyId"},
		},
		Message: "Unique constraint failed",
	})

	newInvReport := database.CreateInvReport(db.ReportTypeStart, "1")
	assert.Nil(t, newInvReport)
}

func TestCreateInventoryReport_NoConnection(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	mock.InventoryReport.Expect(
		client.Client.InventoryReport.CreateOne(
			db.InventoryReport.Type.Set(db.ReportTypeStart),
			db.InventoryReport.Property.Link(db.Property.ID.Equals("1")),
		),
	).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.CreateInvReport(db.ReportTypeStart, "1")
	})
}

func TestCreateRoomState(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	roomState := db.RoomStateModel{
		InnerRoomState: db.InnerRoomState{
			Cleanliness: db.CleanlinessClean,
			State:       db.StateGood,
			Note:        "Test note",
			RoomID:      "1",
		},
	}

	mock.RoomState.Expect(
		client.Client.RoomState.CreateOne(
			db.RoomState.Cleanliness.Set(db.CleanlinessClean),
			db.RoomState.State.Set(db.StateGood),
			db.RoomState.Note.Set("Test note"),
			db.RoomState.Report.Link(db.InventoryReport.ID.Equals("1")),
			db.RoomState.Room.Link(db.Room.ID.Equals("1")),
			db.RoomState.Pictures.Link(db.Image.ID.Equals("1")),
		),
	).Returns(roomState)

	newRoomState := database.CreateRoomState(roomState, []string{"1"}, "1")
	assert.NotNil(t, newRoomState)
	assert.Equal(t, roomState.Cleanliness, newRoomState.Cleanliness)
	assert.Equal(t, roomState.State, newRoomState.State)
	assert.Equal(t, roomState.Note, newRoomState.Note)
}

func TestCreateRoomState_AlreadyExists(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	roomState := db.RoomStateModel{
		InnerRoomState: db.InnerRoomState{
			Cleanliness: db.CleanlinessClean,
			State:       db.StateGood,
			Note:        "Test note",
			RoomID:      "1",
		},
	}

	mock.RoomState.Expect(
		client.Client.RoomState.CreateOne(
			db.RoomState.Cleanliness.Set(db.CleanlinessClean),
			db.RoomState.State.Set(db.StateGood),
			db.RoomState.Note.Set("Test note"),
			db.RoomState.Report.Link(db.InventoryReport.ID.Equals("1")),
			db.RoomState.Room.Link(db.Room.ID.Equals("1")),
			db.RoomState.Pictures.Link(db.Image.ID.Equals("1")),
		),
	).Errors(&protocol.UserFacingError{
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
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	roomState := db.RoomStateModel{
		InnerRoomState: db.InnerRoomState{
			Cleanliness: db.CleanlinessClean,
			State:       db.StateGood,
			Note:        "Test note",
			RoomID:      "1",
		},
	}

	mock.RoomState.Expect(
		client.Client.RoomState.CreateOne(
			db.RoomState.Cleanliness.Set(db.CleanlinessClean),
			db.RoomState.State.Set(db.StateGood),
			db.RoomState.Note.Set("Test note"),
			db.RoomState.Report.Link(db.InventoryReport.ID.Equals("1")),
			db.RoomState.Room.Link(db.Room.ID.Equals("1")),
			db.RoomState.Pictures.Link(db.Image.ID.Equals("1")),
		),
	).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.CreateRoomState(roomState, []string{"1"}, "1")
	})
}

func TestCreateFurnitureState(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	furnitureState := db.FurnitureStateModel{
		InnerFurnitureState: db.InnerFurnitureState{
			Cleanliness: db.CleanlinessClean,
			State:       db.StateGood,
			Note:        "Test note",
			FurnitureID: "1",
		},
	}

	mock.FurnitureState.Expect(
		client.Client.FurnitureState.CreateOne(
			db.FurnitureState.Cleanliness.Set(db.CleanlinessClean),
			db.FurnitureState.State.Set(db.StateGood),
			db.FurnitureState.Note.Set("Test note"),
			db.FurnitureState.Report.Link(db.InventoryReport.ID.Equals("1")),
			db.FurnitureState.Furniture.Link(db.Furniture.ID.Equals("1")),
			db.FurnitureState.Pictures.Link(db.Image.ID.Equals("1")),
		),
	).Returns(furnitureState)

	newFurnitureState := database.CreateFurnitureState(furnitureState, []string{"1"}, "1")
	assert.NotNil(t, newFurnitureState)
	assert.Equal(t, furnitureState.Cleanliness, newFurnitureState.Cleanliness)
	assert.Equal(t, furnitureState.State, newFurnitureState.State)
	assert.Equal(t, furnitureState.Note, newFurnitureState.Note)
}

func TestCreateFurnitureState_AlreadyExists(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	furnitureState := db.FurnitureStateModel{
		InnerFurnitureState: db.InnerFurnitureState{
			Cleanliness: db.CleanlinessClean,
			State:       db.StateGood,
			Note:        "Test note",
			FurnitureID: "1",
		},
	}

	mock.FurnitureState.Expect(
		client.Client.FurnitureState.CreateOne(
			db.FurnitureState.Cleanliness.Set(db.CleanlinessClean),
			db.FurnitureState.State.Set(db.StateGood),
			db.FurnitureState.Note.Set("Test note"),
			db.FurnitureState.Report.Link(db.InventoryReport.ID.Equals("1")),
			db.FurnitureState.Furniture.Link(db.Furniture.ID.Equals("1")),
			db.FurnitureState.Pictures.Link(db.Image.ID.Equals("1")),
		),
	).Errors(&protocol.UserFacingError{
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
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	furnitureState := db.FurnitureStateModel{
		InnerFurnitureState: db.InnerFurnitureState{
			Cleanliness: db.CleanlinessClean,
			State:       db.StateGood,
			Note:        "Test note",
			FurnitureID: "1",
		},
	}

	mock.FurnitureState.Expect(
		client.Client.FurnitureState.CreateOne(
			db.FurnitureState.Cleanliness.Set(db.CleanlinessClean),
			db.FurnitureState.State.Set(db.StateGood),
			db.FurnitureState.Note.Set("Test note"),
			db.FurnitureState.Report.Link(db.InventoryReport.ID.Equals("1")),
			db.FurnitureState.Furniture.Link(db.Furniture.ID.Equals("1")),
			db.FurnitureState.Pictures.Link(db.Image.ID.Equals("1")),
		),
	).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.CreateFurnitureState(furnitureState, []string{"1"}, "1")
	})
}

func TestGetInvReportByPropertyID(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	invReport := db.InventoryReportModel{
		InnerInventoryReport: db.InnerInventoryReport{
			ID:         "1",
			Type:       db.ReportTypeStart,
			PropertyID: "1",
		},
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
	).ReturnsMany([]db.InventoryReportModel{invReport})

	invReports := database.GetInvReportByPropertyID("1")
	assert.Len(t, invReports, 1)
	assert.Equal(t, invReport.ID, invReports[0].ID)
}

func TestGetInvReportByPropertyID_MultipleReports(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	invReport1 := db.InventoryReportModel{
		InnerInventoryReport: db.InnerInventoryReport{
			ID:         "1",
			Type:       db.ReportTypeStart,
			PropertyID: "1",
		},
	}

	invReport2 := db.InventoryReportModel{
		InnerInventoryReport: db.InnerInventoryReport{
			ID:         "2",
			Type:       db.ReportTypeEnd,
			PropertyID: "1",
		},
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
	).ReturnsMany([]db.InventoryReportModel{invReport1, invReport2})

	invReports := database.GetInvReportByPropertyID("1")
	assert.Len(t, invReports, 2)
	assert.Equal(t, invReport1.ID, invReports[0].ID)
	assert.Equal(t, invReport2.ID, invReports[1].ID)
}

func TestGetInvReportByPropertyID_NoReports(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

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
	).ReturnsMany([]db.InventoryReportModel{})

	invReports := database.GetInvReportByPropertyID("1")
	assert.Empty(t, invReports)
}

func TestGetInvReportByPropertyID_NoConnection(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

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
	).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.GetInvReportByPropertyID("1")
	})
}

func TestGetInvReportByID(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	invReport := db.InventoryReportModel{
		InnerInventoryReport: db.InnerInventoryReport{
			ID:         "1",
			Type:       db.ReportTypeStart,
			PropertyID: "1",
		},
	}

	mock.InventoryReport.Expect(
		client.Client.InventoryReport.FindUnique(
			db.InventoryReport.ID.Equals("1"),
		).With(
			db.InventoryReport.Property.Fetch(),
			db.InventoryReport.RoomStates.Fetch().With(db.RoomState.Room.Fetch()).With(db.RoomState.Pictures.Fetch()),
			db.InventoryReport.FurnitureStates.Fetch().With(db.FurnitureState.Furniture.Fetch()).With(db.FurnitureState.Pictures.Fetch()),
		),
	).Returns(invReport)

	foundInvReport := database.GetInvReportByID("1")
	assert.NotNil(t, foundInvReport)
	assert.Equal(t, invReport.ID, foundInvReport.ID)
}

func TestGetInvReportByID_NotFound(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	mock.InventoryReport.Expect(
		client.Client.InventoryReport.FindUnique(
			db.InventoryReport.ID.Equals("1"),
		).With(
			db.InventoryReport.Property.Fetch(),
			db.InventoryReport.RoomStates.Fetch().With(db.RoomState.Room.Fetch()).With(db.RoomState.Pictures.Fetch()),
			db.InventoryReport.FurnitureStates.Fetch().With(db.FurnitureState.Furniture.Fetch()).With(db.FurnitureState.Pictures.Fetch()),
		),
	).Errors(db.ErrNotFound)

	foundInvReport := database.GetInvReportByID("1")
	assert.Nil(t, foundInvReport)
}

func TestGetInvReportByID_NoConnection(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	mock.InventoryReport.Expect(
		client.Client.InventoryReport.FindUnique(
			db.InventoryReport.ID.Equals("1"),
		).With(
			db.InventoryReport.Property.Fetch(),
			db.InventoryReport.RoomStates.Fetch().With(db.RoomState.Room.Fetch()).With(db.RoomState.Pictures.Fetch()),
			db.InventoryReport.FurnitureStates.Fetch().With(db.FurnitureState.Furniture.Fetch()).With(db.FurnitureState.Pictures.Fetch()),
		),
	).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.GetInvReportByID("1")
	})
}

func TestGetLatestInvReport(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	invReport := db.InventoryReportModel{
		InnerInventoryReport: db.InnerInventoryReport{
			ID:         "1",
			Type:       db.ReportTypeStart,
			PropertyID: "1",
		},
	}

	mock.InventoryReport.Expect(
		client.Client.InventoryReport.FindFirst(
			db.InventoryReport.PropertyID.Equals("1"),
		).OrderBy(
			db.InventoryReport.Date.Order(db.SortOrderDesc),
		).With(
			db.InventoryReport.Property.Fetch(),
			db.InventoryReport.RoomStates.Fetch().With(db.RoomState.Room.Fetch()).With(db.RoomState.Pictures.Fetch()),
			db.InventoryReport.FurnitureStates.Fetch().With(db.FurnitureState.Furniture.Fetch()).With(db.FurnitureState.Pictures.Fetch()),
		),
	).Returns(invReport)

	latestInvReport := database.GetLatestInvReport("1")
	assert.NotNil(t, latestInvReport)
	assert.Equal(t, invReport.ID, latestInvReport.ID)
}

func TestGetLatestInvReport_NoReports(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	mock.InventoryReport.Expect(
		client.Client.InventoryReport.FindFirst(
			db.InventoryReport.PropertyID.Equals("1"),
		).OrderBy(
			db.InventoryReport.Date.Order(db.SortOrderDesc),
		).With(
			db.InventoryReport.Property.Fetch(),
			db.InventoryReport.RoomStates.Fetch().With(db.RoomState.Room.Fetch()).With(db.RoomState.Pictures.Fetch()),
			db.InventoryReport.FurnitureStates.Fetch().With(db.FurnitureState.Furniture.Fetch()).With(db.FurnitureState.Pictures.Fetch()),
		),
	).Errors(db.ErrNotFound)

	latestInvReport := database.GetLatestInvReport("1")
	assert.Nil(t, latestInvReport)
}

func TestGetLatestInvReport_NoConnection(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	mock.InventoryReport.Expect(
		client.Client.InventoryReport.FindFirst(
			db.InventoryReport.PropertyID.Equals("1"),
		).OrderBy(
			db.InventoryReport.Date.Order(db.SortOrderDesc),
		).With(
			db.InventoryReport.Property.Fetch(),
			db.InventoryReport.RoomStates.Fetch().With(db.RoomState.Room.Fetch()).With(db.RoomState.Pictures.Fetch()),
			db.InventoryReport.FurnitureStates.Fetch().With(db.FurnitureState.Furniture.Fetch()).With(db.FurnitureState.Pictures.Fetch()),
		),
	).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.GetLatestInvReport("1")
	})
}
