import { initReactI18next } from 'react-i18next'

import i18n from 'i18next'

import en from './locales/en.json'
import fr from './locales/fr.json'

const resources = {
  en: { translation: en },
  fr: { translation: fr }
}

// Vérifie si une langue est déjà stockée dans le localStorage
const storedLang = localStorage.getItem('lang') || 'fr' // Par défaut, 'fr'

i18n.use(initReactI18next).init({
  resources,
  lng: storedLang, // Utilise la langue stockée ou la valeur par défaut
  fallbackLng: 'en',
  interpolation: {
    escapeValue: false
  }
})

// Écouter les changements de langue pour mettre à jour localStorage
i18n.on('languageChanged', lng => {
  localStorage.setItem('lang', lng)
})

export default i18n
