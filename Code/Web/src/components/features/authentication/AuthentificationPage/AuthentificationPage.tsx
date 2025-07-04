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
    <div className={style.headerContainer}>
      <img src={logo} alt="logo Keyz" className={style.headerLogo} />
      <span className={style.headerTitle}>Keyz</span>
    </div>

    <div className={style.contentContainer}>
      <PageTitle title={title} size="title" />
      <PageTitle title={subtitle} size="subtitle" />
    </div>

    <div className={style.childrenContainer}>{children}</div>
  </div>
)

export default AuthentificationPage
