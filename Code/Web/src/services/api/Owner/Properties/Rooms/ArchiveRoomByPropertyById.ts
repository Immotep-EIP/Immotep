import callApi from '@/services/api/apiCaller'

const ArchiveRoomByPropertyById = async (
  PropertyId: string,
  RoomId: string
) => {
  try {
    return await callApi({
      method: 'DELETE',
      endpoint: `owner/properties/${PropertyId}/rooms/${RoomId}/`
    })
  } catch (error) {
    console.error('Error fetching data:', error)
    throw error
  }
}

export default ArchiveRoomByPropertyById
