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
jest.mock('react-i18next', () => ({
  __esModule: true,
  useTranslation: () => ({
    t: (key: string) => key
  }),
  initReactI18next: {
    type: '3rdParty',
    init: jest.fn()
  }
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

    expect(
      screen.getByLabelText('components.input.first_name.label')
    ).toBeInTheDocument()
    expect(
      screen.getByLabelText('components.input.last_name.label')
    ).toBeInTheDocument()
    expect(
      screen.getByLabelText('components.input.email.label')
    ).toBeInTheDocument()
    expect(
      screen.getByPlaceholderText('components.input.password.placeholder')
    ).toBeInTheDocument()
    expect(
      screen.getByPlaceholderText(
        'components.input.confirm_password.placeholder'
      )
    ).toBeInTheDocument()
    expect(screen.getByText('components.button.sign_up')).toBeInTheDocument()
    expect(
      screen.getByText('pages.register.already_have_account')
    ).toBeInTheDocument()
  })

  it('displays error message if passwords do not match', async () => {
    render(<Register />)

    fireEvent.input(
      screen.getByLabelText('components.input.first_name.label'),
      {
        target: { value: 'John' }
      }
    )
    fireEvent.input(screen.getByLabelText('components.input.last_name.label'), {
      target: { value: 'Doe' }
    })
    fireEvent.input(screen.getByLabelText('components.input.email.label'), {
      target: { value: 'john.doe@example.com' }
    })
    fireEvent.input(
      screen.getByPlaceholderText('components.input.password.placeholder'),
      {
        target: { value: 'password123' }
      }
    )
    fireEvent.input(
      screen.getByPlaceholderText(
        'components.input.confirm_password.placeholder'
      ),
      {
        target: { value: 'password321' }
      }
    )

    fireEvent.click(screen.getByText('components.button.sign_up'))

    await waitFor(() =>
      expect(
        screen.getByText('pages.register.confirm_password_error')
      ).toBeInTheDocument()
    )
  })

  it('registers user and redirects to login on success', async () => {
    mockRegister.mockResolvedValueOnce({})
    render(<Register />)

    fireEvent.input(
      screen.getByLabelText('components.input.first_name.label'),
      {
        target: { value: 'John' }
      }
    )
    fireEvent.input(screen.getByLabelText('components.input.last_name.label'), {
      target: { value: 'Doe' }
    })
    fireEvent.input(screen.getByLabelText('components.input.email.label'), {
      target: { value: 'john.doe@example.com' }
    })
    fireEvent.input(
      screen.getByPlaceholderText('components.input.password.placeholder'),
      {
        target: { value: 'password123' }
      }
    )
    fireEvent.input(
      screen.getByPlaceholderText(
        'components.input.confirm_password.placeholder'
      ),
      {
        target: { value: 'password123' }
      }
    )

    fireEvent.click(screen.getByText('components.button.sign_up'))

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

    fireEvent.input(
      screen.getByLabelText('components.input.first_name.label'),
      {
        target: { value: 'John' }
      }
    )
    fireEvent.input(screen.getByLabelText('components.input.last_name.label'), {
      target: { value: 'Doe' }
    })
    fireEvent.input(screen.getByLabelText('components.input.email.label'), {
      target: { value: 'john.doe@example.com' }
    })
    fireEvent.input(
      screen.getByPlaceholderText('components.input.password.placeholder'),
      {
        target: { value: 'password123' }
      }
    )
    fireEvent.input(
      screen.getByPlaceholderText(
        'components.input.confirm_password.placeholder'
      ),
      {
        target: { value: 'password123' }
      }
    )

    fireEvent.click(screen.getByText('components.button.sign_up'))

    await waitFor(() =>
      expect(
        screen.getByText('pages.register.email_already_used')
      ).toBeInTheDocument()
    )

    expect(consoleSpy).toHaveBeenCalledWith(
      'pages.register.register_error',
      expect.anything()
    )

    consoleSpy.mockRestore()
  })

  test('displays error message when form submission fails', async () => {
    render(<Register />)

    fireEvent.click(screen.getByText('components.button.sign_up'))

    await waitFor(() =>
      expect(screen.getByText('pages.register.fill_fields')).toBeInTheDocument()
    )
  })

  it('navigates to Login page when "Sign In" link is clicked', () => {
    render(<Register />)

    const signInLink = screen.getByText('components.button.sign_in')

    fireEvent.keyDown(signInLink, { key: 'Enter', code: 'Enter', charCode: 13 })
  })
})
