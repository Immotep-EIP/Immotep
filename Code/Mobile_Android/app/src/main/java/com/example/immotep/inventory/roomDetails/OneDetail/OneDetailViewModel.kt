package com.example.immotep.inventory.roomDetails.OneDetail

import android.net.Uri
import androidx.compose.runtime.mutableStateListOf
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import androidx.navigation.NavController
import com.example.immotep.apiClient.AiCallInput
import com.example.immotep.apiClient.ApiClient
import com.example.immotep.authService.AuthService
import com.example.immotep.inventory.Cleanliness
import com.example.immotep.inventory.InventoryLocationsTypes
import com.example.immotep.inventory.RoomDetail
import com.example.immotep.inventory.State
import com.example.immotep.login.dataStore
import com.example.immotep.utils.Base64Utils
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.launch
import java.util.Vector

data class RoomDetailsError(
    var name: Boolean = false,
    var comment: Boolean = false,
    var status: Boolean = false,
    var picture: Boolean = false,
    var cleanliness: Boolean = false
)

class OneDetailViewModel : ViewModel() {
    private val _detail = MutableStateFlow(RoomDetail(name = "", id = ""))
    val detail = _detail.asStateFlow()
    val picture = mutableStateListOf<Uri>()
    val entryPictures = mutableStateListOf<String>()
    private val _errors = MutableStateFlow(RoomDetailsError())
    val errors = _errors.asStateFlow()

    fun reset(newDetail : RoomDetail?) {
        picture.clear()
        entryPictures.clear()
        if (newDetail != null) {
            _detail.value = newDetail
            picture.addAll(newDetail.pictures)
            if (newDetail.entryPictures != null) {
                entryPictures.addAll(newDetail.entryPictures)
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

    fun setCleanliness(cleanliness : Cleanliness) {
        _detail.value = _detail.value.copy(cleanliness = cleanliness)
        _errors.value = _errors.value.copy(cleanliness = false)
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
        if (_detail.value.cleanliness == Cleanliness.not_set) {
            error.cleanliness = true
        }
        if (picture.isEmpty()) {
            error.picture = true
        }
        if (error.name || error.comment || error.status || error.picture || error.cleanliness) {
            _errors.value = error
            return
        }
        _detail.value = _detail.value.copy(
            pictures = picture.toTypedArray(),
            completed = true,
            entryPictures = if (isExit) entryPictures.toTypedArray() else null
        )
        onModifyDetail(detail.value)
        reset(null)
    }

    fun onClose(onModifyDetail : (detail : RoomDetail) -> Unit, isExit: Boolean) {
        _detail.value = _detail.value.copy(
            pictures = picture.toTypedArray(),
            entryPictures = if (isExit) entryPictures.toTypedArray() else null
        )
        onModifyDetail(_detail.value)
        reset(null)
    }

    private fun summarize(navController: NavController, propertyId: String) {
        viewModelScope.launch {
            val authService = AuthService(navController.context.dataStore)
            val bearerToken = try {
                authService.getBearerToken()
            } catch (e: Exception) {
                authService.onLogout(navController)
                return@launch
            }
            try {
                val picturesInput = Vector<String>()
                picture.forEach {
                    picturesInput.add(Base64Utils.encodeImageToBase64(it, navController.context))
                }
                val aiResponse = ApiClient.apiService.aiSummarize(
                    authHeader = bearerToken,
                    propertyId = propertyId,
                    summarizeInput = AiCallInput(
                        id = _detail.value.id,
                        pictures = picturesInput,
                        type = InventoryLocationsTypes.furniture
                    )
                )
                _detail.value = _detail.value.copy(
                    cleanliness = aiResponse.cleanliness ?: _detail.value.cleanliness,
                    status = aiResponse.state ?: _detail.value.status,
                    comment = aiResponse.note ?: _detail.value.comment
                )
                _errors.value = RoomDetailsError()
            } catch (e : Exception) {
                println("impossible to analyze ${e.message}")
                e.printStackTrace()
            }
        }
    }

    private fun compare(oldReportId : String, navController: NavController, propertyId: String) {
        viewModelScope.launch {
            val authService = AuthService(navController.context.dataStore)
            val bearerToken = try {
                authService.getBearerToken()
            } catch (e: Exception) {
                authService.onLogout(navController)
                return@launch
            }
            try {
                val picturesInput = Vector<String>()
                picture.forEach {
                    picturesInput.add(Base64Utils.encodeImageToBase64(it, navController.context))
                }
                val aiResponse = ApiClient.apiService.aiCompare(
                    authHeader = bearerToken,
                    propertyId = propertyId,
                    oldReportId = oldReportId,
                    summarizeInput = AiCallInput(
                        id = _detail.value.id,
                        pictures = picturesInput,
                        type = InventoryLocationsTypes.furniture
                    )
                )
                _detail.value = _detail.value.copy(
                    cleanliness = aiResponse.cleanliness ?: _detail.value.cleanliness,
                    status = aiResponse.state ?: _detail.value.status,
                    comment = aiResponse.note ?: _detail.value.comment
                )
            } catch (e : Exception) {
                println("impossible to analyze ${e.message}")
                e.printStackTrace()
            }
        }
    }

    fun summarizeOrCompare(oldReportId : String?, navController: NavController, propertyId: String) {
        if (picture.isEmpty()) {
            _errors.value = _errors.value.copy(picture = true)
            println("picture is empty")
            return
        }
        println(oldReportId)
        if (oldReportId == null) {
            return summarize(navController, propertyId)
        }
        return compare(oldReportId, navController, propertyId)
    }
}