import callApi from '@/services/api/apiCaller'

import endpoints from '@/enums/EndPointEnum'

export const AcceptInvite = async (leaseId: string): Promise<void> => {
  try {
    return await callApi({
      method: 'POST',
      endpoint: endpoints.tenant.invite.accept(leaseId)
    })
  } catch (error) {
    console.error('Error accepting invitation:', error)
    throw error
  }
}

export default AcceptInvite
