package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"immotep/backend/services/database"
	"immotep/backend/utils"
)

func CheckPropertyOwnership(propertyIdUrlParam string) gin.HandlerFunc {
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

		c.Next()
	}
}

func CheckRoomOwnership(propertyIdUrlParam string, roomIdUrlParam string) gin.HandlerFunc {
	return func(c *gin.Context) {
		room := database.GetRoomByID(c.Param(roomIdUrlParam))
		if room == nil || room.PropertyID != c.Param(propertyIdUrlParam) {
			utils.AbortSendError(c, http.StatusNotFound, utils.RoomNotFound, nil)
			return
		}

		c.Next()
	}
}

func CheckFurnitureOwnership(roomIdUrlParam string, furnitureIdUrlParam string) gin.HandlerFunc {
	return func(c *gin.Context) {
		furniture := database.GetFurnitureByID(c.Param(furnitureIdUrlParam))
		if furniture == nil || furniture.RoomID != c.Param(roomIdUrlParam) {
			utils.AbortSendError(c, http.StatusNotFound, utils.FurnitureNotFound, nil)
			return
		}

		c.Next()
	}
}

func CheckInventoryReportOwnership(propertyIdUrlParam string, reportIdUrlParam string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Param(reportIdUrlParam) == "latest" {
			c.Next()
			return
		}
		invrep := database.GetInvReportByID(c.Param(reportIdUrlParam))
		if invrep == nil || invrep.PropertyID != c.Param(propertyIdUrlParam) {
			utils.AbortSendError(c, http.StatusNotFound, utils.InventoryReportNotFound, nil)
			return
		}

		c.Next()
	}
}

func CheckActiveContract(propertyIdUrlParam string) gin.HandlerFunc {
	return func(c *gin.Context) {
		contract := database.GetCurrentActiveContract(c.Param(propertyIdUrlParam))
		if contract == nil {
			utils.AbortSendError(c, http.StatusNotFound, utils.NoActiveContract, nil)
			return
		}

		c.Next()
	}
}

func CheckPendingContract(propertyIdUrlParam string) gin.HandlerFunc {
	return func(c *gin.Context) {
		contract := database.GetCurrentPendingContract(c.Param(propertyIdUrlParam))
		if contract == nil {
			utils.AbortSendError(c, http.StatusNotFound, utils.NoPendingContract, nil)
			return
		}

		c.Next()
	}
}
