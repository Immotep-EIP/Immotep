package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"immotep/backend/prisma/db"
	"immotep/backend/services/database"
	"immotep/backend/utils"
)

func CheckPropertyOwnerOwnership(propertyIdUrlParam string) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := utils.GetClaims(c)
		property := database.GetPropertyByID(c.Param(propertyIdUrlParam))
		if property == nil {
			utils.AbortSendError(c, http.StatusNotFound, utils.PropertyNotFound, nil)
			return
		}
		if property.OwnerID != claims["id"] {
			utils.AbortSendError(c, http.StatusForbidden, utils.PropertyNotYours, nil)
			return
		}

		c.Set("property", *property)
		c.Next()
	}
}

func GetPropertyByLease() gin.HandlerFunc {
	return func(c *gin.Context) {
		lease, _ := c.MustGet("lease").(db.LeaseModel)

		property := database.GetPropertyByID(lease.PropertyID)
		if property == nil {
			utils.AbortSendError(c, http.StatusNotFound, utils.PropertyNotFound, nil)
			return
		}

		c.Set("property", *property)
		c.Next()
	}
}

func CheckRoomPropertyOwnership(propertyIdUrlParam string, roomIdUrlParam string) gin.HandlerFunc {
	return func(c *gin.Context) {
		room := database.GetRoomByID(c.Param(roomIdUrlParam))
		if room == nil || room.PropertyID != c.Param(propertyIdUrlParam) {
			utils.AbortSendError(c, http.StatusNotFound, utils.RoomNotFound, nil)
			return
		}

		c.Set("room", *room)
		c.Next()
	}
}

func CheckFurnitureRoomOwnership(roomIdUrlParam string, furnitureIdUrlParam string) gin.HandlerFunc {
	return func(c *gin.Context) {
		furniture := database.GetFurnitureByID(c.Param(furnitureIdUrlParam))
		if furniture == nil || furniture.RoomID != c.Param(roomIdUrlParam) {
			utils.AbortSendError(c, http.StatusNotFound, utils.FurnitureNotFound, nil)
			return
		}

		c.Set("furniture", *furniture)
		c.Next()
	}
}

func CheckInventoryReportPropertyOwnership(propertyIdUrlParam string, reportIdUrlParam string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var invrep *db.InventoryReportModel
		if c.Param(reportIdUrlParam) == "latest" {
			invrep = database.GetLatestInvReport(c.Param(propertyIdUrlParam))
		} else {
			invrep = database.GetInvReportByID(c.Param(reportIdUrlParam))
		}
		if invrep == nil || invrep.PropertyID != c.Param(propertyIdUrlParam) {
			utils.AbortSendError(c, http.StatusNotFound, utils.InventoryReportNotFound, nil)
			return
		}

		c.Set("invrep", *invrep)
		c.Next()
	}
}
