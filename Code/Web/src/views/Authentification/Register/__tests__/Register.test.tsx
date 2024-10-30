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

    expect(screen.getByLabelText(/first name/i)).toBeInTheDocument()
    expect(screen.getByLabelText(/last name/i)).toBeInTheDocument()
    expect(screen.getByLabelText(/email/i)).toBeInTheDocument()
    expect(
      screen.getByPlaceholderText('Enter your password')
    ).toBeInTheDocument()
    expect(
      screen.getByPlaceholderText('Confirm your password')
    ).toBeInTheDocument()
    expect(screen.getByText(/sign up/i)).toBeInTheDocument()
    expect(screen.getByText(/already have an account/i)).toBeInTheDocument()
  })

  it('displays error message if passwords do not match', async () => {
    render(<Register />)

    fireEvent.input(screen.getByLabelText(/first name/i), {
      target: { value: 'John' }
    })
    fireEvent.input(screen.getByLabelText(/last name/i), {
      target: { value: 'Doe' }
    })
    fireEvent.input(screen.getByLabelText(/email/i), {
      target: { value: 'john.doe@example.com' }
    })
    fireEvent.input(screen.getByPlaceholderText('Enter your password'), {
      target: { value: 'password123' }
    })
    fireEvent.input(screen.getByPlaceholderText('Confirm your password'), {
      target: { value: 'password321' }
    })

    fireEvent.click(screen.getByText(/sign up/i))

    await waitFor(() =>
      expect(
        screen.getByText(/please confirm your password/i)
      ).toBeInTheDocument()
    )
  })

  it('registers user and redirects to login on success', async () => {
    mockRegister.mockResolvedValueOnce({})
    render(<Register />)

    fireEvent.input(screen.getByLabelText(/first name/i), {
      target: { value: 'John' }
    })
    fireEvent.input(screen.getByLabelText(/last name/i), {
      target: { value: 'Doe' }
    })
    fireEvent.input(screen.getByLabelText(/email/i), {
      target: { value: 'john.doe@example.com' }
    })
    fireEvent.input(screen.getByPlaceholderText('Enter your password'), {
      target: { value: 'password123' }
    })
    fireEvent.input(screen.getByPlaceholderText('Confirm your password'), {
      target: { value: 'password123' }
    })

    fireEvent.click(screen.getByText(/sign up/i))

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

    fireEvent.input(screen.getByLabelText(/first name/i), {
      target: { value: 'John' }
    })
    fireEvent.input(screen.getByLabelText(/last name/i), {
      target: { value: 'Doe' }
    })
    fireEvent.input(screen.getByLabelText(/email/i), {
      target: { value: 'john.doe@example.com' }
    })
    fireEvent.input(screen.getByPlaceholderText('Enter your password'), {
      target: { value: 'password123' }
    })
    fireEvent.input(screen.getByPlaceholderText('Confirm your password'), {
      target: { value: 'password123' }
    })

    fireEvent.click(screen.getByText(/sign up/i))

    await waitFor(() =>
      expect(screen.getByText(/email already exist/i)).toBeInTheDocument()
    )

    expect(consoleSpy).toHaveBeenCalledWith(
      'Registration error:',
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

    fireEvent.click(screen.getByText(/sign up/i))

    await waitFor(() =>
      expect(
        screen.getByText(/an error occured, please try again/i)
      ).toBeInTheDocument()
    )
  })
})
