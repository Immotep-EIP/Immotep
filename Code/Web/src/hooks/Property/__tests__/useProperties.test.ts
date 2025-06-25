import { renderHook, act, waitFor } from '@testing-library/react'
import useProperties from '@/hooks/Property/useProperties'
import GetProperties from '@/services/api/Owner/Properties/GetProperties'
import GetPropertyDetails from '@/services/api/Owner/Properties/GetPropertyDetails'
import CreatePropertyFunction from '@/services/api/Owner/Properties/CreateProperty'
import UpdatePropertyFunction from '@/services/api/Owner/Properties/UpdateProperty'
import UpdatePropertyPicture from '@/services/api/Owner/Properties/UpdatePropertyPicture'

jest.mock('@/services/api/Owner/Properties/GetProperties')
jest.mock('@/services/api/Owner/Properties/GetPropertyDetails')
jest.mock('@/services/api/Owner/Properties/CreateProperty')
jest.mock('@/services/api/Owner/Properties/UpdateProperty', () => ({
  __esModule: true,
  default: jest.fn()
}))
jest.mock('@/services/api/Owner/Properties/UpdatePropertyPicture')
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

    const { result } = renderHook(() => useProperties())

    await waitFor(() => {
      expect(result.current.loading).toBe(false)
    })

    expect(GetProperties).toHaveBeenCalled()
    expect(result.current.properties).toEqual(mockProperties)
  })

  it('should fetch property details if propertyId is provided', async () => {
    const mockPropertyDetails = { id: '1', name: 'Property 1' }
    ;(GetPropertyDetails as jest.Mock).mockResolvedValue(mockPropertyDetails)

    const { result } = renderHook(() => useProperties('1'))

    await waitFor(() => {
      expect(result.current.loading).toBe(false)
    })

    expect(GetPropertyDetails).toHaveBeenCalledWith('1')
    expect(result.current.propertyDetails).toEqual(mockPropertyDetails)
  })

  it('should handle property creation and update picture', async () => {
    const newProperty = { id: '3', name: 'New Property', leases: [] }
    ;(CreatePropertyFunction as jest.Mock).mockResolvedValue(newProperty)
    ;(UpdatePropertyPicture as jest.Mock).mockResolvedValue('mocked value')

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
          deposit_price: 2000,
          apartment_number: '640'
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
      deposit_price: 2000,
      apartment_number: '640'
    })
    expect(UpdatePropertyPicture).toHaveBeenCalledWith(
      '3',
      'data:image/jpeg;base64,...'
    )
    expect(result.current.properties).toContainEqual(newProperty)
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
          deposit_price: 2000,
          apartment_number: '640'
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
        deposit_price: 2000,
        apartment_number: '640'
      })
    )
    expect(UpdatePropertyPicture).toHaveBeenCalledWith('3', mockImageBase64)
  })

  it('should fetch and set property details', async () => {
    const mockProperty = { id: '1', name: 'Property 1' }
    ;(GetPropertyDetails as jest.Mock).mockResolvedValue(mockProperty)

    const { result } = renderHook(() => useProperties('1'))

    await waitFor(() => {
      expect(result.current.propertyDetails).toEqual(mockProperty)
    })
  })

  // it('should handle errors when fetching property details', async () => {
  //   const mockError = new Error('API error')
  //   const consoleErrorSpy = jest
  //     .spyOn(console, 'error')
  //     .mockImplementation(() => {})
  //   ;(callApi as jest.Mock).mockRejectedValue(mockError)
  //   ;(GetPropertyDetails as jest.Mock).mockRejectedValue(mockError)

  //   const { result } = renderHook(() => useProperties('1'))

  //   await waitFor(() => {
  //     expect(result.current.error).toBe('API error')
  //   })

  //   consoleErrorSpy.mockRestore()
  // })

  it('should create a property', async () => {
    const mockPropertyData = {
      name: 'New Property',
      address: 'St Test',
      city: 'Test',
      postal_code: 'Test',
      country: 'Test',
      area_sqm: 40,
      rental_price_per_month: 1000,
      deposit_price: 2000,
      apartment_number: '640'
    }
    const mockImageBase64 = 'data:image/png;base64,...'
    const mockCreatedProperty = { id: '3', ...mockPropertyData, leases: [] }

    ;(GetProperties as jest.Mock).mockResolvedValue([])
    ;(CreatePropertyFunction as jest.Mock).mockResolvedValue(
      mockCreatedProperty
    )
    ;(UpdatePropertyPicture as jest.Mock).mockResolvedValue({ success: true })

    const { result } = renderHook(() => useProperties())

    await waitFor(() => {
      expect(result.current.loading).toBe(false)
    })

    await act(async () => {
      await result.current.createProperty(mockPropertyData, mockImageBase64)
    })

    expect(result.current.error).toBeNull()
    expect(result.current.properties).toContainEqual(mockCreatedProperty)
    expect(CreatePropertyFunction).toHaveBeenCalledWith(mockPropertyData)
    expect(UpdatePropertyPicture).toHaveBeenCalledWith('3', mockImageBase64)
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
      deposit_price: 2000,
      apartment_number: '640'
    }

    const mockError = new Error('Property creation failed.')
    ;(GetProperties as jest.Mock).mockResolvedValue([])
    ;(CreatePropertyFunction as jest.Mock).mockRejectedValue(mockError)

    const { result } = renderHook(() => useProperties())

    await waitFor(() => {
      expect(result.current.loading).toBe(false)
    })

    await act(async () => {
      try {
        await result.current.createProperty(mockPropertyData, null)
      } catch (err) {
        expect(err).toEqual(mockError)
      }
    })

    expect(result.current.error).toBe('Property creation failed.')
    expect(result.current.properties).toEqual([])
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
      deposit_price: 2000,
      apartment_number: '640'
    }

    ;(GetProperties as jest.Mock).mockResolvedValue([])
    ;(CreatePropertyFunction as jest.Mock).mockResolvedValue(null)

    const { result } = renderHook(() => useProperties())

    await waitFor(() => {
      expect(result.current.loading).toBe(false)
    })

    await act(async () => {
      try {
        await result.current.createProperty(mockPropertyData, null)
      } catch (err: any) {
        expect(err.message).toBe('Property creation failed.')
      }
    })

    expect(result.current.error).toBe('Property creation failed.')
    expect(result.current.properties).toEqual([])
  })

  it('should update a property and its picture', async () => {
    const mockPropertyData = {
      name: 'Updated Property',
      address: 'Updated St',
      city: 'Updated City',
      postal_code: 'Updated Code',
      country: 'Updated Country',
      area_sqm: 50,
      rental_price_per_month: 1200,
      deposit_price: 2400,
      apartment_number: '641'
    }
    const mockImageBase64 = 'data:image/png;base64,updatedImage'
    const mockUpdatedProperty = { id: '1', ...mockPropertyData }
    const initialProperty = { id: '1', name: 'Initial Property' }

    // Setup initial state with a property
    ;(GetProperties as jest.Mock).mockResolvedValue([initialProperty])
    ;(UpdatePropertyFunction as jest.Mock).mockResolvedValue(
      mockUpdatedProperty
    )
    ;(UpdatePropertyPicture as jest.Mock).mockResolvedValue({ success: true })

    const { result } = renderHook(() => useProperties())

    // Wait for initial load
    await waitFor(() => {
      expect(result.current.loading).toBe(false)
      expect(result.current.properties).toEqual([initialProperty])
    })

    // Update property
    await act(async () => {
      await result.current.updateProperty(
        mockPropertyData,
        mockImageBase64,
        '1'
      )
    })

    // Verify results
    expect(UpdatePropertyFunction).toHaveBeenCalledWith(mockPropertyData, '1')
    expect(UpdatePropertyPicture).toHaveBeenCalledWith('1', mockImageBase64)
    expect(result.current.error).toBeNull()
    expect(result.current.properties).toContainEqual(mockUpdatedProperty)
  })

  it('should handle error when property update fails', async () => {
    const mockPropertyData = {
      name: 'Updated Property',
      address: 'Updated St',
      city: 'Updated City',
      postal_code: 'Updated Code',
      country: 'Updated Country',
      area_sqm: 50,
      rental_price_per_month: 1200,
      deposit_price: 2400,
      apartment_number: '641'
    }
    const mockError = new Error('Property update failed')

    // Setup mocks
    ;(GetProperties as jest.Mock).mockResolvedValue([])
    ;(UpdatePropertyFunction as jest.Mock).mockImplementation(() =>
      Promise.reject(mockError)
    )

    const { result } = renderHook(() => useProperties())

    await waitFor(() => {
      expect(result.current.loading).toBe(false)
    })

    await act(async () => {
      try {
        await result.current.updateProperty(mockPropertyData, null, '1')
      } catch (err) {
        expect(err).toEqual(mockError)
      }
    })

    expect(result.current.error).toBe('Property update failed')
    expect(UpdatePropertyFunction).toHaveBeenCalledWith(mockPropertyData, '1')
  })

  it('should refresh properties list', async () => {
    const updatedProperties = [
      { id: '1', name: 'Updated Property 1' },
      { id: '2', name: 'Updated Property 2' }
    ]

    // Setup initial properties
    ;(GetProperties as jest.Mock)
      .mockResolvedValueOnce([])
      .mockResolvedValueOnce(updatedProperties)

    const { result } = renderHook(() => useProperties())

    // Wait for initial load
    await waitFor(() => {
      expect(result.current.loading).toBe(false)
    })

    // Refresh properties
    await act(async () => {
      await result.current.refreshProperties()
    })

    expect(GetProperties).toHaveBeenCalledTimes(2)
    expect(result.current.properties).toEqual(updatedProperties)
    expect(result.current.error).toBeNull()
  })

  it('should refresh property details', async () => {
    const updatedPropertyDetails = {
      id: '1',
      name: 'Updated Property Details'
    }

    ;(GetPropertyDetails as jest.Mock).mockResolvedValue(updatedPropertyDetails)

    const { result } = renderHook(() => useProperties('1'))

    // Wait for initial load
    await waitFor(() => {
      expect(result.current.loading).toBe(false)
    })

    // Refresh property details
    await act(async () => {
      await result.current.refreshPropertyDetails('1')
    })

    expect(GetPropertyDetails).toHaveBeenCalledWith('1')
    expect(result.current.propertyDetails).toEqual(updatedPropertyDetails)
    expect(result.current.error).toBeNull()
  })
})
