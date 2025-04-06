import qs from 'qs'

import callApi from '@/services/api/apiCaller'
import { UserRegister, UserToken, TokenResponse } from '@/interfaces/User/User'
import endpoints from '@/enums/EndPointEnum'

export const register = async (userInfo: UserRegister) => {
  const endpoint = userInfo.leaseId
    ? endpoints.user.auth.invite(userInfo.leaseId)
    : endpoints.user.auth.register()

  try {
    const response = await callApi({
      method: 'POST',
      endpoint,
      body: userInfo
    })
    return response
  } catch (error) {
    console.error('request error:', error)
    throw error
  }
}

export const loginApi = async (userInfo: UserToken) => {
  try {
    const response = await callApi<UserToken, TokenResponse>({
      method: 'POST',
      endpoint: endpoints.user.auth.token(),
      body: qs.stringify(userInfo),
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded'
      }
    })
    return response
  } catch (error) {
    console.error('request error:', error)
    throw error
  }
}
