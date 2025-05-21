import React from 'react'
import { useTranslation } from 'react-i18next'
import { useNavigate } from 'react-router-dom'

import { Empty, Tooltip } from 'antd'
import { LoadingOutlined } from '@ant-design/icons'

import StatusTag from '@/components/common/Tag/StatusTag'

import { DashboardReminders } from '@/interfaces/Dashboard/Dashboard'

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
              onClick={() => navigate(reminder.link)}
              onKeyDown={e => {
                if (e.key === 'Enter' || e.key === ' ') {
                  e.preventDefault()
                  navigate(reminder.link)
                }
              }}
              role="button"
              tabIndex={0}
              aria-label={`${reminder.title}: ${reminder.advice}`}
            >
              <div className={style.reminderTexts}>
                <Tooltip title={reminder.title} placement="topLeft">
                  <span className={style.titleText}>{reminder.title}</span>
                </Tooltip>
                <div className={style.reminderheader}>
                  <StatusTag
                    value={reminder.priority}
                    colorMap={{
                      urgent: 'red',
                      high: 'red',
                      medium: 'yellow',
                      low: 'green'
                    }}
                    i18nPrefix="pages.real_property_details.tabs.damage.priority"
                    defaultColor="gray"
                  />
                  <Tooltip title={reminder.advice} placement="topLeft">
                    <span className={style.descriptionText}>
                      {reminder.advice}
                    </span>
                  </Tooltip>
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
