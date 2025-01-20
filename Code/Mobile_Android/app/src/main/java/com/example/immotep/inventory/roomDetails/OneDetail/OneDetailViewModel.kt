package com.example.immotep.inventory.roomDetails.OneDetail

import android.net.Uri
import androidx.compose.runtime.mutableStateListOf
import androidx.lifecycle.ViewModel
import com.example.immotep.apiClient.Cleaniness
import com.example.immotep.apiClient.State
import com.example.immotep.inventory.RoomDetail
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.asStateFlow

data class RoomDetailsError(
    var name: Boolean = false,
    var comment: Boolean = false,
    var status: Boolean = false,
    var picture: Boolean = false,
    var exitPicture: Boolean = false,
    var cleaniness: Boolean = false
)

class OneDetailViewModel : ViewModel() {
    private val _detail = MutableStateFlow(RoomDetail(name = "", id = ""))
    val detail = _detail.asStateFlow()
    val picture = mutableStateListOf<Uri>()
    val exitPicture = mutableStateListOf<Uri>()
    private val _errors = MutableStateFlow(RoomDetailsError())
    val errors = _errors.asStateFlow()

    fun reset(newDetail : RoomDetail?) {
        picture.clear()
        exitPicture.clear()
        if (newDetail != null) {
            _detail.value = newDetail
            picture.addAll(newDetail.pictures)
            if (newDetail.exitPictures != null) {
                exitPicture.addAll(newDetail.exitPictures)
            }
        } else {
            _detail.value = RoomDetail(name = "", id = "")
        }
        _errors.value = RoomDetailsError()
    }

    fun setName(name : String) {
        if (name.length > 50) {
            return
        }
        _detail.value = _detail.value.copy(name = name)
        _errors.value = _errors.value.copy(name = false)

    }

    fun setComment(comment : String) {
        if (comment.length > 500) {
            return
        }
        _detail.value = _detail.value.copy(comment = comment)
        _errors.value = _errors.value.copy(comment = false)
    }

    fun setCleaniness(cleaniness : Cleaniness) {
        _detail.value = _detail.value.copy(cleaniness = cleaniness)
        _errors.value = _errors.value.copy(cleaniness = false)
    }

    fun setStatus(status : State) {
        _detail.value = _detail.value.copy(status = status)
        _errors.value = _errors.value.copy(status = false)
    }

    fun addPicture(picture : Uri) {
        this.picture.add(picture)
        _errors.value = _errors.value.copy(picture = false)
    }

    fun removePicture(index : Int) {
        this.picture.removeAt(index)
    }

    fun addExitPicture(picture : Uri) {
        this.exitPicture.add(picture)
        _errors.value = _errors.value.copy(picture = false)
    }

    fun removeExitPicture(index : Int) {
        this.exitPicture.removeAt(index)
    }

    fun onConfirm(onModifyDetail : (detail : RoomDetail) -> Unit, isExit : Boolean) {
        val error = RoomDetailsError()
        if (_detail.value.name.length < 3) {
            error.name = true
        }
        if (_detail.value.comment.length < 3) {
            error.comment = true
        }
        if (_detail.value.status == State.not_set) {
            error.status = true
        }
        if (_detail.value.cleaniness == Cleaniness.not_set) {
            error.cleaniness = true
        }
        if (picture.isEmpty()) {
            error.picture = true
        }
        if (isExit && exitPicture.isEmpty()) {
            error.exitPicture = true
        }
        if (error.name || error.comment || error.status || error.picture || error.exitPicture || error.cleaniness) {
            _errors.value = error
            return
        }
        _detail.value = _detail.value.copy(
            pictures = picture.toTypedArray(),
            completed = true,
            exitPictures = if (isExit) exitPicture.toTypedArray() else null
        )
        onModifyDetail(detail.value)
        reset(null)
    }

    fun onClose(onModifyDetail : (detail : RoomDetail) -> Unit, isExit: Boolean) {
        _detail.value = _detail.value.copy(
            pictures = picture.toTypedArray(),
            exitPictures = if (isExit) exitPicture.toTypedArray() else null
        )
        onModifyDetail(_detail.value)
        reset(null)
    }
}