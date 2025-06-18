package router_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"keyz/backend/router"
)

func TestRoutes1_Welcome(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := router.Routes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Welcome to Keyz API", w.Body.String())
}

func TestRoutes2_Welcome(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := router.TestRoutes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Welcome to Keyz API", w.Body.String())
}
