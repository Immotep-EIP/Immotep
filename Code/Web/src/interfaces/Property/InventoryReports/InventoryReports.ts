import { Room } from '../Room/Room'
import { Furniture } from '../Room/Furniture/Furniture'

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
  property_id: string
  rooms: RoomInventoryReports[]
  type: string
}

export interface CreateInventoryReports {
  rooms: RoomInventoryReports[]
  type: string
}
