import React from 'react'
import logo from '@/assets/icons/ImmotepLogo.svg'
import style from './DividedPage.module.css'

interface DividedPageProps {
  childrenLeft: React.ReactNode
  childrenRight: React.ReactNode
}

const DividedPage: React.FC<DividedPageProps> = ({
  childrenLeft,
  childrenRight
}) => (
  <div className={style.dividedPageContainer}>
    <div className={style.dividedPageLeft}>{childrenLeft}</div>
    <div className={style.dividedPageRight}>
      <div className={style.headerContainer}>
        <img src={logo} alt="logo Immotep" className={style.headerLogo} />
        <span className={style.headerTitle}>Immotep</span>
      </div>
      <div className={style.childrenRightContainer}>{childrenRight}</div>
    </div>
  </div>
)

export default DividedPage
