import qs from 'qs'
import { TokenResponse, UserRegister, UserToken } from '@/interfaces/User/User'
import { register, loginApi } from '@/services/api/Authentification/AuthApi'
import callApi from '@/services/api/apiCaller'

jest.mock('@/services/api/apiCaller')

const mockedCallApi = callApi as jest.MockedFunction<typeof callApi>

describe('AuthApi', () => {
  beforeEach(() => {
    jest.clearAllMocks()
  })

  describe('register', () => {
    it('should call callApi with the correct parameters for registration without contractId', async () => {
      const userInfo: UserRegister = {
        email: 'test@example.com',
        password: 'password123',
        firstname: 'John',
        lastname: 'Doe',
        confirmPassword: 'password123'
      }

      const mockResponse = { data: 'success' }
      mockedCallApi.mockResolvedValueOnce(mockResponse)

      const result = await register(userInfo)

      expect(mockedCallApi).toHaveBeenCalledWith({
        method: 'POST',
        endpoint: 'auth/register/',
        data: userInfo
      })

      expect(result).toEqual(mockResponse)
    })

    it('should call callApi with the correct parameters for registration with contractId', async () => {
      const userInfo: UserRegister = {
        email: 'test@example.com',
        password: 'password123',
        firstname: 'John',
        lastname: 'Doe',
        contractId: '12345',
        confirmPassword: 'password123'
      }

      const mockResponse = { data: 'success' }
      mockedCallApi.mockResolvedValueOnce(mockResponse)

      const result = await register(userInfo)

      expect(mockedCallApi).toHaveBeenCalledWith({
        method: 'POST',
        endpoint: 'auth/invite/12345/',
        data: userInfo
      })

      expect(result).toEqual(mockResponse)
    })

    it('should handle errors during registration', async () => {
      const userInfo: UserRegister = {
        email: 'test@example.com',
        password: 'password123',
        firstname: 'John',
        lastname: 'Doe',
        confirmPassword: 'password123'
      }

      const mockError = new Error('Registration failed')
      mockedCallApi.mockRejectedValueOnce(mockError)

      const consoleErrorSpy = jest
        .spyOn(console, 'error')
        .mockImplementation(() => {})

      await expect(register(userInfo)).rejects.toThrow('Registration failed')

      expect(consoleErrorSpy).toHaveBeenCalledWith('request error:', mockError)

      consoleErrorSpy.mockRestore()
    })
  })

  describe('loginApi', () => {
    it('should call callApi with the correct parameters for login', async () => {
      const userInfo: UserToken = {
        grant_type: 'password',
        username: 'test@example.com',
        password: 'password123'
      }

      const mockResponse: TokenResponse = {
        access_token: 'access-token',
        refresh_token: 'refresh-token',
        expires_in: 3600
      }
      mockedCallApi.mockResolvedValueOnce(mockResponse)

      const result = await loginApi(userInfo)

      expect(mockedCallApi).toHaveBeenCalledWith({
        method: 'POST',
        endpoint: 'auth/token/',
        data: qs.stringify(userInfo),
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded'
        }
      })

      expect(result).toEqual(mockResponse)
    })

    it('should handle errors during login', async () => {
      const userInfo: UserToken = {
        grant_type: 'password',
        username: 'test@example.com',
        password: 'password123'
      }

      const mockError = new Error('Login failed')
      mockedCallApi.mockRejectedValueOnce(mockError)

      const consoleErrorSpy = jest
        .spyOn(console, 'error')
        .mockImplementation(() => {})

      await expect(loginApi(userInfo)).rejects.toThrow('Login failed')

      expect(consoleErrorSpy).toHaveBeenCalledWith('request error:', mockError)

      consoleErrorSpy.mockRestore()
    })
  })
})
