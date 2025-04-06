import callApi from '@/services/api/apiCaller'
import { EndLeaseResponse } from '@/interfaces/Lease/Lease'
import endpoints from '@/enums/EndPointEnum'

const EndLease = async (propertyId: string): Promise<EndLeaseResponse> => {
  try {
    return await callApi<EndLeaseResponse>({
      method: 'PUT',
      endpoint: endpoints.owner.properties.leases.end(propertyId)
    })
  } catch (error) {
    console.error('Error ending lease:', error)
    throw error
  }
}

export default EndLease
