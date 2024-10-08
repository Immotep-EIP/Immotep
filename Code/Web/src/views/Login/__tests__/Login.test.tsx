import '@testing-library/jest-dom'
import { render, screen, fireEvent, act } from '@testing-library/react'
import LoginPage from '@/views/Login/Login'

jest.mock('@/hooks/useNavigation/useNavigation', () => ({
  __esModule: true,
  default: jest.fn(() => ({
    goToHome: jest.fn(),
    goToSignup: jest.fn()
  }))
}))

describe('LoginPage component', () => {
  test('renders all components', () => {
    render(<LoginPage />)

    const title = screen.getByText(/Welcome back/i)
    expect(title).toBeInTheDocument()

    const emailInput = screen.getByPlaceholderText(/Enter your email/i)
    const passwordInput = screen.getByPlaceholderText(/Enter your password/i)
    const signInButton = screen.getByRole('button', { name: /Sign in/i })

    expect(emailInput).toBeInTheDocument()
    expect(passwordInput).toBeInTheDocument()
    expect(signInButton).toBeInTheDocument()
  })

  test('submits the form with correct values', async () => {
    render(<LoginPage />)

    const emailInput = screen.getByPlaceholderText(/Enter your email/i)
    const passwordInput = screen.getByPlaceholderText(/Enter your password/i)
    const signInButton = screen.getByRole('button', { name: /Sign in/i })

    await act(async () => {
      fireEvent.change(emailInput, { target: { value: 'test@example.com' } })
      fireEvent.change(passwordInput, { target: { value: 'password123' } })
      fireEvent.click(signInButton)
    })
  })
})
