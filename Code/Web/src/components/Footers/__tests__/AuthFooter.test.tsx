import '@testing-library/jest-dom'
import { render, screen, fireEvent } from '@testing-library/react'
import AuthFooter from '@/components/Footers/AuthFooter'

describe('AuthFooter component', () => {
  const mockGoTo = jest.fn()

  const setup = () => {
    render(
      <AuthFooter
        goTo={mockGoTo}
        text="Already have an account?"
        buttonText="Login here"
      />
    )
  }

  it('renders the text and button', () => {
    setup()
    expect(screen.getByText('Already have an account?')).toBeInTheDocument()
    expect(screen.getByText('Login here')).toBeInTheDocument()
  })

  it('calls goTo when button is clicked', () => {
    setup()
    const button = screen.getByText('Login here')
    fireEvent.click(button)
    expect(mockGoTo).toHaveBeenCalledTimes(1)
  })
})
