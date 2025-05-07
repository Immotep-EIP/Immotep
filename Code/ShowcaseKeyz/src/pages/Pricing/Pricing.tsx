import { useTranslation } from "react-i18next";
import { useEffect, useRef, useState } from "react";
import style from "./Pricing.module.css";

const prices = [
  {
    title: "Basic",
    price: "$10",
    period: "/month",
    features: [
      "1 property",
      "Basic tenant management",
      "Document storage",
      "Email support",
    ],
    cta: "Get Started",
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
    title: "Enterprise",
    price: "$30",
    period: "/month",
    features: [
      "Unlimited properties",
      "Advanced tenant management",
      "Financial analytics",
      "Document storage & automation",
      "Priority support 24/7",
    ],
    cta: "Get Started",
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
    title: "Pro",
    price: "$20",
    period: "/month",
    features: [
      "5 properties",
      "Advanced tenant management",
      "Basic analytics",
      "Document storage",
      "Email & chat support",
    ],
    cta: "Get Started",
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
              <div className={style.popularBadge}>Most Popular</div>
            )}
            <div className={style.iconContainer}>{price.icon}</div>
            <h2 className={style.priceTitle}>{price.title}</h2>
            <div className={style.priceWrapper}>
              <span className={style.priceAmount}>{price.price}</span>
              <span className={style.pricePeriod}>{price.period}</span>
            </div>
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
            <button className={style.ctaButton}>{price.cta}</button>
          </div>
        ))}
      </div>
    </div>
  );
}

export default PricingPage;
