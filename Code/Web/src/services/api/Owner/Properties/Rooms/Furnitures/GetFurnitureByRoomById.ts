import callApi from '@/services/api/apiCaller'
import { Furniture } from '@/interfaces/Property/Room/Furniture/Furniture'

const GetFurnitureByRoomById = async (
  PropertyId: string,
  RoomId: string,
  FurnitureId: string
) => {
  try {
    return await callApi<Furniture>({
      method: 'GET',
      endpoint: `owner/properties/${PropertyId}/rooms/${RoomId}/furnitures/${FurnitureId}`
    })
  } catch (error) {
    console.error('Error fetching data:', error)
    throw error
  }
}

export default GetFurnitureByRoomById
