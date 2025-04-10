import callApi from '@/services/api/apiCaller'
import {
  CreatePropertyPayload,
  PropertyDetails
} from '@/interfaces/Property/Property'
import UpdatePropertyFunction from '../UpdateProperty'

jest.mock('@/services/api/apiCaller')

describe('UpdatePropertyFunction', () => {
  const mockPropertyData: CreatePropertyPayload = {
    name: 'Updated Property',
    address: 'Updated St Test',
    city: 'Updated City',
    postal_code: 'Updated Code',
    country: 'Updated Country',
    area_sqm: 50,
    rental_price_per_month: 1200,
    deposit_price: 2400,
    apartment_number: '641'
  }

  const mockPropertyId = 'test-property-id'

  const mockUpdatedProperty: PropertyDetails = {
    id: mockPropertyId,
    ...mockPropertyData,
    archived: false,
    created_at: '2024-03-10T10:00:00Z',
    owner_id: '1',
    picture_id: '1',
    nb_damage: 0,
    status: 'active',
    tenant: 'test-tenant',
    start_date: '2024-03-10T10:00:00Z',
    end_date: '2024-03-10T11:00:00Z'
  }

  let consoleErrorSpy: jest.SpyInstance

  beforeEach(() => {
    jest.clearAllMocks()
    // Mock console.error before each test
    consoleErrorSpy = jest.spyOn(console, 'error').mockImplementation(() => {})
  })

  afterEach(() => {
    // Restore console.error after each test
    consoleErrorSpy.mockRestore()
  })

  it('should successfully update a property', async () => {
    ;(callApi as jest.Mock).mockResolvedValue(mockUpdatedProperty)

    const result = await UpdatePropertyFunction(
      mockPropertyData,
      mockPropertyId
    )

    expect(result).toEqual(mockUpdatedProperty)
    expect(callApi).toHaveBeenCalledWith({
      method: 'PUT',
      endpoint: `owner/properties/${mockPropertyId}/`,
      body: mockPropertyData
    })
    expect(callApi).toHaveBeenCalledTimes(1)
    expect(consoleErrorSpy).not.toHaveBeenCalled()
  })

  it('should throw an error when API call fails', async () => {
    const mockError = new Error('Update failed')
    ;(callApi as jest.Mock).mockRejectedValue(mockError)

    await expect(
      UpdatePropertyFunction(mockPropertyData, mockPropertyId)
    ).rejects.toThrow('Update failed')

    expect(callApi).toHaveBeenCalledWith({
      method: 'PUT',
      endpoint: `owner/properties/${mockPropertyId}/`,
      body: mockPropertyData
    })
    expect(consoleErrorSpy).toHaveBeenCalledWith(
      'Error fetching data:',
      mockError
    )
  })

  it('should handle empty property ID', async () => {
    const emptyId = ''
    const mockError = new Error('Invalid property ID')
    ;(callApi as jest.Mock).mockRejectedValue(mockError)

    await expect(
      UpdatePropertyFunction(mockPropertyData, emptyId)
    ).rejects.toThrow()

    expect(callApi).toHaveBeenCalledWith({
      method: 'PUT',
      endpoint: 'owner/properties//',
      body: mockPropertyData
    })
  })

  it('should handle partial property updates', async () => {
    const partialUpdate: Partial<CreatePropertyPayload> = {
      name: 'Updated Name Only',
      rental_price_per_month: 1500
    }

    const mockPartialResponse = {
      ...mockUpdatedProperty,
      ...partialUpdate
    }

    ;(callApi as jest.Mock).mockResolvedValue(mockPartialResponse)

    const result = await UpdatePropertyFunction(
      partialUpdate as CreatePropertyPayload,
      mockPropertyId
    )

    expect(result).toEqual(mockPartialResponse)
    expect(callApi).toHaveBeenCalledWith({
      method: 'PUT',
      endpoint: `owner/properties/${mockPropertyId}/`,
      body: partialUpdate
    })
    expect(consoleErrorSpy).not.toHaveBeenCalled()
  })
})
