import React from "react";
import style from './Text.module.css';

interface PageTitleProps {
  title: string;
  size?: 'title' | 'subtitle';
}

const PageTitle: React.FC<PageTitleProps> = ({ title, size }) => (
  <span
      className={style.pageTitle}
      style={{
        fontSize: size === 'subtitle' ? '1rem' : '1.4rem',
        fontWeight: size === 'subtitle' ? 400 : 500,
        marginBottom: '.5rem',
      }}
    >
      {title}
    </span>
);

export default PageTitle;
