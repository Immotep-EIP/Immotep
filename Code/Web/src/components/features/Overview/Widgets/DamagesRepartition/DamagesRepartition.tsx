import React from 'react'
import { useTranslation } from 'react-i18next'

import { LoadingOutlined } from '@ant-design/icons'

import { DashboardOpenDamages } from '@/interfaces/Dashboard/Dashboard'

import PropertiesIcon from '@/assets/icons/realProperty.svg'
import style from './DamagesRepartition.module.css'

interface DamagesRepartitionProps {
  openDamages: DashboardOpenDamages | null
  loading: boolean
  error: string | null
  height: number
}

const DamagesRepartition: React.FC<DamagesRepartitionProps> = ({
  openDamages,
  loading,
  error,
  height
}) => {
  const { t } = useTranslation()
  const rowHeight = 120
  const pixelHeight = height * rowHeight

  if (loading || openDamages === null) {
    return (
      <div>
        <p>{t('components.loading.loading_data')}</p>
        <LoadingOutlined />
      </div>
    )
  }

  if (error) {
    return <p>{t('widgets.user_info.error_fetching')}</p>
  }

  return (
    <div
      className={style.layoutContainer}
      style={{ height: `${pixelHeight}px` }}
    >
      <div className={style.informationsIcon}>
        <div className={style.informationsIconContainer}>
          <img
            src={PropertiesIcon}
            alt="properties icon"
            style={{
              width: '20px'
            }}
          />
        </div>
      </div>
      <div className={style.contentContainer}>
        <div>
          <span className={style.damagesNumber}>
            {openDamages.nbr_total ?? 0}
          </span>
          <span className={style.damagesNumberText}>
            {(openDamages.nbr_total ?? 0) > 1
              ? t('pages.damage_details.title_damages')
              : t('pages.damage_details.title_damage')}
          </span>
        </div>
        <div className={style.separator} />
        <div>
          <span className={style.damagesNumber}>
            {openDamages.nbr_urgent ?? 0}
            <span className={style.statusText} style={{ color: 'red' }}>
              {t('pages.damage_details.status.urgent')}
            </span>
            {openDamages.nbr_high ?? 0}
            <span className={style.statusText} style={{ color: 'red' }}>
              {t('pages.damage_details.status.high')}
            </span>
            {openDamages.nbr_medium ?? 0}
            <span className={style.statusText} style={{ color: '#F79009' }}>
              {t('pages.damage_details.status.medium')}
            </span>
            {openDamages.nbr_low ?? 0}
            <span className={style.statusText} style={{ color: 'green' }}>
              {t('pages.damage_details.status.low')}
            </span>
          </span>
        </div>
      </div>
    </div>
  )
}

export default DamagesRepartition
