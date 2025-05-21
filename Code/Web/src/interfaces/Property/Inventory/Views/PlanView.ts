import { Inventory } from '../Inventory'

export interface PlanViewProps {
  inventory: Inventory
  layouts: any
  setLayouts: (layouts: any) => void
  showModal: (modal: string, roomId?: string) => void
  handleDeleteFurniture: (roomId: string, furnitureId: string) => void
}
