import callApi from '@/services/api/apiCaller'
import { Room } from '@/interfaces/Property/Room/Room'
import endpoints from '@/enums/EndPointEnum'

const CreateRoomByProperty = async (
  propertyId: string,
  PropertyName: string
) => {
  try {
    return await callApi<Room>({
      method: 'POST',
      endpoint: endpoints.owner.properties.rooms.create(propertyId),
      data: JSON.stringify({ name: PropertyName })
    })
  } catch (error) {
    console.error('Error fetching data:', error)
    throw error
  }
}

export default CreateRoomByProperty
