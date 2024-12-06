import callApi from '@/services/api/apiCaller'
import { GetProperty } from '@/interfaces/Property/Property'

const GetProperties = async () => {
  try {
    return await callApi<GetProperty[]>({
      method: 'GET',
      endpoint: 'owner/properties'
    })
  } catch (error) {
    console.error('Error fetching data:', error)
    throw error
  }
}

export default GetProperties
