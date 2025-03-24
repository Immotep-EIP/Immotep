package middlewares_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"immotep/backend/prisma/db"
	"immotep/backend/router/middlewares"
	"immotep/backend/services"
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
			Damages: []db.DamageModel{{}},
			Leases:  []db.LeaseModel{{}},
		},
	}
}

func TestCheckPropertyOwnership(t *testing.T) {
	gin.SetMode(gin.TestMode)

	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals("1")).With(
			db.Property.Damages.Fetch(),
			db.Property.Leases.Fetch().With(db.Lease.Tenant.Fetch()),
			db.Property.PendingContract.Fetch(),
		),
	).Returns(property)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{gin.Param{Key: "propertyId", Value: "1"}}
	c.Set("oauth.claims", map[string]string{"id": "1"})

	middlewares.CheckPropertyOwnership("propertyId")(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCheckPropertyOwnership_NotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)

	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals("1")).With(
			db.Property.Damages.Fetch(),
			db.Property.Leases.Fetch().With(db.Lease.Tenant.Fetch()),
			db.Property.PendingContract.Fetch(),
		),
	).Errors(db.ErrNotFound)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{gin.Param{Key: "propertyId", Value: "1"}}
	c.Set("oauth.claims", map[string]string{"id": "1"})

	middlewares.CheckPropertyOwnership("propertyId")(c)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestCheckPropertyOwnership_NotYours(t *testing.T) {
	gin.SetMode(gin.TestMode)

	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestProperty("1")
	mock.Property.Expect(
		client.Client.Property.FindUnique(db.Property.ID.Equals("1")).With(
			db.Property.Damages.Fetch(),
			db.Property.Leases.Fetch().With(db.Lease.Tenant.Fetch()),
			db.Property.PendingContract.Fetch(),
		),
	).Returns(property)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{gin.Param{Key: "propertyId", Value: "1"}}
	c.Set("oauth.claims", map[string]string{"id": "2"})

	middlewares.CheckPropertyOwnership("propertyId")(c)
	assert.Equal(t, http.StatusForbidden, w.Code)
}

func TestCheckRoomOwnership(t *testing.T) {
	gin.SetMode(gin.TestMode)

	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	room := db.RoomModel{
		InnerRoom: db.InnerRoom{
			ID:         "1",
			PropertyID: "1",
		},
	}
	mock.Room.Expect(
		client.Client.Room.FindUnique(db.Room.ID.Equals("1")),
	).Returns(room)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{
		{Key: "propertyId", Value: "1"},
		{Key: "roomId", Value: "1"},
	}

	middlewares.CheckRoomOwnership("propertyId", "roomId")(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCheckRoomOwnership_NotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)

	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	mock.Room.Expect(
		client.Client.Room.FindUnique(db.Room.ID.Equals("1")),
	).Errors(db.ErrNotFound)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{
		{Key: "propertyId", Value: "1"},
		{Key: "roomId", Value: "1"},
	}

	middlewares.CheckRoomOwnership("propertyId", "roomId")(c)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestCheckRoomOwnership_NotYours(t *testing.T) {
	gin.SetMode(gin.TestMode)

	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	room := db.RoomModel{
		InnerRoom: db.InnerRoom{
			ID:         "1",
			PropertyID: "2",
		},
	}
	mock.Room.Expect(
		client.Client.Room.FindUnique(db.Room.ID.Equals("1")),
	).Returns(room)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{
		{Key: "propertyId", Value: "1"},
		{Key: "roomId", Value: "1"},
	}

	middlewares.CheckRoomOwnership("propertyId", "roomId")(c)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestCheckFurnitureOwnership(t *testing.T) {
	gin.SetMode(gin.TestMode)

	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	furniture := db.FurnitureModel{
		InnerFurniture: db.InnerFurniture{
			ID:     "1",
			RoomID: "1",
		},
	}
	mock.Furniture.Expect(
		client.Client.Furniture.FindUnique(
			db.Furniture.ID.Equals("1"),
		).With(
			db.Furniture.Room.Fetch(),
		),
	).Returns(furniture)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{
		{Key: "roomId", Value: "1"},
		{Key: "furnitureId", Value: "1"},
	}

	middlewares.CheckFurnitureOwnership("roomId", "furnitureId")(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCheckFurnitureOwnership_NotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)

	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	mock.Furniture.Expect(
		client.Client.Furniture.FindUnique(
			db.Furniture.ID.Equals("1"),
		).With(
			db.Furniture.Room.Fetch(),
		),
	).Errors(db.ErrNotFound)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{
		{Key: "roomId", Value: "1"},
		{Key: "furnitureId", Value: "1"},
	}

	middlewares.CheckFurnitureOwnership("roomId", "furnitureId")(c)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestCheckFurnitureOwnership_NotYours(t *testing.T) {
	gin.SetMode(gin.TestMode)

	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	furniture := db.FurnitureModel{
		InnerFurniture: db.InnerFurniture{
			ID:     "1",
			RoomID: "2",
		},
	}
	mock.Furniture.Expect(
		client.Client.Furniture.FindUnique(
			db.Furniture.ID.Equals("1"),
		).With(
			db.Furniture.Room.Fetch(),
		),
	).Returns(furniture)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{
		{Key: "roomId", Value: "1"},
		{Key: "furnitureId", Value: "1"},
	}

	middlewares.CheckFurnitureOwnership("roomId", "furnitureId")(c)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestCheckInventoryReportOwnership(t *testing.T) {
	gin.SetMode(gin.TestMode)

	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	invReport := db.InventoryReportModel{
		InnerInventoryReport: db.InnerInventoryReport{
			ID:         "1",
			PropertyID: "1",
		},
	}
	mock.InventoryReport.Expect(
		client.Client.InventoryReport.FindUnique(
			db.InventoryReport.ID.Equals("1"),
		).With(
			db.InventoryReport.Property.Fetch(),
			db.InventoryReport.RoomStates.Fetch().With(db.RoomState.Room.Fetch()).With(db.RoomState.Pictures.Fetch()),
			db.InventoryReport.FurnitureStates.Fetch().With(db.FurnitureState.Furniture.Fetch()).With(db.FurnitureState.Pictures.Fetch()),
		),
	).Returns(invReport)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{
		{Key: "propertyId", Value: "1"},
		{Key: "reportId", Value: "1"},
	}

	middlewares.CheckInventoryReportOwnership("propertyId", "reportId")(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCheckInventoryReportOwnership_Latest(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{
		{Key: "propertyId", Value: "1"},
		{Key: "reportId", Value: "latest"},
	}

	middlewares.CheckInventoryReportOwnership("propertyId", "reportId")(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCheckInventoryReportOwnership_NotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)

	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	mock.InventoryReport.Expect(
		client.Client.InventoryReport.FindUnique(
			db.InventoryReport.ID.Equals("1"),
		).With(
			db.InventoryReport.Property.Fetch(),
			db.InventoryReport.RoomStates.Fetch().With(db.RoomState.Room.Fetch()).With(db.RoomState.Pictures.Fetch()),
			db.InventoryReport.FurnitureStates.Fetch().With(db.FurnitureState.Furniture.Fetch()).With(db.FurnitureState.Pictures.Fetch()),
		),
	).Errors(db.ErrNotFound)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{
		{Key: "propertyId", Value: "1"},
		{Key: "reportId", Value: "1"},
	}

	middlewares.CheckInventoryReportOwnership("propertyId", "reportId")(c)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestCheckInventoryReportOwnership_NotYours(t *testing.T) {
	gin.SetMode(gin.TestMode)

	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	invReport := db.InventoryReportModel{
		InnerInventoryReport: db.InnerInventoryReport{
			ID:         "1",
			PropertyID: "2",
		},
	}
	mock.InventoryReport.Expect(
		client.Client.InventoryReport.FindUnique(
			db.InventoryReport.ID.Equals("1"),
		).With(
			db.InventoryReport.Property.Fetch(),
			db.InventoryReport.RoomStates.Fetch().With(db.RoomState.Room.Fetch()).With(db.RoomState.Pictures.Fetch()),
			db.InventoryReport.FurnitureStates.Fetch().With(db.FurnitureState.Furniture.Fetch()).With(db.FurnitureState.Pictures.Fetch()),
		),
	).Returns(invReport)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{
		{Key: "propertyId", Value: "1"},
		{Key: "reportId", Value: "1"},
	}

	middlewares.CheckInventoryReportOwnership("propertyId", "reportId")(c)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

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
		),
	).ReturnsMany([]db.LeaseModel{lease1, lease2})

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{gin.Param{Key: "propertyId", Value: "1"}}

	assert.Panics(t, func() {
		middlewares.CheckActiveLease("propertyId")(c)
	})
}

func TestCheckPendingContract(t *testing.T) {
	gin.SetMode(gin.TestMode)

	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	pendingContract := db.PendingContractModel{
		InnerPendingContract: db.InnerPendingContract{
			ID:         "1",
			PropertyID: "1",
		},
	}
	mock.PendingContract.Expect(
		client.Client.PendingContract.FindUnique(db.PendingContract.PropertyID.Equals("1")),
	).Returns(pendingContract)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{gin.Param{Key: "propertyId", Value: "1"}}

	middlewares.CheckPendingContract("propertyId")(c)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCheckPendingContract_NotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)

	client, mock, ensure := services.ConnectDBTest()
	defer ensure(t)

	mock.PendingContract.Expect(
		client.Client.PendingContract.FindUnique(db.PendingContract.PropertyID.Equals("1")),
	).Errors(db.ErrNotFound)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{gin.Param{Key: "propertyId", Value: "1"}}

	middlewares.CheckPendingContract("propertyId")(c)
	assert.Equal(t, http.StatusNotFound, w.Code)
}
