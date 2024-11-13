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

  const goToForgotPassword = () => {
    navigate(NavigationEnum.FORGOT_PASSWORD)
  }

  const goToOverview = () => {
    navigate(NavigationEnum.OVERVIEW)
  }

  const goToRealProperty = () => {
    navigate(NavigationEnum.REAL_PROPERTY)
  }

  const goToRealPropertyDetails = (id: number) => {
    navigate(`${NavigationEnum.REAL_PROPERTY_DETAILS.replace(':id', String(id))}`)
  }

  const goToMessages = () => {
    navigate(NavigationEnum.MESSAGES)
  }

  const goToSettings = () => {
    navigate(NavigationEnum.SETTINGS)
  }

  const goToMyProfile = () => {
    navigate(NavigationEnum.MY_PROFILE)
  }

  return {
    goToLogin,
    goToSignup,
    goToForgotPassword,
    goToOverview,
    goToRealProperty,
    goToRealPropertyDetails,
    goToMessages,
    goToSettings,
    goToMyProfile
  }
}

export default useNavigation
