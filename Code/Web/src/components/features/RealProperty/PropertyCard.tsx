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
    >
      <Badge.Ribbon
        text={t(
          TenantStatusEnum[
            realProperty!.status as keyof typeof TenantStatusEnum
          ].text || ''
        )}
        color={
          TenantStatusEnum[
            realProperty!.status as keyof typeof TenantStatusEnum
          ].color || 'default'
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
        <b className={style.cardText}>
          {realProperty.name && realProperty.name.length > 35
            ? `${realProperty.name.substring(0, 35)}...`
            : realProperty.name}
        </b>
        <div className={style.cardAddressContainer}>
          <img src={locationIcon} alt="location" className={style.icon} />
          <span className={style.cardText}>
            {realProperty.address &&
            realProperty.postal_code &&
            realProperty.city
              ? (() => {
                  let fullAddress = ''
                  if (realProperty.apartment_number) {
                    fullAddress = `N° ${realProperty.apartment_number} - ${realProperty.address}, ${realProperty.postal_code} ${realProperty.city}`
                  } else {
                    fullAddress = `${realProperty.address}, ${realProperty.postal_code} ${realProperty.city}`
                  }
                  return fullAddress.length > 40
                    ? `${fullAddress.substring(0, 40)}...`
                    : fullAddress
                })()
              : '-----------'}
          </span>
        </div>
      </div>
    </Card>
  )
}

export default CardComponent
