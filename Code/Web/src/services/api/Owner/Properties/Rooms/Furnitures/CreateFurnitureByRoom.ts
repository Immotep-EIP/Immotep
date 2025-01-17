import callApi from '@/services/api/apiCaller'
import { CreateFurniture } from '@/interfaces/Property/Room/Furniture/Furniture'

const CreateFurnitureByRoom = async (
  PropertyId: string,
  RoomId: string,
  data: CreateFurniture
): Promise<CreateFurniture> => {
  try {
    const response = await callApi({
      method: 'POST',
      endpoint: `owner/properties/${PropertyId}/rooms/${RoomId}/furnitures/`,
      data
    })

    return {
      id: response.id,
      name: response.name,
      property_id: response.property_id,
      quantity: response.quantity,
      room_id: response.room_id
    }
  } catch (error) {
    console.error('Error fetching data:', error)
    throw error
  }
}

export default CreateFurnitureByRoom
