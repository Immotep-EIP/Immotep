package com.example.immotep.inventory.rooms

import androidx.compose.runtime.mutableStateListOf
import androidx.lifecycle.ViewModel
import androidx.lifecycle.ViewModelProvider
import androidx.lifecycle.viewModelScope
import com.example.immotep.inventory.Room
import com.example.immotep.inventory.RoomDetail
import com.example.immotep.realProperty.RealPropertyViewModel
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.launch

class RoomsViewModel(
    private val getRooms: () -> Array<Room>,
    private val addRoom: suspend (String) -> String?,
    private val removeRoom: (String) -> Unit,
    private val closeInventory: () -> Unit,
    private val editRoom: (Room) -> Unit,
    private val confirmInventory: () -> Boolean
    ) : ViewModel() {

    private val _currentlyOpenRoom = MutableStateFlow<Room?>(null)
    val currentlyOpenRoom: StateFlow<Room?> = _currentlyOpenRoom.asStateFlow()
    private val _showNotCompletedRooms = MutableStateFlow(false)
    val showNotCompletedRooms: StateFlow<Boolean> = _showNotCompletedRooms.asStateFlow()
    val allRooms = mutableStateListOf<Room>()

    fun handleBaseRooms() {
        allRooms.clear()
        allRooms.addAll(getRooms())
    }

    fun onClose() {
        allRooms.clear()
        closeInventory()
    }

    fun onConfirmInventory() {
        if (!confirmInventory()) {
            _showNotCompletedRooms.value = true
            return
        }
        this.onClose()
    }

    fun addARoom(name: String) {
        viewModelScope.launch {
            val roomId = addRoom(name) ?: return@launch
            val room = Room(id = roomId, name = name)
            allRooms.add(room)
        }
    }

    fun openRoomPanel(room : Room) {
        _currentlyOpenRoom.value = room
    }

    fun closeRoomPanel(updatedRoom: Room) {
        val roomIndex = allRooms.indexOf(allRooms.find { it.id == updatedRoom.id })
        if (roomIndex < 0 || roomIndex >= allRooms.size) return
        editRoom(updatedRoom)
        allRooms[roomIndex] = updatedRoom
        _currentlyOpenRoom.value = null
    }
}

class RoomsViewModelFactory(
    private val getRooms: () -> Array<Room>,
    private val addRoom: suspend (String) -> String?,
    private val removeRoom: (String) -> Unit,
    private val editRoom: (Room) -> Unit,
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