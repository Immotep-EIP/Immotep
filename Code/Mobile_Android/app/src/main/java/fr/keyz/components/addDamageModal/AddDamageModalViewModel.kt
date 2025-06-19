package fr.keyz.components.addDamageModal

import android.net.Uri
import androidx.compose.runtime.mutableStateListOf
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import androidx.navigation.NavController
import fr.keyz.apiCallerServices.Damage
import fr.keyz.apiCallerServices.DamageCallerService
import fr.keyz.apiCallerServices.DamageInput
import fr.keyz.apiCallerServices.DamagePriority
import fr.keyz.apiCallerServices.RoomCallerService
import fr.keyz.apiClient.ApiService
import fr.keyz.utils.Base64Utils
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.launch



class AddDamageModalViewModel(apiService: ApiService, private val navController: NavController) : ViewModel() {
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

    fun submit(addDamage : (Damage) -> Unit, tenantName: String) {
        if (checkBeforeSubmit()) {
            return
        }
        viewModelScope.launch {
            try {
                val roomName = rooms.find { it.id == _form.value.room_id }?.name
                if (roomName == null) {
                    throw Exception("room_not_found")
                }
                val imagesAsBase64 = pictures.map {
                    Base64Utils.encodeImageToBase64(it, navController.context)
                }
                _form.value.pictures.addAll(imagesAsBase64)
                val (id) = _apiCaller.addDamage(_form.value)
                addDamage(_form.value.toDamage(id, roomName, tenantName))
                reset()
            } catch (e : Exception) {
                e.printStackTrace()
            }
        }
    }
}