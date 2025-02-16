package pdf

import (
	"errors"
	"log"
	"strconv"

	"immotep/backend/prisma/db"
	"immotep/backend/services/database"
	"immotep/backend/utils"
)

func NewInventoryReportPDF(invReportId string) ([]byte, error) {
	invReport := database.GetInvReportByID(invReportId)
	if invReport == nil {
		log.Println("No invReport found with id: " + invReportId)
		return nil, errors.New(string(utils.InventoryReportNotFound))
	}
	contract := database.GetCurrentActiveContractWithInfos(invReport.PropertyID)
	if contract == nil {
		log.Println("No active contract found for property with id: " + invReport.PropertyID)
		return nil, errors.New(string(utils.NoActiveContract))
	}

	report := NewPDF()

	report.AddCenteredTitle("Inventory Report", H1)
	report.AddText("ID: " + invReport.ID)
	report.AddText("Date: " + invReport.Date.Format("2006-01-02 15:04:05"))
	report.AddText("Property: " + contract.Property().Name)
	report.AddText("Type: " + string(invReport.Type))

	report.Ln(5)
	report.AddTitle("Lease", H2)
	report.Add2Texts("Owner: "+contract.Property().Owner().Firstname+" "+contract.Property().Owner().Lastname, "Email: "+contract.Property().Owner().Email)
	report.Add2Texts("Tenant: "+contract.Tenant().Firstname+" "+contract.Tenant().Lastname, "Email: "+contract.Tenant().Email)
	contractEndDate, ok := contract.EndDate()
	if ok {
		report.Add2Texts("Start date: "+contract.StartDate.Format("2006-01-02"), "End date: "+contractEndDate.Format("2006-01-02"))
	} else {
		report.Add2Texts("Start date: "+contract.StartDate.Format("2006-01-02"), "End date: None")
	}
	report.AddText("Rent: " + strconv.Itoa(contract.Property().RentalPricePerMonth) + "€")

	addRooms(&report, invReport)

	bytes, err := report.Output()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return bytes, nil
}

func addRooms(report *PDF, invReport *db.InventoryReportModel) {
	report.Ln(5)
	report.AddTitle("Rooms:", H2)
	for _, roomState := range invReport.RoomStates() {
		report.AddLine()
		report.AddTitle(roomState.Room().Name, H3)
		report.AddText("State: " + string(roomState.State))
		report.AddText("Cleanliness: " + string(roomState.Cleanliness))
		report.AddMultiLineText("Note: " + roomState.Note)
		report.Ln(5)
		report.AddImages(roomState.Pictures())

		for _, furnitureState := range invReport.FurnitureStates() {
			if furnitureState.Furniture().RoomID != roomState.RoomID {
				continue
			}
			report.AddTitle("Furniture: "+furnitureState.Furniture().Name+" ("+strconv.Itoa(furnitureState.Furniture().Quantity)+")", H4)
			report.AddText("State: " + string(furnitureState.State))
			report.AddText("Cleanliness: " + string(furnitureState.Cleanliness))
			report.AddMultiLineText("Note: " + furnitureState.Note)
			report.Ln(5)
			report.AddImages(furnitureState.Pictures())
		}
	}
}
