package db

func (u UserModel) Name() string {
	return u.Firstname + " " + u.Lastname
}

func (d DamageModel) IsFixed() bool {
	return d.FixedOwner && d.FixedTenant
}

type FixStatus string

const (
	FixStatusPending                    FixStatus = "pending"
	FixStatusPlanned                    FixStatus = "planned"
	FixStatusAwaitingOwnerConfirmation  FixStatus = "awaiting_owner_confirmation"
	FixStatusAwaitingTenantConfirmation FixStatus = "awaiting_tenant_confirmation"
	FixStatusFixed                      FixStatus = "fixed"
)

func (d DamageModel) FixStatus() FixStatus {
	if d.IsFixed() {
		return FixStatusFixed
	} else if d.FixedTenant && !d.FixedOwner {
		return FixStatusAwaitingOwnerConfirmation
	} else if !d.FixedTenant && d.FixedOwner {
		return FixStatusAwaitingTenantConfirmation
	} else if d.InnerDamage.FixPlannedAt != nil {
		return FixStatusPlanned
	} else {
		return FixStatusPending
	}
}
