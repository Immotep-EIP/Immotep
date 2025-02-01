package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"immotep/backend/models"
	"immotep/backend/prisma/db"
	furnitureservice "immotep/backend/services/furniture"
	imageservice "immotep/backend/services/image"
	inventoryreportservice "immotep/backend/services/inventoryreport"
	roomservice "immotep/backend/services/room"
	"immotep/backend/utils"
)

func checkRoom(roomId string, propertyId string) error {
	room := roomservice.GetByID(roomId)
	if room == nil {
		return errors.New(string(utils.RoomNotFound))
	}
	if room.PropertyID != propertyId {
		return errors.New(string(utils.PropertyNotYours))
	}
	return nil
}

func checkFurniture(furnitureId string, roomId string) error {
	furniture := furnitureservice.GetByID(furnitureId)
	if furniture == nil {
		return errors.New(string(utils.FurnitureNotFound))
	}
	if furniture.RoomID != roomId {
		return errors.New(string(utils.FurnitureNotInThisRoom))
	}
	return nil
}

func getFurnitureStatePictures(f models.FurnitureStateRequest) ([]string, []string) {
	picturesId := make([]string, 0, len(f.Pictures))
	var errorList []string
	for _, pic := range f.Pictures {
		dbImage := models.StringToDbImage(pic)
		if dbImage == nil {
			errorList = append(errorList, string(utils.BadBase64String))
			continue
		}
		newImage := imageservice.Create(*dbImage)
		picturesId = append(picturesId, newImage.ID)
	}
	return picturesId, errorList
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
		picturesId, el := getFurnitureStatePictures(f)
		errorList = append(errorList, el...)
		inventoryreportservice.CreateFurnitureState(fModel, picturesId, invrep.ID)
	}

	return errorList
}

func getRoomStatePictures(r models.RoomStateRequest) ([]string, []string) {
	picturesId := make([]string, 0, len(r.Pictures))
	var errorList []string
	for _, pic := range r.Pictures {
		dbImage := models.StringToDbImage(pic)
		if dbImage == nil {
			errorList = append(errorList, string(utils.BadBase64String))
			continue
		}
		newImage := imageservice.Create(*dbImage)
		picturesId = append(picturesId, newImage.ID)
	}
	return picturesId, errorList
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
		picturesId, el := getRoomStatePictures(r)
		errorList = append(errorList, el...)
		inventoryreportservice.CreateRoomState(rModel, picturesId, invrep.ID)
		errorList = append(errorList, createFurnitureState(invrep, r)...)
	}

	return errorList
}

// CreateInventoryReport godoc
//
//	@Summary		Create a new inventory report
//	@Description	Create a new inventory report for a room
//	@Tags			owner
//	@Accept			json
//	@Produce		json
//	@Param			property_id	path		string							true	"Property ID"
//	@Param			invReport	body		models.InventoryReportRequest	true	"Inventory report data"
//	@Success		201			{object}	models.InventoryReportResponse	"Created inventory report data"
//	@Failure		400			{object}	utils.Error						"Missing fields"
//	@Failure		403			{object}	utils.Error						"Property not yours"
//	@Failure		404			{object}	utils.Error						"Property or room not found"
//	@Failure		500
//	@Security		Bearer
//	@Router			/owner/properties/{property_id}/inventory-reports/ [post]
func CreateInventoryReport(c *gin.Context) {
	var req models.InventoryReportRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendError(c, http.StatusBadRequest, utils.MissingFields, err)
		return
	}

	invrep := inventoryreportservice.Create(db.ReportType(req.Type), c.Param("property_id"))
	if invrep == nil {
		utils.SendError(c, http.StatusConflict, utils.InventoryReportAlreadyExists, nil)
		return
	}

	errorsList := createRoomStates(c, invrep, req)

	c.JSON(http.StatusCreated, errorsList)
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
	reports := inventoryreportservice.GetByPropertyID(c.Param("property_id"))
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
		report = inventoryreportservice.GetLatest(c.Param("property_id"))
	} else {
		report = inventoryreportservice.GetByID(c.Param("report_id"))
	}
	c.JSON(http.StatusOK, models.DbInventoryReportToResponse(*report))
}
