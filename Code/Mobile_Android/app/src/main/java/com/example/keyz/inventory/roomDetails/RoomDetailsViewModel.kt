package com.example.keyz.inventory.roomDetails

import androidx.compose.runtime.mutableStateListOf
import androidx.lifecycle.ViewModel
import com.example.keyz.inventory.Room
import com.example.keyz.inventory.RoomDetail
import kotlinx.coroutines.flow.MutableStateFlow

class RoomDetailsViewModel(
    private val closeRoomPanel : (room : Room) -> Unit,
    private val addDetail : suspend (roomId : String, name : String) -> String?,
    private val removeDetail: (String) -> Unit
)  : ViewModel()  {

    val details = mutableStateListOf<RoomDetail>()
    val currentlyOpenDetail = MutableStateFlow<RoomDetail?>(null)

    fun addBaseDetails(bDetails: Array<RoomDetail>) {
        details.clear()
        details.addAll(bDetails)
    }

    fun onClose(room: Room) {
        room.details = details.toTypedArray()
        closeRoomPanel(room)
        details.clear()
    }

    suspend fun addDetailToRoomDetailPage(name : String, roomId: String) {
        if (details.find { it.name == name } != null) throw Exception("detail_already_exists")
        val detailId = addDetail(roomId, name) ?: throw Exception("impossible_to_add_detail")
        val newDetail = RoomDetail(id = detailId, name = name, newItem = true)
        details.add(newDetail)
    }

    fun removeDetail(detail: RoomDetail) {
        val detailIndex = details.indexOfFirst { it.id == detail.id }
        if (detailIndex < 0 || detailIndex >= details.size) return
        details.removeAt(detailIndex)
    }

    fun onModifyDetail(detail : RoomDetail) {
        val detailIndex = details.indexOfFirst { it.id == detail.id }
        if (detailIndex < 0 || detailIndex >= details.size) return
        details[detailIndex] = detail
        currentlyOpenDetail.value = null
    }

    fun onOpenDetail(detail: RoomDetail) {
        currentlyOpenDetail.value = detail
    }
    //finish this
    fun handleRemoveDetail(detailId : String) {
        try {
            val detailIndex = details.indexOfFirst { it.id == detailId }
            if (detailIndex == -1) return
            removeDetail(detailId)
        } catch(e : Exception) {
            e.printStackTrace()
        }
    }
}
