import endpoints from '@/enums/EndPointEnum'
import callApi from '@/services/api/apiCaller'
import { User } from '@/interfaces/User/User'

const PutUserPicture = async (pictureData: string): Promise<User> => {
  try {
    return await callApi<User, { data: string }>({
      method: 'PUT',
      endpoint: endpoints.user.picture.update(),
      body: JSON.stringify({ data: pictureData })
    })
  } catch (error) {
    console.error('Error updating user picture:', error)
    throw error
  }
}

export default PutUserPicture
