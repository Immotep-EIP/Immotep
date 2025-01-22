import callApi from '@/services/api/apiCaller'
import { Room } from '@/interfaces/Property/Room/Room'

const CreateRoomByProperty = async (id: string, PropertyName: string) => {
  try {
    return await callApi<Room>({
      method: 'POST',
      endpoint: `owner/properties/${id}/rooms/`,
      data: JSON.stringify({ name: PropertyName })
    })
  } catch (error) {
    console.error('Error fetching data:', error)
    throw error
  }
}

export default CreateRoomByProperty
