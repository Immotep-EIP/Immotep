package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PanicRecovery(c *gin.Context, err any) {
	switch e := err.(type) {
	case error:
		c.Abort()
		c.String(http.StatusInternalServerError, e.Error())
	case string:
		c.Abort()
		c.String(http.StatusInternalServerError, e)
	default:
		c.Abort()
		c.String(http.StatusInternalServerError, "An unknown error occurred")
	}
}
