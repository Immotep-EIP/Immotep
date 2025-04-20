import { useState, useEffect } from 'react'
import { message } from 'antd'

import { useTranslation } from 'react-i18next'
import GetRoomsByProperty from '@/services/api/Owner/Properties/Rooms/GetRoomsByProperty'
import GetFurnituresByRoom from '@/services/api/Owner/Properties/Rooms/Furnitures/GetFurnituresByRoom'
import CreateRoomByProperty from '@/services/api/Owner/Properties/Rooms/CreateRoomByProperty'
import CreateFurnitureByRoom from '@/services/api/Owner/Properties/Rooms/Furnitures/CreateFurnitureByRoom'
import DeleteRoomByPropertyById from '@/services/api/Owner/Properties/Rooms/ArchiveRoomByPropertyById'
import DeleteFurnitureByRoom from '@/services/api/Owner/Properties/Rooms/Furnitures/ArchiveFurnitureByRoom'
import { Inventory } from '@/interfaces/Property/Inventory/Inventory'
import { Room } from '@/interfaces/Property/Inventory/Room/Room'

interface FurnitureParams {
  name: string
  quantity: number
}

const useInventory = (propertyId: string) => {
  const [inventory, setInventory] = useState<Inventory>({
    rooms: []
  })
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
      const rooms = await GetRoomsByProperty({ propertyId })
      const roomsData: Room[] = rooms.map(room => ({
        id: room.id || '',
        name: room.name || '',
        type: room.type || 'other',
        property_id: propertyId,
        archived: false,
        furniture: []
      }))

      const furniturePromises = rooms.map(async room => {
        try {
          const furnitures = await GetFurnituresByRoom({
            propertyId,
            roomId: room.id
          })
          return furnitures.map(furniture => ({
            id: furniture.id,
            name: furniture.name,
            quantity: furniture.quantity,
            room_id: room.id,
            property_id: propertyId,
            archived: false
          }))
        } catch (error) {
          console.error(
            `Error fetching furniture for room ${room.name}:`,
            error
          )
          return []
        }
      })

      const furnitureResults = await Promise.all(furniturePromises)

      const roomsWithFurniture = roomsData.map((room, index) => ({
        ...room,
        furniture: furnitureResults[index]
      }))

      setInventory({ rooms: roomsWithFurniture })

      setError(null)
    } catch (error) {
      console.error('Error fetching rooms:', error)
      setError('Failed to load inventory')
    } finally {
      setIsLoading(false)
    }
  }

  const createRoom = async (roomName: string, roomType: string) => {
    try {
      const newRoom = await CreateRoomByProperty(propertyId, roomName, roomType)

      setInventory(prev => ({
        ...prev,
        rooms: [
          ...prev.rooms,
          {
            id: newRoom.id,
            name: roomName,
            type: roomType.toLowerCase(),
            property_id: propertyId,
            archived: false,
            furniture: []
          }
        ]
      }))

      message.success(t('components.messages.room_added'))
      return { success: true, roomId: newRoom.id }
    } catch (error) {
      console.error(t('components.messages.room_added_error'))
      message.error('Failed to add room')
      return { success: false, roomId: null }
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

      setInventory(prev => ({
        ...prev,
        rooms: prev.rooms.map(room => {
          if (room.id === roomId) {
            return {
              ...room,
              furniture: [
                ...room.furniture,
                {
                  id: newFurniture.id as string,
                  name: newFurniture.name,
                  quantity: newFurniture.quantity,
                  room_id: roomId,
                  property_id: propertyId,
                  archived: false
                }
              ]
            }
          }
          return room
        })
      }))

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
      setInventory(prev => ({
        ...prev,
        rooms: prev.rooms.filter(room => room.id !== roomId)
      }))

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
      setInventory(prev => ({
        ...prev,
        rooms: prev.rooms.map(room => {
          if (room.id === roomId) {
            return {
              ...room,
              furniture: room.furniture.filter(f => f.id !== furnitureId)
            }
          }
          return room
        })
      }))
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
