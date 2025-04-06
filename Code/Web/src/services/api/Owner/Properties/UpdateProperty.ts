import callApi from '@/services/api/apiCaller'
import { CreateProperty, PropertyDetails } from '@/interfaces/Property/Property'
import endpoints from '@/enums/EndPointEnum'

const UpdatePropertyFunction = async (
  data: CreateProperty,
  id: string
): Promise<PropertyDetails> => {
  try {
    return await callApi({
      method: 'PUT',
      endpoint: endpoints.owner.properties.update(id),
      body: data
    })
  } catch (error) {
    console.error('Error fetching data:', error)
    throw error
  }
}

export default UpdatePropertyFunction
