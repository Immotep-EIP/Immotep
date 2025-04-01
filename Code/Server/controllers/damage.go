package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"immotep/backend/models"
	"immotep/backend/prisma/db"
	"immotep/backend/services/database"
	"immotep/backend/utils"
)

func getPictures(pics []string) ([]string, error) {
	picturesId := make([]string, 0, len(pics))
	var err error

	for _, pic := range pics {
		dbImage := models.StringToDbImage(pic)
		if dbImage == nil {
			err = errors.New(string(utils.BadBase64String))
		} else {
			newImage := database.CreateImage(*dbImage)
			picturesId = append(picturesId, newImage.ID)
		}
	}
	return picturesId, err
}

// CreateDamage godoc
//
//	@Summary		Create damage
//	@Description	Create a damage to a lease
//	@Tags			damage
//	@Accept			json
//	@Produce		json
//	@Param			property_id	path		string					true	"Property ID"
//	@Param			lease_id	path		string					true	"Lease ID or `current`"
//	@Param			damages		body		models.DamageRequest	true	"Damages to create"
//	@Success		201			{object}	models.DamageResponse	"Damage created"
//	@Failure		400			{object}	utils.Error				"Missing fields"
//	@Failure		403			{object}	utils.Error				"Property not yours"
//	@Failure		404			{object}	utils.Error				"No active lease"
//	@Failure		500
//	@Security		Bearer
//	@Router			/tenant/leases/{lease_id}/damages/ [post]
func CreateDamage(c *gin.Context) {
	var req models.DamageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendError(c, http.StatusBadRequest, utils.MissingFields, err)
		return
	}

	damage := req.ToDbDamage()
	if database.GetRoomByID(damage.RoomID) == nil {
		utils.SendError(c, http.StatusNotFound, utils.RoomNotFound, nil)
		return
	}

	picturesIds, imgErr := getPictures(req.Pictures)

	lease, _ := c.MustGet("lease").(db.LeaseModel)
	res := database.CreateDamage(damage, lease.ID, picturesIds)
	c.JSON(http.StatusCreated, models.DbDamageToCreateResponse(res, imgErr))
}

// GetDamagesByProperty godoc
//
//	@Summary		Get property damages
//	@Description	Get all damages of a property
//	@Tags			damage
//	@Accept			json
//	@Produce		json
//	@Param			property_id	path		string			true	"Property ID"
//	@Success		200			{array}		db.DamageModel	"List of damages"
//	@Failure		403			{object}	utils.Error		"Property not yours"
//	@Failure		404			{object}	utils.Error		"No active lease"
//	@Failure		500
//	@Security		Bearer
//	@Router			/owner/properties/{property_id}/damages/ [get]
func GetDamagesByProperty(c *gin.Context) {
	damages := database.GetDamagesByPropertyID(c.Param("property_id"), false)
	c.JSON(http.StatusOK, utils.Map(damages, models.DbDamageToResponse))
}

// GetFixedDamagesByProperty godoc
//
//	@Summary		Get property fixed damages
//	@Description	Get all fixed damages of a property
//	@Tags			damage
//	@Accept			json
//	@Produce		json
//	@Param			property_id	path		string			true	"Property ID"
//	@Success		200			{array}		db.DamageModel	"List of damages"
//	@Failure		403			{object}	utils.Error		"Property not yours"
//	@Failure		404			{object}	utils.Error		"No active lease"
//	@Failure		500
//	@Security		Bearer
//	@Router			/owner/properties/{property_id}/damages/fixed/ [get]
func GetFixedDamagesByProperty(c *gin.Context) {
	damages := database.GetDamagesByPropertyID(c.Param("property_id"), true)
	c.JSON(http.StatusOK, utils.Map(damages, models.DbDamageToResponse))
}

// GetDamagesByLease godoc
//
//	@Summary		Get lease damages
//	@Description	Get all damages of a lease
//	@Tags			damage
//	@Accept			json
//	@Produce		json
//	@Param			property_id	path		string			true	"Property ID"
//	@Param			lease_id	path		string			true	"Lease ID"
//	@Success		200			{array}		db.DamageModel	"List of damages"
//	@Failure		403			{object}	utils.Error		"Lease not yours"
//	@Failure		404			{object}	utils.Error		"No active lease"
//	@Failure		500
//	@Security		Bearer
//	@Router			/owner/properties/{property_id}/leases/{lease_id}/damages/ [get]
//	@Router			/tenant/leases/{lease_id}/damages/ [get]
func GetDamagesByLease(c *gin.Context) {
	lease, _ := c.MustGet("lease").(db.LeaseModel)
	damages := database.GetDamagesByLeaseID(lease.ID, false)
	c.JSON(http.StatusOK, utils.Map(damages, models.DbDamageToResponse))
}

// GetFixedDamagesByLease godoc
//
//	@Summary		Get lease fixed damages
//	@Description	Get all fixed damages of a lease
//	@Tags			damage
//	@Accept			json
//	@Produce		json
//	@Param			property_id	path		string			true	"Property ID"
//	@Param			lease_id	path		string			true	"Lease ID"
//	@Success		200			{array}		db.DamageModel	"List of damages"
//	@Failure		403			{object}	utils.Error		"Lease not yours"
//	@Failure		404			{object}	utils.Error		"No active lease"
//	@Failure		500
//	@Security		Bearer
//	@Router			/owner/properties/{property_id}/leases/{lease_id}/damages/fixed/ [get]
//	@Router			/tenant/leases/{lease_id}/damages/fixed/ [get]
func GetFixedDamagesByLease(c *gin.Context) {
	lease, _ := c.MustGet("lease").(db.LeaseModel)
	damages := database.GetDamagesByLeaseID(lease.ID, true)
	c.JSON(http.StatusOK, utils.Map(damages, models.DbDamageToResponse))
}

// GetDamage godoc
//
//	@Summary		Get damage
//	@Description	Get a damage
//	@Tags			damage
//	@Accept			json
//	@Produce		json
//	@Param			property_id	path		string					true	"Property ID"
//	@Param			lease_id	path		string					true	"Lease ID"
//	@Param			damage_id	path		string					true	"Damage ID"
//	@Success		200			{object}	models.DamageResponse	"Damage"
//	@Failure		403			{object}	utils.Error				"Lease not yours"
//	@Failure		404			{object}	utils.Error				"Damage not found"
//	@Failure		500
//	@Security		Bearer
//	@Router			/owner/properties/{property_id}/leases/{lease_id}/damages/{damage_id}/ [get]
//	@Router			/tenant/leases/{lease_id}/damages/{damage_id}/ [get]
func GetDamage(c *gin.Context) {
	damage, _ := c.MustGet("damage").(db.DamageModel)
	c.JSON(http.StatusOK, models.DbDamageToResponse(damage))
}

func switchEvent(req models.DamageOwnerUpdateRequest, damage db.DamageModel) *db.DamageModel {
	var newDamage *db.DamageModel

	switch req.Event {
	case models.DamageUpdateEventFixPlanned:
		newDamage = database.UpdateDamageFixPlannedAt(damage.ID, *req.FixPlannedAt)
	case models.DamageUpdateEventFixed:
		newDamage = database.MarkDamageAsFixed(damage.ID)
	case models.DamageUpdateEventRead:
		newDamage = database.MarkDamageAsRead(damage.ID)
	default:
		panic("unknown event")
	}
	return newDamage
}

// UpdateDamageOwner godoc
//
//	@Summary		Update damage from event
//	@Description	Update damage according to event triggered. This can be either fix_planned, fixed or read.
//	@Tags			damage
//	@Accept			json
//	@Produce		json
//	@Param			property_id	path		string							true	"Property ID"
//	@Param			lease_id	path		string							true	"Lease ID"
//	@Param			damage_id	path		string							true	"Damage ID"
//	@Param			damages		body		models.DamageOwnerUpdateRequest	true	"Damage update request with event"
//	@Success		200			{object}	models.DamageResponse			"Updated damage"
//	@Failure		400			{object}	utils.Error						"Missing fields"
//	@Failure		403			{object}	utils.Error						"Property not yours"
//	@Failure		404			{object}	utils.Error						"Damage not found"
//	@Failure		500
//	@Security		Bearer
//	@Router			/owner/properties/{property_id}/leases/{lease_id}/damages/{damage_id}/ [put]
func UpdateDamageOwner(c *gin.Context) {
	var req models.DamageOwnerUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendError(c, http.StatusBadRequest, utils.MissingFields, err)
		return
	}
	if req.Event == models.DamageUpdateEventFixPlanned && req.FixPlannedAt == nil {
		utils.SendError(c, http.StatusBadRequest, utils.MissingFields, errors.New("fix_planned_at is required"))
		return
	}

	damage, _ := c.MustGet("damage").(db.DamageModel)
	if req.Event == models.DamageUpdateEventFixed && damage.InnerDamage.FixedAt != nil {
		utils.SendError(c, http.StatusBadRequest, utils.DamageAlreadyFixed, nil)
		return
	}

	newDamage := switchEvent(req, damage)
	if newDamage == nil {
		utils.SendError(c, http.StatusNotFound, utils.DamageNotFound, nil)
		return
	}
	c.JSON(http.StatusOK, models.DbDamageToResponse(*newDamage))
}

// UpdateDamageTenant godoc
//
//	@Summary		Update damage
//	@Description	Update damage from tenant
//	@Tags			damage
//	@Accept			json
//	@Produce		json
//	@Param			property_id	path		string								true	"Property ID"
//	@Param			lease_id	path		string								true	"Lease ID"
//	@Param			damage_id	path		string								true	"Damage ID"
//	@Param			damages		body		models.DamageTenantUpdateRequest	true	"Damage update request"
//	@Success		200			{object}	models.DamageResponse				"Updated damage"
//	@Failure		400			{object}	utils.Error							"Missing fields"
//	@Failure		403			{object}	utils.Error							"Lease not yours"
//	@Failure		404			{object}	utils.Error							"Damage not found"
//	@Failure		500
//	@Security		Bearer
//	@Router			/tenant/leases/{lease_id}/damages/{damage_id}/ [put]
func UpdateDamageTenant(c *gin.Context) {
	var req models.DamageTenantUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendError(c, http.StatusBadRequest, utils.MissingFields, err)
		return
	}

	damage, _ := c.MustGet("damage").(db.DamageModel)
	picturesIds, imgErr := getPictures(req.AddPictures)

	newDamage := database.UpdateDamage(damage.ID, req, picturesIds)
	if newDamage == nil {
		utils.SendError(c, http.StatusNotFound, utils.DamageNotFound, nil)
		return
	}
	c.JSON(http.StatusOK, models.DbDamageToCreateResponse(*newDamage, imgErr))
}
