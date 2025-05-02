import React from 'react'
import { Badge } from 'antd'
import { useTranslation } from 'react-i18next'
import defaultHouse from '@/assets/images/DefaultHouse.jpg'
import { TenantStatusEnum } from '@/enums/PropertyEnum'
import { PropertyImageProps } from '@/interfaces/Property/Property'
import style from './DetailsPart.module.css'

const PropertyImage: React.FC<PropertyImageProps> = ({
  status,
  picture,
  isLoading
}) => {
  const { t } = useTranslation()

  return (
    <div style={{ width: '55%', height: '400px' }}>
      <Badge.Ribbon
        text={t(
          TenantStatusEnum[status as keyof typeof TenantStatusEnum].text || ''
        )}
        color={
          TenantStatusEnum[status as keyof typeof TenantStatusEnum].color ||
          'default'
        }
      >
        <img
          src={isLoading ? defaultHouse : picture || defaultHouse}
          alt="Property"
          className={style.propertyPicture}
        />
      </Badge.Ribbon>
    </div>
  )
}

export default PropertyImage
