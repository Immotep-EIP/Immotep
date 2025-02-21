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
    fun toRoom() : Room {
        return Room(
            id = id,
            name = name,
            description = comment,
            cleanliness = cleanliness,
            state = status,
            pictures = pictures,
            entryPictures = entryPictures,
            details = arrayOf(),
            completed = completed
        )
    }
}

data class Room (
    var id : String,
    val name : String,
    val description : String = "",
    val cleanliness: Cleanliness = Cleanliness.not_set,
    val state: State = State.not_set,
    val completed : Boolean = false,
    val pictures: Array<Uri> = arrayOf(),
    val entryPictures: Array<String>? = null,
    var details : Array<RoomDetail> = arrayOf()
) {
    fun toInventoryReportRoom(context: Context) : InventoryReportRoom {
        val tmpRoom = InventoryReportRoom(
            id = id,
            name = name,
            state = state,
            cleanliness = cleanliness,
            note = description,
            pictures = Vector(),
            furnitures = Vector()
        )
        pictures.forEach {
            val encodedPicture = Base64Utils.encodeImageToBase64(it, context)
            tmpRoom.pictures.add(encodedPicture)
        }
        details.forEach {
            tmpRoom.furnitures.add(it.toInventoryReportFurniture(context))
        }
        return tmpRoom
    }
    fun toRoomDetail() : RoomDetail {
        return RoomDetail(
            id = id,
            name = name,
            completed = completed,
            comment = description,
            status = state,
            cleanliness = cleanliness,
            pictures = pictures,
            entryPictures = entryPictures
        )
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
        val castedRooms = Vector<Room>()
        rooms.forEach { room ->
            castedRooms.add(room.toRoom(empty = empty))
        }
        return castedRooms.toTypedArray()
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
            cleanliness = cleanliness,
            state = state,
            entryPictures = if (pictures.isEmpty()) null else pictures.toTypedArray(),
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