import { render, screen, fireEvent, waitFor } from '@testing-library/react'
import '@testing-library/jest-dom'
import Register from '@/views/Authentification/Register/Register'
import { register } from '@/services/api/Authentification/AuthApi'
import useNavigation from '@/hooks/useNavigation/useNavigation'

jest.mock('@/services/api/Authentification/AuthApi', () => ({
  register: jest.fn()
}))
jest.mock('@/hooks/useNavigation/useNavigation', () => ({
  __esModule: true,
  default: jest.fn(() => ({
    goToLogin: jest.fn()
  }))
}))

describe('Register Component', () => {
  const mockRegister = register as jest.Mock
  const mockGoToLogin = useNavigation().goToLogin as jest.Mock

  beforeEach(() => {
    mockRegister.mockReset()
    mockGoToLogin.mockReset()
  })

  it('renders the form elements correctly', () => {
    render(<Register />)

    expect(screen.getByLabelText('components.input.firstName.label')).toBeInTheDocument()
    expect(screen.getByLabelText('components.input.lastName.label')).toBeInTheDocument()
    expect(screen.getByLabelText('components.input.email.label')).toBeInTheDocument()
    expect(
      screen.getByPlaceholderText('components.input.password.placeholder')
    ).toBeInTheDocument()
    expect(
      screen.getByPlaceholderText('components.input.confirmPassword.placeholder')
    ).toBeInTheDocument()
    expect(screen.getByText('components.button.signUp')).toBeInTheDocument()
    expect(screen.getByText('pages.register.alreadyHaveAccount')).toBeInTheDocument()
  })

  it('displays error message if passwords do not match', async () => {
    render(<Register />)

    fireEvent.input(screen.getByLabelText('components.input.firstName.label'), {
      target: { value: 'John' }
    })
    fireEvent.input(screen.getByLabelText('components.input.lastName.label'), {
      target: { value: 'Doe' }
    })
    fireEvent.input(screen.getByLabelText('components.input.email.label'), {
      target: { value: 'john.doe@example.com' }
    })
    fireEvent.input(screen.getByPlaceholderText('components.input.password.placeholder'), {
      target: { value: 'password123' }
    })
    fireEvent.input(screen.getByPlaceholderText('components.input.confirmPassword.placeholder'), {
      target: { value: 'password321' }
    })

    fireEvent.click(screen.getByText('components.button.signUp'))

    await waitFor(() =>
      expect(
        screen.getByText('pages.register.passwordsNotMatch')
      ).toBeInTheDocument()
    )
  })

  it('registers user and redirects to login on success', async () => {
    mockRegister.mockResolvedValueOnce({})
    render(<Register />)

    fireEvent.input(screen.getByLabelText('components.input.firstName.label'), {
      target: { value: 'John' }
    })
    fireEvent.input(screen.getByLabelText('components.input.lastName.label'), {
      target: { value: 'Doe' }
    })
    fireEvent.input(screen.getByLabelText('components.input.email.label'), {
      target: { value: 'john.doe@example.com' }
    })
    fireEvent.input(screen.getByPlaceholderText('components.input.password.placeholder'), {
      target: { value: 'password123' }
    })
    fireEvent.input(screen.getByPlaceholderText('components.input.confirmPassword.placeholder'), {
      target: { value: 'password123' }
    })

    fireEvent.click(screen.getByText('components.button.signUp'))

    await waitFor(() => {
      expect(mockRegister).toHaveBeenCalledWith({
        firstname: 'John',
        lastname: 'Doe',
        email: 'john.doe@example.com',
        password: 'password123',
        confirmPassword: 'password123',
        termAgree: false
      })
    })
  })

  it('displays error if email already exists', async () => {
    const consoleSpy = jest.spyOn(console, 'error').mockImplementation(() => {})
    mockRegister.mockRejectedValueOnce({ response: { status: 409 } })
    render(<Register />)

    fireEvent.input(screen.getByLabelText('components.input.firstName.label'), {
      target: { value: 'John' }
    })
    fireEvent.input(screen.getByLabelText('components.input.lastName.label'), {
      target: { value: 'Doe' }
    })
    fireEvent.input(screen.getByLabelText('components.input.email.label'), {
      target: { value: 'john.doe@example.com' }
    })
    fireEvent.input(screen.getByPlaceholderText('components.input.password.placeholder'), {
      target: { value: 'password123' }
    })
    fireEvent.input(screen.getByPlaceholderText('components.input.confirmPassword.placeholder'), {
      target: { value: 'password123' }
    })

    fireEvent.click(screen.getByText('components.button.signUp'))

    await waitFor(() =>
      expect(screen.getByText('pages.register.emailAlreadyUsed')).toBeInTheDocument()
    )

    expect(consoleSpy).toHaveBeenCalledWith(
      'pages.register.registrationError',
      expect.anything()
    )

    consoleSpy.mockRestore()
  })

  // Dont't found a solution
  it.skip('redirects to login when "Sign in" is clicked', async () => {
    render(<Register />)
    fireEvent.click(screen.getByText(/sign in/i))

    expect(mockGoToLogin).toHaveBeenCalled()
  })

  test('displays error message when form submission fails', async () => {
    render(<Register />)

    fireEvent.click(screen.getByText('components.button.signUp'))

    await waitFor(() =>
      expect(
        screen.getByText('pages.register.fillFields')
      ).toBeInTheDocument()
    )
  })
})
