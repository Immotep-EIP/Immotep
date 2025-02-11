import { renderHook, act, waitFor } from '@testing-library/react'
import useProperties from '@/hooks/useEffect/useProperties'
import GetProperties from '@/services/api/Owner/Properties/GetProperties'
import GetPropertyDetails from '@/services/api/Owner/Properties/GetPropertyDetails'
import CreatePropertyFunction from '@/services/api/Owner/Properties/CreateProperty'
import UpdatePropertyPicture from '@/services/api/Owner/Properties/UpdatePropertyPicture'
import {
  savePropertiesToDB,
  getPropertiesFromDB
} from '@/utils/cache/property/indexedDB'
import callApi from '@/services/api/apiCaller'

jest.mock('@/services/api/Owner/Properties/GetProperties')
jest.mock('@/services/api/Owner/Properties/GetPropertyDetails')
jest.mock('@/services/api/Owner/Properties/CreateProperty', () => jest.fn())
jest.mock('@/services/api/Owner/Properties/UpdatePropertyPicture')
jest.mock('@/utils/cache/property/indexedDB')
jest.mock('@/services/api/apiCaller')

describe('useProperties', () => {
  const mockProperties = [
    { id: '1', name: 'Property 1' },
    { id: '2', name: 'Property 2' }
  ]

  beforeEach(() => {
    jest.clearAllMocks()
  })

  it('should fetch properties on mount', async () => {
    ;(GetProperties as jest.Mock).mockResolvedValue(mockProperties)
    ;(getPropertiesFromDB as jest.Mock).mockResolvedValue([])
    ;(savePropertiesToDB as jest.Mock).mockResolvedValue('mocked value')

    const { result } = renderHook(() => useProperties())

    await waitFor(() => {
      expect(result.current.loading).toBe(false)
    })

    expect(GetProperties).toHaveBeenCalled()
    expect(savePropertiesToDB).toHaveBeenCalledWith(mockProperties)
    expect(result.current.properties).toEqual(mockProperties)
  })

  it('should fetch property details if propertyId is provided', async () => {
    const mockPropertyDetails = { id: '1', name: 'Property 1' }
    ;(GetPropertyDetails as jest.Mock).mockResolvedValue(mockPropertyDetails)
    ;(getPropertiesFromDB as jest.Mock).mockResolvedValue([])
    ;(savePropertiesToDB as jest.Mock).mockResolvedValue('mocked value')

    const { result } = renderHook(() => useProperties('1'))

    await waitFor(() => {
      expect(result.current.loading).toBe(false)
    })

    expect(GetPropertyDetails).toHaveBeenCalledWith('1')
    expect(savePropertiesToDB).toHaveBeenCalledWith([mockPropertyDetails])
    expect(result.current.propertyDetails).toEqual(mockPropertyDetails)
  })

  it('should handle property creation and update picture', async () => {
    const newProperty = { id: '3', name: 'New Property' }
    ;(CreatePropertyFunction as jest.Mock).mockResolvedValue(newProperty)
    ;(UpdatePropertyPicture as jest.Mock).mockResolvedValue('mocked value')
    ;(savePropertiesToDB as jest.Mock).mockResolvedValue('mocked value')

    const { result } = renderHook(() => useProperties())

    await act(async () =>
      result.current.createProperty(
        {
          name: 'New Property',
          address: 'St Test',
          city: 'Test',
          postal_code: 'Test',
          country: 'Test',
          area_sqm: 40,
          rental_price_per_month: 1000,
          deposit_price: 2000
        },
        'data:image/jpeg;base64,...'
      )
    )

    expect(CreatePropertyFunction).toHaveBeenCalledWith({
      name: 'New Property',
      address: 'St Test',
      city: 'Test',
      postal_code: 'Test',
      country: 'Test',
      area_sqm: 40,
      rental_price_per_month: 1000,
      deposit_price: 2000
    })
    expect(UpdatePropertyPicture).toHaveBeenCalledWith('3', '...')
    expect(savePropertiesToDB).toHaveBeenCalledWith([newProperty])
    expect(result.current.properties).toContainEqual(newProperty)
  })

  it('should use cached properties if available', async () => {
    ;(getPropertiesFromDB as jest.Mock).mockResolvedValue(mockProperties)

    const { result } = renderHook(() => useProperties())

    await waitFor(() => {
      expect(result.current.loading).toBe(false)
    })

    expect(GetProperties).not.toHaveBeenCalled()
    expect(result.current.properties).toEqual(mockProperties)
  })

  it('should handle errors gracefully', async () => {
    ;(GetProperties as jest.Mock).mockRejectedValue(new Error('API error'))

    const { result } = renderHook(() => useProperties())

    await waitFor(() => {
      expect(result.current.loading).toBe(false)
    })
    ;(GetProperties as jest.Mock).mockRejectedValue(new Error('API error'))
  })

  it('should create a property and upload a picture', async () => {
    const mockCreateProperty = { id: '3', name: 'New Property' }
    const mockImageBase64 = 'data:image/jpeg;base64,...'
    ;(CreatePropertyFunction as jest.Mock).mockResolvedValue(mockCreateProperty)
    ;(UpdatePropertyPicture as jest.Mock).mockResolvedValue('mocked value')
    ;(savePropertiesToDB as jest.Mock).mockResolvedValue('mocked value')

    const { result } = renderHook(() => useProperties())

    await act(async () =>
      result.current.createProperty(
        {
          name: 'New Property',
          address: 'St Test',
          city: 'Test',
          postal_code: 'Test',
          country: 'Test',
          area_sqm: 40,
          rental_price_per_month: 1000,
          deposit_price: 2000
        },
        mockImageBase64
      )
    )

    expect(CreatePropertyFunction).toHaveBeenCalledWith(
      expect.objectContaining({
        name: 'New Property',
        address: 'St Test',
        city: 'Test',
        postal_code: 'Test',
        country: 'Test',
        area_sqm: 40,
        rental_price_per_month: 1000,
        deposit_price: 2000
      })
    )
    expect(UpdatePropertyPicture).toHaveBeenCalledWith('3', '...')
  })

  it('should fetch properties from cache or API', async () => {
    const mockProperties = [{ id: '1', name: 'Property 1' }]
    ;(getPropertiesFromDB as jest.Mock).mockResolvedValue(mockProperties)
    ;(GetProperties as jest.Mock).mockResolvedValue(mockProperties)

    const { result } = renderHook(() => useProperties())

    await waitFor(() => {
      expect(result.current.properties).toEqual(mockProperties)
    })
    expect(GetProperties).not.toHaveBeenCalled()
  })

  it('should fetch and set property details', async () => {
    const mockProperty = { id: '1', name: 'Property 1' }
    ;(getPropertiesFromDB as jest.Mock).mockResolvedValue([mockProperty])
    ;(GetPropertyDetails as jest.Mock).mockResolvedValue(mockProperty)

    const { result } = renderHook(() => useProperties('1'))

    await waitFor(() => {
      expect(result.current.propertyDetails).toEqual(mockProperty)
    })
  })

  it('should save fetched properties to the database', async () => {
    const mockProperties = [{ id: '1', name: 'Property 1' }]
    ;(getPropertiesFromDB as jest.Mock).mockResolvedValue([])
    ;(GetProperties as jest.Mock).mockResolvedValue(mockProperties)
    ;(savePropertiesToDB as jest.Mock).mockResolvedValue('mocked value')

    const { result } = renderHook(() => useProperties())

    await waitFor(() => {
      expect(savePropertiesToDB).toHaveBeenCalledWith(mockProperties)
    })
  })

  it('should handle errors when fetching property details', async () => {
    const mockError = new Error('API error')
    const consoleErrorSpy = jest
      .spyOn(console, 'error')
      .mockImplementation(() => {})
    ;(callApi as jest.Mock).mockRejectedValue(mockError)
    ;(GetPropertyDetails as jest.Mock).mockRejectedValue(mockError)

    const { result } = renderHook(() => useProperties('1'))

    await waitFor(() => {
      expect(result.current.error).toBe('API error')
    })

    consoleErrorSpy.mockRestore()
  })

  it('should create a property and update state and DB', async () => {
    const mockPropertyData = {
      name: 'New Property',
      address: 'St Test',
      city: 'Test',
      postal_code: 'Test',
      country: 'Test',
      area_sqm: 40,
      rental_price_per_month: 1000,
      deposit_price: 2000
    }
    const mockImageBase64 = 'data:image/png;base64,...'
    const mockCreatedProperty = {}

    ;(CreatePropertyFunction as jest.Mock).mockResolvedValue(
      mockCreatedProperty
    )
    ;(UpdatePropertyPicture as jest.Mock).mockResolvedValue(null)
    ;(savePropertiesToDB as jest.Mock).mockResolvedValue(undefined)

    const { result } = renderHook(() => useProperties())

    await act(async () => {
      await result.current.createProperty(mockPropertyData, mockImageBase64)
    })

    expect(result.current.properties).toContain(mockCreatedProperty)
    expect(savePropertiesToDB).toHaveBeenCalledWith([mockCreatedProperty])
    expect(result.current.error).toBeNull()
  })

  it('should handle error if property creation fails', async () => {
    const mockPropertyData = {
      name: 'New Property',
      address: 'St Test',
      city: 'Test',
      postal_code: 'Test',
      country: 'Test',
      area_sqm: 40,
      rental_price_per_month: 1000,
      deposit_price: 2000
    }

    const mockError = new Error('Property creation failed.')
    ;(CreatePropertyFunction as jest.Mock).mockRejectedValue(mockError)

    const { result } = renderHook(() => useProperties())

    await act(async () => {
      try {
        await result.current.createProperty(mockPropertyData, null)
        throw new Error('Expected error was not thrown')
      } catch (err: any) {
        expect(err.message).toBe('Property creation failed.')
      }
    })

    await waitFor(() => {
      expect(result.current.error).toBe('Property creation failed.')
    })
  })

  it('should handle error if fetching properties fails', async () => {
    ;(GetProperties as jest.Mock).mockRejectedValue(
      new Error('Error fetching properties:')
    )

    const { result } = renderHook(() => useProperties())

    await waitFor(() => {
      expect(result.current.error).toBe('Error fetching properties:')
    })

    expect(result.current.properties).toEqual([])
    expect(result.current.loading).toBe(false)
  })

  it('should throw error when created property is falsy', async () => {
    const mockPropertyData = {
      name: 'New Property',
      address: 'St Test',
      city: 'Test',
      postal_code: 'Test',
      country: 'Test',
      area_sqm: 40,
      rental_price_per_month: 1000,
      deposit_price: 2000
    }

    ;(GetProperties as jest.Mock).mockResolvedValue([])
    ;(CreatePropertyFunction as jest.Mock).mockResolvedValue(null)
    ;(getPropertiesFromDB as jest.Mock).mockResolvedValue([])

    const { result } = renderHook(() => useProperties())

    await waitFor(() => {
      expect(result.current.loading).toBe(false)
    })

    await act(async () => {
      try {
        await result.current.createProperty(mockPropertyData, null)
        throw new Error('Expected error was not thrown')
      } catch (err: any) {
        expect(err.message).toBe('Property creation failed.')
      }
    })

    expect(result.current.error).toBe('Property creation failed.')
    expect(result.current.loading).toBe(false)
    expect(result.current.properties).toEqual([])
  })
})
