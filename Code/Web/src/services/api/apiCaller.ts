import axios from 'axios'

import { loginApi } from '@/services/api/Authentification/AuthApi'
import { saveData, deleteData } from '@/utils/cache/localStorage'
import AuthEnum from '@/enums/AuthEnum'
import NavigationEnum from '@/enums/NavigationEnum'
import { ApiCallerParams } from '@/interfaces/Api/callApi'

const API_BASE_URL = `${process.env.VITE_API_URL}` || 'http://localhost:3001/v1'

const api = axios.create({
  baseURL: API_BASE_URL
})

api.interceptors.request.use(
  config => {
    const modifiedConfig = { ...config }

    const accessToken =
      localStorage.getItem('access_token') ||
      sessionStorage.getItem('access_token')
    if (accessToken) {
      modifiedConfig.headers = {
        ...modifiedConfig.headers,
        Authorization: `Bearer ${accessToken}`
      }
    }
    return modifiedConfig
  },
  error => Promise.reject(error)
)

api.interceptors.response.use(
  response => response,
  async error => {
    const originalRequest = error.config

    if (
      error.response &&
      error.response.status === 401 &&
      // eslint-disable-next-line no-underscore-dangle
      !originalRequest._retry
    ) {
      // eslint-disable-next-line no-underscore-dangle
      originalRequest._retry = true
      const refreshToken =
        localStorage.getItem('refresh_token') ||
        sessionStorage.getItem('refresh_token')
      const tokenExpiry =
        localStorage.getItem('expires_in') ||
        sessionStorage.getItem('expires_in')

      const now = Date.now()
      if (refreshToken && tokenExpiry) {
        const expiryTime = parseInt(tokenExpiry, 10) + AuthEnum.SECURETIMER
        if (now < expiryTime) {
          try {
            const response = await loginApi({
              grant_type: 'refresh_token',
              refresh_token: refreshToken
            })

            const newAccessToken = response.access_token
            const newRefreshToken = response.refresh_token
            const newExpiresIn = response.expires_in

            deleteData()
            saveData(newAccessToken, newRefreshToken, newExpiresIn)

            originalRequest.headers.Authorization = `Bearer ${newAccessToken}`

            return api(originalRequest)
          } catch (refreshError) {
            console.error('Refresh token error:', refreshError)
            deleteData()
            window.location.href = NavigationEnum.LOGIN
            return Promise.reject(refreshError)
          }
        }
        deleteData()
        window.location.href = NavigationEnum.LOGIN
        return Promise.reject(error)
      }
      deleteData()
      return Promise.reject(error)
    }
    return Promise.reject(error)
  }
)

const callApi = async <TResponse, TBody = unknown>({
  method,
  endpoint,
  body,
  headers,
  params
}: ApiCallerParams<TResponse, TBody>): Promise<TResponse> => {
  try {
    const response = await api.request({
      method,
      url: `/${endpoint}`,
      data: body,
      params,
      headers: {
        Accept: 'application/json',
        ...headers
      }
    })

    return response.data as TResponse
  } catch (error) {
    console.error(`Error during API call to ${endpoint}:`, error)
    throw error
  }
}

export default callApi
