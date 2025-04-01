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

func getPictures(req models.DamageRequest) ([]string, error) {
	picturesId := make([]string, 0, len(req.Pictures))
	var err error

	for _, pic := range req.Pictures {
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

	pictures, imgErr := getPictures(req)

	lease, _ := c.MustGet("lease").(db.LeaseModel)
	res := database.CreateDamage(damage, lease.ID, pictures)
	c.JSON(http.StatusCreated, models.DbDamageCreateToResponse(res, imgErr))
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
