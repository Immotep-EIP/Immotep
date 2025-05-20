import React from 'react'
import { useTranslation } from 'react-i18next'
import { useNavigate, useLocation } from 'react-router-dom'
import style from './DetailsPart.module.css'
import PageTitle from '@/components/ui/PageText/Title'
import returnIcon from '@/assets/icons/retour.svg'
import NavigationEnum from '@/enums/NavigationEnum'

const DamageHeader: React.FC = () => {
  const { t } = useTranslation()
  const navigate = useNavigate()
  const location = useLocation()

  const navigateToPropertyDetails = () => {
    const path = location.pathname

    const match = path.match(/\/real-property\/details\/([^/]+)/)

    if (match && match[1]) {
      const propertyId = match[1]
      navigate(NavigationEnum.REAL_PROPERTY_DETAILS.replace(':id', propertyId))
    } else {
      window.history.back()
    }
  }

  return (
    <div className={style.moreInfosContainer}>
      <div className={style.titleContainer}>
        <div
          className={style.returnButtonContainer}
          onClick={navigateToPropertyDetails}
          tabIndex={0}
          role="button"
          onKeyDown={e => {
            if (e.key === 'Enter') {
              navigateToPropertyDetails()
            }
          }}
        >
          <img src={returnIcon} alt="Return" className={style.returnIcon} />
        </div>
        <PageTitle
          title={t('pages.damage_details.title')}
          size="title"
          margin={false}
        />
      </div>
    </div>
  )
}

export default DamageHeader
