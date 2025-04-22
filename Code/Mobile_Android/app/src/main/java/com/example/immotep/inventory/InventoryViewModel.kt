package com.example.immotep.inventory

import android.content.Context
import androidx.compose.runtime.mutableStateListOf
import androidx.lifecycle.ViewModel
import androidx.lifecycle.ViewModelProvider
import androidx.lifecycle.viewModelScope
import androidx.navigation.NavController
import com.example.immotep.apiCallerServices.AddRoomInput
import com.example.immotep.apiCallerServices.FurnitureCallerService
import com.example.immotep.apiCallerServices.FurnitureInput
import com.example.immotep.apiCallerServices.InventoryCallerService
import com.example.immotep.apiCallerServices.InventoryReportInput
import com.example.immotep.apiCallerServices.RoomCallerService
import com.example.immotep.apiCallerServices.RoomType
import com.example.immotep.apiClient.ApiService
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.launch
import kotlinx.coroutines.sync.Mutex
import java.util.Vector
import java.util.concurrent.atomic.AtomicBoolean

/**
 * class InventoryViewModel, made to be with the Inventory screen and handle al of his logic
 */

class InventoryViewModel(
    private val navController: NavController,
    apiService: ApiService
) : ViewModel() {
    /**
     * InventoryApiErrors, made for store the errors that can happen during the api calls
     */
    data class InventoryApiErrors(
        var getAllRooms : Boolean = false,
        var getLastInventoryReport : Boolean = false,
        var errorRoomName : String? = null,
        var createInventoryReport : Boolean = false,
    )

    private val inventoryApiCaller = InventoryCallerService(apiService, navController)
    private val roomApiCaller = RoomCallerService(apiService, navController)
    private val furnitureApiCaller = FurnitureCallerService(apiService, navController)

    private var _propertyId : String? = null
    private var _leaseId: String? = null
    private val _inventoryErrors = MutableStateFlow(InventoryApiErrors())

    private val rooms = mutableStateListOf<Room>()

    private var nonModifiedRooms: Vector<Room> = Vector()

    val inventoryErrors = _inventoryErrors.asStateFlow()

    fun loadInventoryFromRooms(rooms : Array<Room>) {
        this.rooms.clear()
        nonModifiedRooms.clear()
        this.rooms.addAll(rooms)
        rooms.forEach {
            nonModifiedRooms.add(it.copy())
        }
    }

    fun setPropertyIdAndLeaseId(propertyId: String, leaseId: String) {
        _propertyId = propertyId
        _leaseId = leaseId
    }

    /**
     * getRooms, made to get the current rooms of the property
     * @return Array<Room>, the rooms of the inventory
     */
    fun getRooms() : Array<Room> {
        return rooms.toTypedArray()
    }

    /**
     * addRoom, made to add a room to the property
     *
     */
    suspend fun addRoom(name: String, roomType: RoomType, onError : () -> Unit) : String? {
        if (_propertyId == null) {
            return null
        }
        try {
            val (id) = roomApiCaller.addRoom(
                _propertyId!!,
                AddRoomInput(name = name, type = roomType),
            )
            val room = Room(id = id, name = name)
            rooms.add(room)
            return id
        } catch (e: Exception) {
            onError()
            println("Impossible to add a room ${e.message}")
            e.printStackTrace()
            return null
        }
    }

    suspend fun addFurnitureCall(roomId: String, name: String, onError : () -> Unit) : String? {
        if (_propertyId == null) {
            return null
        }
        try {
            val createdFurniture = furnitureApiCaller.addFurniture(
                _propertyId!!,
                roomId,
                FurnitureInput(name, 1),
            )
            return createdFurniture.id
        } catch(e : Exception) {
            onError()
            return null
        }
    }

    fun removeRoom(roomId: String) {
        val roomIndex = rooms.indexOf(rooms.find { it.id == roomId })
        if (roomIndex < 0 || roomIndex >= rooms.size) return
        rooms.removeAt(roomIndex)
    }

    fun editRoom(room: Room) {
        val roomIndex = rooms.indexOf(rooms.find { it.id == room.id })
        if (roomIndex < 0 || roomIndex >= rooms.size) return
        rooms[roomIndex] = room
    }

    fun onClose() {
        _inventoryErrors.value = InventoryApiErrors()
        rooms.clear()
        nonModifiedRooms.clear()
        _propertyId = null
    }


    private fun roomsToInventoryReport(oldReportId: String?) : InventoryReportInput {
        val inventoryReportInput = InventoryReportInput(
            type = if (oldReportId == null) "start" else "end",
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

    fun sendInventory(oldReportId : String?, setNewValueOfInventory : (Array<Room>, reportId : String) -> Unit) : Boolean {
        if (!checkIfAllAreCompleted() || _propertyId == null || _leaseId == null) return false
        viewModelScope.launch {
            try {
                val inventoryReport = roomsToInventoryReport(oldReportId)
                val newReport = inventoryApiCaller.createInventoryReport(
                    _propertyId!!,
                    inventoryReport,
                    _leaseId!!
                )
                setNewValueOfInventory(rooms.toTypedArray(), newReport.id)
                onClose()
                return@launch
            } catch (e : Exception) {
                println("Error sending inventory ${e.message}")
                e.printStackTrace()
                _inventoryErrors.value = _inventoryErrors.value.copy(createInventoryReport = true)
            }
        }
        return true
    }
}