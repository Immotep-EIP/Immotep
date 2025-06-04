import { useTranslation } from 'react-i18next'

import SubtitledElement from '@/components/ui/SubtitledElement/SubtitledElement'
import { usePropertyContext } from '@/context/propertyContext'

import { PropertyDetails } from '@/interfaces/Property/Property'
import { Lease } from '@/interfaces/Property/Lease/Lease'

import style from './DetailsPart.module.css'

const PropertyInfo = ({
  propertyData
}: {
  propertyData: PropertyDetails & { leases: Lease[] }
}) => {
  const { t } = useTranslation()
  const { selectedLease } = usePropertyContext()

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
            {selectedLease?.tenant_name || '-----------'}
            {' - '}
            {selectedLease?.tenant_email || '-----------'}
          </span>
        </SubtitledElement>
      </div>
      <div className={style.details}>
        <SubtitledElement
          subtitleKey={t('pages.real_property_details.informations.dates')}
        >
          <span className={style.detailsText}>
            {selectedLease?.start_date
              ? formatDate(selectedLease.start_date)
              : '...'}
            {' - '}
            {selectedLease?.end_date
              ? formatDate(selectedLease.end_date)
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
