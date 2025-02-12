import React, { useEffect, useState } from 'react'
import { useLocation } from 'react-router-dom'
import { Button, message, Modal, Tabs, TabsProps, Tag } from 'antd'
import { useTranslation } from 'react-i18next'

import defaultHouse from '@/assets/images/DefaultHouse.jpg'
import appartmentIcon from '@/assets/icons/appartement.png'
import locationIcon from '@/assets/icons/location.png'
import tenantIcon from '@/assets/icons/tenant.png'
import dateIcon from '@/assets/icons/date.png'

import InviteTenantModal from '@/components/DetailsPage/InviteTenantModal'
import GetPropertyDetails from '@/services/api/Owner/Properties/GetPropertyDetails'
import { PropertyDetails } from '@/interfaces/Property/Property'
import returnIcon from '@/assets/icons/retour.png'

import { PropertyIdProvider } from '@/context/propertyIdContext'
import GetPropertyPicture from '@/services/api/Owner/Properties/GetPropertyPicture'
import StopCurrentContract from '@/services/api/Owner/Properties/StopCurrentContract'
import useImageCache from '@/hooks/useEffect/useImageCache'
import AboutTab from './tabs/1AboutTab'
import DamageTab from './tabs/2DamageTab'
import InventoryTab from './tabs/3InventoryTab'
import DocumentsTab from './tabs/4DocumentsTab'
import style from './RealPropertyDetails.module.css'

const HeaderPart: React.FC<{ propertyData: PropertyDetails | null }> = ({
  propertyData
}) => {
  const { t } = useTranslation()

  const { data: picture, isLoading } = useImageCache(
    propertyData?.id || '',
    GetPropertyPicture
  )

  if (!propertyData) {
    return null
  }

  return (
    <div className={style.headerPartContainer}>
      <div className={style.imageContainer}>
        <img
          src={isLoading ? defaultHouse : picture || defaultHouse}
          alt="Property"
          className={style.image}
        />
      </div>
      <div className={style.detailsContainer}>
        <div className={style.details}>
          <img
            src={appartmentIcon}
            alt="Appartment"
            className={style.detailsIcon}
          />
          <span className={style.detailsText}>{propertyData.name}</span>
        </div>
        <div className={style.details}>
          <img
            src={locationIcon}
            alt="Location"
            className={style.detailsIcon}
          />
          <span className={style.detailsText}>
            {`${propertyData.address}, ${propertyData.postal_code} ${propertyData.city}`}
          </span>
        </div>
        <div className={style.details}>
          <img src={tenantIcon} alt="Tenant" className={style.detailsIcon} />
          <span className={style.detailsText}>
            {propertyData.tenant ? propertyData.tenant : '-----------'}
          </span>
        </div>
        <div className={style.details}>
          <img src={dateIcon} alt="Date" className={style.detailsIcon} />
          <span className={style.detailsText}>
            {propertyData.start_date
              ? `${new Date(propertyData.start_date).toLocaleDateString('fr-FR', { day: 'numeric', month: 'long', year: 'numeric' })}`
              : '...'}
            {' - '}
            {propertyData.end_date
              ? `${new Date(propertyData.end_date).toLocaleDateString('fr-FR', { day: 'numeric', month: 'long', year: 'numeric' })}`
              : '...'}
          </span>
        </div>
      </div>

      <div className={style.moreInfosContainer}>
        <Tag color={propertyData.nb_damage > 0 ? 'red' : 'green'}>
          {propertyData.nb_damage || 0}{' '}
          {t('pages.real_property.damage.waiting')}
        </Tag>
        <Tag color={propertyData.status === 'available' ? 'green' : 'red'}>
          {propertyData.status === 'available'
            ? t('pages.real_property.status.available')
            : t('pages.real_property.status.unavailable')}
        </Tag>
      </div>
    </div>
  )
}

interface ChildrenComponentProps {
  t: (key: string) => string
}

const ChildrenComponent: React.FC<ChildrenComponentProps> = ({ t }) => {
  // const onChange = (key: string) => {
  //   console.log(key)
  // }

  const items: TabsProps['items'] = [
    {
      key: '1',
      label: t('components.button.about'),
      children: <AboutTab />
    },
    {
      key: '2',
      label: t('components.button.damage'),
      children: <DamageTab />
    },
    {
      key: '3',
      label: t('components.button.inventory'),
      children: <InventoryTab />
    },
    {
      key: '4',
      label: t('components.button.documents'),
      children: <DocumentsTab />
    }
  ]

  return (
    <div className={style.childrenContainer}>
      <Tabs style={{ width: '100%' }} defaultActiveKey="1" items={items} />
    </div>
  )
}

const RealPropertyDetails: React.FC = () => {
  const { t } = useTranslation()
  const location = useLocation()
  const { id } = location.state || {}
  const [propertyData, setPropertyData] = useState<PropertyDetails | null>(null)

  const [isModalOpen, setIsModalOpen] = useState(false)
  const [modalRemoveTenant, setModalRemoveTenant] = useState(false)

  const fetchData = async () => {
    const req = await GetPropertyDetails(id)
    if (req) {
      setPropertyData(req)
    } else {
      message.error(t('pages.real_property_details.error_fetching_data'))
    }
  }

  useEffect(() => {
    fetchData()
  }, [id])

  const showModal = () => setIsModalOpen(true)
  const handleCancel = () => setIsModalOpen(false)

  const showModalRemoveTenant = () => setModalRemoveTenant(true)
  const handleCancelRemoveTenant = () => setModalRemoveTenant(false)
  const handleRemoveTenant = async () => {
    await StopCurrentContract(id)
    setModalRemoveTenant(false)
    message.success(t('components.modal.end_contract.success'))
    await fetchData()
  }

  return (
    <div className={style.pageContainer}>
      <div className={style.buttonsContainer}>
        <div
          className={style.returnButtonContainer}
          onClick={() => window.history.back()}
          tabIndex={0}
          onKeyDown={e => {
            if (e.key === 'Enter') {
              window.history.back()
            }
          }}
          role="button"
        >
          <img src={returnIcon} alt="Return" className={style.returnIcon} />
          <span className={style.returnText}>
            {t('components.button.return')}
          </span>
        </div>

        <div className={style.actionButtonsContainer}>
          <Button
            type="primary"
            onClick={showModal}
            disabled={propertyData?.status !== 'available'}
          >
            {t('components.button.add_tenant')}
          </Button>
          <Button
            type="primary"
            danger
            disabled={propertyData?.status !== 'unavailable'}
            onClick={showModalRemoveTenant}
          >
            {t('components.button.end_contract')}
          </Button>
        </div>
      </div>

      <Modal
        title={t('components.modal.end_contract.title')}
        open={modalRemoveTenant}
        onCancel={handleCancelRemoveTenant}
        footer={[
          <Button key="cancel" onClick={handleCancelRemoveTenant}>
            {t('components.button.cancel')}
          </Button>,
          <Button key="ok" type="primary" onClick={handleRemoveTenant}>
            {t('components.button.confirm')}
          </Button>
        ]}
      >
        {t('components.modal.end_contract.description')}
      </Modal>

      <InviteTenantModal
        isOpen={isModalOpen}
        onClose={handleCancel}
        propertyId={id}
      />

      <HeaderPart propertyData={propertyData} />
      <PropertyIdProvider id={id}>
        <ChildrenComponent t={t} />
      </PropertyIdProvider>
    </div>
  )
}

export default RealPropertyDetails
