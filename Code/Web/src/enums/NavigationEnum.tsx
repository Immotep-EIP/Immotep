const NavigationEnum = {
  // ! AUTHENTIFICATION
  LOGIN: '/',
  LOGIN_TENANT: '/login/invite/:contractId',
  REGISTER_WITHOUT_CONTRACT: '/register',
  REGISTER_TENANT: '/register/invite/:contractId',
  FORGOT_PASSWORD: '/forgot-password',

  // ! MAIN LAYOUT - SIDEBAR
  OVERVIEW: '/overview',
  REAL_PROPERTY: '/real-property',
  REAL_PROPERTY_DETAILS: '/real-property/details',
  MESSAGES: '/messages',

  // ! MAIN LAYOUT - HEADER
  SETTINGS: '/settings',
  MY_PROFILE: '/my-profile',

  // ! SUCCESS PAGE
  SUCCESS_REGISTER_TENANT: '/success-register-tenant',
  SUCCESS_LOGIN_TENANT: '/success-login-tenant'
}

export default NavigationEnum
