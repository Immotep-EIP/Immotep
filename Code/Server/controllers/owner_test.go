package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/steebchen/prisma-client-go/engine/protocol"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"immotep/backend/controllers"
	"immotep/backend/database"
	"immotep/backend/models"
	"immotep/backend/prisma/db"
	"immotep/backend/utils"
)

func BuildTestProperty(id string) db.PropertyModel {
	return db.PropertyModel{
		InnerProperty: db.InnerProperty{
			ID:         id,
			Name:       "Test",
			Address:    "Test",
			City:       "Test",
			PostalCode: "Test",
			Country:    "Test",
			CreatedAt:  time.Now(),
			OwnerID:    "1",
		},
	}
}

func BuildTestPendingContract() db.PendingContractModel {
	end := time.Now().Add(time.Hour)
	return db.PendingContractModel{
		InnerPendingContract: db.InnerPendingContract{
			ID:          "1",
			TenantEmail: "test.test@example.com",
			StartDate:   time.Now(),
			EndDate:     &end,
			PropertyID:  "1",
			CreatedAt:   time.Now(),
		},
	}
}

func TestInviteTenant(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals(property.ID)),
	).Returns(property)

	pendingContract := BuildTestPendingContract()
	mock.PendingContract.Expect(
		client.Client.PendingContract.CreateOne(
			db.PendingContract.TenantEmail.Set(pendingContract.TenantEmail),
			db.PendingContract.StartDate.Set(pendingContract.StartDate),
			db.PendingContract.Property.Link(db.Property.ID.Equals(property.ID)),
			db.PendingContract.EndDate.SetIfPresent(pendingContract.InnerPendingContract.EndDate),
		),
	).Returns(pendingContract)

	req := models.InviteRequest{
		TenantEmail: pendingContract.TenantEmail,
		StartDate:   pendingContract.StartDate,
		EndDate:     pendingContract.InnerPendingContract.EndDate,
	}
	b, err := json.Marshal(req)
	require.NoError(t, err)

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(b))
	c.Request.Header.Set("Content-Type", "application/json")
	c.Set("oauth.claims", map[string]string{"id": "1"})
	c.Params = gin.Params{{Key: "propertyId", Value: property.ID}}

	controllers.InviteTenant(c)
	assert.Equal(t, http.StatusOK, w.Code)
	var resp models.InviteResponse
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.JSONEq(t, resp.ID, pendingContract.ID)
}

func TestInviteTenantMissingField(t *testing.T) {
	property := BuildTestProperty("1")
	pendingContract := BuildTestPendingContract()
	req := models.InviteRequest{
		StartDate: pendingContract.StartDate,
		EndDate:   pendingContract.InnerPendingContract.EndDate,
	}
	b, err := json.Marshal(req)
	require.NoError(t, err)

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(b))
	c.Request.Header.Set("Content-Type", "application/json")
	c.Set("oauth.claims", map[string]string{"id": "1"})
	c.Params = gin.Params{{Key: "propertyId", Value: property.ID}}

	controllers.InviteTenant(c)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	var resp utils.Error
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.MissingFields, resp.Code)
}

func TestInviteTenantPropertyNotFound(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals("wrong")),
	).Errors(db.ErrNotFound)

	pendingContract := BuildTestPendingContract()

	req := models.InviteRequest{
		TenantEmail: pendingContract.TenantEmail,
		StartDate:   pendingContract.StartDate,
		EndDate:     pendingContract.InnerPendingContract.EndDate,
	}
	b, err := json.Marshal(req)
	require.NoError(t, err)

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(b))
	c.Request.Header.Set("Content-Type", "application/json")
	c.Set("oauth.claims", map[string]string{"id": "1"})
	c.Params = gin.Params{{Key: "propertyId", Value: "wrong"}}

	controllers.InviteTenant(c)
	assert.Equal(t, http.StatusNotFound, w.Code)
	var resp utils.Error
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.PropertyNotFound, resp.Code)
}

func TestInviteTenantPropertyNotYours(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals(property.ID)),
	).Returns(property)

	pendingContract := BuildTestPendingContract()

	req := models.InviteRequest{
		TenantEmail: pendingContract.TenantEmail,
		StartDate:   pendingContract.StartDate,
		EndDate:     pendingContract.InnerPendingContract.EndDate,
	}
	b, err := json.Marshal(req)
	require.NoError(t, err)

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(b))
	c.Request.Header.Set("Content-Type", "application/json")
	c.Set("oauth.claims", map[string]string{"id": "wrong"})
	c.Params = gin.Params{{Key: "propertyId", Value: property.ID}}

	controllers.InviteTenant(c)
	assert.Equal(t, http.StatusForbidden, w.Code)
	var resp utils.Error
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.PropertyNotYours, resp.Code)
}

func TestInviteTenantAlreadyExists(t *testing.T) {
	client, mock, ensure := database.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals(property.ID)),
	).Returns(property)

	pendingContract := BuildTestPendingContract()
	mock.PendingContract.Expect(
		client.Client.PendingContract.CreateOne(
			db.PendingContract.TenantEmail.Set(pendingContract.TenantEmail),
			db.PendingContract.StartDate.Set(pendingContract.StartDate),
			db.PendingContract.Property.Link(db.Property.ID.Equals(property.ID)),
			db.PendingContract.EndDate.SetIfPresent(pendingContract.InnerPendingContract.EndDate),
		),
	).Errors(&protocol.UserFacingError{
		IsPanic:   false,
		ErrorCode: "P2002", // https://www.prisma.io/docs/orm/reference/error-reference
		Meta: protocol.Meta{
			Target: []any{"tenant_email"},
		},
		Message: "Unique constraint failed",
	})

	req := models.InviteRequest{
		TenantEmail: pendingContract.TenantEmail,
		StartDate:   pendingContract.StartDate,
		EndDate:     pendingContract.InnerPendingContract.EndDate,
	}
	b, err := json.Marshal(req)
	require.NoError(t, err)

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(b))
	c.Request.Header.Set("Content-Type", "application/json")
	c.Set("oauth.claims", map[string]string{"id": "1"})
	c.Params = gin.Params{{Key: "propertyId", Value: property.ID}}

	controllers.InviteTenant(c)
	assert.Equal(t, http.StatusConflict, w.Code)
	var resp utils.Error
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, utils.InviteAlreadyExists, resp.Code)
}
