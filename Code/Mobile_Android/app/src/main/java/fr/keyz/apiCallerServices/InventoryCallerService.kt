package fr.keyz.apiCallerServices

import androidx.navigation.NavController
import fr.keyz.apiClient.ApiService
import fr.keyz.inventory.InventoryReportOutput
import fr.keyz.inventory.InventoryReportRoom
import retrofit2.HttpException
import java.util.Vector

data class InventoryReportInput(
    val type: String,
    val rooms: Vector<InventoryReportRoom>
)

data class CreatedInventoryReport(
    val date: String,
    val errors : Array<String>,
    val id: String,
    val lease_id: String,
    val pdf_data: String,
    val pdf_name: String,
    val property_id: String,
    val type: String
)



class InventoryCallerService(
    apiService: ApiService,
    navController: NavController,
) : ApiCallerService(apiService, navController) {

    suspend fun createInventoryReport(
        propertyId: String,
        inventoryReportInput: InventoryReportInput,
        leaseId : String,
    ) : CreatedInventoryReport {
            try {
                return apiService.inventoryReport(
                    authHeader = getBearerToken(),
                    propertyId = propertyId,
                    leaseId = "current",
                    inventoryReportInput = inventoryReportInput
                )
            } catch (e : HttpException) {
                println("error response inventory report ${e.message()}")
            } catch (e : Exception) {
                println("error inventory report unknown error ${e.message}")
            }
        throw ApiCallerServiceException("500")
    }

    suspend fun getAllInventoryReports(propertyId: String) : Array<InventoryReportOutput> {
        return changeRetrofitExceptionByApiCallerException{
            apiService.getAllInventoryReports(getBearerToken(), propertyId)
        }
    }

    suspend fun getLastInventoryReport(propertyId: String) : InventoryReportOutput {
        return changeRetrofitExceptionByApiCallerException {
            apiService.getInventoryReportByIdOrLatest(getBearerToken(), propertyId, "latest")
        }
    }

    suspend fun getInventoryReportById(propertyId: String, reportId: String) : InventoryReportOutput {
        return changeRetrofitExceptionByApiCallerException {
            apiService.getInventoryReportByIdOrLatest(
                getBearerToken(),
                propertyId,
                reportId
            )
        }
    }
}