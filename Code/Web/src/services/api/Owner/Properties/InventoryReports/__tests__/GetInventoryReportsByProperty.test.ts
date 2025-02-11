import callApi from '@/services/api/apiCaller'
import GetInventoryReportsByProperty from '../GetInventoryReportsByProperty'

jest.mock('@/services/api/apiCaller')

describe('GetInventoryReportsByProperty', () => {
  const mockPropertyId = '123'
  const mockResponse = [
    {
      id: '1',
      title: 'Test Report 1',
      description: 'Test Description 1'
    },
    {
      id: '2',
      title: 'Test Report 2',
      description: 'Test Description 2'
    }
  ]

  beforeEach(() => {
    jest.clearAllMocks()
  })

  it('should fetch inventory reports successfully', async () => {
    ;(callApi as jest.Mock).mockResolvedValue(mockResponse)

    const result = await GetInventoryReportsByProperty(mockPropertyId)

    expect(callApi).toHaveBeenCalledWith({
      method: 'GET',
      endpoint: `owner/properties/${mockPropertyId}/inventory-reports/`
    })
    expect(result).toEqual(mockResponse)
  })

  it('should handle errors when fetching inventory reports', async () => {
    const mockError = new Error('Failed to fetch')
    const consoleErrorSpy = jest
      .spyOn(console, 'error')
      .mockImplementation(() => {})
    ;(callApi as jest.Mock).mockRejectedValue(mockError)

    await expect(GetInventoryReportsByProperty(mockPropertyId)).rejects.toThrow(
      'Failed to fetch'
    )
    expect(callApi).toHaveBeenCalledWith({
      method: 'GET',
      endpoint: `owner/properties/${mockPropertyId}/inventory-reports/`
    })

    consoleErrorSpy.mockRestore()
  })
})
