import React from 'react'
import { useTranslation } from 'react-i18next'
import { LoadingOutlined } from '@ant-design/icons'
import { Badge, Empty } from 'antd'
import {
  DashboardOpenDamages,
  DashboardOpenDamagesToFix
} from '@/interfaces/Dashboard/Dashboard'
import PriorityTag from '@/components/common/PriorityTag'
import toLocaleDate from '@/utils/date/toLocaleDate'
import style from './OpenDamages.module.css'

interface OpenDamagesProps {
  openDamages: DashboardOpenDamages | null
  loading: boolean
  error: string | null
  height: number
}

const OpenDamages: React.FC<OpenDamagesProps> = ({
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
      <div className={style.contentContainer}>
        {openDamages.list_to_fix.length === 0 ? (
          <Empty
            image={Empty.PRESENTED_IMAGE_SIMPLE}
            description={t('widgets.reminders.no_reminders')}
            className={style.empty}
          />
        ) : (
          openDamages.list_to_fix.map((damage: DashboardOpenDamagesToFix) => (
            <div key={damage.id} className={style.damageItem}>
              <div className={style.damageInformationsContainer}>
                <PriorityTag priority={damage.priority} />
                <span className={style.dateText}>
                  {toLocaleDate(damage.created_at)}
                </span>
              </div>
              {!damage.read ? (
                <Badge
                  className={style.damageComment}
                  color="blue"
                  text={damage.comment}
                  style={{ fontWeight: 700 }}
                />
              ) : (
                <span className={style.damageCommentWithoutBadge}>{damage.comment}</span>
              )}
            </div>
          ))
        )}
      </div>
    </div>
  )
}

export default OpenDamages
