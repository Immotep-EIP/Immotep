package validators

import (
	"github.com/go-playground/validator/v10"
	"immotep/backend/prisma/db"
)

var Type validator.Func = func(fl validator.FieldLevel) bool {
	p, ok := fl.Field().Interface().(db.ReportType)
	if !ok {
		return false
	}
	switch p {
	case db.ReportTypeStart, db.ReportTypeMiddle, db.ReportTypeEnd:
		return true
	default:
		return false
	}
}

var State validator.Func = func(fl validator.FieldLevel) bool {
	p, ok := fl.Field().Interface().(db.State)
	if !ok {
		return false
	}
	switch p {
	case db.StateBroken, db.StateNeedsRepair, db.StateBad, db.StateMedium, db.StateGood, db.StateNew:
		return true
	default:
		return false
	}
}

var Cleanliness validator.Func = func(fl validator.FieldLevel) bool {
	p, ok := fl.Field().Interface().(db.Cleanliness)
	if !ok {
		return false
	}
	switch p {
	case db.CleanlinessDirty, db.CleanlinessMedium, db.CleanlinessClean:
		return true
	default:
		return false
	}
}
