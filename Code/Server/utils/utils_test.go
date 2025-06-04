package utils_test

import (
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"immotep/backend/utils"
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
	result := utils.Map(input, strconv.Itoa)
	assert.Equal(t, expected, result)
}

func TestHashPassword(t *testing.T) {
	password := "mysecretpassword"
	hash, err := utils.HashPassword(password)
	require.NoError(t, err)

	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	require.NoError(t, err)
}

func TestCheckPasswordHash(t *testing.T) {
	password := "mysecretpassword"
	hash, err := utils.HashPassword(password)
	require.NoError(t, err)

	match := utils.CheckPasswordHash(password, hash)
	assert.True(t, match)

	wrongPassword := "wrongpassword"
	match = utils.CheckPasswordHash(wrongPassword, hash)
	assert.False(t, match)
}

func TestMapIf(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	condition := func(n int) bool { return n%2 == 0 }
	transform := strconv.Itoa
	expected := []string{"2", "4"}

	result := utils.MapIf(input, condition, transform)
	assert.Equal(t, expected, result)
}

func TestTernary(t *testing.T) {
	assert.Equal(t, "yes", utils.Ternary(true, "yes", "no"))
	assert.Equal(t, "no", utils.Ternary(false, "yes", "no"))

	assert.Equal(t, 10, utils.Ternary(true, 10, 20))
	assert.Equal(t, 20, utils.Ternary(false, 10, 20))

	assert.InEpsilon(t, 3.14, utils.Ternary(true, 3.14, 2.71), 0.0)
	assert.InEpsilon(t, 2.71, utils.Ternary(false, 3.14, 2.71), 0.0)
}

func TestCountIf(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6}
	condition := func(n int) bool { return n%2 == 0 }
	expected := 3

	result := utils.CountIf(input, condition)
	assert.Equal(t, expected, result)

	condition = func(n int) bool { return n > 3 }
	expected = 3

	result = utils.CountIf(input, condition)
	assert.Equal(t, expected, result)

	condition = func(n int) bool { return n < 0 }
	expected = 0

	result = utils.CountIf(input, condition)
	assert.Equal(t, expected, result)
}

func TestPtr(t *testing.T) {
	value := 42
	ptr := utils.Ptr(value)
	require.NotNil(t, ptr)
	assert.Equal(t, value, *ptr)

	str := "hello"
	ptrStr := utils.Ptr(str)
	require.NotNil(t, ptrStr)
	assert.Equal(t, str, *ptrStr)

	boolean := true
	ptrBool := utils.Ptr(boolean)
	require.NotNil(t, ptrBool)
	assert.Equal(t, boolean, *ptrBool)
}

func TestFilter_Ints(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6}
	condition := func(n int) bool { return n%2 == 0 }
	expected := []int{2, 4, 6}

	result := utils.Filter(input, condition)
	assert.Equal(t, expected, result)
}

func TestFilter_Strings(t *testing.T) {
	input := []string{"apple", "banana", "cherry", "date"}
	condition := func(s string) bool { return len(s) > 5 }
	expected := []string{"banana", "cherry"}

	result := utils.Filter(input, condition)
	assert.Equal(t, expected, result)
}

func TestFilter_EmptySlice(t *testing.T) {
	input := []int{}
	condition := func(n int) bool { return n > 0 }
	expected := []int{}

	result := utils.Filter(input, condition)
	assert.Equal(t, expected, result)
}

func TestFilter_NoMatch(t *testing.T) {
	input := []int{1, 3, 5, 7}
	condition := func(n int) bool { return n%2 == 0 }
	expected := []int{}

	result := utils.Filter(input, condition)
	assert.Equal(t, expected, result)
}

func TestFilter_AllMatch(t *testing.T) {
	input := []int{2, 4, 6, 8}
	condition := func(n int) bool { return n%2 == 0 }
	expected := []int{2, 4, 6, 8}

	result := utils.Filter(input, condition)
	assert.Equal(t, expected, result)
}
