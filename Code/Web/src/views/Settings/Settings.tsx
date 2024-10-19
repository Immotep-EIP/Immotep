import React from 'react'
import i18n from 'i18next'
import { Button } from 'antd'
import style from './Settings.module.css'

const Settings: React.FC = () => {

  const switchLanguage = (language: string) => {
    i18n.changeLanguage(language)
  }

  return (
    <div className={style.layoutContainer}>
      <Button color="primary" size="large" onClick={() => switchLanguage('fr')}>
        fr
      </Button>
      <Button color="primary" size="large" onClick={() => switchLanguage('en')}>
        en
      </Button>
    </div>
  )
}

export default Settings
