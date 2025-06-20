# LanguageSetter

## Overview

`LanguageSetter` is a utility class designed to manage the app’s current language preference using Jetpack DataStore for persistent storage.

---

## Constructor

* `LanguageSetter(dataStore: DataStore<Preferences>)`
  Takes a DataStore instance managing `Preferences` for storing language settings.

---

## Functions

### `suspend fun setLanguage(language: String)`

* Saves the provided language code (e.g., `"en"`, `"fr"`) to the DataStore under the `CURRENT_LANGUAGE` key.
* Uses `edit` to update the stored preferences asynchronously.

### `suspend fun getLanguage(): String`

* Retrieves the stored language code from DataStore.
* Returns the stored value or `"en"` as a default fallback if no language is set.
* Uses Kotlin Flow operators to map and fetch the first emitted value.

---

## Companion Object

* `CURRENT_LANGUAGE`
  A `Preferences.Key<String>` used as the key to store/retrieve the current language string in DataStore.

---

## Summary

* Encapsulates language preference storage logic with asynchronous suspend functions.
* Provides a simple API for saving and reading the app’s current language setting persistently.
* Defaults gracefully to English (`"en"`) if no language preference exists.
