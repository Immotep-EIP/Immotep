import callApi from '@/services/api/apiCaller'
import { Lease } from '@/interfaces/Lease/Lease'
import endpoints from '@/enums/EndPointEnum'

const GetLeaseById = async (leaseId: string = 'current') => {
  try {
    return await callApi<Lease>({
      method: 'GET',
      endpoint: endpoints.tenant.leases.byId(leaseId)
    })
  } catch (error) {
    console.error('Error fetching lease:', error)
    throw error
  }
}

export default GetLeaseById
