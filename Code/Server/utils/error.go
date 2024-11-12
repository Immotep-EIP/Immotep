package utils

import (
	"github.com/gin-gonic/gin"
)

type ErrorCode string

const (
	InvalidPassword         ErrorCode = "invalid-password"
	CannotFetchUser         ErrorCode = "cannot-fetch-user"
	UserNotFound            ErrorCode = "user-not-found"
	CannotCreateUser        ErrorCode = "cannot-create-user"
	NoClaims                ErrorCode = "no-claims"
	CannotDecodeUser        ErrorCode = "cannot-decode-user"
	MissingFields           ErrorCode = "missing-fields"
	CannotHashPassword      ErrorCode = "cannot-hash-password"
	EmailAlreadyExists      ErrorCode = "email-already-exists"
	TestError               ErrorCode = "test-error"
	TooManyRequests         ErrorCode = "too-many-requests"
	PendingContractNotFound ErrorCode = "pending-contract-not-found"
	// User must have the same email as the pending contract
	UserSameEmailAsPendingContract ErrorCode = "user-same-email-as-pending-contract"
	// A contract alread exists for that user and property
	ContractAlreadyExist ErrorCode = "contract-already-exists"
	PropertyNotFound     ErrorCode = "property-not-found"
	PropertyNotYours     ErrorCode = "property-is-not-yours"
	InviteAlreadyExists  ErrorCode = "invite-already-exists-for-email"
	NotAnOwner           ErrorCode = "not-an-owner"
	NotATenant           ErrorCode = "not-a-tenant"
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
