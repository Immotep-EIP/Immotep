import React from 'react'
import { useTranslation } from 'react-i18next'
import i18n from 'i18next'
import { Segmented } from 'antd'
import style from './Settings.module.css'

const Settings: React.FC = () => {
  const { t } = useTranslation()
  const switchLanguage = (language: string) => {
    let lang = ''

    switch (language) {
      case 'fr' as string:
        lang = 'fr'
        break
      case 'en' as string:
        lang = 'en'
        break
      default:
        lang = 'fr'
        break
    }
    i18n.changeLanguage(lang)
  }

  return (
    <div className={style.layoutContainer}>
      <div className={style.settingsContainer}>

        <div className={style.settingsItem}>
          {t('pages.settings.language')}
          <Segmented
            options={[
              { label: t('pages.settings.fr'), value: 'fr' },
              { label: t('pages.settings.en'), value: 'en' }
            ]}
            value={i18n.language}
            onChange={(value) => switchLanguage(value as string)}
          />
        </div>

      </div>
    </div>
  )
}

export default Settings
