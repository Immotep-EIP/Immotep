import callApi from '@/services/api/apiCaller'
import { CreateFurniture } from '@/interfaces/Property/Room/Furniture/Furniture'
import endpoints from '@/enums/EndPointEnum'

const CreateFurnitureByRoom = async (
  propertyId: string,
  roomId: string,
  data: CreateFurniture
): Promise<CreateFurniture> => {
  try {
    const response = await callApi({
      method: 'POST',
      endpoint: endpoints.owner.properties.rooms.furnitures.create(
        propertyId,
        roomId
      ),
      body: data
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
