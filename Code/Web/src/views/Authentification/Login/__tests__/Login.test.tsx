import React from 'react'
import '@testing-library/jest-dom'
import { render, screen, fireEvent, waitFor } from '@testing-library/react'
import { HelmetProvider } from 'react-helmet-async'
import Login from '@/views/Authentification/Login/Login'
import { useAuth } from '@/context/authContext'
import useNavigation from '@/hooks/Navigation/useNavigation'

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

jest.mock('react-i18next', () => ({
  useTranslation: () => ({
    t: (key: string) => key
  }),
  Trans: ({ children }: { children: React.ReactNode }) => children
}))

describe('Login Component', () => {
  const mockLogin = jest.fn()
  const mockGoToSignup = jest.fn()
  const mockGoToOverview = jest.fn()
  const mockGoToForgotPassword = jest.fn()

  beforeEach(() => {
    ;(useAuth as jest.Mock).mockReturnValue({
      login: mockLogin
    })
    ;(useNavigation as jest.Mock).mockReturnValue({
      goToSignup: mockGoToSignup,
      goToOverview: mockGoToOverview,
      goToForgotPassword: mockGoToForgotPassword
    })
    jest.clearAllMocks()
  })

  // eslint-disable-next-line no-undef
  const renderWithHelmet = (component: React.ReactNode) =>
    render(<HelmetProvider>{component}</HelmetProvider>)

  it('renders the Login component with form fields', () => {
    renderWithHelmet(<Login />)

    expect(
      screen.getByLabelText('components.input.email.label')
    ).toBeInTheDocument()
    expect(
      screen.getByLabelText('components.input.password.label')
    ).toBeInTheDocument()
    expect(
      screen.getByRole('button', { name: 'components.button.sign_in' })
    ).toBeInTheDocument()
    expect(screen.getByText('pages.login.title')).toBeInTheDocument()
    expect(screen.getByText('pages.login.description')).toBeInTheDocument()
  })

  it('navigates to Forgot Password page when "Enter" key is pressed on "Forgot Password" link', () => {
    renderWithHelmet(<Login />)

    const forgotPasswordLink = screen.getByText(
      'components.button.ask_forgot_password'
    )

    fireEvent.keyDown(forgotPasswordLink, {
      key: 'Enter',
      code: 'Enter',
      charCode: 13
    })

    expect(mockGoToForgotPassword).toHaveBeenCalledTimes(1)
  })

  it('navigates to Signup page when "Sign Up" is clicked', () => {
    renderWithHelmet(<Login />)
    fireEvent.click(screen.getByText('components.button.sign_up'))
    expect(mockGoToSignup).toHaveBeenCalledTimes(1)
  })

  it('navigates to Forgot Password page when the link is clicked', () => {
    renderWithHelmet(<Login />)
    fireEvent.click(screen.getByText('components.button.ask_forgot_password'))
    expect(mockGoToForgotPassword).toHaveBeenCalledTimes(1)
  })

  it('shows an error message when fields are empty on form submission', async () => {
    renderWithHelmet(<Login />)
    fireEvent.click(
      screen.getByRole('button', { name: 'components.button.sign_in' })
    )

    await waitFor(() => {
      expect(
        screen.getByText('components.input.email.error')
      ).toBeInTheDocument()
      expect(
        screen.getByText('components.input.password.error')
      ).toBeInTheDocument()
    })
  })

  it('submits the form successfully with valid inputs', async () => {
    renderWithHelmet(<Login />)

    const emailInput = screen.getByLabelText('components.input.email.label')
    const passwordInput = screen.getByLabelText(
      'components.input.password.label'
    )
    const submitButton = screen.getByRole('button', {
      name: 'components.button.sign_in'
    })

    fireEvent.change(emailInput, { target: { value: 'user@example.com' } })
    fireEvent.change(passwordInput, { target: { value: 'password123' } })
    fireEvent.click(submitButton)

    await waitFor(() => {
      expect(mockLogin).toHaveBeenCalledWith({
        username: 'user@example.com',
        password: 'password123',
        grant_type: 'password',
        rememberMe: false
      })
    })

    expect(mockGoToOverview).toHaveBeenCalledTimes(1)
  })

  it('shows an error message when login fails with 401', async () => {
    mockLogin.mockRejectedValueOnce({ response: { status: 401 } })

    renderWithHelmet(<Login />)

    const emailInput = screen.getByLabelText('components.input.email.label')
    const passwordInput = screen.getByLabelText(
      'components.input.password.label'
    )
    const submitButton = screen.getByRole('button', {
      name: 'components.button.sign_in'
    })

    fireEvent.change(emailInput, { target: { value: 'wrong@example.com' } })
    fireEvent.change(passwordInput, { target: { value: 'wrongpassword' } })
    fireEvent.click(submitButton)

    await waitFor(() => {
      expect(
        screen.getByText('pages.login.connection_error')
      ).toBeInTheDocument()
    })
  })

  it('removes tokens from sessionStorage on mount', () => {
    sessionStorage.setItem('access_token', 'dummy_access')
    sessionStorage.setItem('refresh_token', 'dummy_refresh')
    sessionStorage.setItem('expires_in', '12345')

    renderWithHelmet(<Login />)

    expect(sessionStorage.getItem('access_token')).toBeNull()
    expect(sessionStorage.getItem('refresh_token')).toBeNull()
    expect(sessionStorage.getItem('expires_in')).toBeNull()
  })

  it('navigates to overview if tokens exist in localStorage', () => {
    localStorage.setItem('access_token', 'dummy_access')
    localStorage.setItem('refresh_token', 'dummy_refresh')
    localStorage.setItem('expires_in', '12345')

    renderWithHelmet(<Login />)

    expect(mockGoToOverview).toHaveBeenCalledTimes(1)
  })

  it('calls the goToForgotPassword when the forgot password link is clicked', () => {
    renderWithHelmet(<Login />)
    fireEvent.click(screen.getByText('components.button.ask_forgot_password'))
    expect(mockGoToForgotPassword).toHaveBeenCalledTimes(1)
  })

  it('checks rememberMe functionality', async () => {
    renderWithHelmet(<Login />)

    const emailInput = screen.getByLabelText('components.input.email.label')
    const passwordInput = screen.getByLabelText(
      'components.input.password.label'
    )
    const rememberMeCheckbox = screen.getByLabelText(
      'components.button.remember_me'
    )
    const submitButton = screen.getByRole('button', {
      name: 'components.button.sign_in'
    })

    fireEvent.change(emailInput, { target: { value: 'user@example.com' } })
    fireEvent.change(passwordInput, { target: { value: 'password123' } })
    fireEvent.click(rememberMeCheckbox)
    fireEvent.click(submitButton)

    await waitFor(() => {
      expect(mockLogin).toHaveBeenCalledWith({
        username: 'user@example.com',
        password: 'password123',
        grant_type: 'password',
        rememberMe: true
      })
    })
  })

  it('navigates to Signup page when "Sign Up" link is clicked', () => {
    renderWithHelmet(<Login />)

    const signUpLink = screen.getByText('components.button.sign_up')

    fireEvent.click(signUpLink)

    expect(mockGoToSignup).toHaveBeenCalledTimes(1)
  })

  it('navigates to Signup page when "Sign Up" link is pressed with Enter key', () => {
    renderWithHelmet(<Login />)

    const signUpLink = screen.getByText('components.button.sign_up')

    fireEvent.keyDown(signUpLink, { key: 'Enter', code: 'Enter', charCode: 13 })
  })
})
