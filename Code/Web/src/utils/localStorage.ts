export const saveData = (
  accessToken: string,
  refreshToken: string,
  expiresIn: number
) => {
  const expiryTime = Date.now() + expiresIn * 1000
  localStorage.setItem('access_token', accessToken)
  localStorage.setItem('refresh_token', refreshToken)
  localStorage.setItem('expires_in', expiryTime.toString())
}

export const deleteData = () => {
  localStorage.removeItem('access_token')
  localStorage.removeItem('refresh_token')
  localStorage.removeItem('expires_in')
}
