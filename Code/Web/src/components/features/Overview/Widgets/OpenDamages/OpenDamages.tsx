import React from 'react'
import { useTranslation } from 'react-i18next'
import { useNavigate } from 'react-router-dom'
import { LoadingOutlined } from '@ant-design/icons'
import { Badge, Empty, Tooltip } from 'antd'
import {
  DashboardOpenDamages,
  DashboardOpenDamagesToFix
} from '@/interfaces/Dashboard/Dashboard'
import PriorityTag from '@/components/common/PriorityTag'
import toLocaleDate from '@/utils/date/toLocaleDate'
import NavigationEnum from '@/enums/NavigationEnum'
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
  const navigate = useNavigate()
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
        {openDamages?.list_to_fix === null ||
        openDamages?.list_to_fix?.length === 0 ? (
          <div className={style.emptyContainer}>
            <Empty
              image={Empty.PRESENTED_IMAGE_SIMPLE}
              description={t('widgets.reminders.no_reminders')}
              className={style.empty}
            />
          </div>
        ) : (
          openDamages?.list_to_fix?.map((damage: DashboardOpenDamagesToFix) => (
            <div
              key={damage.id}
              className={style.damageItem}
              onClick={() => {
                navigate(
                  NavigationEnum.DAMAGE_DETAILS.replace(
                    ':id',
                    damage.property_id
                  ).replace(':damageId', damage.id)
                )
              }}
              onKeyDown={e => {
                if (e.key === 'Enter' || e.key === ' ') {
                  e.preventDefault()
                  navigate(
                    NavigationEnum.DAMAGE_DETAILS.replace(
                      ':id',
                      damage.property_id
                    ).replace(':damageId', damage.id)
                  )
                }
              }}
              role="button"
              tabIndex={0}
              aria-label={`${damage.created_at}: ${damage.comment}`}
            >
              <div className={style.damageInformationsContainer}>
                <Tooltip
                  title={`${damage.property_name || '-'} > ${damage.room_name || '-'}`}
                  placement="topLeft"
                >
                  <span className={style.damageInfosContainer}>
                    {damage.property_name || '-'} {'>'}{' '}
                    {damage.room_name || '-'}
                  </span>
                </Tooltip>
                <span className={style.dateText}>
                  {toLocaleDate(damage.created_at)}
                </span>
              </div>
              {!damage.read ? (
                <Badge
                  className={style.damageComment}
                  color="blue"
                  text={
                    <Tooltip title={damage.comment} placement="topLeft">
                      <span style={{ fontWeight: 700 }}>{damage.comment}</span>
                    </Tooltip>
                  }
                />
              ) : (
                <div className={style.damageCommentContainer}>
                  <Tooltip title={damage.comment} placement="topLeft">
                    <span className={style.damageCommentWithoutBadge}>
                      {damage.comment}
                    </span>
                  </Tooltip>
                  <PriorityTag priority={damage.priority} />
                </div>
              )}
            </div>
          ))
        )}
      </div>
    </div>
  )
}

export default OpenDamages
