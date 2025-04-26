import React from 'react'
import '@testing-library/jest-dom'
import { render, screen, fireEvent, waitFor } from '@testing-library/react'
import { message } from 'antd'
import RealPropertyUpdate from '@/views/RealProperty/update/RealPropertyUpdate' // Corrected import path
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
    deposit_price: 2000
  }

  const mockSetIsModalUpdateOpen = jest.fn()
  const mockSetIsPropertyUpdated = jest.fn()

  beforeEach(() => {
    jest.clearAllMocks()
    ;(useProperties as jest.Mock).mockReturnValue({
      loading: false,
      updateProperty: jest.fn()
    })
    ;(useImageUpload as jest.Mock).mockReturnValue({
      uploadProps: {},
      imageBase64: 'mocked-image-base64'
    })
    ;(useImageCache as jest.Mock).mockReturnValue({
      updateCache: jest.fn()
    })
  })

  const renderComponent = () =>
    render(
      <RealPropertyUpdate
        propertyData={mockPropertyData}
        isModalUpdateOpen
        setIsModalUpdateOpen={mockSetIsModalUpdateOpen}
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
      expect(useImageCache().updateCache).toHaveBeenCalledWith(
        'mocked-image-base64'
      )
      expect(mockSetIsModalUpdateOpen).toHaveBeenCalledWith(false)
      expect(message.success).toHaveBeenCalledWith(
        'pages.real_property.update_real_property.property_updated'
      )
      expect(mockSetIsPropertyUpdated).toHaveBeenCalledWith(true)
    })
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

    // Clear all form fields
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
})
