import React from 'react'
import { useTranslation } from 'react-i18next'

import { Input, Select, Space } from 'antd'
import { SearchOutlined } from '@ant-design/icons'

import { Card } from '@/components/common'

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
    <section aria-labelledby="property-filters-title">
      <h2 id="property-filters-title" className="sr-only">
        {t('components.property.filters.title')}
      </h2>

      <Card className={style.filtersCard}>
        <Space direction="vertical" size="middle" style={{ width: '100%' }}>
          <div role="search">
            <Search
              placeholder={t('components.input.search_property.placeholder')}
              allowClear
              enterButton={<SearchOutlined aria-hidden="true" />}
              value={filters.searchQuery}
              onChange={e =>
                setFilters(f => ({ ...f, searchQuery: e.target.value }))
              }
              style={{ width: '100%' }}
              aria-label={t('components.property.filters.search')}
              id="property-search"
            />
          </div>

          <fieldset style={{ border: 'none', padding: 0, margin: 0 }}>
            <legend className="sr-only">
              {t('components.property.filters.filter_options')}
            </legend>

            <Space wrap role="group" aria-labelledby="filter-group-label">
              <span id="filter-group-label" className="sr-only">
                {t('components.property.filters.filter_controls')}
              </span>

              <Select
                style={{ width: 200 }}
                value={filters.surfaceRange}
                onChange={value =>
                  setFilters(f => ({ ...f, surfaceRange: value }))
                }
                options={surfaceRangeOptions}
                placeholder={t('components.select.surface.placeholder')}
                aria-label={t('components.property.filters.surface_filter')}
                id="surface-range-filter"
              />

              <Select
                style={{ width: 200 }}
                value={filters.status}
                onChange={value => setFilters(f => ({ ...f, status: value }))}
                options={statusOptions}
                placeholder={t('components.select.status.placeholder')}
                aria-label={t('components.property.filters.status_filter')}
                id="status-filter"
              />
            </Space>
          </fieldset>
        </Space>
      </Card>
    </section>
  )
}

export default PropertyFilterCard
