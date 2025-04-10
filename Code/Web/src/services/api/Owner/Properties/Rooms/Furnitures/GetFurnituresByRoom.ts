import callApi from '@/services/api/apiCaller'
import { Furniture } from '@/interfaces/Property/Room/Furniture/Furniture'
import endpoints from '@/enums/EndPointEnum'

const GetFurnituresByRoom = async (
  propertyId: string,
  roomId: string
): Promise<Furniture[]> => {
  try {
    return await callApi<Furniture[]>({
      method: 'GET',
      endpoint: endpoints.owner.properties.rooms.furnitures.list(
        propertyId,
        roomId
      )
    })
  } catch (error) {
    console.error('Error fetching data:', error)
    throw error
  }
}

export default GetFurnituresByRoom
