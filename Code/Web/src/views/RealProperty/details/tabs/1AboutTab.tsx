import React from "react";
import { useTranslation } from "react-i18next";
import style from "./1AboutTab.module.css";

const AboutTab: React.FC = () => {
  const { t } = useTranslation();

  return (
    <div className={style.tabContent}>
      <span>{t("About")}</span>
    </div>
  );
}

export default AboutTab;
