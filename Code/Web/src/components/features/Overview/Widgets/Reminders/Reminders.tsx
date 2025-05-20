import React from 'react'
import { useTranslation } from 'react-i18next'
import { useNavigate } from 'react-router-dom'
import { Empty } from 'antd'
import { LoadingOutlined } from '@ant-design/icons'
import { DashboardReminders } from '@/interfaces/Dashboard/Dashboard'
import PriorityTag from '@/components/common/PriorityTag'
import style from './Reminders.module.css'

interface RemindersProps {
  reminders: DashboardReminders[] | null
  loading: boolean
  error: string | null
  height: number
}

const Reminders: React.FC<RemindersProps> = ({
  reminders,
  loading,
  error,
  height
}: RemindersProps) => {
  const { t } = useTranslation()
  const navigate = useNavigate()
  const rowHeight = 120
  const pixelHeight = height * rowHeight

  const navigateTo = (link: string) => {
    const cleanLink = link.replace(/^(https?:\/\/)?([^/]+)\//, '/')

    const formattedLink = cleanLink.startsWith('/')
      ? cleanLink
      : `/${cleanLink}`

    navigate(formattedLink)
  }

  if (loading || reminders === null) {
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
        {reminders.length === 0 ? (
          <Empty
            image={Empty.PRESENTED_IMAGE_SIMPLE}
            description={t('widgets.reminders.no_reminders')}
            className={style.empty}
          />
        ) : (
          reminders.map(reminder => (
            <div
              key={reminder.id}
              className={style.reminderItem}
              onClick={() => navigateTo(reminder.link)}
              onKeyDown={e => {
                if (e.key === 'Enter' || e.key === ' ') {
                  e.preventDefault()
                  navigateTo(reminder.link)
                }
              }}
              role="button"
              tabIndex={0}
              aria-label={`${reminder.title}: ${reminder.advice}`}
            >
              <div className={style.reminderTexts}>
                <span className={style.titleText}>{reminder.title}</span>
                <div className={style.reminderheader}>
                  <PriorityTag priority={reminder.priority} />
                  <span className={style.descriptionText}>
                    {reminder.advice}
                  </span>
                </div>
              </div>
            </div>
          ))
        )}
      </div>
    </div>
  )
}

export default Reminders
