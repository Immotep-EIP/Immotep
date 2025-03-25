import callApi from '@/services/api/apiCaller'
import CreateRoomByProperty from '../CreateRoomByProperty'

jest.mock('@/services/api/apiCaller')

describe('CreateRoomByProperty', () => {
  const mockPropertyId = '123'
  const mockRoomName = 'Living Room'
  const mockResponse = { id: '1', name: mockRoomName }

  beforeEach(() => {
    jest.clearAllMocks()
  })

  it('should create room successfully', async () => {
    ;(callApi as jest.Mock).mockResolvedValue(mockResponse)

    const result = await CreateRoomByProperty(mockPropertyId, mockRoomName)

    expect(callApi).toHaveBeenCalledWith({
      method: 'POST',
      endpoint: `owner/properties/${mockPropertyId}/rooms/`,
      data: JSON.stringify({ name: mockRoomName })
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
      CreateRoomByProperty(mockPropertyId, mockRoomName)
    ).rejects.toThrow('Failed to create')
    expect(callApi).toHaveBeenCalledWith({
      method: 'POST',
      endpoint: `owner/properties/${mockPropertyId}/rooms/`,
      data: JSON.stringify({ name: mockRoomName })
    })

    consoleErrorSpy.mockRestore()
  })
})
