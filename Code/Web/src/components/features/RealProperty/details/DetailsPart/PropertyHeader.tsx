import React from 'react'
import { useTranslation } from 'react-i18next'

import { Dropdown, MenuProps, Select, Tag } from 'antd'
import { MoreOutlined } from '@ant-design/icons'

import PageTitle from '@/components/ui/PageText/Title'
import { Button } from '@/components/common'
import { usePropertyContext } from '@/context/propertyContext'
import useNavigation from '@/hooks/Navigation/useNavigation'
import toLocaleDate from '@/utils/date/toLocaleDate'

import { PropertyHeaderProps } from '@/interfaces/Property/Property'
import PropertyStatusEnum from '@/enums/PropertyEnum'

import returnIcon from '@/assets/icons/retour.svg'
import style from './DetailsPart.module.css'

const PropertyHeader: React.FC<PropertyHeaderProps> = ({
  onShowModal,
  onShowModalUpdate,
  onEndContract,
  onCancelInvitation,
  onRemoveProperty,
  onRecoverProperty,
  propertyStatus,
  propertyArchived
}) => {
  const { t } = useTranslation()
  const { goToRealProperty } = useNavigation()
  const { property, selectedLeaseId, setSelectedLeaseId } = usePropertyContext()

  const onLeaseChange = (value: string) => {
    setSelectedLeaseId(value)
  }

  const items: MenuProps['items'] = [
    {
      key: '1',
      label: t('components.button.add_tenant'),
      onClick: onShowModal,
      disabled:
        propertyStatus === PropertyStatusEnum.UNAVAILABLE ||
        propertyStatus === PropertyStatusEnum.INVITATION_SENT ||
        propertyArchived
    },
    {
      key: '2',
      label: t('components.button.end_contract'),
      onClick: onEndContract,
      danger: true,
      disabled: propertyStatus !== PropertyStatusEnum.UNAVAILABLE
    },
    {
      key: '3',
      label: t('components.button.cancel_invitation'),
      onClick: onCancelInvitation,
      danger: true,
      disabled:
        propertyStatus !== PropertyStatusEnum.INVITATION_SENT ||
        propertyArchived
    },
    {
      key: '4',
      label: t('components.button.edit_property'),
      onClick: onShowModalUpdate
    },
    {
      key: '5',
      label: propertyArchived
        ? t('components.button.unarchive_property')
        : t('components.button.archive_property'),
      danger: true,
      onClick: propertyArchived ? onRecoverProperty : onRemoveProperty,
      disabled: propertyStatus === PropertyStatusEnum.UNAVAILABLE
    }
  ]

  const handleBack = () => {
    if (propertyArchived) {
      goToRealProperty(true)
    } else {
      goToRealProperty()
    }
  }

  return (
    <header className={style.moreInfosContainer} role="banner">
      <div className={style.titleContainer}>
        <Button
          className={style.returnButtonContainer}
          type="text"
          style={{ border: 'none', backgroundColor: 'transparent' }}
          onClick={handleBack}
          aria-label={`${t('components.button.return')} - ${t('pages.real_property_details.title')}`}
        >
          <img
            src={returnIcon}
            alt="Return icon"
            className={style.returnIcon}
            aria-hidden="true"
          />
        </Button>
        <PageTitle
          id="property-header-title"
          title={t('pages.real_property_details.title')}
          size="title"
          margin={false}
        />
      </div>
      <div
        style={{ display: 'flex', alignItems: 'center', gap: 16 }}
        role="toolbar"
        aria-label={t('pages.real_property_details.title')}
      >
        {property?.leases && property?.leases.length > 0 && (
          <Select
            value={selectedLeaseId}
            onChange={onLeaseChange}
            style={{ width: 280, textAlign: 'left' }}
            allowClear={property?.leases.length > 1 || !property?.lease?.active}
            aria-label={t('pages.real_property_details.select_default')}
            options={[
              ...(property?.leases?.map(lease => ({
                value: lease.id,
                label: (
                  <div
                    style={{
                      display: 'flex',
                      alignItems: 'center',
                      gap: '8px'
                    }}
                  >
                    <span>
                      {toLocaleDate(lease.start_date, 'short')}
                      {' - '}
                      {(lease.end_date &&
                        toLocaleDate(lease.end_date, 'short')) ||
                        '...'}
                    </span>
                    {lease.active && (
                      <Tag color="blue" style={{ margin: 0 }}>
                        {t('pages.real_property_details.current')}
                      </Tag>
                    )}
                  </div>
                )
              })) || [])
            ]}
            placeholder={t('pages.real_property_details.select_default')}
          />
        )}
        <Dropdown
          menu={{ items }}
          trigger={['click']}
          placement="bottomRight"
          aria-label={t('components.button.more_actions')}
        >
          <Button
            type="text"
            aria-label={t('components.button.more_actions')}
            aria-expanded="false"
            aria-haspopup="true"
            style={{
              border: '0.4px solid rgba(0, 0, 0, 0.8)',
              borderRadius: 4,
              padding: 6,
              display: 'inline-flex',
              alignItems: 'center',
              cursor: 'pointer',
              background: 'transparent'
            }}
          >
            <MoreOutlined />
          </Button>
        </Dropdown>
      </div>
    </header>
  )
}

export default PropertyHeader
