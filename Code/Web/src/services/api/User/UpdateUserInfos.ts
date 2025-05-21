import callApi from '@/services/api/apiCaller'
import { UpdateUserInfoPayload, User } from '@/interfaces/User/User'
import endpoints from '@/enums/EndPointEnum'

export const UpdateUserInfos = async (
  data: UpdateUserInfoPayload
): Promise<User> => {
  try {
    return await callApi<User, UpdateUserInfoPayload>({
      method: 'PUT',
      endpoint: endpoints.user.profile.get(),
      body: JSON.stringify(data)
    })
  } catch (error) {
    console.error('Error fetching data:', error)
    throw error
  }
}

export default UpdateUserInfos
