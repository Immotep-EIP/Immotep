import callApi from '@/services/api/apiCaller'
import endpoints from '@/enums/EndPointEnum'
import { ROOM_TYPES, isValidRoomType } from '@/utils/types/roomTypes'

const CreateRoomByProperty = async (
  propertyId: string,
  roomName: string,
  roomType: string
): Promise<{ id: string }> => {
  try {
    if (!isValidRoomType(roomType.toLowerCase())) {
      throw new Error(
        `Invalid room type. Must be one of: ${ROOM_TYPES.join(', ')}`
      )
    }

    return await callApi<{ id: string }, { name: string; type: string }>({
      method: 'POST',
      endpoint: endpoints.owner.properties.rooms.create(propertyId),
      body: JSON.stringify({
        name: roomName,
        type: roomType.toLowerCase()
      })
    })
  } catch (error) {
    console.error('Error creating room:', error)
    throw error
  }
}

export default CreateRoomByProperty
