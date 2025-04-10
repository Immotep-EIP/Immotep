import callApi from '@/services/api/apiCaller'
import endpoints from '@/enums/EndPointEnum'
import { PropertyDetails } from '@/interfaces/Property/Property'

const ArchiveFurnitureByRoom = async (
  propertyId: string,
  roomId: string,
  furnitureId: string
): Promise<PropertyDetails> => {
  try {
    return await callApi<PropertyDetails, { archive: boolean }>({
      method: 'PUT',
      endpoint: endpoints.owner.properties.rooms.furnitures.archive(
        propertyId,
        roomId,
        furnitureId
      ),
      body: {
        archive: true
      }
    })
  } catch (error) {
    console.error('Error fetching data:', error)
    throw error
  }
}

export default ArchiveFurnitureByRoom
