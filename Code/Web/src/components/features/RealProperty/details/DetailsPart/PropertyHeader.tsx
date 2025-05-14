import React from 'react'
import { Button, Dropdown, MenuProps } from 'antd'
import { MoreOutlined } from '@ant-design/icons'
import { useTranslation } from 'react-i18next'
import returnIcon from '@/assets/icons/retour.svg'
import PageTitle from '@/components/ui/PageText/Title'
import { PropertyHeaderProps } from '@/interfaces/Property/Property'
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

  const items: MenuProps['items'] = [
    {
      key: '1',
      label: t('components.button.add_tenant'),
      onClick: onShowModal,
      disabled:
        propertyStatus === 'UNAVAILABLE' ||
        propertyStatus === 'INVITATION_SENT' ||
        propertyArchived
    },
    {
      key: '2',
      label: t('components.button.end_contract'),
      onClick: onEndContract,
      danger: true,
      disabled: propertyStatus !== 'UNAVAILABLE' || propertyArchived
    },
    {
      key: '3',
      label: t('components.button.cancel_invitation'),
      onClick: onCancelInvitation,
      danger: true,
      disabled: propertyStatus !== 'INVITATION_SENT' || propertyArchived
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
      onClick: propertyArchived ? onRecoverProperty : onRemoveProperty
    }
  ]

  return (
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
  )
}

export default PropertyHeader
