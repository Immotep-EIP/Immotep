package router_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"immotep/backend/router"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRoutes_Welcome(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := router.Routes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Welcome to Immotep API", w.Body.String())
}
