import React from 'react'

import useNavigation from '@/hooks/Navigation/useNavigation'
import useImageCache from '@/hooks/Image/useImageCache'
import GetPropertyPicture from '@/services/api/Owner/Properties/GetPropertyPicture'

import { CardComponentProps } from '@/interfaces/Property/Property'
import { TenantStatusEnum } from '@/enums/PropertyEnum'
import { Card, Badge } from '@/components/common'

import defaultHouse from '@/assets/images/DefaultHouse.jpg'
import locationIcon from '@/assets/icons/location.svg'
import style from './PropertyCard.module.css'

const CardComponent: React.FC<CardComponentProps> = ({ realProperty, t }) => {
  const { goToRealPropertyDetails } = useNavigation()

  const { data: picture, isLoading } = useImageCache(
    realProperty.id,
    GetPropertyPicture
  )

  const statusText = t(
    TenantStatusEnum[realProperty!.status as keyof typeof TenantStatusEnum]
      .text || ''
  )

  const propertyAddress = (() => {
    if (realProperty.address && realProperty.postal_code && realProperty.city) {
      let fullAddress = ''
      if (realProperty.apartment_number) {
        fullAddress = `N° ${realProperty.apartment_number} - ${realProperty.address}, ${realProperty.postal_code} ${realProperty.city}`
      } else {
        fullAddress = `${realProperty.address}, ${realProperty.postal_code} ${realProperty.city}`
      }
      return fullAddress
    }
    return t('components.property.card.no_address')
  })()

  const displayAddress =
    propertyAddress.length > 40
      ? `${propertyAddress.substring(0, 40)}...`
      : propertyAddress

  const displayName =
    realProperty.name && realProperty.name.length > 35
      ? `${realProperty.name.substring(0, 35)}...`
      : realProperty.name

  return (
    <Card
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
      padding="none"
      aria-label={`${t('components.property.card.view_details')} ${realProperty.name} - ${statusText} - ${propertyAddress}`}
    >
      <Badge.Ribbon
        text={statusText}
        color={
          TenantStatusEnum[
            realProperty!.status as keyof typeof TenantStatusEnum
          ].color || 'default'
        }
      >
        <div className={style.cardPictureContainer}>
          <img
            src={isLoading ? defaultHouse : picture || defaultHouse}
            alt={`${t('components.property.card.image_alt')} ${realProperty.name || t('components.property.card.unnamed_property')}`}
            className={style.cardPicture}
          />
        </div>
      </Badge.Ribbon>
      <div className={style.cardInfoContainer}>
        <header>
          <h3 className={style.cardText} title={realProperty.name}>
            {displayName}
          </h3>
        </header>
        <div className={style.cardAddressContainer}>
          <img
            src={locationIcon}
            alt=""
            className={style.icon}
            aria-hidden="true"
          />
          <address className={style.cardText} title={propertyAddress}>
            {displayAddress}
          </address>
        </div>
      </div>
    </Card>
  )
}

export default CardComponent
