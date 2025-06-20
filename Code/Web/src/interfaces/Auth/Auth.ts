import React from 'react'

import { TokenResponse, UserTokenPayload, User } from '../User/User'

export interface AuthentificationPageProps {
  title: string
  subtitle: string
  children: React.ReactNode
}

export interface AuthProviderProps {
  children: React.ReactNode
}

export interface AuthContextType {
  isAuthenticated: boolean
  login: (user: UserTokenPayload, leaseId: string | undefined) => Promise<TokenResponse>
  logout: () => void
  user: User | null
  updateUser: (newUserData: Partial<User>) => void
}
