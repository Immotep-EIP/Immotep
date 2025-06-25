import callApi from '@/services/api/apiCaller'
import endpoints from '@/enums/EndPointEnum'
import { Lease } from '@/interfaces/Property/Lease/Lease'
import GetLeasesByProperty from '../GetLeasesByProperty'

jest.mock('@/services/api/apiCaller')

describe('GetLeasesByProperty', () => {
  const mockPropertyId = '123'
  const mockLeases: Lease[] = [
    {
      id: '1',
      start_date: '2024-01-01',
      end_date: '2024-12-31',
      active: true,
      property_id: mockPropertyId,
      property_name: 'Test Property',
      created_at: '2024-01-01T00:00:00Z',
      owner_id: 'owner1',
      owner_email: 'owner@test.com',
      owner_name: 'Owner Test',
      tenant_id: 'tenant1',
      tenant_email: 'tenant@test.com',
      tenant_name: 'Tenant Test'
    },
    {
      id: '2',
      start_date: '2023-01-01',
      end_date: '2023-12-31',
      active: false,
      property_id: mockPropertyId,
      property_name: 'Test Property',
      created_at: '2023-01-01T00:00:00Z',
      owner_id: 'owner1',
      owner_email: 'owner@test.com',
      owner_name: 'Owner Test',
      tenant_id: 'tenant2',
      tenant_email: 'tenant2@test.com',
      tenant_name: 'Tenant Test 2'
    }
  ]

  beforeEach(() => {
    jest.clearAllMocks()
    jest.spyOn(console, 'error').mockImplementation(() => {})
  })

  afterEach(() => {
    jest.restoreAllMocks()
  })

  it('should fetch leases successfully', async () => {
    ;(callApi as jest.Mock).mockResolvedValue(mockLeases)

    const result = await GetLeasesByProperty(mockPropertyId)

    expect(callApi).toHaveBeenCalledWith({
      method: 'GET',
      endpoint: endpoints.owner.properties.leases.list(mockPropertyId)
    })
    expect(result).toEqual(mockLeases)
  })

  it('should handle API errors', async () => {
    const mockError = new Error('API Error')
    ;(callApi as jest.Mock).mockRejectedValue(mockError)

    await expect(GetLeasesByProperty(mockPropertyId)).rejects.toThrow(
      'API Error'
    )
    expect(callApi).toHaveBeenCalledWith({
      method: 'GET',
      endpoint: endpoints.owner.properties.leases.list(mockPropertyId)
    })
    expect(console.error).toHaveBeenCalledWith(
      'Error fetching leases:',
      mockError
    )
  })

  it('should handle empty response', async () => {
    ;(callApi as jest.Mock).mockResolvedValue([])

    const result = await GetLeasesByProperty(mockPropertyId)

    expect(callApi).toHaveBeenCalledWith({
      method: 'GET',
      endpoint: endpoints.owner.properties.leases.list(mockPropertyId)
    })
    expect(result).toEqual([])
  })
})
