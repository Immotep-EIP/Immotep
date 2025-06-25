# Translation Management with i18n

---

**Purpose:**  
This document outlines the translation management system implemented in the Keyz web application. It provides guidelines for managing multilingual content, adding new translations, and maintaining consistency across the application's interface.

---

## Overview

Our application uses **i18n** ([i18next Documentation](https://www.i18next.com/)) to manage translations and provide multilingual support. Texts are stored in JSON files for each supported language. Currently, the application supports:

- **French** (fr)
- **English** (en)

We follow the **snake_case** naming convention for all translation keys to ensure consistency and readability.

---

## Key Features

1. **Language Resources**  
   Translations are stored in two JSON files:

   - `en.json` for English
   - `fr.json` for French

2. **Default Language**  
   The application defaults to **French (fr)** if no language is set in the browser's localStorage.

3. **Language Persistence**  
   The selected language is saved in **localStorage** to retain user preferences across sessions.

4. **Fallback Language**  
   If a key is missing in the selected language, the application falls back to **English (en)**.

5. **Dynamic Language Switching**  
   The application updates its UI dynamically when the language is changed.
