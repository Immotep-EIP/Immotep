package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"immotep/backend/models"
	contractservice "immotep/backend/services/contract"
	propertyservice "immotep/backend/services/property"
	"immotep/backend/utils"
)

// InviteTenant godoc
//
//	@Summary		Invite tenant to owner's property
//	@Description	Invite tenant to owner's property
//	@Tags			owner
//	@Accept			json
//	@Produce		json
//	@Param			propertyId	path		string					true	"Property ID"
//	@Param			user		body		models.InviteRequest	true	"Invite params"
//	@Success		200			{object}	models.InviteResponse	"Created invite"
//	@Failure		400			{object}	utils.Error				"Missing fields"
//	@Failure		403			{object}	utils.Error				"Property is not yours"
//	@Failure		404			{object}	utils.Error				"Property not found"
//	@Failure		409			{object}	utils.Error				"Invite already exists for this email"
//	@Failure		500
//	@Security		Bearer
//	@Router			/owner/send-invite/{propertyId} [post]
func InviteTenant(c *gin.Context) {
	claims := utils.GetClaims(c)

	var inviteReq models.InviteRequest
	err := c.ShouldBindBodyWithJSON(&inviteReq)
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, utils.MissingFields, err)
		return
	}

	property := propertyservice.GetByID(c.Param("propertyId"))
	if property == nil {
		utils.SendError(c, http.StatusNotFound, utils.PropertyNotFound, nil)
		return
	}
	if property.OwnerID != claims["id"] {
		utils.SendError(c, http.StatusForbidden, utils.PropertyNotYours, nil)
		return
	}

	pendingContract := contractservice.CreatePending(inviteReq.ToDbPendingContract(), *property)
	if pendingContract == nil {
		utils.SendError(c, http.StatusConflict, utils.InviteAlreadyExists, nil)
		return
	}

	// TODO send email

	c.JSON(http.StatusOK, models.DbPendingContractToResponse(*pendingContract))
}
