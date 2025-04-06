import callApi from '@/services/api/apiCaller'

const EndLease = async (propertyId: string) => {
  try {
    return await callApi({
      method: 'PUT',
      endpoint: `owner/properties/${propertyId}/leases/current/end/`
    })
  } catch (error) {
    console.error('Error ending lease:', error)
    throw error
  }
}

export default EndLease
