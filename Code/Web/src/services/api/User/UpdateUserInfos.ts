import callApi from '@/services/api/apiCaller'
import { User } from '@/interfaces/User/User'

interface UpdateUserInfosProps {
  firstname: string
  lastname: string
}

export const UpdateUserInfos = async (data: UpdateUserInfosProps) => {
  try {
    const response = await callApi<User>({
      method: 'PUT',
      endpoint: 'profile',
      data: JSON.stringify(data),
    })
    return response
  } catch (error) {
    console.error('Error fetching data:', error)
    throw error
  }
}

export default UpdateUserInfos
