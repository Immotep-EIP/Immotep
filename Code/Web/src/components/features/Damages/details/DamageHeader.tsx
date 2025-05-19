import React from 'react'
import { useTranslation } from 'react-i18next'
import style from './DetailsPart.module.css'
import PageTitle from '@/components/ui/PageText/Title'
import returnIcon from '@/assets/icons/retour.svg'

const DamageHeader: React.FC = () => {
  const { t } = useTranslation()
  return (
    <div className={style.moreInfosContainer}>
      <div className={style.titleContainer}>
        <div
          className={style.returnButtonContainer}
          onClick={() => window.history.back()}
          tabIndex={0}
          role="button"
          onKeyDown={e => {
            if (e.key === 'Enter') {
              window.history.back()
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
