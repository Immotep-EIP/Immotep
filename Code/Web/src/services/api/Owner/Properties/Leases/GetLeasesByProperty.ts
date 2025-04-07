import callApi from '@/services/api/apiCaller'
import { Lease } from '@/interfaces/Lease/Lease'
import endpoints from '@/enums/EndPointEnum'

const GetLeasesByProperty = async (propertyId: string): Promise<Lease[]> => {
  try {
    return await callApi<Lease[]>({
      method: 'GET',
      endpoint: endpoints.owner.properties.leases.list(propertyId)
    })
  } catch (error) {
    console.error('Error fetching leases:', error)
    throw error
  }
}

export default GetLeasesByProperty
