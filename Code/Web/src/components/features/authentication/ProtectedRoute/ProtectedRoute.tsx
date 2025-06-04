import React, { useEffect } from 'react'
import { Outlet } from 'react-router-dom'

import { useAuth } from '@/context/authContext'

const ProtectedRoute: React.FC = () => {
  const { isAuthenticated, logout } = useAuth()
  useEffect(() => {
    if (!isAuthenticated) logout()
  }, [isAuthenticated, logout])
  return <Outlet />
}

export default ProtectedRoute
