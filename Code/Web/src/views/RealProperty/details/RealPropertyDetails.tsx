import React, { useEffect, useState } from 'react'
import { useLocation } from 'react-router-dom'
import {
  Button,
  MenuProps,
  message,
  Modal,
  Tabs,
  TabsProps,
  Dropdown,
  Badge
} from 'antd'
import { MoreOutlined } from '@ant-design/icons'
import { useTranslation } from 'react-i18next'

import defaultHouse from '@/assets/images/DefaultHouse.jpg'

import InviteTenantModal from '@/components/DetailsPage/InviteTenantModal'
import { PropertyDetails } from '@/interfaces/Property/Property'
import returnIcon from '@/assets/icons/retour.svg'

import { PropertyIdProvider } from '@/context/propertyIdContext'
import GetPropertyPicture from '@/services/api/Owner/Properties/GetPropertyPicture'
import useImageCache from '@/hooks/Property/useImageCache'
import PageMeta from '@/components/PageMeta/PageMeta'
import useProperties from '@/hooks/Property/useProperties'
import ArchiveProperty from '@/services/api/Owner/Properties/ArchiveProperty'
import CancelTenantInvitation from '@/services/api/Owner/Properties/CancelTenantInvitation'
import useNavigation from '@/hooks/Navigation/useNavigation'
import PageTitle from '@/components/PageText/Title'
import SubtitledElement from '@/components/SubtitledElement/SubtitledElement'
import PropertyStatusEnum, { TenantStatusEnum } from '@/enums/PropertyEnum'
import EndLease from '@/services/api/Owner/Properties/Leases/EndLease'
import DocumentsTab from './tabs/1DocumentsTab'
import DamageTab from './tabs/3DamageTab'
import InventoryTab from './tabs/2InventoryTab'
import style from './RealPropertyDetails.module.css'
import RealPropertyUpdate from '../update/RealPropertyUpdate'

interface ChildrenComponentProps {
  t: (key: string) => string
}

const ChildrenComponent: React.FC<ChildrenComponentProps> = ({ t }) => {
  const items: TabsProps['items'] = [
    {
      key: '1',
      label: t('components.button.documents'),
      children: <DocumentsTab />
    },
    {
      key: '2',
      label: t('components.button.inventory'),
      children: <InventoryTab />
    },
    {
      key: '3',
      label: t('components.button.damage'),
      children: <DamageTab />
    }
  ]

  return (
    <div className={style.childrenContainer}>
      <Tabs style={{ width: '100%' }} defaultActiveKey="1" items={items} />
    </div>
  )
}

interface DetailsPartProps {
  propertyData: PropertyDetails | null
  showModal: () => void
  propertyId: string
  showModalUpdate: () => void
}

const DetailsPart: React.FC<DetailsPartProps> = ({
  propertyData,
  showModal,
  propertyId,
  showModalUpdate
}) => {
  const { t } = useTranslation()
  const { goToRealProperty } = useNavigation()
  const { data: picture, isLoading } = useImageCache(
    propertyData?.id || '',
    GetPropertyPicture
  )
  const { refreshPropertyDetails } = useProperties()

  const removeProperty = async () => {
    Modal.confirm({
      title: t('components.modal.archive_property.title'),
      content: t('components.modal.archive_property.description'),
      okText: t('components.button.confirm'),
      cancelText: t('components.button.cancel'),
      okButtonProps: { danger: true },
      onOk: async () => {
        if (!propertyData) {
          message.error('Property ID is missing.')
          return
        }
        try {
          await ArchiveProperty(propertyData.id)
          message.success(t('components.modal.archive_property.success'))
          goToRealProperty()
        } catch (error) {
          console.error('Error deleting property:', error)
          message.error(t('components.modal.archive_property.error'))
        }
      }
    })
  }

  const endContract = async () => {
    Modal.confirm({
      title: t('components.modal.end_contract.title'),
      content: t('components.modal.end_contract.description'),
      okText: t('components.button.confirm'),
      cancelText: t('components.button.cancel'),
      okButtonProps: { danger: true },
      onOk: async () => {
        try {
          if (!propertyData) {
            message.error('Property ID is missing.')
            return
          }
          await EndLease(propertyData?.id || '')
          await refreshPropertyDetails(propertyData.id)
          message.success(t('components.modal.end_contract.success'))
        } catch (error) {
          console.error('Error ending contract:', error)
          message.error(t('components.modal.end_contract.error'))
        }
      }
    })
  }

  const cancelInvitation = async () => {
    Modal.confirm({
      title: t('components.modal.cancel_invitation.title'),
      content: t('components.modal.cancel_invitation.description'),
      okText: t('components.button.confirm'),
      cancelText: t('components.button.cancel'),
      okButtonProps: { danger: true },
      onOk: async () => {
        try {
          if (!propertyData) {
            message.error('Property ID is missing.')
            return
          }
          await CancelTenantInvitation(propertyData.id)
          await refreshPropertyDetails(propertyData.id)
          message.success(t('components.modal.cancel_invitation.success'))
        } catch (error) {
          console.error('Error cancelling invitation:', error)
          message.error(t('components.modal.cancel_invitation.error'))
        }
      }
    })
  }

  const items: MenuProps['items'] = [
    {
      key: '1',
      label: t('components.button.add_tenant'),
      onClick: () => {
        showModal()
      },
      disabled:
        propertyData?.status === PropertyStatusEnum.UNAVAILABLE ||
        propertyData?.status === PropertyStatusEnum.INVITATION_SENT
    },
    {
      key: '2',
      label: t('components.button.end_contract'),
      onClick: () => {
        endContract()
      },
      danger: true,
      disabled: propertyData?.status !== PropertyStatusEnum.UNAVAILABLE
    },
    {
      key: '3',
      label: t('components.button.cancel_invitation'),
      onClick: () => {
        cancelInvitation()
      },
      danger: true,
      disabled: propertyData?.status !== PropertyStatusEnum.INVITATION_SENT
    },
    {
      key: '4',
      label: t('components.button.edit_property'),
      onClick: () => {
        showModalUpdate()
      }
      // disabled: true
    },
    {
      key: '5',
      label: t('components.button.archive_property'),
      danger: true,
      onClick: () => {
        removeProperty()
      }
    }
  ]

  return (
    <div className={style.mainContainer}>
      <div className={style.moreInfosContainer}>
        <div
          className={style.returnButtonContainer}
          onClick={() => window.history.back()}
          tabIndex={0}
          role="button"
          onKeyDown={e => {
            if (e.key === 'Enter') {
              window.history.back()
            }
          }}
        >
          <img src={returnIcon} alt="Return" className={style.returnIcon} />
          <PageTitle
            title={t('pages.real_property_details.title')}
            size="title"
            margin={false}
          />
        </div>
        <Dropdown menu={{ items }} trigger={['click']} placement="bottomRight">
          <Button
            type="text"
            icon={<MoreOutlined />}
            className={style.actionButton}
          />
        </Dropdown>
      </div>
      <div className={style.headerInformationContainer}>
        <div className={style.pictureContainer}>
          <Badge.Ribbon
            text={t(
              TenantStatusEnum[
                propertyData!.status as keyof typeof TenantStatusEnum
              ].text || ''
            )}
            color={
              TenantStatusEnum[
                propertyData!.status as keyof typeof TenantStatusEnum
              ].color || 'default'
            }
          >
            <img
              src={isLoading ? defaultHouse : picture || defaultHouse}
              alt="Property"
              className={style.propertyPicture}
            />
          </Badge.Ribbon>
        </div>
        <div className={style.informationsContainer}>
          <div className={style.details}>
            <SubtitledElement
              subtitleKey={t('pages.real_property_details.informations.name')}
            >
              <span className={style.detailsText}>{propertyData?.name}</span>
            </SubtitledElement>
          </div>
          <div className={style.details}>
            <SubtitledElement
              subtitleKey={t(
                'pages.real_property_details.informations.address'
              )}
            >
              <span className={style.detailsText}>
                {propertyData?.apartment_number
                  ? `N°${propertyData?.apartment_number} - ${propertyData?.address}, ${propertyData?.postal_code} ${propertyData?.city}`
                  : `${propertyData?.address}, ${propertyData?.postal_code} ${propertyData?.city}`}
              </span>
            </SubtitledElement>
          </div>
          <div className={style.details}>
            <SubtitledElement
              subtitleKey={t('pages.real_property_details.informations.tenant')}
            >
              <span className={style.detailsText}>
                {propertyData?.lease?.tenant_email
                  ? propertyData?.lease.tenant_email
                  : '-----------'}
              </span>
            </SubtitledElement>
          </div>
          <div className={style.details}>
            <SubtitledElement
              subtitleKey={t('pages.real_property_details.informations.dates')}
            >
              <span className={style.detailsText}>
                {propertyData?.lease?.start_date
                  ? `${new Date(propertyData.lease.start_date).toLocaleDateString('fr-FR', { day: 'numeric', month: 'long', year: 'numeric' })}`
                  : '...'}
                {' - '}
                {propertyData?.lease?.end_date
                  ? `${new Date(propertyData.lease.end_date).toLocaleDateString('fr-FR', { day: 'numeric', month: 'long', year: 'numeric' })}`
                  : '...'}
              </span>
            </SubtitledElement>
          </div>
          <div className={style.details}>
            <SubtitledElement
              subtitleKey={t('pages.real_property_details.informations.area')}
            >
              <span className={style.detailsText}>
                {propertyData?.area_sqm} m²
              </span>
            </SubtitledElement>
          </div>
          <div className={style.details}>
            <SubtitledElement
              subtitleKey={t('pages.real_property_details.informations.rental')}
            >
              <span className={style.detailsText}>
                {propertyData?.rental_price_per_month} €
              </span>
            </SubtitledElement>
          </div>
          <div className={style.details}>
            <SubtitledElement
              subtitleKey={t(
                'pages.real_property_details.informations.deposit'
              )}
            >
              <span className={style.detailsText}>
                {propertyData?.deposit_price} €
              </span>
            </SubtitledElement>
          </div>
        </div>
      </div>
      <PropertyIdProvider id={propertyId}>
        <ChildrenComponent t={t} />
      </PropertyIdProvider>
    </div>
  )
}

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
        <DetailsPart
          propertyData={propertyData}
          showModal={showModal}
          propertyId={id}
          showModalUpdate={showModalUpdate}
        />
        {propertyData && (
          <>
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
