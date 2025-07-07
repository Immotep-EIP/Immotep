import React, { useEffect, useState } from 'react'
import { Switch } from 'antd'
import { useTranslation } from 'react-i18next'
import { useSearchParams } from 'react-router-dom'

import useProperties from '@/hooks/Property/useProperties'
import { Button, Empty } from '@/components/common'
import PageTitle from '@/components/ui/PageText/Title'
import CardPropertyLoader from '@/components/ui/Loader/CardPropertyLoader'
import PageMeta from '@/components/ui/PageMeta/PageMeta'
import CardComponent from '@/components/features/RealProperty/PropertyCard'
import PropertyFilterCard from '@/components/features/RealProperty/PropertyFilterCard'
import RealPropertyCreate from './create/RealPropertyCreate'

import style from './RealProperty.module.css'

const RealPropertyPage: React.FC = () => {
  const { t } = useTranslation()

  const [searchParams, setSearchParams] = useSearchParams()
  const archiveParam = searchParams.get('archive')
  const [showArchived, setShowArchived] = useState(archiveParam === 'true')
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

  const handleArchiveToggle = (checked: boolean) => {
    setShowArchived(checked)
    if (checked) {
      setSearchParams({ archive: 'true' })
    } else {
      setSearchParams({})
    }
  }

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
    return (
      <section
        role="alert"
        aria-live="assertive"
        aria-labelledby="properties-error-title"
      >
        <h1 id="properties-error-title" className="sr-only">
          {t('pages.real_property.error_title')}
        </h1>
        <p>{t('pages.real_property.error.error_fetching_data')}</p>
      </section>
    )
  }

  return (
    <>
      <PageMeta
        title={t('pages.real_property.document_title')}
        description={t('pages.real_property.document_description')}
        keywords="real property, Property info, Keyz"
      />
      <main
        className={style.pageContainer}
        aria-labelledby="properties-page-title"
      >
        <header className={style.pageHeader}>
          <PageTitle
            title={t('pages.real_property.title')}
            size="title"
            id="properties-page-title"
          />
          <div
            className={style.headerActions}
            role="toolbar"
            aria-label={t('pages.real_property.toolbar_aria')}
          >
            <div className={style.archiveFilter}>
              <Switch
                checked={showArchived}
                onChange={handleArchiveToggle}
                checkedChildren={t('components.switch.show_archived')}
                unCheckedChildren={t('components.switch.show_active')}
                aria-label={t('pages.real_property.archive_toggle_aria')}
                aria-describedby="archive-help"
              />
              <div id="archive-help" className="sr-only">
                {t('pages.real_property.archive_toggle_help')}
              </div>
            </div>
            <Button
              onClick={() => setShowModalCreate(true)}
              aria-label={t('pages.real_property.add_property_aria')}
            >
              {t('components.button.add_real_property')}
            </Button>
          </div>
        </header>

        <PropertyFilterCard
          filters={filters}
          setFilters={setFilters}
          surfaceRangeOptions={surfaceRangeOptions}
          statusOptions={statusOptions}
        />

        <section
          className={style.cardsContainer}
          aria-labelledby="properties-list-title"
        >
          <h2 id="properties-list-title" className="sr-only">
            {t('pages.real_property.properties_list_title')}
          </h2>
          {loading && <CardPropertyLoader cards={12} />}
          {!loading && filteredProperties.length === 0 && (
            <div
              className={style.emptyContainer}
              role="status"
              aria-live="polite"
              aria-labelledby="properties-empty-title"
            >
              <h3 id="properties-empty-title" className="sr-only">
                {t('pages.real_property.empty_title')}
              </h3>
              <Empty description={t('components.messages.no_properties')} />
            </div>
          )}
          {filteredProperties.map(realProperty => (
            <CardComponent
              key={realProperty.id}
              realProperty={realProperty}
              t={t}
            />
          ))}
        </section>
        <RealPropertyCreate
          showModalCreate={showModalCreate}
          setShowModalCreate={setShowModalCreate}
          setIsPropertyCreated={setIsPropertyCreated}
        />
      </main>
    </>
  )
}

export default RealPropertyPage
