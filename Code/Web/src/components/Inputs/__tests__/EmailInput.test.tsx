import '@testing-library/jest-dom'
import { render, screen, fireEvent } from '@testing-library/react'
import { Form, Button } from 'antd'
import EmailInput from '@/components/Inputs/EmailInput'

describe('EmailInput component', () => {
  test('renders EmailInput with label and placeholder', () => {
    render(
      <Form>
        <EmailInput />
      </Form>
    )

    const emailLabel = screen.getByLabelText(/email/i)
    const emailInput = screen.getByPlaceholderText(/enter your email/i)

    expect(emailLabel).toBeInTheDocument()
    expect(emailInput).toBeInTheDocument()
  })

  test('shows error message when email is required and not filled', async () => {
    const onFinish = jest.fn()

    render(
      <Form onFinish={onFinish}>
        <EmailInput />
        <Button type="primary" htmlType="submit">
          Sign in
        </Button>{' '}
      </Form>
    )

    fireEvent.click(screen.getByRole('button', { name: /sign in/i }))

    const errorMessage = await screen.findByText(/please enter your email/i)
    expect(errorMessage).toBeInTheDocument()
  })
})
