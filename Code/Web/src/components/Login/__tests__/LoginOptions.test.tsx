import '@testing-library/jest-dom'
import { render, screen, fireEvent } from '@testing-library/react'
import { Form } from 'antd'
import LoginOptions from '@/components/Login/LoginOptions'

describe('LoginOptions component', () => {
  const goToHomeMock = jest.fn()

  const renderLoginOptions = () => {
    render(
      <Form>
        <LoginOptions goToHome={goToHomeMock} />
      </Form>
    )
  }

  test('renders the "Keep me signed" checkbox', () => {
    renderLoginOptions()

    const checkbox = screen.getByLabelText(/keep me signed/i)
    expect(checkbox).toBeInTheDocument()
  })

  test('renders the "Forgot password?" button', () => {
    renderLoginOptions()

    const forgotPasswordButton = screen.getByText(/forgot password \?/i)
    expect(forgotPasswordButton).toBeInTheDocument()
  })

  test('calls goToHome when "Forgot password?" button is clicked', () => {
    renderLoginOptions()

    const forgotPasswordButton = screen.getByText(/forgot password \?/i)
    fireEvent.click(forgotPasswordButton)

    expect(goToHomeMock).toHaveBeenCalledTimes(1)
  })
})
