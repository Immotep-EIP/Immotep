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
              1. Informations sur l'éditeur
            </h2>
            <p className={style.paragraph}>Ce site web est édité par Keyz.</p>
            <p className={style.paragraph}>
              Numéro de TVA intracommunautaire : /
            </p>
            <p className={style.paragraph}>
              Directeur de la publication : Oscar FRANK, Co Founder
            </p>
          </div>

          <div className={style.section}>
            <h2 className={style.sectionTitle}>2. Hébergement</h2>
            <p className={style.paragraph}>
              Ce site est hébergé par LWS, société située 2, rue Jules Ferry,
              88190 Golbey
            </p>
          </div>

          <div className={style.section}>
            <h2 className={style.sectionTitle}>3. Propriété intellectuelle</h2>
            <p className={style.paragraph}>
              L'ensemble des éléments figurant sur ce site (textes, images,
              logos, etc.) sont protégés par les lois françaises et
              internationales relatives à la propriété intellectuelle. Toute
              reproduction ou représentation totale ou partielle de ce site ou
              de tout ou partie des éléments y figurant est strictement
              interdite sans l'autorisation préalable de Keyz SAS.
            </p>
          </div>

          <div className={style.section}>
            <h2 className={style.sectionTitle}>4. Données personnelles</h2>
            <p className={style.paragraph}>
              Les informations recueillies sur ce site font l'objet d'un
              traitement informatique destiné à Keyz pour la gestion des clients
              et prospects. Conformément à la loi « informatique et libertés »
              du 6 janvier 1978 modifiée et au Règlement Général sur la
              Protection des Données (RGPD), vous disposez d'un droit d'accès,
              de rectification, et d'opposition aux informations qui vous
              concernent.
            </p>
            <p className={style.paragraph}>
              Pour exercer ces droits, veuillez nous contacter à l'adresse :
              contact@keyz-app.fr
            </p>
          </div>

          <div className={style.section}>
            <h2 className={style.sectionTitle}>5. Contact</h2>
            <p className={style.paragraph}>
              Pour toute question concernant ces mentions légales, vous pouvez
              nous contacter :
            </p>
            <div className={style.contactInfo}>
              <span>Par email : 2. contact@keyz-app.fr</span>
              <span>Par téléphone : /</span>
              <span>Par courrier : /</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

export default LegalMentionsPage;
