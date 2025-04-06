import callApi from '@/services/api/apiCaller'
import { Room } from '@/interfaces/Property/Room/Room'
import endpoints from '@/enums/EndPointEnum'

const GetRoomsByProperty = async (propertyId: string) => {
  try {
    return await callApi<Room[]>({
      method: 'GET',
      endpoint: endpoints.owner.properties.rooms.list(propertyId)
    })
  } catch (error) {
    console.error('Error fetching data:', error)
    throw error
  }
}

export default GetRoomsByProperty
