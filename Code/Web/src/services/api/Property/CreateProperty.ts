import callApi from '@/services/api/apiCaller'
import { CreateProperty } from '@/interfaces/Property/Property'

const CreatePropertyFunction = async (data: CreateProperty) => {
  try {
    return await callApi({
      method: 'POST',
      endpoint: 'owner/properties',
      data
    })
  } catch (error) {
    console.error('Error fetching data:', error)
    throw error
  }
}

export default CreatePropertyFunction
