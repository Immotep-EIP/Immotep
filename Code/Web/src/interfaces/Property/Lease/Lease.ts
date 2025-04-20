export interface Lease {
  active: boolean
  created_at: string
  start_date: string
  end_date: string
  id: string
  owner_email: string
  owner_id: string
  owner_name: string
  property_id: string
  property_name: string
  tenant_email: string
  tenant_id: string
  tenant_name: string
}
