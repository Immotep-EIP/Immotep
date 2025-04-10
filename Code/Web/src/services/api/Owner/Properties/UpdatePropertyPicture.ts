import callApi from '@/services/api/apiCaller'
import endpoints from '@/enums/EndPointEnum'
import { PropertyDetails } from '@/interfaces/Property/Property'

const UpdatePropertyPicture = async (
  id: string,
  pictureData: string
): Promise<PropertyDetails> => {
  try {
    return await callApi<PropertyDetails, { data: string }>({
      method: 'PUT',
      endpoint: endpoints.owner.properties.picture(id),
      body: JSON.stringify({ data: pictureData })
    })
  } catch (error) {
    console.error('Error updating property picture:', error)
    throw error
  }
}

export default UpdatePropertyPicture
