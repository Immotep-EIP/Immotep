package middlewares_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"immotep/backend/prisma/db"
	"immotep/backend/router/middlewares"
	"immotep/backend/services"
	"immotep/backend/services/database"
)

func TestCheckLeaseInvite(t *testing.T) {
	gin.SetMode(gin.TestMode)
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	leaseInvite := db.LeaseInviteModel{
		InnerLeaseInvite: db.InnerLeaseInvite{
			ID:         "1",
			PropertyID: "1",
		},
	}
	m.LeaseInvite.Expect(database.MockGetCurrentLeaseInvite(c)).Returns(leaseInvite)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Params = gin.Params{gin.Param{Key: "propertyId", Value: "1"}}

	middlewares.CheckLeaseInvite("propertyId")(ctx)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCheckLeaseInvite_NotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.LeaseInvite.Expect(database.MockGetCurrentLeaseInvite(c)).Errors(db.ErrNotFound)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Params = gin.Params{gin.Param{Key: "propertyId", Value: "1"}}

	middlewares.CheckLeaseInvite("propertyId")(ctx)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestCheckLeaseOwnership_CurrentLease(t *testing.T) {
	gin.SetMode(gin.TestMode)
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	lease := db.LeaseModel{
		InnerLease: db.InnerLease{
			ID:         "1",
			PropertyID: "1",
			Active:     true,
		},
	}
	m.Lease.Expect(database.MockGetCurrentActiveLeaseByProperty(c)).ReturnsMany([]db.LeaseModel{lease})

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Params = gin.Params{
		gin.Param{Key: "propertyId", Value: "1"},
		gin.Param{Key: "leaseId", Value: "current"},
	}

	middlewares.CheckLeasePropertyOwnership("propertyId", "leaseId")(ctx)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCheckLeaseOwnership_LeaseByID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	lease := db.LeaseModel{
		InnerLease: db.InnerLease{
			ID:         "1",
			PropertyID: "1",
		},
	}
	m.Lease.Expect(database.MockGetLeaseByID(c)).Returns(lease)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Params = gin.Params{
		gin.Param{Key: "propertyId", Value: "1"},
		gin.Param{Key: "leaseId", Value: "1"},
	}

	middlewares.CheckLeasePropertyOwnership("propertyId", "leaseId")(ctx)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCheckLeaseOwnership_LeaseNotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.Lease.Expect(database.MockGetLeaseByID(c)).Errors(db.ErrNotFound)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Params = gin.Params{
		gin.Param{Key: "propertyId", Value: "1"},
		gin.Param{Key: "leaseId", Value: "1"},
	}

	middlewares.CheckLeasePropertyOwnership("propertyId", "leaseId")(ctx)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestCheckLeaseOwnership_PropertyMismatch(t *testing.T) {
	gin.SetMode(gin.TestMode)
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	lease := db.LeaseModel{
		InnerLease: db.InnerLease{
			ID:         "1",
			PropertyID: "2",
		},
	}
	m.Lease.Expect(database.MockGetLeaseByID(c)).Returns(lease)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Params = gin.Params{
		gin.Param{Key: "propertyId", Value: "1"},
		gin.Param{Key: "leaseId", Value: "1"},
	}

	middlewares.CheckLeasePropertyOwnership("propertyId", "leaseId")(ctx)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestCheckDocumentOwnership(t *testing.T) {
	gin.SetMode(gin.TestMode)
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	lease := db.LeaseModel{
		InnerLease: db.InnerLease{
			ID: "1",
		},
	}
	doc := db.DocumentModel{
		InnerDocument: db.InnerDocument{
			ID:      "1",
			LeaseID: "1",
		},
	}
	m.Document.Expect(database.MockGetDocumentByID(c)).Returns(doc)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Set("lease", lease)
	ctx.Params = gin.Params{gin.Param{Key: "docId", Value: "1"}}

	middlewares.CheckDocumentLeaseOwnership("docId")(ctx)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCheckDocumentOwnership_DocumentNotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	lease := db.LeaseModel{
		InnerLease: db.InnerLease{
			ID:         "1",
			PropertyID: "1",
		},
	}
	m.Document.Expect(database.MockGetDocumentByID(c)).Errors(db.ErrNotFound)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Set("lease", lease)
	ctx.Params = gin.Params{gin.Param{Key: "docId", Value: "1"}}

	middlewares.CheckDocumentLeaseOwnership("docId")(ctx)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestCheckDocumentOwnership_LeaseMismatch(t *testing.T) {
	gin.SetMode(gin.TestMode)
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	lease := db.LeaseModel{
		InnerLease: db.InnerLease{
			ID:         "1",
			PropertyID: "1",
		},
	}
	doc := db.DocumentModel{
		InnerDocument: db.InnerDocument{
			ID:      "1",
			LeaseID: "2",
		},
	}
	m.Document.Expect(database.MockGetDocumentByID(c)).Returns(doc)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Set("lease", lease)
	ctx.Params = gin.Params{gin.Param{Key: "docId", Value: "1"}}

	middlewares.CheckDocumentLeaseOwnership("docId")(ctx)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestCheckLeaseTenantOwnership_CurrentLease(t *testing.T) {
	gin.SetMode(gin.TestMode)
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	lease := db.LeaseModel{
		InnerLease: db.InnerLease{
			ID:         "1",
			PropertyID: "1",
			TenantID:   "1",
			Active:     true,
		},
	}
	m.Lease.Expect(database.MockGetCurrentActiveLeaseByTenant(c)).ReturnsMany([]db.LeaseModel{lease})

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Set("oauth.claims", map[string]string{"id": "1"})
	ctx.Params = gin.Params{gin.Param{Key: "leaseId", Value: "current"}}

	middlewares.CheckLeaseTenantOwnership("leaseId")(ctx)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCheckLeaseTenantOwnership_LeaseByID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	lease := db.LeaseModel{
		InnerLease: db.InnerLease{
			ID:       "1",
			TenantID: "1",
		},
	}
	m.Lease.Expect(database.MockGetLeaseByID(c)).Returns(lease)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Set("oauth.claims", map[string]string{"id": "1"})
	ctx.Params = gin.Params{gin.Param{Key: "leaseId", Value: "1"}}

	middlewares.CheckLeaseTenantOwnership("leaseId")(ctx)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCheckLeaseTenantOwnership_LeaseNotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.Lease.Expect(database.MockGetLeaseByID(c)).Errors(db.ErrNotFound)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Set("oauth.claims", map[string]string{"id": "1"})
	ctx.Params = gin.Params{gin.Param{Key: "leaseId", Value: "1"}}

	middlewares.CheckLeaseTenantOwnership("leaseId")(ctx)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestCheckLeaseTenantOwnership_TenantMismatch(t *testing.T) {
	gin.SetMode(gin.TestMode)
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	lease := db.LeaseModel{
		InnerLease: db.InnerLease{
			ID:       "1",
			TenantID: "2",
		},
	}
	m.Lease.Expect(database.MockGetLeaseByID(c)).Returns(lease)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Set("oauth.claims", map[string]string{"id": "1"})
	ctx.Params = gin.Params{gin.Param{Key: "leaseId", Value: "1"}}

	middlewares.CheckLeaseTenantOwnership("leaseId")(ctx)
	assert.Equal(t, http.StatusNotFound, w.Code)
}
