import callApi from '@/services/api/apiCaller'
import CreateInventoryReportByProperty from '../CreateInventoryReportByProperty'

jest.mock('@/services/api/apiCaller')

describe('CreateInventoryReportByProperty', () => {
  const mockPropertyId = '123'
  const mockData = {
    title: 'Test Report',
    description: 'Test Description',
    rooms: [],
    type: 'entry'
  }
  const mockResponse = { id: '1', ...mockData }

  beforeEach(() => {
    jest.clearAllMocks()
  })

  it('should create inventory report successfully', async () => {
    ;(callApi as jest.Mock).mockResolvedValue(mockResponse)

    const result = await CreateInventoryReportByProperty(
      mockPropertyId,
      mockData
    )

    expect(callApi).toHaveBeenCalledWith({
      method: 'POST',
      endpoint: `owner/properties/${mockPropertyId}/inventory-reports/`,
      data: mockData
    })
    expect(result).toEqual(mockResponse)
  })

  it('should handle errors when creating inventory report', async () => {
    const mockError = new Error('Failed to create')
    const consoleErrorSpy = jest
      .spyOn(console, 'error')
      .mockImplementation(() => {})
    ;(callApi as jest.Mock).mockRejectedValue(mockError)

    await expect(
      CreateInventoryReportByProperty(mockPropertyId, mockData)
    ).rejects.toThrow('Failed to create')
    expect(callApi).toHaveBeenCalledWith({
      method: 'POST',
      endpoint: `owner/properties/${mockPropertyId}/inventory-reports/`,
      data: mockData
    })

    consoleErrorSpy.mockRestore()
  })
})
