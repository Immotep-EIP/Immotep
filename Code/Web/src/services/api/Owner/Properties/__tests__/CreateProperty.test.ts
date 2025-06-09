import CreatePropertyFunction from '@/services/api/Owner/Properties/CreateProperty'
import callApi from '@/services/api/apiCaller'
import {
  CreatePropertyPayload,
  PropertyDetails
} from '@/interfaces/Property/Property'

jest.mock('@/services/api/apiCaller')

const mockedCallApi = callApi as jest.MockedFunction<typeof callApi>

describe('CreatePropertyFunction', () => {
  beforeEach(() => {
    jest.clearAllMocks()
  })

  it('should call callApi with the correct parameters and return the response', async () => {
    const mockResponse: PropertyDetails = {
      id: '1',
      owner_id: 'owner123',
      archived: false,
      name: 'Test Property',
      address: '123 Test St',
      city: 'Test City',
      postal_code: '12345',
      country: 'Test Country',
      area_sqm: 100,
      rental_price_per_month: 1500,
      deposit_price: 3000,
      picture_id: 'pic123',
      created_at: new Date().toDateString(),
      nb_damage: 0,
      status: 'available',
      apartment_number: '914',
      start_date: '2023-01-01',
      end_date: '2023-12-31',
      tenant: 'tenant123',
      lease: {
        active: true,
        created_at: new Date().toDateString(),
        start_date: '2023-01-01',
        end_date: '2023-12-31',
        id: 'lease123',
        owner_email: 'owner@example.com',
        owner_id: 'owner123',
        owner_name: 'John Owner',
        property_id: '1',
        property_name: 'Test Property',
        tenant_email: 'tenant123',
        tenant_id: 'tenant123',
        tenant_name: 'Jane Tenant'
      },
      leases: []
    }
    mockedCallApi.mockResolvedValueOnce(mockResponse)

    const propertyData: CreatePropertyPayload = {
      name: 'Test Property',
      address: '123 Test St',
      city: 'Test City',
      postal_code: '12345',
      country: 'Test Country',
      area_sqm: 100,
      rental_price_per_month: 1500,
      deposit_price: 3000,
      apartment_number: '914'
    }

    const result = await CreatePropertyFunction(propertyData)

    expect(mockedCallApi).toHaveBeenCalledWith({
      method: 'POST',
      endpoint: 'owner/properties/',
      body: propertyData
    })

    expect(result).toEqual(mockResponse)
  })

  it('should handle errors during property creation', async () => {
    const mockError = new Error('API call failed')
    mockedCallApi.mockRejectedValueOnce(mockError)

    const propertyData: CreatePropertyPayload = {
      name: 'Test Property',
      address: '123 Test St',
      city: 'Test City',
      postal_code: '12345',
      country: 'Test Country',
      area_sqm: 100,
      rental_price_per_month: 1500,
      deposit_price: 3000,
      apartment_number: '914'
    }

    const consoleErrorSpy = jest
      .spyOn(console, 'error')
      .mockImplementation(() => {})

    await expect(CreatePropertyFunction(propertyData)).rejects.toThrow(
      'API call failed'
    )

    expect(consoleErrorSpy).toHaveBeenCalledWith(
      'Error fetching data:',
      mockError
    )

    consoleErrorSpy.mockRestore()
  })
})
