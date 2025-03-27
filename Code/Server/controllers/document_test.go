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
			ID:         "1",
			Name:       "Test Document",
			Data:       []byte("Test Data"),
			ContractID: "1",
			CreatedAt:  time.Now(),
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
			db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
			db.Property.PendingContract.Fetch(),
		),
	).Returns(property)

	activeContract := BuildTestContract()
	mock.Contract.Expect(
		client.Client.Contract.FindMany(
			db.Contract.PropertyID.Equals("1"),
			db.Contract.Active.Equals(true),
		),
	).ReturnsMany([]db.ContractModel{activeContract})

	documents := []db.DocumentModel{BuildTestDocument()}
	mock.Document.Expect(
		client.Client.Document.FindMany(
			db.Document.ContractID.Equals(activeContract.ID),
		),
	).ReturnsMany(documents)

	r := router.TestRoutes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/owner/properties/1/documents/", nil)
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
			db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
			db.Property.PendingContract.Fetch(),
		),
	).Returns(property)

	r := router.TestRoutes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/owner/properties/1/documents/", nil)
	req.Header.Set("Oauth.claims.id", "2")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusForbidden, w.Code)
	var errorResponse utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
	require.NoError(t, err)
	assert.Equal(t, utils.PropertyNotYours, errorResponse.Code)
}

func TestGetPropertyDocuments_NoActiveContract(t *testing.T) {
	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals(property.ID)).With(
			db.Property.Damages.Fetch(),
			db.Property.Contracts.Fetch().With(db.Contract.Tenant.Fetch()),
			db.Property.PendingContract.Fetch(),
		),
	).Returns(property)

	mock.Contract.Expect(
		client.Client.Contract.FindMany(
			db.Contract.PropertyID.Equals("1"),
			db.Contract.Active.Equals(true),
		),
	).Errors(db.ErrNotFound)

	r := router.TestRoutes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/owner/properties/1/documents/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)
	var errorResponse utils.Error
	err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
	require.NoError(t, err)
	assert.Equal(t, utils.NoActiveContract, errorResponse.Code)
}
