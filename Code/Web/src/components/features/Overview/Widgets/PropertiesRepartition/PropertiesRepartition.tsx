import React from 'react'
import { LoadingOutlined } from '@ant-design/icons'

import { useTranslation } from 'react-i18next'
import PropertiesIcon from '@/assets/icons/realProperty.svg'
import style from './PropertiesRepartition.module.css'
import { DashboardProperties } from '@/interfaces/Dashboard/Dashboard'

interface PropertiesRepartitionProps {
  properties: DashboardProperties | null
  loading: boolean
  error: string | null
  height: number
}

const PropertiesRepartition: React.FC<PropertiesRepartitionProps> = ({
  properties,
  loading,
  error,
  height
}) => {
  const rowHeight = 120
  const pixelHeight = height * rowHeight
  const { t } = useTranslation()

  if (loading || properties === null) {
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
          <span className={style.propertiesNumber}>
            {properties?.nbr_total ?? 0}
          </span>
          <span
            className={style.propertiesNumberText}
            style={{ color: '#4caf50' }}
          >
            {(properties?.nbr_available ?? 0) > 1
              ? t('pages.real_property.status.availables')
              : t('pages.real_property.status.available')}
          </span>
        </div>
        <div className={style.separator} />
        <div>
          <span className={style.propertiesNumber}>
            {properties?.nbr_occupied ?? 0}
          </span>
          <span
            className={style.propertiesNumberText}
            style={{ color: '#f44336' }}
          >
            {(properties?.nbr_occupied ?? 0) > 1
              ? t('pages.real_property.status.unavailables')
              : t('pages.real_property.status.unavailable')}
          </span>
        </div>
      </div>
    </div>
  )
}

export default PropertiesRepartition
