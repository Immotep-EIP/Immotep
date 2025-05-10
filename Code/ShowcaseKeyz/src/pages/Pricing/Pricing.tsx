import { useTranslation } from "react-i18next";
import { useEffect, useRef, useState } from "react";
import style from "./Pricing.module.css";

function PricingPage() {
  const { t } = useTranslation();
  const titleRef = useRef<HTMLHeadingElement>(null);
  const [isVisible, setIsVisible] = useState(false);
  const [hoveredCard, setHoveredCard] = useState<number | null>(null);

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

  const handleCtaClick = () => {
    // Option 1: Pour naviguer vers la page de contact dans une SPA
    const contactElement = document.getElementById("contact-us");
    if (contactElement) {
      contactElement.scrollIntoView({ behavior: "smooth" });
    } else {
      // Option 2: Si le contact est sur une autre page, naviguer vers cette page avec l'ancre
      window.location.href = "/#contact-us";
    }
  };

  const prices = [
    {
      title: t("pricing.plans.basic.title"),
      price: "4,99€",
      period: t("pricing.period"),
      features: [
        t("pricing.plans.basic.features.feature1"),
        t("pricing.plans.basic.features.feature2"),
        t("pricing.plans.basic.features.feature3"),
        t("pricing.plans.basic.features.feature4"),
        t("pricing.plans.basic.features.feature5"),
      ],
      // cta: t("pricing.cta"), // ! TODO : Uncomment this and remove the next line when the CTA is ready
      cta: t("pricing.contact_us"),
      popular: false,
      icon: (
        <svg viewBox="0 0 24 24" fill="none" className={style.planIcon}>
          <path
            d="M3 9l9-7 9 7v11a2 2 0 01-2 2H5a2 2 0 01-2-2V9z"
            stroke="currentColor"
            strokeWidth="2"
            strokeLinecap="round"
            strokeLinejoin="round"
          />
          <path
            d="M9 22V12h6v10"
            stroke="currentColor"
            strokeWidth="2"
            strokeLinecap="round"
            strokeLinejoin="round"
          />
        </svg>
      ),
    },
    {
      title: t("pricing.plans.premium.title"),
      price: "4,99€",
      period: t("pricing.period"),
      features: [
        t("pricing.plans.premium.features.feature1"),
        t("pricing.plans.premium.features.feature2"),
        t("pricing.plans.premium.features.feature3"),
        t("pricing.plans.premium.features.feature4"),
        t("pricing.plans.premium.features.feature5"),
      ],
      // cta: t("pricing.cta"), // ! TODO : Uncomment this and remove the next line when the CTA is ready
      cta: t("pricing.contact_us"),
      popular: true,
      icon: (
        <svg viewBox="0 0 24 24" fill="none" className={style.planIcon}>
          <path
            d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5m0 0v-5a2 2 0 012-2h2a2 2 0 012 2v5m0 0h-6"
            stroke="currentColor"
            strokeWidth="2"
            strokeLinecap="round"
            strokeLinejoin="round"
          />
        </svg>
      ),
    },
    {
      title: t("pricing.plans.pro.title"),
      price: "28,99€",
      period: t("pricing.period"),
      features: [
        t("pricing.plans.pro.features.feature1"),
        t("pricing.plans.pro.features.feature2"),
        t("pricing.plans.pro.features.feature3"),
        t("pricing.plans.pro.features.feature4"),
        t("pricing.plans.pro.features.feature5"),
        t("pricing.plans.pro.features.feature6"),
      ],
      cta: t("pricing.contact_us"),
      popular: false,
      icon: (
        <svg viewBox="0 0 24 24" fill="none" className={style.planIcon}>
          <path
            d="M2 9a2 2 0 012-2h16a2 2 0 012 2v9a2 2 0 01-2 2H4a2 2 0 01-2-2V9z"
            stroke="currentColor"
            strokeWidth="2"
            strokeLinecap="round"
            strokeLinejoin="round"
          />
          <path
            d="M16 13a2 2 0 100-4 2 2 0 000 4z"
            stroke="currentColor"
            strokeWidth="2"
            strokeLinecap="round"
            strokeLinejoin="round"
          />
          <path
            d="M22 9V7a2 2 0 00-2-2H4a2 2 0 00-2 2v2"
            stroke="currentColor"
            strokeWidth="2"
            strokeLinecap="round"
            strokeLinejoin="round"
          />
        </svg>
      ),
    },
  ];

  return (
    <div className={style.pageContainer}>
      <div className={style.decorCircle1}></div>
      <div className={style.decorCircle2}></div>

      <div className={style.titleContainer}>
        <h1 ref={titleRef} className={style.title}>
          {t("pricing.title")}
        </h1>
        <div
          className={`${style.titleUnderline} ${
            isVisible ? style.animate : ""
          }`}
        />
        <p className={style.subtitle}>{t("pricing.subtitle")}</p>
      </div>

      <div className={style.pricingContainer}>
        {prices.map((price, index) => (
          <div
            key={index}
            className={`${style.priceCard} ${
              price.popular ? style.popular : ""
            } ${hoveredCard === index ? style.hovered : ""}`}
            onMouseEnter={() => setHoveredCard(index)}
            onMouseLeave={() => setHoveredCard(null)}
          >
            {price.popular && (
              <div className={style.popularBadge}>
                {t("pricing.most_popular")}
              </div>
            )}
            <div className={style.iconContainer}>{price.icon}</div>
            <h2 className={style.priceTitle}>{price.title}</h2>
            {price.price && (
              <div className={style.priceWrapper}>
                <span className={style.pricePrefix}>{t("pricing.from")} </span>
                <span className={style.priceAmount}>{price.price}</span>
                <span className={style.pricePeriod}>{price.period}</span>
              </div>
            )}
            {index === 1 && (
              <div className={style.additionalInfo}>
                {t("pricing.plans.premium.additional_info")}
              </div>
            )}
            <ul className={style.featuresList}>
              {price.features.map((feature, idx) => (
                <li key={idx} className={style.featureItem}>
                  <svg
                    className={style.checkIcon}
                    viewBox="0 0 24 24"
                    fill="none"
                  >
                    <path
                      d="M5 13l4 4L19 7"
                      stroke="currentColor"
                      strokeWidth="2"
                      strokeLinecap="round"
                      strokeLinejoin="round"
                    />
                  </svg>
                  {feature}
                </li>
              ))}
            </ul>
            <button onClick={handleCtaClick} className={style.ctaButton}>
              {price.cta}
            </button>
          </div>
        ))}
      </div>
    </div>
  );
}

export default PricingPage;
