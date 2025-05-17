export interface ReminderType {
  id: string
  priority: string
  title: string
  advice: string
  link: string
}

export interface PropertyType {
  id: string
  name: string
  address: string
  city: string
  postal_code: string
  country: string
  area_sqm: number
  rental_price_per_month: number
  deposit_price: number
  created_at: string
  archived: boolean
  owner_id: string
}

export interface DamageType {
  id: string
  comment: string
  priority: string
  read: boolean
  created_at: string
  updated_at: string
  fix_planned_at: string
  fixed_owner: boolean
  fixed_tenant: boolean
  lease_id: string
  room_id: string
}

export interface Dashboard {
  reminders: ReminderType[]
  properties: {
    nbr_total: number
    nbr_archived: number
    nbr_occupied: number
    nbr_available: number
    nbr_pending_invites: number
    list_recently_added: PropertyType[]
  } | null
  open_damages: {
    nbr_total: number
    nbr_urgent: number
    nbr_high: number
    nbr_medium: number
    nbr_low: number
    nbr_planned_to_fix_this_week: number
    list_to_fix: DamageType[]
  } | null
}

export interface UseDashboardReturn {
  reminders: ReminderType[]
  properties: PropertyType[] | null
  open_damages: DamageType[] | null
  loading: boolean
  error: string | null
  refreshDashboard: () => Promise<void>
}
