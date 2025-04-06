import callApi from '@/services/api/apiCaller'
import { User } from '@/interfaces/User/User'
import endpoints from '@/enums/EndPointEnum'

export const getUsers = async () => {
  try {
    const response = await callApi<User>({
      method: 'GET',
      endpoint: endpoints.user.list()
    })
    return response
  } catch (error) {
    console.error('Error fetching data:', error)
    throw error
  }
}

export default getUsers
