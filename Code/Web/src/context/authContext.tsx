import React, {
  createContext,
  useContext,
  useState,
  useEffect,
  ReactNode,
  useMemo
} from 'react'
import { useNavigate } from 'react-router-dom'

import { UserToken, TokenResponse } from '@/interfaces/User/User'
import { loginApi } from '@/services/api/Authentification/AuthApi'
import { saveData, deleteData } from '@/utils/localStorage'

interface AuthContextType {
  isAuthenticated: boolean
  login: (user: UserToken) => Promise<TokenResponse>
  logout: () => void
}

const AuthContext = createContext<AuthContextType | undefined>(undefined)

interface AuthProviderProps {
  children: ReactNode
}

export const AuthProvider: React.FC<AuthProviderProps> = ({ children }) => {
  const [isAuthenticated, setIsAuthenticated] = useState(false)
  const [loading, setLoading] = useState(true)
  const navigate = useNavigate()

  useEffect(() => {
    const accessToken = localStorage.getItem('access_token')
    const refreshToken = localStorage.getItem('refresh_token')

    if (accessToken && refreshToken) setIsAuthenticated(true)
    else setIsAuthenticated(false)
    setLoading(false)
  }, [])

  const login = async (userInfo: UserToken) => {
    try {
      const response = await loginApi(userInfo)
      setIsAuthenticated(true)
      saveData(
        response.access_token,
        response.refresh_token,
        response.expires_in
      )
      return response
    } catch (error) {
      console.error('login error:', error)
      deleteData()
      setIsAuthenticated(false)
      throw error
    }
  }

  const logout = () => {
    setIsAuthenticated(false)
    deleteData()
    navigate('/')
  }

  const value = useMemo(
    () => ({
      isAuthenticated,
      login,
      logout
    }),
    [isAuthenticated]
  )

  return (
    <AuthContext.Provider value={value}>
      {!loading ? children : <div>Loading...</div>}{' '}
    </AuthContext.Provider>
  )
}

export const useAuth = () => {
  const context = useContext(AuthContext)
  if (context === undefined)
    throw new Error('useAuth must be used within an AuthProvider')
  return context
}
