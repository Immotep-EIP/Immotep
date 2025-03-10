package com.example.immotep.apiCallerServices

import androidx.navigation.NavController
import com.example.immotep.apiClient.ApiService


data class ProfileResponse(
    val id: String,
    val email: String,
    val firstname: String,
    val lastname: String,
    val role: String,
    val created_at: String,
    val updated_at: String,
)

data class ProfileUpdateInput(
    val email: String,
    val firstname : String,
    val lastname: String
)

class ProfileCallerService(
    apiService: ApiService,
    navController: NavController,
) : ApiCallerService(apiService, navController) {

    suspend fun getProfile(onError : () -> Unit) : ProfileResponse {
        try {
            val profile = apiService.getProfile(getBearerToken())
            return profile
        } catch (e : Exception) {
            onError()
            throw e
        }
    }

    suspend fun updateProfile(profileUpdateInput: ProfileUpdateInput, onError : () -> Unit) {
        try {
            apiService.updateProfile(getBearerToken(), profileUpdateInput)
        } catch (e : Exception) {
            onError()
            throw e
        }
    }
}