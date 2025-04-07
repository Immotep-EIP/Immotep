import callApi from '@/services/api/apiCaller'
import endpoints from '@/enums/EndPointEnum'
import { PropertyDetails } from '@/interfaces/Property/Property'

const ArchiveRoomByPropertyById = async (
  propertyId: string,
  roomId: string
): Promise<PropertyDetails> => {
  try {
    return await callApi<PropertyDetails, { archive: boolean }>({
      method: 'PUT',
      endpoint: endpoints.owner.properties.rooms.archive(propertyId, roomId),
      body: {
        archive: true
      }
    })
  } catch (error) {
    console.error('Error fetching data:', error)
    throw error
  }
}

export default ArchiveRoomByPropertyById
