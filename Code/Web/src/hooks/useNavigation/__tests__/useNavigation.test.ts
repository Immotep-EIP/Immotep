import { renderHook } from '@testing-library/react'
import { useNavigate } from 'react-router-dom'
import NavigationEnum from '@/enums/NavigationEnum'
import useNavigation from '../useNavigation'

// Mock react-router-dom
jest.mock('react-router-dom', () => ({
  useNavigate: jest.fn()
}))

describe('useNavigation', () => {
  const mockNavigate = jest.fn()

  beforeEach(() => {
    jest.clearAllMocks()
    ;(useNavigate as jest.Mock).mockReturnValue(mockNavigate)
  })

  it('should navigate to login page', () => {
    const { result } = renderHook(() => useNavigation())
    result.current.goToLogin()
    expect(mockNavigate).toHaveBeenCalledWith(NavigationEnum.LOGIN)
  })

  it('should navigate to signup page', () => {
    const { result } = renderHook(() => useNavigation())
    result.current.goToSignup()
    expect(mockNavigate).toHaveBeenCalledWith(NavigationEnum.REGISTER)
  })

  it('should navigate to forgot password page', () => {
    const { result } = renderHook(() => useNavigation())
    result.current.goToForgotPassword()
    expect(mockNavigate).toHaveBeenCalledWith(NavigationEnum.FORGOT_PASSWORD)
  })

  it('should navigate to overview page', () => {
    const { result } = renderHook(() => useNavigation())
    result.current.goToOverview()
    expect(mockNavigate).toHaveBeenCalledWith(NavigationEnum.OVERVIEW)
  })

  it('should navigate to real property page', () => {
    const { result } = renderHook(() => useNavigation())
    result.current.goToRealProperty()
    expect(mockNavigate).toHaveBeenCalledWith(NavigationEnum.REAL_PROPERTY)
  })

  it('should navigate to real property create page', () => {
    const { result } = renderHook(() => useNavigation())
    result.current.goToRealPropertyCreate()
    expect(mockNavigate).toHaveBeenCalledWith(
      NavigationEnum.REAL_PROPERTY_CREATE
    )
  })

  it('should navigate to real property details page with id', () => {
    const { result } = renderHook(() => useNavigation())
    const testId = '123'
    result.current.goToRealPropertyDetails(testId)
    expect(mockNavigate).toHaveBeenCalledWith(
      NavigationEnum.REAL_PROPERTY_DETAILS,
      {
        state: { id: testId }
      }
    )
  })

  it('should navigate to messages page', () => {
    const { result } = renderHook(() => useNavigation())
    result.current.goToMessages()
    expect(mockNavigate).toHaveBeenCalledWith(NavigationEnum.MESSAGES)
  })

  it('should navigate to settings page', () => {
    const { result } = renderHook(() => useNavigation())
    result.current.goToSettings()
    expect(mockNavigate).toHaveBeenCalledWith(NavigationEnum.SETTINGS)
  })

  it('should navigate to my profile page', () => {
    const { result } = renderHook(() => useNavigation())
    result.current.goToMyProfile()
    expect(mockNavigate).toHaveBeenCalledWith(NavigationEnum.MY_PROFILE)
  })
})
