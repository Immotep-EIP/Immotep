package com.example.immotep.inventory.roomDetails

import androidx.lifecycle.ViewModel
import androidx.lifecycle.ViewModelProvider
import com.example.immotep.inventory.Room
import com.example.immotep.inventory.RoomDetail
import com.example.immotep.inventory.rooms.RoomsViewModel
import com.example.immotep.realProperty.RealPropertyViewModel

class RoomDetailsViewModel(
    private val closeRoomPanel : (roomIndex: Int, details: Array<RoomDetail>) -> Unit,
    private val baseDetails: Array<RoomDetail>,
    private val roomIndex: Int
) {

    val details = mutableListOf<RoomDetail>()

    init {
        details.clear()
        details.addAll(baseDetails)
    }

    fun onClose() {
        closeRoomPanel(roomIndex, details.toTypedArray())
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
    private val baseDetails: Array<RoomDetail>,
    private val roomIndex: Int
) :
    ViewModelProvider.Factory {
    override fun <T : ViewModel> create(modelClass: Class<T>): T {
        if (modelClass.isAssignableFrom(RoomDetailsViewModel::class.java)) {
            @Suppress("UNCHECKED_CAST")
            return RoomDetailsViewModel(closeRoomPanel, baseDetails, roomIndex) as T
        }
        throw IllegalArgumentException("Unknown ViewModel class")
    }
}