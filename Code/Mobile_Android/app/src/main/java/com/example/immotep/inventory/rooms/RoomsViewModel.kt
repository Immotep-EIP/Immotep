package com.example.immotep.inventory.rooms

import androidx.lifecycle.ViewModel
import androidx.lifecycle.ViewModelProvider
import com.example.immotep.inventory.Room
import com.example.immotep.inventory.RoomDetail
import com.example.immotep.realProperty.RealPropertyViewModel
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.asStateFlow

class RoomsViewModel(
    private val rooms: Array<Room>,
    val addRoom: (String) -> Unit,
    val removeRoom: (Int) -> Unit,
    private val editRoom: (Int, Room) -> Unit) : ViewModel() {

    val _currentlyOpenRoomIndex = MutableStateFlow<Int?>(null)
    val currentlyOpenRoomIndex: StateFlow<Int?> = _currentlyOpenRoomIndex.asStateFlow()

    fun openRoomPanel(roomIndex: Int) {
        _currentlyOpenRoomIndex.value = roomIndex
    }

    fun closeRoomPanel(roomIndex: Int, details: Array<RoomDetail>) {
        if (roomIndex < 0 || roomIndex >= rooms.size) return
        val roomSelected = rooms[roomIndex]
        details.copyInto(roomSelected.details)
        editRoom(roomIndex, roomSelected)
    }
}

class RoomsViewModelFactory(
    private val rooms: Array<Room>,
    private val addRoom: (String) -> Unit,
    private val removeRoom: (Int) -> Unit,
    private val editRoom: (Int, Room) -> Unit
    ) :
    ViewModelProvider.Factory {
    override fun <T : ViewModel> create(modelClass: Class<T>): T {
        if (modelClass.isAssignableFrom(RoomsViewModel::class.java)) {
            @Suppress("UNCHECKED_CAST")
            return RoomsViewModel(rooms, addRoom, removeRoom, editRoom) as T
        }
        throw IllegalArgumentException("Unknown ViewModel class")
    }
}