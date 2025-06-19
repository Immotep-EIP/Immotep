package fr.keyz.apiCallerServices

import androidx.navigation.NavController
import fr.keyz.apiClient.ApiService
import fr.keyz.apiClient.CreateOrUpdateResponse
import fr.keyz.inventory.Cleanliness
import fr.keyz.inventory.RoomDetail
import fr.keyz.inventory.State

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
            cleanliness = Cleanliness.not_set,
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
    ) : Array<FurnitureOutput> {
        return changeRetrofitExceptionByApiCallerException {
            apiService.getAllFurnitures(getBearerToken(), propertyId, roomId)
        }
    }

    suspend fun addFurniture(
        propertyId: String,
        roomId: String,
        furniture: FurnitureInput,
    ) : CreateOrUpdateResponse {
        return changeRetrofitExceptionByApiCallerException {
            apiService.addFurniture(getBearerToken(), propertyId, roomId, furniture)
        }
    }
}