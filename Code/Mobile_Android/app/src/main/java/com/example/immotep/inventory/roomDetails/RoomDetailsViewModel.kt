package com.example.immotep.inventory.roomDetails

import androidx.compose.runtime.mutableStateListOf
import androidx.lifecycle.ViewModel
import androidx.lifecycle.ViewModelProvider
import com.example.immotep.inventory.Room
import com.example.immotep.inventory.RoomDetail
import com.example.immotep.inventory.rooms.RoomsViewModel
import com.example.immotep.realProperty.RealPropertyViewModel
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.asStateFlow

class RoomDetailsViewModel(
    private val closeRoomPanel : (roomIndex: Int, details: Array<RoomDetail>) -> Unit,
)  : ViewModel()  {

    val details = mutableStateListOf<RoomDetail>()
    val currentlyOpenDetailIndex = MutableStateFlow<Int?>(null)

    fun addBaseDetails(bDetails: Array<RoomDetail>) {
        details.clear()
        details.addAll(bDetails)
    }

    fun onClose(roomIndex : Int) {
        closeRoomPanel(roomIndex, details.toTypedArray())
        details.clear()
    }

    fun addDetail(name : String) {
        val newDetail = RoomDetail(name)
        details.add(newDetail)
    }

    fun removeDetail(detailIndex: Int) {
        if (detailIndex < 0 || detailIndex >= details.size) return
        details.removeAt(detailIndex)
    }

    fun onConfirmDetail(detailIndex: Int, detailValue : RoomDetail) {
        if (detailIndex < 0 || detailIndex >= details.size) return
        details[detailIndex] = detailValue
        details[detailIndex].completed = true
    }
}

class RoomDetailsViewModelFactory(
    private val closeRoomPanel : (roomIndex: Int, details: Array<RoomDetail>) -> Unit,
) :
    ViewModelProvider.Factory {
    override fun <T : ViewModel> create(modelClass: Class<T>): T {
        if (modelClass.isAssignableFrom(RoomDetailsViewModel::class.java)) {
            @Suppress("UNCHECKED_CAST")
            return RoomDetailsViewModel(closeRoomPanel) as T
        }
        throw IllegalArgumentException("Unknown ViewModel class")
    }
}