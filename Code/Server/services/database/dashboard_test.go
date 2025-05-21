package database_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"immotep/backend/prisma/db"
	"immotep/backend/services"
	"immotep/backend/services/database"
)

func TestGetAllDatasFromProperties(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	m.Property.Expect(database.MockGetAllDatasFromProperties(c)).ReturnsMany([]db.PropertyModel{property})

	result := database.GetAllDatasFromProperties("1")
	assert.Len(t, result, 1)
	assert.Equal(t, property.ID, result[0].ID)
	assert.Equal(t, property.Name, result[0].Name)
	assert.Equal(t, property.OwnerID, result[0].OwnerID)
}

func TestGetAllDatasFromProperties_PanicOnError(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.Property.Expect(database.MockGetAllDatasFromProperties(c)).Errors(errors.New("connection failed"))

	assert.Panics(t, func() {
		database.GetAllDatasFromProperties("1")
	})
}
