package com.example.immotep.realProperty

import androidx.compose.runtime.mutableStateListOf
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import androidx.navigation.NavController
import com.example.immotep.apiCallerServices.AddPropertyInput
import com.example.immotep.apiCallerServices.Property
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
    private val navController: NavController,
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
    private val _apiError = MutableStateFlow(WhichApiError.NONE)

    val isLoading: StateFlow<Boolean> = _isLoading.asStateFlow()
    val apiError: StateFlow<WhichApiError> = _apiError.asStateFlow()
    val properties = mutableStateListOf<Property>()


    fun getProperties() {
        viewModelScope.launch {
            _apiError.value = WhichApiError.NONE
            _isLoading.value = true
            properties.clear()
            try {
                properties.addAll(apiCaller.getPropertiesAsProperties({ _apiError.value = WhichApiError.GET_PROPERTIES }))
            } catch (e : Exception) {
                println("error getting properties ${e.message}")
                e.printStackTrace()
            } finally {
                _isLoading.value = false
            }
        }
    }

    suspend fun addProperty(propertyForm: AddPropertyInput) {
        val newProperty = apiCaller.addProperty(propertyForm) { _apiError.value = WhichApiError.ADD_PROPERTY }
        properties.add(newProperty)
        _apiError.value = WhichApiError.NONE
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
                _apiError.value = WhichApiError.NONE
            } catch (e : Exception) {
                println("error deleting property ${e.message}")
                e.printStackTrace()
            }
        }
    }
}
