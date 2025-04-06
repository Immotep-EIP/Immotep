import callApi from '@/services/api/apiCaller'
import EndLease from '../EndLease'

jest.mock('@/services/api/apiCaller')

const mockedCallApi = callApi as jest.MockedFunction<typeof callApi>

describe('EndLease', () => {
  beforeEach(() => {
    jest.clearAllMocks()
  })

  it('should call callApi with the correct parameters and return the response', async () => {
    const mockResponse = {
      success: true,
      message: 'Lease ended'
    }
    mockedCallApi.mockResolvedValueOnce(mockResponse)

    const propertyId = '123'

    const result = await EndLease(propertyId)

    expect(mockedCallApi).toHaveBeenCalledWith({
      method: 'PUT',
      endpoint: `owner/properties/${propertyId}/leases/current/end/`
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

    await expect(EndLease(propertyId)).rejects.toThrow('API call failed')

    expect(consoleErrorSpy).toHaveBeenCalledWith(
      'Error ending lease:',
      mockError
    )

    consoleErrorSpy.mockRestore()
  })
})
