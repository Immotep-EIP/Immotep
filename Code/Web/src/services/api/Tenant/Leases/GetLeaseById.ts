import callApi from '@/services/api/apiCaller'
import { Lease } from '@/interfaces/Lease/Lease'

const GetLeaseById = async (leaseId: string = 'current') => {
  try {
    return await callApi<Lease>({
      method: 'GET',
      endpoint: `tenant/leases/${leaseId}/`
    })
  } catch (error) {
    console.error('Error fetching lease:', error)
    throw error
  }
}

export default GetLeaseById
