import { useTranslation } from "react-i18next";
import { useEffect, useRef, useState } from "react";
import Illustration from "../../assets/propertyDetails.png";
import style from "./OurApplication.module.css";

function OurApplicationPage() {
  const { t } = useTranslation();
  const titleRef = useRef<HTMLHeadingElement>(null);
  const [isVisible, setIsVisible] = useState(false);

  useEffect(() => {
    const observer = new IntersectionObserver(
      ([entry]) => {
        if (entry.isIntersecting) {
          setIsVisible(true);
        } else {
          setIsVisible(false);
        }
      },
      {
        threshold: 0.3,
        rootMargin: "0px 0px -100px 0px",
      }
    );
    if (titleRef.current) {
      observer.observe(titleRef.current);
    }
    return () => {
      if (titleRef.current) {
        observer.unobserve(titleRef.current);
      }
    };
  }, []);

  return (
    <div className={style.pageContainer}>
      <div className={style.decorGradient1}></div>
      <div className={style.decorGradient2}></div>
      <div className={style.decorDots}></div>
      <div className={style.titleContainer}>
        <h1 ref={titleRef} className={style.title}>
          {t("our_application.title")}
        </h1>
        <div
          className={`${style.titleUnderline} ${
            isVisible ? style.animate : ""
          }`}
        />
        <p className={style.subtitle}>{t("our_application.subtitle")}</p>
      </div>

      <div className={style.featureShowcase}>
        <div className={style.showcaseImageContainer}>
          <div className={style.showcaseImageDecor}></div>
          <img
            src={Illustration}
            alt={t("our_application.image_alt")}
            className={style.showcaseImage}
          />
        </div>
        <div className={style.showcaseContent}>
          <h2 className={style.showcaseTitle}>
            {t("our_application.showcase_title")}
          </h2>
          <p className={style.showcaseText}>
            {t("our_application.showcase_text_1")}
          </p>
          <p className={style.showcaseText}>
            {t("our_application.showcase_text_2")}
          </p>
          <div className={style.showcaseHighlights}>
            <span className={style.highlightItem}>
              <span className={style.highlightIcon}>✓</span>
              {t("our_application.highlight_1")}
            </span>
            <span className={style.highlightItem}>
              <span className={style.highlightIcon}>✓</span>
              {t("our_application.highlight_2")}
            </span>
            <span className={style.highlightItem}>
              <span className={style.highlightIcon}>✓</span>
              {t("our_application.highlight_3")}
            </span>
            <span className={style.highlightItem}>
              <span className={style.highlightIcon}>✓</span>
              {t("our_application.highlight_4")}
            </span>
          </div>
        </div>
      </div>

      <div className={style.devicesSection}>
        <h2 className={style.devicesSectionTitle}>
          {t("our_application.devices_title")}
        </h2>
        <div className={style.devicesContainer}>
          <div className={style.deviceItem}>
            <svg className={style.deviceIcon} viewBox="0 0 24 24" fill="none">
              <path
                d="M18 2H6a2 2 0 00-2 2v16a2 2 0 002 2h12a2 2 0 002-2V4a2 2 0 00-2-2z"
                stroke="currentColor"
                strokeWidth="2"
                strokeLinecap="round"
                strokeLinejoin="round"
              />
              <path
                d="M12 18h.01"
                stroke="currentColor"
                strokeWidth="2"
                strokeLinecap="round"
                strokeLinejoin="round"
              />
            </svg>
            <span>{t("our_application.device_mobile")}</span>
          </div>
          <div className={style.deviceItem}>
            <svg className={style.deviceIcon} viewBox="0 0 24 24" fill="none">
              <path
                d="M21 16H3a2 2 0 01-2-2V4a2 2 0 012-2h18a2 2 0 012 2v10a2 2 0 01-2 2z"
                stroke="currentColor"
                strokeWidth="2"
                strokeLinecap="round"
                strokeLinejoin="round"
              />
              <path
                d="M12 20v2M17 20v2M7 20v2"
                stroke="currentColor"
                strokeWidth="2"
                strokeLinecap="round"
                strokeLinejoin="round"
              />
            </svg>
            <span>{t("our_application.device_computer")}</span>
          </div>
        </div>
      </div>

      <div className={style.ctaSection}>
        <h2 className={style.ctaTitle}>{t("our_application.cta_title")}</h2>
        <p className={style.ctaText}>{t("our_application.cta_text")}</p>
        <a href="#contact-us" className={style.ctaButton}>
          {t("our_application.contact_us_now")}
          <svg className={style.arrowIcon} viewBox="0 0 24 24" fill="none">
            <path
              d="M5 12h14M12 5l7 7-7 7"
              stroke="currentColor"
              strokeWidth="2"
              strokeLinecap="round"
              strokeLinejoin="round"
            />
          </svg>
        </a>
      </div>
    </div>
  );
}

export default OurApplicationPage;
