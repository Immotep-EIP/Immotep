import { UpdateUserInfos } from '@/services/api/User/UpdateUserInfos'
import callApi from '@/services/api/apiCaller'
import { User } from '@/interfaces/User/User'

jest.mock('@/services/api/apiCaller')

describe('UpdateUserInfos', () => {
  beforeEach(() => {
    jest.clearAllMocks()
  })

  it('should call callApi with the correct parameters and return the response', async () => {
    const consoleErrorSpy = jest
      .spyOn(console, 'error')
      .mockImplementation(() => {})

    const mockResponse: User = {
      id: '1',
      firstname: 'John',
      lastname: 'Doe',
      email: 'john.doe@example.com',
      role: 'user',
      created_at: new Date(),
      updated_at: new Date()
    }
    ;(callApi as jest.Mock).mockResolvedValue(mockResponse)

    const userData: User = {
      id: '1',
      firstname: 'John',
      lastname: 'Doe',
      email: 'john.doe@example.com',
      role: 'user',
      created_at: new Date(),
      updated_at: new Date()
    }

    const result = await UpdateUserInfos(userData)

    expect(callApi).toHaveBeenCalledWith({
      method: 'PUT',
      endpoint: 'profile/',
      body: JSON.stringify(userData)
    })

    expect(result).toEqual(mockResponse)

    consoleErrorSpy.mockRestore()
  })

  it('should throw an error if callApi fails', async () => {
    const consoleErrorSpy = jest
      .spyOn(console, 'error')
      .mockImplementation(() => {})

    const mockError = new Error('API call failed')
    ;(callApi as jest.Mock).mockRejectedValue(mockError)

    const userData: User = {
      id: '1',
      firstname: 'John',
      lastname: 'Doe',
      email: 'john.doe@example.com',
      role: 'user',
      created_at: new Date(),
      updated_at: new Date()
    }

    await expect(UpdateUserInfos(userData)).rejects.toThrow('API call failed')

    expect(callApi).toHaveBeenCalledWith({
      method: 'PUT',
      endpoint: 'profile/',
      body: JSON.stringify(userData)
    })

    consoleErrorSpy.mockRestore()
  })

  it('should log an error if callApi fails', async () => {
    const mockError = new Error('API call failed')
    ;(callApi as jest.Mock).mockRejectedValue(mockError)

    const consoleErrorSpy = jest
      .spyOn(console, 'error')
      .mockImplementation(() => {})

    const userData: User = {
      id: '1',
      firstname: 'John',
      lastname: 'Doe',
      email: 'john.doe@example.com',
      role: 'user',
      created_at: new Date(),
      updated_at: new Date()
    }

    await expect(UpdateUserInfos(userData)).rejects.toThrow('API call failed')

    expect(consoleErrorSpy).toHaveBeenCalledWith(
      'Error fetching data:',
      mockError
    )

    consoleErrorSpy.mockRestore()
  })
})
