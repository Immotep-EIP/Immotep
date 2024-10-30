import callApi from '@/services/api/apiCaller'
import { User } from '@/interfaces/User/User'

export const getUserProfile = async () => {
  try {
    const response = await callApi<User>({
      method: 'GET',
      endpoint: 'profile'
    })
    return response
  } catch (error) {
    console.error('Error fetching data:', error)
    throw error
  }
}

export default getUserProfile
