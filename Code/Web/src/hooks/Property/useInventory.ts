import { useState, useEffect } from 'react'
import { message } from 'antd'

import { useTranslation } from 'react-i18next'
import GetRoomsByProperty from '@/services/api/Owner/Properties/Rooms/GetRoomsByProperty'
import GetFurnituresByRoom from '@/services/api/Owner/Properties/Rooms/Furnitures/GetFurnituresByRoom'
import CreateRoomByProperty from '@/services/api/Owner/Properties/Rooms/CreateRoomByProperty'
import CreateFurnitureByRoom from '@/services/api/Owner/Properties/Rooms/Furnitures/CreateFurnitureByRoom'
import DeleteRoomByPropertyById from '@/services/api/Owner/Properties/Rooms/ArchiveRoomByPropertyById'
import DeleteFurnitureByRoom from '@/services/api/Owner/Properties/Rooms/Furnitures/ArchiveFurnitureByRoom'

interface Room {
  roomId: string
  roomName: string
  stuffs: Array<{
    id: string
    name: string
  }>
}

interface FurnitureParams {
  name: string
  quantity: number
}

const useInventory = (propertyId: string) => {
  const [inventory, setInventory] = useState<Room[]>([])
  const [isLoading, setIsLoading] = useState(false)
  const [error, setError] = useState<string | null>(null)
  const { t } = useTranslation()

  const fetchInventory = async () => {
    if (!propertyId) {
      setError('Property ID is missing')
      return
    }

    setIsLoading(true)
    try {
      const rooms = await GetRoomsByProperty(propertyId)
      const inventoryData = await Promise.all(
        rooms.map(async room => {
          try {
            const furnitures = await GetFurnituresByRoom(propertyId, room.id)
            return {
              roomId: room.id,
              roomName: room.name,
              stuffs: furnitures.map(furniture => ({
                id: furniture.id,
                name: furniture.name,
                quantity: furniture.quantity
              }))
            }
          } catch (error) {
            console.error(
              `Error fetching furniture for room ${room.name}:`,
              error
            )
            return { roomId: room.id, roomName: room.name, stuffs: [] }
          }
        })
      )
      setInventory(inventoryData)
      setError(null)
    } catch (error) {
      console.error('Error fetching rooms:', error)
      setError('Failed to load inventory')
    } finally {
      setIsLoading(false)
    }
  }

  const createRoom = async (roomName: string) => {
    try {
      const newRoom = await CreateRoomByProperty(propertyId, roomName)
      setInventory(prev => [
        ...prev,
        { roomId: newRoom.id, roomName: newRoom.name, stuffs: [] }
      ])
      message.success(t('components.messages.room_added'))
      return true
    } catch (error) {
      console.error(t('components.messages.room_added_error'))
      message.error('Failed to add room')
      return false
    }
  }

  const createFurniture = async (
    roomId: string,
    furniture: FurnitureParams
  ) => {
    try {
      const newFurniture = await CreateFurnitureByRoom(
        propertyId,
        roomId,
        furniture
      )
      setInventory(prev =>
        prev.map(room =>
          room.roomId === roomId
            ? {
                ...room,
                stuffs: [
                  ...room.stuffs,
                  {
                    id: newFurniture.id as string,
                    name: newFurniture.name,
                    quantity: newFurniture.quantity
                  }
                ]
              }
            : room
        )
      )
      message.success(t('components.messages.furniture_added'))
      return true
    } catch (error) {
      console.error(t('components.messages.furniture_added_error'))
      message.error('Failed to add item')
      return false
    }
  }

  const deleteRoom = async (roomId: string) => {
    try {
      await DeleteRoomByPropertyById(propertyId, roomId)
      setInventory(prev => prev.filter(room => room.roomId !== roomId))
      message.success(t('components.messages.room_deleted'))
      return true
    } catch (error) {
      console.error(t('components.messages.room_deleted_error'))
      message.error('Failed to delete room')
      return false
    }
  }

  const deleteFurniture = async (roomId: string, furnitureId: string) => {
    try {
      await DeleteFurnitureByRoom(propertyId, roomId, furnitureId)
      setInventory(prev =>
        prev.map(room =>
          room.roomId === roomId
            ? {
                ...room,
                stuffs: room.stuffs.filter(stuff => stuff.id !== furnitureId)
              }
            : room
        )
      )
      message.success(t('components.messages.furniture_deleted'))
      return true
    } catch (error) {
      console.error(t('components.messages.furniture_deleted_error'))
      message.error('Failed to delete item')
      return false
    }
  }

  useEffect(() => {
    fetchInventory()
  }, [propertyId])

  return {
    inventory,
    isLoading,
    error,
    createRoom,
    createFurniture,
    deleteRoom,
    deleteFurniture,
    refreshInventory: fetchInventory
  }
}

export default useInventory
