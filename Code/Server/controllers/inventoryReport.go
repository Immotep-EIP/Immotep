package controllers

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"immotep/backend/models"
	"immotep/backend/prisma/db"
	"immotep/backend/services/database"
	"immotep/backend/services/pdf"
	"immotep/backend/utils"
)

func checkRoom(roomId string, propertyId string) error {
	room := database.GetRoomByID(roomId)
	if room == nil {
		return errors.New(string(utils.RoomNotFound))
	}
	if room.PropertyID != propertyId {
		return errors.New(string(utils.PropertyNotYours))
	}
	return nil
}

func checkFurniture(furnitureId string, roomId string) error {
	furniture := database.GetFurnitureByID(furnitureId)
	if furniture == nil {
		return errors.New(string(utils.FurnitureNotFound))
	}
	if furniture.RoomID != roomId {
		return errors.New(string(utils.FurnitureNotInThisRoom))
	}
	return nil
}

func getFurnitureStatePictures(f models.FurnitureStateRequest) ([]string, error) {
	picturesId := make([]string, 0, len(f.Pictures))
	var err error

	for _, pic := range f.Pictures {
		dbImage := models.StringToDbImage(pic)
		if dbImage == nil {
			err = errors.New(string(utils.BadBase64String))
			continue
		}
		newImage := database.CreateImage(*dbImage)
		picturesId = append(picturesId, newImage.ID)
	}
	return picturesId, err
}

func createFurnitureState(invrep *db.InventoryReportModel, room models.RoomStateRequest) []string {
	var errorList []string

	for _, f := range room.Furnitures {
		if err := checkFurniture(f.ID, room.ID); err != nil {
			errorList = append(errorList, err.Error())
			continue
		}

		fModel := db.FurnitureStateModel{
			InnerFurnitureState: db.InnerFurnitureState{
				FurnitureID: f.ID,
				ReportID:    invrep.ID,
				Cleanliness: db.Cleanliness(f.Cleanliness),
				State:       db.State(f.State),
				Note:        f.Note,
			},
		}
		picturesId, err := getFurnitureStatePictures(f)
		if err != nil {
			errorList = append(errorList, err.Error())
		}
		database.CreateFurnitureState(fModel, picturesId, invrep.ID)
	}

	return errorList
}

func getRoomStatePictures(r models.RoomStateRequest) ([]string, error) {
	picturesId := make([]string, 0, len(r.Pictures))
	var err error

	for _, pic := range r.Pictures {
		dbImage := models.StringToDbImage(pic)
		if dbImage == nil {
			err = errors.New(string(utils.BadBase64String))
			continue
		}
		newImage := database.CreateImage(*dbImage)
		picturesId = append(picturesId, newImage.ID)
	}
	return picturesId, err
}

func createRoomStates(c *gin.Context, invrep *db.InventoryReportModel, req models.InventoryReportRequest) []string {
	var errorList []string

	for _, r := range req.Rooms {
		if err := checkRoom(r.ID, c.Param("property_id")); err != nil {
			errorList = append(errorList, err.Error())
			continue
		}

		rModel := db.RoomStateModel{
			InnerRoomState: db.InnerRoomState{
				RoomID:      r.ID,
				ReportID:    invrep.ID,
				Cleanliness: db.Cleanliness(r.Cleanliness),
				State:       db.State(r.State),
				Note:        r.Note,
			},
		}
		picturesId, err := getRoomStatePictures(r)
		if err != nil {
			errorList = append(errorList, err.Error())
		}
		database.CreateRoomState(rModel, picturesId, invrep.ID)
		errorList = append(errorList, createFurnitureState(invrep, r)...)
	}

	return errorList
}

func createInvReportPDF(invRepId string, contract db.ContractModel) (*db.DocumentModel, error) {
	invReport := database.GetInvReportByID(invRepId)
	docBytes, err := pdf.NewInventoryReportPDF(*invReport, contract)
	if err != nil || docBytes == nil {
		return nil, err
	}

	res := database.CreateDocument(db.DocumentModel{
		InnerDocument: db.InnerDocument{
			Name:       "inventory_report_" + time.Now().Format("2006-01-02") + "_" + invRepId + ".pdf",
			Data:       docBytes,
			ContractID: contract.ID,
		},
	})
	return &res, nil
}

// CreateInventoryReport godoc
//
//	@Summary		Create a new inventory report
//	@Description	Create a new inventory report for a room
//	@Tags			owner
//	@Accept			json
//	@Produce		json
//	@Param			property_id	path		string									true	"Property ID"
//	@Param			invReport	body		models.InventoryReportRequest			true	"Inventory report data"
//	@Success		201			{object}	models.CreateInventoryReportResponse	"Created inventory report data"
//	@Failure		400			{object}	utils.Error								"Missing fields"
//	@Failure		403			{object}	utils.Error								"Property not yours"
//	@Failure		404			{object}	utils.Error								"Property or room not found"
//	@Failure		500
//	@Security		Bearer
//	@Router			/owner/properties/{property_id}/inventory-reports/ [post]
func CreateInventoryReport(c *gin.Context) {
	var req models.InventoryReportRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendError(c, http.StatusBadRequest, utils.MissingFields, err)
		return
	}

	propertyId := c.Param("property_id")

	contract := database.GetCurrentActiveContractWithInfos(propertyId)
	if contract == nil {
		utils.SendError(c, http.StatusNotFound, utils.NoActiveContract, nil)
		return
	}

	invrep := database.CreateInvReport(db.ReportType(req.Type), propertyId)
	if invrep == nil {
		utils.SendError(c, http.StatusConflict, utils.InventoryReportAlreadyExists, nil)
		return
	}

	errorsList := createRoomStates(c, invrep, req)

	irPdf, err := createInvReportPDF(invrep.ID, *contract)
	if err != nil {
		errorsList = append(errorsList, err.Error())
	}

	c.JSON(http.StatusCreated, models.DbInventoryReportToCreateResponse(*invrep, irPdf, errorsList))
}

// GetInventoryReportsByProperty godoc
//
//	@Summary		Get all inventory reports for a property
//	@Description	Get all inventory reports for a property
//	@Tags			owner
//	@Accept			json
//	@Produce		json
//	@Param			property_id	path		string							true	"Property ID"
//	@Success		200			{array}		models.InventoryReportResponse	"List of inventory reports"
//	@Failure		403			{object}	utils.Error						"Property not found"
//	@Failure		404			{object}	utils.Error						"Property not found"
//	@Failure		500
//	@Security		Bearer
//	@Router			/owner/properties/{property_id}/inventory-reports/ [get]
func GetInventoryReportsByProperty(c *gin.Context) {
	reports := database.GetInvReportByPropertyID(c.Param("property_id"))
	c.JSON(http.StatusOK, utils.Map(reports, models.DbInventoryReportToResponse))
}

// GetInventoryReportByID godoc
//
//	@Summary		Get inventory report by ID
//	@Description	Get inventory report information by its ID or get the latest one
//	@Tags			owner
//	@Accept			json
//	@Produce		json
//	@Param			property_id	path		string							true	"Property ID"
//	@Param			report_id	path		string							true	"Report ID or 'latest' to get the latest one"
//	@Success		200			{object}	models.InventoryReportResponse	"Inventory report data"
//	@Failure		403			{object}	utils.Error						"Property not yours"
//	@Failure		404			{object}	utils.Error						"Inventory report not found"
//	@Failure		500
//	@Security		Bearer
//	@Router			/owner/properties/{property_id}/inventory-reports/{report_id}/ [get]
func GetInventoryReportByID(c *gin.Context) {
	var report *db.InventoryReportModel
	if c.Param("report_id") == "latest" {
		report = database.GetLatestInvReport(c.Param("property_id"))
	} else {
		report = database.GetInvReportByID(c.Param("report_id"))
	}
	c.JSON(http.StatusOK, models.DbInventoryReportToResponse(*report))
}
