import callApi from '@/services/api/apiCaller'
import endpoints from '@/enums/EndPointEnum'

const UpdatePropertyPicture = async (id: string, pictureData: string) => {
  try {
    return await callApi({
      method: 'PUT',
      endpoint: endpoints.owner.properties.picture(id),
      data: JSON.stringify({ data: pictureData })
    })
  } catch (error) {
    console.error('Error updating property picture:', error)
    throw error
  }
}

export default UpdatePropertyPicture
