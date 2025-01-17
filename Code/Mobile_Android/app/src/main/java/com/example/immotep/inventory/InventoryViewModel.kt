package com.example.immotep.inventory

import android.net.Uri
import androidx.compose.runtime.mutableStateListOf
import androidx.lifecycle.ViewModel
import androidx.lifecycle.ViewModelProvider
import androidx.lifecycle.viewModelScope
import androidx.navigation.NavController
import com.example.immotep.apiClient.AddRoomInput
import com.example.immotep.apiClient.ApiClient
import com.example.immotep.apiClient.FurnitureInput
import com.example.immotep.apiClient.InventoryReportFurniture
import com.example.immotep.apiClient.InventoryReportInput
import com.example.immotep.apiClient.InventoryReportRoom
import com.example.immotep.authService.AuthService
import com.example.immotep.inventory.rooms.RoomsViewModel
import com.example.immotep.login.dataStore
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.launch
import java.util.Vector

data class RoomDetail(
    var id : String? = null,
    var name : String = "",
    var completed : Boolean = false,
    var comment : String = "",
    var status : String = "",
    var cleaniness : String = "",
    val pictures : Array<Uri> = arrayOf(),
    val exitPictures : Array<Uri>? = null
)

data class Room (
    var id : String? = null,
    val name : String = "",
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
    val inventoryOpen = _inventoryOpen.asStateFlow()

    val rooms = mutableStateListOf<Room>()

    fun setInventoryOpen(value: InventoryOpenValues) {
        _inventoryOpen.value = value
    }

    fun getRooms() : Array<Room> {
        return rooms.toTypedArray()
    }

    fun addRoom(name: String) {
        val room = Room(name = name)
        rooms.add(room)
    }

    fun removeRoom(roomIndex: Int) {
        if (roomIndex < 0 || roomIndex >= rooms.size) return
        rooms.removeAt(roomIndex)
    }

    fun editRoom(roomIndex: Int, room: Room) {
        if (roomIndex < 0 || roomIndex >= rooms.size) return
        rooms[roomIndex] = room
    }

    fun onClose() {
        //rooms.clear()
    }

    fun getBaseRooms() {
        viewModelScope.launch {
            val authService = AuthService(navController.context.dataStore)
            val bearerToken : String = try {
                authService.getBearerToken()
            } catch (e : Exception) {
                e.printStackTrace()
                authService.onLogout(navController)
                ""
            }
            rooms.clear()
            try {
                val rooms = ApiClient.apiService.getAllRooms(bearerToken, propertyId)
                rooms.forEach {
                    val roomsDetails = ApiClient.apiService.getAllFurnitures(bearerToken, propertyId, it.id)
                    val room = Room(it.id, it.name, "", roomsDetails.map { detail -> RoomDetail(
                        id = detail.id,
                        name = detail.name,
                    )}.toTypedArray())
                    this@InventoryViewModel.rooms.add(room)
                }
            } catch (e : Exception) {
                e.printStackTrace()
            }
        }
    }

    private suspend fun createNewFurnitures(roomId : String, furnitures : Array<RoomDetail>, bearerToken : String) {
        try {
            furnitures.forEach {
                if (it.id == null) {
                    val createdFurniture = ApiClient.apiService.addFurniture(
                        bearerToken,
                        propertyId,
                        roomId,
                        FurnitureInput(it.name, 1)
                    )
                    it.id = createdFurniture.id
                }
            }
            furnitures.forEach {
                println("furni created ${it.id}, ${it.name} !")
            }
        } catch(e : Exception) {
            e.printStackTrace()
        }
    }
    private suspend fun createNewRooms(roomsToCheck : Array<Room>, bearerToken : String){
        try {
            roomsToCheck.forEach {
                if (it.id == null) {
                    val createdRoom = ApiClient.apiService.addRoom(
                        bearerToken,
                        propertyId,
                        AddRoomInput(it.name)
                    )
                    it.id = createdRoom.id
                }
                this.createNewFurnitures(it.id!!, it.details, bearerToken)
                it.details.forEach {
                    println("check for furni ${it.id}, ${it.name}")
                }
            }
        } catch(e : Exception) {
            e.printStackTrace()
        }
    }

    private fun roomsToInventoryReport(openValue: InventoryOpenValues) : InventoryReportInput {
        val inventoryReportInput = InventoryReportInput(
            type = if (openValue === InventoryOpenValues.ENTRY) "start" else "end",
            rooms = Vector()
        )
        rooms.forEach { room ->
            val tmpRoom = InventoryReportRoom(
                id = room.id!!,
                state = room.details[0].status,
                cleanliness = room.details[0].cleaniness,
                furnitures = Vector()
            )
            room.details.forEach { detail ->
                tmpRoom.furnitures.add(
                    InventoryReportFurniture(
                        state = detail.status,
                        cleanliness = detail.cleaniness,
                        id = detail.id!!
                    )
                )
            }
            inventoryReportInput.rooms.add(tmpRoom)
        }
        return inventoryReportInput
    }

    private fun checkIfAllAreCompleted() : Boolean {
        rooms.forEach { room ->
            room.details.forEach { detail ->
                if (!detail.completed) {
                    return false
                }
                }
        }
        return true
    }

    fun sendInventory() {
        if (!checkIfAllAreCompleted()) return
        viewModelScope.launch {
            var bearerToken = ""
            try {
                val authService = AuthService(navController.context.dataStore)
                bearerToken = authService.getBearerToken()
            } catch (e : Exception) {
                e.printStackTrace()
            }
            val roomsToSend = rooms.toTypedArray()
            createNewRooms(roomsToSend, bearerToken)
            setInventoryOpen(InventoryOpenValues.CLOSED)
        }
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