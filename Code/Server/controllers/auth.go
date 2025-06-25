package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maxzerbini/oauth"
	"keyz/backend/models"
	"keyz/backend/prisma/db"
	"keyz/backend/services/database"
	"keyz/backend/utils"
)

// TokenAuth godoc
//
//	@Summary		Authenticate user
//	@Description	Authenticate user with email and password
//	@Tags			auth
//	@Accept			x-www-form-urlencoded
//	@Produce		json
//	@Param			grant_type		formData	string		true	"password / refresh_token"
//	@Param			username		formData	string		false	"User email"
//	@Param			password		formData	string		false	"User password"
//	@Param			refresh_token	formData	string		false	"Refresh token"
//	@Success		200				{object}	oauth.Any	"Token data"
//	@Failure		400				{object}	oauth.Any	"Invalid grant_type"
//	@Failure		401				{object}	oauth.Any	"Unauthorized"
//	@Failure		500
//	@Router			/auth/token/ [post]
func TokenAuth(s *oauth.OAuthBearerServer) func(c *gin.Context) {
	return s.UserCredentials
}

// RegisterOwner godoc
//
//	@Summary		Create a new owner
//	@Description	Create a new user with owner role
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			user	body		models.UserRequest	true	"Owner user data"
//	@Success		201		{object}	models.IdResponse	"Created user ID"
//	@Failure		400		{object}	utils.Error			"Missing fields"
//	@Failure		409		{object}	utils.Error			"Email already exists"
//	@Failure		500
//	@Router			/auth/register/ [post]
func RegisterOwner(c *gin.Context) {
	var userReq models.UserRequest
	err := c.ShouldBindBodyWithJSON(&userReq)
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, utils.MissingFields, err)
		return
	}

	userReq.Password, err = utils.HashPassword(userReq.Password)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, utils.CannotHashPassword, err)
		return
	}

	user := database.CreateUser(userReq.ToDbUser(), db.RoleOwner)
	if user == nil {
		utils.SendError(c, http.StatusConflict, utils.EmailAlreadyExists, nil)
		return
	}
	c.JSON(http.StatusCreated, models.IdResponse{ID: user.ID})
}

// RegisterTenant godoc
//
//	@Summary		Create a new tenant
//	@Description	Answer an invite from an owner with an invite link by creating a new user with tenant role
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string				true	"Pending lease ID"
//	@Param			user	body		models.UserRequest	true	"Tenant user data"
//	@Success		201		{object}	models.IdResponse	"Created user ID"
//	@Failure		400		{object}	utils.Error			"Missing fields"
//	@Failure		404		{object}	utils.Error			"Pending lease not found"
//	@Failure		409		{object}	utils.Error			"Email already exists"
//	@Failure		500
//	@Router			/auth/invite/{id}/ [post]
func RegisterTenant(c *gin.Context) {
	var userReq models.UserRequest
	err := c.ShouldBindBodyWithJSON(&userReq)
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, utils.MissingFields, err)
		return
	}
	userReq.Email = utils.SanitizeEmail(userReq.Email)

	leaseInvite := database.GetLeaseInviteById(c.Param("id"))
	if leaseInvite == nil {
		utils.SendError(c, http.StatusNotFound, utils.InviteNotFound, nil)
		return
	}
	if leaseInvite.TenantEmail != userReq.Email {
		utils.SendError(c, http.StatusBadRequest, utils.UserSameEmailAsInvite, nil)
		return
	}

	if database.GetCurrentActiveLeaseByProperty(leaseInvite.PropertyID) != nil {
		utils.SendError(c, http.StatusConflict, utils.PropertyNotAvailable, nil)
		return
	}

	userReq.Password, err = utils.HashPassword(userReq.Password)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, utils.CannotHashPassword, err)
		return
	}

	user := database.CreateUser(userReq.ToDbUser(), db.RoleTenant)
	if user == nil {
		utils.SendError(c, http.StatusConflict, utils.EmailAlreadyExists, nil)
		return
	}

	_ = database.CreateLease(*leaseInvite, *user)
	c.JSON(http.StatusCreated, models.IdResponse{ID: user.ID})
}

// AcceptInvite godoc
//
//	@Summary		Accept an invite
//	@Description	Answer an invite from an owner with an invite link by accepting the invite
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			id	path	string	true	"Pending lease ID"
//	@Success		204	"Accepted"
//	@Failure		403	{object}	utils.Error	"Not a tenant"
//	@Failure		404	{object}	utils.Error	"Pending lease not found"
//	@Failure		409	{object}	utils.Error	"Property not available or tenant already has lease"
//	@Failure		500
//	@Router			/tenant/invite/{id}/ [post]
func AcceptInvite(c *gin.Context) {
	claims := utils.GetClaims(c)
	user := database.GetUserByID(claims["id"])
	if user == nil || user.Role != db.RoleTenant {
		utils.SendError(c, http.StatusForbidden, utils.NotATenant, nil)
		return
	}

	leaseInvite := database.GetLeaseInviteById(c.Param("id"))
	if leaseInvite == nil {
		utils.SendError(c, http.StatusNotFound, utils.InviteNotFound, nil)
		return
	}
	if leaseInvite.TenantEmail != user.Email {
		utils.SendError(c, http.StatusForbidden, utils.UserSameEmailAsInvite, nil)
		return
	}

	if database.GetCurrentActiveLeaseByProperty(leaseInvite.PropertyID) != nil {
		utils.SendError(c, http.StatusConflict, utils.PropertyNotAvailable, nil)
		return
	}
	if database.GetCurrentActiveLeaseByTenant(user.ID) != nil {
		utils.SendError(c, http.StatusConflict, utils.TenantAlreadyHasLease, nil)
		return
	}

	_ = database.CreateLease(*leaseInvite, *user)
	c.Status(http.StatusNoContent)
}
