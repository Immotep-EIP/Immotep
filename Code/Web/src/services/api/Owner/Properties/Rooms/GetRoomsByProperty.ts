import callApi from '@/services/api/apiCaller'
import { Room } from '@/interfaces/Property/Room/Room'

const GetRoomsByProperty = async (id: string) => {
  try {
    return await callApi<Room[]>({
      method: 'GET',
      endpoint: `owner/properties/${id}/rooms/`
    })
  } catch (error) {
    console.error('Error fetching data:', error)
    throw error
  }
}

export default GetRoomsByProperty
