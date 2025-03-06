package com.example.immotep.apiCallerServices

import androidx.navigation.NavController
import com.example.immotep.apiClient.AddRoomInput
import com.example.immotep.apiClient.ApiService
import com.example.immotep.inventory.Room
import com.example.immotep.inventory.RoomDetail

data class RoomOutput(
    val id : String,
    val name : String,
    val property_id : String,
)

class RoomCallerService(
    apiService: ApiService,
    navController: NavController
) : ApiCallerService(apiService, navController) {
    private val furnitureCaller = FurnitureCallerService(apiService, navController)

    suspend fun getAllRooms(propertyId: String, onError : () -> Unit) : Array<RoomOutput> {
        try {
            val rooms = apiService.getAllRooms(getBearerToken(), propertyId)
            return rooms
        } catch (e: Exception) {
            onError()
            throw e
        }
    }

    suspend fun getAllRoomsWithFurniture(
        propertyId: String,
        onError : () -> Unit,
        onErrorRoomFurniture : (String) -> Unit) : Array<Room>
    {
        val rooms = try {
            this.getAllRooms(propertyId, onError)
        } catch (e: Exception) {
            onError()
            throw e
        }
        val newRooms = mutableListOf<Room>()
        rooms.forEach {
            try {
                val roomsDetails = furnitureCaller.getFurnituresByRoomId(
                    propertyId,
                    it.id,
                    { onErrorRoomFurniture(it.name) }
                )
                val room = Room(
                    id = it.id,
                    name = it.name,
                    description = "",
                    details = roomsDetails.map { detail ->
                        RoomDetail(
                            id = detail.id,
                            name = detail.name,
                        )
                    }.toTypedArray()
                )
                newRooms.add(room)
            } catch (e: Exception) {
                println("Error during get all rooms with furniture ${e.message}")
            }
        }
        return newRooms.toTypedArray()
    }

    suspend fun addRoom(propertyId: String, room: AddRoomInput, onError : () -> Unit) : RoomOutput {
        try {
            val createdRoom = apiService.addRoom(getBearerToken(), propertyId, room)
            return createdRoom
        } catch (e: Exception) {
            onError()
            throw e
        }
    }
}