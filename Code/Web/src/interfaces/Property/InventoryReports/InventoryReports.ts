import { Room } from '../Inventory/Room/Room'
import { Furniture } from '../Inventory/Room/Furniture/Furniture'

export interface FurnitureInventoryReports extends Furniture {
  cleanliness: string
  note: string
  state: string
}

export interface RoomInventoryReports extends Room {
  cleanliness: string
  note: string
  state: string
  furnitures: FurnitureInventoryReports[]
}

export interface InventoryReports {
  date: string
  id: string
  lease_id: string
  property_id: string
  rooms: RoomInventoryReports[]
  type: string
}

export interface CreateInventoryReportsPayload {
  rooms: RoomInventoryReports[]
  type: string
}

export interface InventoryReportsResponse {
  date: string
  errors: string[]
  id: string
  lease_id: string
  pdf_data: string
  pdf_name: string
  property_id: string
  type: string
}
