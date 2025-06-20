package validators_test

import (
	"reflect"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"keyz/backend/prisma/db"
	"keyz/backend/router/validators"
)

type MockFieldLevel struct {
	validator.FieldLevel
	Val any
}

func (m MockFieldLevel) Field() reflect.Value {
	return reflect.ValueOf(m.Val)
}

func TestRoomType(t *testing.T) {
	validTypes := []db.RoomType{
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
		db.RoomTypeOther,
	}
	for _, typ := range validTypes {
		assert.True(t, validators.RoomType(MockFieldLevel{Val: typ}))
	}
	assert.False(t, validators.RoomType(MockFieldLevel{Val: "invalid"}))
}

func TestReportType(t *testing.T) {
	validTypes := []db.ReportType{
		db.ReportTypeStart,
		db.ReportTypeMiddle,
		db.ReportTypeEnd,
	}
	for _, typ := range validTypes {
		assert.True(t, validators.ReportType(MockFieldLevel{Val: typ}))
	}
	assert.False(t, validators.ReportType(MockFieldLevel{Val: "invalid"}))
}

func TestState(t *testing.T) {
	validStates := []db.State{
		db.StateBroken,
		db.StateNeedsRepair,
		db.StateBad,
		db.StateMedium,
		db.StateGood,
		db.StateNew,
	}
	for _, state := range validStates {
		assert.True(t, validators.State(MockFieldLevel{Val: state}))
	}
	assert.False(t, validators.State(MockFieldLevel{Val: "invalid"}))
}

func TestCleanliness(t *testing.T) {
	validClean := []db.Cleanliness{
		db.CleanlinessDirty,
		db.CleanlinessMedium,
		db.CleanlinessClean,
	}
	for _, c := range validClean {
		assert.True(t, validators.Cleanliness(MockFieldLevel{Val: c}))
	}
	assert.False(t, validators.Cleanliness(MockFieldLevel{Val: "invalid"}))
}

func TestPriority(t *testing.T) {
	validPriorities := []db.Priority{
		db.PriorityLow,
		db.PriorityMedium,
		db.PriorityHigh,
		db.PriorityUrgent,
	}
	for _, p := range validPriorities {
		assert.True(t, validators.Priority(MockFieldLevel{Val: p}))
	}
	assert.False(t, validators.Priority(MockFieldLevel{Val: "invalid"}))
}
