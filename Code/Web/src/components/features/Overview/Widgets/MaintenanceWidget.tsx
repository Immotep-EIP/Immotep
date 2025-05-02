import React, { useState } from 'react'
import { List, Button, Tag } from 'antd'
import { CheckOutlined } from '@ant-design/icons'

import { useTranslation } from 'react-i18next'
import { MaintenanceTask, WidgetProps } from '@/interfaces/Widgets/Widgets.ts'
import style from './MaintenanceWidget.module.css'

const MaintenanceWidget: React.FC<WidgetProps> = ({ height }) => {
  const { t } = useTranslation()
  const rowHeight = 70
  const pixelHeight = height * rowHeight

  const [tasks, setTasks] = useState<MaintenanceTask[]>([
    {
      id: 1,
      description: 'Réparation de la chaudière',
      priority: 'high',
      completed: false
    },
    {
      id: 2,
      description: 'Entretien des jardins',
      priority: 'medium',
      completed: false
    },
    {
      id: 3,
      description: "Révision de l'ascenseur",
      priority: 'high',
      completed: false
    },
    {
      id: 4,
      description: 'Peinture des couloirs',
      priority: 'low',
      completed: false
    }
  ])

  const markAsCompleted = (id: number) => {
    const updatedTasks = tasks.map(task =>
      task.id === id ? { ...task, completed: true } : task
    )
    setTasks(updatedTasks)
  }

  const renderPriorityTag = (priority: 'high' | 'medium' | 'low') => {
    const colors = {
      high: 'red',
      medium: 'orange',
      low: 'green'
    }
    return <Tag color={colors[priority]} />
  }

  return (
    <div
      className={style.maintenanceWidgetContainer}
      style={{ height: `${pixelHeight}px` }}
    >
      <h4 className={style.maintenanceWidgetTitle}>
        {t('widgets.maintenance.title')}
      </h4>
      <div className={style.maintenanceWidgetScrollList}>
        <List
          dataSource={tasks}
          aria-label={t('listDamage')}
          renderItem={task => (
            <List.Item className={style.maintenanceWidgetListItem}>
              <div className={style.maintenanceWidgetTask}>
                {renderPriorityTag(task.priority)}
                <span
                  className={
                    task.completed
                      ? style.maintenanceWidgetTaskCompleted
                      : style.maintenanceWidgetTaskPending
                  }
                >
                  {task.description}
                </span>
              </div>
              {!task.completed && (
                <Button
                  type="primary"
                  size="small"
                  icon={<CheckOutlined />}
                  onClick={() => markAsCompleted(task.id)}
                >
                  {t('widgets.maintenance.complete_button')}
                </Button>
              )}
            </List.Item>
          )}
        />
      </div>
    </div>
  )
}

export default MaintenanceWidget
