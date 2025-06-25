import callApi from '@/services/api/apiCaller'

import endpoints from '@/enums/EndPointEnum'

const DeleteDocument = async (
  propertyId: string,
  documentId: string,
  leaseId: string = 'current'
) => {
  try {
    return await callApi<void>({
      method: 'DELETE',
      endpoint: endpoints.owner.properties.leases.deleteDocument(
        propertyId,
        leaseId,
        documentId
      )
    })
  } catch (error) {
    console.error('Error deleting document:', error)
    throw error
  }
}

export default DeleteDocument
