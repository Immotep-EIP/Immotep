import { useTranslation } from 'react-i18next'

import { Result } from 'antd'

import logo from '@/assets/images/KeyzLogo.svg'
import style from './SuccesPageRegisterTenant.module.css'

const SuccessPage = () => {
  const { t } = useTranslation()

  return (
    <div className={style.successPage}>
      <div className={style.headerContainer}>
        <img src={logo} alt="logo Keyz" className={style.headerLogo} />
        <span className={style.headerTitle}>Keyz</span>
      </div>
      <div className={style.contentContainer}>
        <Result
          status="success"
          title={t('pages.register_tenant.register_success')}
          subTitle={t('pages.register_tenant.register_success_description')}
        />
      </div>
    </div>
  )
}

export default SuccessPage
