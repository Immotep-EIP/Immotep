import React, { useState, useEffect } from 'react'
import { useTranslation } from 'react-i18next'

import { Modal, message, Tabs } from 'antd'

import { PropertyIdProvider } from '@/context/propertyIdContext'
import useImageCache from '@/hooks/Image/useImageCache'
import useProperties from '@/hooks/Property/useProperties'
import useNavigation from '@/hooks/Navigation/useNavigation'
import GetPropertyPicture from '@/services/api/Owner/Properties/GetPropertyPicture'
import CancelTenantInvitation from '@/services/api/Owner/Properties/CancelTenantInvitation'
import EndLease from '@/services/api/Owner/Properties/Leases/EndLease'
import ArchiveProperty from '@/services/api/Owner/Properties/ArchiveProperty'
import UnarchiveProperty from '@/services/api/Owner/Properties/UnarchiveProperty'
import DocumentsTab from '@/views/RealProperty/details/tabs/1DocumentsTab'
import DamageTab from '@/views/RealProperty/details/tabs/3DamageTab'
import InventoryTab from '@/views/RealProperty/details/tabs/2InventoryTab'

import { DetailsPartProps } from '@/interfaces/Property/Property'
import PropertyStatusEnum from '@/enums/PropertyEnum'
import PropertyHeader from './PropertyHeader'
import PropertyImage from './PropertyImage'
import PropertyInfo from './PropertyInfo'

import style from './DetailsPart.module.css'

interface ChildrenComponentProps {
  t: (key: string) => string
  propertyStatus: string
}

const ChildrenComponent: React.FC<ChildrenComponentProps> = ({
  t,
  propertyStatus
}) => {
  const items = [
    {
      key: '1',
      label: t('components.button.documents'),
      children: <DocumentsTab status={propertyStatus} />
    },
    {
      key: '2',
      label: t('components.button.inventory'),
      children: <InventoryTab />
    },
    {
      key: '3',
      label: t('components.button.damage'),
      children: <DamageTab status={propertyStatus} />
    }
  ]

  return (
    <div className={style.childrenContainer}>
      <Tabs style={{ width: '100%' }} defaultActiveKey="1" items={items} />
    </div>
  )
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

  const [currentStatus, setCurrentStatus] = useState(propertyData?.status || '')

  useEffect(() => {
    if (propertyData?.status) {
      setCurrentStatus(propertyData.status)
    }
  }, [propertyData])

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

  const recoverProperty = async () => {
    Modal.confirm({
      title: t('components.modal.unarchive_property.title'),
      content: t('components.modal.unarchive_property.description'),
      okText: t('components.button.confirm'),
      cancelText: t('components.button.cancel'),
      okButtonProps: { danger: true },
      onOk: async () => {
        if (!propertyData) {
          message.error('Property ID is missing.')
          return
        }
        try {
          await UnarchiveProperty(propertyData.id)
          message.success(t('components.modal.unarchive_property.success'))
          goToRealProperty()
        } catch (error) {
          console.error('Error deleting property:', error)
          message.error(t('components.modal.unarchive_property.error'))
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
          setCurrentStatus(PropertyStatusEnum.AVAILABLE)
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
          setCurrentStatus(PropertyStatusEnum.AVAILABLE)
          await refreshPropertyDetails(propertyData.id)
          message.success(t('components.modal.cancel_invitation.success'))
        } catch (error) {
          console.error('Error cancelling invitation:', error)
          message.error(t('components.modal.cancel_invitation.error'))
        }
      }
    })
  }

  return (
    <div className={style.mainContainer}>
      <PropertyHeader
        onShowModal={showModal}
        onShowModalUpdate={showModalUpdate}
        onEndContract={endContract}
        onCancelInvitation={cancelInvitation}
        onRemoveProperty={removeProperty}
        onRecoverProperty={recoverProperty}
        propertyStatus={currentStatus}
        propertyArchived={propertyData.archived || false}
      />
      <div className={style.headerInformationContainer}>
        <PropertyImage
          status={currentStatus}
          picture={picture}
          isLoading={isLoading}
        />
        <PropertyInfo propertyData={propertyData} />
      </div>
      <PropertyIdProvider id={propertyId}>
        <ChildrenComponent t={t} propertyStatus={currentStatus} />
      </PropertyIdProvider>
    </div>
  )
}

export default DetailsPart
