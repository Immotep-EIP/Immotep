import React from 'react'
import { render, screen, fireEvent } from '@testing-library/react'
import { HelmetProvider } from 'react-helmet-async'
import { useTranslation } from 'react-i18next'
import useNavigation from '@/hooks/Navigation/useNavigation'
import Lost from '../Lost'

jest.mock('react-i18next', () => ({
  useTranslation: jest.fn()
}))

jest.mock('@/hooks/Navigation/useNavigation', () => ({
  __esModule: true,
  default: jest.fn()
}))

describe('Lost', () => {
  const mockT = jest.fn(key => key)
  const mockGoToOverview = jest.fn()

  beforeEach(() => {
    jest.clearAllMocks()
    ;(useTranslation as jest.Mock).mockReturnValue({ t: mockT })
    ;(useNavigation as jest.Mock).mockReturnValue({
      goToOverview: mockGoToOverview
    })
  })

  const renderWithHelmet = (component: React.ReactNode) =>
    render(<HelmetProvider>{component}</HelmetProvider>)

  it('should render 404 page correctly', () => {
    renderWithHelmet(<Lost />)

    expect(screen.getByText('404')).toBeInTheDocument()
    expect(screen.getByText('pages.lost.page_not_found')).toBeInTheDocument()
    expect(screen.getByText('pages.lost.back_home')).toBeInTheDocument()
  })

  it('should set correct meta tags', () => {
    renderWithHelmet(<Lost />)

    expect(mockT).toHaveBeenCalledWith('pages.lost.document_title')
    expect(mockT).toHaveBeenCalledWith('pages.lost.document_description')
  })

  it('should navigate to overview when button is clicked', () => {
    renderWithHelmet(<Lost />)

    const button = screen.getByText('pages.lost.back_home')
    fireEvent.click(button)

    expect(mockGoToOverview).toHaveBeenCalled()
  })
})
