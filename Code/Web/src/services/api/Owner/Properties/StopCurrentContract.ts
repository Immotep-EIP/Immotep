import callApi from '@/services/api/apiCaller'

const StopCurrentContract = async (id: string) => {
  try {
    return await callApi({
      method: 'PUT',
      endpoint: `owner/properties/${id}/end-contract`,
    })
  } catch (error) {
    console.error('Error stopping current contract:', error)
    throw error
  }
}

export default StopCurrentContract
