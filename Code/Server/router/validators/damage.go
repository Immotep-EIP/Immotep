package validators

import (
	"github.com/go-playground/validator/v10"
	"keyz/backend/prisma/db"
)

var Priority validator.Func = func(fl validator.FieldLevel) bool {
	p, ok := fl.Field().Interface().(db.Priority)
	if !ok {
		return false
	}
	switch p {
	case db.PriorityLow, db.PriorityMedium, db.PriorityHigh, db.PriorityUrgent:
		return true
	default:
		return false
	}
}
