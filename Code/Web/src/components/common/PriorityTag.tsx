import React from 'react'
import { Tag } from 'antd'
import { useTranslation } from 'react-i18next'

const PriorityTag: React.FC<{ priority: string }> = ({ priority }) => {
  const { t } = useTranslation()
  const priorityLower = priority.toLowerCase()

  const getTagColor = (priority: string): string => {
    if (!priority) {
      return 'default'
    }
    const lowerPriority = priority.toLowerCase()

    switch (lowerPriority) {
      case 'urgent':
        return 'red'
      case 'high':
        return 'red'
      case 'medium':
        return 'yellow'
      case 'low':
        return 'green'
      default:
        return 'green'
    }
  }

  return (
    <Tag color={getTagColor(priorityLower)}>
      {t(
        `pages.real_property_details.tabs.damage.priority.${priority.toLowerCase()}`
      )}
    </Tag>
  )
}

export default PriorityTag
