package com.example.immotep.inventory.rooms

import androidx.compose.runtime.mutableStateListOf
import androidx.lifecycle.ViewModel
import androidx.lifecycle.ViewModelProvider
import com.example.immotep.inventory.Room
import com.example.immotep.inventory.RoomDetail
import com.example.immotep.realProperty.RealPropertyViewModel
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.asStateFlow

class RoomsViewModel(
    private val getRooms: () -> Array<Room>,
    private val addRoom: (String) -> Unit,
    private val removeRoom: (Int) -> Unit,
    private val closeInventory: () -> Unit,
    private val editRoom: (Int, Room) -> Unit,
    private val confirmInventory: () -> Boolean
    ) : ViewModel() {

    private val _currentlyOpenRoomIndex = MutableStateFlow<Int?>(null)
    val currentlyOpenRoomIndex: StateFlow<Int?> = _currentlyOpenRoomIndex.asStateFlow()
    private val _showNotCompletedRooms = MutableStateFlow(false)
    val showNotCompletedRooms: StateFlow<Boolean> = _showNotCompletedRooms.asStateFlow()
    val allRooms = mutableStateListOf<Room>()

    fun handleBaseRooms() {
        allRooms.clear()
        allRooms.addAll(getRooms())
    }

    fun onConfirmInventory() {
        if (!confirmInventory()) {
            _showNotCompletedRooms.value = true
            return
        }
        allRooms.clear()
        closeInventory()
    }

    fun onClose() {
        allRooms.clear()
        closeInventory()
    }

    fun addARoom(name: String) {
        val room = Room(name = name)
        addRoom(room.name)
        allRooms.add(room)
    }

    fun openRoomPanel(roomIndex: Int) {
        if (roomIndex < 0 || roomIndex >= allRooms.size) return
        _currentlyOpenRoomIndex.value = roomIndex
    }

    fun closeRoomPanel(roomIndex: Int, details: Array<RoomDetail>) {
        if (roomIndex < 0 || roomIndex >= allRooms.size) return
        val roomSelected = allRooms[roomIndex]
        roomSelected.details = details
        editRoom(roomIndex, roomSelected)
        allRooms[roomIndex] = roomSelected
        _currentlyOpenRoomIndex.value = null
    }
}

class RoomsViewModelFactory(
    private val getRooms: () -> Array<Room>,
    private val addRoom: (String) -> Unit,
    private val removeRoom: (Int) -> Unit,
    private val editRoom: (Int, Room) -> Unit,
    private val closeInventory: () -> Unit,
    private val confirmInventory: () -> Boolean
    ) :
    ViewModelProvider.Factory {
    override fun <T : ViewModel> create(modelClass: Class<T>): T {
        if (modelClass.isAssignableFrom(RoomsViewModel::class.java)) {
            @Suppress("UNCHECKED_CAST")
            return RoomsViewModel(getRooms, addRoom, removeRoom, closeInventory, editRoom, confirmInventory) as T
        }
        throw IllegalArgumentException("Unknown ViewModel class")
    }
}