import callApi from '@/services/api/apiCaller'
import { User } from '@/interfaces/User/User'

export const PutUserPicture = async (data: string) => {
  try {
    const response = await callApi<User>({
      method: 'PUT',
      endpoint: 'profile/picture/',
      data
    })
    return response
  } catch (error) {
    console.error('Error fetching data:', error)
    throw error
  }
}

export default PutUserPicture
