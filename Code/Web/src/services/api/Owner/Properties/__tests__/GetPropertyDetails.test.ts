import callApi from '@/services/api/apiCaller'
import { PropertyDetails } from '@/interfaces/Property/Property'
import GetPropertyDetails from '../GetPropertyDetails'

jest.mock('@/services/api/apiCaller')

describe('GetPropertyDetails', () => {
  const mockProperty: PropertyDetails = {
    id: '1',
    name: 'Property 1'
  } as PropertyDetails

  beforeEach(() => {
    jest.clearAllMocks()
  })

  it('should fetch property details successfully', async () => {
    ;(callApi as jest.Mock).mockResolvedValue(mockProperty)

    const result = await GetPropertyDetails('1')

    expect(callApi).toHaveBeenCalledWith({
      method: 'GET',
      endpoint: 'owner/properties/1/'
    })
    expect(result).toEqual(mockProperty)
  })

  it('should handle errors gracefully', async () => {
    const mockError = new Error('Network error')
    const consoleErrorSpy = jest
      .spyOn(console, 'error')
      .mockImplementation(() => {})
    ;(callApi as jest.Mock).mockRejectedValue(mockError)

    await expect(GetPropertyDetails('1')).rejects.toThrow('Network error')

    expect(callApi).toHaveBeenCalledWith({
      method: 'GET',
      endpoint: 'owner/properties/1/'
    })

    consoleErrorSpy.mockRestore()
  })
})
