package com.example.immotep.inventory

import android.content.Context
import android.net.Uri
import com.example.immotep.apiClient.Cleanliness
import com.example.immotep.apiClient.State
import com.example.immotep.utils.Base64Utils
import java.util.Vector

data class RoomDetail(
    var id : String,
    var name : String,
    var completed : Boolean = false,
    var comment : String = "",
    var status : State = State.not_set,
    var cleanliness : Cleanliness = Cleanliness.not_set,
    val pictures : Array<Uri> = arrayOf(),
    val exitPictures : Array<Uri>? = null,
) {
    fun toInventoryReportFurniture(context: Context) : InventoryReportFurniture {
        val base64Utils = Base64Utils(Uri.EMPTY)
        val tmpFurniturePictures = Vector<String>()
        pictures.forEach { uri ->
            base64Utils.setFileUri(uri)
            val encodedPicture = base64Utils.encodeImageToBase64(context)
            tmpFurniturePictures.add(encodedPicture)
        }
        return InventoryReportFurniture(
            state = status,
            cleanliness = cleanliness,
            id = id,
            note = comment,
            pictures = tmpFurniturePictures
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
            } else {
                tmpRoom.furnitures.add(it.toInventoryReportFurniture(context))
            }
        }
        return tmpRoom
    }
}

enum class InventoryOpenValues {
    ENTRY, EXIT, CLOSED
}

data class InventoryReportOutput(
    val date: String,
    val id: String,
    val property_id: String,
    val rooms: Vector<InventoryReportRoom>,
    val type: String
)

data class InventoryReportFurniture(
    val id: String,
    val cleanliness: Cleanliness,
    val note: String,
    val pictures: Vector<String>,
    val state: State,
)

data class InventoryReportRoom(
    val id: String,
    val cleanliness: Cleanliness,
    val state: State,
    val note: String,
    val pictures: Vector<String>,
    val furnitures: Vector<InventoryReportFurniture>
)