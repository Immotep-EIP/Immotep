import React from 'react'
import { LoadingOutlined } from '@ant-design/icons'

import { useTranslation } from 'react-i18next'
import PropertiesIcon from '@/assets/icons/realProperty.svg'
import style from './PropertiesNumber.module.css'
import { DashboardProperties } from '@/interfaces/Dashboard/Dashboard'

interface PropertiesNumberProps {
  properties: DashboardProperties | null
  loading: boolean
  error: string | null
  height: number
}

const PropertiesNumber: React.FC<PropertiesNumberProps> = ({
  properties,
  loading,
  error,
  height
}: PropertiesNumberProps) => {
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
      <div className={style.informationsContainer}>
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
        <span className={style.propertiesNumber}>
          {properties?.nbr_total && properties?.nbr_total < 9
            ? `0${properties?.nbr_total}`
            : properties?.nbr_total}
        </span>
        <span className={style.propertiesNumberText}>
          {t('widgets.properties_number.real_properties')}
        </span>
      </div>
    </div>
  )
}

export default PropertiesNumber
