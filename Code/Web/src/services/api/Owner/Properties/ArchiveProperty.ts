import callApi from '@/services/api/apiCaller'
import endpoints from '@/enums/EndPointEnum'
import { PropertyDetails } from '@/interfaces/Property/Property'

const ArchiveProperty = async (
  propertyId: string
): Promise<PropertyDetails> => {
  try {
    return await callApi<PropertyDetails, { archive: boolean }>({
      method: 'PUT',
      endpoint: endpoints.owner.properties.archive(propertyId),
      body: {
        archive: true
      }
    })
  } catch (error) {
    console.error('Error deleting property:', error)
    throw error
  }
}

export default ArchiveProperty
