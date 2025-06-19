package fr.keyz.apiCallerServices

import androidx.navigation.NavController
import fr.keyz.apiClient.ApiService
import fr.keyz.apiClient.ArchiveInput
import fr.keyz.apiClient.CreateOrUpdateResponse
import fr.keyz.inventory.Room
import fr.keyz.inventory.RoomDetail


enum class RoomType {
    dressing,
    laundryroom,
    bedroom,
    playroom,
    bathroom,
    toilet,
    livingroom,
    diningroom,
    kitchen,
    hallway,
    balcony,
    cellar,
    garage,
    storage,
    office,
    other
}

data class AddRoomInput(
    val name : String,
    val type : RoomType
)

data class RoomOutput(
    val id : String,
    val name : String,
    val property_id : String,
    val type: String,
    val archived : Boolean
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
            if (this.isOwner()) {
                apiService.getAllRooms(getBearerToken(), propertyId)
            } else {
                apiService.getAllRoomsTenant(getBearerToken(), "current")
            }
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

    suspend fun archiveRoom(propertyId: String, roomId : String) = changeRetrofitExceptionByApiCallerException {
        apiService.archiveRoom(getBearerToken(), propertyId, roomId, ArchiveInput())
    }
}