import callApi from '@/services/api/apiCaller'

import { Furniture } from '@/interfaces/Property/Inventory/Room/Furniture/Furniture'
import { GetFurnituresByRoomParams } from '@/interfaces/Api/callApi'
import endpoints from '@/enums/EndPointEnum'

const GetFurnituresByRoom = async ({
  propertyId,
  roomId,
  archive = false
}: GetFurnituresByRoomParams): Promise<Furniture[]> => {
  try {
    return await callApi<Furniture[]>({
      method: 'GET',
      endpoint: endpoints.owner.properties.rooms.furnitures.list(
        propertyId,
        roomId
      ),
      params: {
        archive
      }
    })
  } catch (error) {
    console.error('Error fetching data:', error)
    throw error
  }
}

export default GetFurnituresByRoom
