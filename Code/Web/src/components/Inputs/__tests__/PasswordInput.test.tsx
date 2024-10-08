import '@testing-library/jest-dom'
import { render, screen, fireEvent } from '@testing-library/react'
import { Form, Button } from 'antd'
import PasswordInput from '@/components/Inputs/PasswordInput'

describe('PasswordInput component', () => {
  test('renders PasswordInput with label and placeholder', () => {
    render(
      <Form>
        <PasswordInput />
      </Form>
    )

    const passwordLabel = screen.getByLabelText(/password/i)
    const passwordInput = screen.getByPlaceholderText(/enter your password/i)

    expect(passwordLabel).toBeInTheDocument()
    expect(passwordInput).toBeInTheDocument()
  })

  test('shows error message when password is required and not filled', async () => {
    const onFinish = jest.fn()

    render(
      <Form onFinish={onFinish}>
        <PasswordInput />
        <Button type="primary" htmlType="submit">
          Sign in
        </Button>{' '}
      </Form>
    )

    fireEvent.click(screen.getByRole('button', { name: /sign in/i }))

    const errorMessage = await screen.findByText(/please enter your password/i)
    expect(errorMessage).toBeInTheDocument()
  })
})
