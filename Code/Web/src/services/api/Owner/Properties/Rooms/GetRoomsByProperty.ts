import callApi from '@/services/api/apiCaller'

import { Room } from '@/interfaces/Property/Inventory/Room/Room'
import { GetRoomsByPropertyParams } from '@/interfaces/Api/callApi'
import endpoints from '@/enums/EndPointEnum'

const GetRoomsByProperty = async ({
  propertyId,
  archive = false
}: GetRoomsByPropertyParams): Promise<Room[]> => {
  try {
    return await callApi<Room[]>({
      method: 'GET',
      endpoint: endpoints.owner.properties.rooms.list(propertyId),
      params: {
        archive
      }
    })
  } catch (error) {
    console.error('Error fetching data:', error)
    throw error
  }
}

export default GetRoomsByProperty
