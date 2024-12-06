export interface GetProperty {
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
  picture: string
  created_at: string
}

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
  picture: string
  created_at: string
}
