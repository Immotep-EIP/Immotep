import { InviteTenants } from '@/services/api/Tenant/InviteTenant'
import callApi from '@/services/api/apiCaller'
import {
  InviteTenant,
  InviteTenantResponse
} from '@/interfaces/Tenant/InviteTenant'

jest.mock('@/services/api/apiCaller')

const mockedCallApi = callApi as jest.MockedFunction<typeof callApi>

describe('InviteTenants', () => {
  beforeEach(() => {
    jest.clearAllMocks()
  })

  it('should call callApi with the correct parameters and return the response', async () => {
    const mockResponse: InviteTenantResponse = {
      created_at: new Date().toDateString(),
      end_date: '2023-12-31T00:00:00Z',
      id: '123',
      property_id: '456',
      start_date: '2023-10-01T00:00:00Z',
      tenant_email: 'tenant@example.com'
    }
    mockedCallApi.mockResolvedValueOnce(mockResponse)

    const tenantInfo: InviteTenant = {
      propertyId: '123',
      tenant_email: 'tenant@example.com',
      start_date: new Date().toDateString()
    }

    const result = await InviteTenants(tenantInfo)

    expect(mockedCallApi).toHaveBeenCalledWith({
      method: 'POST',
      endpoint: `owner/properties/${tenantInfo.propertyId}/send-invite/`,
      body: tenantInfo
    })

    expect(result).toEqual(mockResponse)
  })

  it('should handle errors during tenant invitation', async () => {
    const mockError = new Error('API call failed')
    mockedCallApi.mockRejectedValueOnce(mockError)

    const tenantInfo: InviteTenant = {
      propertyId: '123',
      tenant_email: 'tenant@example.com',
      start_date: '2023-10-01T00:00:00Z'
    }

    const consoleErrorSpy = jest
      .spyOn(console, 'error')
      .mockImplementation(() => {})

    await expect(InviteTenants(tenantInfo)).rejects.toThrow('API call failed')

    expect(consoleErrorSpy).toHaveBeenCalledWith(
      'Error fetching data:',
      mockError
    )

    consoleErrorSpy.mockRestore()
  })

  it('should throw error when propertyId is missing', async () => {
    const tenantInfo = {
      tenant_email: 'tenant@example.com',
      start_date: '2023-10-01T00:00:00Z'
    } as InviteTenant

    const consoleErrorSpy = jest
      .spyOn(console, 'error')
      .mockImplementation(() => {})

    await expect(InviteTenants(tenantInfo)).rejects.toThrow(
      'Property ID is required'
    )
    expect(mockedCallApi).not.toHaveBeenCalled()

    consoleErrorSpy.mockRestore()
  })
})
