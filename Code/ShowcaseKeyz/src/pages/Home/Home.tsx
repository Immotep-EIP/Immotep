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
        <div className={style.shape4}>
          <svg
            viewBox="0 0 24 24"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
            className={style.keySvg}
          >
            <path
              d="M9 22H15C20 22 22 20 22 15V9C22 4 20 2 15 2H9C4 2 2 4 2 9V15C2 20 4 22 9 22Z"
              stroke="rgba(255, 255, 255, 0.5)"
              strokeWidth="1.5"
              strokeLinecap="round"
              strokeLinejoin="round"
            />
            <path
              d="M16.28 13.61C15.15 14.74 13.53 15.09 12.1 14.64L9.51 17.22C9.33 17.41 8.96 17.53 8.69 17.49L7.49 17.33C7.09 17.28 6.73 16.9 6.67 16.51L6.51 15.31C6.47 15.05 6.6 14.68 6.78 14.49L9.36 11.91C8.92 10.48 9.26 8.86 10.39 7.73C12.01 6.11 14.65 6.11 16.28 7.73C17.9 9.34 17.9 11.98 16.28 13.61Z"
              stroke="rgba(255, 255, 255, 0.5)"
              strokeWidth="1.5"
              strokeMiterlimit="10"
              strokeLinecap="round"
              strokeLinejoin="round"
            />
            <path
              d="M10.45 16.28L9.6 15.42"
              stroke="rgba(255, 255, 255, 0.5)"
              strokeWidth="1.5"
              strokeMiterlimit="10"
              strokeLinecap="round"
              strokeLinejoin="round"
            />
            <path
              d="M13.3944 10.7H13.4034"
              stroke="rgba(255, 255, 255, 0.5)"
              strokeWidth="2"
              strokeLinecap="round"
              strokeLinejoin="round"
            />
          </svg>
        </div>
      </div>
    </div>
  );
}

export default HomePage;
