import React from "react";
import { useTranslation } from "react-i18next";
import style from "./4DocumentsTab.module.css";

const DocumentsTab: React.FC = () => {
  const { t } = useTranslation();

  return (
    <div className={style.tabContent}>
      <span>{t("Documents")}</span>
    </div>
  );
}

export default DocumentsTab;
