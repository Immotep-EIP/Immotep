import callApi from '@/services/api/apiCaller'

import endpoints from '@/enums/EndPointEnum'

const CancelTenantInvitation = async (propertyId: string): Promise<void> => {
  try {
    return await callApi({
      method: 'DELETE',
      endpoint: endpoints.owner.properties.tenant.cancelInvite(propertyId)
    })
  } catch (error) {
    console.error('Error cancelling invitation to tenant:', error)
    throw error
  }
}

export default CancelTenantInvitation
