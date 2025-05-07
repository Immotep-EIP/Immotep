import { useTranslation } from "react-i18next";
import { useEffect, useRef, useState } from "react";
import style from "./Features.module.css";

const features = [
  {
    title: "Feature 1",
    description:
      "This description must make around 2 or 3 lines. It is a long description.",
    icon: "https://simpleicons.org/icons/github.svg",
  },
  {
    title: "Feature 2",
    description:
      "This description must make around 2 or 3 lines. It is a long description.",
    icon: "https://simpleicons.org/icons/github.svg",
  },
  {
    title: "Feature 3",
    description:
      "This description must make around 2 or 3 lines. It is a long description.",
    icon: "https://simpleicons.org/icons/github.svg",
  },
  {
    title: "Feature 4",
    description:
      "This description must make around 2 or 3 lines. It is a long description.",
    icon: "https://simpleicons.org/icons/github.svg",
  },
  {
    title: "Feature 5",
    description:
      "This description must make around 2 or 3 lines. It is a long description.",
    icon: "https://simpleicons.org/icons/github.svg",
  },
  {
    title: "Feature 6",
    description:
      "This description must make around 2 or 3 lines. It is a long description.",
    icon: "https://simpleicons.org/icons/github.svg",
  },
];

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

  return (
    <div className={style.pageContainer}>
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
            <img
              src={feature.icon}
              alt={feature.title}
              className={style.icon}
            />
            <h2 className={style.featureTitle}>{feature.title}</h2>
            <p className={style.featureDescription}>{feature.description}</p>
          </div>
        ))}
      </div>
    </div>
  );
}

export default FeaturesPage;
