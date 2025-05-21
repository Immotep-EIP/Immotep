import callApi from '@/services/api/apiCaller'
import { PropertyPictureResponse } from '@/interfaces/Property/Property'
import endpoints from '@/enums/EndPointEnum'

const GetPropertyPicture = async (
  id: string
): Promise<PropertyPictureResponse> => {
  try {
    return await callApi<PropertyPictureResponse>({
      method: 'GET',
      endpoint: endpoints.owner.properties.picture(id)
    })
  } catch (error) {
    console.error('Error fetching data:', error)
    throw error
  }
}

export default GetPropertyPicture
