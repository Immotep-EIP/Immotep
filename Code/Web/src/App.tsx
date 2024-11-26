import React from 'react'
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'

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
import RealPropertyCreate from '@/views/RealProperty/create/RealPropertyCreate.tsx'
import RealPropertyDetails from '@/views/RealProperty/details/RealPropertyDetails.tsx'
import MessagesContent from '@/views/Messages/Messages'

// ! MAIN LAYOUT - HEADER
import Settings from '@/views/Settings/Settings'
import MyProfile from '@/views/MyProfile/MyProfile'
import ProtectedRoute from './components/ProtectedRoute/ProtectedRoute'

import Lost from './views/Lost/Lost.tsx'

const App: React.FC = () => (
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

        <Route
          path={NavigationEnum.REAL_PROPERTY_CREATE}
          element={<RealPropertyCreate />}
        />

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
)

export default App
