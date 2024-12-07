import React from 'react'
import { Button } from 'antd'
import { useTranslation } from 'react-i18next'

import useNavigation from '@/hooks/useNavigation/useNavigation'
import useFetchProperties from "@/hooks/useEffect/useFetchProperties.ts";

import appartmentIcon from '@/assets/icons/appartement.png'
import locationIcon from '@/assets/icons/location.png'
import tenantIcon from '@/assets/icons/tenant.png'
import dateIcon from '@/assets/icons/date.png'

import PageTitle from "@/components/PageText/Title.tsx";
import defaultHouse from '@/assets/images/DefaultHouse.jpg'
import style from './RealProperty.module.css'

interface CardComponentProps {
  realProperty: any
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
        soon available
        {/* <Tag
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
        )} */}
      </div>

      {/* SECOND PART */}
      <div className={style.pictureContainer}>
        <img
          src={realProperty.image || defaultHouse}
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
          <img src={locationIcon} alt="location" className={style.icon} />
          {/* <span className={style.text}>
            {realProperty.address && realProperty.postal_code && realProperty.city
              ? `${realProperty.address}, ${realProperty.postal_code} ${realProperty.city}`
              : '-----------'}
          </span> */}
          <span>
            {realProperty.address && realProperty.postal_code && realProperty.city
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
          <img src={tenantIcon} alt="location" className={style.icon} />
          <span>
            soon available
            {/* {realProperty.tenants.length > 0
              ? realProperty.tenants.map(tenant => tenant.name).join(' & ')
              : '-----------'} */}
          </span>
        </div>
        <div className={style.informations}>
          <img src={dateIcon} alt="location" className={style.icon} />
          <span>
            soon available
            {/* {realProperty.startDate && realProperty.endDate
              ? `${new Date(realProperty.startDate).toLocaleDateString('fr-FR', { day: 'numeric', month: 'long', year: 'numeric' })} - ${new Date(realProperty.endDate).toLocaleDateString('fr-FR', { day: 'numeric', month: 'long', year: 'numeric' })}`
              : '-----------'} */}
          </span>
        </div>
      </div>
    </div>
  )
}

const RealPropertyPage: React.FC = () => {
  const { t } = useTranslation()
  const { goToRealPropertyCreate } = useNavigation()
  const { properties, loading, error } = useFetchProperties();

  if (loading) {
    return <p>{t("generals.loading")}</p>;
  }

  if (error) {
    return <p>{t("pages.property.error.errorFetchingData")}</p>;
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
