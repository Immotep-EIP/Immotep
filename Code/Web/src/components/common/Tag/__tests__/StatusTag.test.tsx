import { render, screen } from '@testing-library/react'
import '@testing-library/jest-dom'
import StatusTag from '../StatusTag'

// Mock react-i18next
const mockT = jest.fn()
jest.mock('react-i18next', () => ({
  useTranslation: () => ({
    t: mockT
  })
}))

describe('StatusTag Component', () => {
  beforeEach(() => {
    jest.clearAllMocks()
  })

  afterEach(() => {
    document.body.innerHTML = ''
  })

  describe('basic rendering', () => {
    it('renders with value and colorMap', () => {
      const colorMap = { active: 'green', inactive: 'red' }

      render(<StatusTag value="active" colorMap={colorMap} />)

      expect(screen.getByText('active')).toBeInTheDocument()
      const tag = screen.getByText('active').closest('.ant-tag')
      expect(tag).toHaveClass('ant-tag-green')
    })

    it('renders with uppercase value converted to lowercase for colorMap lookup', () => {
      const colorMap = { active: 'green', inactive: 'red' }

      render(<StatusTag value="ACTIVE" colorMap={colorMap} />)

      expect(screen.getByText('ACTIVE')).toBeInTheDocument()
      const tag = screen.getByText('ACTIVE').closest('.ant-tag')
      expect(tag).toHaveClass('ant-tag-green')
    })

    it('uses defaultColor when value not found in colorMap', () => {
      const colorMap = { active: 'green' }

      render(<StatusTag value="unknown" colorMap={colorMap} />)

      expect(screen.getByText('unknown')).toBeInTheDocument()
      const tag = screen.getByText('unknown').closest('.ant-tag')
      expect(tag).toHaveClass('ant-tag-default')
    })

    it('uses custom defaultColor when provided', () => {
      const colorMap = { active: 'green' }

      render(
        <StatusTag value="unknown" colorMap={colorMap} defaultColor="blue" />
      )

      expect(screen.getByText('unknown')).toBeInTheDocument()
      const tag = screen.getByText('unknown').closest('.ant-tag')
      expect(tag).toHaveClass('ant-tag-blue')
    })
  })

  describe('translation with i18nPrefix', () => {
    it('translates value when i18nPrefix is provided and translation exists', () => {
      const colorMap = { active: 'green' }
      mockT.mockReturnValue('Actif')

      render(
        <StatusTag value="active" colorMap={colorMap} i18nPrefix="status" />
      )

      expect(mockT).toHaveBeenCalledWith('status.active')
      expect(screen.getByText('Actif')).toBeInTheDocument()
    })

    it('uses original value when translation does not exist (starts with prefix)', () => {
      const colorMap = { unknown: 'gray' }
      mockT.mockReturnValue('status.unknown') // Translation not found, returns key

      render(
        <StatusTag value="unknown" colorMap={colorMap} i18nPrefix="status" />
      )

      expect(mockT).toHaveBeenCalledWith('status.unknown')
      expect(screen.getByText('unknown')).toBeInTheDocument()
    })

    it('handles mixed case value with translation', () => {
      const colorMap = { pending: 'orange' }
      mockT.mockReturnValue('En attente')

      render(
        <StatusTag value="PENDING" colorMap={colorMap} i18nPrefix="status" />
      )

      expect(mockT).toHaveBeenCalledWith('status.pending')
      expect(screen.getByText('En attente')).toBeInTheDocument()
      const tag = screen.getByText('En attente').closest('.ant-tag')
      expect(tag).toHaveClass('ant-tag-orange')
    })

    it('works without i18nPrefix', () => {
      const colorMap = { active: 'green' }

      render(<StatusTag value="active" colorMap={colorMap} />)

      expect(mockT).not.toHaveBeenCalled()
      expect(screen.getByText('active')).toBeInTheDocument()
    })
  })

  describe('edge cases', () => {
    it('handles empty string value', () => {
      const colorMap = { '': 'gray' }

      const { container } = render(<StatusTag value="" colorMap={colorMap} />)

      const tag = container.querySelector('.ant-tag')
      expect(tag).toBeInTheDocument()
      expect(tag).toHaveClass('ant-tag-has-color')
      expect(tag).toHaveStyle('background-color: gray')
    })

    it('handles undefined value', () => {
      const colorMap = { active: 'green' }

      const { container } = render(
        <StatusTag value={undefined as any} colorMap={colorMap} />
      )

      const tag = container.querySelector('.ant-tag')
      expect(tag).toBeInTheDocument()
      expect(tag).toHaveClass('ant-tag-default')
    })

    it('handles null value', () => {
      const colorMap = { active: 'green' }

      const { container } = render(
        <StatusTag value={null as any} colorMap={colorMap} />
      )

      const tag = container.querySelector('.ant-tag')
      expect(tag).toBeInTheDocument()
      expect(tag).toHaveClass('ant-tag-default')
    })

    it('handles undefined colorMap', () => {
      render(<StatusTag value="active" colorMap={undefined as any} />)

      expect(screen.getByText('active')).toBeInTheDocument()
      const tag = screen.getByText('active').closest('.ant-tag')
      expect(tag).toHaveClass('ant-tag-default')
    })

    it('handles null colorMap', () => {
      render(<StatusTag value="active" colorMap={null as any} />)

      expect(screen.getByText('active')).toBeInTheDocument()
      const tag = screen.getByText('active').closest('.ant-tag')
      expect(tag).toHaveClass('ant-tag-default')
    })

    it('handles empty colorMap', () => {
      render(<StatusTag value="active" colorMap={{}} />)

      expect(screen.getByText('active')).toBeInTheDocument()
      const tag = screen.getByText('active').closest('.ant-tag')
      expect(tag).toHaveClass('ant-tag-default')
    })
  })

  describe('complex scenarios', () => {
    it('handles special characters in value', () => {
      const colorMap = { 'in-progress': 'yellow', on_hold: 'orange' }

      render(<StatusTag value="in-progress" colorMap={colorMap} />)

      expect(screen.getByText('in-progress')).toBeInTheDocument()
      const tag = screen.getByText('in-progress').closest('.ant-tag')
      expect(tag).toHaveClass('ant-tag-yellow')
    })

    it('handles numeric values as strings', () => {
      const colorMap = { '1': 'green', '0': 'red' }

      render(<StatusTag value="1" colorMap={colorMap} />)

      expect(screen.getByText('1')).toBeInTheDocument()
      const tag = screen.getByText('1').closest('.ant-tag')
      expect(tag).toHaveClass('ant-tag-green')
    })

    it('translates with special characters in prefix', () => {
      const colorMap = { active: 'green' }
      mockT.mockReturnValue('Status Actif')

      render(
        <StatusTag
          value="active"
          colorMap={colorMap}
          i18nPrefix="custom.status"
        />
      )

      expect(mockT).toHaveBeenCalledWith('custom.status.active')
      expect(screen.getByText('Status Actif')).toBeInTheDocument()
    })

    it('handles case where translation returns a string starting with prefix', () => {
      const colorMap = { test: 'blue' }
      mockT.mockReturnValue('prefix.test') // i18next returns the key when no translation is found

      render(<StatusTag value="test" colorMap={colorMap} i18nPrefix="prefix" />)

      expect(mockT).toHaveBeenCalledWith('prefix.test')
      // When translation returns a string starting with prefix, it means no translation was found
      // so the original value should be displayed
      expect(screen.getByText('test')).toBeInTheDocument()
    })

    it('combines all props correctly', () => {
      const colorMap = {
        confirmed: 'green',
        pending: 'orange',
        cancelled: 'red'
      }
      mockT.mockReturnValue('Confirmé')

      render(
        <StatusTag
          value="CONFIRMED"
          colorMap={colorMap}
          i18nPrefix="booking.status"
          defaultColor="purple"
        />
      )

      expect(mockT).toHaveBeenCalledWith('booking.status.confirmed')
      expect(screen.getByText('Confirmé')).toBeInTheDocument()
      const tag = screen.getByText('Confirmé').closest('.ant-tag')
      expect(tag).toHaveClass('ant-tag-green')
    })
  })

  describe('colorMap lookup edge cases', () => {
    it('handles colorMap with falsy values', () => {
      const colorMap = {
        active: 'green',
        inactive: '',
        disabled: null as any,
        unknown: undefined as any
      }

      render(<StatusTag value="inactive" colorMap={colorMap} />)

      const tag = screen.getByText('inactive').closest('.ant-tag')
      expect(tag).toHaveClass('ant-tag-default') // Falls back to default when color is falsy
    })

    it('preserves original value casing in display when no translation', () => {
      const colorMap = { mixed_case: 'blue' }

      render(<StatusTag value="Mixed_Case" colorMap={colorMap} />)

      expect(screen.getByText('Mixed_Case')).toBeInTheDocument()
      const tag = screen.getByText('Mixed_Case').closest('.ant-tag')
      expect(tag).toHaveClass('ant-tag-blue')
    })
  })
})
