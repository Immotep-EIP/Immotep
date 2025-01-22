package com.example.immotep.inventory

import android.content.Context
import android.net.Uri
import com.example.immotep.utils.Base64Utils
import java.util.Vector

//enums classes

enum class Cleanliness {
    not_set,
    dirty,
    medium,
    clean,
}

enum class State {
    not_set,
    broken,
    needsRepair,
    bad,
    medium,
    good,
    new
}

enum class InventoryLocationsTypes {
    room,
    furniture
}

enum class InventoryOpenValues {
    ENTRY, EXIT, CLOSED
}

data class RoomDetail(
    var id : String,
    var name : String,
    var completed : Boolean = false,
    var comment : String = "",
    var status : State = State.not_set,
    var cleanliness : Cleanliness = Cleanliness.not_set,
    val pictures : Array<Uri> = arrayOf(),
    val entryPictures : Array<String>? = null,
) {
    fun toInventoryReportFurniture(context: Context) : InventoryReportFurniture {
        val tmpFurniturePictures = Vector<String>()
        pictures.forEach { uri ->
            val encodedPicture = Base64Utils.encodeImageToBase64(uri, context)
            tmpFurniturePictures.add(encodedPicture)
        }
        return InventoryReportFurniture(
            state = status,
            name = name,
            cleanliness = cleanliness,
            id = id,
            note = comment,
            pictures = tmpFurniturePictures,
            quantity = 1
        )
    }
}

data class Room (
    var id : String,
    val name : String,
    val description : String = "",
    var details : Array<RoomDetail> = arrayOf()
) {
    fun toInventoryReportRoom(context: Context) : InventoryReportRoom {
        val tmpRoom = InventoryReportRoom(
            id = id,
            name = name,
            state = details[0].status,
            cleanliness = details[0].cleanliness,
            note = details[0].comment,
            pictures = Vector(),
            furnitures = Vector()
        )
        var addedPicture = false
        details.forEach {
            if (!addedPicture) {
                val tmpRoomDetail = it.toInventoryReportFurniture(context)
                tmpRoom.pictures.add(tmpRoomDetail.pictures[0])
                addedPicture = true
                tmpRoom.furnitures.add(tmpRoomDetail)
            } else {
                tmpRoom.furnitures.add(it.toInventoryReportFurniture(context))
            }
        }
        return tmpRoom
    }
}


data class InventoryReportOutput(
    val date: String,
    val id: String,
    val property_id: String,
    val rooms: Array<InventoryReportRoom>,
    val type: String
) {
    fun getRoomsAsRooms(empty : Boolean = false) : Array<Room> {
        val castedRooms = Array(rooms.size) {
            Room(id = "", name = "")
        }
        rooms.forEachIndexed { index, room ->
            castedRooms[index] = room.toRoom(empty = empty)
        }
        return castedRooms
    }
}


data class InventoryReportRoom(
    val id: String,
    val cleanliness: Cleanliness,
    val state: State,
    val note: String,
    val pictures: Vector<String>,
    val furnitures: Vector<InventoryReportFurniture>,
    val name: String
) {
    fun toRoom(empty : Boolean) : Room {
        val tmpRoom = Room(
            id = id,
            name = name,
            description = note,
            details = Array(furnitures.size) {
                RoomDetail(id = "", name = "")
            },
        )
        furnitures.forEachIndexed {
            index, furniture ->
            tmpRoom.details[index] = furniture.toRoomDetail(empty = empty)
        }
        return tmpRoom
    }
}

data class InventoryReportFurniture(
    val id: String,
    val cleanliness: Cleanliness,
    val note: String,
    val pictures: Vector<String>,
    val state: State,
    val name: String,
    val quantity: Int
) {
    fun toRoomDetail(empty : Boolean) : RoomDetail {
        val tmpRoomDetail = RoomDetail(
            id = id,
            name = name,
            completed = false,
            comment = if (empty) "" else note,
            status = if (empty) State.not_set else state,
            cleanliness = if (empty) Cleanliness.not_set else cleanliness,
            pictures = arrayOf(),
            entryPictures = pictures.toTypedArray()
        )
        return tmpRoomDetail
    }
}