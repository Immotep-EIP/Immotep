package utils

import (
	"github.com/gin-gonic/gin"
)

type ErrorCode string

const (
	InvalidPassword		ErrorCode = "invalid-password"
	CannotFetchUser		ErrorCode = "cannot-fetch-user"
	CannotFindUser		ErrorCode = "cannot-find-user"
	CannotCreateUser	ErrorCode = "cannot-create-user"
	NoClaims			ErrorCode = "no-claims"
	CannotDecodeUser	ErrorCode = "cannot-decode-user"
	MissingFields		ErrorCode = "missing-fields"
	CannotHashPassword	ErrorCode = "cannot-hash-password"
	EmailAlreadyExists	ErrorCode = "email-already-exists"
	TestError			ErrorCode = "test-error"
)

type Error struct {
	Code ErrorCode	`json:"code"`
	Err  string		`json:"error"`
}

func SendError(c *gin.Context, httpStatus int, code ErrorCode, err error) {
	if err == nil {
		c.JSON(httpStatus, Error{code, ""})
	} else {
		c.Error(err)
		c.JSON(httpStatus, Error{code, err.Error()})
	}
}
