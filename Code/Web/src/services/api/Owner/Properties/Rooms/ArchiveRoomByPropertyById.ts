import callApi from '@/services/api/apiCaller'
import endpoints from '@/enums/EndPointEnum'

const ArchiveRoomByPropertyById = async (
  propertyId: string,
  roomId: string
) => {
  try {
    return await callApi({
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
