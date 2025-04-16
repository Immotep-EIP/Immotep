package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"immotep/backend/models"
	"immotep/backend/prisma/db"
	"immotep/backend/services/database"
	"immotep/backend/utils"
)

func getPictures(pics []string) ([]string, error) {
	picturesId := make([]string, 0, len(pics))

	for i, pic := range pics {
		dbImage := models.StringToDbImage(pic)
		if dbImage == nil {
			return nil, errors.New("invalid base64 string at index " + strconv.Itoa(i))
		}
		newImage := database.CreateImage(*dbImage)
		picturesId = append(picturesId, newImage.ID)
	}
	return picturesId, nil
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
//	@Success		201			{object}	models.IdResponse		"Created damage ID"
//	@Failure		400			{object}	utils.Error				"Missing fields or bad base64 string"
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
	if imgErr != nil {
		utils.SendError(c, http.StatusBadRequest, utils.BadBase64String, imgErr)
		return
	}

	lease, _ := c.MustGet("lease").(db.LeaseModel)
	res := database.CreateDamage(damage, lease.ID, picturesIds)
	c.JSON(http.StatusCreated, models.IdResponse{ID: res.ID})
}

// GetDamagesByProperty godoc
//
//	@Summary		Get property damages
//	@Description	Get all damages of a property, optionally filtered by fixed status
//	@Tags			damage
//	@Accept			json
//	@Produce		json
//	@Param			property_id	path		string					true	"Property ID"
//	@Param			fixed		query		boolean					false	"Filter by fixed status (default: false)"
//	@Success		200			{array}		models.DamageResponse	"List of damages"
//	@Failure		403			{object}	utils.Error				"Property not yours"
//	@Failure		404			{object}	utils.Error				"No active lease"
//	@Failure		500
//	@Security		Bearer
//	@Router			/owner/properties/{property_id}/damages/ [get]
func GetDamagesByProperty(c *gin.Context) {
	fixed := c.DefaultQuery("fixed", "false") == utils.Strue
	damages := database.GetDamagesByPropertyID(c.Param("property_id"), fixed)
	c.JSON(http.StatusOK, utils.Map(damages, models.DbDamageToResponse))
}

// GetDamagesByLease godoc
//
//	@Summary		Get lease damages
//	@Description	Get all damages of a lease, optionally filtered by fixed status
//	@Tags			damage
//	@Accept			json
//	@Produce		json
//	@Param			property_id	path		string					true	"Property ID"
//	@Param			lease_id	path		string					true	"Lease ID"
//	@Param			fixed		query		boolean					false	"Filter by fixed status (default: false)"
//	@Success		200			{array}		models.DamageResponse	"List of damages"
//	@Failure		403			{object}	utils.Error				"Lease not yours"
//	@Failure		404			{object}	utils.Error				"No active lease"
//	@Failure		500
//	@Security		Bearer
//	@Router			/owner/properties/{property_id}/leases/{lease_id}/damages/ [get]
//	@Router			/tenant/leases/{lease_id}/damages/ [get]
func GetDamagesByLease(c *gin.Context) {
	fixed := c.DefaultQuery("fixed", "false") == utils.Strue
	lease, _ := c.MustGet("lease").(db.LeaseModel)
	damages := database.GetDamagesByLeaseID(lease.ID, fixed)
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

// UpdateDamageOwner godoc
//
//	@Summary		Update damage for owner
//	@Description	Update damage on owner side. Owner can only update the read status and the fix planned date.
//	@Tags			damage
//	@Accept			json
//	@Produce		json
//	@Param			property_id	path		string							true	"Property ID"
//	@Param			lease_id	path		string							true	"Lease ID"
//	@Param			damage_id	path		string							true	"Damage ID"
//	@Param			damages		body		models.DamageOwnerUpdateRequest	true	"Damage update request"
//	@Success		200			{object}	models.IdResponse				"Updated damage ID"
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

	damage, _ := c.MustGet("damage").(db.DamageModel)
	if damage.IsFixed() {
		utils.SendError(c, http.StatusBadRequest, utils.CannotUpdateFixedDamage, nil)
		return
	}

	newDamage := database.UpdateDamageOwner(damage, req)
	if newDamage == nil {
		utils.SendError(c, http.StatusConflict, utils.DamageAlreadyExists, nil)
		return
	}
	c.JSON(http.StatusOK, models.IdResponse{ID: newDamage.ID})
}

// UpdateDamageTenant godoc
//
//	@Summary		Update damage for tenant
//	@Description	Update damage on tenant side. Tenant can only update comment, priority and add new pictures.
//	@Tags			damage
//	@Accept			json
//	@Produce		json
//	@Param			property_id	path		string								true	"Property ID"
//	@Param			lease_id	path		string								true	"Lease ID"
//	@Param			damage_id	path		string								true	"Damage ID"
//	@Param			damages		body		models.DamageTenantUpdateRequest	true	"Damage update request"
//	@Success		200			{object}	models.IdResponse					"Updated damage ID"
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
	if damage.IsFixed() {
		utils.SendError(c, http.StatusBadRequest, utils.CannotUpdateFixedDamage, nil)
		return
	}

	picturesIds, imgErr := getPictures(req.AddPictures)
	if imgErr != nil {
		utils.SendError(c, http.StatusBadRequest, utils.BadBase64String, imgErr)
		return
	}

	newDamage := database.UpdateDamageTenant(damage, req, picturesIds)
	if newDamage == nil {
		utils.SendError(c, http.StatusConflict, utils.DamageAlreadyExists, nil)
		return
	}
	c.JSON(http.StatusOK, models.IdResponse{ID: newDamage.ID})
}

// FixDamage godoc
//
//	@Summary		Fix damage for one user
//	@Description	Fix a damage for a tenant or owner. When both users have fixed the damage, it will be marked as fixed.
//	@Tags			damage
//	@Accept			json
//	@Produce		json
//	@Param			property_id	path		string				true	"Property ID"
//	@Param			lease_id	path		string				true	"Lease ID"
//	@Param			damage_id	path		string				true	"Damage ID"
//	@Success		200			{object}	models.IdResponse	"Fixed damage ID"
//	@Failure		400			{object}	utils.Error			"Missing fields"
//	@Failure		403			{object}	utils.Error			"Lease not yours"
//	@Failure		404			{object}	utils.Error			"Damage not found"
//	@Failure		500
//	@Security		Bearer
//	@Router			/owner/properties/{property_id}/leases/{lease_id}/damages/{damage_id}/fix/ [put]
//	@Router			/tenant/leases/{lease_id}/damages/{damage_id}/fix/ [put]
func FixDamage(c *gin.Context) {
	claims := utils.GetClaims(c)

	damage, _ := c.MustGet("damage").(db.DamageModel)
	if damage.IsFixed() {
		utils.SendError(c, http.StatusBadRequest, utils.DamageAlreadyFixed, nil)
		return
	}

	newDamage := database.MarkDamageAsFixed(damage, db.Role(claims["role"]))
	c.JSON(http.StatusOK, models.IdResponse{ID: newDamage.ID})
}
