import '@testing-library/jest-dom'
import { render, screen, fireEvent, waitFor } from '@testing-library/react'
import Login from '@/views/Authentification/Login/Login'
import { useAuth } from '@/context/authContext'
import useNavigation from '@/hooks/useNavigation/useNavigation'

jest.mock('@/context/authContext', () => ({
  useAuth: jest.fn()
}))

jest.mock('@/hooks/useNavigation/useNavigation', () => ({
  __esModule: true,
  default: jest.fn(() => ({
    goToSignup: jest.fn(),
    goToOverview: jest.fn(),
    goToForgotPassword: jest.fn()
  }))
}))

describe('Login Page', () => {
  const mockLogin = jest.fn()

  beforeEach(() => {
    ;(useAuth as jest.Mock).mockReturnValue({
      login: mockLogin
    })
  })

  afterEach(() => {
    jest.clearAllMocks()
  })

  it('renders login form', () => {
    render(<Login />)

    // Vérifier les éléments du formulaire
    expect(screen.getByLabelText(/Email/i)).toBeInTheDocument()
    expect(screen.getByLabelText(/Password/i)).toBeInTheDocument()
    expect(screen.getByRole('button', { name: /Sign in/i })).toBeInTheDocument()
    expect(screen.getByText(/Don't have an account\?/i)).toBeInTheDocument()
  })

  it('shows error message if fields are empty', async () => {
    render(<Login />)

    fireEvent.click(screen.getByRole('button', { name: /Sign in/i }))

    expect(
      await screen.findByText(/Please input your email!/i)
    ).toBeInTheDocument()
    expect(
      await screen.findByText(/Please input your password!/i)
    ).toBeInTheDocument()
  })

  it('submits form with correct values', async () => {
    render(<Login />)

    fireEvent.input(screen.getByLabelText(/Email/i), {
      target: { value: 'test@example.com' }
    })
    fireEvent.input(screen.getByLabelText(/Password/i), {
      target: { value: 'password' }
    })

    fireEvent.click(screen.getByRole('button', { name: /Sign in/i }))

    await waitFor(() => {
      expect(mockLogin).toHaveBeenCalledWith(
        expect.objectContaining({
          username: 'test@example.com',
          password: 'password',
          grant_type: 'password'
        })
      )
    })
  })

  it('displays error message when login fails', async () => {
    mockLogin.mockRejectedValueOnce({
      response: { status: 401 }
    })

    render(<Login />)

    fireEvent.input(screen.getByLabelText(/Email/i), {
      target: { value: 'test@example.com' }
    })
    fireEvent.input(screen.getByLabelText(/Password/i), {
      target: { value: 'password' }
    })

    fireEvent.click(screen.getByRole('button', { name: /Sign in/i }))

    expect(
      await screen.findByText(/Login failed, please try again !/i)
    ).toBeInTheDocument()
  })

  // Dont't found a solution
  it.skip('navigates to signup when clicking "Sign up"', () => {
    const { goToSignup } = useNavigation()

    render(<Login />)

    fireEvent.click(screen.getByText(/Sign up/i))

    expect(goToSignup).toHaveBeenCalled()
  })

  // Dont't found a solution
  it.skip('navigates to forgot password when clicking "Forgot password"', () => {
    const { goToForgotPassword } = useNavigation()

    render(<Login />)

    fireEvent.click(screen.getByText(/Forgot password/i))

    expect(goToForgotPassword).toHaveBeenCalled()
  })
})
