import '@testing-library/jest-dom'
import { render, screen, fireEvent, waitFor } from '@testing-library/react'
import RealPropertyPage from '@/views/RealProperty/RealProperty'
import useProperties from '@/hooks/Property/useProperties'
import useNavigation from '@/hooks/Navigation/useNavigation'
import useImageCache from '@/hooks/Image/useImageCache'
import RealPropertyCreate from '@/views/RealProperty/create/RealPropertyCreate'

// Mock TenantStatusEnum
jest.mock('@/enums/PropertyEnum', () => ({
  TenantStatusEnum: {
    active: { text: 'Active', color: 'green' },
    archived: { text: 'Archived', color: 'red' }
  }
}))

// Mock other dependencies
jest.mock('react-i18next', () => ({
  useTranslation: () => ({
    t: (key: string) => {
      if (key === 'components.button.add_real_property')
        return 'Add Real Property'
      return key
    }
  })
}))

jest.mock('@/hooks/Property/useProperties', () => ({
  __esModule: true,
  default: jest.fn()
}))

jest.mock('@/hooks/Navigation/useNavigation', () => ({
  __esModule: true,
  default: jest.fn()
}))

jest.mock('@/hooks/Image/useImageCache', () => ({
  __esModule: true,
  default: jest.fn()
}))

jest.mock('@/services/api/Owner/Properties/GetPropertyPicture', () => ({
  __esModule: true,
  default: jest.fn()
}))

jest.mock('@/views/RealProperty/create/RealPropertyCreate', () => ({
  __esModule: true,
  default: jest.fn(() => <div>RealPropertyCreate Mock</div>)
}))

// Mock components
jest.mock('@/components/ui/PageMeta/PageMeta', () => ({
  __esModule: true,
  default: jest.fn(() => null)
}))

jest.mock('@/components/ui/PageText/Title', () => ({
  __esModule: true,
  default: jest.fn(() => <h1>PageTitle Mock</h1>)
}))

jest.mock('@/components/ui/Loader/CardPropertyLoader', () => ({
  __esModule: true,
  default: jest.fn(() => <div>CardPropertyLoader Mock</div>)
}))

describe('RealPropertyPage', () => {
  const mockProperties = [
    {
      id: '1',
      name: 'Test Property 1',
      address: '123 Main St',
      postal_code: '12345',
      city: 'Test City',
      status: 'active',
      apartment_number: '4B'
    },
    {
      id: '2',
      name: 'Test Property 2',
      address: '456 Oak Ave',
      postal_code: '67890',
      city: 'Another City',
      status: 'archived'
    }
  ]

  const mockGoToRealPropertyDetails = jest.fn()
  const mockRefreshProperties = jest.fn()

  beforeEach(() => {
    ;(useProperties as jest.Mock).mockReturnValue({
      properties: mockProperties,
      loading: false,
      error: null,
      refreshProperties: mockRefreshProperties
    })
    ;(useNavigation as jest.Mock).mockReturnValue({
      goToRealPropertyDetails: mockGoToRealPropertyDetails
    })
    ;(useImageCache as jest.Mock).mockReturnValue({
      data: 'mocked-image-url',
      isLoading: false
    })
    ;(RealPropertyCreate as jest.Mock).mockImplementation(
      ({ showModalCreate }) =>
        showModalCreate ? <div>RealPropertyCreate Mock</div> : null
    )
  })

  afterEach(() => {
    jest.clearAllMocks()
  })

  it('renders the page title and header correctly', () => {
    render(<RealPropertyPage />)
    expect(screen.getByText('PageTitle Mock')).toBeInTheDocument()
    expect(
      screen.getByRole('button', {
        name: 'Add Real Property'
      })
    ).toBeInTheDocument()
  })

  it('displays the switch for archived properties', () => {
    render(<RealPropertyPage />)
    expect(
      screen.getByText('components.switch.show_active')
    ).toBeInTheDocument()
  })

  it('shows loading state when properties are loading', () => {
    ;(useProperties as jest.Mock).mockReturnValueOnce({
      properties: [],
      loading: true,
      error: null,
      refreshProperties: mockRefreshProperties
    })
    render(<RealPropertyPage />)
    expect(screen.getByText('CardPropertyLoader Mock')).toBeInTheDocument()
  })

  it('displays error message when there is an error fetching properties', () => {
    ;(useProperties as jest.Mock).mockReturnValueOnce({
      properties: [],
      loading: false,
      error: new Error('Fetch error'),
      refreshProperties: mockRefreshProperties
    })
    render(<RealPropertyPage />)
    expect(
      screen.getByText('pages.real_property.error.error_fetching_data')
    ).toBeInTheDocument()
  })

  it('shows empty state when there are no properties', () => {
    ;(useProperties as jest.Mock).mockReturnValueOnce({
      properties: [],
      loading: false,
      error: null,
      refreshProperties: mockRefreshProperties
    })
    render(<RealPropertyPage />)
    expect(
      screen.getByText('components.messages.no_properties')
    ).toBeInTheDocument()
  })

  it('renders property cards when properties exist', () => {
    render(<RealPropertyPage />)
    expect(screen.getByText('Test Property 1')).toBeInTheDocument()
    expect(screen.getByText('Test Property 2')).toBeInTheDocument()
  })

  it('navigates to property details when a card is clicked', () => {
    render(<RealPropertyPage />)
    fireEvent.click(screen.getByText('Test Property 1'))
    expect(mockGoToRealPropertyDetails).toHaveBeenCalledWith('1')
  })

  it('opens the create modal when add button is clicked', () => {
    render(<RealPropertyPage />)
    fireEvent.click(
      screen.getByRole('button', {
        name: 'components.button.add_real_property'
      })
    )
    expect(screen.getByText('RealPropertyCreate Mock')).toBeInTheDocument()
  })

  // it('toggles archived properties when switch is clicked', () => {
  //   const { rerender } = render(<RealPropertyPage />)
  //   const switchElement = screen.getByRole('switch')

  //   // Initial call
  //   expect(useProperties).toHaveBeenCalledWith(null, false)

  //   // Toggle switch
  //   fireEvent.click(switchElement)

  //   // Verify the hook was called with showArchived=true
  //   expect(useProperties).toHaveBeenCalledWith(null, true)
  // })

  it('displays the correct address format with apartment number', () => {
    render(<RealPropertyPage />)
    expect(screen.getByText(/N° 4B - 123 Main St/)).toBeInTheDocument()
  })

  it('displays the correct address format without apartment number', () => {
    ;(useProperties as jest.Mock).mockReturnValueOnce({
      properties: [
        {
          ...mockProperties[0],
          apartment_number: undefined
        }
      ],
      loading: false,
      error: null,
      refreshProperties: mockRefreshProperties
    })
    render(<RealPropertyPage />)
    expect(screen.getByText(/123 Main St/)).toBeInTheDocument()
  })

  it('refreshes properties after a new property is created', async () => {
    ;(RealPropertyCreate as jest.Mock).mockImplementationOnce(
      ({ setIsPropertyCreated }) => {
        setTimeout(() => {
          setIsPropertyCreated(true)
        }, 0)
        return <div>RealPropertyCreate Mock</div>
      }
    )

    fireEvent.click(
      screen.getByRole('button', {
        name: 'components.button.add_real_property'
      })
    )

    await waitFor(() => {
      expect(mockRefreshProperties).toHaveBeenCalled()
    })
  })
})
