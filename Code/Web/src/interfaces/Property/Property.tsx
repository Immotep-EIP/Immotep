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
  end_date: string
  id: string
  name: string
  nb_damage: number
  owner_id: string
  picture_id: string
  postal_code: string
  rental_price_per_month: number
  start_date: string
  status: string
  tenant: string
}
