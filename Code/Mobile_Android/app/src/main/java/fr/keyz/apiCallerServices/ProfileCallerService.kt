package fr.keyz.apiCallerServices

import androidx.navigation.NavController
import fr.keyz.apiClient.ApiService


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

    suspend fun getProfile() : ProfileResponse {
        return changeRetrofitExceptionByApiCallerException {
            apiService.getProfile(getBearerToken())
        }
    }

    suspend fun updateProfile(profileUpdateInput: ProfileUpdateInput) {
        return changeRetrofitExceptionByApiCallerException {
            apiService.updateProfile(getBearerToken(), profileUpdateInput)
        }
    }
}