import callApi from '@/services/api/apiCaller'
import { Room } from '@/interfaces/Property/Room/Room'

const GetRoomByPropertyById = async (PropertyId: string, RoomId: string) => {
  try {
    return await callApi<Room>({
      method: 'GET',
      endpoint: `owner/properties/${PropertyId}/rooms/${RoomId}`
    })
  } catch (error) {
    console.error('Error fetching data:', error)
    throw error
  }
}

export default GetRoomByPropertyById
