import callApi from '@/services/api/apiCaller'
import { CreateProperty, PropertyDetails } from '@/interfaces/Property/Property'

const UpdatePropertyFunction = async (
  data: CreateProperty,
  id: string
): Promise<PropertyDetails> => {
  try {
    return await callApi({
      method: 'PUT',
      endpoint: `owner/properties/${id}/`,
      data
    })
  } catch (error) {
    console.error('Error fetching data:', error)
    throw error
  }
}

export default UpdatePropertyFunction
