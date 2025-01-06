package com.example.immotep.inventory

import android.net.Uri
import androidx.compose.runtime.mutableStateListOf
import androidx.lifecycle.ViewModel

data class RoomDetail(
    var name : String = "",
    var completed : Boolean = false,
    var comment : String = "",
    var status : String = "",
    val pictures : Array<Uri> = arrayOf()
)

data class Room (
    val name : String = "",
    val description : String = "",
    val details : Array<RoomDetail> = arrayOf()
)

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
}