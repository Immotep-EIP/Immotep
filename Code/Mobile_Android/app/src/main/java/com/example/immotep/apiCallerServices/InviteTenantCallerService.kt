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
        return changeRetrofitExceptionByApiCallerException {
            this.apiService.inviteTenant(this.getBearerToken(), propertyId, inviteInput)
        }
    }

    suspend fun cancelInvite(propertyId: String) {
        return changeRetrofitExceptionByApiCallerException {
            this.apiService.cancelTenantInvitation(this.getBearerToken(), propertyId)
        }
    }
}