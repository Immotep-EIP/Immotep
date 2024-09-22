package utils

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CheckPasswordHash(password string, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}
