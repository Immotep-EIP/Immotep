import React, { useEffect, useState } from 'react'
import { Button, Empty, Typography, Switch } from 'antd'
import { useTranslation } from 'react-i18next'

import useProperties from '@/hooks/Property/useProperties'
import PageTitle from '@/components/ui/PageText/Title'
import CardPropertyLoader from '@/components/ui/Loader/CardPropertyLoader'
import PageMeta from '@/components/ui/PageMeta/PageMeta'
import CardComponent from '@/components/features/RealProperty/PropertyCard'
import PropertyFilterCard from '@/components/features/RealProperty/PropertyFilterCard'
import RealPropertyCreate from './create/RealPropertyCreate'

import style from './RealProperty.module.css'

const RealPropertyPage: React.FC = () => {
  const { t } = useTranslation()
  const [showArchived, setShowArchived] = useState(false)
  const [filters, setFilters] = useState({
    searchQuery: '',
    surfaceRange: 'all',
    status: 'all'
  })
  const { properties, loading, error, refreshProperties } = useProperties(
    null,
    showArchived
  )
  const [showModalCreate, setShowModalCreate] = useState(false)
  const [isPropertyCreated, setIsPropertyCreated] = useState(false)

  useEffect(() => {
    if (isPropertyCreated) {
      refreshProperties()
      setIsPropertyCreated(false)
    }
  }, [isPropertyCreated, refreshProperties])

  const surfaceRangeOptions = [
    { value: 'all', label: t('components.select.surface.all') },
    { value: '0-50', label: '0-50m²' },
    { value: '51-100', label: '51-100m²' },
    { value: '101-150', label: '101-150m²' },
    { value: '151-200', label: '151-200m²' },
    { value: '201+', label: '201m²+' }
  ]

  const statusOptions = [
    { value: 'all', label: t('components.select.status.all') },
    { value: 'available', label: t('components.select.status.available') },
    {
      value: 'unavailable',
      label: t('components.select.status.unavailable')
    },
    {
      value: 'invitation_sent',
      label: t('components.select.status.invitation_sent')
    }
  ]

  const filteredProperties = properties.filter(property => {
    const searchLower = filters.searchQuery.toLowerCase()
    const matchesSearch =
      property.name.toLowerCase().includes(searchLower) ||
      (property.country &&
        property.country.toLowerCase().includes(searchLower)) ||
      (property.city && property.city.toLowerCase().includes(searchLower))

    const matchesSurface =
      filters.surfaceRange === 'all' ||
      (filters.surfaceRange === '201+'
        ? property.area_sqm >= 201
        : property.area_sqm >=
            parseInt(filters.surfaceRange.split('-')[0], 10) &&
          property.area_sqm <= parseInt(filters.surfaceRange.split('-')[1], 10))

    const matchesStatus =
      filters.status === 'all' || property.status === filters.status

    return matchesSearch && matchesSurface && matchesStatus
  })

  if (error) {
    return <p>{t('pages.real_property.error.error_fetching_data')}</p>
  }

  return (
    <>
      <PageMeta
        title={t('pages.real_property.document_title')}
        description={t('pages.real_property.document_description')}
        keywords="real property, Property info, Keyz"
      />
      <div className={style.pageContainer}>
        <div className={style.pageHeader}>
          <PageTitle title={t('pages.real_property.title')} size="title" />
          <div className={style.headerActions}>
            <div className={style.archiveFilter}>
              <Switch
                checked={showArchived}
                onChange={setShowArchived}
                checkedChildren={t('components.switch.show_archived')}
                unCheckedChildren={t('components.switch.show_active')}
              />
            </div>
            <Button type="primary" onClick={() => setShowModalCreate(true)}>
              {t('components.button.add_real_property')}
            </Button>
          </div>
        </div>

        <PropertyFilterCard
          filters={filters}
          setFilters={setFilters}
          surfaceRangeOptions={surfaceRangeOptions}
          statusOptions={statusOptions}
        />

        <div className={style.cardsContainer}>
          {loading && <CardPropertyLoader cards={12} />}
          {!loading && filteredProperties.length === 0 && (
            <div className={style.emptyContainer}>
              <Empty
                image="https://gw.alipayobjects.com/zos/antfincdn/ZHrcdLPrvN/empty.svg"
                description={
                  <Typography.Text>
                    {t('components.messages.no_properties')}
                  </Typography.Text>
                }
              />
            </div>
          )}
          {filteredProperties.map(realProperty => (
            <CardComponent
              key={realProperty.id}
              realProperty={realProperty}
              t={t}
            />
          ))}
        </div>
        <RealPropertyCreate
          showModalCreate={showModalCreate}
          setShowModalCreate={setShowModalCreate}
          setIsPropertyCreated={setIsPropertyCreated}
        />
      </div>
    </>
  )
}

export default RealPropertyPage
