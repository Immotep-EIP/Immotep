import { render, screen } from '@testing-library/react'
import '@testing-library/jest-dom'
import AuthentificationPage from '@/components/AuthentificationPage/AuthentificationPage'

jest.mock('@/assets/icons/ImmotepLogo.svg', () => 'mockedLogoPath')

describe('AuthentificationPage', () => {
  const title = 'Login'
  const subtitle = 'Please enter your credentials'

  it('renders the page title and subtitle', () => {
    render(
      <AuthentificationPage title={title} subtitle={subtitle}>
        <div>Content</div>
      </AuthentificationPage>
    )

    expect(screen.getByText(title)).toBeInTheDocument()
    expect(screen.getByText(subtitle)).toBeInTheDocument()
  })

  it('renders children correctly', () => {
    render(
      <AuthentificationPage title={title} subtitle={subtitle}>
        <div>Content</div>
      </AuthentificationPage>
    )

    expect(screen.getByText('Content')).toBeInTheDocument()
  })

  it('renders the logo', () => {
    render(
      <AuthentificationPage title={title} subtitle={subtitle}>
        <div>Content</div>
      </AuthentificationPage>
    )

    const logo = screen.getByAltText('logo Immotep')
    expect(logo).toBeInTheDocument()
    expect(logo).toHaveAttribute('src', 'mockedLogoPath')
  })
})
