import callApi from '@/services/api/apiCaller'
import { Room } from '@/interfaces/Property/Room/Room'
import endpoints from '@/enums/EndPointEnum'

const CreateRoomByProperty = async (
  propertyId: string,
  propertyName: string
): Promise<Room> => {
  try {
    return await callApi<Room, { name: string }>({
      method: 'POST',
      endpoint: endpoints.owner.properties.rooms.create(propertyId),
      body: JSON.stringify({ name: propertyName })
    })
  } catch (error) {
    console.error('Error fetching data:', error)
    throw error
  }
}

export default CreateRoomByProperty
