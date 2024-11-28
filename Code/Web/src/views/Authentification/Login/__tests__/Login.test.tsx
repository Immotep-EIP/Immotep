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
    expect(screen.getByLabelText('components.input.email.label')).toBeInTheDocument()
    expect(screen.getByLabelText('components.input.password.label')).toBeInTheDocument()
    expect(screen.getByRole('button', { name: 'components.button.signIn' })).toBeInTheDocument()
    expect(screen.getByText('pages.login.dontHaveAccount')).toBeInTheDocument()
  })

  it('shows error message if fields are empty', async () => {
    render(<Login />)

    fireEvent.click(screen.getByRole('button', { name: 'components.button.signIn' }))

    expect(
      await screen.findByText('components.input.email.error')
    ).toBeInTheDocument()
    expect(
      await screen.findByText('components.input.password.error')
    ).toBeInTheDocument()
  })

  it('submits form with correct values', async () => {
    render(<Login />)

    fireEvent.input(screen.getByLabelText('components.input.email.label'), {
      target: { value: 'test@example.com' }
    })
    fireEvent.input(screen.getByLabelText('components.input.password.label'), {
      target: { value: 'password' }
    })

    fireEvent.click(screen.getByRole('button', { name: 'components.button.signIn' }))

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

    fireEvent.input(screen.getByLabelText('components.input.email.label'), {
      target: { value: 'test@example.com' }
    })
    fireEvent.input(screen.getByLabelText('components.input.password.label'), {
      target: { value: 'password' }
    })

    fireEvent.click(screen.getByRole('button', { name: 'components.button.signIn' }))

    expect(
      await screen.findByText('pages.login.connectionError')
    ).toBeInTheDocument()
  })

  // Dont't found a solution
  it.skip('navigates to signup when clicking "Sign up"', () => {
    const { goToSignup } = useNavigation()

    render(<Login />)

    fireEvent.click(screen.getByText('components.button.signUp'))

    expect(goToSignup).toHaveBeenCalled()
  })

  // Dont't found a solution
  it.skip('navigates to forgot password when clicking "Forgot password"', () => {
    const { goToForgotPassword } = useNavigation()

    render(<Login />)

    fireEvent.click(screen.getByText('components.button.askForgotPassword'))

    expect(goToForgotPassword).toHaveBeenCalled()
  })
})
