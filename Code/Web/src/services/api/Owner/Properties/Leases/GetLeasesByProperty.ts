import callApi from '@/services/api/apiCaller'
import { Lease } from '@/interfaces/Lease/Lease'

const GetLeasesByProperty = async (propertyId: string) => {
  try {
    return await callApi<Lease[]>({
      method: 'GET',
      endpoint: `owner/properties/${propertyId}/leases/`
    })
  } catch (error) {
    console.error('Error fetching leases:', error)
    throw error
  }
}

export default GetLeasesByProperty
