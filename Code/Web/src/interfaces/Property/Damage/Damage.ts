export interface Damage {
  id: string
  lease_id: string
  tenant_name: string
  room_id: string
  room_name: string
  comment: string
  priority: string
  read: string
  created_at: string
  updated_at: string
  fix_status: string
  fix_planned_at: string
  pictures: string[]
}

export interface DamageDetails {
  read?: string
  fix_planned_at?: string
}

export interface UseDamagesReturn {
  damages: Damage[] | null
  loading: boolean
  error: string | null
  refreshDamages: (propertyId: string) => Promise<void>
  damage: Damage | null
  getDamageByID: (propertyId: string, damageId: string) => Promise<void>
  updateDamage: (
    propertyId: string,
    damageId: string,
    data: DamageDetails
  ) => Promise<void>
}
