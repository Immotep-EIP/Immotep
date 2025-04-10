import callApi from '@/services/api/apiCaller'
import {
  CreateFurniturePayload,
  CreateFurnitureResponse
} from '@/interfaces/Property/Room/Furniture/Furniture'
import endpoints from '@/enums/EndPointEnum'

const CreateFurnitureByRoom = async (
  propertyId: string,
  roomId: string,
  data: CreateFurniturePayload
): Promise<CreateFurnitureResponse> => {
  try {
    return await callApi<CreateFurnitureResponse, CreateFurniturePayload>({
      method: 'POST',
      endpoint: endpoints.owner.properties.rooms.furnitures.create(
        propertyId,
        roomId
      ),
      body: data
    })
  } catch (error) {
    console.error('Error fetching data:', error)
    throw error
  }
}

export default CreateFurnitureByRoom
