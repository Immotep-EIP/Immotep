package com.example.immotep.inventory

import android.net.Uri
import androidx.compose.runtime.mutableStateListOf
import androidx.lifecycle.ViewModel
import androidx.lifecycle.ViewModelProvider
import androidx.lifecycle.viewModelScope
import androidx.navigation.NavController
import com.example.immotep.apiClient.ApiClient
import com.example.immotep.apiClient.CreateRoomInput
import com.example.immotep.authService.AuthService
import com.example.immotep.inventory.rooms.RoomsViewModel
import com.example.immotep.login.dataStore
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.launch

data class RoomDetail(
    var name : String = "",
    var completed : Boolean = false,
    var comment : String = "",
    var status : String = "",
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
    val rooms = mutableStateListOf<Room>()
    fun addRoom(name: String) {
        val room = Room(name = name)
        println(room)
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
        rooms.clear()
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
                    val room = Room(it.id, it.name, "")
                    this@InventoryViewModel.rooms.add(room)
                }
            } catch (e : Exception) {
                e.printStackTrace()
            }
        }
    }

    private suspend fun createNewRooms(roomsToCheck : Array<Room>, bearerToken : String) {
        try {
            roomsToCheck.forEach {
                if (it.id != null) {
                    return
                }
                val createdRoom = ApiClient.apiService.createRoom(bearerToken, propertyId, CreateRoomInput(it.name))
                it.id = createdRoom.id
            }
            println(roomsToCheck)
        } catch(e : Exception) {
            e.printStackTrace()
        }
    }

    fun sendInventory(openValue: InventoryOpenValues) {
        viewModelScope.launch {
            var bearerToken = ""
            try {
                val authService: AuthService = AuthService(navController.context.dataStore)
                bearerToken = authService.getBearerToken()
            } catch (e : Exception) {
                e.printStackTrace()
            }
            val roomsToSend = rooms.toTypedArray()
            createNewRooms(roomsToSend, bearerToken)
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