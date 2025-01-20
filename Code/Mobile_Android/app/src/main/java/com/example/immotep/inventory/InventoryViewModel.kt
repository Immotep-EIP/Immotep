package com.example.immotep.inventory

import android.net.Uri
import androidx.compose.runtime.mutableStateListOf
import androidx.lifecycle.ViewModel
import androidx.lifecycle.ViewModelProvider
import androidx.lifecycle.viewModelScope
import androidx.navigation.NavController
import com.example.immotep.apiClient.AddRoomInput
import com.example.immotep.apiClient.ApiClient
import com.example.immotep.apiClient.Cleaniness
import com.example.immotep.apiClient.FurnitureInput
import com.example.immotep.apiClient.InventoryReportFurniture
import com.example.immotep.apiClient.InventoryReportInput
import com.example.immotep.apiClient.InventoryReportRoom
import com.example.immotep.apiClient.State
import com.example.immotep.authService.AuthService
import com.example.immotep.login.dataStore
import com.example.immotep.utils.Base64Utils
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.launch
import java.util.Optional
import java.util.Vector

data class RoomDetail(
    var id : String,
    var name : String,
    var completed : Boolean = false,
    var comment : String = "",
    var status : State = State.not_set,
    var cleaniness : Cleaniness = Cleaniness.not_set,
    val pictures : Array<Uri> = arrayOf(),
    val exitPictures : Array<Uri>? = null
)

data class Room (
    var id : String,
    val name : String,
    val description : String = "",
    var details : Array<RoomDetail> = arrayOf()
)

enum class InventoryOpenValues {
    ENTRY, EXIT, CLOSED
}

class InventoryViewModel(
    private val navController: NavController,
    private val propertyId : String,
) : ViewModel() {
    private val _inventoryOpen = MutableStateFlow(InventoryOpenValues.CLOSED)
    private var baseRooms: Array<Room> = arrayOf()
    val inventoryOpen = _inventoryOpen.asStateFlow()

    val rooms = mutableStateListOf<Room>()

    private suspend fun getBearerToken() : String? {
        val authService = AuthService(navController.context.dataStore)
        val bearerToken : String? = try {
            authService.getBearerToken()
        } catch (e : Exception) {
            e.printStackTrace()
            authService.onLogout(navController)
            null
        }
        return bearerToken
    }

    fun setInventoryOpen(value: InventoryOpenValues) {
        _inventoryOpen.value = value
    }

    fun getRooms() : Array<Room> {
        return rooms.toTypedArray()
    }

    suspend fun addRoom(name: String) : String? {
        val bearerToken = getBearerToken() ?: return null
        try {
            val createdRoom = ApiClient.apiService.addRoom(
                bearerToken,
                propertyId,
                AddRoomInput(name = name)
            )
            val room = Room(id = createdRoom.id, name = name)
            rooms.add(room)
            return createdRoom.id
        } catch (e: Exception) {
            println("Impossible to add a room ${e.message}")
            e.printStackTrace()
            return null
        }
    }

    suspend fun addFurniture(roomId: String, name: String) : String? {
        val bearerToken = getBearerToken() ?: return null
        val createdFurniture = ApiClient.apiService.addFurniture(
            bearerToken,
            propertyId,
            roomId,
            FurnitureInput(name, 1)
        )
        return createdFurniture.id
    }

    fun removeRoom(roomId: String) {
        val roomIndex = rooms.indexOf(rooms.find { it.id == roomId })
        if (roomIndex < 0 || roomIndex >= rooms.size) return
        rooms.removeAt(roomIndex)
    }

    fun editRoom(roomId: String, room: Room) {
        val roomIndex = rooms.indexOf(rooms.find { it.id == roomId })
        if (roomIndex < 0 || roomIndex >= rooms.size) return
        rooms[roomIndex] = room
    }

    fun onClose() {
        rooms.clear()
        rooms.addAll(this.baseRooms)
    }

    fun closeInventory() {
        rooms.forEach {
            it.details.forEach { detail ->
                detail.completed = false
            }
        }
        baseRooms = rooms.toTypedArray()
    }

    fun getBaseRooms() {
        viewModelScope.launch {
            val bearerToken = getBearerToken() ?: return@launch
            rooms.clear()
            try {
                val rooms = ApiClient.apiService.getAllRooms(bearerToken, propertyId)
                val newRooms = mutableListOf<Room>()
                rooms.forEach {
                    val roomsDetails = ApiClient.apiService.getAllFurnitures(bearerToken, propertyId, it.id)
                    val room = Room(it.id, it.name, "", roomsDetails.map { detail -> RoomDetail(
                        id = detail.id,
                        name = detail.name,
                    )}.toTypedArray())
                    newRooms.add(room)
                }
                this@InventoryViewModel.rooms.addAll(newRooms)
                baseRooms = newRooms.toTypedArray()
            } catch (e : Exception) {
                println("Error during get base rooms ${e.message}")
                e.printStackTrace()
            }
        }
    }

    private fun roomsToInventoryReport(openValue: InventoryOpenValues) : InventoryReportInput {
        val inventoryReportInput = InventoryReportInput(
            type = if (openValue === InventoryOpenValues.ENTRY) "start" else "end",
            rooms = Vector()
        )
        val base64Utils = Base64Utils(Uri.EMPTY)
        rooms.forEach { room ->
            val tmpRoom = InventoryReportRoom(
                id = room.id,
                state = room.details[0].status,
                cleanliness = room.details[0].cleaniness,
                note = room.details[0].comment,
                pictures = Vector(),
                furnitures = Vector()
            )
            var addedPicture = false
            room.details.forEach { detail ->
                val tmpFurniturePictures = Vector<String>()
                detail.pictures.forEach { uri ->
                    base64Utils.setFileUri(uri)
                    val encodedPicture = base64Utils.encodeImageToBase64(navController.context)
                    println("picture : $encodedPicture")
                    if (!addedPicture) {
                        tmpRoom.pictures.add(encodedPicture)
                        addedPicture = true
                    }
                    tmpFurniturePictures.add(encodedPicture)
                }
                tmpRoom.furnitures.add(
                    InventoryReportFurniture(
                        state = detail.status,
                        cleanliness = detail.cleaniness,
                        id = detail.id,
                        note = detail.comment,
                        pictures = tmpFurniturePictures
                    )
                )
            }
            inventoryReportInput.rooms.add(tmpRoom)
        }
        return inventoryReportInput
    }

    private fun checkIfAllAreCompleted() : Boolean {
        rooms.forEach { room ->
            if (room.details.isEmpty()) {
                return false
            }
            room.details.forEach { detail ->
                if (!detail.completed) {
                    return false
                }
            }
        }
        return true
    }

    fun sendInventory() : Boolean {
        if (!checkIfAllAreCompleted()) return false
        viewModelScope.launch {
            val bearerToken = getBearerToken() ?: return@launch
            try {
                val inventoryReport = roomsToInventoryReport(inventoryOpen.value)
                ApiClient.apiService.inventoryReport(bearerToken, propertyId, inventoryReport)
                closeInventory()
                setInventoryOpen(InventoryOpenValues.CLOSED)
                return@launch
            } catch (e : Exception) {
                println("Error sending inventory ${e.message}")
                e.printStackTrace()
            }
        }
        return true
    }
}

class InventoryViewModelFactory(
    private val navController: NavController,
    private val propertyId : String,
) :
    ViewModelProvider.Factory {
    override fun <T : ViewModel> create(modelClass: Class<T>): T {
        if (modelClass.isAssignableFrom(InventoryViewModel::class.java)) {
            @Suppress("UNCHECKED_CAST")
            return InventoryViewModel(navController, propertyId) as T
        }
        throw IllegalArgumentException("Unknown ViewModel class")
    }
}