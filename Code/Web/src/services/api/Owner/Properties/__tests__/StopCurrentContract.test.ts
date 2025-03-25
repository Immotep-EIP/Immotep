import StopCurrentContract from '@/services/api/Owner/Properties/StopCurrentContract'
import callApi from '@/services/api/apiCaller'

jest.mock('@/services/api/apiCaller')

const mockedCallApi = callApi as jest.MockedFunction<typeof callApi>

describe('StopCurrentContract', () => {
  beforeEach(() => {
    jest.clearAllMocks()
  })

  it('should call callApi with the correct parameters and return the response', async () => {
    const mockResponse = {
      success: true,
      message: 'Contract stopped successfully'
    }
    mockedCallApi.mockResolvedValueOnce(mockResponse)

    const propertyId = '123'

    const result = await StopCurrentContract(propertyId)

    expect(mockedCallApi).toHaveBeenCalledWith({
      method: 'PUT',
      endpoint: `owner/properties/${propertyId}/end-contract/`
    })

    expect(result).toEqual(mockResponse)
  })

  it('should handle errors during contract termination', async () => {
    const mockError = new Error('API call failed')
    mockedCallApi.mockRejectedValueOnce(mockError)

    const propertyId = '123'

    const consoleErrorSpy = jest
      .spyOn(console, 'error')
      .mockImplementation(() => {})

    await expect(StopCurrentContract(propertyId)).rejects.toThrow(
      'API call failed'
    )

    expect(consoleErrorSpy).toHaveBeenCalledWith(
      'Error stopping current contract:',
      mockError
    )

    consoleErrorSpy.mockRestore()
  })
})
