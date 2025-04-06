import callApi from '@/services/api/apiCaller'
import endpoints from '@/enums/EndPointEnum'

const ArchiveFurnitureByRoom = async (
  propertyId: string,
  roomId: string,
  furnitureId: string
) => {
  try {
    return await callApi({
      method: 'PUT',
      endpoint: endpoints.owner.properties.rooms.furnitures.archive(
        propertyId,
        roomId,
        furnitureId
      ),
      data: {
        archive: true
      }
    })
  } catch (error) {
    console.error('Error fetching data:', error)
    throw error
  }
}

export default ArchiveFurnitureByRoom
