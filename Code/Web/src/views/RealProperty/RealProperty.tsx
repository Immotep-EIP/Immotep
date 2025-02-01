import React from 'react'
import { Button, Tag } from 'antd'
import { useTranslation } from 'react-i18next'

import useNavigation from '@/hooks/useNavigation/useNavigation'
import useProperties from '@/hooks/useEffect/useProperties.ts'

import appartmentIcon from '@/assets/icons/appartement.png'
import locationIcon from '@/assets/icons/location.png'
import tenantIcon from '@/assets/icons/tenant.png'
import dateIcon from '@/assets/icons/date.png'

import PageTitle from '@/components/PageText/Title.tsx'
import defaultHouse from '@/assets/images/DefaultHouse.jpg'
import GetPropertyPicture from '@/services/api/Owner/Properties/GetPropertyPicture'
import useImageCache from '@/hooks/useEffect/useImageCache'
import style from './RealProperty.module.css'

interface CardComponentProps {
  realProperty: any
  t: (key: string) => string
}

const CardComponent: React.FC<CardComponentProps> = ({ realProperty, t }) => {
  const { goToRealPropertyDetails } = useNavigation()

  const { data: picture, isLoading } = useImageCache(
    realProperty.id,
    GetPropertyPicture
  )

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
        <Tag color={realProperty.status === 'available' ? 'green' : 'red'}>
          {realProperty.status === 'available'
            ? t('pages.real_property.status.available')
            : t('pages.real_property.status.unavailable')}
        </Tag>
        <Tag color={realProperty.nb_damage > 0 ? 'red' : 'green'}>
          {realProperty.nb_damage || 0}{' '}
          {t('pages.real_property.damage.waiting')}
        </Tag>
      </div>

      {/* SECOND PART */}
      <div className={style.pictureContainer}>
        <img
          src={isLoading ? defaultHouse : picture || defaultHouse}
          alt="property"
          className={style.picture}
        />
      </div>

      {/* THIRD PART */}
      <div className={style.informationsContainer}>
        <div className={style.informations}>
          <img src={appartmentIcon} alt="location" className={style.icon} />
          <span>
            {(() => {
              if (realProperty.name) {
                return realProperty.name.length > 40
                  ? `${realProperty.name.substring(0, 40)}...`
                  : realProperty.name
              }
              return '-----------'
            })()}
          </span>
        </div>
        <div className={style.informations}>
          <img src={locationIcon} alt="locationIcon" className={style.icon} />
          <span>
            {realProperty.address &&
            realProperty.postal_code &&
            realProperty.city
              ? (() => {
                  const fullAddress = `${realProperty.address}, ${realProperty.postal_code} ${realProperty.city}`
                  return fullAddress.length > 40
                    ? `${fullAddress.substring(0, 40)}...`
                    : fullAddress
                })()
              : '-----------'}
          </span>
        </div>
        <div className={style.informations}>
          <img src={tenantIcon} alt="tenantIcon" className={style.icon} />
          <span>
            {realProperty.tenant ? realProperty.tenant : '-----------'}
          </span>
        </div>
        <div className={style.informations}>
          <img src={dateIcon} alt="dateIcon" className={style.icon} />
          <span>
            {realProperty.start_date
              ? `${new Date(realProperty.start_date).toLocaleDateString('fr-FR', { day: 'numeric', month: 'long', year: 'numeric' })}`
              : '...'}
            {' - '}
            {realProperty.end_date
              ? `${new Date(realProperty.end_date).toLocaleDateString('fr-FR', { day: 'numeric', month: 'long', year: 'numeric' })}`
              : '...'}
          </span>
        </div>
      </div>
    </div>
  )
}

const RealPropertyPage: React.FC = () => {
  const { t } = useTranslation()
  const { goToRealPropertyCreate } = useNavigation()
  const { properties, loading, error } = useProperties()

  if (loading) {
    return <p>{t('generals.loading')}</p>
  }

  if (error) {
    return <p>{t('pages.real_property.error.error_fetching_data')}</p>
  }

  return (
    <div className={style.pageContainer}>
      <div className={style.pageHeader}>
        <PageTitle title={t('pages.real_property.title')} size="title" />
        <Button type="primary" onClick={goToRealPropertyCreate}>
          {t('components.button.add_real_property')}
        </Button>
      </div>
      <div className={style.cardsContainer}>
        {properties.map(realProperty => (
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
