package com.example.immotep.AuthService

import androidx.datastore.core.DataStore
import androidx.datastore.preferences.core.Preferences
import androidx.datastore.preferences.core.edit
import androidx.datastore.preferences.core.stringPreferencesKey
import com.example.immotep.ApiClient.ApiClient
import com.example.immotep.components.decodeRetroFitMessagesToHttpCodes
import kotlinx.coroutines.flow.Flow
import kotlinx.coroutines.flow.firstOrNull
import kotlinx.coroutines.flow.map

class AuthService(
    private val dataStore: DataStore<Preferences>,
) {
    fun isAuthenticated(): Flow<Boolean> =
        dataStore.data.map {
            it.contains(ACCESS_TOKEN)
        }

    suspend fun onLogin(
        username: String,
        password: String,
    ) {
        val response =
            try {
                ApiClient.apiService.login(username = username, password = password)
            } catch (e: Exception) {
                val code = decodeRetroFitMessagesToHttpCodes(e)
                throw Exception("Failed to login,$code")
            }
        this.store(response.access_token, response.refresh_token)
    }

    suspend fun refreshToken(): String {
        val refreshToken = dataStore.data.map { it[REFRESH_TOKEN] }.firstOrNull()
        if (refreshToken == null) {
            throw IllegalArgumentException("no refresh token stored")
        }
        try {
            val res = ApiClient.apiService.refreshToken(refreshToken = refreshToken)
            this.store(res.access_token, res.refresh_token)
            return res.access_token
        } catch (e: Exception) {
            val code = decodeRetroFitMessagesToHttpCodes(e)
            throw Exception("Failed to refresh,$code")
        }
    }

    private suspend fun store(
        accessToken: String,
        refreshToken: String?,
    ) {
        dataStore.edit {
            it[ACCESS_TOKEN] = accessToken
        }
        if (refreshToken != null) {
            dataStore.edit {
                it[REFRESH_TOKEN] = refreshToken
            }
        }
    }

    suspend fun getToken(): String =
        dataStore.data
            .map { it[ACCESS_TOKEN] }
            .firstOrNull()
            ?: throw IllegalArgumentException("no token stored")

    suspend fun onLogout() {
        dataStore.edit {
            it.remove(ACCESS_TOKEN)
            it.remove(REFRESH_TOKEN)
        }
    }

    companion object {
        val ACCESS_TOKEN = stringPreferencesKey("access_token")
        val REFRESH_TOKEN = stringPreferencesKey("refresh_token")
    }
}
