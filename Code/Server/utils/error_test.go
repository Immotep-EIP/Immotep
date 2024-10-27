package utils_test

import (
	"errors"
	"immotep/backend/utils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSendError_NoError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/test", func(c *gin.Context) {
		utils.SendError(c, http.StatusBadRequest, utils.InvalidPassword, nil)
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.JSONEq(t, `{"code":"invalid-password","error":""}`, w.Body.String())
}

func TestSendError_WithError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/test", func(c *gin.Context) {
		utils.SendError(c, http.StatusInternalServerError, utils.CannotFetchUser, errors.New("database error"))
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.JSONEq(t, `{"code":"cannot-fetch-user","error":"database error"}`, w.Body.String())
}
