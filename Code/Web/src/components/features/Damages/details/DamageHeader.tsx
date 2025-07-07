import React from 'react'
import { useTranslation } from 'react-i18next'
import { useNavigate, useLocation } from 'react-router-dom'

import { Button } from '@/components/common'
import PageTitle from '@/components/ui/PageText/Title'

import NavigationEnum from '@/enums/NavigationEnum'

import returnIcon from '@/assets/icons/retour.svg'
import style from './DetailsPart.module.css'

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
    <header className={style.moreInfosContainer} role="banner">
      <div className={style.titleContainer}>
        <Button
          type="text"
          style={{ border: 'none', backgroundColor: 'transparent' }}
          className={style.returnButtonContainer}
          onClick={navigateToPropertyDetails}
          aria-label={`${t('components.button.return')} - ${t('pages.damage_details.title')}`}
        >
          <img
            src={returnIcon}
            alt="Return icon"
            className={style.returnIcon}
            aria-hidden="true"
          />
        </Button>
        <PageTitle
          title={t('pages.damage_details.title')}
          size="title"
          margin={false}
          id="damage-header-title"
        />
      </div>
    </header>
  )
}

export default DamageHeader
