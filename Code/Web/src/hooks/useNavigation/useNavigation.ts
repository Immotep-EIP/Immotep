import { useNavigate } from 'react-router-dom'
import NavigationEnum from '@/enums/NavigationEnum'

const useNavigation = () => {
  const navigate = useNavigate()

  const goToLogin = () => {
    navigate(NavigationEnum.LOGIN)
  }

  const goToSignup = () => {
    navigate(NavigationEnum.REGISTER)
  }

  const goToOverview = () => {
    navigate(NavigationEnum.OVERVIEW)
  }

  const goToRealProperty = () => {
    navigate(NavigationEnum.REAL_PROPERTY)
  }

  const goToMessages = () => {
    navigate(NavigationEnum.MESSAGES)
  }

  return {
    goToLogin,
    goToSignup,
    goToOverview,
    goToRealProperty,
    goToMessages
  }
}

export default useNavigation
