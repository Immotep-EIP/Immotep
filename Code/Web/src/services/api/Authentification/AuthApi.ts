import qs from 'qs'

import callApi from '@/services/api/apiCaller'
import {
  UserRegisterPayload,
  UserTokenPayload,
  TokenResponse,
  User
} from '@/interfaces/User/User'
import endpoints from '@/enums/EndPointEnum'

export const register = async (
  userInfo: UserRegisterPayload
): Promise<User> => {
  const endpoint = userInfo.leaseId
    ? endpoints.user.auth.invite(userInfo.leaseId)
    : endpoints.user.auth.register()

  try {
    return await callApi<User, UserRegisterPayload>({
      method: 'POST',
      endpoint,
      body: userInfo
    })
  } catch (error) {
    console.error('request error:', error)
    throw error
  }
}

export const loginApi = async (
  userInfo: UserTokenPayload
): Promise<TokenResponse> => {
  try {
    return await callApi<TokenResponse, UserTokenPayload>({
      method: 'POST',
      endpoint: endpoints.user.auth.token(),
      body: qs.stringify(userInfo),
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded'
      }
    })
  } catch (error) {
    console.error('request error:', error)
    throw error
  }
}
