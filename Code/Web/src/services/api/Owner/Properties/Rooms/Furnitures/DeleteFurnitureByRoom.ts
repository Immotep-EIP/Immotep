import callApi from '@/services/api/apiCaller'

const DeleteFurnitureByRoom = async (
  PropertyId: string,
  RoomId: string,
  FurnitureId: string
) => {
  try {
    return await callApi({
      method: 'DELETE',
      endpoint: `owner/properties/${PropertyId}/rooms/${RoomId}/furnitures/${FurnitureId}/`
    })
  } catch (error) {
    console.error('Error fetching data:', error)
    throw error
  }
}

export default DeleteFurnitureByRoom
