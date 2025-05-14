package com.example.keyz.authService

import androidx.datastore.core.DataStore
import androidx.datastore.preferences.core.Preferences
import androidx.datastore.preferences.core.edit
import androidx.datastore.preferences.core.stringPreferencesKey
import androidx.navigation.NavController
import com.example.keyz.apiClient.ApiService
import com.example.keyz.components.decodeRetroFitMessagesToHttpCodes
import kotlinx.coroutines.flow.firstOrNull
import kotlinx.coroutines.flow.map
import java.time.LocalDateTime
import java.time.format.DateTimeFormatter

//input and output classes

data class LoginResponse(
    val access_token: String,
    val refresh_token: String,
    val token_type: String,
    val expires_in: Int,
    val properties: Map<String, Any>,
)

data class RegistrationInput(
    val email: String,
    val password: String,
    val firstName: String,
    val lastName: String,
)

data class RegistrationResponse(
    val id: String,
    val email: String,
    val firstname: String,
    val lastname: String,
    val role: String,
    val created_at: String,
    val updated_at: String,
)

class AuthService(
    private val dataStore: DataStore<Preferences>,
    private val apiService : ApiService
) {
    suspend fun onLogin(
        username: String,
        password: String,
    ) {
        val response =
            try {
                apiService.login(username = username, password = password)
            } catch (e: Exception) {
                val code = decodeRetroFitMessagesToHttpCodes(e)
                throw Exception("Failed to login,$code")
            }
        this.store(response.access_token, response.refresh_token, response.expires_in)
        try {
            val profile = apiService.getProfile(this.getBearerToken())
            println("profile role ${profile.role}")
            if (profile.role != "owner") {
                dataStore.edit {
                    it[IS_OWNER] = "false"
                }
            } else {
                dataStore.edit {
                    it[IS_OWNER] = "true"
                }
            }
        } catch (e : Exception) {
            e.printStackTrace()
        }
    }
    
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

    private suspend fun refreshToken() {
        val refreshToken = dataStore.data.map { it[REFRESH_TOKEN] }.firstOrNull()
        if (refreshToken == null) {
            throw IllegalArgumentException("no refresh token stored")
        }
        val response =
            try {
                apiService.refreshToken(refreshToken = refreshToken)
            } catch (e: Exception) {
                val code = decodeRetroFitMessagesToHttpCodes(e)
                throw Exception("Failed to refresh token,$code")
            }
        this.store(response.access_token, response.refresh_token, response.expires_in)
    }

    private suspend fun isAccessTokenExpired(): Boolean {
        val expirationTime = dataStore.data.map { it[EXPIRES_IN] }.firstOrNull()
        val expTime = LocalDateTime.parse(expirationTime, DateTimeFormatter.ofPattern("yyyy-MM-dd HH:mm:ss"))
        val actTime = LocalDateTime.now()
        return expirationTime == null || actTime.isAfter(expTime)
    }

    suspend fun register(registrationInput: RegistrationInput) : RegistrationResponse {
        try {
            return apiService.register(registrationInput)
        } catch (e: Exception) {
            val code = decodeRetroFitMessagesToHttpCodes(e)
            throw Exception("Failed to register,$code")
        }
    }

    suspend fun isUserOwner() : Boolean {
        val isOwner = dataStore.data
            .map { it[IS_OWNER] }
            .firstOrNull()
            ?: return true
        return isOwner == "true"
    }

    companion object {
        val ACCESS_TOKEN = stringPreferencesKey("access_token")
        val REFRESH_TOKEN = stringPreferencesKey("refresh_token")
        val EXPIRES_IN = stringPreferencesKey("expires_in")
        val IS_OWNER = stringPreferencesKey("is_owner")
    }
}

