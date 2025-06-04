import '@testing-library/jest-dom'
import { render, screen, fireEvent, waitFor } from '@testing-library/react'
import { act } from 'react'
import { MemoryRouter } from 'react-router-dom'
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

// Mock translations with a more comprehensive approach
jest.mock('react-i18next', () => ({
  useTranslation: () => ({
    t: (key: string) => {
      const translations = {
        'components.button.add_real_property': 'Add Real Property',
        'pages.real_property.error.error_fetching_data': 'Error fetching data',
        'components.messages.no_properties': 'No properties found',
        'components.switch.show_archived': 'Show Archived',
        'components.switch.show_active': 'Show Active',
        'components.select.surface.all': 'All Surfaces',
        'components.select.status.all': 'All Statuses',
        'components.select.status.available': 'Available',
        'components.select.status.unavailable': 'Unavailable',
        'components.select.status.invitation_sent': 'Invitation Sent',
        'pages.real_property.title': 'Real Properties'
      }
      return Object.prototype.hasOwnProperty.call(translations, key)
        ? (translations as Record<string, string>)[key]
        : key
    }
  })
}))

// Mock hooks with better type definitions
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

// Mock components with better implementations
jest.mock('@/views/RealProperty/create/RealPropertyCreate', () => ({
  __esModule: true,
  default: jest.fn(props => (
    <div
      data-testid="real-property-create-modal"
      onClick={() => props.setIsPropertyCreated(true)}
      onKeyDown={e => {
        if (e.key === 'Enter' || e.key === ' ') {
          props.setIsPropertyCreated(true)
        }
      }}
      role="button"
      tabIndex={0}
      aria-label="Real Property Create Modal"
    >
      RealPropertyCreate Mock
      <button type="button" onClick={() => props.setShowModalCreate(false)}>
        Close Modal
      </button>
      <button type="button" onClick={() => props.setIsPropertyCreated(true)}>
        Create Property
      </button>
    </div>
  ))
}))

jest.mock('@/components/ui/PageMeta/PageMeta', () => ({
  __esModule: true,
  default: jest.fn(({ title }) => <div data-testid="page-meta">{title}</div>)
}))

jest.mock('@/components/ui/PageText/Title', () => ({
  __esModule: true,
  default: jest.fn(({ title }) => <h1>{title}</h1>)
}))

jest.mock('@/components/ui/Loader/CardPropertyLoader', () => ({
  __esModule: true,
  default: jest.fn(({ cards }) => (
    <div data-testid="card-property-loader">Loading {cards} cards</div>
  ))
}))

jest.mock('@/components/features/RealProperty/PropertyFilterCard', () => ({
  __esModule: true,
  default: jest.fn(props => (
    <div data-testid="property-filter-card">
      <input
        data-testid="search-input"
        value={props.filters.searchQuery}
        onChange={e =>
          props.setFilters({ ...props.filters, searchQuery: e.target.value })
        }
      />
      <select
        data-testid="surface-select"
        value={props.filters.surfaceRange}
        onChange={e =>
          props.setFilters({ ...props.filters, surfaceRange: e.target.value })
        }
      >
        {props.surfaceRangeOptions.map(
          (option: { value: string; label: string }) => (
            <option key={option.value} value={option.value}>
              {option.label}
            </option>
          )
        )}
      </select>
      <select
        data-testid="status-select"
        value={props.filters.status}
        onChange={e =>
          props.setFilters({ ...props.filters, status: e.target.value })
        }
      >
        {props.statusOptions.map((option: { value: string; label: string }) => (
          <option key={option.value} value={option.value}>
            {option.label}
          </option>
        ))}
      </select>
    </div>
  ))
}))

jest.mock('@/components/features/RealProperty/PropertyCard', () => ({
  __esModule: true,
  default: jest.fn(({ realProperty }) => (
    <div
      data-testid={`property-card-${realProperty.id}`}
      className="property-card"
    >
      <h3>{realProperty.name}</h3>
      <p>
        {realProperty.apartment_number
          ? `N° ${realProperty.apartment_number} - ${realProperty.address}`
          : realProperty.address}
      </p>
      <p>{realProperty.city}</p>
      <p>{realProperty.status}</p>
    </div>
  ))
}))

describe('RealPropertyPage', () => {
  const mockProperties = [
    {
      id: '1',
      name: 'Test Property 1',
      address: '123 Main St',
      postal_code: '12345',
      city: 'Test City',
      country: 'Test Country',
      area_sqm: 75,
      status: 'available',
      apartment_number: '4B'
    },
    {
      id: '2',
      name: 'Test Property 2',
      address: '456 Oak Ave',
      postal_code: '67890',
      city: 'Another City',
      country: 'Test Country',
      area_sqm: 120,
      status: 'unavailable'
    },
    {
      id: '3',
      name: 'Test Property 3',
      address: '789 Pine St',
      postal_code: '54321',
      city: 'Third City',
      country: 'Another Country',
      area_sqm: 210,
      status: 'invitation_sent'
    }
  ]

  const mockGoToRealPropertyDetails = jest.fn()
  const mockRefreshProperties = jest.fn()

  beforeEach(() => {
    jest.clearAllMocks()

    // Setup default hook mock implementations
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

    // Make RealPropertyCreate a proper controllable component
    ;(RealPropertyCreate as jest.Mock).mockImplementation(
      ({ showModalCreate, setShowModalCreate, setIsPropertyCreated }) =>
        showModalCreate ? (
          <div
            data-testid="real-property-create-modal"
            role="dialog"
            aria-modal="true"
            aria-labelledby="create-property-modal-title"
          >
            <h2 id="create-property-modal-title">Create New Property</h2>
            <button type="button" onClick={() => setShowModalCreate(false)}>
              Close
            </button>
            <button
              type="button"
              onClick={() => {
                setIsPropertyCreated(true)
                setShowModalCreate(false)
              }}
            >
              Create
            </button>
          </div>
        ) : null
    )
  })

  it('renders the page title and header correctly', () => {
    render(
      <MemoryRouter>
        <RealPropertyPage />
      </MemoryRouter>
    )
    expect(screen.getByText('Real Properties')).toBeInTheDocument()
    expect(
      screen.getByRole('button', {
        name: 'Add Real Property'
      })
    ).toBeInTheDocument()
  })

  it('displays the switch for archived properties', () => {
    render(
      <MemoryRouter>
        <RealPropertyPage />
      </MemoryRouter>
    )
    expect(screen.getByText('Show Active')).toBeInTheDocument()

    // Test the switch functionality
    const switchElement = screen.getByRole('switch')
    expect(switchElement).toBeInTheDocument()

    // Initial state
    expect(useProperties).toHaveBeenLastCalledWith(null, false)

    // Click switch to show archived
    fireEvent.click(switchElement)

    // Re-render with new state
    expect(useProperties).toHaveBeenCalledTimes(2)
  })

  it('shows loading state when properties are loading', () => {
    ;(useProperties as jest.Mock).mockReturnValueOnce({
      properties: [],
      loading: true,
      error: null,
      refreshProperties: mockRefreshProperties
    })
    render(
      <MemoryRouter>
        <RealPropertyPage />
      </MemoryRouter>
    )
    expect(screen.getByTestId('card-property-loader')).toBeInTheDocument()
  })

  it('displays error message when there is an error fetching properties', () => {
    ;(useProperties as jest.Mock).mockReturnValueOnce({
      properties: [],
      loading: false,
      error: new Error('Fetch error'),
      refreshProperties: mockRefreshProperties
    })
    render(
      <MemoryRouter>
        <RealPropertyPage />
      </MemoryRouter>
    )
    expect(screen.getByText('Error fetching data')).toBeInTheDocument()
  })

  it('shows empty state when there are no properties', () => {
    ;(useProperties as jest.Mock).mockReturnValueOnce({
      properties: [],
      loading: false,
      error: null,
      refreshProperties: mockRefreshProperties
    })
    render(
      <MemoryRouter>
        <RealPropertyPage />
      </MemoryRouter>
    )
    expect(screen.getByText('No properties found')).toBeInTheDocument()
  })

  it('renders property cards when properties exist', () => {
    render(
      <MemoryRouter>
        <RealPropertyPage />
      </MemoryRouter>
    )
    expect(screen.getByText('Test Property 1')).toBeInTheDocument()
    expect(screen.getByText('Test Property 2')).toBeInTheDocument()
    expect(screen.getByText('Test Property 3')).toBeInTheDocument()
  })

  it('filters properties by search query', async () => {
    render(
      <MemoryRouter>
        <RealPropertyPage />
      </MemoryRouter>
    )

    // Initial state should show all properties
    expect(screen.getByText('Test Property 1')).toBeInTheDocument()
    expect(screen.getByText('Test Property 2')).toBeInTheDocument()
    expect(screen.getByText('Test Property 3')).toBeInTheDocument()

    // Filter by property name
    const searchInput = screen.getByTestId('search-input')
    fireEvent.change(searchInput, { target: { value: 'Property 1' } })

    // Only Property 1 should be visible
    expect(screen.getByText('Test Property 1')).toBeInTheDocument()
    expect(screen.queryByText('Test Property 2')).not.toBeInTheDocument()
    expect(screen.queryByText('Test Property 3')).not.toBeInTheDocument()
  })

  it('filters properties by surface range', () => {
    render(
      <MemoryRouter>
        <RealPropertyPage />
      </MemoryRouter>
    )

    // Initial state should show all properties
    expect(screen.getByText('Test Property 1')).toBeInTheDocument() // 75 sqm
    expect(screen.getByText('Test Property 2')).toBeInTheDocument() // 120 sqm
    expect(screen.getByText('Test Property 3')).toBeInTheDocument() // 210 sqm

    // Filter by surface range 51-100
    const surfaceSelect = screen.getByTestId('surface-select')
    fireEvent.change(surfaceSelect, { target: { value: '51-100' } })

    // Only Property 1 should be visible
    expect(screen.getByText('Test Property 1')).toBeInTheDocument()
    expect(screen.queryByText('Test Property 2')).not.toBeInTheDocument()
    expect(screen.queryByText('Test Property 3')).not.toBeInTheDocument()

    // Filter by surface range 101-150
    fireEvent.change(surfaceSelect, { target: { value: '101-150' } })

    // Only Property 2 should be visible
    expect(screen.queryByText('Test Property 1')).not.toBeInTheDocument()
    expect(screen.getByText('Test Property 2')).toBeInTheDocument()
    expect(screen.queryByText('Test Property 3')).not.toBeInTheDocument()

    // Filter by surface range 201+
    fireEvent.change(surfaceSelect, { target: { value: '201+' } })

    // Only Property 3 should be visible
    expect(screen.queryByText('Test Property 1')).not.toBeInTheDocument()
    expect(screen.queryByText('Test Property 2')).not.toBeInTheDocument()
    expect(screen.getByText('Test Property 3')).toBeInTheDocument()
  })

  it('filters properties by status', () => {
    render(
      <MemoryRouter>
        <RealPropertyPage />
      </MemoryRouter>
    )

    // Initial state should show all properties
    expect(screen.getByText('Test Property 1')).toBeInTheDocument() // available
    expect(screen.getByText('Test Property 2')).toBeInTheDocument() // unavailable
    expect(screen.getByText('Test Property 3')).toBeInTheDocument() // invitation_sent

    // Filter by status available
    const statusSelect = screen.getByTestId('status-select')
    fireEvent.change(statusSelect, { target: { value: 'available' } })

    // Only Property 1 should be visible
    expect(screen.getByText('Test Property 1')).toBeInTheDocument()
    expect(screen.queryByText('Test Property 2')).not.toBeInTheDocument()
    expect(screen.queryByText('Test Property 3')).not.toBeInTheDocument()

    // Filter by status unavailable
    fireEvent.change(statusSelect, { target: { value: 'unavailable' } })

    // Only Property 2 should be visible
    expect(screen.queryByText('Test Property 1')).not.toBeInTheDocument()
    expect(screen.getByText('Test Property 2')).toBeInTheDocument()
    expect(screen.queryByText('Test Property 3')).not.toBeInTheDocument()
  })

  it('opens the create modal when add button is clicked', () => {
    render(
      <MemoryRouter>
        <RealPropertyPage />
      </MemoryRouter>
    )

    // Modal should not be visible initially
    expect(
      screen.queryByTestId('real-property-create-modal')
    ).not.toBeInTheDocument()

    // Click the add button
    fireEvent.click(screen.getByRole('button', { name: 'Add Real Property' }))

    // Modal should now be visible
    expect(screen.getByTestId('real-property-create-modal')).toBeInTheDocument()

    // Close the modal
    fireEvent.click(screen.getByRole('button', { name: 'Close' }))

    // Modal should be hidden again
    expect(
      screen.queryByTestId('real-property-create-modal')
    ).not.toBeInTheDocument()
  })

  it('refreshes properties after a new property is created', async () => {
    render(
      <MemoryRouter>
        <RealPropertyPage />
      </MemoryRouter>
    )

    // Open the modal
    fireEvent.click(screen.getByRole('button', { name: 'Add Real Property' }))

    // Create a new property
    fireEvent.click(screen.getByRole('button', { name: 'Create' }))

    // Wait for the refreshProperties to be called
    await waitFor(() => {
      expect(mockRefreshProperties).toHaveBeenCalled()
    })

    // Modal should be closed
    expect(
      screen.queryByTestId('real-property-create-modal')
    ).not.toBeInTheDocument()
  })

  it('correctly handles the archived property toggle', async () => {
    render(
      <MemoryRouter>
        <RealPropertyPage />
      </MemoryRouter>
    )

    // Initial call should be with showArchived=false
    expect(useProperties).toHaveBeenLastCalledWith(null, false)

    // Toggle switch
    const switchElement = screen.getByRole('switch')

    // Act inside an act block to handle state updates
    await act(async () => {
      fireEvent.click(switchElement)
    })

    // Re-render the component after state update
    render(
      <MemoryRouter>
        <RealPropertyPage />
      </MemoryRouter>
    )

    // After toggle, the useProperties hook should be called with showArchived=true
    expect(useProperties).toHaveBeenCalledWith(null, true)
  })

  it('displays the correct address format with apartment number', () => {
    render(
      <MemoryRouter>
        <RealPropertyPage />
      </MemoryRouter>
    )
    expect(screen.getByText(/N° 4B - 123 Main St/)).toBeInTheDocument()
  })

  it('displays the correct address format without apartment number', () => {
    ;(useProperties as jest.Mock).mockReturnValueOnce({
      properties: [
        {
          ...mockProperties[1] // Property without apartment number
        }
      ],
      loading: false,
      error: null,
      refreshProperties: mockRefreshProperties
    })
    render(
      <MemoryRouter>
        <RealPropertyPage />
      </MemoryRouter>
    )
    expect(screen.getByText('456 Oak Ave')).toBeInTheDocument()
    expect(screen.queryByText(/N° .* - 456 Oak Ave/)).not.toBeInTheDocument()
  })
})
