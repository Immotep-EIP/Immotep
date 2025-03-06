package com.example.immotep.apiCallerServices

import androidx.navigation.NavController
import com.example.immotep.apiClient.ApiService
import com.example.immotep.inventory.Cleanliness
import com.example.immotep.inventory.InventoryLocationsTypes
import com.example.immotep.inventory.State
import java.util.Vector

data class AiCallInput(
    val id : String,
    val pictures : Vector<String>,
    val type : InventoryLocationsTypes
)

data class AiCallOutput(
    val cleanliness: Cleanliness?,
    val note: String?,
    val state: State?
)

class AICallerService(
    apiService: ApiService,
    navController: NavController
) : ApiCallerService(apiService, navController) {

    suspend fun summarize(input: AiCallInput, propertyId : String, onError : () -> Unit) : AiCallOutput {
        try {
            return apiService.aiSummarize(this.getBearerToken(), propertyId, input)
        } catch(e : Exception) {
            onError()
            throw e
        }
    }

    suspend fun compare(input: AiCallInput, propertyId: String, oldReportId : String, onError: () -> Unit) : AiCallOutput {
        try {
           return apiService.aiCompare(this.getBearerToken(), propertyId, oldReportId, input)
        } catch(e : Exception) {
            onError()
            throw e
        }
    }
}