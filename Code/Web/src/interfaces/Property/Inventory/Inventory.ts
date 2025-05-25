import { Room } from './Room/Room'

export interface Inventory {
  rooms: Room[]
}

export interface InventoryStats {
  totalRooms: number
  totalFurniture: number
  roomsByType: {
    [key: string]: number
  }
}

export interface InventoryFilters {
  searchQuery: string
  selectedRoomType: string
  viewMode: 'grid' | 'list' | 'plan'
}

export interface InventoryState {
  inventory: Inventory
  isLoading: boolean
  error: string | null
  filters: InventoryFilters
  layouts: any
}

export interface InventoryActions {
  createRoom: (name: string) => Promise<boolean>
  createFurniture: (
    roomId: string,
    furniture: { name: string; quantity: number }
  ) => Promise<boolean>
  deleteRoom: (roomId: string) => Promise<boolean>
  deleteFurniture: (roomId: string, furnitureId: string) => Promise<boolean>
  updateFilters: (filters: Partial<InventoryFilters>) => void
  updateLayouts: (layouts: any) => void
}

export interface InventoryControlsProps {
  setSearchQuery: (query: string) => void
  selectedRoomType: string
  setSelectedRoomType: (type: string) => void
  viewMode: 'grid' | 'list' | 'plan'
  setViewMode: (mode: 'grid' | 'list' | 'plan') => void
  showModal: (modal: string) => void
  roomTypes: { value: string; label: string; color?: string }[]
}
