import callApi from '@/services/api/apiCaller'
import { Lease } from '@/interfaces/Lease/Lease'

const GetLeaseByTenant = async () => {
  try {
    return await callApi<Lease[]>({
      method: 'GET',
      endpoint: `tenant/leases/`
    })
  } catch (error) {
    console.error('Error fetching lease:', error)
    throw error
  }
}

export default GetLeaseByTenant
