package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	propertyservice "immotep/backend/services/property"
	roomservice "immotep/backend/services/room"
	"immotep/backend/utils"
)

func CheckPropertyOwnership(propertyIdUrlParam string) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := utils.GetClaims(c)
		property := propertyservice.GetByID(c.Param(propertyIdUrlParam))
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

func CheckRoomExists(propertyIdUrlParam string, roomIdUrlParam string) gin.HandlerFunc {
	return func(c *gin.Context) {
		room := roomservice.GetByID(c.Param(roomIdUrlParam))
		if room == nil || room.PropertyID != c.Param(propertyIdUrlParam) {
			utils.AbortSendError(c, http.StatusNotFound, utils.RoomNotFound, nil)
			return
		}

		c.Next()
	}
}
