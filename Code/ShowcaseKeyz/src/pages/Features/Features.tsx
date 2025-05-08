import { useTranslation } from "react-i18next";
import { useEffect, useRef, useState } from "react";
import IA from "../../assets/ia.svg";
import Generator from "../../assets/generator.svg";
import Documents from "../../assets/documents.svg";
import Messages from "../../assets/messages.svg";
import Damage from "../../assets/damage.svg";
import Dashboard from "../../assets/dashboard.svg";
import style from "./Features.module.css";

function FeaturesPage() {
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

  const features = [
    {
      title: t("features.feature1.title"),
      description: t("features.feature1.description"),
      icon: IA,
    },
    {
      title: t("features.feature2.title"),
      description: t("features.feature2.description"),
      icon: Generator,
    },
    {
      title: t("features.feature3.title"),
      description: t("features.feature3.description"),
      icon: Documents,
    },
    {
      title: t("features.feature4.title"),
      description: t("features.feature4.description"),
      icon: Messages,
    },
    {
      title: t("features.feature5.title"),
      description: t("features.feature5.description"),
      icon: Damage,
    },
    {
      title: t("features.feature6.title"),
      description: t("features.feature6.description"),
      icon: Dashboard,
    },
  ];

  return (
    <div className={style.pageContainer}>
      <div className={style.decorShape1}></div>
      <div className={style.decorShape2}></div>
      <div className={style.decorDots}></div>
      <div className={style.decorTriangle}></div>
      <div className={style.titleContainer}>
        <h1 ref={titleRef} className={style.title}>
          {t("features.title")}
        </h1>
        <div
          className={`${style.titleUnderline} ${
            isVisible ? style.animate : ""
          }`}
        />
      </div>
      <div className={style.featuresContainer}>
        {features.map((feature, index) => (
          <div key={index} className={style.featureCard}>
            <div className={style.iconContainer}>
              <img
                src={feature.icon}
                alt={`${feature.title} icon`}
                className={style.icon}
              />
            </div>
            <h2 className={style.featureTitle}>{feature.title}</h2>
            <p className={style.featureDescription}>{feature.description}</p>
          </div>
        ))}
      </div>
    </div>
  );
}

export default FeaturesPage;
