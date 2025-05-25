import callApi from '@/services/api/apiCaller'

import endpoints from '@/enums/EndPointEnum'
import { PropertyDetails } from '@/interfaces/Property/Property'

const UnarchiveProperty = async (
  propertyId: string
): Promise<PropertyDetails> => {
  try {
    return await callApi<PropertyDetails, { archive: boolean }>({
      method: 'PUT',
      endpoint: endpoints.owner.properties.archive(propertyId),
      body: {
        archive: false
      }
    })
  } catch (error) {
    console.error('Error deleting property:', error)
    throw error
  }
}

export default UnarchiveProperty
