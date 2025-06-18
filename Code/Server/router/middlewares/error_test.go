package middlewares_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"keyz/backend/router/middlewares"
)

func TestPanicRecovery_WithError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	err := gin.Error{Err: assert.AnError}
	middlewares.PanicRecovery(c, err)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, assert.AnError.Error(), w.Body.String())
}

func TestPanicRecovery_WithString(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	middlewares.PanicRecovery(c, "test error")

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "test error", w.Body.String())
}

func TestPanicRecovery_WithUnknownType(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	middlewares.PanicRecovery(c, 123)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "An unknown error occurred", w.Body.String())
}
