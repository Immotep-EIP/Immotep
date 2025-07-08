import React from 'react'
import { useTranslation } from 'react-i18next'

import logo from '@/assets/images/KeyzLogo.svg'
import style from './DividedPage.module.css'

interface DividedPageProps {
  childrenLeft: React.ReactNode
  childrenRight: React.ReactNode
}

const DividedPage: React.FC<DividedPageProps> = ({
  childrenLeft,
  childrenRight
}) => {
  const { t } = useTranslation()

  return (
    <div className={style.dividedPageContainer}>
      <aside
        className={style.dividedPageLeft}
        aria-labelledby="left-content-title"
      >
        <h2 id="left-content-title" className="sr-only">
          {t('components.layout.divided_page.left_content')}
        </h2>
        {childrenLeft}
      </aside>

      <main className={style.dividedPageRight}>
        <header className={style.headerContainer} role="banner">
          <img
            src={logo}
            alt={t('components.layout.divided_page.logo_alt')}
            className={style.headerLogo}
          />
          <h1 className={style.headerTitle}>Keyz</h1>
        </header>

        <section
          className={style.childrenRightContainer}
          aria-labelledby="main-content-title"
        >
          <h2 id="main-content-title" className="sr-only">
            {t('components.layout.divided_page.main_content')}
          </h2>
          {childrenRight}
        </section>
      </main>
    </div>
  )
}

export default DividedPage
