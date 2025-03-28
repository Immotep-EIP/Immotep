package database

import (
	"immotep/backend/prisma/db"
	"immotep/backend/services"
)

func CreateInvReport(repType db.ReportType, propertyId string) *db.InventoryReportModel {
	pdb := services.DBclient
	newInvReport, err := pdb.Client.InventoryReport.CreateOne(
		db.InventoryReport.Type.Set(repType),
		db.InventoryReport.Property.Link(db.Property.ID.Equals(propertyId)),
	).Exec(pdb.Context)
	if err != nil {
		if _, is := db.IsErrUniqueConstraint(err); is {
			return nil
		}
		panic(err)
	}
	return newInvReport
}

func CreateRoomState(roomState db.RoomStateModel, picturesId []string, invReportID string) *db.RoomStateModel {
	params := make([]db.RoomStateSetParam, 0, len(picturesId))
	for _, id := range picturesId {
		params = append(params, db.RoomState.Pictures.Link(db.Image.ID.Equals(id)))
	}

	pdb := services.DBclient
	newRoomState, err := pdb.Client.RoomState.CreateOne(
		db.RoomState.Cleanliness.Set(roomState.Cleanliness),
		db.RoomState.State.Set(roomState.State),
		db.RoomState.Note.Set(roomState.Note),
		db.RoomState.Report.Link(db.InventoryReport.ID.Equals(invReportID)),
		db.RoomState.Room.Link(db.Room.ID.Equals(roomState.RoomID)),
		params...,
	).Exec(pdb.Context)
	if err != nil {
		if _, is := db.IsErrUniqueConstraint(err); is {
			return nil
		}
		panic(err)
	}
	return newRoomState
}

func CreateFurnitureState(furnitureState db.FurnitureStateModel, picturesId []string, invReportID string) *db.FurnitureStateModel {
	params := make([]db.FurnitureStateSetParam, 0, len(picturesId))
	for _, id := range picturesId {
		params = append(params, db.FurnitureState.Pictures.Link(db.Image.ID.Equals(id)))
	}

	pdb := services.DBclient
	newFurnitureState, err := pdb.Client.FurnitureState.CreateOne(
		db.FurnitureState.Cleanliness.Set(furnitureState.Cleanliness),
		db.FurnitureState.State.Set(furnitureState.State),
		db.FurnitureState.Note.Set(furnitureState.Note),
		db.FurnitureState.Report.Link(db.InventoryReport.ID.Equals(invReportID)),
		db.FurnitureState.Furniture.Link(db.Furniture.ID.Equals(furnitureState.FurnitureID)),
		params...,
	).Exec(pdb.Context)
	if err != nil {
		if _, is := db.IsErrUniqueConstraint(err); is {
			return nil
		}
		panic(err)
	}
	return newFurnitureState
}

func GetInvReportByPropertyID(propertyID string) []db.InventoryReportModel {
	pdb := services.DBclient
	invReports, err := pdb.Client.InventoryReport.FindMany(
		db.InventoryReport.PropertyID.Equals(propertyID),
	).OrderBy(
		db.InventoryReport.Date.Order(db.SortOrderDesc),
	).With(
		db.InventoryReport.Property.Fetch(),
		db.InventoryReport.RoomStates.Fetch().With(db.RoomState.Room.Fetch()).With(db.RoomState.Pictures.Fetch()),
		db.InventoryReport.FurnitureStates.Fetch().With(db.FurnitureState.Furniture.Fetch()).With(db.FurnitureState.Pictures.Fetch()),
	).Exec(pdb.Context)
	if err != nil {
		panic(err)
	}
	return invReports
}

func GetInvReportByID(id string) *db.InventoryReportModel {
	pdb := services.DBclient
	invReport, err := pdb.Client.InventoryReport.FindUnique(
		db.InventoryReport.ID.Equals(id),
	).With(
		db.InventoryReport.Property.Fetch(),
		db.InventoryReport.RoomStates.Fetch().With(db.RoomState.Room.Fetch()).With(db.RoomState.Pictures.Fetch()),
		db.InventoryReport.FurnitureStates.Fetch().With(db.FurnitureState.Furniture.Fetch()).With(db.FurnitureState.Pictures.Fetch()),
	).Exec(pdb.Context)
	if err != nil {
		if db.IsErrNotFound(err) {
			return nil
		}
		panic(err)
	}
	return invReport
}

func GetLatestInvReport(propertyID string) *db.InventoryReportModel {
	pdb := services.DBclient
	invReport, err := pdb.Client.InventoryReport.FindFirst(
		db.InventoryReport.PropertyID.Equals(propertyID),
	).OrderBy(
		db.InventoryReport.Date.Order(db.SortOrderDesc),
	).With(
		db.InventoryReport.Property.Fetch(),
		db.InventoryReport.RoomStates.Fetch().With(db.RoomState.Room.Fetch()).With(db.RoomState.Pictures.Fetch()),
		db.InventoryReport.FurnitureStates.Fetch().With(db.FurnitureState.Furniture.Fetch()).With(db.FurnitureState.Pictures.Fetch()),
	).Exec(pdb.Context)
	if err != nil {
		if db.IsErrNotFound(err) {
			return nil
		}
		panic(err)
	}
	return invReport
}
