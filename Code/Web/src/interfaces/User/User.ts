export interface UserRegisterPayload {
  firstname: string
  lastname: string
  email: string
  password: string
  confirmPassword: string
  leaseId?: string
}

export interface UserTokenPayload {
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
  properties: {
    id: string
    role: string
  }
}

export interface User {
  id: string
  email: string
  firstname: string
  lastname: string
  role: string
  profile_picture_id: string
  created_at: Date
  updated_at: Date
}

export interface UserPictureResponse {
  id: string
  created_at: string
  data: string
}

export interface UpdateUserInfoPayload {
  firstname: string
  lastname: string
}
