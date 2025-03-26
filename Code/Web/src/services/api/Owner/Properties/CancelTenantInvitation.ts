import callApi from '@/services/api/apiCaller'

const CancelTenantInvitation = async (propertyId: string) => {
  try {
    return await callApi({
      method: 'DELETE',
      endpoint: `owner/properties/${propertyId}/cancel-invite/`
    })
  } catch (error) {
    console.error('Error cancelling invitation to tenant:', error)
    throw error
  }
}

export default CancelTenantInvitation
