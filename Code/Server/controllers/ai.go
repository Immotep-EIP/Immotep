package controllers

import (
	"encoding/base64"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"keyz/backend/models"
	"keyz/backend/prisma/db"
	"keyz/backend/services/chatgpt"
	"keyz/backend/services/database"
	"keyz/backend/utils"
)

const chatGPTerror = "error"

func imagesToBase64Strings(images []db.ImageModel) []string {
	res := make([]string, len(images))
	for i, img := range images {
		res[i] = "data:image/" + string(img.Type) + ";base64," + base64.StdEncoding.EncodeToString(img.Data)
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
//	@Router			/owner/properties/{property_id}/leases/{lease_id}/inventory-reports/summarize/ [post]
func GenerateSummary(c *gin.Context) {
	var req models.SummarizeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendError(c, http.StatusBadRequest, utils.MissingFields, err)
		return
	}

	switch req.Type {
	case "room":
		handleRoomSummary(c, req)
	case "furniture":
		handleFurnitureSummary(c, req)
	}
}

func handleRoomSummary(c *gin.Context, req models.SummarizeRequest) {
	room := database.GetRoomByID(req.Id)
	chatGPTres, err := chatgpt.SummarizeRoom(room.Name, req.Pictures)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, utils.ErrorRequestChatGPTAPI, err)
		return
	}
	splitted := strings.Split(chatGPTres, "|")
	if len(splitted) == 2 && splitted[0] == chatGPTerror {
		utils.SendError(c, http.StatusBadRequest, utils.ErrorRequestChatGPTAPI, errors.New(splitted[1]))
		return
	} else if len(splitted) != 3 {
		log.Println(chatGPTres)
		utils.SendError(c, http.StatusInternalServerError, utils.ErrorRequestChatGPTAPI, errors.New("unexpected response format from ChatGPT"))
		return
	}
	resp := models.SummarizeResponse{
		State:       splitted[0],
		Cleanliness: splitted[1],
		Note:        splitted[2],
	}
	c.JSON(http.StatusOK, resp)
}

func handleFurnitureSummary(c *gin.Context, req models.SummarizeRequest) {
	furniture := database.GetFurnitureByID(req.Id)
	chatGPTres, err := chatgpt.SummarizeFurniture(furniture.Name, req.Pictures)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, utils.ErrorRequestChatGPTAPI, err)
		return
	}
	splitted := strings.Split(chatGPTres, "|")
	if len(splitted) == 2 && splitted[0] == chatGPTerror {
		utils.SendError(c, http.StatusBadRequest, utils.ErrorRequestChatGPTAPI, errors.New(splitted[1]))
		return
	} else if len(splitted) != 3 {
		log.Println(chatGPTres)
		utils.SendError(c, http.StatusInternalServerError, utils.ErrorRequestChatGPTAPI, errors.New("unexpected response format from ChatGPT"))
		return
	}
	resp := models.SummarizeResponse{
		State:       splitted[0],
		Cleanliness: splitted[1],
		Note:        splitted[2],
	}
	c.JSON(http.StatusOK, resp)
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
//	@Router			/owner/properties/{property_id}/leases/{lease_id}/inventory-reports/compare/{old_report_id}/ [post]
func GenerateComparison(c *gin.Context) {
	var req models.CompareRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendError(c, http.StatusBadRequest, utils.MissingFields, err)
		return
	}

	oldReport := database.GetInvReportByID(c.Param("old_report_id"))
	if oldReport == nil {
		utils.SendError(c, http.StatusNotFound, utils.InventoryReportNotFound, nil)
		return
	}

	switch req.Type {
	case "room":
		handleRoomComparison(c, req, oldReport)
	case "furniture":
		handleFurnitureComparison(c, req, oldReport)
	}
}

func handleRoomComparison(c *gin.Context, req models.CompareRequest, oldReport *db.InventoryReportModel) {
	for _, rs := range oldReport.RoomStates() {
		if rs.RoomID == req.Id {
			handleRoomComparison2(c, req, rs)
			return
		}
	}
	utils.SendError(c, http.StatusNotFound, utils.RoomNotFound, nil)
}

func handleRoomComparison2(c *gin.Context, req models.CompareRequest, rs db.RoomStateModel) {
	room := database.GetRoomByID(req.Id)
	chatGPTres, err := chatgpt.CompareRoom(room.Name, rs, imagesToBase64Strings(rs.Pictures()), req.Pictures)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, utils.ErrorRequestChatGPTAPI, err)
		return
	}
	splitted := strings.Split(chatGPTres, "|")
	if len(splitted) == 2 && splitted[0] == chatGPTerror {
		utils.SendError(c, http.StatusBadRequest, utils.ErrorRequestChatGPTAPI, errors.New(splitted[1]))
		return
	} else if len(splitted) != 3 {
		log.Println(chatGPTres)
		utils.SendError(c, http.StatusInternalServerError, utils.ErrorRequestChatGPTAPI, errors.New("unexpected response format from ChatGPT"))
		return
	}
	resp := models.SummarizeResponse{
		State:       splitted[0],
		Cleanliness: splitted[1],
		Note:        splitted[2],
	}
	c.JSON(http.StatusOK, resp)
}

func handleFurnitureComparison(c *gin.Context, req models.CompareRequest, oldReport *db.InventoryReportModel) {
	for _, fs := range oldReport.FurnitureStates() {
		if fs.FurnitureID == req.Id {
			handleFurnitureComparison2(c, req, fs)
			return
		}
	}
	utils.SendError(c, http.StatusNotFound, utils.FurnitureNotFound, nil)
}

func handleFurnitureComparison2(c *gin.Context, req models.CompareRequest, fs db.FurnitureStateModel) {
	furniture := database.GetFurnitureByID(req.Id)
	chatGPTres, err := chatgpt.CompareFurniture(furniture.Name, fs, imagesToBase64Strings(fs.Pictures()), req.Pictures)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, utils.ErrorRequestChatGPTAPI, err)
		return
	}
	splitted := strings.Split(chatGPTres, "|")
	if len(splitted) == 2 && splitted[0] == chatGPTerror {
		utils.SendError(c, http.StatusBadRequest, utils.ErrorRequestChatGPTAPI, errors.New(splitted[1]))
		return
	} else if len(splitted) != 3 {
		log.Println(chatGPTres)
		utils.SendError(c, http.StatusInternalServerError, utils.ErrorRequestChatGPTAPI, errors.New("unexpected response format from ChatGPT"))
		return
	}
	resp := models.SummarizeResponse{
		State:       splitted[0],
		Cleanliness: splitted[1],
		Note:        splitted[2],
	}
	c.JSON(http.StatusOK, resp)
}
