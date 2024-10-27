export interface UserRegister {
  firstname: string
  lastname: string
  email: string
  password: string
  confirmPassword: string
}

export interface UserToken {
  username?: string
  password?: string
  grant_type: string
  refresh_token?: string
  rememberMe?: boolean
}

export interface TokenResponse {
  access_token: string
  refresh_token: string
  expires_in: number
}

export interface User {
  id: string
  email: string
  firstname: string
  lastname: string
  role: string
  created_at: Date
  updated_at: Date
}
