import '@testing-library/jest-dom'
import { render, screen, fireEvent, waitFor } from '@testing-library/react'
import { message } from 'antd'
import RealPropertyCreate from '@/views/RealProperty/create/RealPropertyCreate'
import useProperties from '@/hooks/Property/useProperties'
import useImageUpload from '@/hooks/Image/useImageUpload'

jest.mock('@/hooks/Property/useProperties')
jest.mock('@/hooks/Image/useImageUpload')
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

describe('RealPropertyCreate Component', () => {
  const mockSetShowModalCreate = jest.fn()
  const mockSetIsPropertyCreated = jest.fn()

  beforeEach(() => {
    jest.clearAllMocks()
    ;(useProperties as jest.Mock).mockReturnValue({
      loading: false,
      createProperty: jest.fn()
    })
    ;(useImageUpload as jest.Mock).mockReturnValue({
      uploadProps: {},
      imageBase64: 'mocked-image-base64',
      resetImage: jest.fn()
    })
  })

  const renderComponent = () =>
    render(
      <RealPropertyCreate
        showModalCreate
        setShowModalCreate={mockSetShowModalCreate}
        setIsPropertyCreated={mockSetIsPropertyCreated}
      />
    )

  it('renders the modal with form fields', () => {
    renderComponent()

    expect(
      screen.getByText('pages.real_property.add_real_property.document_title')
    ).toBeInTheDocument()
    expect(
      screen.getByText('components.input.property_name.label')
    ).toBeInTheDocument()
    expect(
      screen.getByText('components.input.apartment_number.label')
    ).toBeInTheDocument()
    expect(
      screen.getByText('components.input.address.label')
    ).toBeInTheDocument()
    expect(
      screen.getByText('components.input.zip_code.label')
    ).toBeInTheDocument()
    expect(screen.getByText('components.input.city.label')).toBeInTheDocument()
    expect(
      screen.getByText('components.input.country.label')
    ).toBeInTheDocument()
    expect(screen.getByText('components.input.area.label')).toBeInTheDocument()
    expect(
      screen.getByText('components.input.rental.label')
    ).toBeInTheDocument()
    expect(
      screen.getByText('components.input.deposit.label')
    ).toBeInTheDocument()
    expect(
      screen.getByText('components.input.picture.label')
    ).toBeInTheDocument()
  })

  it('closes the modal when cancel button is clicked', () => {
    renderComponent()

    fireEvent.click(screen.getByText('components.button.cancel'))

    expect(mockSetShowModalCreate).toHaveBeenCalledWith(false)
  })

  it('submits the form and creates the property successfully', async () => {
    const mockCreateProperty = useProperties().createProperty as jest.Mock
    mockCreateProperty.mockResolvedValueOnce({})

    renderComponent()

    // Remplir les champs du formulaire
    fireEvent.change(
      screen.getByLabelText('components.input.property_name.label'),
      {
        target: { value: 'Test Property' }
      }
    )
    fireEvent.change(
      screen.getByLabelText('components.input.apartment_number.label'),
      {
        target: { value: '4B' }
      }
    )
    fireEvent.change(screen.getByLabelText('components.input.address.label'), {
      target: { value: '123 Main St' }
    })
    fireEvent.change(screen.getByLabelText('components.input.zip_code.label'), {
      target: { value: '12345' }
    })
    fireEvent.change(screen.getByLabelText('components.input.city.label'), {
      target: { value: 'Test City' }
    })
    fireEvent.change(screen.getByLabelText('components.input.country.label'), {
      target: { value: 'Test Country' }
    })
    fireEvent.change(screen.getByLabelText('components.input.area.label'), {
      target: { value: 100 }
    })
    fireEvent.change(screen.getByLabelText('components.input.rental.label'), {
      target: { value: 1000 }
    })
    fireEvent.change(screen.getByLabelText('components.input.deposit.label'), {
      target: { value: 2000 }
    })

    fireEvent.click(screen.getByText('components.button.add'))

    await waitFor(() => {
      expect(mockCreateProperty).toHaveBeenCalledWith(
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
        'mocked-image-base64'
      )
      expect(mockSetShowModalCreate).toHaveBeenCalledWith(false)
      expect(message.success).toHaveBeenCalledWith(
        'pages.real_property.add_real_property.property_created'
      )
      expect(mockSetIsPropertyCreated).toHaveBeenCalledWith(true)
    })
  })

  it('shows an error message when form submission fails', async () => {
    const mockCreateProperty = useProperties().createProperty as jest.Mock
    mockCreateProperty.mockRejectedValueOnce(new Error('Create failed'))

    renderComponent()

    // Remplir les champs du formulaire
    fireEvent.change(
      screen.getByLabelText('components.input.property_name.label'),
      {
        target: { value: 'Test Property' }
      }
    )
    fireEvent.change(
      screen.getByLabelText('components.input.apartment_number.label'),
      {
        target: { value: '4B' }
      }
    )
    fireEvent.change(screen.getByLabelText('components.input.address.label'), {
      target: { value: '123 Main St' }
    })
    fireEvent.change(screen.getByLabelText('components.input.zip_code.label'), {
      target: { value: '12345' }
    })
    fireEvent.change(screen.getByLabelText('components.input.city.label'), {
      target: { value: 'Test City' }
    })
    fireEvent.change(screen.getByLabelText('components.input.country.label'), {
      target: { value: 'Test Country' }
    })
    fireEvent.change(screen.getByLabelText('components.input.area.label'), {
      target: { value: 100 }
    })
    fireEvent.change(screen.getByLabelText('components.input.rental.label'), {
      target: { value: 1000 }
    })
    fireEvent.change(screen.getByLabelText('components.input.deposit.label'), {
      target: { value: 2000 }
    })

    fireEvent.click(screen.getByText('components.button.add'))

    await waitFor(() => {
      expect(message.error).toHaveBeenCalledWith(
        'pages.real_property.add_real_property.error_property_created'
      )
    })
  })

  it('shows an error message when form fields are not filled', async () => {
    renderComponent()

    fireEvent.click(screen.getByText('components.button.add'))

    await waitFor(() => {
      expect(message.error).toHaveBeenCalledWith(
        'pages.real_property.add_real_property.fill_all_fields'
      )
    })
  })
})
