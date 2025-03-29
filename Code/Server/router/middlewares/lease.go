package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"immotep/backend/prisma/db"
	"immotep/backend/services/database"
	"immotep/backend/utils"
)

func CheckActiveLeaseByProperty(propertyIdUrlParam string) gin.HandlerFunc {
	return func(c *gin.Context) {
		lease := database.GetCurrentActiveLeaseByProperty(c.Param(propertyIdUrlParam))
		if lease == nil {
			utils.AbortSendError(c, http.StatusNotFound, utils.NoActiveLease, nil)
			return
		}

		c.Set("lease", *lease)
		c.Next()
	}
}

func CheckActiveLeaseByTenant() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := utils.GetClaims(c)
		lease := database.GetCurrentActiveLeaseByTenant(claims["id"])
		if lease == nil {
			utils.AbortSendError(c, http.StatusNotFound, utils.NoActiveLease, nil)
			return
		}

		c.Set("lease", *lease)
		c.Next()
	}
}

func CheckLeaseInvite(propertyIdUrlParam string) gin.HandlerFunc {
	return func(c *gin.Context) {
		leaseInvite := database.GetCurrentLeaseInvite(c.Param(propertyIdUrlParam))
		if leaseInvite == nil {
			utils.AbortSendError(c, http.StatusNotFound, utils.NoLeaseInvite, nil)
			return
		}

		c.Next()
	}
}

func CheckLeasePropertyOwnership(propertyIdUrlParam string, leaseIdUrlParam string) gin.HandlerFunc {
	return func(c *gin.Context) {
		leaseId := c.Param(leaseIdUrlParam)
		if leaseId == "current" {
			CheckActiveLeaseByProperty(propertyIdUrlParam)(c)
			return
		}

		lease := database.GetLeaseByID(leaseId)
		if lease == nil || lease.PropertyID != c.Param(propertyIdUrlParam) {
			utils.AbortSendError(c, http.StatusNotFound, utils.LeaseNotFound, nil)
			return
		}

		c.Set("lease", *lease)
		c.Next()
	}
}

func CheckLeaseTenantOwnership(leaseIdUrlParam string) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := utils.GetClaims(c)
		leaseId := c.Param(leaseIdUrlParam)
		if leaseId == "current" {
			CheckActiveLeaseByTenant()(c)
			return
		}

		lease := database.GetLeaseByID(leaseId)
		if lease == nil || lease.TenantID != claims["id"] {
			utils.AbortSendError(c, http.StatusNotFound, utils.LeaseNotFound, nil)
			return
		}

		c.Set("lease", *lease)
		c.Next()
	}
}

func CheckDocumentLeaseOwnership(docIdUrlParam string) gin.HandlerFunc {
	return func(c *gin.Context) {
		lease, _ := c.MustGet("lease").(db.LeaseModel)

		doc := database.GetDocumentByID(c.Param(docIdUrlParam))
		if doc == nil || doc.LeaseID != lease.ID {
			utils.AbortSendError(c, http.StatusNotFound, utils.DocumentNotFound, nil)
			return
		}

		c.Set("document", *doc)
		c.Next()
	}
}
