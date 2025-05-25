import callApi from '@/services/api/apiCaller'

import { Lease } from '@/interfaces/Property/Lease/Lease'
import endpoints from '@/enums/EndPointEnum'

const GetLeaseByTenant = async (): Promise<Lease[]> => {
  try {
    return await callApi<Lease[]>({
      method: 'GET',
      endpoint: endpoints.tenant.leases.list()
    })
  } catch (error) {
    console.error('Error fetching lease:', error)
    throw error
  }
}

export default GetLeaseByTenant
