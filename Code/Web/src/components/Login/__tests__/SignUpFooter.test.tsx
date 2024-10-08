import '@testing-library/jest-dom'
import { render, screen, fireEvent } from '@testing-library/react'
import SignUpFooter from '@/components/Login/SignUpFooter'

describe('SignUpFooter component', () => {
  test('renders the "Sign up" link', () => {
    render(<SignUpFooter goToSignup={jest.fn()} />)

    const signUpLink = screen.getByText(/sign up/i)
    expect(signUpLink).toBeInTheDocument()
  })

  test('calls goToSignup when "Sign up" button is clicked', () => {
    const goToSignupMock = jest.fn()
    render(<SignUpFooter goToSignup={goToSignupMock} />)

    const signUpLink = screen.getByText(/sign up/i)
    fireEvent.click(signUpLink)

    expect(goToSignupMock).toHaveBeenCalledTimes(1)
  })
})
