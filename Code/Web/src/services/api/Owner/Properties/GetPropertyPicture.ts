import callApi from '@/services/api/apiCaller'
import { PropertyPictureResponse } from '@/interfaces/Property/Property'

const GetPropertyPicture = async (id: string) => {
  try {
    return await callApi<PropertyPictureResponse>({
      method: 'GET',
      endpoint: `owner/properties/${id}/picture`
    })
  } catch (error) {
    console.error('Error fetching data:', error)
    throw error
  }
}

export default GetPropertyPicture
