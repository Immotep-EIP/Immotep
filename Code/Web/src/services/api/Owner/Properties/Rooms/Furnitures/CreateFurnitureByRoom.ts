import callApi from '@/services/api/apiCaller'
import { CreateFurniture } from '@/interfaces/Property/Room/Furniture/Furniture'

const CreateFurnitureByRoom = async (
  PropertyId: string,
  RoomId: string,
  data: CreateFurniture
) => {
  try {
    return await callApi({
      method: 'POST',
      endpoint: `owner/properties/${PropertyId}/rooms/${RoomId}/furnitures/`,
      data
    })
  } catch (error) {
    console.error('Error fetching data:', error)
    throw error
  }
}

export default CreateFurnitureByRoom
