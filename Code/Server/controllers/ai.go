package controllers

import (
	"encoding/base64"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"immotep/backend/models"
	"immotep/backend/prisma/db"
	chatgptservice "immotep/backend/services/chatgpt"
	furnitureservice "immotep/backend/services/furniture"
	inventoryreportservice "immotep/backend/services/inventoryreport"
	roomservice "immotep/backend/services/room"
	"immotep/backend/utils"
)

func imagesToBase64Strings(images []db.ImageModel) []string {
	res := make([]string, len(images))
	for i, img := range images {
		res[i] = base64.StdEncoding.EncodeToString(img.Data)
	}
	return res
}

// GenerateSummary godoc
//
//	@Summary		Generate summary from photo
//	@Description	Generate summary from photo for first inventory report
//	@Tags			ai
//	@Accept			json
//	@Produce		json
//	@Param			property_id	path		string						true	"Property ID"
//	@Param			aiData		body		models.SummarizeRequest		true	"Summarize data"
//	@Success		201			{object}	models.SummarizeResponse	"Summarize data"
//	@Failure		400			{object}	utils.Error					"Missing fields"
//	@Failure		403			{object}	utils.Error					"Property not yours"
//	@Failure		404			{object}	utils.Error					"Property not found"
//	@Failure		500
//	@Security		Bearer
//	@Router			/owner/properties/{property_id}/inventory-reports/summarize/ [post]
func GenerateSummary(c *gin.Context) {
	var req models.SummarizeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendError(c, http.StatusBadRequest, utils.MissingFields, err)
		return
	}

	if req.Type == "room" {
		room := roomservice.GetByID(req.Id)
		chatGPTres, err := chatgptservice.SummarizeRoom(room.Name, req.Pictures)
		if err != nil {
			utils.SendError(c, http.StatusInternalServerError, utils.ErrorRequestChatGPTAPI, err)
			return
		}
		splitted := strings.Split(chatGPTres, "|")
		resp := models.SummarizeResponse{
			State:       splitted[0],
			Cleanliness: splitted[1],
			Note:        splitted[2],
		}
		c.JSON(http.StatusOK, resp)
	} else {
		furniture := furnitureservice.GetByID(req.Id)
		chatGPTres, err := chatgptservice.SummarizeFurniture(furniture.Name, req.Pictures)
		if err != nil {
			utils.SendError(c, http.StatusInternalServerError, utils.ErrorRequestChatGPTAPI, err)
			return
		}
		splitted := strings.Split(chatGPTres, "|")
		resp := models.SummarizeResponse{
			State:       splitted[0],
			Cleanliness: splitted[1],
			Note:        splitted[2],
		}
		c.JSON(http.StatusOK, resp)
	}
}

// GenerateComparison godoc
//
//	@Summary		Generate comparison from photo
//	@Description	Generate comparison from photo for last inventory report
//	@Tags			ai
//	@Accept			json
//	@Produce		json
//	@Param			property_id		path		string							true	"Property ID"
//	@Param			old_report_id	path		string							true	"Previous report ID to compare with"
//	@Param			aiData			body		models.SummarizeRequest			true	"Compare data"
//	@Success		201				{object}	models.InventoryReportResponse	"Created inventory report data"
//	@Failure		400				{object}	utils.Error						"Missing fields"
//	@Failure		403				{object}	utils.Error						"Property not yours"
//	@Failure		404				{object}	utils.Error						"Property or old report not found"
//	@Failure		500
//	@Security		Bearer
//	@Router			/owner/properties/{property_id}/inventory-reports/compare/{old_report_id}/ [post]
func GenerateComparison(c *gin.Context) {
	var req models.CompareRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendError(c, http.StatusBadRequest, utils.MissingFields, err)
		return
	}

	oldReport := inventoryreportservice.GetByID(c.Param("old_report_id"))
	if oldReport == nil {
		utils.SendError(c, http.StatusNotFound, utils.InventoryReportNotFound, nil)
		return
	}

	if req.Type == "room" {
		handleRoomComparison(c, req, oldReport)
	} else if req.Type == "furniture" {
		handleFurnitureComparison(c, req, oldReport)
	}
}

func handleRoomComparison(c *gin.Context, req models.CompareRequest, oldReport *db.InventoryReportModel) {
	for _, rs := range oldReport.RoomStates() {
		if rs.RoomID == req.Id {
			room := roomservice.GetByID(req.Id)
			chatGPTres, err := chatgptservice.CompareRoom(room.Name, rs, imagesToBase64Strings(rs.Pictures()), req.Pictures)
			if err != nil {
				utils.SendError(c, http.StatusInternalServerError, utils.ErrorRequestChatGPTAPI, err)
				return
			}
			log.Println(chatGPTres)
			splitted := strings.Split(chatGPTres, "|")
			resp := models.SummarizeResponse{
				State:       splitted[0],
				Cleanliness: splitted[1],
				Note:        splitted[2],
			}
			c.JSON(http.StatusOK, resp)
			return
		}
	}
	utils.SendError(c, http.StatusNotFound, utils.RoomNotFound, nil)
}

func handleFurnitureComparison(c *gin.Context, req models.CompareRequest, oldReport *db.InventoryReportModel) {
	for _, fs := range oldReport.FurnitureStates() {
		if fs.FurnitureID == req.Id {
			furniture := furnitureservice.GetByID(req.Id)
			chatGPTres, err := chatgptservice.CompareFurniture(furniture.Name, fs, imagesToBase64Strings(fs.Pictures()), req.Pictures)
			if err != nil {
				utils.SendError(c, http.StatusInternalServerError, utils.ErrorRequestChatGPTAPI, err)
				return
			}
			log.Println(chatGPTres)
			splitted := strings.Split(chatGPTres, "|")
			resp := models.SummarizeResponse{
				State:       splitted[0],
				Cleanliness: splitted[1],
				Note:        splitted[2],
			}
			c.JSON(http.StatusOK, resp)
			return
		}
	}
	utils.SendError(c, http.StatusNotFound, utils.FurnitureNotFound, nil)
}
