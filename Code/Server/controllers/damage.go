package controllers

import (
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
	"immotep/backend/models"
	"immotep/backend/prisma/db"
	"immotep/backend/services/database"
	"immotep/backend/services/minio"
	"immotep/backend/utils"
)

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

	lease, _ := c.MustGet("lease").(db.LeaseModel)
	res := database.CreateDamage(damage, lease.ID)
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
	c.JSON(http.StatusOK, utils.Map(damages, func(damage db.DamageModel) models.DamageResponse {
		return models.DbDamageToResponse(damage, minio.GetImageURLs(damage.Pictures))
	}))
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
	c.JSON(http.StatusOK, utils.Map(damages, func(damage db.DamageModel) models.DamageResponse {
		return models.DbDamageToResponse(damage, minio.GetImageURLs(damage.Pictures))
	}))
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
	c.JSON(http.StatusOK, models.DbDamageToResponse(damage, minio.GetImageURLs(damage.Pictures)))
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

	newDamage := database.UpdateDamageTenant(damage, req)
	if newDamage == nil {
		utils.SendError(c, http.StatusConflict, utils.DamageAlreadyExists, nil)
		return
	}
	c.JSON(http.StatusOK, models.IdResponse{ID: newDamage.ID})
}

// AddPicturesToDamage godoc
//
//	@Summary		Add pictures to damage
//	@Description	Add pictures to a damage. Only tenant can add pictures.
//	@Tags			damage
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			lease_id	path		string				true	"Lease ID"
//	@Param			damage_id	path		string				true	"Damage ID"
//	@Param			pictures	formData	[]file				true	"Files to upload"
//	@Success		200			{object}	models.IdResponse	"Updated damage ID"
//	@Failure		400			{object}	utils.Error			"Missing fields"
//	@Failure		403			{object}	utils.Error			"Lease not yours"
//	@Failure		404			{object}	utils.Error			"Damage not found"
//	@Failure		500
//	@Security		Bearer
//	@Router			/tenant/leases/{lease_id}/damages/{damage_id}/pictures/ [post]
func AddPicturesToDamage(c *gin.Context) {
	damage, _ := c.MustGet("damage").(db.DamageModel)
	if damage.IsFixed() {
		utils.SendError(c, http.StatusBadRequest, utils.CannotUpdateFixedDamage, nil)
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, utils.MissingFile, err)
		return
	}

	files := form.File["pictures"]
	if len(files) == 0 {
		utils.SendError(c, http.StatusBadRequest, utils.MissingFile, err)
		return
	}

	picturePaths := getDamagePicturesPath(damage, files)
	newDamage := database.AddPicturesToDamage(damage, picturePaths)
	c.JSON(http.StatusOK, models.IdResponse{ID: newDamage.ID})
}

func getDamagePicturesPath(damage db.DamageModel, files []*multipart.FileHeader) []string {
	picturePaths := make([]string, len(files))
	for i, file := range files {
		fileInfo := minio.UploadDamageImage(damage.ID, file)
		picturePaths[i] = fileInfo.Key
	}
	return picturePaths
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
