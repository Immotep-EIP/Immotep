import callApi from '@/services/api/apiCaller'
import { Lease } from '@/interfaces/Lease/Lease'

const GetLeaseById = async (
  propertyId: string,
  leaseId: string = 'current'
) => {
  try {
    return await callApi<Lease>({
      method: 'GET',
      endpoint: `owner/properties/${propertyId}/leases/${leaseId}/`
    })
  } catch (error) {
    console.error('Error fetching lease:', error)
    throw error
  }
}

export default GetLeaseById
