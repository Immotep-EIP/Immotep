package utils_test

import (
	"immotep/backend/utils"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestGetClaims(t *testing.T) {
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(nil)
	expectedClaims := map[string]string{"sub": "1234567890", "name": "John Doe", "admin": "true"}
	c.Set("oauth.claims", expectedClaims)

	claims := utils.GetClaims(c)
	assert.Equal(t, expectedClaims, claims)
}

func TestMap(t *testing.T) {
	input := []int{1, 2, 3, 4}
	expected := []string{"1", "2", "3", "4"}
	result := utils.Map(input, func(i int) string {
		return strconv.Itoa(i)
	})
	assert.Equal(t, expected, result)
}

func TestHashPassword(t *testing.T) {
	password := "mysecretpassword"
	hash, err := utils.HashPassword(password)
	assert.NoError(t, err)

	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	assert.NoError(t, err)
}

func TestCheckPasswordHash(t *testing.T) {
	password := "mysecretpassword"
	hash, err := utils.HashPassword(password)
	assert.NoError(t, err)

	match := utils.CheckPasswordHash(password, hash)
	assert.True(t, match)

	wrongPassword := "wrongpassword"
	match = utils.CheckPasswordHash(wrongPassword, hash)
	assert.False(t, match)
}
