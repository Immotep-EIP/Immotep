import React, { useEffect, useState } from 'react'
import { useLocation } from 'react-router-dom'
import { useTranslation } from 'react-i18next'
import { Spin } from 'antd'

import useProperties from '@/hooks/Property/useProperties'
import RealPropertyUpdate from '@/views/RealProperty/update/RealPropertyUpdate'
import InviteTenantModal from '@/components/features/RealProperty/details/DetailsPart/InviteTenantModal'
import PageMeta from '@/components/ui/PageMeta/PageMeta'
import DetailsPart from '@/components/features/RealProperty/details/DetailsPart/DetailsPart'
import { PropertyProvider } from '@/context/propertyContext'

import style from './RealPropertyDetails.module.css'

const RealPropertyDetails: React.FC = () => {
  const { t } = useTranslation()
  const location = useLocation()
  let { id } = location.state || {}
  if (!id) {
    const pathParts = location.pathname.split('/')
    const lastPart = pathParts[pathParts.length - 1]
    const idFromUrl = lastPart.split('?')[0]
    if (!idFromUrl) {
      throw new Error('Property ID not found in location state or URL')
    }
    id = idFromUrl
  }
  const [isModalOpen, setIsModalOpen] = useState(false)
  const showModal = () => setIsModalOpen(true)

  const [isModalUpdateOpen, setIsModalUpdateOpen] = useState(false)
  const showModalUpdate = () => setIsModalUpdateOpen(true)
  const [isPropertyUpdated, setIsPropertyUpdated] = useState(false)

  const [refreshKey, setRefreshKey] = useState(0)

  const {
    propertyDetails: propertyData,
    refreshPropertyDetails,
    loading,
    error
  } = useProperties(id)

  const handleCancel = (invitationSent: boolean) => {
    setIsModalOpen(false)
    if (invitationSent) {
      refreshPropertyDetails(id)
      setRefreshKey(prev => prev + 1)
    }
  }

  useEffect(() => {
    if (isPropertyUpdated) {
      refreshPropertyDetails(id)
      setRefreshKey(prev => prev + 1)
      setIsPropertyUpdated(false)
    }
  }, [isPropertyUpdated, id, refreshPropertyDetails])

  useEffect(() => {
    if (propertyData) {
      setRefreshKey(prev => prev + 1)
    }
  }, [propertyData])

  if (loading) {
    return (
      <section
        role="status"
        aria-live="polite"
        aria-labelledby="property-loading-title"
      >
        <h1 id="property-loading-title" className="sr-only">
          {t('pages.real_property_details.loading_title')}
        </h1>
        <Spin size="large" />
        <div>{t('pages.real_property_details.loading_title')}</div>
      </section>
    )
  }

  if (error) {
    return (
      <section
        role="alert"
        aria-live="assertive"
        aria-labelledby="property-error-title"
      >
        <h1 id="property-error-title" className="sr-only">
          {t('pages.real_property_details.error_title')}
        </h1>
        <div>
          {t('pages.real_property_details.error_title', { message: error })}
        </div>
      </section>
    )
  }

  return (
    <>
      <PageMeta
        title={t('pages.real_property_details.document_title')}
        description={t('pages.real_property_details.document_description')}
        keywords="real property details, Property info, Keyz"
      />

      <main
        className={style.pageContainer}
        aria-labelledby="property-details-title"
      >
        <h1 id="property-details-title" className="sr-only">
          {t('pages.real_property_details.page_title')}
        </h1>
        {propertyData && (
          <PropertyProvider property={propertyData}>
            <DetailsPart
              key={refreshKey}
              propertyData={propertyData}
              showModal={showModal}
              propertyId={id}
              showModalUpdate={showModalUpdate}
              refreshPropertyDetails={refreshPropertyDetails}
            />
            <InviteTenantModal
              isOpen={isModalOpen}
              onClose={handleCancel}
              propertyId={id}
            />
            <RealPropertyUpdate
              isModalUpdateOpen={isModalUpdateOpen}
              setIsModalUpdateOpen={setIsModalUpdateOpen}
              propertyData={propertyData}
              setIsPropertyUpdated={setIsPropertyUpdated}
            />
          </PropertyProvider>
        )}
      </main>
    </>
  )
}

export default RealPropertyDetails
