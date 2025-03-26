import callApi from '@/services/api/apiCaller'
import GetPropertyPicture from '../GetPropertyPicture'

jest.mock('@/services/api/apiCaller')

describe('GetPropertyPicture', () => {
  const mockId = '123'
  const mockResponse = { picture_url: 'test-url' }

  beforeEach(() => {
    jest.clearAllMocks()
  })

  it('should fetch property picture successfully', async () => {
    ;(callApi as jest.Mock).mockResolvedValue(mockResponse)

    const result = await GetPropertyPicture(mockId)

    expect(callApi).toHaveBeenCalledWith({
      method: 'GET',
      endpoint: `owner/properties/${mockId}/picture/`
    })
    expect(result).toEqual(mockResponse)
  })

  it('should handle errors when fetching property picture', async () => {
    const mockError = new Error('Failed to fetch')
    const consoleErrorSpy = jest
      .spyOn(console, 'error')
      .mockImplementation(() => {})
    ;(callApi as jest.Mock).mockRejectedValue(mockError)

    await expect(GetPropertyPicture(mockId)).rejects.toThrow('Failed to fetch')
    expect(callApi).toHaveBeenCalledWith({
      method: 'GET',
      endpoint: `owner/properties/${mockId}/picture/`
    })

    consoleErrorSpy.mockRestore()
  })
})
