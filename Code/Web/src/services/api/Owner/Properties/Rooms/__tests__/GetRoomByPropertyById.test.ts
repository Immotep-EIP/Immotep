import callApi from '@/services/api/apiCaller'
import GetRoomByPropertyById from '../GetRoomByPropertyById'

jest.mock('@/services/api/apiCaller')

describe('GetRoomByPropertyById', () => {
  const mockPropertyId = '123'
  const mockRoomId = '456'
  const mockResponse = { id: mockRoomId, name: 'Living Room' }

  beforeEach(() => {
    jest.clearAllMocks()
  })

  it('should fetch room by id successfully', async () => {
    ;(callApi as jest.Mock).mockResolvedValue(mockResponse)

    const result = await GetRoomByPropertyById(mockPropertyId, mockRoomId)

    expect(callApi).toHaveBeenCalledWith({
      method: 'GET',
      endpoint: `owner/properties/${mockPropertyId}/rooms/${mockRoomId}/`
    })
    expect(result).toEqual(mockResponse)
  })

  it('should handle errors when fetching room', async () => {
    const mockError = new Error('Failed to fetch')
    const consoleErrorSpy = jest
      .spyOn(console, 'error')
      .mockImplementation(() => {})
    ;(callApi as jest.Mock).mockRejectedValue(mockError)

    await expect(
      GetRoomByPropertyById(mockPropertyId, mockRoomId)
    ).rejects.toThrow('Failed to fetch')
    expect(callApi).toHaveBeenCalledWith({
      method: 'GET',
      endpoint: `owner/properties/${mockPropertyId}/rooms/${mockRoomId}/`
    })

    consoleErrorSpy.mockRestore()
  })
})
