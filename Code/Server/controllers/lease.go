package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"immotep/backend/models"
	"immotep/backend/prisma/db"
	"immotep/backend/services/database"
	"immotep/backend/services/minio"
	"immotep/backend/utils"
)

// GetAllLeasesByProperty godoc
//
//	@Summary		Get leases by property
//	@Description	Get all leases (active and inactive) for a property
//	@Tags			lease
//	@Accept			json
//	@Produce		json
//	@Param			property_id	path		string					true	"Property ID"
//	@Success		200			{array}		models.LeaseResponse	"List of leases"
//	@Failure		404			{object}	utils.Error				"No leases found"
//	@Failure		500
//	@Security		Bearer
//	@Router			/owner/properties/{property_id}/leases/ [get]
func GetAllLeasesByProperty(c *gin.Context) {
	leases := database.GetLeasesByProperty(c.Param("property_id"))
	c.JSON(http.StatusOK, utils.Map(leases, models.DbLeaseToResponse))
}

// GetAllLeasesByTenant godoc
//
//	@Summary		Get leases by tenant
//	@Description	Get all leases (active and inactive) for a tenant
//	@Tags			lease
//	@Accept			json
//	@Produce		json
//	@Param			tenant_id	path		string					true	"Tenant ID"
//	@Success		200			{array}		models.LeaseResponse	"List of leases"
//	@Failure		404			{object}	utils.Error				"No leases found"
//	@Failure		500
//	@Security		Bearer
//	@Router			/tenant/leases/ [get]
func GetAllLeasesByTenant(c *gin.Context) {
	claims := utils.GetClaims(c)
	leases := database.GetLeasesByTenant(claims["id"])
	c.JSON(http.StatusOK, utils.Map(leases, models.DbLeaseToResponse))
}

// GetLease godoc
//
//	@Summary		Get lease by ID
//	@Description	Get lease by ID or current active lease
//	@Tags			lease
//	@Accept			json
//	@Produce		json
//	@Param			property_id	path		string					true	"Property ID"
//	@Param			lease_id	path		string					true	"Lease ID or `current`"
//	@Success		200			{object}	models.LeaseResponse	"Lease"
//	@Failure		404			{object}	utils.Error				"Lease not found"
//	@Failure		500
//	@Security		Bearer
//	@Router			/owner/properties/{property_id}/leases/{lease_id}/ [get]
//	@Router			/tenant/leases/{lease_id}/ [get]
func GetLease(c *gin.Context) {
	lease, _ := c.MustGet("lease").(db.LeaseModel)
	c.JSON(http.StatusOK, models.DbLeaseToResponse(lease))
}

// EndLease godoc
//
//	@Summary		End lease
//	@Description	End active lease for a property
//	@Tags			lease
//	@Accept			json
//	@Produce		json
//	@Param			property_id	path	string	true	"Property ID"
//	@Param			lease_id	path	string	true	"Mandatory: `current`"
//	@Success		204			"Lease ended"
//	@Failure		400			{object}	utils.Error	"Cannot end non current lease"
//	@Failure		403			{object}	utils.Error	"Property is not yours"
//	@Failure		404			{object}	utils.Error	"No active lease"
//	@Failure		500
//	@Security		Bearer
//	@Router			/owner/properties/{property_id}/leases/{lease_id}/end/ [put]
func EndLease(c *gin.Context) {
	leaseId := c.Param("lease_id")
	if leaseId != "current" {
		utils.SendError(c, http.StatusBadRequest, utils.CannotEndNonCurrentLease, nil)
		return
	}

	currentActive, _ := c.MustGet("lease").(db.LeaseModel)
	_, ok := currentActive.EndDate()
	now := time.Now().Truncate(time.Minute)
	database.EndLease(currentActive.ID, utils.Ternary(ok, nil, &now))
	c.Status(http.StatusNoContent)
}

// UploadLeaseDocument godoc
//
//	@Summary		Upload document
//	@Description	Upload a document to a lease
//	@Tags			document
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			property_id	path		string				true	"Property ID"
//	@Param			lease_id	path		string				true	"Lease ID or `current`"
//	@Param			document	formData	file				true	"Document file"
//	@Success		201			{object}	models.IdResponse	"Updated lease ID"
//	@Failure		400			{object}	utils.Error			"Missing fields"
//	@Failure		403			{object}	utils.Error			"Property not yours"
//	@Failure		404			{object}	utils.Error			"No active lease"
//	@Failure		500
//	@Security		Bearer
//	@Router			/owner/properties/{property_id}/leases/{lease_id}/docs/ [post]
//	@Router			/tenant/leases/{lease_id}/docs/ [post]
func UploadLeaseDocument(c *gin.Context) {
	lease, _ := c.MustGet("lease").(db.LeaseModel)
	file, err := c.FormFile("document")
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, utils.MissingFile, err)
		return
	}

	fileInfo := minio.UploadLeaseDocument(lease.ID, file)
	newLease := database.AddDocumentToLease(lease, fileInfo.Key)
	c.JSON(http.StatusOK, models.IdResponse{ID: newLease.ID})
}

// GetAllDocumentsByLease godoc
//
//	@Summary		Get lease documents
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
	docs := minio.GetDocuments(lease.Documents)
	c.JSON(http.StatusOK, docs)
}

// DeleteLeaseDocument godoc
//
//	@Summary		Delete document
//	@Description	Delete a document by its name
//	@Tags			document
//	@Accept			json
//	@Produce		json
//	@Param			property_id	path		string					true	"Property ID"
//	@Param			lease_id	path		string					true	"Lease ID or `current`"
//	@Param			doc_name	path		string					true	"Document name"
//	@Success		204			{object}	models.DocumentResponse	"Document deleted"
//	@Failure		403			{object}	utils.Error				"Property not yours"
//	@Failure		404			{object}	utils.Error				"Document not found"
//	@Failure		500
//	@Security		Bearer
//	@Router			/owner/properties/{property_id}/leases/{lease_id}/docs/{doc_name}/ [delete]
func DeleteLeaseDocument(c *gin.Context) {
	lease, _ := c.MustGet("lease").(db.LeaseModel)
	if minio.DeleteLeaseDocument(lease.ID, c.Param("doc_name")) {
		c.Status(http.StatusNoContent)
	} else {
		utils.SendError(c, http.StatusNotFound, utils.DocumentNotFound, nil)
	}
}
