import callApi from '@/services/api/apiCaller'
import { PropertyDetails } from '@/interfaces/Property/Property'

const GetPropertyDetails = async (id: string) => {
  try {
    return await callApi<PropertyDetails>({
      method: 'GET',
      endpoint: `owner/properties/${id}`,
    })
  } catch (error) {
    console.error('Error fetching data:', error)
    throw error
  }
}

export default GetPropertyDetails
