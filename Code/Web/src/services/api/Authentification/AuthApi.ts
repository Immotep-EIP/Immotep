import qs from 'qs'

import callApi from '@/services/api/apiCaller'
import { UserRegister, UserToken, TokenResponse } from '@/interfaces/User/User'

export const register = async (userInfo: UserRegister) => {
  const endpoint = userInfo.contractId
      ? `auth/invite/${userInfo.contractId}`
      : 'auth/register';

  try {
    const response = await callApi({
      method: 'POST',
      endpoint,
      data: userInfo
    })
    return response
  } catch (error) {
    console.error('request error:', error)
    throw error
  }
}

export const loginApi = async (userInfo: UserToken) => {
  try {
    const response = await callApi< UserToken, TokenResponse>({
      method: 'POST',
      endpoint: 'auth/token',
      data: qs.stringify(userInfo),
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
