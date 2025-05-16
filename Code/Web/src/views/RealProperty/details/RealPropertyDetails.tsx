import React, { useEffect, useState } from 'react'
import { useLocation } from 'react-router-dom'
import { useTranslation } from 'react-i18next'

import InviteTenantModal from '@/components/features/RealProperty/details/DetailsPart/InviteTenantModal'
import PageMeta from '@/components/ui/PageMeta/PageMeta'
import useProperties from '@/hooks/Property/useProperties'
import DetailsPart from '@/components/features/RealProperty/details/DetailsPart/DetailsPart'
import RealPropertyUpdate from '../update/RealPropertyUpdate'
import style from './RealPropertyDetails.module.css'

const RealPropertyDetails: React.FC = () => {
  const { t } = useTranslation()
  const location = useLocation()
  const { id } = location.state || {}
  const [isModalOpen, setIsModalOpen] = useState(false)
  const showModal = () => setIsModalOpen(true)

  const [isModalUpdateOpen, setIsModalUpdateOpen] = useState(false)
  const showModalUpdate = () => setIsModalUpdateOpen(true)
  const [isPropertyUpdated, setIsPropertyUpdated] = useState(false)

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
    }
  }

  useEffect(() => {
    if (isPropertyUpdated) {
      refreshPropertyDetails(id)
      setIsPropertyUpdated(false)
    }
  }, [isPropertyUpdated, id, refreshPropertyDetails])

  if (loading) {
    return <div>{t('components.loading')}</div>
  }

  if (error) {
    return <div>{t('components.error', { message: error })}</div>
  }

  return (
    <>
      <PageMeta
        title={t('pages.real_property_details.document_title')}
        description={t('pages.real_property_details.document_description')}
        keywords="real property details, Property info, Keyz"
      />

      <div className={style.pageContainer}>
        {propertyData && (
          <>
            <DetailsPart
              propertyData={propertyData}
              showModal={showModal}
              propertyId={id}
              showModalUpdate={showModalUpdate}
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
          </>
        )}
      </div>
    </>
  )
}

export default RealPropertyDetails
