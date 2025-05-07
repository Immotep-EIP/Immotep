import React from 'react'
import { useTranslation } from 'react-i18next'
import { Empty } from 'antd'
import { WidgetProps } from '@/interfaces/Widgets/Widgets.ts'
import style from './Reminders.module.css'

const Reminders: React.FC<WidgetProps> = ({ height }) => {
  const { t } = useTranslation()
  const rowHeight = 120
  const pixelHeight = height * rowHeight

  const reminders = [
    {
      id: 1,
      title: 'End of lease',
      description:
        'Your contract on property "Logement étudiant KM0" is about to expire. Think about the exit inventory.'
    },
    {
      id: 2,
      title: 'Empty property',
      description: 'Your property "The Stables" is empty.'
    },
    {
      id: 3,
      title: 'New message',
      description:
        'You have a new message from "John Doe" regarding your property "The Stables".'
    },
    {
      id: 4,
      title: 'New message',
      description:
        'You have a new message from "John Doe" regarding your property "The Stables".'
    }
  ]

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
            <div key={reminder.id} className={style.reminderItem}>
              <div className={style.reminderTexts}>
                <span className={style.titleText}>{reminder.title}</span>
                <span className={style.descriptionText}>
                  {reminder.description}
                </span>
              </div>
            </div>
          ))
        )}
      </div>
    </div>
  )
}

export default Reminders
