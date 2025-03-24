package controllers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"immotep/backend/models"
	"immotep/backend/prisma/db"
	"immotep/backend/router"
	"immotep/backend/services"
	"immotep/backend/utils"
)

func BuildTestDocument() db.DocumentModel {
	return db.DocumentModel{
		InnerDocument: db.InnerDocument{
			ID:        "1",
			Name:      "Test Document",
			Data:      []byte("Test Data"),
			LeaseID:   "1",
			CreatedAt: time.Now(),
		},
	}
}

func TestGetPropertyDocuments(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals(property.ID)).With(
			db.Property.Damages.Fetch(),
			db.Property.Leases.Fetch().With(db.Lease.Tenant.Fetch()),
			db.Property.LeaseInvite.Fetch(),
		),
	).Returns(property)

	activeLease := BuildTestLease("1")
	mock.Lease.Expect(
		client.Client.Lease.FindMany(
			db.Lease.PropertyID.Equals("1"),
			db.Lease.Active.Equals(true),
		).With(
			db.Lease.Tenant.Fetch(),
		),
	).ReturnsMany([]db.LeaseModel{activeLease})

	documents := []db.DocumentModel{BuildTestDocument()}
	mock.Document.Expect(
		client.Client.Document.FindMany(
			db.Document.LeaseID.Equals(activeLease.ID),
		),
	).ReturnsMany(documents)

	r := router.TestRoutes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/owner/properties/1/leases/current/documents/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	var resp []models.DocumentResponse
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, documents[0].ID, resp[0].ID)
}

func TestGetPropertyDocuments_NotYours(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals(property.ID)).With(
			db.Property.Damages.Fetch(),
			db.Property.Leases.Fetch().With(db.Lease.Tenant.Fetch()),
			db.Property.LeaseInvite.Fetch(),
		),
	).Returns(property)

	r := router.TestRoutes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/owner/properties/1/leases/current/documents/", nil)
	req.Header.Set("Oauth.claims.id", "2")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusForbidden, w.Code)
	var errorResponse utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
	require.NoError(t, err)
	assert.Equal(t, utils.PropertyNotYours, errorResponse.Code)
}

func TestGetPropertyDocuments_NoActiveLease(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals(property.ID)).With(
			db.Property.Damages.Fetch(),
			db.Property.Leases.Fetch().With(db.Lease.Tenant.Fetch()),
			db.Property.LeaseInvite.Fetch(),
		),
	).Returns(property)

	mock.Lease.Expect(
		client.Client.Lease.FindMany(
			db.Lease.PropertyID.Equals("1"),
			db.Lease.Active.Equals(true),
		).With(
			db.Lease.Tenant.Fetch(),
		),
	).Errors(db.ErrNotFound)

	r := router.TestRoutes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/owner/properties/1/leases/current/documents/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)
	var errorResponse utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
	require.NoError(t, err)
	assert.Equal(t, utils.NoActiveLease, errorResponse.Code)
}
