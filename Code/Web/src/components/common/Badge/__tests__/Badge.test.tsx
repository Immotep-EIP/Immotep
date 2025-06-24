import { render, screen } from '@testing-library/react'
import '@testing-library/jest-dom'
import { Badge as AntBadge } from 'antd'
import Badge from '../Badge'

describe('Badge Component', () => {
  afterEach(() => {
    document.body.innerHTML = ''
  })

  describe('basic rendering', () => {
    it('renders children correctly', () => {
      render(
        <Badge>
          <div>Badge content</div>
        </Badge>
      )

      expect(screen.getByText('Badge content')).toBeInTheDocument()
    })

    it('renders without children', () => {
      const { container } = render(<Badge count={5} />)

      const badgeElement = container.querySelector('.ant-badge')
      expect(badgeElement).toBeInTheDocument()
    })

    it('renders with count prop', () => {
      render(
        <Badge count={10}>
          <div>Content</div>
        </Badge>
      )

      expect(screen.getByTitle('10')).toBeInTheDocument()
      expect(screen.getByText('Content')).toBeInTheDocument()
    })

    it('renders with dot prop', () => {
      const { container } = render(
        <Badge dot>
          <div>Content</div>
        </Badge>
      )

      expect(screen.getByText('Content')).toBeInTheDocument()
      const dotBadge = container.querySelector('.ant-badge-dot')
      expect(dotBadge).toBeInTheDocument()
    })

    it('renders with text prop', () => {
      render(<Badge status="success" text="Badge text" />)

      expect(screen.getByText('Badge text')).toBeInTheDocument()
    })
  })

  describe('status badges', () => {
    it('renders with success status', () => {
      const { container } = render(<Badge status="success" text="Success" />)

      expect(screen.getByText('Success')).toBeInTheDocument()
      const statusBadge = container.querySelector('.ant-badge-status-success')
      expect(statusBadge).toBeInTheDocument()
    })

    it('renders with error status', () => {
      const { container } = render(<Badge status="error" text="Error" />)

      expect(screen.getByText('Error')).toBeInTheDocument()
      const statusBadge = container.querySelector('.ant-badge-status-error')
      expect(statusBadge).toBeInTheDocument()
    })

    it('renders with warning status', () => {
      const { container } = render(<Badge status="warning" text="Warning" />)

      expect(screen.getByText('Warning')).toBeInTheDocument()
      const statusBadge = container.querySelector('.ant-badge-status-warning')
      expect(statusBadge).toBeInTheDocument()
    })

    it('renders with processing status', () => {
      const { container } = render(
        <Badge status="processing" text="Processing" />
      )

      expect(screen.getByText('Processing')).toBeInTheDocument()
      const statusBadge = container.querySelector(
        '.ant-badge-status-processing'
      )
      expect(statusBadge).toBeInTheDocument()
    })

    it('renders with default status', () => {
      const { container } = render(<Badge status="default" text="Default" />)

      expect(screen.getByText('Default')).toBeInTheDocument()
      const statusBadge = container.querySelector('.ant-badge-status-default')
      expect(statusBadge).toBeInTheDocument()
    })
  })

  describe('color and styling', () => {
    it('renders with custom color', () => {
      const { container } = render(
        <Badge count={1} color="red">
          <div>Content</div>
        </Badge>
      )

      const badge = container.querySelector('.ant-badge-count')
      expect(badge).toHaveStyle('background: rgb(245, 34, 45)')
    })

    it('renders with custom className', () => {
      const { container } = render(
        <Badge className="custom-badge" count={5}>
          <div>Content</div>
        </Badge>
      )

      const badgeWrapper = container.querySelector('.custom-badge')
      expect(badgeWrapper).toBeInTheDocument()
    })

    it('renders with custom style', () => {
      const customStyle = { fontSize: '16px' }
      const { container } = render(
        <Badge style={customStyle} count={3}>
          <div>Content</div>
        </Badge>
      )

      const badge = container.querySelector('.ant-badge')
      expect(badge).toBeInTheDocument()

      expect(screen.getByText('Content')).toBeInTheDocument()
    })

    it('renders with size prop', () => {
      const { container } = render(
        <Badge count={8} size="small">
          <div>Content</div>
        </Badge>
      )

      const badge = container.querySelector('.ant-badge-count')
      expect(badge).toHaveClass('ant-badge-count-sm')
    })
  })

  describe('overflowCount and showZero', () => {
    it('shows overflow count correctly', () => {
      render(
        <Badge count={100} overflowCount={99}>
          <div>Content</div>
        </Badge>
      )

      expect(screen.getByTitle('100')).toBeInTheDocument()
      expect(screen.getByText('99+')).toBeInTheDocument()
    })

    it('handles showZero prop', () => {
      render(
        <Badge count={0} showZero>
          <div>Content</div>
        </Badge>
      )

      expect(screen.getByTitle('0')).toBeInTheDocument()
      expect(screen.getByText('0')).toBeInTheDocument()
    })

    it('hides zero count when showZero is false', () => {
      const { container } = render(
        <Badge count={0} showZero={false}>
          <div>Content</div>
        </Badge>
      )

      const countBadge = container.querySelector('.ant-badge-count')
      expect(countBadge).not.toBeInTheDocument()
    })
  })

  describe('offset positioning', () => {
    it('renders with custom offset', () => {
      const { container } = render(
        <Badge count={5} offset={[10, 10]}>
          <div>Content</div>
        </Badge>
      )

      const badge = container.querySelector('.ant-badge-count')
      expect(badge).toHaveStyle('right: -10px; margin-top: 10px')
    })
  })

  describe('Badge.Ribbon static property', () => {
    it('exposes Ribbon component from AntBadge', () => {
      expect(Badge.Ribbon).toBe(AntBadge.Ribbon)
    })

    it('renders Badge.Ribbon correctly', () => {
      render(
        <Badge.Ribbon text="Ribbon Text">
          <div data-testid="ribbon-content">Ribbon Content</div>
        </Badge.Ribbon>
      )

      expect(screen.getByText('Ribbon Text')).toBeInTheDocument()
      expect(screen.getByTestId('ribbon-content')).toBeInTheDocument()
    })

    it('renders Badge.Ribbon with color', () => {
      const { container } = render(
        <Badge.Ribbon text="Ribbon" color="blue">
          <div>Content</div>
        </Badge.Ribbon>
      )

      const ribbon = container.querySelector('.ant-ribbon')
      expect(ribbon).toHaveStyle('background: rgb(22, 119, 255)')
    })

    it('renders Badge.Ribbon with placement', () => {
      const { container } = render(
        <Badge.Ribbon text="Ribbon" placement="start">
          <div>Content</div>
        </Badge.Ribbon>
      )

      const ribbon = container.querySelector('.ant-ribbon-placement-start')
      expect(ribbon).toBeInTheDocument()
    })
  })

  describe('edge cases', () => {
    it('handles undefined children', () => {
      const { container } = render(<Badge count={5}>{undefined}</Badge>)

      const badge = container.querySelector('.ant-badge')
      expect(badge).toBeInTheDocument()
    })

    it('handles null children', () => {
      const { container } = render(<Badge count={5}>{null}</Badge>)

      const badge = container.querySelector('.ant-badge')
      expect(badge).toBeInTheDocument()
    })

    it('handles empty string children', () => {
      render(<Badge count={5} />)

      const badge = screen.getByTitle('5')
      expect(badge).toBeInTheDocument()
    })

    it('handles multiple children', () => {
      render(
        <Badge count={3}>
          <div>First child</div>
          <span>Second child</span>
        </Badge>
      )

      expect(screen.getByText('First child')).toBeInTheDocument()
      expect(screen.getByText('Second child')).toBeInTheDocument()
      expect(screen.getByTitle('3')).toBeInTheDocument()
    })

    it('handles complex children structure', () => {
      render(
        <Badge count={7}>
          <div>
            <h3>Title</h3>
            <p>Description</p>
            <button type="button">Action</button>
          </div>
        </Badge>
      )

      expect(screen.getByText('Title')).toBeInTheDocument()
      expect(screen.getByText('Description')).toBeInTheDocument()
      expect(screen.getByText('Action')).toBeInTheDocument()
      expect(screen.getByTitle('7')).toBeInTheDocument()
    })

    it('handles count of 0 without showZero', () => {
      const { container } = render(
        <Badge count={0}>
          <div>Content</div>
        </Badge>
      )

      expect(screen.getByText('Content')).toBeInTheDocument()
      const countBadge = container.querySelector('.ant-badge-count')
      expect(countBadge).not.toBeInTheDocument()
    })

    it('handles negative count', () => {
      render(
        <Badge count={-5}>
          <div>Content</div>
        </Badge>
      )

      expect(screen.getByText('Content')).toBeInTheDocument()
    })
  })

  describe('props forwarding', () => {
    it('forwards all props to AntBadge', () => {
      render(
        <Badge
          count={42}
          dot={false}
          size="default"
          title="Custom title"
          data-testid="badge-wrapper"
        >
          <div>Content</div>
        </Badge>
      )

      expect(screen.getByTitle('Custom title')).toBeInTheDocument()
      expect(screen.getByTestId('badge-wrapper')).toBeInTheDocument()
      expect(screen.getByText('Content')).toBeInTheDocument()
    })

    it('handles boolean props correctly', () => {
      const { container } = render(
        <Badge dot count={1}>
          <div>Content</div>
        </Badge>
      )

      const dotBadge = container.querySelector('.ant-badge-dot')
      expect(dotBadge).toBeInTheDocument()
      expect(screen.getByText('Content')).toBeInTheDocument()
    })

    it('handles function props', () => {
      const countRender = jest.fn(() => <span>Custom Count</span>)

      render(
        <Badge count={countRender as any}>
          <div>Content</div>
        </Badge>
      )

      expect(screen.getByText('Content')).toBeInTheDocument()
    })
  })

  describe('accessibility', () => {
    it('provides proper accessibility attributes', () => {
      render(
        <Badge count={5} title="5 notifications">
          <button type="button">Notifications</button>
        </Badge>
      )

      expect(screen.getByTitle('5 notifications')).toBeInTheDocument()
      expect(screen.getByRole('button')).toBeInTheDocument()
    })

    it('handles aria labels correctly', () => {
      const { container } = render(
        <Badge count={3} aria-label="3 unread messages">
          <div>Messages</div>
        </Badge>
      )

      const badge = container.querySelector('[aria-label="3 unread messages"]')
      expect(badge).toBeInTheDocument()
    })
  })
})
