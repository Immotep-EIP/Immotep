package com.example.immotep.inventory

import android.net.Uri
import androidx.compose.runtime.mutableStateListOf
import androidx.lifecycle.ViewModel
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.asStateFlow

data class RoomDetail(
    var name : String = "",
    var completed : Boolean = false,
    var comment : String = "",
    var status : String = "",
    val pictures : Array<Uri> = arrayOf(),
    val exitPictures : Array<Uri>? = null
)

data class Room (
    val name : String = "",
    val description : String = "",
    var details : Array<RoomDetail> = arrayOf()
)

enum class InventoryOpenValues {
    ENTRY, EXIT, CLOSED
}

class InventoryViewModel : ViewModel() {
    val rooms = mutableStateListOf<Room>()
    fun addRoom(name: String) {
        val room = Room(name)
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

    fun sendInventory(openValue: InventoryOpenValues) {

    }
}