export interface Lease {
  active: boolean
  created_at: string
  end_date: string
  id: string
  owner_email: string
  owner_id: string
  owner_name: string
  property_id: string
  property_name: string
  start_date: string
  tenant_email: string
  tenant_id: string
  tenant_name: string
}

export interface EndLeaseResponse {
  success: boolean
  message?: string
}
