import { useTranslation } from "react-i18next";
import style from "./LegalStyle.module.css";

function LegalMentionsPage() {
  const { t } = useTranslation();

  return (
    <div className={style.pageContainer}>
      <div className={style.contentWrapper}>
        <div className={style.titleContainer}>
          <h1 className={style.title}>{t("legal_mentions.title")}</h1>
          <div className={style.titleUnderline}></div>
        </div>

        <div className={style.contentSection}>
          <div className={style.section}>
            <h2 className={style.sectionTitle}>
              {t("legal_mentions.section1.title")}
            </h2>
            <p className={style.paragraph}>
              {t("legal_mentions.section1.paragraph1")}
            </p>
            <p className={style.paragraph}>
              {t("legal_mentions.section1.paragraph2")}
            </p>
            <p className={style.paragraph}>
              {t("legal_mentions.section1.paragraph3")}
            </p>
          </div>

          <div className={style.section}>
            <h2 className={style.sectionTitle}>
              {t("legal_mentions.section2.title")}
            </h2>
            <p className={style.paragraph}>
              {t("legal_mentions.section2.paragraph1")}
            </p>
          </div>

          <div className={style.section}>
            <h2 className={style.sectionTitle}>
              {t("legal_mentions.section3.title")}
            </h2>
            <p className={style.paragraph}>
              {t("legal_mentions.section3.paragraph1")}
            </p>
          </div>

          <div className={style.section}>
            <h2 className={style.sectionTitle}>
              {t("legal_mentions.section4.title")}
            </h2>
            <p className={style.paragraph}>
              {t("legal_mentions.section4.paragraph1")}
            </p>
            <p className={style.paragraph}>
              {t("legal_mentions.section4.paragraph2")}
            </p>
          </div>

          <div className={style.section}>
            <h2 className={style.sectionTitle}>
              {t("legal_mentions.section5.title")}
            </h2>
            <p className={style.paragraph}>
              {t("legal_mentions.section5.paragraph1")}
            </p>
            <div className={style.contactInfo}>
              <span>{t("legal_mentions.section5.email")}</span>
              <span>{t("legal_mentions.section5.phone")}</span>
              <span>{t("legal_mentions.section5.mail")}</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

export default LegalMentionsPage;
