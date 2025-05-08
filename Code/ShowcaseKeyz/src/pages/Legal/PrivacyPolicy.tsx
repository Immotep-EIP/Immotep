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
            <h2 className={style.sectionTitle}>1. Introduction</h2>
            <p className={style.paragraph}>
              Chez Keyz, nous prenons la protection de vos données personnelles
              très au sérieux. Cette politique de confidentialité explique
              comment nous collectons, utilisons, partageons et protégeons vos
              informations lorsque vous utilisez notre site web et notre
              application.
            </p>
            <p className={style.paragraph}>
              En utilisant nos services, vous acceptez les pratiques décrites
              dans cette politique de confidentialité. Nous vous encourageons à
              la lire attentivement.
            </p>
          </div>

          <div className={style.section}>
            <h2 className={style.sectionTitle}>2. Données collectées</h2>
            <p className={style.paragraph}>
              Nous collectons les informations suivantes :
            </p>
            <ul className={style.list}>
              <li className={style.listItem}>
                Informations d'identification : nom, prénom, adresse email,
                numéro de téléphone
              </li>
              <li className={style.listItem}>
                Informations relatives à vos biens immobiliers : adresse,
                caractéristiques, photos, documents
              </li>
              <li className={style.listItem}>
                Informations sur les locataires : coordonnées, documents
                administratifs
              </li>
              <li className={style.listItem}>
                Données de paiement : pour la facturation de nos services
              </li>
              <li className={style.listItem}>
                Données d'utilisation : comment vous interagissez avec notre
                plateforme
              </li>
            </ul>
          </div>

          <div className={style.section}>
            <h2 className={style.sectionTitle}>3. Utilisation des données</h2>
            <p className={style.paragraph}>
              Nous utilisons vos données personnelles pour :
            </p>
            <ul className={style.list}>
              <li className={style.listItem}>
                Fournir, maintenir et améliorer nos services
              </li>
              <li className={style.listItem}>
                Traiter vos transactions et gérer votre compte
              </li>
              <li className={style.listItem}>
                Vous envoyer des notifications importantes concernant nos
                services
              </li>
              <li className={style.listItem}>
                Vous proposer un support client personnalisé
              </li>
              <li className={style.listItem}>
                Améliorer et personnaliser votre expérience utilisateur
              </li>
            </ul>
          </div>

          <div className={style.section}>
            <h2 className={style.sectionTitle}>4. Partage des données</h2>
            <p className={style.paragraph}>
              Nous ne vendons pas vos données personnelles à des tiers. Nous
              pouvons les partager dans les situations suivantes :
            </p>
            <ul className={style.list}>
              <li className={style.listItem}>
                Avec nos prestataires de services qui nous aident à fournir nos
                services
              </li>
              <li className={style.listItem}>
                Pour respecter la loi, les réglementations applicables ou les
                procédures légales
              </li>
              <li className={style.listItem}>
                Pour protéger la sécurité de nos utilisateurs et du public
              </li>
              <li className={style.listItem}>
                Dans le cadre d'une transaction d'entreprise (fusion,
                acquisition, etc.)
              </li>
            </ul>
          </div>

          <div className={style.section}>
            <h2 className={style.sectionTitle}>5. Vos droits</h2>
            <p className={style.paragraph}>
              Conformément au Règlement Général sur la Protection des Données
              (RGPD), vous disposez des droits suivants :
            </p>
            <ul className={style.list}>
              <li className={style.listItem}>
                Droit d'accès à vos données personnelles
              </li>
              <li className={style.listItem}>
                Droit de rectification des informations inexactes
              </li>
              <li className={style.listItem}>
                Droit à l'effacement de vos données dans certaines conditions
              </li>
              <li className={style.listItem}>
                Droit à la limitation du traitement
              </li>
              <li className={style.listItem}>
                Droit à la portabilité de vos données
              </li>
              <li className={style.listItem}>
                Droit d'opposition au traitement
              </li>
            </ul>
          </div>

          <div className={style.section}>
            <h2 className={style.sectionTitle}>6. Contact</h2>
            <p className={style.paragraph}>
              Pour toute question concernant cette politique de confidentialité
              ou pour exercer vos droits, veuillez nous contacter :
            </p>
            <div className={style.contactInfo}>
              <span>Par email : contact@keyz-app.fr</span>
              <span>Par courrier : /</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

export default PrivacyPolicyPage;
