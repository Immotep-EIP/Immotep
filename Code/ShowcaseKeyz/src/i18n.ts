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
        title: "Notre Application",
        subtitle:
          "Une solution complète pour la gestion locative, accessible depuis n'importe quel appareil",
        image_alt: "Application Keyz en action",
        showcase_title: "Simplifiez votre gestion immobilière",
        showcase_text_1:
          "Keyz transforme votre façon de gérer vos propriétés en unifiant toutes les fonctionnalités essentielles dans une interface intuitive et puissante. Que vous soyez propriétaire d'un seul bien ou gestionnaire d'un portefeuille immobilier, notre solution s'adapte à vos besoins.",
        showcase_text_2:
          "Grâce à des outils de suivi financier, de gestion locative et de communication centralisés, vous économiserez du temps et réduirez les erreurs tout en augmentant la satisfaction de vos locataires.",
        highlight_1: "Interface intuitive",
        highlight_2: "Données sécurisées",
        highlight_3: "Mises à jour régulières",
        highlight_4: "Support réactif",
        devices_title: "Disponible sur tous vos appareils",
        device_mobile: "Mobile",
        device_computer: "Ordinateur",
        device_tablet: "Tablette",
        cta_title: "Prêt à simplifier votre gestion locative ?",
        cta_text: "Contactez-nous pour en savoir plus",
        cta_button: "Commencer maintenant",
        contact_us_now: "Nous contacter dès maintenant",
      },
      pricing: {
        title: "Nos abonnements",
        subtitle: "Choisissez l'abonnement qui vous convient le mieux",
      },
      contact_us: {
        title: "Nous contacter",
        firstname: "Prénom*",
        lastname: "Nom*",
        your_email: "Adresse e-mail*",
        message: "Votre message*",
        send_message: "Envoyer le message",
        firstname_placeholder: "Votre prénom",
        lastname_placeholder: "Votre nom",
        your_email_placeholder: "Votre adresse e-mail",
        message_placeholder: "Veuillez saisir votre message",
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
        },
      },
      privacy_policy: {
        title: "Politique de confidentialité",
        content:
          "Cette politique de confidentialité décrit comment Keyz collecte, utilise et protège vos informations personnelles lorsque vous utilisez notre application. Nous nous engageons à respecter votre vie privée et à protéger vos données.",
      },
      legal_mentions: {
        title: "Mentions légales",
        content:
          "Les mentions légales de Keyz incluent des informations sur l'éditeur du site, l'hébergeur, et les conditions d'utilisation. Pour toute question, veuillez nous contacter.",
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
        subtitle:
          "A complete solution for property management, accessible from any device",
        image_alt: "Keyz application in action",
        showcase_title: "Simplify your property management",
        showcase_text_1:
          "Keyz transforms how you manage your properties by unifying all essential features in an intuitive and powerful interface. Whether you own a single property or manage a real estate portfolio, our solution adapts to your needs.",
        showcase_text_2:
          "With centralized financial tracking, rental management, and communication tools, you'll save time, reduce errors, and increase tenant satisfaction.",
        highlight_1: "Intuitive interface",
        highlight_2: "Secure data",
        highlight_3: "Regular updates",
        highlight_4: "Responsive support",
        devices_title: "Available on all your devices",
        device_mobile: "Mobile",
        device_computer: "Computer",
        device_tablet: "Tablet",
        cta_title: "Ready to simplify your property management?",
        cta_text: "Contact us to learn more",
        cta_button: "Get started now",
        contact_us_now: "Contact us now",
      },
      pricing: {
        title: "Pricing",
        subtitle: "Choose the plan that suits you best",
      },
      contact_us: {
        title: "Contact Us",
        firstname: "First Name*",
        lastname: "Last Name*",
        your_email: "Your Email*",
        message: "Your Message*",
        send_message: "Send Message",
        firstname_placeholder: "Your first name",
        lastname_placeholder: "Your last name",
        your_email_placeholder: "Your email address",
        message_placeholder: "Please enter your message",
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
        },
      },
      privacy_policy: {
        title: "Privacy Policy",
        content:
          "This privacy policy describes how Keyz collects, uses, and protects your personal information when you use our application. We are committed to respecting your privacy and protecting your data.",
      },
      legal_mentions: {
        title: "Legal Notice",
        content:
          "The legal notice of Keyz includes information about the publisher of the site, the host, and the terms of use. For any questions, please contact us.",
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
        subtitle:
          "Eine komplette Lösung für die Immobilienverwaltung, von jedem Gerät aus zugänglich",
        image_alt: "Keyz-Anwendung in Aktion",
        showcase_title: "Vereinfachen Sie Ihre Immobilienverwaltung",
        showcase_text_1:
          "Keyz verändert die Art und Weise, wie Sie Ihre Immobilien verwalten, indem es alle wesentlichen Funktionen in einer intuitiven und leistungsstarken Benutzeroberfläche vereint. Ob Sie Eigentümer einer einzelnen Immobilie sind oder ein Immobilienportfolio verwalten, unsere Lösung passt sich Ihren Bedürfnissen an.",
        showcase_text_2:
          "Mit zentralisierten Tools für Finanzverfolgung, Mietverwaltung und Kommunikation sparen Sie Zeit, reduzieren Fehler und erhöhen die Zufriedenheit Ihrer Mieter.",
        highlight_1: "Intuitive Benutzeroberfläche",
        highlight_2: "Sichere Daten",
        highlight_3: "Regelmäßige Updates",
        highlight_4: "Reaktionsschneller Support",
        devices_title: "Verfügbar auf allen Ihren Geräten",
        device_mobile: "Mobiltelefon",
        device_computer: "Computer",
        device_tablet: "Tablet",
        cta_title: "Bereit, Ihre Immobilienverwaltung zu vereinfachen?",
        cta_text: "Kontaktieren Sie uns, um mehr zu erfahren",
        cta_button: "Jetzt starten",
        contact_us_now: "Jetzt kontaktieren",
      },
      pricing: {
        title: "Preise",
        subtitle: "Wählen Sie den Plan, der am besten zu Ihnen passt",
      },
      contact_us: {
        title: "Kontaktiere uns",
        firstname: "Vorname*",
        lastname: "Nachname*",
        your_email: "Ihre E-Mail*",
        message: "Ihre Nachricht*",
        send_message: "Nachricht senden",
        firstname_placeholder: "Ihr Vorname",
        lastname_placeholder: "Ihr Nachname",
        your_email_placeholder: "Ihre E-Mail-Adresse",
        message_placeholder: "Bitte geben Sie Ihre Nachricht ein",
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
        },
      },
      privacy_policy: {
        title: "Datenschutzrichtlinie",
        content:
          "Diese Datenschutzrichtlinie beschreibt, wie Keyz Ihre persönlichen Informationen sammelt, verwendet und schützt, wenn Sie unsere Anwendung nutzen. Wir verpflichten uns, Ihre Privatsphäre zu respektieren und Ihre Daten zu schützen.",
      },
      legal_mentions: {
        title: "Impressum",
        content:
          "Das Impressum von Keyz enthält Informationen über den Herausgeber der Website, den Host und die Nutzungsbedingungen. Bei Fragen kontaktieren Sie uns bitte.",
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
