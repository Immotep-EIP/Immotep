import React from 'react'
import { useTranslation } from 'react-i18next'
import style from './3DamageTab.module.css'

const DamageTab: React.FC = () => {
  const { t } = useTranslation()

  return (
    <div className={style.tabContent}>
      <span>{t('Damage')}</span>
    </div>
  )
}

export default DamageTab
