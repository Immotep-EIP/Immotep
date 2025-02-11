import callApi from '@/services/api/apiCaller'

const ArchiveProperty = async (propertyId: string) => {
  try {
    return await callApi({
      method: 'DELETE',
      endpoint: `owner/properties/${propertyId}/`
    })
  } catch (error) {
    console.error('Error deleting property:', error)
    throw error
  }
}

export default ArchiveProperty
