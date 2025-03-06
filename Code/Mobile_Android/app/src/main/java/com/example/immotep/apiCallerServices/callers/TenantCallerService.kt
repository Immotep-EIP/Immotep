package com.example.immotep.apiCallerServices.callers

import androidx.navigation.NavController
import com.example.immotep.apiCallerServices.ApiCallerService
import com.example.immotep.apiClient.ApiService

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

class TenantCallerService(
    apiService: ApiService,
    navController: NavController
) : ApiCallerService(apiService, navController) {

    suspend fun invite(propertyId : String, inviteInput: InviteInput, onError : () -> Unit) : InviteOutput {
        try {
            return this.apiService.inviteTenant(this.getBearerToken(), propertyId, inviteInput)
        } catch (e: Exception) {
            onError()
            throw e
        }
    }
}