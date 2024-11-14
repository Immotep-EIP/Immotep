import React from 'react'
import { Button, Tag } from 'antd'
import { useTranslation } from 'react-i18next'
import { RealProperty } from '@/interfaces/Property/Property.tsx'
import useNavigation from '@/hooks/useNavigation/useNavigation'

import style from './RealProperty.module.css'

import locationIcon from '../../assets/icons/location.png'
import tenantIcon from '../../assets/icons/tenant.png'
import dateIcon from '../../assets/icons/date.png'

import fakeData from '../../fakeDatas/RealProperties.tsx'

interface CardComponentProps {
  realProperty: RealProperty
  t: (key: string) => string
}

const CardComponent: React.FC<CardComponentProps> = ({ realProperty, t }) => {
  const { goToRealPropertyDetails } = useNavigation()

  return (
    <div
      key={realProperty.id}
      className={style.card}
      role="button"
      tabIndex={0}
      onClick={() => goToRealPropertyDetails(realProperty.id)}
      onKeyDown={e => {
        if (e.key === 'Enter' || e.key === ' ') {
          goToRealPropertyDetails(realProperty.id)
        }
      }}
    >
      {/* FIRST PART */}
      <div className={style.statusContainer}>
        <Tag
          color={
            realProperty.status === 'pages.property.status.available'
              ? 'green'
              : 'red'
          }
        >
          {t(realProperty.status)}
        </Tag>
        {realProperty.damages.some(damage => !damage.read) && (
          <Tag color="red">
            ({realProperty.damages.filter(damage => !damage.read).length}){' '}
            {t('pages.property.damage.unread')}
          </Tag>
        )}
      </div>

      {/* SECOND PART */}
      <div className={style.pictureContainer}>
        <img
          src={realProperty.image}
          alt="property"
          className={style.picture}
        />
      </div>

      {/* THIRD PART */}
      <div className={style.informationsContainer}>
        <div className={style.informations}>
          <img src={locationIcon} alt="location" className={style.icon} />
          <span>
            {realProperty.adress && realProperty.zipCode && realProperty.city
              ? (() => {
                  const fullAddress = `${realProperty.adress}, ${realProperty.zipCode} ${realProperty.city}`
                  return fullAddress.length > 40
                    ? `${fullAddress.substring(0, 40)}...`
                    : fullAddress
                })()
              : '-----------'}
          </span>
        </div>
        <div className={style.informations}>
          <img src={tenantIcon} alt="location" className={style.icon} />
          <span>
            {realProperty.tenants.length > 0
              ? realProperty.tenants.map(tenant => tenant.name).join(' & ')
              : '-----------'}
          </span>
        </div>
        <div className={style.informations}>
          <img src={dateIcon} alt="location" className={style.icon} />
          <span>
            {realProperty.startDate && realProperty.endDate
              ? `${new Date(realProperty.startDate).toLocaleDateString('fr-FR', { day: 'numeric', month: 'long', year: 'numeric' })} - ${new Date(realProperty.endDate).toLocaleDateString('fr-FR', { day: 'numeric', month: 'long', year: 'numeric' })}`
              : '-----------'}
          </span>
        </div>
      </div>
    </div>
  )
}

const RealPropertyPage: React.FC = () => {
  const { t } = useTranslation()
  const { goToRealPropertyCreate } = useNavigation()

  return (
    <div className={style.pageContainer}>
      <div className={style.pageHeader}>
        <span className={style.pageTitle}>
          {t('pages.real_property.title')}
        </span>
        <Button type="primary" onClick={goToRealPropertyCreate}>
          {t('components.button.add_real_property')}
        </Button>
      </div>
      <div className={style.cardsContainer}>
        {fakeData.map(realProperty => (
          <CardComponent
            key={realProperty.id}
            realProperty={realProperty}
            t={t}
          />
        ))}
      </div>
    </div>
  )
}

export default RealPropertyPage
