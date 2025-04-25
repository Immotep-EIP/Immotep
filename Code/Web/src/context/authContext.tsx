import React, {
  createContext,
  useContext,
  useState,
  useEffect,
  useMemo
} from 'react'
import { useNavigate } from 'react-router-dom'
import { LoadingOutlined } from '@ant-design/icons'

import { UserTokenPayload, User } from '@/interfaces/User/User'
import { loginApi } from '@/services/api/Authentification/AuthApi'
import getUserProfile from '@/services/api/User/GetUserProfile'
import { saveData, deleteData } from '@/utils/cache/localStorage'
import { AuthContextType, AuthProviderProps } from '@/interfaces/Auth/Auth'
import NavigationEnum from '@/enums/NavigationEnum'

const AuthContext = createContext<AuthContextType | undefined>(undefined)

export const AuthProvider: React.FC<AuthProviderProps> = ({ children }) => {
  const [isAuthenticated, setIsAuthenticated] = useState(false)
  const [loading, setLoading] = useState(true)
  const [user, setUser] = useState<User | null>(null)
  const navigate = useNavigate()

  const updateUser = async (newUserData: Partial<User>) => {
    if (user) {
      const updatedUser = { ...user, ...newUserData }
      setUser(updatedUser)
    }
  }

  useEffect(() => {
    const initializeAuth = async () => {
      const accessToken =
        sessionStorage.getItem('access_token') ||
        localStorage.getItem('access_token')

      const refreshToken =
        sessionStorage.getItem('refresh_token') ||
        localStorage.getItem('refresh_token')

      try {
        if (accessToken && refreshToken) {
          setIsAuthenticated(true)
          const profile = await getUserProfile()
          setUser(profile)
        } else {
          setIsAuthenticated(false)
          deleteData()
        }
      } catch (err) {
        console.error('Error during auth initialization:', err)
        setIsAuthenticated(false)
        deleteData()
      } finally {
        setLoading(false)
      }
    }

    initializeAuth()
  }, [])

  const login = async (userInfo: UserTokenPayload) => {
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
    if ('serviceWorker' in navigator) {
      navigator.serviceWorker.ready.then(registration => {
        if (registration.active) {
          registration.active.postMessage({
            type: 'LOGOUT'
          })
        }
      })
    }
    deleteData()
    navigate(NavigationEnum.LOGIN)
  }

  const value = useMemo(
    () => ({
      isAuthenticated,
      login,
      logout,
      user,
      updateUser
    }),
    [isAuthenticated, user]
  )

  return (
    <AuthContext.Provider value={value}>
      {!loading ? (
        children
      ) : (
        <div>
          <LoadingOutlined />
        </div>
      )}{' '}
    </AuthContext.Provider>
  )
}

export const useAuth = () => {
  const context = useContext(AuthContext)
  if (context === undefined)
    throw new Error('useAuth must be used within an AuthProvider')
  return context
}
