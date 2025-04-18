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

    suspend fun summarize(input: AiCallInput, propertyId : String) : AiCallOutput {
        return handleRetrofitExceptions {
            apiService.aiSummarize(this.getBearerToken(), propertyId, input)
        }
    }

    suspend fun compare(input: AiCallInput, propertyId: String, oldReportId : String) : AiCallOutput {
       return handleRetrofitExceptions {
           apiService.aiCompare(this.getBearerToken(), propertyId, oldReportId, input)
       }
    }
}