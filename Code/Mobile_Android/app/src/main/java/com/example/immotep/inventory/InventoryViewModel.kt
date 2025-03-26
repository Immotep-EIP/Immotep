package com.example.immotep.inventory

import androidx.compose.runtime.mutableStateListOf
import androidx.lifecycle.ViewModel
import androidx.lifecycle.ViewModelProvider
import androidx.lifecycle.viewModelScope
import androidx.navigation.NavController
import com.example.immotep.apiCallerServices.FurnitureCallerService
import com.example.immotep.apiCallerServices.FurnitureInput
import com.example.immotep.apiCallerServices.InventoryCallerService
import com.example.immotep.apiCallerServices.InventoryReportInput
import com.example.immotep.apiCallerServices.RoomCallerService
import com.example.immotep.apiClient.AddRoomInput
import com.example.immotep.apiClient.ApiService
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.launch
import java.util.Vector

/**
 * class InventoryViewModel, made to be with the Inventory screen and handle al of his logic
 */

class InventoryViewModel(
    private val navController: NavController,
    private var propertyId : String,
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
    private val _inventoryOpen = MutableStateFlow(InventoryOpenValues.CLOSED)
    private val _inventoryErrors = MutableStateFlow(InventoryApiErrors())
    private val _oldReportId = MutableStateFlow<String?>(null)
    private val _cannotMakeExitInventory = MutableStateFlow(false)

    private val rooms = mutableStateListOf<Room>()
    private val lastInventoryRooms = mutableStateListOf<Room>()

    private var baseRooms: Vector<Room> = Vector()
    private var lastInventoryBaseRooms: Vector<Room> = Vector()

    val cannotMakeExitInventory = _cannotMakeExitInventory.asStateFlow()
    val inventoryOpen = _inventoryOpen.asStateFlow()
    val oldReportId = _oldReportId.asStateFlow()
    val inventoryErrors = _inventoryErrors.asStateFlow()

    /**
     * setInventoryOpen, made to set the inventory open or close
     * @param value, the value to set the inventory to
     * @return void
     */
    fun setInventoryOpen(value: InventoryOpenValues) {
        if (inventoryErrors.value.getAllRooms || inventoryErrors.value.errorRoomName != null) {
            return
        }
        if (value === InventoryOpenValues.EXIT
            &&
            (_oldReportId.value == null
                    || lastInventoryBaseRooms.isEmpty()
                    || inventoryErrors.value.getLastInventoryReport)
            ) {
            _cannotMakeExitInventory.value = true
            return
        }
        _inventoryOpen.value = value
    }

    /**
     * closeCannotMakeExitInventory, made to close the cannot make exit inventory error
     * @return void
     */
    fun closeCannotMakeExitInventory() {
        _cannotMakeExitInventory.value = false
    }

    /**
     * getRooms, made to get the current rooms of the property
     * @return Array<Room>, the rooms of the inventory
     */
    fun getRooms() : Array<Room> {
        if (inventoryOpen.value === InventoryOpenValues.EXIT) {
            return lastInventoryRooms.toTypedArray()
        }
        return rooms.toTypedArray()
    }

    /**
     * addRoom, made to add a room to the property
     *
     */
    suspend fun addRoom(name: String, onError : () -> Unit) : String? {
        try {
            val createdRoom = roomApiCaller.addRoom(
                propertyId,
                AddRoomInput(name = name),
                onError = onError
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

    suspend fun addFurnitureCall(roomId: String, name: String, onError : () -> Unit) : String? {
        try {
            val createdFurniture = furnitureApiCaller.addFurniture(
                propertyId,
                roomId,
                FurnitureInput(name, 1),
                onError
            )
            return createdFurniture.id
        } catch(e : Exception) {
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
        if (inventoryOpen.value == InventoryOpenValues.ENTRY) {
            rooms.clear()
            this.baseRooms.forEach {
                rooms.add(it.copy())
                val tmpRoom = rooms.last()
                for (i in tmpRoom.details.indices) {
                    tmpRoom.details[i] = it.details[i].copy(completed = false)
                }
            }
            return
        }
        if (inventoryOpen.value == InventoryOpenValues.EXIT) {
            lastInventoryRooms.clear()
            this.lastInventoryBaseRooms.forEach {
                lastInventoryRooms.add(it.copy())
                val tmpRoom = lastInventoryRooms.last()
                for (i in tmpRoom.details.indices) {
                    tmpRoom.details[i] = it.details[i].copy(completed = false)
                }
            }
        }
    }

    private fun closeInventory() {
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

    private suspend fun getLastInventory() {
        lastInventoryRooms.clear()
        lastInventoryBaseRooms.clear()
        try {
            val inventoryReport = inventoryApiCaller.getLastInventoryReport(
                propertyId
            ) {}
            val lastInventoryRoomsAsRooms = inventoryReport.getRoomsAsRooms(empty = true)
            _oldReportId.value = inventoryReport.id
            lastInventoryRooms.addAll(lastInventoryRoomsAsRooms)
            lastInventoryRoomsAsRooms.forEach {
                lastInventoryBaseRooms.add(it.copy())
            }
        } catch (e : Exception) {
            println("Error during get last inventory ${e.message}")
            e.printStackTrace()
        }
    }

    fun getBaseRooms(propertyId: String) {
        this.propertyId = propertyId
        _inventoryErrors.value = InventoryApiErrors()
        viewModelScope.launch {
            rooms.clear()
            baseRooms.clear()
            try {
                val newRooms = roomApiCaller.getAllRoomsWithFurniture(
                    propertyId,
                    { _inventoryErrors.value = _inventoryErrors.value.copy(getAllRooms = true) },
                    { _inventoryErrors.value = _inventoryErrors.value.copy(errorRoomName = it) }
                )
                this@InventoryViewModel.rooms.addAll(newRooms)
                newRooms.forEach {
                    baseRooms.add(it.copy())
                }
            } catch (e : Exception) {
                println("Error during get base rooms ${e.message}")
                e.printStackTrace()
            }
            getLastInventory()
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
            try {
                val inventoryReport = roomsToInventoryReport(inventoryOpen.value)
                inventoryApiCaller.createInventoryReport(
                    propertyId,
                    inventoryReport,
                    { _inventoryErrors.value = _inventoryErrors.value.copy(createInventoryReport = true) }
                )
                closeInventory()
                setInventoryOpen(InventoryOpenValues.CLOSED)
                getLastInventory()
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
    private val apiService: ApiService
) :
    ViewModelProvider.Factory {
    override fun <T : ViewModel> create(modelClass: Class<T>): T {
        if (modelClass.isAssignableFrom(InventoryViewModel::class.java)) {
            @Suppress("UNCHECKED_CAST")
            return InventoryViewModel(navController, propertyId, apiService) as T
        }
        throw IllegalArgumentException("Unknown ViewModel class")
    }
}