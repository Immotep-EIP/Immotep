package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"immotep/backend/models"
	"immotep/backend/prisma/db"
	"immotep/backend/services/database"
	"immotep/backend/utils"
)

// GetLeaseDocuments godoc
//
//	@Summary		Get property documents
//	@Description	Get all documents of a lease related to a property
//	@Tags			owner
//	@Accept			json
//	@Produce		json
//	@Param			property_id	path		string					true	"Property ID"
//	@Param			lease_id	path		string					true	"Lease ID or `current`"
//	@Success		200			{array}		models.DocumentResponse	"List of documents"
//	@Failure		403			{object}	utils.Error				"Property not yours"
//	@Failure		404			{object}	utils.Error				"No active lease"
//	@Failure		500
//	@Security		Bearer
//	@Router			/owner/properties/{property_id}/leases/{lease_id}/documents/ [get]
func GetLeaseDocuments(c *gin.Context) {
	lease, _ := c.MustGet("lease").(db.LeaseModel)
	documents := database.GetLeaseDocuments(lease.ID)
	c.JSON(http.StatusOK, utils.Map(documents, models.DbDocumentToResponse))
}
