export interface Furniture {
  archived: boolean
  id: string
  name: string
  property_id: string
  quantity: number
  room_id: string
}

export interface CreateFurniturePayload {
  name: string
  quantity: number
}

export interface CreateFurnitureResponse {
  archived: boolean
  id: string
  name: string
  property_id: string
  quantity: number
  room_id: string
}

export interface AddFurnitureModalProps {
  isOpen: boolean
  onOk: () => void
  onCancel: () => void
  form: any
}

export interface FurnitureParams {
  name: string
  quantity: number
}
