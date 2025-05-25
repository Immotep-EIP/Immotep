import callApi from '@/services/api/apiCaller'

import {
  CreatePropertyPayload,
  PropertyDetails
} from '@/interfaces/Property/Property'
import endpoints from '@/enums/EndPointEnum'

const UpdatePropertyFunction = async (
  data: CreatePropertyPayload,
  id: string
): Promise<PropertyDetails> => {
  try {
    return await callApi<PropertyDetails, CreatePropertyPayload>({
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
