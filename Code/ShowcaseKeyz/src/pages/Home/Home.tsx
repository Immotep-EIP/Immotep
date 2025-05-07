import style from "./Home.module.css";
import { useTranslation } from "react-i18next";

function HomePage() {
  const { t } = useTranslation();

  return (
    <div className={style.pageContainer}>
      <div className={style.pageContent}>
        <h1 className={style.title}>{t("home.title")}</h1>
        <h6 className={style.description}>{t("home.description")}</h6>
        <a className={style.button} href="#features">
          {t("home.button")}
        </a>
      </div>
      <div className={style.shapes}>
        <div className={style.shape1}></div>
        <div className={style.shape2}></div>
        <div className={style.shape3}></div>
      </div>
    </div>
  );
}

export default HomePage;
