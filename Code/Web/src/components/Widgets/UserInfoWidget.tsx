import React from 'react'
import { LoadingOutlined } from '@ant-design/icons'

import { useTranslation } from 'react-i18next'
import useProperties from '@/hooks/useEffect/useProperties.ts'
import { useAuth } from '@/context/authContext'
import { WidgetProps } from '@/interfaces/Widgets/Widgets.ts'
import style from './UserInfoWidget.module.css'

const UserInfoWidget: React.FC<WidgetProps> = ({ height }) => {
  const rowHeight = 70
  const pixelHeight = height * rowHeight
  const { user } = useAuth()
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
      {user ? (
        <div key={user.id}>
          <p>
            {t('widgets.user_info.title')} {user.firstname} !
          </p>
          <p>
            {[
              t('widgets.user_info.properties_number'),
              properties.length,
              t('widgets.user_info.real_properties')
            ].join(' ')}
          </p>
        </div>
      ) : (
        <p>{t('widgets.user_info.noUser')}</p>
      )}
    </div>
  )
}

export default UserInfoWidget
