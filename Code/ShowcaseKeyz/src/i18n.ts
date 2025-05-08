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
        feature1: {
          title: "État des lieux assisté par IA",
          description:
            "États des lieux précis guidés par IA avec pré-diagnostics.",
        },
        feature2: {
          title: "Générateur de documents",
          description: "Modèles légaux de quittances et baux personnalisables.",
        },
        feature3: {
          title: "Gestion documentaire",
          description: "Stockage sécurisé avec e-signature pour vos contrats.",
        },
        feature4: {
          title: "Messagerie sécurisée",
          description: "Communication cryptée avec archivage automatique.",
        },
        feature5: {
          title: "Gestion des sinistres",
          description:
            "Déclaration et suivi des sinistres avec preuves photos.",
        },
        feature6: {
          title: "Tableau de bord intégré",
          description: "Vue d'ensemble en temps réel de vos locations.",
        },
      },
      our_application: {
        title: "Notre Application",
        subtitle:
          "La plateforme tout-en-un pour propriétaires et gestionnaires immobiliers, disponible sur mobile et ordinateur",
        image_alt: "Capture d'écran de l'interface Keyz",
        showcase_title: "Révolutionnez votre gestion locative",
        showcase_text_1:
          "Keyz centralise l'ensemble du processus locatif dans une seule application : état des lieux numérique, gestion des documents, suivi des paiements et communication avec les locataires. Adapté aussi bien aux particuliers qu'aux professionnels de l'immobilier.",
        showcase_text_2:
          "Notre technologie vous permet d'automatiser les tâches répétitives, de garantir la conformité légale et d'améliorer la relation avec vos locataires - le tout avec un gain de temps considérable.",
        highlight_1: "Interface intuitive",
        highlight_2: "Données sécurisées",
        highlight_3: "Mises à jour régulières",
        highlight_4: "Outil tout-en-un",
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
        title: "Nos Tarifs",
        subtitle:
          "Des tarifs simples et transparents pour tous vos besoins immobiliers",
        period: "/mois",
        cta: "Commencer",
        contact_us: "Nous contacter",
        most_popular: "Le plus populaire",
        from: "À partir de",
        plans: {
          basic: {
            title: "Basic",
            features: {
              feature1: "Limité à 1 seul logement",
              feature2: "1 compte locataire par logement",
              feature3: "États des lieux guidés par IA",
              feature4: "5 Go de stockage par logement",
              feature5: "Toutes les fonctionnalités de notre solution",
            },
          },
          premium: {
            title: "Premium",
            features: {
              feature1: "1 logement inclus",
              feature2: "1 compte locataire par logement",
              feature3: "États des lieux guidés par IA",
              feature4: "5 Go de stockage par logement",
              feature5: "Toutes les fonctionnalités de notre solution",
            },
            additional_info: "(+2,49€ par logement supplémentaire)",
          },
          pro: {
            title: "Pro",
            features: {
              feature1: "À partir de 10 logements",
              feature2: "1 compte locataire par logement",
              feature3: "États des lieux guidés par IA",
              feature4: "8 Go de stockage par logement",
              feature5: "Toutes les fonctionnalités de notre solution",
              feature6: "Devis personnalisé sur mesure",
            },
            price_note: "Sur demande",
          },
        },
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
      legal_mentions: {
        title: "Mentions légales",
        section1: {
          title: "1. Informations sur l'éditeur",
          paragraph1: "Ce site web est édité par Keyz.",
          paragraph2: "Numéro de TVA intracommunautaire : /",
          paragraph3: "Directeur de la publication : Oscar FRANK, Co Founder",
        },
        section2: {
          title: "2. Hébergement",
          paragraph1:
            "Ce site est hébergé par LWS, société située 2, rue Jules Ferry, 88190 Golbey",
        },
        section3: {
          title: "3. Propriété intellectuelle",
          paragraph1:
            "L'ensemble des éléments figurant sur ce site (textes, images, logos, etc.) sont protégés par les lois françaises et internationales relatives à la propriété intellectuelle. Toute reproduction ou représentation totale ou partielle de ce site ou de tout ou partie des éléments y figurant est strictement interdite sans l'autorisation préalable de Keyz SAS.",
        },
        section4: {
          title: "4. Données personnelles",
          paragraph1:
            "Les informations recueillies sur ce site font l'objet d'un traitement informatique destiné à Keyz pour la gestion des clients et prospects. Conformément à la loi « informatique et libertés » du 6 janvier 1978 modifiée et au Règlement Général sur la Protection des Données (RGPD), vous disposez d'un droit d'accès, de rectification, et d'opposition aux informations qui vous concernent.",
          paragraph2:
            "Pour exercer ces droits, veuillez nous contacter à l'adresse : contact@keyz-app.fr",
        },
        section5: {
          title: "5. Contact",
          paragraph1:
            "Pour toute question concernant ces mentions légales, vous pouvez nous contacter :",
          email: "Par email : contact@keyz-app.fr",
          phone: "Par téléphone : /",
          mail: "Par courrier : /",
        },
      },
      privacy_policy: {
        title: "Politique de confidentialité",
        section1: {
          title: "1. Introduction",
          paragraph1:
            "Chez Keyz, nous prenons la protection de vos données personnelles très au sérieux. Cette politique de confidentialité explique comment nous collectons, utilisons, partageons et protégeons vos informations lorsque vous utilisez notre site web et notre application.",
          paragraph2:
            "En utilisant nos services, vous acceptez les pratiques décrites dans cette politique de confidentialité. Nous vous encourageons à la lire attentivement.",
        },
        section2: {
          title: "2. Données collectées",
          paragraph1: "Nous collectons les informations suivantes :",
          list: {
            item1:
              "Informations d'identification : nom, prénom, adresse email, numéro de téléphone",
            item2:
              "Informations relatives à vos biens immobiliers : adresse, caractéristiques, photos, documents",
            item3:
              "Informations sur les locataires : coordonnées, documents administratifs",
            item4: "Données de paiement : pour la facturation de nos services",
            item5:
              "Données d'utilisation : comment vous interagissez avec notre plateforme",
          },
        },
        section3: {
          title: "3. Utilisation des données",
          paragraph1: "Nous utilisons vos données personnelles pour :",
          list: {
            item1: "Fournir, maintenir et améliorer nos services",
            item2: "Traiter vos transactions et gérer votre compte",
            item3:
              "Vous envoyer des notifications importantes concernant nos services",
            item4: "Vous proposer un support client personnalisé",
            item5: "Améliorer et personnaliser votre expérience utilisateur",
          },
        },
        section4: {
          title: "4. Partage des données",
          paragraph1:
            "Nous ne vendons pas vos données personnelles à des tiers. Nous pouvons les partager dans les situations suivantes :",
          list: {
            item1:
              "Avec nos prestataires de services qui nous aident à fournir nos services",
            item2:
              "Pour respecter la loi, les réglementations applicables ou les procédures légales",
            item3: "Pour protéger la sécurité de nos utilisateurs et du public",
            item4:
              "Dans le cadre d'une transaction d'entreprise (fusion, acquisition, etc.)",
          },
        },
        section5: {
          title: "5. Vos droits",
          paragraph1:
            "Conformément au Règlement Général sur la Protection des Données (RGPD), vous disposez des droits suivants :",
          list: {
            item1: "Droit d'accès à vos données personnelles",
            item2: "Droit de rectification des informations inexactes",
            item3:
              "Droit à l'effacement de vos données dans certaines conditions",
            item4: "Droit à la limitation du traitement",
            item5: "Droit à la portabilité de vos données",
            item6: "Droit d'opposition au traitement",
          },
        },
        section6: {
          title: "6. Contact",
          paragraph1:
            "Pour toute question concernant cette politique de confidentialité ou pour exercer vos droits, veuillez nous contacter :",
          email: "Par email : contact@keyz-app.fr",
          mail: "Par courrier : /",
        },
      },
      demo: {
        title: "Démonstration",
        coming_soon: "Bientôt disponible",
        development_message:
          "Notre démo interactive est en cours de développement. Notre équipe travaille activement pour vous offrir une expérience immersive qui vous permettra de découvrir toutes les fonctionnalités de Keyz.",
        completed: "complété",
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
        feature1: {
          title: "AI-Assisted Property Inspection",
          description: "Precise inspections guided by AI with pre-diagnostics.",
        },
        feature2: {
          title: "Document Generator",
          description: "Customizable legal templates for leases and receipts.",
        },
        feature3: {
          title: "Document Management",
          description: "Secure storage with e-signature for contracts.",
        },
        feature4: {
          title: "Secure Messaging",
          description: "Encrypted communication with auto-archiving.",
        },
        feature5: {
          title: "Damage Reporting",
          description: "Report and track damages with photo evidence.",
        },
        feature6: {
          title: "Dashboard",
          description: "Real-time overview of your rentals.",
        },
      },
      our_application: {
        title: "Our Application",
        subtitle:
          "The all-in-one platform for landlords and property managers, available on mobile and computer",
        image_alt: "Keyz interface screenshot",
        showcase_title: "Revolutionize your rental management",
        showcase_text_1:
          "Keyz centralizes the entire rental process in one app: digital inspections, document management, payment tracking, and tenant communication. Suitable for both individual owners and real estate professionals.",
        showcase_text_2:
          "Our technology helps you automate repetitive tasks, ensure legal compliance, and improve tenant relationships - all while saving significant time.",
        highlight_1: "Intuitive interface",
        highlight_2: "Secure data",
        highlight_3: "Regular updates",
        highlight_4: "All-in-one tool",
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
        title: "Our Pricing",
        subtitle:
          "Simple and transparent pricing for all your real estate needs",
        period: "/month",
        cta: "Get Started",
        contact_us: "Contact Us",
        most_popular: "Most popular",
        from: "From",
        plans: {
          basic: {
            title: "Basic",
            features: {
              feature1: "Limited to 1 property",
              feature2: "1 tenant account per property",
              feature3: "AI-guided property inspections",
              feature4: "5 GB storage per property",
              feature5: "All features of our solution",
            },
          },
          premium: {
            title: "Premium",
            features: {
              feature1: "1 property included",
              feature2: "1 tenant account per property",
              feature3: "AI-guided property inspections",
              feature4: "5 GB storage per property",
              feature5: "All features of our solution",
            },
            additional_info: "(+€2.49 per additional property)",
          },
          pro: {
            title: "Pro",
            features: {
              feature1: "Starting from 10 properties",
              feature2: "1 tenant account per property",
              feature3: "AI-guided property inspections",
              feature4: "8 GB storage per property",
              feature5: "All features of our solution",
              feature6: "Custom pricing on demand",
            },
            price_note: "Custom pricing",
          },
        },
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
      legal_mentions: {
        title: "Legal Notice",
        section1: {
          title: "1. Publisher Information",
          paragraph1: "This website is published by Keyz.",
          paragraph2: "VAT identification number: /",
          paragraph3: "Publication Director: Oscar FRANK, Co Founder",
        },
        section2: {
          title: "2. Hosting",
          paragraph1:
            "This site is hosted by LWS, a company located at 2, rue Jules Ferry, 88190 Golbey, France",
        },
        section3: {
          title: "3. Intellectual Property",
          paragraph1:
            "All elements on this site (texts, images, logos, etc.) are protected by French and international laws relating to intellectual property. Any reproduction or representation, in whole or in part, of this site or any of its elements is strictly prohibited without the prior authorization of Keyz SAS.",
        },
        section4: {
          title: "4. Personal Data",
          paragraph1:
            "The information collected on this site is subject to computer processing intended for Keyz for customer and prospect management. In accordance with the French Data Protection Act of January 6, 1978, as amended, and the General Data Protection Regulation (GDPR), you have the right to access, rectify, and object to information concerning you.",
          paragraph2:
            "To exercise these rights, please contact us at: contact@keyz-app.fr",
        },
        section5: {
          title: "5. Contact",
          paragraph1:
            "For any questions regarding this legal notice, you can contact us:",
          email: "By email: contact@keyz-app.fr",
          phone: "By phone: /",
          mail: "By mail: /",
        },
      },
      privacy_policy: {
        title: "Privacy Policy",
        section1: {
          title: "1. Introduction",
          paragraph1:
            "At Keyz, we take the protection of your personal data very seriously. This privacy policy explains how we collect, use, share, and protect your information when you use our website and application.",
          paragraph2:
            "By using our services, you accept the practices described in this privacy policy. We encourage you to read it carefully.",
        },
        section2: {
          title: "2. Data Collected",
          paragraph1: "We collect the following information:",
          list: {
            item1:
              "Identification information: name, surname, email address, phone number",
            item2:
              "Information relating to your properties: address, characteristics, photos, documents",
            item3:
              "Information about tenants: contact details, administrative documents",
            item4: "Payment data: for billing our services",
            item5: "Usage data: how you interact with our platform",
          },
        },
        section3: {
          title: "3. Use of Data",
          paragraph1: "We use your personal data to:",
          list: {
            item1: "Provide, maintain, and improve our services",
            item2: "Process your transactions and manage your account",
            item3: "Send you important notifications regarding our services",
            item4: "Offer you personalized customer support",
            item5: "Improve and personalize your user experience",
          },
        },
        section4: {
          title: "4. Data Sharing",
          paragraph1:
            "We do not sell your personal data to third parties. We may share it in the following situations:",
          list: {
            item1:
              "With our service providers who help us provide our services",
            item2:
              "To comply with laws, applicable regulations, or legal proceedings",
            item3: "To protect the safety of our users and the public",
            item4:
              "As part of a business transaction (merger, acquisition, etc.)",
          },
        },
        section5: {
          title: "5. Your Rights",
          paragraph1:
            "In accordance with the General Data Protection Regulation (GDPR), you have the following rights:",
          list: {
            item1: "Right to access your personal data",
            item2: "Right to rectify inaccurate information",
            item3: "Right to erase your data under certain conditions",
            item4: "Right to restriction of processing",
            item5: "Right to data portability",
            item6: "Right to object to processing",
          },
        },
        section6: {
          title: "6. Contact",
          paragraph1:
            "For any questions regarding this privacy policy or to exercise your rights, please contact us:",
          email: "By email: contact@keyz-app.fr",
          mail: "By mail: /",
        },
      },
      demo: {
        title: "Demo",
        coming_soon: "Coming Soon",
        development_message:
          "Our interactive demo is currently under development. Our team is actively working to provide you with an immersive experience that will allow you to discover all the features of Keyz.",
        completed: "completed",
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
        feature1: {
          title: "KI-gestützte Wohnungsübergabe",
          description: "Präzise Übergaben mit KI-Anleitung und Vorab-Checks.",
        },
        feature2: {
          title: "Dokumentengenerator",
          description: "Anpassbare Vorlagen für Mietverträge und Quittungen.",
        },
        feature3: {
          title: "Dokumentenverwaltung",
          description: "Sichere Aufbewahrung mit E-Signatur für Verträge.",
        },
        feature4: {
          title: "Sichere Kommunikation",
          description: "Verschlüsselter Nachrichtenaustausch mit Archivierung.",
        },
        feature5: {
          title: "Schadensmeldung",
          description: "Schäden melden und verfolgen mit Fotobeweis.",
        },
        feature6: {
          title: "Dashboard",
          description: "Echtzeit-Überblick über Ihre Mietobjekte.",
        },
      },
      our_application: {
        title: "Unsere Anwendung",
        subtitle:
          "Die All-in-One-Plattform für Vermieter und Immobilienverwalter, verfügbar auf Mobilgeräten und Computern",
        image_alt: "Keyz Interface-Screenshot",
        showcase_title: "Revolutionieren Sie Ihre Mietverwaltung",
        showcase_text_1:
          "Keyz bündelt den gesamten Mietprozess in einer App: digitale Übergabeprotokolle, Dokumentenverwaltung, Zahlungsverfolgung und Mieterkommunikation. Geeignet für private Eigentümer und Immobilienprofis.",
        showcase_text_2:
          "Unsere Technologie automatisiert repetitive Aufgaben, gewährleistet Rechtskonformität und verbessert die Mieterbeziehung - mit erheblicher Zeitersparnis.",
        highlight_1: "Intuitive Benutzeroberfläche",
        highlight_2: "Sichere Daten",
        highlight_3: "Regelmäßige Updates",
        highlight_4: "All-in-One-Tool",
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
        title: "Unsere Preise",
        subtitle:
          "Einfache und transparente Preise für alle Ihre Immobilienbedürfnisse",
        period: "/Monat",
        cta: "Jetzt starten",
        contact_us: "Kontaktieren Sie uns",
        most_popular: "Am beliebtesten",
        from: "Ab",
        plans: {
          basic: {
            title: "Basic",
            features: {
              feature1: "Begrenzt auf 1 Immobilie",
              feature2: "1 Mieterkonto pro Immobilie",
              feature3: "KI-gestützte Wohnungsübergaben",
              feature4: "5 GB Speicher pro Immobilie",
              feature5: "Alle Funktionen unserer Lösung",
            },
          },
          premium: {
            title: "Premium",
            features: {
              feature1: "1 Immobilie inklusive",
              feature2: "1 Mieterkonto pro Immobilie",
              feature3: "KI-gestützte Wohnungsübergaben",
              feature4: "5 GB Speicher pro Immobilie",
              feature5: "Alle Funktionen unserer Lösung",
            },
            additional_info: "(+2,49€ pro zusätzliche Immobilie)",
          },
          pro: {
            title: "Pro",
            features: {
              feature1: "Ab 10 Immobilien",
              feature2: "1 Mieterkonto pro Immobilie",
              feature3: "KI-gestützte Wohnungsübergaben",
              feature4: "8 GB Speicher pro Immobilie",
              feature5: "Alle Funktionen unserer Lösung",
              feature6: "Individuelles Angebot auf Anfrage",
            },
            price_note: "Auf Anfrage",
          },
        },
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
      legal_mentions: {
        title: "Impressum",
        section1: {
          title: "1. Herausgeberinformationen",
          paragraph1: "Diese Website wird von Keyz veröffentlicht.",
          paragraph2: "Umsatzsteuer-Identifikationsnummer: /",
          paragraph3: "Veröffentlichungsdirektor: Oscar FRANK, Mitbegründer",
        },
        section2: {
          title: "2. Hosting",
          paragraph1:
            "Diese Website wird von LWS gehostet, einem Unternehmen mit Sitz in 2, rue Jules Ferry, 88190 Golbey, Frankreich",
        },
        section3: {
          title: "3. Geistiges Eigentum",
          paragraph1:
            "Alle Elemente auf dieser Website (Texte, Bilder, Logos usw.) sind durch französische und internationale Gesetze zum geistigen Eigentum geschützt. Jede Reproduktion oder Darstellung, ganz oder teilweise, dieser Website oder eines ihrer Elemente ist ohne vorherige Genehmigung von Keyz SAS strengstens untersagt.",
        },
        section4: {
          title: "4. Persönliche Daten",
          paragraph1:
            "Die auf dieser Website gesammelten Informationen unterliegen einer computergestützten Verarbeitung, die für Keyz zur Verwaltung von Kunden und Interessenten bestimmt ist. Gemäß dem französischen Datenschutzgesetz vom 6. Januar 1978 in seiner geänderten Fassung und der Datenschutz-Grundverordnung (DSGVO) haben Sie das Recht auf Zugang, Berichtigung und Widerspruch gegen Informationen, die Sie betreffen.",
          paragraph2:
            "Um diese Rechte auszuüben, kontaktieren Sie uns bitte unter: contact@keyz-app.fr",
        },
        section5: {
          title: "5. Kontakt",
          paragraph1:
            "Bei Fragen zu diesem Impressum können Sie uns kontaktieren:",
          email: "Per E-Mail: contact@keyz-app.fr",
          phone: "Per Telefon: /",
          mail: "Per Post: /",
        },
      },
      privacy_policy: {
        title: "Datenschutzrichtlinie",
        section1: {
          title: "1. Einführung",
          paragraph1:
            "Bei Keyz nehmen wir den Schutz Ihrer persönlichen Daten sehr ernst. Diese Datenschutzrichtlinie erklärt, wie wir Ihre Informationen sammeln, verwenden, teilen und schützen, wenn Sie unsere Website und Anwendung nutzen.",
          paragraph2:
            "Durch die Nutzung unserer Dienste akzeptieren Sie die in dieser Datenschutzrichtlinie beschriebenen Praktiken. Wir empfehlen Ihnen, sie sorgfältig zu lesen.",
        },
        section2: {
          title: "2. Gesammelte Daten",
          paragraph1: "Wir sammeln folgende Informationen:",
          list: {
            item1:
              "Identifikationsinformationen: Name, Vorname, E-Mail-Adresse, Telefonnummer",
            item2:
              "Informationen zu Ihren Immobilien: Adresse, Eigenschaften, Fotos, Dokumente",
            item3:
              "Informationen über Mieter: Kontaktdaten, Verwaltungsdokumente",
            item4: "Zahlungsdaten: für die Abrechnung unserer Dienstleistungen",
            item5: "Nutzungsdaten: wie Sie mit unserer Plattform interagieren",
          },
        },
        section3: {
          title: "3. Verwendung der Daten",
          paragraph1: "Wir verwenden Ihre persönlichen Daten, um:",
          list: {
            item1:
              "Unsere Dienste bereitzustellen, zu erhalten und zu verbessern",
            item2:
              "Ihre Transaktionen zu verarbeiten und Ihr Konto zu verwalten",
            item3:
              "Ihnen wichtige Benachrichtigungen bezüglich unserer Dienste zu senden",
            item4: "Ihnen personalisierten Kundensupport anzubieten",
            item5:
              "Ihre Benutzererfahrung zu verbessern und zu personalisieren",
          },
        },
        section4: {
          title: "4. Datenweitergabe",
          paragraph1:
            "Wir verkaufen Ihre persönlichen Daten nicht an Dritte. Wir können sie in folgenden Situationen teilen:",
          list: {
            item1:
              "Mit unseren Dienstleistern, die uns bei der Bereitstellung unserer Dienste helfen",
            item2:
              "Um Gesetze, geltende Vorschriften oder rechtliche Verfahren einzuhalten",
            item3:
              "Um die Sicherheit unserer Benutzer und der Öffentlichkeit zu schützen",
            item4:
              "Im Rahmen einer Unternehmenstransaktion (Fusion, Übernahme usw.)",
          },
        },
        section5: {
          title: "5. Ihre Rechte",
          paragraph1:
            "Gemäß der Datenschutz-Grundverordnung (DSGVO) haben Sie folgende Rechte:",
          list: {
            item1: "Recht auf Zugang zu Ihren persönlichen Daten",
            item2: "Recht auf Berichtigung ungenauer Informationen",
            item3:
              "Recht auf Löschung Ihrer Daten unter bestimmten Bedingungen",
            item4: "Recht auf Einschränkung der Verarbeitung",
            item5: "Recht auf Datenübertragbarkeit",
            item6: "Recht auf Widerspruch gegen die Verarbeitung",
          },
        },
        section6: {
          title: "6. Kontakt",
          paragraph1:
            "Bei Fragen zu dieser Datenschutzrichtlinie oder zur Ausübung Ihrer Rechte kontaktieren Sie uns bitte:",
          email: "Per E-Mail: contact@keyz-app.fr",
          mail: "Per Post: /",
        },
      },
      demo: {
        title: "Demo",
        coming_soon: "Demnächst verfügbar",
        development_message:
          "Unsere interaktive Demo wird derzeit entwickelt. Unser Team arbeitet aktiv daran, Ihnen ein immersives Erlebnis zu bieten, das Ihnen alle Funktionen von Keyz zeigen wird.",
        completed: "abgeschlossen",
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
