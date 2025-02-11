import PutUserPicture from '@/services/api/User/PutUserPicture'
import callApi from '@/services/api/apiCaller'

// Mock the callApi function
jest.mock('@/services/api/apiCaller')

describe('PutUserPicture', () => {
  beforeEach(() => {
    jest.clearAllMocks()
  })

  it('should call callApi with the correct parameters and return the response', async () => {
    const mockResponse = { success: true }
    ;(callApi as jest.Mock).mockResolvedValue(mockResponse)

    const consoleErrorSpy = jest
      .spyOn(console, 'error')
      .mockImplementation(() => {})

    const pictureData = 'base64-encoded-image-data'

    const result = await PutUserPicture(pictureData)

    expect(callApi).toHaveBeenCalledWith({
      method: 'PUT',
      endpoint: 'profile/picture/',
      data: JSON.stringify({ data: pictureData })
    })

    expect(result).toEqual(mockResponse)

    consoleErrorSpy.mockRestore()
  })

  it('should throw an error if callApi fails', async () => {
    const mockError = new Error('API call failed')
    ;(callApi as jest.Mock).mockRejectedValue(mockError)

    const consoleErrorSpy = jest
      .spyOn(console, 'error')
      .mockImplementation(() => {})

    const pictureData = 'base64-encoded-image-data'

    await expect(PutUserPicture(pictureData)).rejects.toThrow('API call failed')

    expect(callApi).toHaveBeenCalledWith({
      method: 'PUT',
      endpoint: 'profile/picture/',
      data: JSON.stringify({ data: pictureData })
    })

    consoleErrorSpy.mockRestore()
  })

  it('should log an error if callApi fails', async () => {
    const mockError = new Error('API call failed')
    ;(callApi as jest.Mock).mockRejectedValue(mockError)

    const consoleErrorSpy = jest
      .spyOn(console, 'error')
      .mockImplementation(() => {})

    const pictureData = 'base64-encoded-image-data'

    await expect(PutUserPicture(pictureData)).rejects.toThrow('API call failed')

    expect(consoleErrorSpy).toHaveBeenCalledWith(
      'Error updating property picture:',
      mockError
    )

    consoleErrorSpy.mockRestore()
  })
})
