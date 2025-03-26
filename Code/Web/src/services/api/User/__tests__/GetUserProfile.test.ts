import callApi from '@/services/api/apiCaller'
import getUserProfile from '../GetUserProfile'

jest.mock('@/services/api/apiCaller')

describe('getUserProfile', () => {
  const mockResponse = {
    id: '123',
    email: 'test@example.com',
    first_name: 'John',
    last_name: 'Doe'
  }

  beforeEach(() => {
    jest.clearAllMocks()
  })

  it('should fetch user profile successfully', async () => {
    ;(callApi as jest.Mock).mockResolvedValue(mockResponse)

    const result = await getUserProfile()

    expect(callApi).toHaveBeenCalledWith({
      method: 'GET',
      endpoint: 'profile/'
    })
    expect(result).toEqual(mockResponse)
  })

  it('should handle errors when fetching user profile', async () => {
    const mockError = new Error('Failed to fetch')
    const consoleErrorSpy = jest
      .spyOn(console, 'error')
      .mockImplementation(() => {})
    ;(callApi as jest.Mock).mockRejectedValue(mockError)

    await expect(getUserProfile()).rejects.toThrow('Failed to fetch')

    expect(callApi).toHaveBeenCalledWith({
      method: 'GET',
      endpoint: 'profile/'
    })

    consoleErrorSpy.mockRestore()
  })
})
