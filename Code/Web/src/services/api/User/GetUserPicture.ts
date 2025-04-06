import callApi from '@/services/api/apiCaller'
import { UserPictureResponse } from '@/interfaces/User/User'
import endpoints from '@/enums/EndPointEnum'

const GetUserPicture = async (id: string) => {
  try {
    return await callApi<UserPictureResponse>({
      method: 'GET',
      endpoint: endpoints.user.picture.get(id)
    })
  } catch (error) {
    console.error('Error fetching data:', error)
    throw error
  }
}

export default GetUserPicture
