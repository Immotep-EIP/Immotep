import { saveData, deleteData } from '@/utils/cache/localStorage'

describe('saveData', () => {
  beforeEach(() => {
    localStorage.clear()
    sessionStorage.clear()

    jest.spyOn(Date, 'now').mockReturnValue(1739049951351)
  })

  afterEach(() => {
    jest.restoreAllMocks()
  })

  it('should save data to sessionStorage when rememberMe is false', () => {
    const accessToken = 'test-access-token'
    const refreshToken = 'test-refresh-token'
    const expiresIn = 3600

    saveData(accessToken, refreshToken, expiresIn, false)

    expect(sessionStorage.getItem('access_token')).toBe(accessToken)
    expect(sessionStorage.getItem('refresh_token')).toBe(refreshToken)
    expect(sessionStorage.getItem('expires_in')).toBe(
      (Date.now() + expiresIn * 1000).toString()
    )

    expect(localStorage.getItem('access_token')).toBeNull()
    expect(localStorage.getItem('refresh_token')).toBeNull()
    expect(localStorage.getItem('expires_in')).toBeNull()
  })

  it('should save data to localStorage when rememberMe is true', () => {
    const accessToken = 'test-access-token'
    const refreshToken = 'test-refresh-token'
    const expiresIn = 3600

    saveData(accessToken, refreshToken, expiresIn, true)

    expect(localStorage.getItem('access_token')).toBe(accessToken)
    expect(localStorage.getItem('refresh_token')).toBe(refreshToken)
    expect(localStorage.getItem('expires_in')).toBe(
      (Date.now() + expiresIn * 1000).toString()
    )

    expect(sessionStorage.getItem('access_token')).toBeNull()
    expect(sessionStorage.getItem('refresh_token')).toBeNull()
    expect(sessionStorage.getItem('expires_in')).toBeNull()
  })

  it('should use sessionStorage by default when rememberMe is not provided', () => {
    const accessToken = 'test-access-token'
    const refreshToken = 'test-refresh-token'
    const expiresIn = 3600

    saveData(accessToken, refreshToken, expiresIn)

    expect(sessionStorage.getItem('access_token')).toBe(accessToken)
    expect(sessionStorage.getItem('refresh_token')).toBe(refreshToken)
    expect(sessionStorage.getItem('expires_in')).toBe(
      (Date.now() + expiresIn * 1000).toString()
    )

    expect(localStorage.getItem('access_token')).toBeNull()
    expect(localStorage.getItem('refresh_token')).toBeNull()
    expect(localStorage.getItem('expires_in')).toBeNull()
  })
})

describe('deleteData', () => {
  beforeEach(() => {
    localStorage.clear()
    sessionStorage.clear()
  })

  it('should delete data from both localStorage and sessionStorage', () => {
    localStorage.setItem('access_token', 'test-access-token')
    localStorage.setItem('refresh_token', 'test-refresh-token')
    localStorage.setItem('expires_in', '123456789')

    sessionStorage.setItem('access_token', 'test-access-token')
    sessionStorage.setItem('refresh_token', 'test-refresh-token')
    sessionStorage.setItem('expires_in', '123456789')

    deleteData()

    expect(localStorage.getItem('access_token')).toBeNull()
    expect(localStorage.getItem('refresh_token')).toBeNull()
    expect(localStorage.getItem('expires_in')).toBeNull()

    expect(sessionStorage.getItem('access_token')).toBeNull()
    expect(sessionStorage.getItem('refresh_token')).toBeNull()
    expect(sessionStorage.getItem('expires_in')).toBeNull()
  })
})
