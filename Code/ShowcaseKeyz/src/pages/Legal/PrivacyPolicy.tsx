import { useTranslation } from "react-i18next";
import style from "./LegalStyle.module.css";

function PrivacyPolicyPage() {
  const { t } = useTranslation();

  return (
    <div className={style.pageContainer}>
      <div className={style.contentWrapper}>
        <div className={style.titleContainer}>
          <h1 className={style.title}>{t("privacy_policy.title")}</h1>
          <div className={style.titleUnderline}></div>
        </div>

        <div className={style.contentSection}>
          <div className={style.section}>
            <h2 className={style.sectionTitle}>
              {t("privacy_policy.section1.title")}
            </h2>
            <p className={style.paragraph}>
              {t("privacy_policy.section1.paragraph1")}
            </p>
            <p className={style.paragraph}>
              {t("privacy_policy.section1.paragraph2")}
            </p>
          </div>

          <div className={style.section}>
            <h2 className={style.sectionTitle}>
              {t("privacy_policy.section2.title")}
            </h2>
            <p className={style.paragraph}>
              {t("privacy_policy.section2.paragraph1")}
            </p>
            <ul className={style.list}>
              <li className={style.listItem}>
                {t("privacy_policy.section2.list.item1")}
              </li>
              <li className={style.listItem}>
                {t("privacy_policy.section2.list.item2")}
              </li>
              <li className={style.listItem}>
                {t("privacy_policy.section2.list.item3")}
              </li>
              <li className={style.listItem}>
                {t("privacy_policy.section2.list.item4")}
              </li>
              <li className={style.listItem}>
                {t("privacy_policy.section2.list.item5")}
              </li>
            </ul>
          </div>

          <div className={style.section}>
            <h2 className={style.sectionTitle}>
              {t("privacy_policy.section3.title")}
            </h2>
            <p className={style.paragraph}>
              {t("privacy_policy.section3.paragraph1")}
            </p>
            <ul className={style.list}>
              <li className={style.listItem}>
                {t("privacy_policy.section3.list.item1")}
              </li>
              <li className={style.listItem}>
                {t("privacy_policy.section3.list.item2")}
              </li>
              <li className={style.listItem}>
                {t("privacy_policy.section3.list.item3")}
              </li>
              <li className={style.listItem}>
                {t("privacy_policy.section3.list.item4")}
              </li>
              <li className={style.listItem}>
                {t("privacy_policy.section3.list.item5")}
              </li>
            </ul>
          </div>

          <div className={style.section}>
            <h2 className={style.sectionTitle}>
              {t("privacy_policy.section4.title")}
            </h2>
            <p className={style.paragraph}>
              {t("privacy_policy.section4.paragraph1")}
            </p>
            <ul className={style.list}>
              <li className={style.listItem}>
                {t("privacy_policy.section4.list.item1")}
              </li>
              <li className={style.listItem}>
                {t("privacy_policy.section4.list.item2")}
              </li>
              <li className={style.listItem}>
                {t("privacy_policy.section4.list.item3")}
              </li>
              <li className={style.listItem}>
                {t("privacy_policy.section4.list.item4")}
              </li>
            </ul>
          </div>

          <div className={style.section}>
            <h2 className={style.sectionTitle}>
              {t("privacy_policy.section5.title")}
            </h2>
            <p className={style.paragraph}>
              {t("privacy_policy.section5.paragraph1")}
            </p>
            <ul className={style.list}>
              <li className={style.listItem}>
                {t("privacy_policy.section5.list.item1")}
              </li>
              <li className={style.listItem}>
                {t("privacy_policy.section5.list.item2")}
              </li>
              <li className={style.listItem}>
                {t("privacy_policy.section5.list.item3")}
              </li>
              <li className={style.listItem}>
                {t("privacy_policy.section5.list.item4")}
              </li>
              <li className={style.listItem}>
                {t("privacy_policy.section5.list.item5")}
              </li>
              <li className={style.listItem}>
                {t("privacy_policy.section5.list.item6")}
              </li>
            </ul>
          </div>

          <div className={style.section}>
            <h2 className={style.sectionTitle}>
              {t("privacy_policy.section6.title")}
            </h2>
            <p className={style.paragraph}>
              {t("privacy_policy.section6.paragraph1")}
            </p>
            <div className={style.contactInfo}>
              <span>{t("privacy_policy.section6.email")}</span>
              <span>{t("privacy_policy.section6.mail")}</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

export default PrivacyPolicyPage;
