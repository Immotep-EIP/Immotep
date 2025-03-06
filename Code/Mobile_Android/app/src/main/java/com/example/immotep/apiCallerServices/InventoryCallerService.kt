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

class InventoryCallerService(
    apiService: ApiService,
    navController: NavController,
) : ApiCallerService(apiService, navController) {

    suspend fun createInventoryReport(
        propertyId: String,
        inventoryReportInput: InventoryReportInput,
        onError : () -> Unit
    ) {
        try {
            apiService.inventoryReport(getBearerToken(), propertyId, inventoryReportInput)
        } catch(e : Exception) {
            onError()
            return
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

    suspend fun getLastInventoryReport(propertyId: String, onError : () -> Unit) : InventoryReportOutput {
        try {
            val inventoryReport = apiService.getInventoryReportByIdOrLatest(getBearerToken(), propertyId, "latest")
            return inventoryReport
        } catch (e : Exception) {
            onError()
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