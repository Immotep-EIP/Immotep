package controllers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"keyz/backend/models"
	"keyz/backend/prisma/db"
	"keyz/backend/router"
	"keyz/backend/services"
	"keyz/backend/services/database"
	"keyz/backend/utils"
)

func BuildTestDashboard(id string) db.PropertyModel {
	return db.PropertyModel{
		InnerProperty: db.InnerProperty{
			ID:                  id,
			Name:                "Test",
			Address:             "Test",
			ApartmentNumber:     utils.Ptr("Test"),
			City:                "Test",
			PostalCode:          "Test",
			Country:             "Test",
			AreaSqm:             20.0,
			RentalPricePerMonth: 500,
			DepositPrice:        1000,
			CreatedAt:           time.Now(),
			OwnerID:             "1",
			PictureID:           utils.Ptr("1"),
			Archived:            false,
		},
		RelationsProperty: db.RelationsProperty{
			Owner: &db.UserModel{
				InnerUser: db.InnerUser{
					ID:        "1",
					Firstname: "Test",
					Lastname:  "Test",
					Email:     "test@example.com",
				},
			},
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
							ID:           "1",
							Read:         true,
							CreatedAt:    time.Now().AddDate(0, 0, -10),
							UpdatedAt:    time.Now().AddDate(0, 0, -5),
							FixPlannedAt: utils.Ptr(time.Now().AddDate(0, 0, 1)),
							Priority:     db.PriorityHigh,
							FixedOwner:   false,
							FixedTenant:  false,
							FixedAt:      nil,
						},
						RelationsDamage: db.RelationsDamage{
							Room: &db.RoomModel{
								InnerRoom: db.InnerRoom{
									ID:   "1",
									Name: "Test Room",
								},
							},
						},
					}},
					Reports: []db.InventoryReportModel{{}},
				},
			}},
			LeaseInvite: &db.LeaseInviteModel{},
			Rooms: []db.RoomModel{{
				InnerRoom: db.InnerRoom{
					ID:   "1",
					Name: "Test Room",
				},
			}},
		},
	}
}

func TestGetOwnerDashboard_Success(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	property := BuildTestDashboard("1")
	m.Property.Expect(database.MockGetAllDatasFromProperties(c)).ReturnsMany([]db.PropertyModel{property})

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/owner/dashboard/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	var resp models.DashboardResponse
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.NotNil(t, resp.Reminders)
	assert.NotNil(t, resp.Properties)
	assert.NotNil(t, resp.OpenDamages)
}

func TestGetOwnerDashboard_EmptyProperties(t *testing.T) {
	c, m, ensure := services.ConnectDBTest()
	defer ensure(t)

	m.Property.Expect(database.MockGetAllDatasFromProperties(c)).ReturnsMany([]db.PropertyModel{})

	r := router.TestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/owner/dashboard/", nil)
	req.Header.Set("Oauth.claims.id", "1")
	req.Header.Set("Oauth.claims.role", string(db.RoleOwner))
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	var resp models.DashboardResponse
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.NotNil(t, resp.Reminders)
	assert.NotNil(t, resp.Properties)
	assert.NotNil(t, resp.OpenDamages)
}
