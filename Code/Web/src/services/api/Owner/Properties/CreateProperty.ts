import callApi from '@/services/api/apiCaller'
import { CreateProperty, PropertyDetails } from '@/interfaces/Property/Property'
import endpoints from '@/enums/EndPointEnum'

const CreatePropertyFunction = async (
  data: CreateProperty
): Promise<PropertyDetails> => {
  try {
    return await callApi({
      method: 'POST',
      endpoint: endpoints.owner.properties.create,
      data
    })
  } catch (error) {
    console.error('Error fetching data:', error)
    throw error
  }
}

export default CreatePropertyFunction
