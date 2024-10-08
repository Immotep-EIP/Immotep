import '@testing-library/jest-dom'
import { render, screen } from '@testing-library/react'
import LoginTitle from '@/components/Login/LoginTitle'

describe('LoginTitle component', () => {
  test('renders the "Welcome back" title', () => {
    render(<LoginTitle />)

    const title = screen.getByText(/welcome back/i)
    expect(title).toBeInTheDocument()
  })

  test('renders the sign-in prompt', () => {
    render(<LoginTitle />)

    const text = screen.getByText(/please enter your details to sign in/i)
    expect(text).toBeInTheDocument()
  })
})
