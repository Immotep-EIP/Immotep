package com.example.immotep.apiCallerServices.callers

import androidx.navigation.NavController
import com.example.immotep.apiCallerServices.ApiCallerService
import com.example.immotep.apiClient.ApiClient
import com.example.immotep.apiClient.ApiService

data class FurnitureOutput(
    val id: String,
    val property_id: String,
    val room_id: String,
    val name: String,
    val quantity: Int
)

data class FurnitureInput(
    val name: String,
    val quantity: Int
)

class FurnitureCallerService(
    apiService: ApiService,
    navController: NavController
) : ApiCallerService(apiService, navController) {

    suspend fun getFurnituresByRoomId(
        propertyId: String,
        roomId: String,
        onError : () -> Unit
    ) : Array<FurnitureOutput> {
        try {
            return apiService.getAllFurnitures(getBearerToken(), propertyId, roomId)
        } catch(e : Exception) {
            onError()
            throw e
        }
    }

    suspend fun addFurniture(
        propertyId: String,
        roomId: String,
        furniture: FurnitureInput,
        onError : () -> Unit
    ) : FurnitureOutput {
        try {
            return apiService.addFurniture(getBearerToken(), propertyId, roomId, furniture)
        } catch(e : Exception) {
            onError()
            throw e
        }
    }
}