import React from 'react'

import PageTitle from '@/components/ui/PageText/Title'

import { AuthentificationPageProps } from '@/interfaces/Auth/Auth'

import logo from '@/assets/images/KeyzLogo.svg'
import style from './AuthentificationPage.module.css'

const AuthentificationPage: React.FC<AuthentificationPageProps> = ({
  title,
  subtitle,
  children
}) => (
  <div className={style.pageContainer}>
    <header className={style.headerContainer} role="banner">
      <img
        src={logo}
        alt="Keyz - Property Management Platform Logo"
        className={style.headerLogo}
      />
      <span className={style.headerTitle} aria-label="Keyz Application">
        Keyz
      </span>
    </header>

    <main className={style.contentContainer} role="main">
      <PageTitle title={title} size="title" id="page-title" />
      <PageTitle title={subtitle} size="subtitle" />

      <div
        className={style.childrenContainer}
        role="form"
        aria-labelledby="page-title"
      >
        {children}
      </div>
    </main>
  </div>
)

export default AuthentificationPage
