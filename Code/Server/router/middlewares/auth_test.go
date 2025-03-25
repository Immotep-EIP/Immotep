package middlewares_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"immotep/backend/prisma/db"
	"immotep/backend/router/middlewares"
)

func TestCheckClaims(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("oauth.claims", map[string]string{"id": "1"})

	middlewares.CheckClaims()(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCheckClaimsNoClaims(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	middlewares.CheckClaims()(c)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestAuthorizeOwner(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("oauth.claims", map[string]string{"role": string(db.RoleOwner)})

	middlewares.AuthorizeOwner()(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestAuthorizeOwnerNotAnOwner(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("oauth.claims", map[string]string{"role": string(db.RoleTenant)})

	middlewares.AuthorizeOwner()(c)
	assert.Equal(t, http.StatusForbidden, w.Code)
}

func TestAuthorizeTenant(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("oauth.claims", map[string]string{"role": string(db.RoleTenant)})

	middlewares.AuthorizeTenant()(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestAuthorizeTenantNotAnTenant(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("oauth.claims", map[string]string{"role": string(db.RoleOwner)})

	middlewares.AuthorizeTenant()(c)
	assert.Equal(t, http.StatusForbidden, w.Code)
}

func TestMockClaims(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(http.MethodGet, "/", nil)
	c.Request.Header.Set("Oauth.claims.id", "1")
	c.Request.Header.Set("Oauth.claims.role", string(db.RoleOwner))

	middlewares.MockClaims()(c)

	claims, exists := c.Get("oauth.claims")
	assert.True(t, exists)
	assert.Equal(t, map[string]string{"id": "1", "role": string(db.RoleOwner)}, claims)
}

func TestMockClaimsNoHeaders(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(http.MethodGet, "/", nil)

	middlewares.MockClaims()(c)

	claims, exists := c.Get("oauth.claims")
	assert.True(t, exists)
	assert.Equal(t, map[string]string{"id": "", "role": ""}, claims)
}
