package validators

import (
	"github.com/go-playground/validator/v10"
	"immotep/backend/prisma/db"
)

var RoomType validator.Func = func(fl validator.FieldLevel) bool {
	p, ok := fl.Field().Interface().(db.RoomType)
	if !ok {
		return false
	}
	switch p {
	case
		db.RoomTypeDressing,
		db.RoomTypeLaundryroom,
		db.RoomTypeBedroom,
		db.RoomTypePlayroom,
		db.RoomTypeBathroom,
		db.RoomTypeToilet,
		db.RoomTypeLivingroom,
		db.RoomTypeDiningroom,
		db.RoomTypeKitchen,
		db.RoomTypeHallway,
		db.RoomTypeBalcony,
		db.RoomTypeCellar,
		db.RoomTypeGarage,
		db.RoomTypeStorage,
		db.RoomTypeOffice,
		db.RoomTypeOther:
		return true
	default:
		return false
	}
}
