package controllers

import (
	"errors"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"immotep/backend/models"
	"immotep/backend/prisma/db"
	"immotep/backend/services/database"
	"immotep/backend/services/minio"
	"immotep/backend/services/pdf"
	"immotep/backend/utils"
)

func createInvReportPDF(invrep db.InventoryReportModel, lease db.LeaseModel) (*models.DocumentResponse, error) {
	file, err := pdf.NewInventoryReportPDF(invrep, lease)
	if err != nil || file == nil {
		return nil, err
	}

	fileInfo := minio.UploadLeasePDF(lease.ID, file)
	database.AddDocumentToLease(lease, fileInfo.Key)

	res := minio.GetDocument(fileInfo.Key)
	if res == nil {
		panic("error getting document")
	}
	return res, nil
}

// CreateInventoryReport godoc
//
//	@Summary		Create a new inventory report
//	@Description	Create a new inventory report for a room
//	@Tags			inventory-report
//	@Accept			json
//	@Produce		json
//	@Param			property_id	path		string							true	"Property ID"
//	@Param			lease_id	path		string							true	"Lease ID"
//	@Param			invReport	body		models.InventoryReportRequest	true	"Inventory report data"
//	@Success		201			{object}	models.IdResponse				"Created inventory report ID"
//	@Failure		400			{object}	utils.Error						"Missing fields"
//	@Failure		403			{object}	utils.Error						"Property not yours"
//	@Failure		404			{object}	utils.Error						"Property or room not found"
//	@Failure		500
//	@Security		Bearer
//	@Router			/owner/properties/{property_id}/leases/{lease_id}/inventory-reports/ [post]
func CreateInventoryReport(c *gin.Context) {
	var req models.InventoryReportRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendError(c, http.StatusBadRequest, utils.MissingFields, err)
		return
	}

	if c.Param("lease_id") != "current" {
		utils.SendError(c, http.StatusBadRequest, utils.InvReportMustBeCurrentLease, nil)
		return
	}

	lease, _ := c.MustGet("lease").(db.LeaseModel)
	invrep := database.CreateInvReport(req.Type, lease.ID)
	if invrep == nil {
		utils.SendError(c, http.StatusConflict, utils.InventoryReportAlreadyExists, nil)
		return
	}

	c.JSON(http.StatusCreated, models.IdResponse{ID: invrep.ID})
}

// SubmitInventoryReport godoc
//
//	@Summary		Submit an inventory report
//	@Description	Submit an inventory report
//	@Tags			inventory-report
//	@Accept			json
//	@Produce		json
//	@Param			property_id	path		string									true	"Property ID"
//	@Param			report_id	path		string									true	"Report ID"
//	@Success		200			{object}	models.CreateInventoryReportResponse	"Inventory report data with PDF"
//	@Failure		403			{object}	utils.Error								"Property not yours"
//	@Failure		404			{object}	utils.Error								"Property or inventory report not found"
//	@Failure		500
//	@Security		Bearer
//	@Router			/owner/properties/{property_id}/leases/{lease_id}/inventory-reports/{report_id}/submit/ [post]
func SubmitInventoryReport(c *gin.Context) {
	lease, _ := c.MustGet("lease").(db.LeaseModel)
	invrep, _ := c.MustGet("invrep").(db.InventoryReportModel)

	if invrep.Submitted {
		utils.SendError(c, http.StatusForbidden, utils.CannotModifyInventoryReport, errors.New("inventory report already submitted"))
		return
	}

	invRepPdf, err := createInvReportPDF(invrep, lease)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, utils.ErrorCode(err.Error()), err)
		return
	}
	newInvRep := database.SubmitInventoryReport(invrep)
	c.JSON(http.StatusOK, models.DbInventoryReportToCreateResponse(newInvRep, *invRepPdf))
}

// AddRoomStateToInventoryReport godoc
//
//	@Summary		Add a room state to an inventory report
//	@Description	Add a room state to an inventory report
//	@Tags			inventory-report
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			property_id	path		string				true	"Property ID"
//	@Param			report_id	path		string				true	"Report ID"
//	@Param			room_id		formData	string				true	"Room ID"
//	@Param			state		formData	string				true	"Room state"
//	@Param			cleanliness	formData	string				true	"Room cleanliness"
//	@Param			note		formData	string				true	"Room note"
//	@Param			pictures	formData	[]file				true	"Room pictures"
//	@Success		201			{object}	models.IdResponse	"Created room state ID"
//	@Failure		400			{object}	utils.Error			"Missing fields"
//	@Failure		403			{object}	utils.Error			"Property not yours"
//	@Failure		404			{object}	utils.Error			"Property or room not found"
//	@Failure		409			{object}	utils.Error			"Room state already exists"
//	@Failure		500
//	@Security		Bearer
//	@Router			/owner/properties/{property_id}/leases/{lease_id}/inventory-reports/{report_id}/rooms/ [post]
func AddRoomStateToInventoryReport(c *gin.Context) {
	property, _ := c.MustGet("property").(db.PropertyModel)
	invrep, _ := c.MustGet("invrep").(db.InventoryReportModel)

	if invrep.Submitted {
		utils.SendError(c, http.StatusForbidden, utils.CannotModifyInventoryReport, errors.New("inventory report already submitted"))
		return
	}

	req := models.RoomStateRequest{
		RoomID:      c.PostForm("room_id"),
		State:       db.State(c.PostForm("state")),
		Cleanliness: db.Cleanliness(c.PostForm("cleanliness")),
		Note:        c.PostForm("note"),
	}
	if err := binding.Validator.ValidateStruct(req); err != nil {
		utils.SendError(c, http.StatusBadRequest, utils.MissingFields, err)
		return
	}
	if err := checkRoom(req.RoomID, property.ID); err != nil {
		utils.SendError(c, http.StatusBadRequest, utils.ErrorCode(err.Error()), nil)
		return
	}

	form, _ := c.MultipartForm()
	recPictures := form.File["pictures"]
	if len(recPictures) == 0 {
		utils.SendError(c, http.StatusBadRequest, utils.MissingFile, nil)
		return
	}

	roomState := database.CreateRoomState(req.ToDbRoomState(), invrep.ID)
	if roomState == nil {
		utils.SendError(c, http.StatusConflict, utils.RoomStateAlreadyExists, nil)
		return
	}
	picturePaths := getRoomStatePicturesPath(*roomState, recPictures)
	newRoomState := database.AddPicturesToRoomState(*roomState, picturePaths)
	c.JSON(http.StatusCreated, models.IdResponse{ID: newRoomState.ID})
}

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

func getRoomStatePicturesPath(roomState db.RoomStateModel, files []*multipart.FileHeader) []string {
	picturePaths := make([]string, len(files))
	for i, file := range files {
		fileInfo := minio.UploadRoomStateImage(roomState.ID, file)
		picturePaths[i] = fileInfo.Key
	}
	return picturePaths
}

// AddFurnitureStateToInventoryReport godoc
//
//	@Summary		Add a furniture state to an inventory report
//	@Description	Add a furniture state to an inventory report
//	@Tags			inventory-report
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			property_id		path		string				true	"Property ID"
//	@Param			report_id		path		string				true	"Report ID"
//	@Param			furniture_id	formData	string				true	"Furniture ID"
//	@Param			state			formData	string				true	"Furniture state"
//	@Param			cleanliness		formData	string				true	"Furniture cleanliness"
//	@Param			note			formData	string				true	"Furniture note"
//	@Param			pictures		formData	[]file				true	"Furniture pictures"
//	@Success		201				{object}	models.IdResponse	"Created furniture state ID"
//	@Failure		400				{object}	utils.Error			"Missing fields"
//	@Failure		403				{object}	utils.Error			"Property not yours"
//	@Failure		404				{object}	utils.Error			"Property or furniture not found"
//	@Failure		409				{object}	utils.Error			"Furniture state already exists"
//	@Failure		500
//	@Security		Bearer
//	@Router			/owner/properties/{property_id}/leases/{lease_id}/inventory-reports/{report_id}/furnitures/ [post]
func AddFurnitureStateToInventoryReport(c *gin.Context) {
	property, _ := c.MustGet("property").(db.PropertyModel)
	invrep, _ := c.MustGet("invrep").(db.InventoryReportModel)

	if invrep.Submitted {
		utils.SendError(c, http.StatusForbidden, utils.CannotModifyInventoryReport, errors.New("inventory report already submitted"))
		return
	}

	req := models.FurnitureStateRequest{
		FurnitureID: c.PostForm("furniture_id"),
		State:       db.State(c.PostForm("state")),
		Cleanliness: db.Cleanliness(c.PostForm("cleanliness")),
		Note:        c.PostForm("note"),
	}
	if err := binding.Validator.ValidateStruct(req); err != nil {
		utils.SendError(c, http.StatusBadRequest, utils.MissingFields, err)
		return
	}
	if err := checkFurniture(req.FurnitureID, property.ID); err != nil {
		utils.SendError(c, http.StatusBadRequest, utils.ErrorCode(err.Error()), nil)
		return
	}

	form, _ := c.MultipartForm()
	recPictures := form.File["pictures"]
	if len(recPictures) == 0 {
		utils.SendError(c, http.StatusBadRequest, utils.MissingFile, nil)
		return
	}

	furnitureState := database.CreateFurnitureState(req.ToDbFurnitureState(), invrep.ID)
	if furnitureState == nil {
		utils.SendError(c, http.StatusConflict, utils.FurnitureStateAlreadyExists, nil)
		return
	}
	picturePaths := getFurnitureStatePicturesPath(*furnitureState, recPictures)
	newFurnitureState := database.AddPicturesToFurnitureState(*furnitureState, picturePaths)
	c.JSON(http.StatusCreated, models.IdResponse{ID: newFurnitureState.ID})
}

func checkFurniture(furnitureId string, propertyId string) error {
	furniture := database.GetFurnitureByID(furnitureId)
	if furniture == nil {
		return errors.New(string(utils.FurnitureNotFound))
	}
	if furniture.Room().PropertyID != propertyId {
		return errors.New(string(utils.FurnitureNotInThisProperty))
	}
	return nil
}

func getFurnitureStatePicturesPath(furnitureState db.FurnitureStateModel, files []*multipart.FileHeader) []string {
	picturePaths := make([]string, len(files))
	for i, file := range files {
		fileInfo := minio.UploadFurnitureStateImage(furnitureState.ID, file)
		picturePaths[i] = fileInfo.Key
	}
	return picturePaths
}

// GetAllInventoryReportsByProperty godoc
//
//	@Summary		Get all inventory reports for a property
//	@Description	Get all inventory reports for a property
//	@Tags			inventory-report
//	@Accept			json
//	@Produce		json
//	@Param			property_id	path		string							true	"Property ID"
//	@Success		200			{array}		models.InventoryReportResponse	"List of inventory reports"
//	@Failure		403			{object}	utils.Error						"Property not found"
//	@Failure		404			{object}	utils.Error						"Property not found"
//	@Failure		500
//	@Security		Bearer
//	@Router			/owner/properties/{property_id}/inventory-reports/ [get]
func GetAllInventoryReportsByProperty(c *gin.Context) {
	property, _ := c.MustGet("property").(db.PropertyModel)
	reports := database.GetInvReportsByPropertyID(property.ID)
	c.JSON(http.StatusOK, utils.Map(reports, func(report db.InventoryReportModel) models.InventoryReportResponse {
		return models.DbInventoryReportToResponse(report, fetchImageURLs(report))
	}))
}

// GetAllInventoryReportsByLease godoc
//
//	@Summary		Get all inventory reports for a lease
//	@Description	Get all inventory reports for a lease
//	@Tags			inventory-report
//	@Accept			json
//	@Produce		json
//	@Param			property_id	path		string							true	"Property ID"
//	@Param			lease_id	path		string							true	"Lease ID"
//	@Success		200			{array}		models.InventoryReportResponse	"List of inventory reports"
//	@Failure		403			{object}	utils.Error						"Property not found"
//	@Failure		404			{object}	utils.Error						"Property not found"
//	@Failure		500
//	@Security		Bearer
//	@Router			/owner/properties/{property_id}/leases/{lease_id}/inventory-reports/ [get]
//	@Router			/tenant/leases/{lease_id}/inventory-reports/ [get]
func GetInventoryReportsByLease(c *gin.Context) {
	lease, _ := c.MustGet("lease").(db.LeaseModel)
	reports := database.GetInvReportsByLeaseID(lease.ID)
	c.JSON(http.StatusOK, utils.Map(reports, func(report db.InventoryReportModel) models.InventoryReportResponse {
		return models.DbInventoryReportToResponse(report, fetchImageURLs(report))
	}))
}

// GetInventoryReport godoc
//
//	@Summary		Get inventory report by ID
//	@Description	Get inventory report information by its ID or get the latest one
//	@Tags			inventory-report
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
//	@Router			/owner/properties/{property_id}/leases/{lease_id}/inventory-reports/{report_id}/ [get]
//	@Router			/tenant/leases/{lease_id}/inventory-reports/{report_id}/ [get]
func GetInventoryReport(c *gin.Context) {
	report, _ := c.MustGet("invrep").(db.InventoryReportModel)
	c.JSON(http.StatusOK, models.DbInventoryReportToResponse(report, fetchImageURLs(report)))
}

func fetchImageURLs(report db.InventoryReportModel) map[string]string {
	res := make(map[string]string)
	for _, rs := range report.RoomStates() {
		for _, path := range rs.Pictures {
			res[path] = minio.GetImageURL(path)
		}
	}
	for _, fs := range report.FurnitureStates() {
		for _, path := range fs.Pictures {
			res[path] = minio.GetImageURL(path)
		}
	}
	return res
}
