package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"immotep/backend/models"
	"immotep/backend/prisma/db"
	"immotep/backend/services/database"
	"immotep/backend/utils"
)

// UploadDocument godoc
//
//	@Summary		Upload document
//	@Description	Upload a document to a lease
//	@Tags			document
//	@Accept			json
//	@Produce		json
//	@Param			property_id	path		string					true	"Property ID"
//	@Param			lease_id	path		string					true	"Lease ID or `current`"
//	@Param			doc			body		models.DocumentRequest	true	"Document to upload"
//	@Success		201			{object}	models.DocumentResponse	"Document uploaded"
//	@Failure		400			{object}	utils.Error				"Missing fields"
//	@Failure		403			{object}	utils.Error				"Property not yours"
//	@Failure		404			{object}	utils.Error				"No active lease"
//	@Failure		500
//	@Security		Bearer
//	@Router			/owner/properties/{property_id}/leases/{lease_id}/docs/ [post]
//	@Router			/tenant/leases/{lease_id}/docs/ [post]
func UploadDocument(c *gin.Context) {
	var req models.DocumentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendError(c, http.StatusBadRequest, utils.MissingFields, err)
		return
	}

	doc := req.ToDbDocument()
	if doc == nil {
		utils.SendError(c, http.StatusBadRequest, utils.BadBase64String, nil)
		return
	}

	lease, _ := c.MustGet("lease").(db.LeaseModel)
	res := database.CreateDocument(*doc, lease.ID)
	c.JSON(http.StatusCreated, models.DbDocumentToResponse(res))
}

// GetAllDocumentsByLease godoc
//
//	@Summary		Get property documents
//	@Description	Get all documents of a lease related to a property
//	@Tags			document
//	@Accept			json
//	@Produce		json
//	@Param			property_id	path		string					true	"Property ID"
//	@Param			lease_id	path		string					true	"Lease ID or `current`"
//	@Success		200			{array}		models.DocumentResponse	"List of documents"
//	@Failure		403			{object}	utils.Error				"Property not yours"
//	@Failure		404			{object}	utils.Error				"No active lease"
//	@Failure		500
//	@Security		Bearer
//	@Router			/owner/properties/{property_id}/leases/{lease_id}/docs/ [get]
//	@Router			/tenant/leases/{lease_id}/docs/ [get]
func GetAllDocumentsByLease(c *gin.Context) {
	lease, _ := c.MustGet("lease").(db.LeaseModel)
	documents := database.GetDocumentsByLease(lease.ID)
	c.JSON(http.StatusOK, utils.Map(documents, models.DbDocumentToResponse))
}

// GetDocument godoc
//
//	@Summary		Get document by ID
//	@Description	Get a document by its ID
//	@Tags			document
//	@Accept			json
//	@Produce		json
//	@Param			property_id	path		string					true	"Property ID"
//	@Param			lease_id	path		string					true	"Lease ID or `current`"
//	@Param			doc_id		path		string					true	"Document ID"
//	@Success		200			{object}	models.DocumentResponse	"Document"
//	@Failure		403			{object}	utils.Error				"Property not yours"
//	@Failure		404			{object}	utils.Error				"Document not found"
//	@Failure		500
//	@Security		Bearer
//	@Router			/owner/properties/{property_id}/leases/{lease_id}/docs/{doc_id}/ [get]
//	@Router			/tenant/leases/{lease_id}/docs/{doc_id}/ [get]
func GetDocument(c *gin.Context) {
	doc, _ := c.MustGet("document").(db.DocumentModel)
	c.JSON(http.StatusOK, models.DbDocumentToResponse(doc))
}
