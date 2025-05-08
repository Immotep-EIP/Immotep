package com.example.immotep.addDamageModal

import android.net.Uri
import androidx.compose.runtime.MutableState
import androidx.compose.runtime.mutableStateListOf
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import androidx.navigation.NavController
import com.example.immotep.apiCallerServices.DamageCallerService
import com.example.immotep.apiCallerServices.DamageInput
import com.example.immotep.apiCallerServices.DamagePriority
import com.example.immotep.apiCallerServices.RealPropertyCallerService
import com.example.immotep.apiCallerServices.RoomCallerService
import com.example.immotep.apiClient.ApiService
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.launch



class AddDamageModalViewModel(apiService: ApiService, navController: NavController) : ViewModel() {
    data class DamageInputError(
        var comment: Boolean = false,
        var pictures: Boolean = false,
        var room: Boolean = false
    )
    data class SimplifiedRoom(
        val id : String,
        val name : String
    )
    private val _apiCaller = DamageCallerService(apiService, navController)
    private val _roomApiCaller = RoomCallerService(apiService, navController)
    private val _form = MutableStateFlow(DamageInput())
    private val _formError = MutableStateFlow(DamageInputError())

    val form = _form.asStateFlow()
    val formError = _formError.asStateFlow()

    val pictures = mutableStateListOf<Uri>()
    val rooms = mutableStateListOf<SimplifiedRoom>()

    private fun getBaseRooms() {
        rooms.clear()
        viewModelScope.launch {
            try {
                val newRooms = _roomApiCaller.getAllRooms("current")
                rooms.addAll(newRooms.map { SimplifiedRoom(it.id, it.name) })
            } catch (e : Exception) {
                e.printStackTrace()
            }
        }
    }

    fun reset() {
        _form.value = DamageInput()
        _formError.value = DamageInputError()
        pictures.clear()
        getBaseRooms()
    }

    fun setComment(comment : String) {
        _formError.value = _formError.value.copy(comment = false)
        _form.value = _form.value.copy(comment = comment)
    }

    fun setPriority(priority: DamagePriority) {
        _form.value = _form.value.copy(priority = priority)
    }

    fun setRoomId(roomId : String) {
        _formError.value = _formError.value.copy(room = false)
        _form.value = _form.value.copy(room_id = roomId)
    }

    fun addPicture(picture : Uri) {
        _formError.value = _formError.value.copy(pictures = false)
        pictures.add(picture)
    }

    fun removePicture(pictureIndex : Int) {
        if (pictureIndex < 0 || pictureIndex >= pictures.size) return
        pictures.removeAt(pictureIndex)
    }

    private fun checkBeforeSubmit() : Boolean {
        val error = DamageInputError()
        if (_form.value.comment.isEmpty()) {
            error.comment = true
        }
        if (pictures.isEmpty()) {
            error.pictures = true
        }
        if (_form.value.room_id == null) {
            error.room = true
        }
        _formError.value = error
        return (error.comment || error.pictures || error.room)
    }

    fun submit() {
        if (checkBeforeSubmit()) {
            return
        }
        viewModelScope.launch {
            try {
            } catch (e : Exception) {
                e.printStackTrace()
            }
        }
    }
}