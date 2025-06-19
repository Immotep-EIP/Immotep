import { Dispatch, SetStateAction } from 'react'

import { DBSchema } from 'idb'

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
  leases: Lease[]
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
  refreshPropertyDetails: (propertyId: string) => void
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
  onRecoverProperty: () => void
  propertyStatus: string
  propertyArchived: boolean
}

export interface ImageDBSchema extends DBSchema {
  images: {
    key: string
    value: {
      id: string
      blob: Blob
      timestamp: number
    }
  }
}

export interface PropertyFilterCardProps {
  filters: {
    searchQuery: string
    surfaceRange: string
    status: string
  }
  setFilters: Dispatch<
    SetStateAction<{
      searchQuery: string
      surfaceRange: string
      status: string
    }>
  >
  surfaceRangeOptions: { value: string; label: string }[]
  statusOptions: { value: string; label: string }[]
}
