import callApi from '@/services/api/apiCaller'

import { Furniture } from '@/interfaces/Property/Inventory/Room/Furniture/Furniture'
import endpoints from '@/enums/EndPointEnum'

const GetFurnitureByRoomById = async (
  propertyId: string,
  roomId: string,
  furnitureId: string
): Promise<Furniture> => {
  try {
    return await callApi<Furniture>({
      method: 'GET',
      endpoint: endpoints.owner.properties.rooms.furnitures.byId(
        propertyId,
        roomId,
        furnitureId
      )
    })
  } catch (error) {
    console.error('Error fetching data:', error)
    throw error
  }
}

export default GetFurnitureByRoomById
