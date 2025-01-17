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
    private val editRoom: (Int, Room) -> Unit) : ViewModel() {

    val _currentlyOpenRoomIndex = MutableStateFlow<Int?>(null)
    val currentlyOpenRoomIndex: StateFlow<Int?> = _currentlyOpenRoomIndex.asStateFlow()
    val allRooms = mutableStateListOf<Room>()

    fun handleBaseRooms() {
        allRooms.clear()
        allRooms.addAll(getRooms())
    }

    fun onClose() {
        allRooms.clear()
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
    private val editRoom: (Int, Room) -> Unit
    ) :
    ViewModelProvider.Factory {
    override fun <T : ViewModel> create(modelClass: Class<T>): T {
        if (modelClass.isAssignableFrom(RoomsViewModel::class.java)) {
            @Suppress("UNCHECKED_CAST")
            return RoomsViewModel(getRooms, addRoom, removeRoom, editRoom) as T
        }
        throw IllegalArgumentException("Unknown ViewModel class")
    }
}