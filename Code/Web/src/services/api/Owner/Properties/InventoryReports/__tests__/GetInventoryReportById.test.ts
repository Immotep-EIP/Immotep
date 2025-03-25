import callApi from '@/services/api/apiCaller'
import GetInventoryReportById from '../GetInventoryReportById'

jest.mock('@/services/api/apiCaller')

describe('GetInventoryReportById', () => {
  const mockPropertyId = '123'
  const mockReportId = '456'
  const mockResponse = {
    id: mockReportId,
    title: 'Test Report',
    description: 'Test Description'
  }

  beforeEach(() => {
    jest.clearAllMocks()
  })

  it('should fetch inventory report by id successfully', async () => {
    ;(callApi as jest.Mock).mockResolvedValue(mockResponse)

    const result = await GetInventoryReportById(mockPropertyId, mockReportId)

    expect(callApi).toHaveBeenCalledWith({
      method: 'GET',
      endpoint: `owner/properties/${mockPropertyId}/inventory-reports/${mockReportId}/`
    })
    expect(result).toEqual(mockResponse)
  })

  it('should handle errors when fetching inventory report', async () => {
    const mockError = new Error('Failed to fetch')
    const consoleErrorSpy = jest
      .spyOn(console, 'error')
      .mockImplementation(() => {})
    ;(callApi as jest.Mock).mockRejectedValue(mockError)

    await expect(
      GetInventoryReportById(mockPropertyId, mockReportId)
    ).rejects.toThrow('Failed to fetch')
    expect(callApi).toHaveBeenCalledWith({
      method: 'GET',
      endpoint: `owner/properties/${mockPropertyId}/inventory-reports/${mockReportId}/`
    })

    consoleErrorSpy.mockRestore()
  })
})
