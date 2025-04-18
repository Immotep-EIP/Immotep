package com.example.immotep.apiCallerServices

import androidx.navigation.NavController
import coil.network.HttpException
import com.example.immotep.apiClient.ApiService
import com.example.immotep.apiClient.CreateOrUpdateResponse
import java.io.IOException

//tenant input data classes

data class InviteInput(
    val tenant_email: String,
    val start_date: String,
    val end_date: String,
)

//tenant output data classes

data class InviteOutput(
    val id: String,
    val property_id: String,
    val tenant_email: String,
    val start_date: String,
    val end_date: String,
    val created_at: String
)

class InviteTenantCallerService(
    apiService: ApiService,
    navController: NavController
) : ApiCallerService(apiService, navController) {

    suspend fun invite(propertyId : String, inviteInput: InviteInput) : CreateOrUpdateResponse {
        try {
            return this.apiService.inviteTenant(this.getBearerToken(), propertyId, inviteInput)
        } catch (e: Exception) {
            throw e
        }
    }

    suspend fun cancelInvite(propertyId: String) {
        try {
            this.apiService.cancelTenantInvitation(this.getBearerToken(), propertyId)
        } catch (e: Exception) {
            throw e
        }
    }
}