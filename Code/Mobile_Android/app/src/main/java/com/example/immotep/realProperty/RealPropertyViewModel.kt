package com.example.immotep.realProperty

import androidx.compose.runtime.mutableStateListOf
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import androidx.navigation.NavController
import com.example.immotep.apiCallerServices.AddPropertyInput
import com.example.immotep.apiCallerServices.DetailedProperty
import com.example.immotep.apiCallerServices.GetPropertyResponse
import com.example.immotep.apiCallerServices.RealPropertyCallerService
import com.example.immotep.apiClient.ApiClient
import com.example.immotep.apiClient.ApiService
import com.example.immotep.authService.AuthService
import com.example.immotep.login.dataStore
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.launch


class RealPropertyViewModel(
    navController: NavController,
    apiService: ApiService
) : ViewModel() {
    enum class WhichApiError {
         NONE,
        GET_PROPERTIES,
        ADD_PROPERTY,
        DELETE_PROPERTY
    }
    private val apiCaller = RealPropertyCallerService(apiService, navController)
    private val _isLoading = MutableStateFlow(true)
    private val _propertySelectedDetails = MutableStateFlow<DetailedProperty?>(null)
    private val _apiError = MutableStateFlow(WhichApiError.NONE)

    val isLoading: StateFlow<Boolean> = _isLoading.asStateFlow()
    val apiError: StateFlow<WhichApiError> = _apiError.asStateFlow()
    val propertySelectedDetails = _propertySelectedDetails.asStateFlow()
    val properties = mutableStateListOf<DetailedProperty>()

    fun closeError() {
        _apiError.value = WhichApiError.NONE
    }

    fun getProperties() {
        viewModelScope.launch {
            closeError()
            _isLoading.value = true
            try {
                properties.clear()
                properties.addAll(apiCaller.getPropertiesAsDetailedProperties())
            } catch (e : Exception) {
                _apiError.value = WhichApiError.GET_PROPERTIES
                println("error getting properties ${e.message}")
                e.printStackTrace()
            } finally {
                _isLoading.value = false
            }
        }
    }

    suspend fun addProperty(propertyForm: AddPropertyInput) {
        try {
            val newPropertyId = apiCaller.addProperty(propertyForm)
            properties.add(propertyForm.toDetailedProperty(newPropertyId.id))
            closeError()
        } catch (e : Exception) {
            _apiError.value = WhichApiError.ADD_PROPERTY
            println("error adding property ${e.message}")
        }
    }

    fun deleteProperty(propertyId: String) {
        val index = properties.indexOfFirst { it.id == propertyId }
        if (index == -1) {
            return
        }
        viewModelScope.launch {
            try {
                apiCaller.archiveProperty(propertyId, { _apiError.value = WhichApiError.DELETE_PROPERTY })
                properties.removeAt(index)
                closeError()
            } catch (e : Exception) {
                _apiError.value = WhichApiError.DELETE_PROPERTY
                println("error deleting property ${e.message}")
                e.printStackTrace()
            }
        }
    }

    fun setPropertySelectedDetails(propertyId: String) {
        val index = properties.indexOfFirst { it.id == propertyId }
        if (index == -1) {
            return
        }
        _propertySelectedDetails.value = properties[index]
    }

    fun getBackFromDetails(modifiedProperty : DetailedProperty) {
        val index = properties.indexOfFirst { it.id == modifiedProperty.id }
        if (index == -1) {
            return
        }
        properties[index] = modifiedProperty
        _propertySelectedDetails.value = null
    }
}
