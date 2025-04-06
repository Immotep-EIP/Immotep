import callApi from '@/services/api/apiCaller'
import { Document } from '@/interfaces/Property/Document'

const GetPropertyDocuments = async (
  propertyId: string,
  leaseId: string = 'current'
) => {
  try {
    return await callApi<Document[]>({
      method: 'GET',
      endpoint: `owner/properties/${propertyId}/leases/${leaseId}/docs/`
    })
  } catch (error) {
    console.error('Error fetching documents:', error)
    throw error
  }
}

export default GetPropertyDocuments
