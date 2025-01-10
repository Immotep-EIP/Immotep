package utils

import (
	"github.com/gin-gonic/gin"
)

type ErrorCode string

const (
	InvalidPassword              ErrorCode = "invalid-password"
	CannotFetchUser              ErrorCode = "cannot-fetch-user"
	UserNotFound                 ErrorCode = "user-not-found"
	CannotCreateUser             ErrorCode = "cannot-create-user"
	NoClaims                     ErrorCode = "no-claims"
	CannotDecodeUser             ErrorCode = "cannot-decode-user"
	MissingFields                ErrorCode = "missing-fields"
	CannotHashPassword           ErrorCode = "cannot-hash-password"
	EmailAlreadyExists           ErrorCode = "email-already-exists"
	TestError                    ErrorCode = "test-error"
	TooManyRequests              ErrorCode = "too-many-requests"
	InviteNotFound               ErrorCode = "invite-not-found"
	UserSameEmailAsInvite        ErrorCode = "user-must-have-same-email-as-invite"
	InviteAlreadyExists          ErrorCode = "invite-already-exists-for-email-or-property"
	ContractAlreadyExist         ErrorCode = "contract-already-exists-for-tenant-and-property"
	PropertyNotFound             ErrorCode = "property-not-found"
	PropertyNotYours             ErrorCode = "property-is-not-yours"
	NotAnOwner                   ErrorCode = "not-an-owner"
	NotATenant                   ErrorCode = "not-a-tenant"
	PropertyAlreadyExists        ErrorCode = "property-already-exists"
	PropertyNotAvailable         ErrorCode = "property-not-available"
	FailedLinkImage              ErrorCode = "failed-to-link-image"
	BadBase64String              ErrorCode = "bad-base64-string"
	PropertyPictureNotFound      ErrorCode = "property-picture-not-found"
	UserProfilePictureNotFound   ErrorCode = "user-profile-picture-not-found"
	RoomAlreadyExists            ErrorCode = "room-already-exists"
	RoomNotFound                 ErrorCode = "room-not-found"
	FurnitureNotFound            ErrorCode = "furniture-not-found"
	FurnitureAlreadyExists       ErrorCode = "furniture-already-exists"
	FurnitureNotInThisRoom       ErrorCode = "furniture-not-in-this-room"
	InventoryReportAlreadyExists ErrorCode = "inventory-report-already-exists"
	InventoryReportNotFound      ErrorCode = "inventory-report-not-found"
	RoomStateAlreadyExists       ErrorCode = "room-state-already-exists"
	FurnitureStateAlreadyExists  ErrorCode = "furniture-state-already-exists"
)

type Error struct {
	Code  ErrorCode `json:"code"`
	Error string    `json:"error"`
}

func SendError(c *gin.Context, httpStatus int, code ErrorCode, err error) {
	if err == nil {
		c.JSON(httpStatus, Error{code, ""})
	} else {
		_ = c.Error(err)
		c.JSON(httpStatus, Error{code, err.Error()})
	}
}

func AbortSendError(c *gin.Context, httpStatus int, code ErrorCode, err error) {
	if err == nil {
		c.AbortWithStatusJSON(httpStatus, Error{code, ""})
	} else {
		_ = c.Error(err)
		c.AbortWithStatusJSON(httpStatus, Error{code, err.Error()})
	}
}
