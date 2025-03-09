package com.example.immotep.realProperty.details

import androidx.lifecycle.ViewModel
import androidx.lifecycle.ViewModelProvider
import androidx.lifecycle.viewModelScope
import androidx.navigation.NavController
import com.example.immotep.apiCallerServices.AddPropertyInput
import com.example.immotep.apiCallerServices.DetailedProperty
import com.example.immotep.apiCallerServices.RealPropertyCallerService
import com.example.immotep.apiClient.ApiClient
import com.example.immotep.apiClient.ApiService
import com.example.immotep.authService.AuthService
import com.example.immotep.login.dataStore
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.launch
import java.time.LocalDateTime
import java.time.OffsetDateTime
import java.util.Date

class RealPropertyDetailsViewModel(
    private val navController: NavController,
    apiService: ApiService
) : ViewModel() {
    enum class ApiErrors {
        GET_PROPERTY,
        UPDATE_PROPERTY,
        NONE
    }
    private val apiCaller = RealPropertyCallerService(apiService, navController)
    private var _property = MutableStateFlow(DetailedProperty())
    private val _apiError = MutableStateFlow(ApiErrors.NONE)

    val property: StateFlow<DetailedProperty> = _property.asStateFlow()
    val apiError = _apiError.asStateFlow()

    fun loadProperty(propertyId: String) {
        _apiError.value = ApiErrors.NONE
        viewModelScope.launch {
            try {
                val newProperty = apiCaller.getPropertyWithDetails(propertyId) { _apiError.value = ApiErrors.GET_PROPERTY }
                println("Loading property ${newProperty.appartementNumber}")
                _property.value = newProperty
            } catch (e : Exception) {
                println("Error loading property ${e.message}")
                e.printStackTrace()
            }
        }
    }

    suspend fun editProperty(property: AddPropertyInput, propertyId: String) {
        _apiError.value = ApiErrors.NONE
        _property.value = apiCaller.updateProperty(property, propertyId) { _apiError.value = ApiErrors.UPDATE_PROPERTY }
    }
}
