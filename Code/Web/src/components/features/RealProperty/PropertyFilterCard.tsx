import React from 'react'
import { useTranslation } from 'react-i18next'

import { Input, Select, Space, Card } from 'antd'
import { SearchOutlined } from '@ant-design/icons'

import { PropertyFilterCardProps } from '@/interfaces/Property/Property'

import style from './PropertyFilterCard.module.css'

const { Search } = Input

const PropertyFilterCard: React.FC<PropertyFilterCardProps> = ({
  filters,
  setFilters,
  surfaceRangeOptions,
  statusOptions
}) => {
  const { t } = useTranslation()
  return (
    <Card className={style.filtersCard}>
      <Space direction="vertical" size="middle" style={{ width: '100%' }}>
        <Search
          placeholder={t('components.input.search_property.placeholder')}
          allowClear
          enterButton={<SearchOutlined />}
          value={filters.searchQuery}
          onChange={e =>
            setFilters(f => ({ ...f, searchQuery: e.target.value }))
          }
          style={{ width: '100%' }}
        />
        <Space wrap>
          <Select
            style={{ width: 200 }}
            value={filters.surfaceRange}
            onChange={value => setFilters(f => ({ ...f, surfaceRange: value }))}
            options={surfaceRangeOptions}
            placeholder={t('components.select.surface_range.placeholder')}
          />
          <Select
            style={{ width: 200 }}
            value={filters.status}
            onChange={value => setFilters(f => ({ ...f, status: value }))}
            options={statusOptions}
            placeholder={t('components.select.status.placeholder')}
          />
        </Space>
      </Space>
    </Card>
  )
}

export default PropertyFilterCard
