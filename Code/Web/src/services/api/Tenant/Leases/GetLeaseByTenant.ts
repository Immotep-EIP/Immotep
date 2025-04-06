import callApi from '@/services/api/apiCaller'
import { Lease } from '@/interfaces/Lease/Lease'
import endpoints from '@/enums/EndPointEnum'

const GetLeaseByTenant = async () => {
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
