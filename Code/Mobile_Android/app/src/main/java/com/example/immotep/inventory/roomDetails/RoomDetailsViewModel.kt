package com.example.immotep.inventory.roomDetails

import androidx.compose.runtime.mutableStateListOf
import androidx.lifecycle.ViewModel
import androidx.lifecycle.ViewModelProvider
import androidx.lifecycle.viewModelScope
import com.example.immotep.inventory.Room
import com.example.immotep.inventory.RoomDetail
import com.example.immotep.inventory.rooms.RoomsViewModel
import com.example.immotep.realProperty.RealPropertyViewModel
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.launch

class RoomDetailsViewModel(
    private val closeRoomPanel : (room : Room) -> Unit,
    private val addDetail : suspend (roomId : String, name : String) -> String?,
    private val roomId : String
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

    fun addDetailToRoomDetailPage(name : String) {
        viewModelScope.launch {
            val detailId = addDetail(roomId, name) ?: return@launch
            val newDetail = RoomDetail(id = detailId, name = name)
            details.add(newDetail)
        }
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
}

class RoomDetailsViewModelFactory(
    private val closeRoomPanel : (room : Room) -> Unit,
    private val addDetail : suspend (roomId : String, name : String) -> String?,
    private val roomId : String
) :
    ViewModelProvider.Factory {
    override fun <T : ViewModel> create(modelClass: Class<T>): T {
        if (modelClass.isAssignableFrom(RoomDetailsViewModel::class.java)) {
            @Suppress("UNCHECKED_CAST")
            return RoomDetailsViewModel(closeRoomPanel, addDetail, roomId) as T
        }
        throw IllegalArgumentException("Unknown ViewModel class")
    }
}