import React, { useEffect, useState } from 'react'
import { Badge, Button, Empty, Typography } from 'antd'
import { useTranslation } from 'react-i18next'

import useNavigation from '@/hooks/useNavigation/useNavigation'
import useProperties from '@/hooks/useEffect/useProperties.ts'

import locationIcon from '@/assets/icons/location.png'

import PageTitle from '@/components/PageText/Title.tsx'
import defaultHouse from '@/assets/images/DefaultHouse.jpg'
import GetPropertyPicture from '@/services/api/Owner/Properties/GetPropertyPicture'
import useImageCache from '@/hooks/useEffect/useImageCache'
import CardPropertyLoader from '@/components/Loader/CardPropertyLoader'
import PageMeta from '@/components/PageMeta/PageMeta'
import PropertyStatusEnum from '@/enums/PropertyEnum'
import style from './RealProperty.module.css'
import RealPropertyCreate from './create/RealPropertyCreate'

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
      className={style.card}
      key={realProperty.id}
      role="button"
      tabIndex={0}
      onClick={() => goToRealPropertyDetails(realProperty.id)}
      onKeyDown={e => {
        if (e.key === 'Enter' || e.key === ' ') {
          goToRealPropertyDetails(realProperty.id)
        }
      }}
    >
      <div className={style.cardContentContainer}>
        <Badge.Ribbon
          text={
            realProperty.status === PropertyStatusEnum.AVAILABLE
              ? t('pages.real_property.status.available')
              : t('pages.real_property.status.unavailable')
          }
          color={
            realProperty.status === PropertyStatusEnum.AVAILABLE
              ? 'green'
              : 'red'
          }
        >
          <div className={style.cardPictureContainer}>
            <img
              src={isLoading ? defaultHouse : picture || defaultHouse}
              alt="property"
              className={style.cardPicture}
            />
          </div>
        </Badge.Ribbon>
        <div className={style.cardInfoContainer}>
          <b className={style.cardText}>{realProperty.name}</b>
          <div className={style.cardAddressContainer}>
            <img src={locationIcon} alt="location" className={style.icon} />
            <span className={style.cardText}>
              {realProperty.apartment_number &&
              realProperty.address &&
              realProperty.postal_code &&
              realProperty.city
                ? (() => {
                    const fullAddress = `${realProperty.apartment_number} - ${realProperty.address}, ${realProperty.postal_code} ${realProperty.city}`
                    return fullAddress.length > 40
                      ? `${fullAddress.substring(0, 40)}...`
                      : fullAddress
                  })()
                : '-----------'}
            </span>
          </div>
        </div>
      </div>
    </div>
  )
}

const RealPropertyPage: React.FC = () => {
  const { t } = useTranslation()
  const { properties, loading, error, refreshProperties } = useProperties()
  const [showModalCreate, setShowModalCreate] = useState(false)
  const [isPropertyCreated, setIsPropertyCreated] = useState(false)

  useEffect(() => {
    if (isPropertyCreated) {
      refreshProperties()
      setIsPropertyCreated(false)
    }
  }, [isPropertyCreated, refreshProperties])

  if (error) {
    return <p>{t('pages.real_property.error.error_fetching_data')}</p>
  }

  return (
    <>
      <PageMeta
        title={t('pages.real_property.document_title')}
        description={t('pages.real_property.document_description')}
        keywords="real property, Property info, Immotep"
      />
      <div className={style.pageContainer}>
        <div className={style.pageHeader}>
          <PageTitle title={t('pages.real_property.title')} size="title" />
          <Button type="primary" onClick={() => setShowModalCreate(true)}>
            {t('components.button.add_real_property')}
          </Button>
        </div>

        {loading && <CardPropertyLoader cards={12} />}

        <div className={style.cardsContainer}>
          {properties.length === 0 && (
            <div className={style.emptyContainer}>
              <Empty
                image="https://gw.alipayobjects.com/zos/antfincdn/ZHrcdLPrvN/empty.svg"
                description={
                  <Typography.Text>
                    {t('components.messages.no_properties')}
                  </Typography.Text>
                }
              />
            </div>
          )}
          {properties.map(realProperty => (
            <CardComponent
              key={realProperty.id}
              realProperty={realProperty}
              t={t}
            />
          ))}
        </div>
        <RealPropertyCreate
          showModalCreate={showModalCreate}
          setShowModalCreate={setShowModalCreate}
          setIsPropertyCreated={setIsPropertyCreated}
        />
      </div>
    </>
  )
}

export default RealPropertyPage
