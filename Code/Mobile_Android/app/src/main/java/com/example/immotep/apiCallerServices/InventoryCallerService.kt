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
        try {
            return apiService.inventoryReport(getBearerToken(), propertyId, leaseId, inventoryReportInput)
        } catch(e : Exception) {
            println("Error during create inventory report ${e.message}")
            e.printStackTrace()
            throw e
        }
    }

    suspend fun getAllInventoryReports(propertyId: String, onError : () -> Unit) : Array<InventoryReportOutput> {
        try {
            val inventoryReports = apiService.getAllInventoryReports(getBearerToken(), propertyId)
            return inventoryReports
        } catch (e : Exception) {
            onError()
            throw e
        }
    }

    suspend fun getLastInventoryReport(propertyId: String) : InventoryReportOutput {
        try {
            val inventoryReport = apiService.getInventoryReportByIdOrLatest(getBearerToken(), propertyId, "latest")
            return inventoryReport
        } catch (e : Exception) {
            throw e
        }
    }

    suspend fun getInventoryReportById(propertyId: String, reportId: String, onError : () -> Unit) : InventoryReportOutput {
        try {
            val inventoryReport = apiService.getInventoryReportByIdOrLatest(
                getBearerToken(),
                propertyId,
                reportId
            )
            return inventoryReport
        } catch (e: Exception) {
            onError()
            throw e
        }
    }
}