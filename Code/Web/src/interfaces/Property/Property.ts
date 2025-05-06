import { InviteTenant } from '../Tenant/InviteTenant'
import { Lease } from './Lease/Lease'

export interface CreatePropertyPayload {
  name: string
  apartment_number: string
  address: string
  city: string
  postal_code: string
  country: string
  area_sqm: number
  rental_price_per_month: number
  deposit_price: number
}

export interface PropertyPictureResponse {
  id: string
  created_at: string
  data: string
}

export interface PropertyDetails {
  address: string
  apartment_number: string
  archived: boolean
  area_sqm: number
  city: string
  country: string
  created_at: string
  deposit_price: number
  id: string
  invite?: InviteTenant
  lease: Lease
  name: string
  nb_damage: number
  owner_id: string
  picture_id: string
  postal_code: string
  rental_price_per_month: number
  status: string
  start_date: string
  end_date: string
  tenant: string
}

export interface CardComponentProps {
  realProperty: any
  t: (key: string) => string
}

export interface RealPropertyCreateProps {
  showModalCreate: boolean
  setShowModalCreate: (show: boolean) => void
  setIsPropertyCreated: (isCreated: boolean) => void
}

export interface RealPropertyUpdateProps {
  propertyData: PropertyDetails
  isModalUpdateOpen: boolean
  setIsModalUpdateOpen: (show: boolean) => void
  setIsPropertyUpdated: (isCreated: boolean) => void
}

export interface DetailsPartProps {
  propertyData: PropertyDetails
  showModal: () => void
  propertyId: string
  showModalUpdate: () => void
}

export interface PropertyImageProps {
  status: string
  picture: string | null
  isLoading: boolean
}

export interface PropertyHeaderProps {
  onShowModal: () => void
  onShowModalUpdate: () => void
  onEndContract: () => void
  onCancelInvitation: () => void
  onRemoveProperty: () => void
  propertyStatus: string
}
