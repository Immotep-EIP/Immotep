import callApi from '@/services/api/apiCaller'
import GetUserPicture from '../GetUserPicture'

jest.mock('@/services/api/apiCaller')

describe('GetUserPicture', () => {
  const mockUserId = '123'
  const mockResponse = { picture_url: 'http://example.com/picture.jpg' }

  beforeEach(() => {
    jest.clearAllMocks()
  })

  it('should fetch user picture successfully', async () => {
    ;(callApi as jest.Mock).mockResolvedValue(mockResponse)

    const result = await GetUserPicture(mockUserId)

    expect(callApi).toHaveBeenCalledWith({
      method: 'GET',
      endpoint: `user/${mockUserId}/picture/`
    })
    expect(result).toEqual(mockResponse)
  })

  it('should handle errors when fetching user picture', async () => {
    const mockError = new Error('Failed to fetch')
    const consoleErrorSpy = jest
      .spyOn(console, 'error')
      .mockImplementation(() => {})
    ;(callApi as jest.Mock).mockRejectedValue(mockError)

    await expect(GetUserPicture(mockUserId)).rejects.toThrow('Failed to fetch')

    expect(callApi).toHaveBeenCalledWith({
      method: 'GET',
      endpoint: `user/${mockUserId}/picture/`
    })

    consoleErrorSpy.mockRestore()
  })
})
