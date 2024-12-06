import React from "react";
import { useTranslation } from "react-i18next";
import style from "./3InventoryTab.module.css";

const InventoryTab: React.FC = () => {
  const { t } = useTranslation();

  return (
    <div className={style.tabContent}>
      <span>{t("Inventory")}</span>
    </div>
  );
}

export default InventoryTab;
