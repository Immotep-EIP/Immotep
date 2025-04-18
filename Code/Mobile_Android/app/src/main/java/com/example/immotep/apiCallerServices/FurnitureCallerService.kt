package com.example.immotep.apiCallerServices

import androidx.navigation.NavController
import com.example.immotep.apiClient.ApiService
import com.example.immotep.apiClient.CreateOrUpdateResponse
import com.example.immotep.inventory.RoomDetail
import com.example.immotep.inventory.State

data class FurnitureOutput(
    val id: String,
    val property_id: String,
    val room_id: String,
    val name: String,
    val quantity: Int
) {
    fun toRoomDetail() : RoomDetail {
        return RoomDetail(
            id = id,
            name = name,
            completed = false,
            comment = "",
            status = State.not_set,
            cleanliness = com.example.immotep.inventory.Cleanliness.not_set,
            pictures = arrayOf()
        )
    }
}

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
    ) : CreateOrUpdateResponse {
        try {
            return apiService.addFurniture(getBearerToken(), propertyId, roomId, furniture)
        } catch(e : Exception) {
            throw e
        }
    }
}