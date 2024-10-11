import React from "react";
import logo from '@/assets/icons/ImmotepLogo.svg';
import PageTitle from '@/components/PageText/Text';
import style from './AuthentificationPage.module.css';

interface AuthentificationPageProps {
  title: string;
  subtitle: string;
  children: React.ReactNode;
}

const AuthentificationPage: React.FC<AuthentificationPageProps> = (
  { title, subtitle, children }) => (
    <div className={style.pageContainer}>

      <div className={style.headerContainer}>
        <img src={logo} alt="logo Immotep" className={style.headerLogo} />
        <span className={style.headerTitle}>Immotep</span>
      </div>

      <div className={style.contentContainer}>
        <PageTitle title={title} size="title" />
        <PageTitle title={subtitle} size="subtitle" />

        <div className={style.childrenContainer}>
          {children}
        </div>
      </div>

    </div>
  );

export default AuthentificationPage;
