import callApi from '@/services/api/apiCaller'
import GetFurnituresByRoom from '../GetFurnituresByRoom'

jest.mock('@/services/api/apiCaller')

describe('GetFurnituresByRoom', () => {
  const mockPropertyId = '123'
  const mockRoomId = '456'
  const mockResponse = [
    {
      id: '1',
      name: 'Sofa',
      quantity: 1,
      property_id: mockPropertyId,
      room_id: mockRoomId
    },
    {
      id: '2',
      name: 'Chair',
      quantity: 4,
      property_id: mockPropertyId,
      room_id: mockRoomId
    }
  ]

  beforeEach(() => {
    jest.clearAllMocks()
  })

  it('should fetch furnitures successfully', async () => {
    ;(callApi as jest.Mock).mockResolvedValue(mockResponse)

    const result = await GetFurnituresByRoom(mockPropertyId, mockRoomId)

    expect(callApi).toHaveBeenCalledWith({
      method: 'GET',
      endpoint: `owner/properties/${mockPropertyId}/rooms/${mockRoomId}/furnitures/`
    })
    expect(result).toEqual(mockResponse)
  })

  it('should handle errors when fetching furnitures', async () => {
    const mockError = new Error('Failed to fetch')
    const consoleErrorSpy = jest
      .spyOn(console, 'error')
      .mockImplementation(() => {})
    ;(callApi as jest.Mock).mockRejectedValue(mockError)

    await expect(
      GetFurnituresByRoom(mockPropertyId, mockRoomId)
    ).rejects.toThrow('Failed to fetch')

    expect(callApi).toHaveBeenCalledWith({
      method: 'GET',
      endpoint: `owner/properties/${mockPropertyId}/rooms/${mockRoomId}/furnitures/`
    })

    consoleErrorSpy.mockRestore()
  })
})
