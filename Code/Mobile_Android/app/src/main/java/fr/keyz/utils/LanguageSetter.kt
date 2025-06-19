package fr.keyz.utils

import androidx.datastore.core.DataStore
import androidx.datastore.preferences.core.Preferences
import androidx.datastore.preferences.core.edit
import androidx.datastore.preferences.core.stringPreferencesKey
import kotlinx.coroutines.flow.first
import kotlinx.coroutines.flow.map

class LanguageSetter(
    private val dataStore: DataStore<Preferences>,
) {
    suspend fun setLanguage(language: String) {
        dataStore.edit {
            it[CURRENT_LANGUAGE] = language
        }
    }
    suspend fun getLanguage(): String {
        return dataStore.data.map { it[CURRENT_LANGUAGE] ?: "en" }.first()
    }
    companion object {
        val CURRENT_LANGUAGE = stringPreferencesKey("current_language")
    }
}