"use client";

import { useState, useEffect } from "react";
import { Menu, X } from "lucide-react";
import styles from "./Navbar.module.css";
import { useTranslation } from "react-i18next";
import LanguageSelector from "./LanguageSelector";
import NavigationEnum from "../../enums/NavigationEnum";

export default function Navbar() {
  const { t } = useTranslation();
  const [isMenuOpen, setIsMenuOpen] = useState(false);
  const [scrolled, setScrolled] = useState(false);

  useEffect(() => {
    const path = window.location.pathname;
    if (
      path === NavigationEnum.LEGAL_MENTIONS ||
      path === NavigationEnum.PRIVACY_POLICY
    ) {
      setScrolled(true);
    }
  }, [scrolled]);

  useEffect(() => {
    const handleScroll = () => {
      if (window.scrollY > 50) {
        setScrolled(true);
      } else {
        setScrolled(false);
      }
    };

    window.addEventListener("scroll", handleScroll);
    return () => window.removeEventListener("scroll", handleScroll);
  }, []);

  const toggleMenu = () => {
    setIsMenuOpen(!isMenuOpen);
  };

  return (
    <nav className={`${styles.navbar} ${scrolled ? styles.scrolled : ""}`}>
      <div className={styles.container}>
        <a
          href="/#home"
          className={scrolled ? styles.logoScrolled : styles.logo}
        >
          Keyz
        </a>

        {/* Desktop Navigation */}
        <div className={styles.desktopNav}>
          <a
            href="/#home"
            className={scrolled ? styles.navLinkScrolled : styles.navLink}
          >
            {t("topbar.home")}
          </a>
          <a
            href="/#features"
            className={scrolled ? styles.navLinkScrolled : styles.navLink}
          >
            {t("topbar.features")}
          </a>
          <a
            href="/#application"
            className={scrolled ? styles.navLinkScrolled : styles.navLink}
          >
            {t("topbar.our_app")}
          </a>
          <a
            href="/#pricing"
            className={scrolled ? styles.navLinkScrolled : styles.navLink}
          >
            {t("topbar.pricing")}
          </a>
          <a
            href="/#contact-us"
            className={scrolled ? styles.navLinkScrolled : styles.navLink}
          >
            {t("topbar.contact_us")}
          </a>
          <a href="/demo" className={styles.demoButton}>
            {t("topbar.demo")}
          </a>
          <LanguageSelector />
        </div>

        {/* Mobile menu button */}
        <div className={styles.mobileControls}>
          <button
            onClick={toggleMenu}
            className={styles.menuButton}
            aria-label="Toggle menu"
          >
            {isMenuOpen ? <X size={24} /> : <Menu size={24} />}
          </button>
        </div>
      </div>

      {/* Mobile Navigation */}
      <div className={`${styles.mobileNav} ${isMenuOpen ? styles.open : ""}`}>
        <a
          href="/#home"
          className={styles.mobileNavLink}
          onClick={() => setIsMenuOpen(false)}
        >
          {t("topbar.home")}
        </a>
        <a
          href="/#features"
          className={styles.mobileNavLink}
          onClick={() => setIsMenuOpen(false)}
        >
          {t("topbar.features")}
        </a>
        <a
          href="/#application"
          className={styles.mobileNavLink}
          onClick={() => setIsMenuOpen(false)}
        >
          {t("topbar.our_app")}
        </a>
        <a
          href="/#pricing"
          className={styles.mobileNavLink}
          onClick={() => setIsMenuOpen(false)}
        >
          {t("topbar.pricing")}
        </a>
        <a
          href="/#contact-us"
          className={styles.mobileNavLink}
          onClick={() => setIsMenuOpen(false)}
        >
          {t("topbar.contact_us")}
        </a>
        <a
          href="/demo"
          className={styles.mobileDemoButton}
          onClick={() => setIsMenuOpen(false)}
        >
          {t("topbar.demo")}
        </a>
        <LanguageSelector />
      </div>
    </nav>
  );
}
