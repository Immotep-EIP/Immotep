import '@testing-library/jest-dom'
import { render, screen, fireEvent, waitFor } from '@testing-library/react'
import { message } from 'antd'
import RealPropertyUpdate from '@/views/RealProperty/update/RealPropertyUpdate'
import useProperties from '@/hooks/Property/useProperties'
import useImageUpload from '@/hooks/Image/useImageUpload'
import useImageCache from '@/hooks/Image/useImageCache'

jest.mock('@/hooks/Property/useProperties')
jest.mock('@/hooks/Image/useImageUpload')
jest.mock('@/hooks/Image/useImageCache')
jest.mock('react-i18next', () => ({
  useTranslation: () => ({
    t: (key: string) => key
  })
}))
jest.mock('antd', () => {
  const originalModule = jest.requireActual('antd')
  return {
    ...originalModule,
    message: {
      success: jest.fn(),
      error: jest.fn()
    }
  }
})

describe('RealPropertyUpdate Component', () => {
  const mockPropertyData = {
    id: '1',
    name: 'Test Property',
    apartment_number: '4B',
    address: '123 Main St',
    postal_code: '12345',
    city: 'Test City',
    country: 'Test Country',
    area_sqm: 100,
    rental_price_per_month: 1000,
    deposit_price: 2000,
    archived: false,
    created_at: '2024-03-10T10:00:00Z',
    owner_id: '1',
    picture_id: '1',
    nb_damage: 0,
    status: 'active',
    start_date: '2024-03-10T10:00:00Z',
    end_date: '2024-03-10T11:00:00Z',
    tenant: 'tenant123',
    lease: {
      active: true,
      created_at: '2024-03-10T10:00:00Z',
      start_date: '2024-03-10T10:00:00Z',
      end_date: '2024-03-10T11:00:00Z',
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

  const mockSetIsModalUpdateOpen = jest.fn()
  const mockSetIsPropertyUpdated = jest.fn()
  const mockUpdateCache = jest.fn()

  beforeEach(() => {
    jest.clearAllMocks()
    ;(useProperties as jest.Mock).mockReturnValue({
      loading: false,
      updateProperty: jest.fn()
    })

    // Correction du mock pour éviter l'avertissement "value is not a valid prop"
    ;(useImageUpload as jest.Mock).mockReturnValue({
      uploadProps: {
        // Utiliser fileList au lieu de value
        fileList: [],
        beforeUpload: jest.fn(() => false),
        onChange: jest.fn(),
        onRemove: jest.fn(),
        accept: '.jpg,.jpeg,.png',
        maxCount: 1,
        listType: 'picture-card'
      },
      imageBase64: 'mocked-image-base64',
      resetImage: jest.fn()
    })
    ;(useImageCache as jest.Mock).mockImplementation(() => ({
      updateCache: mockUpdateCache,
      data: null,
      isLoading: false
    }))
  })

  const renderComponent = () =>
    render(
      <RealPropertyUpdate
        isModalUpdateOpen
        setIsModalUpdateOpen={mockSetIsModalUpdateOpen}
        propertyData={mockPropertyData}
        setIsPropertyUpdated={mockSetIsPropertyUpdated}
      />
    )

  it('renders the modal with form fields pre-filled', () => {
    renderComponent()

    expect(
      screen.getByText('pages.real_property.update_real_property.title')
    ).toBeInTheDocument()
    expect(screen.getByDisplayValue('Test Property')).toBeInTheDocument()
    expect(screen.getByDisplayValue('4B')).toBeInTheDocument()
    expect(screen.getByDisplayValue('123 Main St')).toBeInTheDocument()
    expect(screen.getByDisplayValue('12345')).toBeInTheDocument()
    expect(screen.getByDisplayValue('Test City')).toBeInTheDocument()
    expect(screen.getByDisplayValue('Test Country')).toBeInTheDocument()
    expect(screen.getByDisplayValue('100')).toBeInTheDocument()
    expect(screen.getByDisplayValue('1000')).toBeInTheDocument()
    expect(screen.getByDisplayValue('2000')).toBeInTheDocument()
  })

  it('closes the modal when cancel button is clicked', () => {
    renderComponent()

    fireEvent.click(screen.getByText('components.button.cancel'))

    expect(mockSetIsModalUpdateOpen).toHaveBeenCalledWith(false)
  })

  it('closes the modal when clicking the close button', () => {
    renderComponent()

    const closeButton = screen.getByRole('button', { name: /close/i })
    fireEvent.click(closeButton)

    expect(mockSetIsModalUpdateOpen).toHaveBeenCalledWith(false)
  })

  it('submits the form and updates the property successfully', async () => {
    const mockUpdateProperty = useProperties().updateProperty as jest.Mock
    mockUpdateProperty.mockResolvedValueOnce({})

    renderComponent()

    fireEvent.click(screen.getByText('components.button.update'))

    await waitFor(() => {
      expect(mockUpdateProperty).toHaveBeenCalledWith(
        {
          name: 'Test Property',
          apartment_number: '4B',
          address: '123 Main St',
          postal_code: '12345',
          city: 'Test City',
          country: 'Test Country',
          area_sqm: 100,
          rental_price_per_month: 1000,
          deposit_price: 2000
        },
        'mocked-image-base64',
        '1'
      )
      expect(mockUpdateCache).toHaveBeenCalledWith('mocked-image-base64')
      expect(mockSetIsModalUpdateOpen).toHaveBeenCalledWith(false)
      expect(message.success).toHaveBeenCalledWith(
        'pages.real_property.update_real_property.property_updated'
      )
      expect(mockSetIsPropertyUpdated).toHaveBeenCalledWith(true)
    })
  })

  it('handles image data promise correctly', async () => {
    const mockImageData = 'test-image-data'
    ;(useImageUpload as jest.Mock).mockReturnValue({
      uploadProps: {
        // Fournir des props valides pour le composant Upload
        fileList: [],
        beforeUpload: jest.fn(() => false),
        onChange: jest.fn(),
        onRemove: jest.fn(),
        accept: '.jpg,.jpeg,.png',
        maxCount: 1,
        listType: 'picture-card'
      },
      imageBase64: mockImageData,
      resetImage: jest.fn()
    })

    renderComponent()

    await waitFor(() => {
      expect(useImageCache).toHaveBeenCalledWith(
        mockPropertyData.id,
        expect.any(Function)
      )
    })

    const imageCacheCallback = (useImageCache as jest.Mock).mock.calls[0][1]
    const result = await imageCacheCallback()
    expect(result).toEqual({ data: mockImageData })
  })

  it('shows an error message when form submission fails', async () => {
    const mockUpdateProperty = useProperties().updateProperty as jest.Mock
    mockUpdateProperty.mockRejectedValueOnce(new Error('Update failed'))

    renderComponent()

    fireEvent.click(screen.getByText('components.button.update'))

    await waitFor(() => {
      expect(message.error).toHaveBeenCalledWith(
        'pages.real_property.update_real_property.error_property_updated'
      )
    })
  })

  it('shows an error message when form fields are not filled', async () => {
    renderComponent()

    fireEvent.change(screen.getByDisplayValue('Test Property'), {
      target: { value: '' }
    })
    fireEvent.change(screen.getByDisplayValue('4B'), { target: { value: '' } })
    fireEvent.change(screen.getByDisplayValue('123 Main St'), {
      target: { value: '' }
    })
    fireEvent.change(screen.getByDisplayValue('12345'), {
      target: { value: '' }
    })
    fireEvent.change(screen.getByDisplayValue('Test City'), {
      target: { value: '' }
    })
    fireEvent.change(screen.getByDisplayValue('Test Country'), {
      target: { value: '' }
    })
    fireEvent.change(screen.getByDisplayValue('100'), { target: { value: '' } })
    fireEvent.change(screen.getByDisplayValue('1000'), {
      target: { value: '' }
    })
    fireEvent.change(screen.getByDisplayValue('2000'), {
      target: { value: '' }
    })

    fireEvent.click(screen.getByText('components.button.update'))

    await waitFor(() => {
      expect(message.error).toHaveBeenCalledWith(
        'pages.real_property.update_real_property.fill_all_fields'
      )
    })
  })

  it('returns early when propertyData or values are missing', async () => {
    const mockUpdateProperty = useProperties().updateProperty as jest.Mock

    render(
      <RealPropertyUpdate
        propertyData={{} as any}
        isModalUpdateOpen
        setIsModalUpdateOpen={mockSetIsModalUpdateOpen}
        setIsPropertyUpdated={mockSetIsPropertyUpdated}
      />
    )

    fireEvent.click(screen.getByText('components.button.update'))

    await waitFor(() => {
      expect(mockUpdateProperty).not.toHaveBeenCalled()
      expect(mockUpdateCache).not.toHaveBeenCalled()
      expect(mockSetIsModalUpdateOpen).not.toHaveBeenCalled()
      expect(message.success).not.toHaveBeenCalled()
      expect(mockSetIsPropertyUpdated).not.toHaveBeenCalled()
    })
  })
})
