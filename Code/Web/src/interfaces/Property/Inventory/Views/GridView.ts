import { Inventory } from '../Inventory'

export interface GridViewProps {
  inventory: Inventory
  showModal: (modal: string, roomId?: string) => void
  handleDeleteRoom: (roomId: string) => void
  handleDeleteFurniture: (roomId: string, furnitureId: string) => void
}
