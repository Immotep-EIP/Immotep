import { useTranslation } from 'react-i18next'

import SubtitledElement from '@/components/ui/SubtitledElement/SubtitledElement'

import { PropertyDetails } from '@/interfaces/Property/Property'

import style from './DetailsPart.module.css'

const PropertyInfo = ({ propertyData }: { propertyData: PropertyDetails }) => {
  const { t } = useTranslation()

  const formatDate = (date: string) =>
    new Date(date).toLocaleDateString('fr-FR', {
      day: 'numeric',
      month: 'long',
      year: 'numeric'
    })

  return (
    <div className={style.informationsContainer}>
      <div className={style.details}>
        <SubtitledElement
          subtitleKey={t('pages.real_property_details.informations.name')}
        >
          <span className={style.detailsText}>{propertyData.name}</span>
        </SubtitledElement>
      </div>
      <div className={style.details}>
        <SubtitledElement
          subtitleKey={t('pages.real_property_details.informations.address')}
        >
          <span className={style.detailsText}>
            {propertyData.apartment_number
              ? `N°${propertyData.apartment_number} - ${propertyData.address}, ${propertyData.postal_code} ${propertyData.city}`
              : `${propertyData.address}, ${propertyData.postal_code} ${propertyData.city}`}
          </span>
        </SubtitledElement>
      </div>
      <div className={style.details}>
        <SubtitledElement
          subtitleKey={t('pages.real_property_details.informations.tenant')}
        >
          <span className={style.detailsText}>
            {propertyData.lease?.tenant_name || '-----------'}
            {' - '}
            {propertyData.lease?.tenant_email || '-----------'}
          </span>
        </SubtitledElement>
      </div>
      <div className={style.details}>
        <SubtitledElement
          subtitleKey={t('pages.real_property_details.informations.dates')}
        >
          <span className={style.detailsText}>
            {propertyData.lease?.start_date
              ? formatDate(propertyData.lease.start_date)
              : '...'}
            {' - '}
            {propertyData.lease?.end_date
              ? formatDate(propertyData.lease.end_date)
              : '...'}
          </span>
        </SubtitledElement>
      </div>
      <div className={style.details}>
        <SubtitledElement
          subtitleKey={t('pages.real_property_details.informations.area')}
        >
          <span className={style.detailsText}>{propertyData.area_sqm} m²</span>
        </SubtitledElement>
      </div>
      <div className={style.details}>
        <SubtitledElement
          subtitleKey={t('pages.real_property_details.informations.rental')}
        >
          <span className={style.detailsText}>
            {propertyData.rental_price_per_month} €
          </span>
        </SubtitledElement>
      </div>
      <div className={style.details}>
        <SubtitledElement
          subtitleKey={t('pages.real_property_details.informations.deposit')}
        >
          <span className={style.detailsText}>
            {propertyData.deposit_price} €
          </span>
        </SubtitledElement>
      </div>
    </div>
  )
}

export default PropertyInfo
