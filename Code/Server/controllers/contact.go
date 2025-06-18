package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"keyz/backend/models"
	"keyz/backend/services/brevo"
	"keyz/backend/services/database"
	"keyz/backend/utils"
)

// CreateContactMessage godoc
//
//	@Summary		Create contact message
//	@Description	Create a contact message to a lease
//	@Tags			contact-message
//	@Accept			json
//	@Produce		json
//	@Param			contactMessages	body		models.ContactMessageRequest	true	"Message to create"
//	@Success		201				{object}	models.IdResponse				"Created message ID"
//	@Failure		400				{object}	utils.Error						"Missing fields or bad base64 string"
//	@Failure		500
//	@Router			/contact/ [post]
func CreateContactMessage(c *gin.Context) {
	var req models.ContactMessageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendError(c, http.StatusBadRequest, utils.MissingFields, err)
		return
	}

	cm := database.CreateContactMessage(req.ToDbContact())

	res, err := brevo.SendNewContactMessage(cm)
	if err != nil {
		log.Println(res, err.Error())
		utils.SendError(c, http.StatusInternalServerError, utils.FailedSendEmail, err)
		return
	}

	c.JSON(http.StatusCreated, models.IdResponse{ID: cm.ID})
}
