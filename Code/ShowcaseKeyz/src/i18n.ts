import i18n from "i18next";
import { initReactI18next } from "react-i18next";

const resources = {
  fr: {
    translation: {
      topbar: {
        home: "Accueil",
        features: "Nos fonctionnalités",
        our_app: "Notre application",
        pricing: "Nos abonnements",
        contact_us: "Nous contacter",
        demo: "Démo",
      },
      home: {
        title: "Keyz",
        description:
          "La gestion locative simplifiée pour les propriétaires et les locataires.",
        button: "En savoir plus",
      },
      features: {
        title: "Nos fonctionnalités",
      },
      our_application: {
        title: "Notre application",
      },
      pricing: {
        title: "Nos abonnements",
        subtitle: "Choisissez l'abonnement qui vous convient le mieux",
      },
      contact_us: {
        title: "Nous contacter",
        firstname: "Prénom",
        lastname: "Nom",
        your_email: "Adresse e-mail",
        message: "Votre message",
        send_message: "Envoyer le message",
      },
      footer: {
        title: "Keyz",
        description:
          "La gestion locative simplifiée pour les propriétaires et les locataires.",
        pages: {
          title: "Navigation",
          home: "Accueil",
          features: "Nos fonctionnalités",
          our_app: "Notre application",
          pricing: "Nos abonnements",
          contact_us: "Nous contacter",
        },
        legal: {
          title: "Légal",
          mentions_leg: "Mentions légales",
          privacy_pol: "Politique de confidentialité",
          terms_of_ser: "CGU",
        },
        support: {
          title: "Support",
          help: "Aide",
          contact_us: "Nous contacter",
          faq: "FAQ",
        },
      },
    },
  },
  en: {
    translation: {
      topbar: {
        home: "Home",
        features: "Features",
        our_app: "Our App",
        pricing: "Pricing",
        contact_us: "Contact Us",
        demo: "Demo",
      },
      home: {
        title: "Keyz",
        description: "Simplified rental management for owners and tenants.",
        button: "Find out more",
      },
      features: {
        title: "Features",
      },
      our_application: {
        title: "Our Application",
      },
      pricing: {
        title: "Pricing",
        subtitle: "Choose the plan that suits you best",
      },
      contact_us: {
        title: "Contact Us",
        firstname: "First Name",
        lastname: "Last Name",
        your_email: "Your Email",
        message: "Your Message",
        send_message: "Send Message",
      },
      footer: {
        title: "Keyz",
        description: "Simplified rental management for owners and tenants.",
        pages: {
          title: "Navigation",
          home: "Home",
          features: "Features",
          our_app: "Our App",
          pricing: "Pricing",
          contact_us: "Contact Us",
        },
        legal: {
          title: "Legal",
          mentions_leg: "Legal Notice",
          privacy_pol: "Privacy Policy",
          terms_of_ser: "Terms of Service",
        },
        support: {
          title: "Support",
          help: "Help",
          contact_us: "Contact Us",
          faq: "FAQ",
        },
      },
    },
  },
  de: {
    translation: {
      topbar: {
        home: "Startseite",
        features: "Funktionen",
        our_app: "Unsere App",
        pricing: "Preise",
        contact_us: "Kontaktiere uns",
        demo: "Demo",
      },
      home: {
        title: "Keyz",
        description: "Vereinfachte Mietverwaltung für Eigentümer und Mieter.",
        button: "Mehr erfahren",
      },
      features: {
        title: "Funktionen",
      },
      our_application: {
        title: "Unsere Anwendung",
      },
      pricing: {
        title: "Preise",
        subtitle: "Wählen Sie den Plan, der am besten zu Ihnen passt",
      },
      contact_us: {
        title: "Kontaktiere uns",
        firstname: "Vorname",
        lastname: "Nachname",
        your_email: "Ihre E-Mail",
        message: "Ihre Nachricht",
        send_message: "Nachricht senden",
      },
      footer: {
        title: "Keyz",
        description: "Vereinfachte Mietverwaltung für Eigentümer und Mieter.",
        pages: {
          title: "Navigation",
          home: "Startseite",
          features: "Funktionen",
          our_app: "Unsere App",
          pricing: "Preise",
          contact_us: "Kontaktiere uns",
        },
        legal: {
          title: "Rechtliches",
          mentions_leg: "Impressum",
          privacy_pol: "Datenschutzrichtlinie",
          terms_of_ser: "Nutzungsbedingungen",
        },
        support: {
          title: "Support",
          help: "Hilfe",
          contact_us: "Kontaktiere uns",
          faq: "FAQ",
        },
      },
    },
  },
};

i18n.use(initReactI18next).init({
  resources,
  lng: "fr",
  fallbackLng: "fr",
  interpolation: {
    escapeValue: false,
  },
});

export default i18n;
