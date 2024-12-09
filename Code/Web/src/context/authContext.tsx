import React, {
  createContext,
  useContext,
  useState,
  useEffect,
  ReactNode,
  useMemo
} from 'react'
import { useNavigate } from 'react-router-dom'
import { LoadingOutlined } from '@ant-design/icons';


import { UserToken, TokenResponse, User } from '@/interfaces/User/User'
import { loginApi } from '@/services/api/Authentification/AuthApi'
import getUserProfile from '@/services/api/User/GetUserProfile'
import { saveData, deleteData } from '@/utils/localStorage'
import NavigationEnum from '@/enums/NavigationEnum'

interface AuthContextType {
  isAuthenticated: boolean
  login: (user: UserToken) => Promise<TokenResponse>
  logout: () => void
  user: User | null
}

const AuthContext = createContext<AuthContextType | undefined>(undefined)

interface AuthProviderProps {
  children: ReactNode
}

export const AuthProvider: React.FC<AuthProviderProps> = ({ children }) => {
  const [isAuthenticated, setIsAuthenticated] = useState(false)
  const [loading, setLoading] = useState(true)
  const [user, setUser] = useState<User | null>(null)
  const navigate = useNavigate()

  useEffect(() => {
    const accessToken =
      sessionStorage.getItem('access_token') ||
      localStorage.getItem('access_token')

    const refreshToken =
      sessionStorage.getItem('refresh_token') ||
      localStorage.getItem('refresh_token')

    const userInfo = async () => {
      const profile = await getUserProfile()
      setUser(profile)
    }

    if (accessToken && refreshToken) {
      setIsAuthenticated(true)
      userInfo()
    } else {
      setIsAuthenticated(false)
      deleteData()
    }
    setLoading(false)
  }, [])

  const login = async (userInfo: UserToken) => {
    try {
      const response = await loginApi(userInfo)
      setIsAuthenticated(true)
      saveData(
        response.access_token,
        response.refresh_token,
        response.expires_in,
        userInfo.rememberMe
      )
      const profile = await getUserProfile()
      setUser(profile)
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
    navigate(NavigationEnum.LOGIN)
  }

  const value = useMemo(
    () => ({
      isAuthenticated,
      login,
      logout,
      user
    }),
    [isAuthenticated, user]
  )

  return (
    <AuthContext.Provider value={value}>
      {!loading ? children : <div><LoadingOutlined /></div>}{' '}
    </AuthContext.Provider>
  )
}

export const useAuth = () => {
  const context = useContext(AuthContext)
  if (context === undefined)
    throw new Error('useAuth must be used within an AuthProvider')
  return context
}
