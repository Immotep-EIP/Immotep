import React from 'react'
import { Tag } from 'antd'
import { useTranslation } from 'react-i18next'

interface StatusTagProps {
  value: string
  colorMap: Record<string, string>
  i18nPrefix?: string
  defaultColor?: string
}

const StatusTag: React.FC<StatusTagProps> = ({
  value,
  colorMap,
  i18nPrefix,
  defaultColor = 'default'
}) => {
  const { t } = useTranslation()
  const lowerValue = value?.toLowerCase() || ''
  const tagColor = colorMap?.[lowerValue] || defaultColor

  let label = value
  if (i18nPrefix) {
    const translation = t(`${i18nPrefix}.${lowerValue}`)

    if (!translation.startsWith(i18nPrefix)) {
      label = translation
    }
  }

  return <Tag color={tagColor}>{label}</Tag>
}

export default StatusTag
