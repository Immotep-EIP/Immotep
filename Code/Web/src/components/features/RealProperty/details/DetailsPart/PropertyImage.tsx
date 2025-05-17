import React, { memo } from 'react'
import PropTypes from 'prop-types'
import { Badge } from 'antd'
import { useTranslation } from 'react-i18next'
import defaultHouse from '@/assets/images/DefaultHouse.jpg'
import { TenantStatusEnum } from '@/enums/PropertyEnum'
import { PropertyImageProps } from '@/interfaces/Property/Property'
import style from './DetailsPart.module.css'

const PropertyImage: React.FC<PropertyImageProps> = memo(
  ({ status, picture, isLoading }) => {
    const { t } = useTranslation()

    const statusText =
      status && TenantStatusEnum[status as keyof typeof TenantStatusEnum]?.text
        ? t(TenantStatusEnum[status as keyof typeof TenantStatusEnum].text)
        : t('pages.real_property.status.unknown')

    const statusColor =
      status && TenantStatusEnum[status as keyof typeof TenantStatusEnum]?.color
        ? TenantStatusEnum[status as keyof typeof TenantStatusEnum].color
        : 'default'

    return (
      <div style={{ width: '55%', height: '400px' }}>
        <Badge.Ribbon text={statusText} color={statusColor}>
          <img
            src={isLoading ? defaultHouse : picture || defaultHouse}
            alt="Property"
            className={style.propertyPicture}
          />
        </Badge.Ribbon>
      </div>
    )
  }
)

PropertyImage.propTypes = {
  status: PropTypes.string.isRequired,
  picture: PropTypes.string,
  isLoading: PropTypes.bool.isRequired
}

PropertyImage.displayName = 'PropertyImage'
export default PropertyImage
