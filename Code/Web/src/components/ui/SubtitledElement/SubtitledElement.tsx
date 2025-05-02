import React from 'react';
import { useTranslation } from 'react-i18next';
import style from './SubtitledElement.module.css';

interface SubtitledElementProps {
  subtitleKey: string;
  children: React.ReactNode;
  subTitleStyle?: React.CSSProperties;
}

const SubtitledElement = ({ subtitleKey, children, subTitleStyle = {} }: SubtitledElementProps) => {
  const { t } = useTranslation();

  return (
    <div className={style.box} style={subTitleStyle}>
      <span className={style.subtitle}>{t(subtitleKey)}</span>
      {children}
    </div>
  );
};

export default SubtitledElement;