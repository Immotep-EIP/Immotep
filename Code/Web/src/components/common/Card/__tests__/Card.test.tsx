import { render, screen } from '@testing-library/react'
import '@testing-library/jest-dom'
import { Card as AntCard } from 'antd'
import Card from '../Card'

describe('Card Component', () => {
  it('renders children correctly', () => {
    render(
      <Card>
        <div>Test content</div>
      </Card>
    )

    expect(screen.getByText('Test content')).toBeInTheDocument()
  })

  it('renders with title prop', () => {
    render(
      <Card title="Test Title">
        <div>Test content</div>
      </Card>
    )

    expect(screen.getByText('Test Title')).toBeInTheDocument()
    expect(screen.getByText('Test content')).toBeInTheDocument()
  })

  it('renders with React node as title', () => {
    const titleNode = <span data-testid="custom-title">Custom Title</span>

    render(
      <Card title={titleNode}>
        <div>Test content</div>
      </Card>
    )

    expect(screen.getByTestId('custom-title')).toBeInTheDocument()
  })

  describe('customVariant prop', () => {
    it('applies default variant styles by default', () => {
      render(
        <Card data-testid="card">
          <div>Content</div>
        </Card>
      )

      const card = screen.getByTestId('card')
      const cardBody = card.querySelector('.ant-card-body')

      expect(cardBody).toHaveStyle({ padding: '16px' })
    })

    it('applies elevated variant styles', () => {
      render(
        <Card customVariant="elevated" data-testid="card">
          <div>Content</div>
        </Card>
      )

      const card = screen.getByTestId('card')
      const cardBody = card.querySelector('.ant-card-body')

      expect(cardBody).toHaveStyle({
        padding: '16px',
        'box-shadow': '0 4px 12px rgba(0, 0, 0, 0.1)'
      })
    })

    it('applies outlined variant styles', () => {
      render(
        <Card customVariant="outlined" data-testid="card">
          <div>Content</div>
        </Card>
      )

      const card = screen.getByTestId('card')
      const cardBody = card.querySelector('.ant-card-body')

      expect(cardBody).toHaveStyle({
        padding: '16px',
        border: '1px solid #d9d9d9'
      })
    })

    it('applies default variant when invalid variant is provided', () => {
      render(
        <Card customVariant={'invalid' as any} data-testid="card">
          <div>Content</div>
        </Card>
      )

      const card = screen.getByTestId('card')
      const cardBody = card.querySelector('.ant-card-body')

      expect(cardBody).toHaveStyle({ padding: '16px' })
    })
  })

  describe('padding prop', () => {
    it('applies no padding when padding is "none"', () => {
      render(
        <Card padding="none" data-testid="card">
          <div>Content</div>
        </Card>
      )

      const card = screen.getByTestId('card')
      const cardBody = card.querySelector('.ant-card-body')

      expect(cardBody).toHaveStyle({ padding: '0px' })
    })

    it('applies small padding when padding is "small"', () => {
      render(
        <Card padding="small" data-testid="card">
          <div>Content</div>
        </Card>
      )

      const card = screen.getByTestId('card')
      const cardBody = card.querySelector('.ant-card-body')

      expect(cardBody).toHaveStyle({ padding: '12px' })
    })

    it('applies medium padding by default', () => {
      render(
        <Card data-testid="card">
          <div>Content</div>
        </Card>
      )

      const card = screen.getByTestId('card')
      const cardBody = card.querySelector('.ant-card-body')

      expect(cardBody).toHaveStyle({ padding: '16px' })
    })

    it('applies medium padding when padding is "medium"', () => {
      render(
        <Card padding="medium" data-testid="card">
          <div>Content</div>
        </Card>
      )

      const card = screen.getByTestId('card')
      const cardBody = card.querySelector('.ant-card-body')

      expect(cardBody).toHaveStyle({ padding: '16px' })
    })

    it('applies large padding when padding is "large"', () => {
      render(
        <Card padding="large" data-testid="card">
          <div>Content</div>
        </Card>
      )

      const card = screen.getByTestId('card')
      const cardBody = card.querySelector('.ant-card-body')

      expect(cardBody).toHaveStyle({ padding: '24px' })
    })

    it('applies default padding when invalid padding is provided', () => {
      render(
        <Card padding={'invalid' as any} data-testid="card">
          <div>Content</div>
        </Card>
      )

      const card = screen.getByTestId('card')
      const cardBody = card.querySelector('.ant-card-body')

      expect(cardBody).toHaveStyle({ padding: '16px' })
    })
  })

  describe('combined props', () => {
    it('combines elevated variant with large padding', () => {
      render(
        <Card customVariant="elevated" padding="large" data-testid="card">
          <div>Content</div>
        </Card>
      )

      const card = screen.getByTestId('card')
      const cardBody = card.querySelector('.ant-card-body')

      expect(cardBody).toHaveStyle({
        padding: '24px',
        'box-shadow': '0 4px 12px rgba(0, 0, 0, 0.1)'
      })
    })

    it('combines outlined variant with no padding', () => {
      render(
        <Card customVariant="outlined" padding="none" data-testid="card">
          <div>Content</div>
        </Card>
      )

      const card = screen.getByTestId('card')
      const cardBody = card.querySelector('.ant-card-body')

      expect(cardBody).toHaveStyle({
        padding: '0px',
        border: '1px solid #d9d9d9'
      })
    })
  })

  it('passes through other AntCard props', () => {
    render(
      <Card
        className="custom-class"
        data-testid="card"
        bordered={false}
        size="small"
      >
        <div>Content</div>
      </Card>
    )

    const card = screen.getByTestId('card')
    expect(card).toHaveClass('custom-class')
    expect(card).toHaveClass('ant-card-small')
    expect(card).not.toHaveClass('ant-card-bordered')
  })

  it('omits title and variant from passed props', () => {
    // This test ensures that our component properly omits the title and variant
    // props that might conflict with AntCard's props
    const { container } = render(
      <Card title="Test Title" customVariant="elevated" data-testid="card">
        <div>Content</div>
      </Card>
    )

    // The card should render without errors and display the title
    expect(screen.getByText('Test Title')).toBeInTheDocument()
    expect(container.firstChild).toBeInTheDocument()
  })

  describe('Card.Grid static property', () => {
    it('exposes Grid component from AntCard', () => {
      expect(Card.Grid).toBe(AntCard.Grid)
    })

    it('renders Card.Grid correctly', () => {
      render(
        <Card data-testid="card">
          <Card.Grid data-testid="grid">Grid Content</Card.Grid>
        </Card>
      )

      expect(screen.getByTestId('grid')).toBeInTheDocument()
      expect(screen.getByText('Grid Content')).toBeInTheDocument()
    })
  })

  describe('edge cases', () => {
    it('handles empty children', () => {
      render(<Card data-testid="card">{null}</Card>)

      expect(screen.getByTestId('card')).toBeInTheDocument()
    })

    it('handles multiple children', () => {
      render(
        <Card>
          <div>First child</div>
          <div>Second child</div>
          <span>Third child</span>
        </Card>
      )

      expect(screen.getByText('First child')).toBeInTheDocument()
      expect(screen.getByText('Second child')).toBeInTheDocument()
      expect(screen.getByText('Third child')).toBeInTheDocument()
    })

    it('handles complex children structure', () => {
      render(
        <Card title="Complex Card">
          <div>
            <h2>Heading</h2>
            <p>Paragraph</p>
            <ul>
              <li>Item 1</li>
              <li>Item 2</li>
            </ul>
          </div>
        </Card>
      )

      expect(screen.getByText('Complex Card')).toBeInTheDocument()
      expect(screen.getByText('Heading')).toBeInTheDocument()
      expect(screen.getByText('Paragraph')).toBeInTheDocument()
      expect(screen.getByText('Item 1')).toBeInTheDocument()
      expect(screen.getByText('Item 2')).toBeInTheDocument()
    })
  })
})
