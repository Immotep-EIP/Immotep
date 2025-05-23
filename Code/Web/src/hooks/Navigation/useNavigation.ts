import { useNavigate } from 'react-router-dom'

import NavigationEnum from '@/enums/NavigationEnum'

const useNavigation = () => {
  const navigate = useNavigate()

  const goToLogin = () => {
    navigate(NavigationEnum.LOGIN)
  }

  const goToSignup = () => {
    navigate(NavigationEnum.REGISTER_WITHOUT_CONTRACT)
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

  const goToRealPropertyDetails = (id: string) => {
    navigate(NavigationEnum.REAL_PROPERTY_DETAILS.replace(':id', id), {
      state: { id }
    })
  }

  const goToDamageDetails = (id: string, damageId: string) => {
    navigate(
      NavigationEnum.DAMAGE_DETAILS.replace(':id', id).replace(
        ':damageId',
        damageId
      ),
      { state: { id, damageId } }
    )
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

  const goToSuccessRegisterTenant = () => {
    navigate(NavigationEnum.SUCCESS_REGISTER_TENANT)
  }

  const goToSuccessLoginTenant = () => {
    navigate(NavigationEnum.SUCCESS_LOGIN_TENANT)
  }

  return {
    goToLogin,
    goToSignup,
    goToForgotPassword,
    goToOverview,
    goToRealProperty,
    goToRealPropertyDetails,
    goToDamageDetails,
    goToMessages,
    goToSettings,
    goToMyProfile,
    goToSuccessRegisterTenant,
    goToSuccessLoginTenant
  }
}

export default useNavigation
