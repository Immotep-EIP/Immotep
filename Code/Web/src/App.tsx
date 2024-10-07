import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'

import '@/App.css'
import PublicLayout from '@/components/PublicLayout/PublicLayout'
import AuthLayout from '@/components/AuthLayout/AuthLayout'
import Login from '@/views/Login/Login'
import Register from '@/views/Register/Register'
import Home from '@/views/Home/Home'
import Dashboard from '@/views/Dashboard/Dashboard'

const App = () => (
  <Router>
    <Routes>
      <Route
        path="/"
        element={
          <PublicLayout>
            <Home />
          </PublicLayout>
        }
      />
      <Route
        path="/register"
        element={
          <PublicLayout>
            <Register />
          </PublicLayout>
        }
      />
      <Route
        path="/login"
        element={
          <PublicLayout>
            <Login />
          </PublicLayout>
        }
      />
      <Route
        path="/dashboard"
        element={
          <AuthLayout>
            <Dashboard />
          </AuthLayout>
        }
      />
    </Routes>
  </Router>
)

export default App
