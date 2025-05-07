package database

import (
	"immotep/backend/prisma/db"
	"immotep/backend/services"
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

func SubmitInventoryReport(invrep db.InventoryReportModel) db.InventoryReportModel {
	pdb := services.DBclient
	ir, err := pdb.Client.InventoryReport.FindUnique(
		db.InventoryReport.ID.Equals(invrep.ID),
	).With(
		db.InventoryReport.Lease.Fetch(),
		db.InventoryReport.RoomStates.Fetch().With(db.RoomState.Room.Fetch()),
		db.InventoryReport.FurnitureStates.Fetch().With(db.FurnitureState.Furniture.Fetch()),
	).Update(
		db.InventoryReport.Submitted.Set(true),
	).Exec(pdb.Context)
	if err != nil {
		panic(err)
	}
	return *ir
}

func MockSubmitInventoryReport(c *services.PrismaDB) db.InventoryReportMockExpectParam {
	return c.Client.InventoryReport.FindUnique(
		db.InventoryReport.ID.Equals("1"),
	).With(
		db.InventoryReport.Lease.Fetch(),
		db.InventoryReport.RoomStates.Fetch().With(db.RoomState.Room.Fetch()),
		db.InventoryReport.FurnitureStates.Fetch().With(db.FurnitureState.Furniture.Fetch()),
	).Update(
		db.InventoryReport.Submitted.Set(true),
	)
}

func CreateRoomState(roomState db.RoomStateModel, invReportID string) *db.RoomStateModel {
	pdb := services.DBclient
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

func MockCreateRoomState(c *services.PrismaDB, roomState db.RoomStateModel) db.RoomStateMockExpectParam {
	return c.Client.RoomState.CreateOne(
		db.RoomState.Cleanliness.Set(roomState.Cleanliness),
		db.RoomState.State.Set(roomState.State),
		db.RoomState.Note.Set(roomState.Note),
		db.RoomState.Report.Link(db.InventoryReport.ID.Equals("1")),
		db.RoomState.Room.Link(db.Room.ID.Equals("1")),
	)
}

func AddPicturesToRoomState(roomState db.RoomStateModel, picturePaths []string) db.RoomStateModel {
	pdb := services.DBclient
	rs, err := pdb.Client.RoomState.FindUnique(
		db.RoomState.ID.Equals(roomState.ID),
	).Update(
		db.RoomState.Pictures.Push(picturePaths),
	).Exec(pdb.Context)
	if err != nil {
		panic(err)
	}
	return *rs
}

func MockAddPicturesToRoomState(c *services.PrismaDB) db.RoomStateMockExpectParam {
	return c.Client.RoomState.FindUnique(
		db.RoomState.ID.Equals("1"),
	).Update(
		db.RoomState.Pictures.Push([]string{"/path/to/picture.jpg"}),
	)
}

func CreateFurnitureState(furnitureState db.FurnitureStateModel, invReportID string) *db.FurnitureStateModel {
	pdb := services.DBclient
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

func MockCreateFurnitureState(c *services.PrismaDB, furnitureState db.FurnitureStateModel) db.FurnitureStateMockExpectParam {
	return c.Client.FurnitureState.CreateOne(
		db.FurnitureState.Cleanliness.Set(furnitureState.Cleanliness),
		db.FurnitureState.State.Set(furnitureState.State),
		db.FurnitureState.Note.Set(furnitureState.Note),
		db.FurnitureState.Report.Link(db.InventoryReport.ID.Equals("1")),
		db.FurnitureState.Furniture.Link(db.Furniture.ID.Equals("1")),
	)
}

func AddPicturesToFurnitureState(furnitureState db.FurnitureStateModel, picturePaths []string) db.FurnitureStateModel {
	pdb := services.DBclient
	fs, err := pdb.Client.FurnitureState.FindUnique(
		db.FurnitureState.ID.Equals(furnitureState.ID),
	).Update(
		db.FurnitureState.Pictures.Push(picturePaths),
	).Exec(pdb.Context)
	if err != nil {
		panic(err)
	}
	return *fs
}

func MockAddPicturesToFurnitureState(c *services.PrismaDB) db.FurnitureStateMockExpectParam {
	return c.Client.FurnitureState.FindUnique(
		db.FurnitureState.ID.Equals("1"),
	).Update(
		db.FurnitureState.Pictures.Push([]string{"/path/to/picture.jpg"}),
	)
}

func GetInvReportsByPropertyID(propertyID string) []db.InventoryReportModel {
	pdb := services.DBclient
	invReports, err := pdb.Client.InventoryReport.FindMany(
		db.InventoryReport.Lease.Where(db.Lease.PropertyID.Equals(propertyID)),
		db.InventoryReport.Submitted.Equals(true),
	).OrderBy(
		db.InventoryReport.Date.Order(db.SortOrderDesc),
	).With(
		db.InventoryReport.Lease.Fetch(),
		db.InventoryReport.RoomStates.Fetch().With(db.RoomState.Room.Fetch()),
		db.InventoryReport.FurnitureStates.Fetch().With(db.FurnitureState.Furniture.Fetch()),
	).Exec(pdb.Context)
	if err != nil {
		panic(err)
	}
	return invReports
}

func MockGetInvReportByPropertyID(c *services.PrismaDB) db.InventoryReportMockExpectParam {
	return c.Client.InventoryReport.FindMany(
		db.InventoryReport.Lease.Where(db.Lease.PropertyID.Equals("1")),
		db.InventoryReport.Submitted.Equals(true),
	).OrderBy(
		db.InventoryReport.Date.Order(db.SortOrderDesc),
	).With(
		db.InventoryReport.Lease.Fetch(),
		db.InventoryReport.RoomStates.Fetch().With(db.RoomState.Room.Fetch()),
		db.InventoryReport.FurnitureStates.Fetch().With(db.FurnitureState.Furniture.Fetch()),
	)
}

func GetInvReportsByLeaseID(leaseID string) []db.InventoryReportModel {
	pdb := services.DBclient
	invReports, err := pdb.Client.InventoryReport.FindMany(
		db.InventoryReport.LeaseID.Equals(leaseID),
		db.InventoryReport.Submitted.Equals(true),
	).OrderBy(
		db.InventoryReport.Date.Order(db.SortOrderDesc),
	).With(
		db.InventoryReport.Lease.Fetch(),
		db.InventoryReport.RoomStates.Fetch().With(db.RoomState.Room.Fetch()),
		db.InventoryReport.FurnitureStates.Fetch().With(db.FurnitureState.Furniture.Fetch()),
	).Exec(pdb.Context)
	if err != nil {
		panic(err)
	}
	return invReports
}

func MockGetInvReportByLeaseID(c *services.PrismaDB) db.InventoryReportMockExpectParam {
	return c.Client.InventoryReport.FindMany(
		db.InventoryReport.LeaseID.Equals("1"),
		db.InventoryReport.Submitted.Equals(true),
	).OrderBy(
		db.InventoryReport.Date.Order(db.SortOrderDesc),
	).With(
		db.InventoryReport.Lease.Fetch(),
		db.InventoryReport.RoomStates.Fetch().With(db.RoomState.Room.Fetch()),
		db.InventoryReport.FurnitureStates.Fetch().With(db.FurnitureState.Furniture.Fetch()),
	)
}

func GetInvReportByID(id string) *db.InventoryReportModel {
	pdb := services.DBclient
	invReport, err := pdb.Client.InventoryReport.FindUnique(
		db.InventoryReport.ID.Equals(id),
	).With(
		db.InventoryReport.Lease.Fetch(),
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

func MockGetInvReportByID(c *services.PrismaDB) db.InventoryReportMockExpectParam {
	return c.Client.InventoryReport.FindUnique(
		db.InventoryReport.ID.Equals("1"),
	).With(
		db.InventoryReport.Lease.Fetch(),
		db.InventoryReport.RoomStates.Fetch().With(db.RoomState.Room.Fetch()),
		db.InventoryReport.FurnitureStates.Fetch().With(db.FurnitureState.Furniture.Fetch()),
	)
}

func GetLatestInvReportByLease(leaseID string) *db.InventoryReportModel {
	pdb := services.DBclient
	invReport, err := pdb.Client.InventoryReport.FindFirst(
		db.InventoryReport.LeaseID.Equals(leaseID),
		db.InventoryReport.Submitted.Equals(true),
	).OrderBy(
		db.InventoryReport.Date.Order(db.SortOrderDesc),
	).With(
		db.InventoryReport.Lease.Fetch(),
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

func MockGetLatestInvReportByLease(c *services.PrismaDB) db.InventoryReportMockExpectParam {
	return c.Client.InventoryReport.FindFirst(
		db.InventoryReport.LeaseID.Equals("1"),
		db.InventoryReport.Submitted.Equals(true),
	).OrderBy(
		db.InventoryReport.Date.Order(db.SortOrderDesc),
	).With(
		db.InventoryReport.Lease.Fetch(),
		db.InventoryReport.RoomStates.Fetch().With(db.RoomState.Room.Fetch()),
		db.InventoryReport.FurnitureStates.Fetch().With(db.FurnitureState.Furniture.Fetch()),
	)
}

func GetLatestInvReportByProperty(propertyID string) *db.InventoryReportModel {
	pdb := services.DBclient
	invReport, err := pdb.Client.InventoryReport.FindFirst(
		db.InventoryReport.Lease.Where(db.Lease.PropertyID.Equals(propertyID)),
		db.InventoryReport.Submitted.Equals(true),
	).OrderBy(
		db.InventoryReport.Date.Order(db.SortOrderDesc),
	).With(
		db.InventoryReport.Lease.Fetch(),
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

func MockGetLatestInvReportByProperty(c *services.PrismaDB) db.InventoryReportMockExpectParam {
	return c.Client.InventoryReport.FindFirst(
		db.InventoryReport.Lease.Where(db.Lease.PropertyID.Equals("1")),
		db.InventoryReport.Submitted.Equals(true),
	).OrderBy(
		db.InventoryReport.Date.Order(db.SortOrderDesc),
	).With(
		db.InventoryReport.Lease.Fetch(),
		db.InventoryReport.RoomStates.Fetch().With(db.RoomState.Room.Fetch()),
		db.InventoryReport.FurnitureStates.Fetch().With(db.FurnitureState.Furniture.Fetch()),
	)
}
