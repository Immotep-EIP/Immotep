import callApi from '@/services/api/apiCaller'
import { Document } from '@/interfaces/Property/Document'
import endpoints from '@/enums/EndPointEnum'

const GetPropertyDocuments = async (
  propertyId: string,
  leaseId: string = 'current'
) => {
  try {
    return await callApi<Document[]>({
      method: 'GET',
      endpoint: endpoints.owner.properties.leases.documents(propertyId, leaseId)
    })
  } catch (error) {
    console.error('Error fetching documents:', error)
    throw error
  }
}

export default GetPropertyDocuments
