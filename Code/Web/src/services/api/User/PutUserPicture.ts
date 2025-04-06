import endpoints from '@/enums/EndPointEnum'
import callApi from '@/services/api/apiCaller'

const PutUserPicture = async (pictureData: string) => {
  try {
    return await callApi({
      method: 'PUT',
      endpoint: endpoints.user.picture.update(),
      data: JSON.stringify({ data: pictureData })
    })
  } catch (error) {
    console.error('Error updating user picture:', error)
    throw error
  }
}

export default PutUserPicture
