import { useTranslation } from "react-i18next";
import style from "./Demo.module.css";
import { useEffect, useRef, useState } from "react";

function DemoPage() {
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
      <div className={style.decorShape1}></div>
      <div className={style.decorShape2}></div>
      <div className={style.decorDots}></div>
      <div className={style.decorTriangle}></div>
      <div className={style.titleContainer}>
        <h1 ref={titleRef} className={style.title}>
          {t("demo.title")}
        </h1>
        <div
          className={`${style.titleUnderline} ${
            isVisible ? style.animate : ""
          }`}
        />
      </div>
      <div className={style.contentContainer}>
        <div className={style.comingSoonWrapper}>
          <h2 className={style.comingSoonTitle}>{t("demo.coming_soon")}</h2>
          <p className={style.comingSoonText}>
            {t("demo.development_message")}
          </p>

          <div className={style.progressContainer}>
            <div className={style.progressBar}></div>
          </div>
          <div className={style.percentageText}>75%</div>
        </div>
      </div>
    </div>
  );
}

export default DemoPage;
