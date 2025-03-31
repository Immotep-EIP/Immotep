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
)

func TestCheckActiveLease(t *testing.T) {
	gin.SetMode(gin.TestMode)

	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	lease := db.LeaseModel{
		InnerLease: db.InnerLease{
			ID:         "1",
			PropertyID: "1",
			Active:     true,
		},
	}
	mock.Lease.Expect(
		client.Client.Lease.FindMany(
			db.Lease.PropertyID.Equals("1"),
			db.Lease.Active.Equals(true),
		).With(
			db.Lease.Tenant.Fetch(),
			db.Lease.Property.Fetch().With(db.Property.Owner.Fetch()),
		),
	).ReturnsMany([]db.LeaseModel{lease})

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{gin.Param{Key: "propertyId", Value: "1"}}

	middlewares.CheckActiveLeaseByProperty("propertyId")(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCheckActiveLease_NotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)

	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	mock.Lease.Expect(
		client.Client.Lease.FindMany(
			db.Lease.PropertyID.Equals("1"),
			db.Lease.Active.Equals(true),
		).With(
			db.Lease.Tenant.Fetch(),
			db.Lease.Property.Fetch().With(db.Property.Owner.Fetch()),
		),
	).ReturnsMany([]db.LeaseModel{})

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{gin.Param{Key: "propertyId", Value: "1"}}

	middlewares.CheckActiveLeaseByProperty("propertyId")(c)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestCheckActiveLease_MultipleActiveLeases(t *testing.T) {
	gin.SetMode(gin.TestMode)

	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	lease1 := db.LeaseModel{
		InnerLease: db.InnerLease{
			ID:         "1",
			PropertyID: "1",
			Active:     true,
		},
	}
	lease2 := db.LeaseModel{
		InnerLease: db.InnerLease{
			ID:         "2",
			PropertyID: "1",
			Active:     true,
		},
	}
	mock.Lease.Expect(
		client.Client.Lease.FindMany(
			db.Lease.PropertyID.Equals("1"),
			db.Lease.Active.Equals(true),
		).With(
			db.Lease.Tenant.Fetch(),
			db.Lease.Property.Fetch().With(db.Property.Owner.Fetch()),
		),
	).ReturnsMany([]db.LeaseModel{lease1, lease2})

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{gin.Param{Key: "propertyId", Value: "1"}}

	assert.Panics(t, func() {
		middlewares.CheckActiveLeaseByProperty("propertyId")(c)
	})
}

func TestCheckLeaseInvite(t *testing.T) {
	gin.SetMode(gin.TestMode)

	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	leaseInvite := db.LeaseInviteModel{
		InnerLeaseInvite: db.InnerLeaseInvite{
			ID:         "1",
			PropertyID: "1",
		},
	}
	mock.LeaseInvite.Expect(
		client.Client.LeaseInvite.FindUnique(db.LeaseInvite.PropertyID.Equals("1")),
	).Returns(leaseInvite)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{gin.Param{Key: "propertyId", Value: "1"}}

	middlewares.CheckLeaseInvite("propertyId")(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCheckLeaseInvite_NotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)

	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	mock.LeaseInvite.Expect(
		client.Client.LeaseInvite.FindUnique(db.LeaseInvite.PropertyID.Equals("1")),
	).Errors(db.ErrNotFound)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{gin.Param{Key: "propertyId", Value: "1"}}

	middlewares.CheckLeaseInvite("propertyId")(c)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestCheckLeaseOwnership_CurrentLease(t *testing.T) {
	gin.SetMode(gin.TestMode)

	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	lease := db.LeaseModel{
		InnerLease: db.InnerLease{
			ID:         "1",
			PropertyID: "1",
			Active:     true,
		},
	}
	mock.Lease.Expect(
		client.Client.Lease.FindMany(
			db.Lease.PropertyID.Equals("1"),
			db.Lease.Active.Equals(true),
		).With(
			db.Lease.Tenant.Fetch(),
			db.Lease.Property.Fetch().With(db.Property.Owner.Fetch()),
		),
	).ReturnsMany([]db.LeaseModel{lease})

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{
		gin.Param{Key: "propertyId", Value: "1"},
		gin.Param{Key: "leaseId", Value: "current"},
	}

	middlewares.CheckLeasePropertyOwnership("propertyId", "leaseId")(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCheckLeaseOwnership_LeaseByID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	lease := db.LeaseModel{
		InnerLease: db.InnerLease{
			ID:         "1",
			PropertyID: "1",
		},
	}
	mock.Lease.Expect(
		client.Client.Lease.FindUnique(
			db.Lease.ID.Equals("1"),
		).With(
			db.Lease.Tenant.Fetch(),
			db.Lease.Property.Fetch().With(db.Property.Owner.Fetch()),
		),
	).Returns(lease)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{
		gin.Param{Key: "propertyId", Value: "1"},
		gin.Param{Key: "leaseId", Value: "1"},
	}

	middlewares.CheckLeasePropertyOwnership("propertyId", "leaseId")(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCheckLeaseOwnership_LeaseNotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)

	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	mock.Lease.Expect(
		client.Client.Lease.FindUnique(
			db.Lease.ID.Equals("1"),
		).With(
			db.Lease.Tenant.Fetch(),
			db.Lease.Property.Fetch().With(db.Property.Owner.Fetch()),
		),
	).Errors(db.ErrNotFound)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{
		gin.Param{Key: "propertyId", Value: "1"},
		gin.Param{Key: "leaseId", Value: "1"},
	}

	middlewares.CheckLeasePropertyOwnership("propertyId", "leaseId")(c)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestCheckLeaseOwnership_PropertyMismatch(t *testing.T) {
	gin.SetMode(gin.TestMode)

	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	lease := db.LeaseModel{
		InnerLease: db.InnerLease{
			ID:         "1",
			PropertyID: "2",
		},
	}
	mock.Lease.Expect(
		client.Client.Lease.FindUnique(
			db.Lease.ID.Equals("1"),
		).With(
			db.Lease.Tenant.Fetch(),
			db.Lease.Property.Fetch().With(db.Property.Owner.Fetch()),
		),
	).Returns(lease)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{
		gin.Param{Key: "propertyId", Value: "1"},
		gin.Param{Key: "leaseId", Value: "1"},
	}

	middlewares.CheckLeasePropertyOwnership("propertyId", "leaseId")(c)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestCheckDocumentOwnership(t *testing.T) {
	gin.SetMode(gin.TestMode)

	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	lease := db.LeaseModel{
		InnerLease: db.InnerLease{
			ID: "1",
		},
	}
	doc := db.DocumentModel{
		InnerDocument: db.InnerDocument{
			ID:      "doc1",
			LeaseID: "1",
		},
	}

	mock.Document.Expect(
		client.Client.Document.FindUnique(
			db.Document.ID.Equals("doc1"),
		),
	).Returns(doc)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("lease", lease)
	c.Params = gin.Params{gin.Param{Key: "docId", Value: "doc1"}}

	middlewares.CheckDocumentLeaseOwnership("docId")(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCheckDocumentOwnership_DocumentNotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)

	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	lease := db.LeaseModel{
		InnerLease: db.InnerLease{
			ID:         "1",
			PropertyID: "1",
		},
	}

	mock.Document.Expect(
		client.Client.Document.FindUnique(
			db.Document.ID.Equals("doc1"),
		),
	).Errors(db.ErrNotFound)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("lease", lease)
	c.Params = gin.Params{gin.Param{Key: "docId", Value: "doc1"}}

	middlewares.CheckDocumentLeaseOwnership("docId")(c)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestCheckDocumentOwnership_LeaseMismatch(t *testing.T) {
	gin.SetMode(gin.TestMode)

	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	lease := db.LeaseModel{
		InnerLease: db.InnerLease{
			ID:         "1",
			PropertyID: "1",
		},
	}
	doc := db.DocumentModel{
		InnerDocument: db.InnerDocument{
			ID:      "doc1",
			LeaseID: "2",
		},
	}

	mock.Document.Expect(
		client.Client.Document.FindUnique(
			db.Document.ID.Equals("doc1"),
		),
	).Returns(doc)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("lease", lease)
	c.Params = gin.Params{gin.Param{Key: "docId", Value: "doc1"}}

	middlewares.CheckDocumentLeaseOwnership("docId")(c)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestCheckActiveLeaseByTenant(t *testing.T) {
	gin.SetMode(gin.TestMode)

	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	lease := db.LeaseModel{
		InnerLease: db.InnerLease{
			ID:         "1",
			PropertyID: "1",
			TenantID:   "1",
			Active:     true,
		},
	}
	mock.Lease.Expect(
		client.Client.Lease.FindMany(
			db.Lease.TenantID.Equals("1"),
			db.Lease.Active.Equals(true),
		).With(
			db.Lease.Tenant.Fetch(),
			db.Lease.Property.Fetch().With(db.Property.Owner.Fetch()),
		),
	).ReturnsMany([]db.LeaseModel{lease})

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("oauth.claims", map[string]string{"id": "1"})

	middlewares.CheckActiveLeaseByTenant()(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCheckActiveLeaseByTenant_NotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)

	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	mock.Lease.Expect(
		client.Client.Lease.FindMany(
			db.Lease.TenantID.Equals("1"),
			db.Lease.Active.Equals(true),
		).With(
			db.Lease.Tenant.Fetch(),
			db.Lease.Property.Fetch().With(db.Property.Owner.Fetch()),
		),
	).ReturnsMany([]db.LeaseModel{})

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("oauth.claims", map[string]string{"id": "1"})

	middlewares.CheckActiveLeaseByTenant()(c)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestCheckActiveLeaseByTenant_MultipleActiveLeases(t *testing.T) {
	gin.SetMode(gin.TestMode)

	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	lease1 := db.LeaseModel{
		InnerLease: db.InnerLease{
			ID:         "1",
			PropertyID: "1",
			TenantID:   "1",
			Active:     true,
		},
	}
	lease2 := db.LeaseModel{
		InnerLease: db.InnerLease{
			ID:         "2",
			PropertyID: "2",
			TenantID:   "1",
			Active:     true,
		},
	}
	mock.Lease.Expect(
		client.Client.Lease.FindMany(
			db.Lease.TenantID.Equals("1"),
			db.Lease.Active.Equals(true),
		).With(
			db.Lease.Tenant.Fetch(),
			db.Lease.Property.Fetch().With(db.Property.Owner.Fetch()),
		),
	).ReturnsMany([]db.LeaseModel{lease1, lease2})

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("oauth.claims", map[string]string{"id": "1"})

	assert.Panics(t, func() {
		middlewares.CheckActiveLeaseByTenant()(c)
	})
}

func TestCheckLeaseTenantOwnership_CurrentLease(t *testing.T) {
	gin.SetMode(gin.TestMode)

	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	lease := db.LeaseModel{
		InnerLease: db.InnerLease{
			ID:         "1",
			PropertyID: "1",
			TenantID:   "1",
			Active:     true,
		},
	}
	mock.Lease.Expect(
		client.Client.Lease.FindMany(
			db.Lease.TenantID.Equals("1"),
			db.Lease.Active.Equals(true),
		).With(
			db.Lease.Tenant.Fetch(),
			db.Lease.Property.Fetch().With(db.Property.Owner.Fetch()),
		),
	).ReturnsMany([]db.LeaseModel{lease})

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("oauth.claims", map[string]string{"id": "1"})
	c.Params = gin.Params{gin.Param{Key: "leaseId", Value: "current"}}

	middlewares.CheckLeaseTenantOwnership("leaseId")(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCheckLeaseTenantOwnership_LeaseByID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	lease := db.LeaseModel{
		InnerLease: db.InnerLease{
			ID:       "1",
			TenantID: "1",
		},
	}
	mock.Lease.Expect(
		client.Client.Lease.FindUnique(
			db.Lease.ID.Equals("1"),
		).With(
			db.Lease.Tenant.Fetch(),
			db.Lease.Property.Fetch().With(db.Property.Owner.Fetch()),
		),
	).Returns(lease)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("oauth.claims", map[string]string{"id": "1"})
	c.Params = gin.Params{gin.Param{Key: "leaseId", Value: "1"}}

	middlewares.CheckLeaseTenantOwnership("leaseId")(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCheckLeaseTenantOwnership_LeaseNotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)

	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	mock.Lease.Expect(
		client.Client.Lease.FindUnique(
			db.Lease.ID.Equals("1"),
		).With(
			db.Lease.Tenant.Fetch(),
			db.Lease.Property.Fetch().With(db.Property.Owner.Fetch()),
		),
	).Errors(db.ErrNotFound)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("oauth.claims", map[string]string{"id": "1"})
	c.Params = gin.Params{gin.Param{Key: "leaseId", Value: "1"}}

	middlewares.CheckLeaseTenantOwnership("leaseId")(c)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestCheckLeaseTenantOwnership_TenantMismatch(t *testing.T) {
	gin.SetMode(gin.TestMode)

	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	lease := db.LeaseModel{
		InnerLease: db.InnerLease{
			ID:       "1",
			TenantID: "2",
		},
	}
	mock.Lease.Expect(
		client.Client.Lease.FindUnique(
			db.Lease.ID.Equals("1"),
		).With(
			db.Lease.Tenant.Fetch(),
			db.Lease.Property.Fetch().With(db.Property.Owner.Fetch()),
		),
	).Returns(lease)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("oauth.claims", map[string]string{"id": "1"})
	c.Params = gin.Params{gin.Param{Key: "leaseId", Value: "1"}}

	middlewares.CheckLeaseTenantOwnership("leaseId")(c)
	assert.Equal(t, http.StatusNotFound, w.Code)
}
