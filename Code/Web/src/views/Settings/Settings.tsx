import React, { useState } from 'react'
import { useTranslation } from 'react-i18next'
import i18n from 'i18next'
import { Segmented } from 'antd'
import style from './Settings.module.css'

const Settings: React.FC = () => {
  const { t } = useTranslation()
  const [activeLanguage] = useState(i18n.language)

  const switchLanguage = (language: string) => {
    let lang = ''

    switch (language) {
      case t('pages.settings.fr') as string:
        lang = 'fr'
        break
      case t('pages.settings.en') as string:
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
            options={[t('pages.settings.fr'), t('pages.settings.en')]}
            defaultValue={t(`pages.settings.${activeLanguage}`)}
            onChange={value => {
              switchLanguage(value)
            }}
          />
        </div>

      </div>
    </div>
  )
}

export default Settings
