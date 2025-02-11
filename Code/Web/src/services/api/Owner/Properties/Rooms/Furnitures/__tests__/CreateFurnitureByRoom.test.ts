import callApi from '@/services/api/apiCaller'
import CreateFurnitureByRoom from '../CreateFurnitureByRoom'

jest.mock('@/services/api/apiCaller')

describe('CreateFurnitureByRoom', () => {
  const mockPropertyId = '123'
  const mockRoomId = '456'
  const mockData = {
    name: 'Sofa',
    quantity: 1,
    property_id: mockPropertyId,
    room_id: mockRoomId
  }
  const mockResponse = {
    id: '789',
    ...mockData
  }

  beforeEach(() => {
    jest.clearAllMocks()
  })

  it('should create furniture successfully', async () => {
    ;(callApi as jest.Mock).mockResolvedValue(mockResponse)

    const result = await CreateFurnitureByRoom(
      mockPropertyId,
      mockRoomId,
      mockData
    )

    expect(callApi).toHaveBeenCalledWith({
      method: 'POST',
      endpoint: `owner/properties/${mockPropertyId}/rooms/${mockRoomId}/furnitures/`,
      data: mockData
    })
    expect(result).toEqual(mockResponse)
  })

  it('should handle errors when creating furniture', async () => {
    const mockError = new Error('Failed to create')
    const consoleErrorSpy = jest
      .spyOn(console, 'error')
      .mockImplementation(() => {})
    ;(callApi as jest.Mock).mockRejectedValue(mockError)

    await expect(
      CreateFurnitureByRoom(mockPropertyId, mockRoomId, mockData)
    ).rejects.toThrow('Failed to create')

    expect(callApi).toHaveBeenCalledWith({
      method: 'POST',
      endpoint: `owner/properties/${mockPropertyId}/rooms/${mockRoomId}/furnitures/`,
      data: mockData
    })

    consoleErrorSpy.mockRestore()
  })
})
