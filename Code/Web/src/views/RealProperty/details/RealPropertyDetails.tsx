import React, { useState } from 'react'
import { useLocation } from 'react-router-dom'
import {
  Button,
} from 'antd'
import { useTranslation } from 'react-i18next'

import InviteTenantModal from '@/components/DetailsPage/InviteTenantModal'
import style from './RealPropertyDetails.module.css'

const RealPropertyDetails: React.FC = () => {
  const { t } = useTranslation()
  const location = useLocation()
  const { id } = location.state || {}

  const [isModalOpen, setIsModalOpen] = useState(false)

  const showModal = () => setIsModalOpen(true)
  const handleCancel = () => setIsModalOpen(false)

  return (
    <div className={style.pageContainer}>
      <span>Details for Real Property ID: {id}</span>
      <Button type="primary" onClick={showModal}>
        {t('components.button.addTenant')}
      </Button>

      <InviteTenantModal isOpen={isModalOpen} onClose={handleCancel} propertyId={id} />
    </div>
  )
}

export default RealPropertyDetails
