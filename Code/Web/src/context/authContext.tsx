import React, {
  createContext,
  useContext,
  useState,
  useEffect,
  ReactNode,
  useMemo
} from 'react'
import { useNavigate } from 'react-router-dom'
import { LoadingOutlined } from '@ant-design/icons'

import { UserToken, TokenResponse, User } from '@/interfaces/User/User'
import { loginApi } from '@/services/api/Authentification/AuthApi'
import getUserProfile from '@/services/api/User/GetUserProfile'
import { saveData, deleteData } from '@/utils/localStorage'
import NavigationEnum from '@/enums/NavigationEnum'
import { getUserFromDB, saveUserToDB } from '@/utils/cache/user/indexDB'

interface AuthContextType {
  isAuthenticated: boolean
  login: (user: UserToken) => Promise<TokenResponse>
  logout: () => void
  user: User | null
  updateUser: (newUserData: Partial<User>) => void
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

  const updateUser = async (newUserData: Partial<User>) => {
    if (user) {
      const updatedUser = { ...user, ...newUserData }
      setUser(updatedUser)
      await saveUserToDB(updatedUser)
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

          const cachedUser = await getUserFromDB()
          if (cachedUser) {
            setUser(cachedUser)
          } else {
            const profile = await getUserProfile()
            setUser(profile)
            await saveUserToDB(profile)
          }
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
      await saveUserToDB(profile)
      return response
    } catch (error) {
      console.error('login error:', error)
      deleteData()
      setIsAuthenticated(false)
      throw error
    }
  }

  const deleteAllDatabases = async () => {
    const databases = await window.indexedDB.databases()
    databases.forEach(db => {
      if (db.name) {
        indexedDB.deleteDatabase(db.name)
      }
    })
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
    deleteAllDatabases()
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
