package utils

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

const Strue = "true"
const Sfalse = "false"

func GetClaims(c *gin.Context) map[string]string {
	return c.GetStringMapString("oauth.claims")
}

func Map[T, V any](slice []T, transform func(T) V) []V {
	res := make([]V, len(slice))
	for i, t := range slice {
		res[i] = transform(t)
	}
	return res
}

func MapIf[T, V any](slice []T, condition func(T) bool, transform func(T) V) []V {
	res := make([]V, 0, len(slice))
	for _, t := range slice {
		if condition(t) {
			res = append(res, transform(t))
		}
	}
	return res
}

func Ternary[T any](condition bool, yes T, no T) T {
	if condition {
		return yes
	}
	return no
}

func CountIf[T any](slice []T, condition func(T) bool) int {
	count := 0
	for _, elem := range slice {
		if condition(elem) {
			count++
		}
	}
	return count
}

func Filter[T any](slice []T, condition func(T) bool) []T {
	res := make([]T, 0, len(slice))
	for _, elem := range slice {
		if condition(elem) {
			res = append(res, elem)
		}
	}
	return res
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Ptr[T any](v T) *T {
	return &v
}
