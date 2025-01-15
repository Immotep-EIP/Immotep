package com.example.immotep.inventory.roomDetails.OneDetail

import android.net.Uri
import androidx.compose.runtime.mutableStateListOf
import androidx.lifecycle.ViewModel
import com.example.immotep.inventory.RoomDetail
import kotlinx.coroutines.flow.MutableStateFlow

data class RoomDetailsError(
    var name: Boolean = false,
    var comment: Boolean = false,
    var status: Boolean = false,
    var picture: Boolean = false
)

class OneDetailViewModel : ViewModel() {
    val detail = MutableStateFlow<RoomDetail>(RoomDetail())
    val picture = mutableStateListOf<Uri>()
    val errors = MutableStateFlow<RoomDetailsError>(RoomDetailsError())

    fun reset(newDetail : RoomDetail?) {
        if (newDetail != null) {
            detail.value = newDetail
        } else {
            detail.value = RoomDetail()
        }
    }

    fun setName(name : String) {
        if (name.length > 50) {
            return
        }
        detail.value.name = name
    }

    fun setComment(comment : String) {
        if (comment.length > 500) {
            return
        }
        detail.value.comment = comment
    }

    fun setStatus(status : String) {
        detail.value.status = status
    }

    fun addPicture(picture : Uri) {
        this.picture.add(picture)
    }

    fun removePicture(index : Int) {
        this.picture.removeAt(index)
    }

    fun onConfirm(onModifyDetail : (detailIndex : Int, detail : RoomDetail) -> Unit, index : Int, baseDetail : RoomDetail) {
        val error = RoomDetailsError()
        if (detail.value.name.length < 3) {
            error.name = true
        }
        if (detail.value.comment.length < 3) {
            error.comment = true
        }
        if (detail.value.status.length < 3) {
            error.status = true
        }
        if (detail.value.pictures.size < 1) {
            error.picture = true
        }
        if (error.name || error.comment || error.status || error.picture) {
            errors.value = error
        }
        detail.value.completed = true
        onModifyDetail(index, detail.value)
        reset(null)
    }

    fun onClose(onModifyDetail : (detailIndex : Int, detail : RoomDetail) -> Unit, index : Int) {
        onModifyDetail(index, detail.value)
        reset(null)
    }


}