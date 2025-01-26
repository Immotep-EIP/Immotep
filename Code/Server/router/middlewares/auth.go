package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"immotep/backend/prisma/db"
	"immotep/backend/utils"
)

func MockClaims() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("oauth.claims", map[string]string{
			"id":   c.GetHeader("Oauth.claims.id"),
			"role": c.GetHeader("Oauth.claims.role"),
		})
		c.Next()
	}
}

func CheckClaims() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := utils.GetClaims(c)
		if claims == nil {
			utils.AbortSendError(c, http.StatusUnauthorized, utils.NoClaims, nil)
			return
		}

		c.Next()
	}
}

func AuthorizeOwner() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := utils.GetClaims(c)
		if claims["role"] != string(db.RoleOwner) {
			utils.AbortSendError(c, http.StatusForbidden, utils.NotAnOwner, nil)
			return
		}

		c.Next()
	}
}

func AuthorizeTenant() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := utils.GetClaims(c)
		if claims["role"] != string(db.RoleTenant) {
			utils.AbortSendError(c, http.StatusForbidden, utils.NotATenant, nil)
			return
		}

		c.Next()
	}
}
