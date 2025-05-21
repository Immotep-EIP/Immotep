import callApi from '@/services/api/apiCaller'

import { Document } from '@/interfaces/Property/Document'
import endpoints from '@/enums/EndPointEnum'

const UploadDocument = async (
  payload: string,
  propertyId: string,
  leaseId: string = endpoints.owner.properties.leases.current
) => {
  try {
    return callApi<Document>({
      method: 'POST',
      endpoint: endpoints.owner.properties.leases.documents(
        propertyId,
        leaseId
      ),
      body: payload
    })
  } catch (error) {
    console.error('Error uploading document:', error)
    throw error
  }
}

export default UploadDocument
