package middlewares_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"immotep/backend/prisma/db"
	"immotep/backend/router/middlewares"
	"immotep/backend/services"
	"immotep/backend/services/database"
	"immotep/backend/utils"
)

func BuildTestProperty(id string) db.PropertyModel {
	return db.PropertyModel{
		InnerProperty: db.InnerProperty{
			ID:                  id,
			Name:                "Test",
			Address:             "Test",
			City:                "Test",
			PostalCode:          "Test",
			Country:             "Test",
			AreaSqm:             20.0,
			RentalPricePerMonth: 500,
			DepositPrice:        1000,
			CreatedAt:           time.Now(),
			OwnerID:             "1",
			PictureID:           utils.Ptr("1"),
		},
		RelationsProperty: db.RelationsProperty{
			Leases: []db.LeaseModel{{
				InnerLease: db.InnerLease{
					ID:     "1",
					Active: true,
				},
				RelationsLease: db.RelationsLease{
					Tenant: &db.UserModel{
						InnerUser: db.InnerUser{
							Firstname: "Test",
							Lastname:  "Test",
						},
					},
					Damages: []db.DamageModel{{
						InnerDamage: db.InnerDamage{
							ID:      "1",
							FixedAt: nil,
						}},
					},
				},
			}},
			LeaseInvite: &db.LeaseInviteModel{},
		},
	}
}

func TestCheckPropertyOwnership(t *testing.T) {
	gin.SetMode(gin.TestMode)
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	m.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Params = gin.Params{gin.Param{Key: "propertyId", Value: "1"}}
	ctx.Set("oauth.claims", map[string]string{"id": "1"})

	middlewares.CheckPropertyOwnerOwnership("propertyId")(ctx)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCheckPropertyOwnership_NotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.Property.Expect(database.MockGetPropertyByID(c)).Errors(db.ErrNotFound)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Params = gin.Params{gin.Param{Key: "propertyId", Value: "1"}}
	ctx.Set("oauth.claims", map[string]string{"id": "1"})

	middlewares.CheckPropertyOwnerOwnership("propertyId")(ctx)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestCheckPropertyOwnership_NotYours(t *testing.T) {
	gin.SetMode(gin.TestMode)
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	m.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Params = gin.Params{gin.Param{Key: "propertyId", Value: "1"}}
	ctx.Set("oauth.claims", map[string]string{"id": "2"})

	middlewares.CheckPropertyOwnerOwnership("propertyId")(ctx)
	assert.Equal(t, http.StatusForbidden, w.Code)
}

func TestCheckRoomOwnership(t *testing.T) {
	gin.SetMode(gin.TestMode)
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	room := db.RoomModel{
		InnerRoom: db.InnerRoom{
			ID:         "1",
			PropertyID: "1",
		},
	}
	m.Room.Expect(database.MockGetRoomByID(c)).Returns(room)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Params = gin.Params{
		{Key: "propertyId", Value: "1"},
		{Key: "roomId", Value: "1"},
	}

	middlewares.CheckRoomPropertyOwnership("propertyId", "roomId")(ctx)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCheckRoomOwnership_NotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.Room.Expect(database.MockGetRoomByID(c)).Errors(db.ErrNotFound)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Params = gin.Params{
		{Key: "propertyId", Value: "1"},
		{Key: "roomId", Value: "1"},
	}

	middlewares.CheckRoomPropertyOwnership("propertyId", "roomId")(ctx)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestCheckRoomOwnership_NotYours(t *testing.T) {
	gin.SetMode(gin.TestMode)
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	room := db.RoomModel{
		InnerRoom: db.InnerRoom{
			ID:         "1",
			PropertyID: "2",
		},
	}
	m.Room.Expect(database.MockGetRoomByID(c)).Returns(room)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Params = gin.Params{
		{Key: "propertyId", Value: "1"},
		{Key: "roomId", Value: "1"},
	}

	middlewares.CheckRoomPropertyOwnership("propertyId", "roomId")(ctx)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestCheckFurnitureOwnership(t *testing.T) {
	gin.SetMode(gin.TestMode)
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	furniture := db.FurnitureModel{
		InnerFurniture: db.InnerFurniture{
			ID:     "1",
			RoomID: "1",
		},
	}
	m.Furniture.Expect(database.MockGetFurnitureByID(c)).Returns(furniture)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Params = gin.Params{
		{Key: "roomId", Value: "1"},
		{Key: "furnitureId", Value: "1"},
	}

	middlewares.CheckFurnitureRoomOwnership("roomId", "furnitureId")(ctx)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCheckFurnitureOwnership_NotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.Furniture.Expect(database.MockGetFurnitureByID(c)).Errors(db.ErrNotFound)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Params = gin.Params{
		{Key: "roomId", Value: "1"},
		{Key: "furnitureId", Value: "1"},
	}

	middlewares.CheckFurnitureRoomOwnership("roomId", "furnitureId")(ctx)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestCheckFurnitureOwnership_NotYours(t *testing.T) {
	gin.SetMode(gin.TestMode)
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	furniture := db.FurnitureModel{
		InnerFurniture: db.InnerFurniture{
			ID:     "1",
			RoomID: "2",
		},
	}
	m.Furniture.Expect(database.MockGetFurnitureByID(c)).Returns(furniture)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Params = gin.Params{
		{Key: "roomId", Value: "1"},
		{Key: "furnitureId", Value: "1"},
	}

	middlewares.CheckFurnitureRoomOwnership("roomId", "furnitureId")(ctx)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestCheckInventoryReportOwnership(t *testing.T) {
	gin.SetMode(gin.TestMode)
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	invReport := db.InventoryReportModel{
		InnerInventoryReport: db.InnerInventoryReport{
			ID:         "1",
			PropertyID: "1",
		},
	}
	m.InventoryReport.Expect(database.MockGetInvReportByID(c)).Returns(invReport)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Params = gin.Params{
		{Key: "propertyId", Value: "1"},
		{Key: "reportId", Value: "1"},
	}

	middlewares.CheckInventoryReportPropertyOwnership("propertyId", "reportId")(ctx)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCheckInventoryReportOwnership_Latest(t *testing.T) {
	gin.SetMode(gin.TestMode)
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	invReport := db.InventoryReportModel{
		InnerInventoryReport: db.InnerInventoryReport{
			ID:         "1",
			PropertyID: "1",
		},
	}
	m.InventoryReport.Expect(database.MockGetLatestInvReport(c)).Returns(invReport)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Params = gin.Params{
		{Key: "propertyId", Value: "1"},
		{Key: "reportId", Value: "latest"},
	}

	middlewares.CheckInventoryReportPropertyOwnership("propertyId", "reportId")(ctx)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCheckInventoryReportOwnership_LatestNotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.InventoryReport.Expect(database.MockGetLatestInvReport(c)).Errors(db.ErrNotFound)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Params = gin.Params{
		{Key: "propertyId", Value: "1"},
		{Key: "reportId", Value: "latest"},
	}

	middlewares.CheckInventoryReportPropertyOwnership("propertyId", "reportId")(ctx)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestCheckInventoryReportOwnership_NotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.InventoryReport.Expect(database.MockGetInvReportByID(c)).Errors(db.ErrNotFound)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Params = gin.Params{
		{Key: "propertyId", Value: "1"},
		{Key: "reportId", Value: "1"},
	}

	middlewares.CheckInventoryReportPropertyOwnership("propertyId", "reportId")(ctx)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestCheckInventoryReportOwnership_NotYours(t *testing.T) {
	gin.SetMode(gin.TestMode)
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	invReport := db.InventoryReportModel{
		InnerInventoryReport: db.InnerInventoryReport{
			ID:         "1",
			PropertyID: "2",
		},
	}
	m.InventoryReport.Expect(database.MockGetInvReportByID(c)).Returns(invReport)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Params = gin.Params{
		{Key: "propertyId", Value: "1"},
		{Key: "reportId", Value: "1"},
	}

	middlewares.CheckInventoryReportPropertyOwnership("propertyId", "reportId")(ctx)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestGetPropertyByLease(t *testing.T) {
	gin.SetMode(gin.TestMode)
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	lease := db.LeaseModel{
		InnerLease: db.InnerLease{
			ID:         "1",
			PropertyID: "1",
		},
	}
	property := BuildTestProperty("1")
	m.Property.Expect(database.MockGetPropertyByID(c)).Returns(property)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Set("lease", lease)

	middlewares.GetPropertyByLease()(ctx)
	require.Equal(t, http.StatusOK, w.Code)
	p, ok := ctx.Get("property")
	require.True(t, ok)
	assert.NotNil(t, p)
	assert.IsType(t, db.PropertyModel{}, p)
}

func TestGetPropertyByLease_PropertyNotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	lease := db.LeaseModel{
		InnerLease: db.InnerLease{
			ID:         "1",
			PropertyID: "1",
		},
	}
	m.Property.Expect(database.MockGetPropertyByID(c)).Errors(db.ErrNotFound)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Set("lease", lease)

	middlewares.GetPropertyByLease()(ctx)
	assert.Equal(t, http.StatusNotFound, w.Code)
}
