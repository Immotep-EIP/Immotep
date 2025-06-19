package fr.keyz.apiCallerServices

import androidx.navigation.NavController
import fr.keyz.apiClient.ApiService
import fr.keyz.inventory.Cleanliness
import fr.keyz.inventory.InventoryLocationsTypes
import fr.keyz.inventory.State
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

    suspend fun summarize(input: AiCallInput, propertyId : String, leaseId : String) : AiCallOutput {
        return changeRetrofitExceptionByApiCallerException {
            apiService.aiSummarize(this.getBearerToken(), propertyId, leaseId, input)
        }
    }

    suspend fun compare(input: AiCallInput, propertyId: String, oldReportId : String, leaseId: String) : AiCallOutput {
       return changeRetrofitExceptionByApiCallerException {
           apiService.aiCompare(
               authHeader = this.getBearerToken(),
               propertyId = propertyId,
               oldReportId = oldReportId,
               leaseId = leaseId,
               summarizeInput = input
           )
       }
    }
}