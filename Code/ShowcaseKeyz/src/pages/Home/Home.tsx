import style from "./Home.module.css";
import { useTranslation } from "react-i18next";
import { useEffect, useState } from "react";

function HomePage() {
  const { t } = useTranslation();
  const [prefersReducedMotion, setPrefersReducedMotion] = useState(false);

  useEffect(() => {
    const mql = window.matchMedia("(prefers-reduced-motion: reduce)");
    setPrefersReducedMotion(mql.matches);

    const handleChange = (e: MediaQueryListEvent) => {
      setPrefersReducedMotion(e.matches);
    };

    mql.addEventListener("change", handleChange);
    return () => mql.removeEventListener("change", handleChange);
  }, []);

  return (
    <section
      id="home"
      className={style.pageContainer}
      aria-labelledby="home-title"
    >
      <div className={style.pageContent}>
        <h1 id="home-title" className={style.title}>
          {t("home.title")}
        </h1>
        <p className={style.description}>{t("home.description")}</p>
        <a
          className={style.button}
          href="#features"
          aria-label={t("home.button_aria")}
        >
          {t("home.button")}
        </a>
      </div>

      <div className={style.floatingElements} aria-hidden="true">
        <div
          className={`${style.floatingIcon} ${style.key1} ${
            prefersReducedMotion ? style.noAnimation : ""
          }`}
        >
          <svg
            viewBox="0 0 24 24"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
          >
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

        <div
          className={`${style.floatingIcon} ${style.key2} ${
            prefersReducedMotion ? style.noAnimation : ""
          }`}
        >
          <svg
            viewBox="0 0 24 24"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              d="M16.28 13.61C15.15 14.74 13.53 15.09 12.1 14.64L9.51 17.22C9.33 17.41 8.96 17.53 8.69 17.49L7.49 17.33C7.09 17.28 6.73 16.9 6.67 16.51L6.51 15.31C6.47 15.05 6.6 14.68 6.78 14.49L9.36 11.91C8.92 10.48 9.26 8.86 10.39 7.73C12.01 6.11 14.65 6.11 16.28 7.73C17.9 9.34 17.9 11.98 16.28 13.61Z"
              stroke="rgba(255, 255, 255, 0.4)"
              strokeWidth="1.5"
              strokeMiterlimit="10"
              strokeLinecap="round"
              strokeLinejoin="round"
            />
            <path
              d="M13.3944 10.7H13.4034"
              stroke="rgba(255, 255, 255, 0.4)"
              strokeWidth="2"
              strokeLinecap="round"
              strokeLinejoin="round"
            />
          </svg>
        </div>

        <div
          className={`${style.floatingIcon} ${style.house1} ${
            prefersReducedMotion ? style.noAnimation : ""
          }`}
        >
          <svg
            viewBox="0 0 24 24"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              d="M3 9L12 2L21 9V22H3V9Z"
              stroke="rgba(255, 255, 255, 0.4)"
              strokeWidth="1.5"
              strokeLinecap="round"
              strokeLinejoin="round"
            />
            <path
              d="M9 22V12H15V22"
              stroke="rgba(255, 255, 255, 0.4)"
              strokeWidth="1.5"
              strokeLinecap="round"
              strokeLinejoin="round"
            />
          </svg>
        </div>

        <div
          className={`${style.floatingIcon} ${style.building1} ${
            prefersReducedMotion ? style.noAnimation : ""
          }`}
        >
          <svg
            viewBox="0 0 24 24"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              d="M19 2H5C3.89543 2 3 2.89543 3 4V20C3 21.1046 3.89543 22 5 22H19C20.1046 22 21 21.1046 21 20V4C21 2.89543 20.1046 2 19 2Z"
              stroke="rgba(255, 255, 255, 0.35)"
              strokeWidth="1.5"
              strokeLinecap="round"
              strokeLinejoin="round"
            />
            <path
              d="M7 22V17"
              stroke="rgba(255, 255, 255, 0.35)"
              strokeWidth="1.5"
              strokeLinecap="round"
              strokeLinejoin="round"
            />
            <path
              d="M17 22V17"
              stroke="rgba(255, 255, 255, 0.35)"
              strokeWidth="1.5"
              strokeLinecap="round"
              strokeLinejoin="round"
            />
            <path
              d="M7 12H17"
              stroke="rgba(255, 255, 255, 0.35)"
              strokeWidth="1.5"
              strokeLinecap="round"
              strokeLinejoin="round"
            />
            <path
              d="M7 7H17"
              stroke="rgba(255, 255, 255, 0.35)"
              strokeWidth="1.5"
              strokeLinecap="round"
              strokeLinejoin="round"
            />
          </svg>
        </div>

        <div
          className={`${style.floatingIcon} ${style.building2} ${
            prefersReducedMotion ? style.noAnimation : ""
          }`}
        >
          <svg
            viewBox="0 0 24 24"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              d="M1 22H23"
              stroke="rgba(255, 255, 255, 0.3)"
              strokeWidth="1.5"
              strokeLinecap="round"
              strokeLinejoin="round"
            />
            <path
              d="M6 22V4C6 3.44772 6.44772 3 7 3H17C17.5523 3 18 3.44772 18 4V22"
              stroke="rgba(255, 255, 255, 0.3)"
              strokeWidth="1.5"
              strokeLinecap="round"
              strokeLinejoin="round"
            />
            <path
              d="M10 8H14"
              stroke="rgba(255, 255, 255, 0.3)"
              strokeWidth="1.5"
              strokeLinecap="round"
              strokeLinejoin="round"
            />
            <path
              d="M10 12H14"
              stroke="rgba(255, 255, 255, 0.3)"
              strokeWidth="1.5"
              strokeLinecap="round"
              strokeLinejoin="round"
            />
            <path
              d="M10 16H14"
              stroke="rgba(255, 255, 255, 0.3)"
              strokeWidth="1.5"
              strokeLinecap="round"
              strokeLinejoin="round"
            />
          </svg>
        </div>

        <div
          className={`${style.floatingIcon} ${style.house2} ${
            prefersReducedMotion ? style.noAnimation : ""
          }`}
        >
          <svg
            viewBox="0 0 24 24"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              d="M4 10L12 3L20 10L20 20H4L4 10Z"
              stroke="rgba(255, 255, 255, 0.25)"
              strokeWidth="1.5"
              strokeLinecap="round"
              strokeLinejoin="round"
            />
            <path
              d="M12 3L12 7"
              stroke="rgba(255, 255, 255, 0.25)"
              strokeWidth="1.5"
              strokeLinecap="round"
              strokeLinejoin="round"
            />
          </svg>
        </div>
      </div>

      <div className={style.shapes} aria-hidden="true">
        <div className={style.shape1}></div>
        <div className={style.shape2}></div>
        <div className={style.shape3}></div>
      </div>
    </section>
  );
}

export default HomePage;
