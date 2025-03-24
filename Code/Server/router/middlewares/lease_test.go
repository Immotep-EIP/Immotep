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
		),
	).ReturnsMany([]db.LeaseModel{lease})

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{gin.Param{Key: "propertyId", Value: "1"}}

	middlewares.CheckActiveLease("propertyId")(c)
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
		),
	).ReturnsMany([]db.LeaseModel{})

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{gin.Param{Key: "propertyId", Value: "1"}}

	middlewares.CheckActiveLease("propertyId")(c)
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
		),
	).ReturnsMany([]db.LeaseModel{lease1, lease2})

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{gin.Param{Key: "propertyId", Value: "1"}}

	assert.Panics(t, func() {
		middlewares.CheckActiveLease("propertyId")(c)
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
		),
	).ReturnsMany([]db.LeaseModel{lease})

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{
		gin.Param{Key: "propertyId", Value: "1"},
		gin.Param{Key: "leaseId", Value: "current"},
	}

	middlewares.CheckLeaseOwnership("propertyId", "leaseId")(c)
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
		),
	).Returns(lease)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{
		gin.Param{Key: "propertyId", Value: "1"},
		gin.Param{Key: "leaseId", Value: "1"},
	}

	middlewares.CheckLeaseOwnership("propertyId", "leaseId")(c)
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
		),
	).Errors(db.ErrNotFound)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{
		gin.Param{Key: "propertyId", Value: "1"},
		gin.Param{Key: "leaseId", Value: "1"},
	}

	middlewares.CheckLeaseOwnership("propertyId", "leaseId")(c)
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
		),
	).Returns(lease)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{
		gin.Param{Key: "propertyId", Value: "1"},
		gin.Param{Key: "leaseId", Value: "1"},
	}

	middlewares.CheckLeaseOwnership("propertyId", "leaseId")(c)
	assert.Equal(t, http.StatusNotFound, w.Code)
}
