import callApi from '@/services/api/apiCaller'
import { PropertyDetails } from '@/interfaces/Property/Property'
import endpoints from '@/enums/EndPointEnum'

const GetProperties = async (
  archive: boolean = false
): Promise<PropertyDetails[]> => {
  try {
    return await callApi<PropertyDetails[]>({
      method: 'GET',
      endpoint: endpoints.owner.properties.list,
      params: {
        archive
      }
    })
  } catch (error) {
    console.error('Error fetching data:', error)
    throw error
  }
}

export default GetProperties
