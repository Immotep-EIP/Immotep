export interface CreateProperty {
  name: string
  address: string
  city: string
  postal_code: string
  country: string
  area_sqm: number
  rental_price_per_month: number
  deposit_price: number
}

export interface PropertyDetails {
  id: string
  owner_id: string
  name: string
  address: string
  city: string
  postal_code: string
  country: string
  area_sqm: number
  rental_price_per_month: number
  deposit_price: number
  picture_id: string
  created_at: string
  nb_damage: number
  status: string
  tenant: string
  start_date: string
  end_date: string
}

export interface PropertyPictureResponse {
  id: string
  created_at: string
  data: string
}
