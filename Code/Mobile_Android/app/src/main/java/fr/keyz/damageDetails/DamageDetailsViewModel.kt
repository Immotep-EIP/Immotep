package fr.keyz.damageDetails

import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import androidx.navigation.NavController
import fr.keyz.apiCallerServices.ApiCallerServiceException
import fr.keyz.apiCallerServices.Damage
import fr.keyz.apiCallerServices.DamageCallerService
import fr.keyz.apiCallerServices.DamageStatus
import fr.keyz.apiCallerServices.UpdateDamageInput
import fr.keyz.apiClient.ApiService
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.launch
import java.text.SimpleDateFormat
import java.time.Instant
import java.time.OffsetDateTime
import java.time.ZoneOffset
import java.util.Date
import java.util.Locale

class DamageDetailsViewModel(apiService: ApiService, navController: NavController) : ViewModel() {
    private val apiCallerService = DamageCallerService(apiService, navController)
    private val _currentDamage = MutableStateFlow<Damage?>(null)
    private val _isLoading = MutableStateFlow(false)
    private val _apiError = MutableStateFlow<Int?>(null)

    val currentDamage = _currentDamage.asStateFlow()
    val isLoading = _isLoading.asStateFlow()
    val apiError = _apiError.asStateFlow()

    fun getDamage(propertyId: String?, leaseId: String, damageId: String) {
        viewModelScope.launch {
            try {
                _apiError.value = null
                _isLoading.value = true
                _currentDamage.value = apiCallerService.getDamage(propertyId ?: "", leaseId, damageId)
            } catch (e: ApiCallerServiceException) {
                _apiError.value = e.getCode()
            } finally {
                _isLoading.value = false
            }
        }
    }

    fun onSubmitUpdateDamageResolution(date: Long?, propertyId: String?) {
        if (currentDamage.value == null || propertyId.isNullOrEmpty()) return
        viewModelScope.launch {
            try {
                _isLoading.value = true
                if (date == null) {
                    apiCallerService.fixDamage(
                        propertyId = propertyId,
                        leaseId =  currentDamage.value!!.leaseId,
                        damageId =  currentDamage.value!!.id,
                    )
                    _currentDamage.value = _currentDamage.value?.copy(fixStatus = DamageStatus.AWAITING_TENANT_CONFIRMATION)
                } else {
                    val formatter = SimpleDateFormat("yyyy-MM-dd'T'HH:mm:ss.SSSXXX", Locale.getDefault())
                    val updateInput = UpdateDamageInput(
                        fix_planned_at = formatter.format(Date(date)),
                        read = true
                    )
                    apiCallerService.updateDamageOwner(
                        propertyId = propertyId,
                        leaseId =  _currentDamage.value!!.leaseId,
                        damageId = _currentDamage.value!!.id,
                        updateDamageInput = updateInput
                    )
                    _currentDamage.value = _currentDamage.value?.copy(
                        fixStatus = DamageStatus.PLANNED,
                        fixPlannedAt = OffsetDateTime.ofInstant(Instant.ofEpochMilli(date), ZoneOffset.UTC),
                    )
                }
            } catch (e: ApiCallerServiceException) {
                _apiError.value = e.getCode()
            } catch (e : Exception) {
                println("Unexpected error in onSubmitUpdateDamageResolution: ${e.message}")
                _apiError.value = 500
            } finally {
                _isLoading.value = false
            }
        }
    }

    fun onConfirm(propertyId: String?) {
        if (
            currentDamage.value == null ||
            (currentDamage.value!!.fixStatus != DamageStatus.AWAITING_TENANT_CONFIRMATION &&
            currentDamage.value!!.fixStatus != DamageStatus.AWAITING_OWNER_CONFIRMATION)) return
        viewModelScope.launch {
            try {
                _isLoading.value = true
                apiCallerService.fixDamage(
                    propertyId,
                    currentDamage.value!!.leaseId,
                    currentDamage.value!!.id
                )
                getDamage(propertyId, currentDamage.value!!.leaseId, currentDamage.value!!.id)
            } catch (e: ApiCallerServiceException) {
                _apiError.value = e.getCode()
            } finally {
                _isLoading.value = false
            }
        }
    }
}