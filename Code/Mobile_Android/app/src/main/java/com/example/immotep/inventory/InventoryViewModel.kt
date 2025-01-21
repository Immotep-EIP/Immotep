package com.example.immotep.inventory

import androidx.compose.runtime.mutableStateListOf
import androidx.lifecycle.ViewModel
import androidx.lifecycle.ViewModelProvider
import androidx.lifecycle.viewModelScope
import androidx.navigation.NavController
import com.example.immotep.apiClient.AddRoomInput
import com.example.immotep.apiClient.ApiClient
import com.example.immotep.apiClient.FurnitureInput
import com.example.immotep.apiClient.InventoryReportInput
import com.example.immotep.authService.AuthService
import com.example.immotep.login.dataStore
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.launch
import java.util.Vector

class InventoryViewModel(
    private val navController: NavController,
    private val propertyId : String,
) : ViewModel() {
    private val _inventoryOpen = MutableStateFlow(InventoryOpenValues.CLOSED)
    private var baseRooms: Vector<Room> = Vector()
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
        this.baseRooms.forEach {
            it.details.forEach { detail ->
                println(detail.completed)
            }
        }
        this.baseRooms.forEach {
            rooms.add(it.copy())
        }
    }

    fun closeInventory() {
        rooms.forEach {
            it.details.forEach { detail ->
                detail.completed = false
            }
        }
        baseRooms.clear()
        rooms.forEach {
            baseRooms.add(it.copy())
        }
    }

    fun getBaseRooms() {
        viewModelScope.launch {
            val bearerToken = getBearerToken() ?: return@launch
            rooms.clear()
            baseRooms.clear()
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
                newRooms.forEach {
                    baseRooms.add(it.copy())
                }
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
        rooms.forEach {
            inventoryReportInput.rooms.add(it.toInventoryReportRoom(navController.context))
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

    fun setRoomsAsLastInventory() {
        viewModelScope.launch {
            val bearerToken = getBearerToken() ?: return@launch
            try {
                val inventoryReports = ApiClient.apiService.getAllInventoryReports(bearerToken, propertyId)
            } catch (e : Exception) {
                println("Error getting last inventory ${e.message}")
                e.printStackTrace()
            }
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