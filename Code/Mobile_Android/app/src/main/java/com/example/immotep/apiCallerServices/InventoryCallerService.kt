package com.example.immotep.apiCallerServices

import androidx.navigation.NavController
import com.example.immotep.apiClient.ApiService
import com.example.immotep.inventory.InventoryReportOutput
import com.example.immotep.inventory.InventoryReportRoom
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
        return changeRetrofitExceptionByApiCallerException {
            apiService.inventoryReport(
                authHeader = getBearerToken(),
                propertyId = propertyId,
                leaseId = "current",
                inventoryReportInput = inventoryReportInput
            )
        }
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