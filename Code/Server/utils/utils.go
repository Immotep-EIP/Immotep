package utils

import (
	"github.com/gin-gonic/gin"
)

func GetClaims(c *gin.Context) map[string]string {
	return c.GetStringMapString("oauth.claims")
}

func Map[T, V any](ts []T, fn func(T) V) []V {
    result := make([]V, len(ts))
    for i, t := range ts {
        result[i] = fn(t)
    }
    return result
}
