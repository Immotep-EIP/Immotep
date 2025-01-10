export interface Furniture {
  id: string
  name: string
  property_id: string
  quantity: number
  room_id: string
}

export interface CreateFurniture {
  name: string
  quantity: number
}
