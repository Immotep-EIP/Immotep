package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"immotep/backend/prisma/db"
	"immotep/backend/services/database"
	"immotep/backend/utils"
)

const CurrentLeaseID = "current"

func CheckLeaseInvite(propertyIdUrlParam string) gin.HandlerFunc {
	return func(c *gin.Context) {
		leaseInvite := database.GetCurrentLeaseInvite(c.Param(propertyIdUrlParam))
		if leaseInvite == nil {
			utils.AbortSendError(c, http.StatusNotFound, utils.NoLeaseInvite, nil)
			return
		}

		c.Set("invite", *leaseInvite)
		c.Next()
	}
}

func CheckLeasePropertyOwnership(propertyIdUrlParam string, leaseIdUrlParam string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var lease *db.LeaseModel
		propertyId := c.Param(propertyIdUrlParam)
		leaseId := c.Param(leaseIdUrlParam)

		if leaseId == CurrentLeaseID {
			lease = database.GetCurrentActiveLeaseByProperty(propertyId)
		} else {
			lease = database.GetLeaseByID(leaseId)
		}

		if lease == nil || lease.PropertyID != propertyId {
			utils.AbortSendError(c, http.StatusNotFound, utils.Ternary(leaseId == CurrentLeaseID, utils.NoActiveLease, utils.LeaseNotFound), nil)
			return
		}

		c.Set("lease", *lease)
		c.Next()
	}
}

func CheckLeaseTenantOwnership(leaseIdUrlParam string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var lease *db.LeaseModel
		claims := utils.GetClaims(c)
		leaseId := c.Param(leaseIdUrlParam)

		if leaseId == CurrentLeaseID {
			lease = database.GetCurrentActiveLeaseByTenant(claims["id"])
		} else {
			lease = database.GetLeaseByID(leaseId)
		}

		if lease == nil || lease.TenantID != claims["id"] {
			utils.AbortSendError(c, http.StatusNotFound, utils.Ternary(leaseId == CurrentLeaseID, utils.NoActiveLease, utils.LeaseNotFound), nil)
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

func CheckDamageLeaseOwnership(damageIdUrlParam string) gin.HandlerFunc {
	return func(c *gin.Context) {
		lease, _ := c.MustGet("lease").(db.LeaseModel)

		damage := database.GetDamageByID(c.Param(damageIdUrlParam))
		if damage == nil || damage.LeaseID != lease.ID {
			utils.AbortSendError(c, http.StatusNotFound, utils.DamageNotFound, nil)
			return
		}

		c.Set("damage", *damage)
		c.Next()
	}
}
