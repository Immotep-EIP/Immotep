import { render, screen, fireEvent, waitFor, act } from '@testing-library/react'
import { HelmetProvider } from 'react-helmet-async'
import { message } from 'antd'
import useNavigation from '@/hooks/Navigation/useNavigation'
import ForgotPassword from '../ForgotPassword'

jest.mock('react-i18next', () => ({
  useTranslation: jest.fn().mockReturnValue({
    t: (key: string) => key
  })
}))

jest.mock('@/hooks/useNavigation/useNavigation', () => ({
  __esModule: true,
  default: () => ({
    goToLogin: jest.fn()
  })
}))

jest.mock('antd', () => ({
  ...jest.requireActual('antd'),
  message: {
    success: jest.fn(),
    error: jest.fn()
  }
}))

// eslint-disable-next-line no-undef
const renderWithHelmet = (component: React.ReactNode) =>
  render(<HelmetProvider>{component}</HelmetProvider>)

describe('ForgotPassword', () => {
  beforeEach(() => {
    renderWithHelmet(<ForgotPassword />)
  })

  test('renders correctly', () => {
    expect(screen.getByText('pages.forgot_password.title')).toBeInTheDocument()
    expect(
      screen.getByText('pages.forgot_password.description')
    ).toBeInTheDocument()
    expect(
      screen.getByLabelText('components.input.email.label')
    ).toBeInTheDocument()
    expect(
      screen.getByLabelText('components.input.email_confirmation.label')
    ).toBeInTheDocument()
    expect(screen.getByText('components.button.send_email')).toBeInTheDocument()
  })

  test('displays error if emails do not match', async () => {
    const emailInput = screen.getByLabelText('components.input.email.label')
    const emailConfirmationInput = screen.getByLabelText(
      'components.input.email_confirmation.label'
    )
    const submitButton = screen.getByText('components.button.send_email')

    fireEvent.change(emailInput, { target: { value: 'test@example.com' } })
    fireEvent.change(emailConfirmationInput, {
      target: { value: 'mismatch@example.com' }
    })

    await act(async () => {
      fireEvent.click(submitButton)
    })

    await waitFor(() =>
      expect(message.error).toHaveBeenCalledWith(
        'pages.forgot_password.emails_dont_match'
      )
    )
  })

  test('displays success message on correct form submission', async () => {
    const emailInput = screen.getByLabelText('components.input.email.label')
    const emailConfirmationInput = screen.getByLabelText(
      'components.input.email_confirmation.label'
    )
    const submitButton = screen.getByText('components.button.send_email')

    fireEvent.change(emailInput, { target: { value: 'test@example.com' } })
    fireEvent.change(emailConfirmationInput, {
      target: { value: 'test@example.com' }
    })

    await act(async () => {
      fireEvent.click(submitButton)
    })

    await waitFor(() =>
      expect(message.success).toHaveBeenCalledWith(
        'pages.forgot_password.send_email_success'
      )
    )
  })

  test('displays error if fields are empty on submission', async () => {
    const submitButton = screen.getByText('components.button.send_email')

    await act(async () => {
      fireEvent.click(submitButton)
    })

    await waitFor(() =>
      expect(message.error).toHaveBeenCalledWith(
        'pages.forgot_password.fill_fields'
      )
    )
  })

  test('calls goToLogin on successful form submission', async () => {
    const emailInput = screen.getByLabelText('components.input.email.label')
    const emailConfirmationInput = screen.getByLabelText(
      'components.input.email_confirmation.label'
    )
    const submitButton = screen.getByText('components.button.send_email')

    const { goToLogin } = useNavigation()

    expect(goToLogin).not.toHaveBeenCalled()

    fireEvent.change(emailInput, { target: { value: 'test@example.com' } })
    fireEvent.change(emailConfirmationInput, {
      target: { value: 'test@example.com' }
    })

    await act(async () => {
      fireEvent.click(submitButton)
    })
  })
})
