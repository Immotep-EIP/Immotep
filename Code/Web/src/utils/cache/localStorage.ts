export const saveData = (
  accessToken: string,
  refreshToken: string,
  expiresIn: number,
  rememberMe: boolean = false
) => {
  const storage = rememberMe ? localStorage : sessionStorage
  const expiryTime = Date.now() + expiresIn * 1000
  storage.setItem('access_token', accessToken)
  storage.setItem('refresh_token', refreshToken)
  storage.setItem('expires_in', expiryTime.toString())
}

export const deleteData = () => {
  sessionStorage.removeItem('access_token')
  sessionStorage.removeItem('refresh_token')
  sessionStorage.removeItem('expires_in')

  localStorage.removeItem('access_token')
  localStorage.removeItem('refresh_token')
  localStorage.removeItem('expires_in')
}
