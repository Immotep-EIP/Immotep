import React from 'react'
import { LoadingOutlined } from '@ant-design/icons'

import { useTranslation } from 'react-i18next'
import useProperties from '@/hooks/Property/useProperties'
import { WidgetProps } from '@/interfaces/Widgets/Widgets.ts'
import PropertiesIcon from '@/assets/icons/realProperty.svg'
import style from './PropertiesRepartition.module.css'

const PropertiesRepartition: React.FC<WidgetProps> = ({ height }) => {
  const rowHeight = 120
  const pixelHeight = height * rowHeight
  const { t } = useTranslation()
  const { properties, loading, error } = useProperties()

  if (loading) {
    return (
      <div>
        <p>{t('generals.loading')}</p>
        <LoadingOutlined />
      </div>
    )
  }

  if (error) {
    return <p>{t('widgets.user_info.error_fetching')}</p>
  }

  const availableProperties = properties.filter(
    property => property.lease === null
  ).length
  const unavailableProperties = properties.filter(
    property => property.lease !== null
  ).length

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
            {availableProperties < 9
              ? `0${availableProperties}`
              : availableProperties || '-'}
          </span>
          <span
            className={style.propertiesNumberText}
            style={{ color: '#4caf50' }}
          >
            {availableProperties > 1
              ? t('pages.real_property.status.availables')
              : t('pages.real_property.status.available')}
          </span>
        </div>
        <div className={style.separator} />
        <div>
          <span className={style.propertiesNumber}>
            {unavailableProperties < 9
              ? `0${unavailableProperties}`
              : unavailableProperties || '-'}
          </span>
          <span
            className={style.propertiesNumberText}
            style={{ color: '#f44336' }}
          >
            {availableProperties > 1
              ? t('pages.real_property.status.unavailables')
              : t('pages.real_property.status.unavailable')}
          </span>
        </div>
      </div>
    </div>
  )
}

export default PropertiesRepartition
