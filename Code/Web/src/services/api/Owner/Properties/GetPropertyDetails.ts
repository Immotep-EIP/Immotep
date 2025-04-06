import callApi from '@/services/api/apiCaller'
import { PropertyDetails } from '@/interfaces/Property/Property'
import endpoints from '@/enums/EndPointEnum'

const GetPropertyDetails = async (id: string) => {
  try {
    return await callApi<PropertyDetails>({
      method: 'GET',
      endpoint: endpoints.owner.properties.details(id)
    })
  } catch (error) {
    console.error('Error fetching data:', error)
    throw error
  }
}

export default GetPropertyDetails
