package com.example.immotep.authService

import androidx.datastore.core.DataStore
import androidx.datastore.preferences.core.Preferences
import androidx.datastore.preferences.core.edit
import androidx.datastore.preferences.core.stringPreferencesKey
import androidx.navigation.NavController
import com.example.immotep.apiClient.ApiClient
import com.example.immotep.components.decodeRetroFitMessagesToHttpCodes
import kotlinx.coroutines.flow.firstOrNull
import kotlinx.coroutines.flow.map
import java.time.LocalDate
import java.time.LocalDateTime
import java.time.format.DateTimeFormatter
import java.util.Date

class AuthService(
    private val dataStore: DataStore<Preferences>,
) {
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
        this.store(response.access_token, response.refresh_token, response.expires_in)
    }

    /*
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
    */

    private suspend fun store(
        accessToken: String,
        refreshToken: String?,
        expiresIn: Int
    ) {
        val expirationTime = LocalDateTime.now().plusSeconds(expiresIn.toLong() - (5 * 60))
        val formatter = DateTimeFormatter.ofPattern("yyyy-MM-dd HH:mm:ss")
        dataStore.edit {
            it[ACCESS_TOKEN] = accessToken
            it[EXPIRES_IN] = expirationTime.format(formatter)
        }
        if (refreshToken != null) {
            dataStore.edit {
                it[REFRESH_TOKEN] = refreshToken
            }
        }
    }

    suspend fun getToken(): String {
        if (isAccessTokenExpired()) {
            refreshToken()
        }
        val token = dataStore.data
            .map { it[ACCESS_TOKEN] }
            .firstOrNull()
            ?: throw IllegalArgumentException("no token stored")
        return token
    }

    suspend fun getBearerToken(): String = "Bearer ${this.getToken()}"

    suspend fun deleteToken() {
        dataStore.edit {
            it.remove(ACCESS_TOKEN)
            it.remove(REFRESH_TOKEN)
        }
    }

    suspend fun onLogout(navController: NavController) {
        this.deleteToken()
        navController.navigate("login")
    }

    suspend fun refreshToken() {
        val refreshToken = dataStore.data.map { it[REFRESH_TOKEN] }.firstOrNull()
        if (refreshToken == null) {
            throw IllegalArgumentException("no refresh token stored")
        }
        val response =
            try {
                ApiClient.apiService.refreshToken(refreshToken = refreshToken)
            } catch (e: Exception) {
                val code = decodeRetroFitMessagesToHttpCodes(e)
                throw Exception("Failed to refresh token,$code")
            }
        this.store(response.access_token, response.refresh_token, response.expires_in)
    }

    suspend fun isAccessTokenExpired(): Boolean {
        val expirationTime = dataStore.data.map { it[EXPIRES_IN] }.firstOrNull()
        val expTime = LocalDateTime.parse(expirationTime, DateTimeFormatter.ofPattern("yyyy-MM-dd HH:mm:ss"))
        val actTime = LocalDateTime.now()
        if (expirationTime == null || actTime.isAfter(expTime)) {
            return true
        }
        return false
    }

    companion object {
        val ACCESS_TOKEN = stringPreferencesKey("access_token")
        val REFRESH_TOKEN = stringPreferencesKey("refresh_token")
        val EXPIRES_IN = stringPreferencesKey("expires_in")
    }
}

