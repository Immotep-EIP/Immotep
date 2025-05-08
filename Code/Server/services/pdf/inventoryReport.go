package pdf

import (
	"log"
	"strconv"

	"immotep/backend/prisma/db"
	"immotep/backend/services/filesystem"
)

func NewInventoryReportPDF(invReport db.InventoryReportModel, lease db.LeaseModel) (*File, error) {
	report := NewPDF()

	report.AddCenteredTitle("Inventory Report", H1)
	report.AddText("ID: " + invReport.ID)
	report.AddText("Date: " + invReport.Date.Format("2006-01-02 15:04:05"))
	report.AddText("Property: " + lease.Property().Name)
	report.AddText("Type: " + string(invReport.Type))

	report.Ln(5)
	report.AddTitle("Lease", H2)
	report.Add2Texts("Owner: "+lease.Property().Owner().Name(), "Email: "+lease.Property().Owner().Email)
	report.Add2Texts("Tenant: "+lease.Tenant().Name(), "Email: "+lease.Tenant().Email)
	leaseEndDate, ok := lease.EndDate()
	if ok {
		report.Add2Texts("Start date: "+lease.StartDate.Format("2006-01-02"), "End date: "+leaseEndDate.Format("2006-01-02"))
	} else {
		report.Add2Texts("Start date: "+lease.StartDate.Format("2006-01-02"), "End date: None")
	}
	report.AddText("Rent: " + strconv.FormatFloat(lease.Property().RentalPricePerMonth, 'f', 2, 64) + "€")

	addRooms(&report, invReport)

	f, err := report.Output()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	f.Filename = "inventory_report_" + invReport.Date.Format("2006-01-02") + "_" + invReport.ID + ".pdf"
	return f, nil
}

func addRooms(report *PDF, invReport db.InventoryReportModel) {
	report.Ln(5)
	report.AddTitle("Rooms:", H2)
	for _, roomState := range invReport.RoomStates() {
		report.AddLine()
		report.AddTitle(roomState.Room().Name, H3)
		report.AddText("State: " + string(roomState.State))
		report.AddText("Cleanliness: " + string(roomState.Cleanliness))
		report.AddMultiLineText("Note: " + roomState.Note)
		report.Ln(5)
		report.AddImages(filesystem.GetImageObjs(roomState.Pictures))

		for _, furnitureState := range invReport.FurnitureStates() {
			if furnitureState.Furniture().RoomID != roomState.RoomID {
				continue
			}
			report.AddTitle("Furniture: "+furnitureState.Furniture().Name+" ("+strconv.Itoa(furnitureState.Furniture().Quantity)+")", H4)
			report.AddText("State: " + string(furnitureState.State))
			report.AddText("Cleanliness: " + string(furnitureState.Cleanliness))
			report.AddMultiLineText("Note: " + furnitureState.Note)
			report.Ln(5)
			report.AddImages(filesystem.GetImageObjs(furnitureState.Pictures))
		}
	}
}
