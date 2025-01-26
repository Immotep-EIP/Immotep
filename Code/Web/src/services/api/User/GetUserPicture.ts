import callApi from '@/services/api/apiCaller'
import { UserPictureResponse } from '@/interfaces/User/User'

const GetUserPicture = async (id: string) => {
  try {
    return await callApi<UserPictureResponse>({
      method: 'GET',
      endpoint: `user/${id}/picture/`
    })
  } catch (error) {
    console.error('Error fetching data:', error)
    throw error
  }
}

export default GetUserPicture
