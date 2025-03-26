import UpdatePropertyPicture from '@/services/api/Owner/Properties/UpdatePropertyPicture'
import callApi from '@/services/api/apiCaller'

jest.mock('@/services/api/apiCaller')

const mockedCallApi = callApi as jest.MockedFunction<typeof callApi>

describe('UpdatePropertyPicture', () => {
  beforeEach(() => {
    jest.clearAllMocks()
  })

  it('should call callApi with the correct parameters and return the response', async () => {
    const mockResponse = {
      success: true,
      message: 'Picture updated successfully'
    }
    mockedCallApi.mockResolvedValueOnce(mockResponse)

    const propertyId = '123'
    const pictureData = 'base64-encoded-image-data'

    const result = await UpdatePropertyPicture(propertyId, pictureData)

    expect(mockedCallApi).toHaveBeenCalledWith({
      method: 'PUT',
      endpoint: `owner/properties/${propertyId}/picture/`,
      data: JSON.stringify({ data: pictureData })
    })

    expect(result).toEqual(mockResponse)
  })

  it('should handle errors during picture update', async () => {
    const mockError = new Error('API call failed')
    mockedCallApi.mockRejectedValueOnce(mockError)

    const propertyId = '123'
    const pictureData = 'base64-encoded-image-data'

    const consoleErrorSpy = jest
      .spyOn(console, 'error')
      .mockImplementation(() => {})

    await expect(
      UpdatePropertyPicture(propertyId, pictureData)
    ).rejects.toThrow('API call failed')

    expect(consoleErrorSpy).toHaveBeenCalledWith(
      'Error updating property picture:',
      mockError
    )

    consoleErrorSpy.mockRestore()
  })
})
