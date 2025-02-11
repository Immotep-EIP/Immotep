import callApi from '@/services/api/apiCaller'
import GetFurnitureByRoomById from '../GetFurnitureByRoomById'

jest.mock('@/services/api/apiCaller')

describe('GetFurnitureByRoomById', () => {
  const mockPropertyId = '123'
  const mockRoomId = '456'
  const mockFurnitureId = '789'
  const mockResponse = {
    id: mockFurnitureId,
    name: 'Sofa',
    quantity: 1,
    property_id: mockPropertyId,
    room_id: mockRoomId
  }

  beforeEach(() => {
    jest.clearAllMocks()
  })

  it('should fetch furniture by id successfully', async () => {
    ;(callApi as jest.Mock).mockResolvedValue(mockResponse)

    const result = await GetFurnitureByRoomById(
      mockPropertyId,
      mockRoomId,
      mockFurnitureId
    )

    expect(callApi).toHaveBeenCalledWith({
      method: 'GET',
      endpoint: `owner/properties/${mockPropertyId}/rooms/${mockRoomId}/furnitures/${mockFurnitureId}/`
    })
    expect(result).toEqual(mockResponse)
  })

  it('should handle errors when fetching furniture', async () => {
    const mockError = new Error('Failed to fetch')
    const consoleErrorSpy = jest
      .spyOn(console, 'error')
      .mockImplementation(() => {})
    ;(callApi as jest.Mock).mockRejectedValue(mockError)

    await expect(
      GetFurnitureByRoomById(mockPropertyId, mockRoomId, mockFurnitureId)
    ).rejects.toThrow('Failed to fetch')

    expect(callApi).toHaveBeenCalledWith({
      method: 'GET',
      endpoint: `owner/properties/${mockPropertyId}/rooms/${mockRoomId}/furnitures/${mockFurnitureId}/`
    })

    consoleErrorSpy.mockRestore()
  })
})
