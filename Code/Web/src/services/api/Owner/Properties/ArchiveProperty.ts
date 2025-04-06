import callApi from '@/services/api/apiCaller'
import endpoints from '@/enums/EndPointEnum'

const ArchiveProperty = async (propertyId: string) => {
  try {
    return await callApi({
      method: 'PUT',
      endpoint: endpoints.owner.properties.archive(propertyId),
      data: {
        archive: true
      }
    })
  } catch (error) {
    console.error('Error deleting property:', error)
    throw error
  }
}

export default ArchiveProperty
