import callApi from '@/services/api/apiCaller'

export const AcceptInvite = async (contractId: string) => {
  try {
    return await callApi({
      method: 'POST',
      endpoint: `tenant/invite/${contractId}/`
    })
  } catch (error) {
    console.error('Error accepting invitation:', error)
    throw error
  }
}

export default AcceptInvite
