import { useTranslation } from "react-i18next";
import InstagramIcon from "../../assets/Instagram.png";
import FacebookIcon from "../../assets/Facebook.png";
import LinkedinIcon from "../../assets/Linkedin.png";
import style from "./Footer.module.css";

const Footer = () => {
  const { t } = useTranslation();

  return (
    <div className={style.footerContainer}>
      <div className={style.columnContainer}>
        <span className={style.appTitle}>{t("footer.title")}</span>
        <span className={style.appDesc}>{t("footer.description")}</span>
        <div className={style.iconsContainer}>
          <div className={style.iconContainer}>
            <img
              src={InstagramIcon}
              alt="Instagram"
              className={style.icon}
              onClick={() =>
                window.open("https://www.instagram.com/app.keyz/", "_blank")
              }
            />
          </div>
          <div className={style.iconContainer}>
            <img
              src={FacebookIcon}
              alt="Facebook"
              className={style.icon}
              onClick={() =>
                window.open(
                  "https://www.facebook.com/profile.php?id=61575702764536",
                  "_blank"
                )
              }
            />
          </div>
          <div className={style.iconContainer}>
            <img
              src={LinkedinIcon}
              alt="LinkedIn"
              className={style.icon}
              onClick={() =>
                window.open(
                  "https://www.linkedin.com/company/keyz-app/",
                  "_blank"
                )
              }
            />
          </div>
        </div>
      </div>

      <div className={style.columnContainer}>
        <span className={style.colTitle}>{t("footer.pages.title")}</span>
        <a href="#home" className={style.colDesc}>
          {t("footer.pages.home")}
        </a>
        <a href="#features" className={style.colDesc}>
          {t("footer.pages.features")}
        </a>
        <a href="#application" className={style.colDesc}>
          {t("footer.pages.our_app")}
        </a>
        <a href="#pricing" className={style.colDesc}>
          {t("footer.pages.pricing")}
        </a>
        <a href="#contact-us" className={style.colDesc}>
          {t("footer.pages.contact_us")}
        </a>
      </div>
      <div className={style.columnContainer}>
        <span className={style.colTitle}>{t("footer.legal.title")}</span>
        <span className={style.colDesc}>{t("footer.legal.mentions_leg")}</span>
        <span className={style.colDesc}>{t("footer.legal.privacy_pol")}</span>
        <span className={style.colDesc}>{t("footer.legal.terms_of_ser")}</span>
      </div>
      {/* <div className={style.columnContainer}>
        <span className={style.colTitle}>{t("footer.support.title")}</span>
        <span className={style.colDesc}>{t("footer.support.help")}</span>
        <span className={style.colDesc}>{t("footer.support.contact_us")}</span>
        <span className={style.colDesc}>{t("footer.support.faq")}</span>
      </div> */}
    </div>
  );
};

export default Footer;
