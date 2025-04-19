import { Furniture } from './Furniture/Furniture'

export interface Room {
  archived: boolean
  id: string
  name: string
  property_id: string
  type: string
  furniture: Furniture[]
}

export interface AddRoomModalProps {
  isOpen: boolean
  onOk: () => void
  onCancel: () => void
  form: any
  roomTypes: { value: string; label: string }[]
}
