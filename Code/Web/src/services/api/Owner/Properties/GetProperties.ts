import callApi from '@/services/api/apiCaller'
import { PropertyDetails } from '@/interfaces/Property/Property'

const GetProperties = async () => {
  try {
    return await callApi<PropertyDetails[]>({
      method: 'GET',
      endpoint: 'owner/properties/'
    })
  } catch (error) {
    console.error('Error fetching data:', error)
    throw error
  }
}

export default GetProperties
