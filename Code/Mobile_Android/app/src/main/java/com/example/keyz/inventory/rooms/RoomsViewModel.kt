package com.example.keyz.inventory.rooms

import androidx.compose.runtime.mutableStateListOf
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.example.keyz.apiCallerServices.RoomType
import com.example.keyz.inventory.Room
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.launch

class RoomsViewModel(
    private val getRooms: () -> Array<Room>,
    private val addRoom: suspend (String, RoomType) -> String?,
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

    suspend fun addARoom(name: String, type : RoomType) {
        if (allRooms.find { it.name == name } != null) throw Exception("room_already_exists")
        val roomId = addRoom(name, type) ?: throw Exception("impossible_to_add_room")
        val room = Room(id = roomId, name = name, newItem = true)
        allRooms.add(room)
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
