import callApi from '@/services/api/apiCaller'
import { ROOM_TYPES } from '@/utils/types/roomTypes'
import CreateRoomByProperty from '../CreateRoomByProperty'

jest.mock('@/services/api/apiCaller')

describe('CreateRoomByProperty', () => {
  const mockPropertyId = '123'
  const mockRoomName = 'Living Room'
  const mockRoomType = 'bedroom'
  const mockResponse = { id: '1', name: mockRoomName }

  beforeEach(() => {
    jest.clearAllMocks()
  })

  it('should create room successfully', async () => {
    ;(callApi as jest.Mock).mockResolvedValue(mockResponse)

    const result = await CreateRoomByProperty(
      mockPropertyId,
      mockRoomName,
      mockRoomType
    )

    expect(callApi).toHaveBeenCalledWith({
      method: 'POST',
      endpoint: `owner/properties/${mockPropertyId}/rooms/`,
      body: JSON.stringify({
        name: mockRoomName,
        type: mockRoomType.toLowerCase()
      })
    })
    expect(result).toEqual(mockResponse)
  })

  it('should handle errors when creating room', async () => {
    const mockError = new Error('Failed to create')
    const consoleErrorSpy = jest
      .spyOn(console, 'error')
      .mockImplementation(() => {})
    ;(callApi as jest.Mock).mockRejectedValue(mockError)

    await expect(
      CreateRoomByProperty(mockPropertyId, mockRoomName, mockRoomType)
    ).rejects.toThrow('Failed to create')
    expect(callApi).toHaveBeenCalledWith({
      method: 'POST',
      endpoint: `owner/properties/${mockPropertyId}/rooms/`,
      body: JSON.stringify({
        name: mockRoomName,
        type: mockRoomType.toLowerCase()
      })
    })

    consoleErrorSpy.mockRestore()
  })

  it('should throw error for invalid room type', async () => {
    const invalidRoomType = 'invalid_type'
    const expectedError = `Invalid room type. Must be one of: ${ROOM_TYPES.join(', ')}`
    const consoleErrorSpy = jest
      .spyOn(console, 'error')
      .mockImplementation(() => {})

    await expect(
      CreateRoomByProperty(mockPropertyId, mockRoomName, invalidRoomType)
    ).rejects.toThrow(expectedError)

    expect(callApi).not.toHaveBeenCalled()
    expect(consoleErrorSpy).toHaveBeenCalledWith(
      'Error creating room:',
      expect.any(Error)
    )

    consoleErrorSpy.mockRestore()
  })

  it('should handle room type case insensitivity', async () => {
    ;(callApi as jest.Mock).mockResolvedValue(mockResponse)

    const result = await CreateRoomByProperty(
      mockPropertyId,
      mockRoomName,
      'BEDROOM'
    )

    expect(callApi).toHaveBeenCalledWith({
      method: 'POST',
      endpoint: `owner/properties/${mockPropertyId}/rooms/`,
      body: JSON.stringify({ name: mockRoomName, type: 'bedroom' })
    })
    expect(result).toEqual(mockResponse)
  })

  it('should handle all valid room types', async () => {
    ;(callApi as jest.Mock).mockResolvedValue(mockResponse)

    await Promise.all(
      ROOM_TYPES.map(roomType =>
        CreateRoomByProperty(mockPropertyId, mockRoomName, roomType)
      )
    )

    ROOM_TYPES.forEach(roomType => {
      expect(callApi).toHaveBeenCalledWith({
        method: 'POST',
        endpoint: `owner/properties/${mockPropertyId}/rooms/`,
        body: JSON.stringify({
          name: mockRoomName,
          type: roomType.toLowerCase()
        })
      })
    })
  })
})
