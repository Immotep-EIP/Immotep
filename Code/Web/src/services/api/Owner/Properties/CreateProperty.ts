import callApi from '@/services/api/apiCaller'
import {
  CreatePropertyPayload,
  PropertyDetails
} from '@/interfaces/Property/Property'
import endpoints from '@/enums/EndPointEnum'

const CreatePropertyFunction = async (
  data: CreatePropertyPayload
): Promise<PropertyDetails> => {
  try {
    return await callApi<PropertyDetails, CreatePropertyPayload>({
      method: 'POST',
      endpoint: endpoints.owner.properties.create,
      body: data
    })
  } catch (error) {
    console.error('Error fetching data:', error)
    throw error
  }
}

export default CreatePropertyFunction
