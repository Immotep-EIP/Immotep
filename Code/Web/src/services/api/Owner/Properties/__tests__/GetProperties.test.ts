import callApi from '@/services/api/apiCaller'
import { PropertyDetails } from '@/interfaces/Property/Property'
import GetProperties from '../GetProperties'

jest.mock('@/services/api/apiCaller')

describe('GetProperties', () => {
  const mockProperties: PropertyDetails[] = [
    { id: '1', name: 'Property 1' } as PropertyDetails,
    { id: '2', name: 'Property 2' } as PropertyDetails
  ]

  beforeEach(() => {
    jest.clearAllMocks()
  })

  it('should fetch properties successfully', async () => {
    ;(callApi as jest.Mock).mockResolvedValue(mockProperties)

    const result = await GetProperties()

    expect(callApi).toHaveBeenCalledWith({
      method: 'GET',
      endpoint: 'owner/properties/',
      params: {
        archive: false
      }
    })
    expect(result).toEqual(mockProperties)
  })

  it('should handle errors gracefully', async () => {
    const mockError = new Error('Network error')
    const consoleErrorSpy = jest
      .spyOn(console, 'error')
      .mockImplementation(() => {})
    ;(callApi as jest.Mock).mockRejectedValue(mockError)

    await expect(GetProperties()).rejects.toThrow('Network error')

    expect(callApi).toHaveBeenCalledWith({
      method: 'GET',
      endpoint: 'owner/properties/',
      params: {
        archive: false
      }
    })

    consoleErrorSpy.mockRestore()
  })
})
