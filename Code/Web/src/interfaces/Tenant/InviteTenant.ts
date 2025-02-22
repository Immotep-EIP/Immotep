import { PropertyDetails } from '../Property/Property'

export interface InviteTenant {
  propertyId: string
  end_date?: string
  start_date: string
  tenant_email: string
}

export interface InviteTenantResponse {
  created_at: string
  end_date: string
  id: string
  property_id: string
  start_date: string
  tenant_email: string
}

export interface InviteTenantModalProps {
  isOpen: boolean
  onClose: (invitationSent: boolean) => void
  property: PropertyDetails
}
