package com.example.immotep.apiCallerServices

import androidx.navigation.NavController
import com.example.immotep.apiClient.AddRoomInput
import com.example.immotep.apiClient.ApiService
import com.example.immotep.apiClient.CreateOrUpdateResponse
import com.example.immotep.inventory.Room
import com.example.immotep.inventory.RoomDetail

data class RoomOutput(
    val id : String,
    val name : String,
    val property_id : String,
) {
    fun toRoom(details : Array<RoomDetail>?) : Room {
        return Room(
            id = id,
            name = name,
            details = details ?: arrayOf()
        )
    }
}

class RoomCallerService(
    apiService: ApiService,
    navController: NavController
) : ApiCallerService(apiService, navController) {
    private val furnitureCaller = FurnitureCallerService(apiService, navController)

    suspend fun getAllRooms(propertyId: String) : Array<RoomOutput> =
        changeRetrofitExceptionByApiCallerException {
            apiService.getAllRooms(getBearerToken(), propertyId)
        }

    suspend fun getAllRoomsWithFurniture(
        propertyId: String,
        onErrorRoomFurniture : (String) -> Unit
    ) : Array<Room>
    {
        val rooms = try {
            this.getAllRooms(propertyId)
        } catch (e: ApiCallerServiceException) {
            throw e
        }
        val newRooms = mutableListOf<Room>()
        changeRetrofitExceptionByApiCallerException {
            rooms.forEach {
                try {
                    val roomsDetails = furnitureCaller.getFurnituresByRoomId(
                        propertyId,
                        it.id
                    )
                    val room =
                        it.toRoom(roomsDetails.map { roomDetail -> roomDetail.toRoomDetail() }
                            .toTypedArray())
                    newRooms.add(room)
                } catch (e: Exception) {
                    onErrorRoomFurniture(it.name)
                    println("Error during get all rooms with furniture ${e.message}")
                }
            }
        }
        return newRooms.toTypedArray()
    }

    suspend fun addRoom(propertyId: String, room: AddRoomInput) : CreateOrUpdateResponse = changeRetrofitExceptionByApiCallerException {
        apiService.addRoom(getBearerToken(), propertyId, room)
    }
}