import { useNavigate } from 'react-router-dom'

const useNavigation = () => {
  const navigate = useNavigate()

  const goToLogin = () => {
    navigate('/login')
  }

  const goToSignup = () => {
    navigate('/register')
  }

  const goToHome = () => {
    navigate('/')
  }

  const goToDashboard = () => {
    navigate('/dashboard')
  }

  return {
    goToLogin,
    goToSignup,
    goToHome,
    goToDashboard
  }
}

export default useNavigation
