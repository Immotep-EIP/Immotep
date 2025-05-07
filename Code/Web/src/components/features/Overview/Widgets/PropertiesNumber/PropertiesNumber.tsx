import React from 'react'
import { LoadingOutlined } from '@ant-design/icons'

import { useTranslation } from 'react-i18next'
import useProperties from '@/hooks/Property/useProperties'
import { WidgetProps } from '@/interfaces/Widgets/Widgets.ts'
import PropertiesIcon from '@/assets/icons/realProperty.svg'
import style from './PropertiesNumber.module.css'

const PropertiesNumber: React.FC<WidgetProps> = ({ height }) => {
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
          {properties.length && properties.length < 9
            ? `0${properties.length}`
            : properties.length}
        </span>
        <span className={style.propertiesNumberText}>
          {t('widgets.properties_number.real_properties')}
        </span>
      </div>
    </div>
  )
}

export default PropertiesNumber
