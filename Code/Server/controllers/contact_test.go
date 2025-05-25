package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"immotep/backend/models"
	"immotep/backend/prisma/db"
	"immotep/backend/router"
	"immotep/backend/services"
	"immotep/backend/services/database"
)

func BuildTestContactMessage() db.ContactMessageModel {
	return db.ContactMessageModel{
		InnerContactMessage: db.InnerContactMessage{
			ID:        "1",
			Firstname: "John",
			Lastname:  "Doe",
			Email:     "john.doe@example.com",
			Subject:   "Test Subject",
			Message:   "This is a test message.",
		},
	}
}

func TestCreateContactMessage(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	cm := BuildTestContactMessage()
	m.ContactMessage.Expect(database.MockCreateContactMessage(c, cm)).Returns(cm)

	reqBody := models.ContactMessageRequest{
		Firstname: cm.Firstname,
		Lastname:  cm.Lastname,
		Email:     cm.Email,
		Subject:   cm.Subject,
		Message:   cm.Message,
	}
	b, err := json.Marshal(reqBody)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/v1/contact/", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	var resp models.IdResponse
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, cm.ID, resp.ID)
}

func TestCreateContactMessage_MissingFields(t *testing.T) {
	reqBody := models.ContactMessageRequest{
		Message: "Test",
	}
	b, err := json.Marshal(reqBody)
	require.NoError(t, err)

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	reqHttp, _ := http.NewRequest(http.MethodPost, "/v1/contact/", bytes.NewReader(b))
	reqHttp.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, reqHttp)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}
