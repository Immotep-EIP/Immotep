package database

import (
	"keyz/backend/prisma/db"
	"keyz/backend/services"
)

func CreateInvReport(repType db.ReportType, leaseId string) *db.InventoryReportModel {
	pdb := services.DBclient
	newInvReport, err := pdb.Client.InventoryReport.CreateOne(
		db.InventoryReport.Type.Set(repType),
		db.InventoryReport.Lease.Link(db.Lease.ID.Equals(leaseId)),
	).With(
		db.InventoryReport.Lease.Fetch(),
	).Exec(pdb.Context)
	if err != nil {
		if _, is := db.IsErrUniqueConstraint(err); is {
			return nil
		}
		panic(err)
	}
	return newInvReport
}

func MockCreateInventoryReport(c *services.PrismaDB, invRep db.InventoryReportModel) db.InventoryReportMockExpectParam {
	return c.Client.InventoryReport.CreateOne(
		db.InventoryReport.Type.Set(invRep.Type),
		db.InventoryReport.Lease.Link(db.Lease.ID.Equals("1")),
	).With(
		db.InventoryReport.Lease.Fetch(),
	)
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

func MockCreateRoomState(c *services.PrismaDB, roomState db.RoomStateModel) db.RoomStateMockExpectParam {
	return c.Client.RoomState.CreateOne(
		db.RoomState.Cleanliness.Set(roomState.Cleanliness),
		db.RoomState.State.Set(roomState.State),
		db.RoomState.Note.Set(roomState.Note),
		db.RoomState.Report.Link(db.InventoryReport.ID.Equals("1")),
		db.RoomState.Room.Link(db.Room.ID.Equals("1")),
		db.RoomState.Pictures.Link(db.Image.ID.Equals("1")),
	)
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

func MockCreateFurnitureState(c *services.PrismaDB, furnitureState db.FurnitureStateModel) db.FurnitureStateMockExpectParam {
	return c.Client.FurnitureState.CreateOne(
		db.FurnitureState.Cleanliness.Set(furnitureState.Cleanliness),
		db.FurnitureState.State.Set(furnitureState.State),
		db.FurnitureState.Note.Set(furnitureState.Note),
		db.FurnitureState.Report.Link(db.InventoryReport.ID.Equals("1")),
		db.FurnitureState.Furniture.Link(db.Furniture.ID.Equals("1")),
		db.FurnitureState.Pictures.Link(db.Image.ID.Equals("1")),
	)
}

func GetInvReportsByPropertyID(propertyID string) []db.InventoryReportModel {
	pdb := services.DBclient
	invReports, err := pdb.Client.InventoryReport.FindMany(
		db.InventoryReport.Lease.Where(db.Lease.PropertyID.Equals(propertyID)),
	).OrderBy(
		db.InventoryReport.Date.Order(db.SortOrderDesc),
	).With(
		db.InventoryReport.Lease.Fetch(),
		db.InventoryReport.RoomStates.Fetch().With(db.RoomState.Room.Fetch()).With(db.RoomState.Pictures.Fetch()),
		db.InventoryReport.FurnitureStates.Fetch().With(db.FurnitureState.Furniture.Fetch()).With(db.FurnitureState.Pictures.Fetch()),
	).Exec(pdb.Context)
	if err != nil {
		panic(err)
	}
	return invReports
}

func MockGetInvReportByPropertyID(c *services.PrismaDB) db.InventoryReportMockExpectParam {
	return c.Client.InventoryReport.FindMany(
		db.InventoryReport.Lease.Where(db.Lease.PropertyID.Equals("1")),
	).OrderBy(
		db.InventoryReport.Date.Order(db.SortOrderDesc),
	).With(
		db.InventoryReport.Lease.Fetch(),
		db.InventoryReport.RoomStates.Fetch().With(db.RoomState.Room.Fetch()).With(db.RoomState.Pictures.Fetch()),
		db.InventoryReport.FurnitureStates.Fetch().With(db.FurnitureState.Furniture.Fetch()).With(db.FurnitureState.Pictures.Fetch()),
	)
}

func GetInvReportsByLeaseID(leaseID string) []db.InventoryReportModel {
	pdb := services.DBclient
	invReports, err := pdb.Client.InventoryReport.FindMany(
		db.InventoryReport.LeaseID.Equals(leaseID),
	).OrderBy(
		db.InventoryReport.Date.Order(db.SortOrderDesc),
	).With(
		db.InventoryReport.Lease.Fetch(),
		db.InventoryReport.RoomStates.Fetch().With(db.RoomState.Room.Fetch()).With(db.RoomState.Pictures.Fetch()),
		db.InventoryReport.FurnitureStates.Fetch().With(db.FurnitureState.Furniture.Fetch()).With(db.FurnitureState.Pictures.Fetch()),
	).Exec(pdb.Context)
	if err != nil {
		panic(err)
	}
	return invReports
}

func MockGetInvReportsByLeaseID(c *services.PrismaDB) db.InventoryReportMockExpectParam {
	return c.Client.InventoryReport.FindMany(
		db.InventoryReport.LeaseID.Equals("1"),
	).OrderBy(
		db.InventoryReport.Date.Order(db.SortOrderDesc),
	).With(
		db.InventoryReport.Lease.Fetch(),
		db.InventoryReport.RoomStates.Fetch().With(db.RoomState.Room.Fetch()).With(db.RoomState.Pictures.Fetch()),
		db.InventoryReport.FurnitureStates.Fetch().With(db.FurnitureState.Furniture.Fetch()).With(db.FurnitureState.Pictures.Fetch()),
	)
}

func GetInvReportByID(id string) *db.InventoryReportModel {
	pdb := services.DBclient
	invReport, err := pdb.Client.InventoryReport.FindUnique(
		db.InventoryReport.ID.Equals(id),
	).With(
		db.InventoryReport.Lease.Fetch(),
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

func MockGetInvReportByID(c *services.PrismaDB) db.InventoryReportMockExpectParam {
	return c.Client.InventoryReport.FindUnique(
		db.InventoryReport.ID.Equals("1"),
	).With(
		db.InventoryReport.Lease.Fetch(),
		db.InventoryReport.RoomStates.Fetch().With(db.RoomState.Room.Fetch()).With(db.RoomState.Pictures.Fetch()),
		db.InventoryReport.FurnitureStates.Fetch().With(db.FurnitureState.Furniture.Fetch()).With(db.FurnitureState.Pictures.Fetch()),
	)
}

func GetLatestInvReportByLease(leaseID string) *db.InventoryReportModel {
	pdb := services.DBclient
	invReport, err := pdb.Client.InventoryReport.FindFirst(
		db.InventoryReport.LeaseID.Equals(leaseID),
	).OrderBy(
		db.InventoryReport.Date.Order(db.SortOrderDesc),
	).With(
		db.InventoryReport.Lease.Fetch(),
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

func MockGetLatestInvReportByLease(c *services.PrismaDB) db.InventoryReportMockExpectParam {
	return c.Client.InventoryReport.FindFirst(
		db.InventoryReport.LeaseID.Equals("1"),
	).OrderBy(
		db.InventoryReport.Date.Order(db.SortOrderDesc),
	).With(
		db.InventoryReport.Lease.Fetch(),
		db.InventoryReport.RoomStates.Fetch().With(db.RoomState.Room.Fetch()).With(db.RoomState.Pictures.Fetch()),
		db.InventoryReport.FurnitureStates.Fetch().With(db.FurnitureState.Furniture.Fetch()).With(db.FurnitureState.Pictures.Fetch()),
	)
}

func GetLatestInvReportByProperty(propertyID string) *db.InventoryReportModel {
	pdb := services.DBclient
	invReport, err := pdb.Client.InventoryReport.FindFirst(
		db.InventoryReport.Lease.Where(db.Lease.PropertyID.Equals(propertyID)),
	).OrderBy(
		db.InventoryReport.Date.Order(db.SortOrderDesc),
	).With(
		db.InventoryReport.Lease.Fetch(),
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

func MockGetLatestInvReport(c *services.PrismaDB) db.InventoryReportMockExpectParam {
	return c.Client.InventoryReport.FindFirst(
		db.InventoryReport.Lease.Where(db.Lease.PropertyID.Equals("1")),
	).OrderBy(
		db.InventoryReport.Date.Order(db.SortOrderDesc),
	).With(
		db.InventoryReport.Lease.Fetch(),
		db.InventoryReport.RoomStates.Fetch().With(db.RoomState.Room.Fetch()).With(db.RoomState.Pictures.Fetch()),
		db.InventoryReport.FurnitureStates.Fetch().With(db.FurnitureState.Furniture.Fetch()).With(db.FurnitureState.Pictures.Fetch()),
	)
}
