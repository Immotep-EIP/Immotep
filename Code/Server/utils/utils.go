package utils

import (
	"github.com/gin-gonic/gin"
)

func GetClaims(c *gin.Context) map[string]string {
	return c.GetStringMapString("oauth.claims")
}
