import callApi from '@/services/api/apiCaller'
import endpoints from '@/enums/EndPointEnum'

export const AcceptInvite = async (contractId: string) => {
  try {
    return await callApi({
      method: 'POST',
      endpoint: endpoints.tenant.invite.accept(contractId)
    })
  } catch (error) {
    console.error('Error accepting invitation:', error)
    throw error
  }
}

export default AcceptInvite
