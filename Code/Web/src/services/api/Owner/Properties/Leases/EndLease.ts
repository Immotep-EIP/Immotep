import callApi from '@/services/api/apiCaller'
import endpoints from '@/enums/EndPointEnum'

const EndLease = async (propertyId: string): Promise<void> => {
  try {
    return await callApi({
      method: 'PUT',
      endpoint: endpoints.owner.properties.leases.end(propertyId)
    })
  } catch (error) {
    console.error('Error ending lease:', error)
    throw error
  }
}

export default EndLease
