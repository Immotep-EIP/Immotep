import callApi from '@/services/api/apiCaller'
import { Lease } from '@/interfaces/Property/Lease/Lease'
import endpoints from '@/enums/EndPointEnum'

const GetLeaseById = async (
  propertyId: string,
  leaseId: string = 'current'
): Promise<Lease> => {
  try {
    return await callApi<Lease>({
      method: 'GET',
      endpoint: endpoints.owner.properties.leases.byId(propertyId, leaseId)
    })
  } catch (error) {
    console.error('Error fetching lease:', error)
    throw error
  }
}

export default GetLeaseById
