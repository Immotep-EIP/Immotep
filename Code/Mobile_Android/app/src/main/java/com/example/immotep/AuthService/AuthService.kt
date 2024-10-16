package com.example.immotep.AuthService

import androidx.datastore.core.DataStore
import androidx.datastore.preferences.core.Preferences
import androidx.datastore.preferences.core.edit
import androidx.datastore.preferences.core.stringPreferencesKey
import com.example.immotep.ApiClient.ApiClient
import kotlinx.coroutines.flow.Flow
import kotlinx.coroutines.flow.firstOrNull
import kotlinx.coroutines.flow.map
import retrofit2.awaitResponse

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
        println("before calling the api")
        val response =
            try {
                ApiClient.apiService.login(username = username, password = password)
            } catch (e: Exception) {
                println(e.message)
                var code = -1
                if (e.message != null) {
                    code = this.decodeRetroFitMessagesToHttpCodes(e.message.toString())
                }
                throw Exception("Failed to login," + code.toString())
            }
        println("access tok" + response.access_token)
        this.store(response.access_token, response.refresh_token)
    }

    suspend fun refreshToken(): String {
        val refreshToken = dataStore.data.map { it[REFRESH_TOKEN] }.firstOrNull()
        if (refreshToken == null) {
            throw IllegalArgumentException("no refresh token stored")
        }
        val response = ApiClient.apiService.refreshToken(refreshToken = refreshToken).awaitResponse()
        if (!response.isSuccessful) {
            throw Exception("Failed to refresh token")
        }
        val body = response.body() ?: throw Exception("No body in response")
        this.store(body.access_token, body.refresh_token)
        return body.access_token
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

    private fun decodeRetroFitMessagesToHttpCodes(msg: String): Int {
        if (msg.startsWith("Failed to connect")) {
            return 500
        }
        if (!msg.startsWith("HTTP ")) {
            return -1
        }
        val splitedMessage = msg.split(' ')
        if (splitedMessage.size < 3) {
            return -1
        }
        try {
            val code = splitedMessage[1].toInt()
            return code
        } catch (e: Exception) {
            return -1
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
