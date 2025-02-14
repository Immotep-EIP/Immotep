import React from 'react'
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'
import { HelmetProvider } from 'react-helmet-async'

import './translation/i18n.tsx'

import NavigationEnum from '@/enums/NavigationEnum'
import MainLayout from '@/components/MainLayout/MainLayout.tsx'

// ! AUTHENTIFICATION
import { AuthProvider } from '@/context/authContext'
import Login from '@/views/Authentification/Login/Login.tsx'
import Register from '@/views/Authentification/Register/Register.tsx'
import ForgotPassword from '@/views/Authentification/ForgotPassword/ForgotPassword.tsx'

// ! MAIN LAYOUT - SIDEBAR
import OverviewContent from '@/views/Overview/Overview'
import RealPropertyContent from '@/views/RealProperty/RealProperty'
import RealPropertyDetails from '@/views/RealProperty/details/RealPropertyDetails.tsx'
import MessagesContent from '@/views/Messages/Messages'

// ! MAIN LAYOUT - HEADER
import Settings from '@/views/Settings/Settings'
import MyProfile from '@/views/MyProfile/MyProfile'
import ProtectedRoute from './components/ProtectedRoute/ProtectedRoute'

import Lost from './views/Lost/Lost.tsx'

if ('serviceWorker' in navigator) {
  ;(async () => {
    try {
      const registration = await navigator.serviceWorker.register('/sw.js', {
        scope: './'
      })
      if (registration.installing) {
        console.log('Service worker installing')
      } else if (registration.waiting) {
        console.log('Service worker installed')
      } else if (registration.active) {
        console.log('Service worker active')
      }
    } catch (error) {
      console.error(`Registration failed with ${error}`)
    }
  })()
}

const App: React.FC = () => (
  <HelmetProvider>
    <Router>
      <AuthProvider>
        <Routes>
          <Route path={NavigationEnum.LOGIN} element={<Login />} />
          <Route path={NavigationEnum.REGISTER} element={<Register />} />
          <Route path={NavigationEnum.REGISTER_TENANT} element={<Register />} />
          <Route
            path={NavigationEnum.FORGOT_PASSWORD}
            element={<ForgotPassword />}
          />
          <Route path="*" element={<Lost />} />

          <Route element={<MainLayout />}>
            <Route element={<ProtectedRoute />}>
              <Route
                path={NavigationEnum.OVERVIEW}
                element={<OverviewContent />}
              />
              <Route
                path={NavigationEnum.REAL_PROPERTY}
                element={<RealPropertyContent />}
              />
              <Route
                path={NavigationEnum.REAL_PROPERTY_DETAILS}
                element={<RealPropertyDetails />}
              />
              <Route
                path={NavigationEnum.MESSAGES}
                element={<MessagesContent />}
              />
              <Route path={NavigationEnum.SETTINGS} element={<Settings />} />
              <Route path={NavigationEnum.MY_PROFILE} element={<MyProfile />} />
            </Route>
          </Route>
        </Routes>
      </AuthProvider>
    </Router>
  </HelmetProvider>
)

export default App
