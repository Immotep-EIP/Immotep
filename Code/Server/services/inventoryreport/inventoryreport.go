package inventoryreportservice

import (
	"immotep/backend/database"
	"immotep/backend/prisma/db"
)

func Create(repType db.ReportType, propertyId string) *db.InventoryReportModel {
	pdb := database.DBclient
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

func CreateRoomState(roomState db.RoomStateModel, invReportID string) *db.RoomStateModel {
	pdb := database.DBclient
	newRoomState, err := pdb.Client.RoomState.CreateOne(
		db.RoomState.Cleanliness.Set(roomState.Cleanliness),
		db.RoomState.State.Set(roomState.State),
		db.RoomState.Note.Set(roomState.Note),
		db.RoomState.Report.Link(db.InventoryReport.ID.Equals(invReportID)),
		db.RoomState.Room.Link(db.Room.ID.Equals(roomState.RoomID)),
	).Exec(pdb.Context)
	if err != nil {
		if _, is := db.IsErrUniqueConstraint(err); is {
			return nil
		}
		panic(err)
	}
	return newRoomState
}

func CreateFurnitureState(furnitureState db.FurnitureStateModel, invReportID string) *db.FurnitureStateModel {
	pdb := database.DBclient
	newFurnitureState, err := pdb.Client.FurnitureState.CreateOne(
		db.FurnitureState.Cleanliness.Set(furnitureState.Cleanliness),
		db.FurnitureState.State.Set(furnitureState.State),
		db.FurnitureState.Note.Set(furnitureState.Note),
		db.FurnitureState.Report.Link(db.InventoryReport.ID.Equals(invReportID)),
		db.FurnitureState.Furniture.Link(db.Furniture.ID.Equals(furnitureState.FurnitureID)),
	).Exec(pdb.Context)
	if err != nil {
		if _, is := db.IsErrUniqueConstraint(err); is {
			return nil
		}
		panic(err)
	}
	return newFurnitureState
}

func GetByPropertyID(propertyID string) []db.InventoryReportModel {
	pdb := database.DBclient
	invReports, err := pdb.Client.InventoryReport.FindMany(
		db.InventoryReport.PropertyID.Equals(propertyID),
	).With(
		db.InventoryReport.Property.Fetch(),
		db.InventoryReport.RoomStates.Fetch().With(db.RoomState.Room.Fetch()),
		db.InventoryReport.FurnitureStates.Fetch().With(db.FurnitureState.Furniture.Fetch()),
	).Exec(pdb.Context)
	if err != nil {
		panic(err)
	}
	return invReports
}

func GetByID(id string) *db.InventoryReportModel {
	pdb := database.DBclient
	invReport, err := pdb.Client.InventoryReport.FindUnique(
		db.InventoryReport.ID.Equals(id),
	).With(
		db.InventoryReport.Property.Fetch(),
		db.InventoryReport.RoomStates.Fetch().With(db.RoomState.Room.Fetch()),
		db.InventoryReport.FurnitureStates.Fetch().With(db.FurnitureState.Furniture.Fetch()),
	).Exec(pdb.Context)
	if err != nil {
		if db.IsErrNotFound(err) {
			return nil
		}
		panic(err)
	}
	return invReport
}
