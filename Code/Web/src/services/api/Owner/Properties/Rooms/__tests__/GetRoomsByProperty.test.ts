import callApi from '@/services/api/apiCaller'
import GetRoomsByProperty from '../GetRoomsByProperty'

jest.mock('@/services/api/apiCaller')

describe('GetRoomsByProperty', () => {
  const mockPropertyId = '123'
  const mockResponse = [
    { id: '1', name: 'Living Room' },
    { id: '2', name: 'Bedroom' }
  ]

  beforeEach(() => {
    jest.clearAllMocks()
  })

  it('should fetch rooms successfully', async () => {
    ;(callApi as jest.Mock).mockResolvedValue(mockResponse)

    const result = await GetRoomsByProperty({ propertyId: mockPropertyId })

    expect(callApi).toHaveBeenCalledWith({
      method: 'GET',
      endpoint: `owner/properties/${mockPropertyId}/rooms/`,
      params: {
        archive: false
      }
    })
    expect(result).toEqual(mockResponse)
  })

  it('should handle errors when fetching rooms', async () => {
    const mockError = new Error('Failed to fetch')
    const consoleErrorSpy = jest
      .spyOn(console, 'error')
      .mockImplementation(() => {})
    ;(callApi as jest.Mock).mockRejectedValue(mockError)

    await expect(
      GetRoomsByProperty({ propertyId: mockPropertyId })
    ).rejects.toThrow('Failed to fetch')
    expect(callApi).toHaveBeenCalledWith({
      method: 'GET',
      endpoint: `owner/properties/${mockPropertyId}/rooms/`,
      params: {
        archive: false
      }
    })

    consoleErrorSpy.mockRestore()
  })
})
