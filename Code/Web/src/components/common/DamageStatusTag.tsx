import React from 'react'
import { Tag } from 'antd'
import { useTranslation } from 'react-i18next'

const DamageStatusTag: React.FC<{ status: string }> = ({ status }) => {
  const { t } = useTranslation()
  const statusLower = status.toLowerCase()

  const getTagColor = (status: string): string => {
    if (!status) {
      return 'default'
    }
    const lowerStatus = status.toLowerCase()

    switch (lowerStatus) {
      case 'pending':
        return 'red'
      case 'planned':
        return 'orange'
      case 'awaiting_owner_confirmation':
        return 'blue'
      case 'awaiting_tenant_confirmation':
        return 'purple'
      case 'fixed':
        return 'green'
      default:
        return 'gray'
    }
  }

  return (
    <Tag color={getTagColor(statusLower)}>
      {t(
        `pages.real_property_details.tabs.damage.status.${status.toLowerCase()}`
      )}
    </Tag>
  )
}

export default DamageStatusTag
