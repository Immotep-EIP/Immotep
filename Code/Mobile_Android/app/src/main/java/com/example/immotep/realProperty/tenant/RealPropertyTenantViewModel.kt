package com.example.immotep.realProperty.tenant

import androidx.compose.runtime.mutableStateOf
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import androidx.navigation.NavController
import com.example.immotep.apiCallerServices.DetailedProperty
import com.example.immotep.apiCallerServices.RealPropertyCallerService
import com.example.immotep.apiClient.ApiService
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.launch

class RealPropertyTenantViewModel(apiService: ApiService, navController: NavController) : ViewModel() {
    private val realPropertyCallerService = RealPropertyCallerService(
        apiService = apiService,
        navController = navController
    )
    private val _property = MutableStateFlow<DetailedProperty?>(null)
    private val _isLoading = MutableStateFlow(false)

    val property = _property.asStateFlow()
    val isLoading = _isLoading.asStateFlow()

    fun loadProperty() {
        viewModelScope.launch {
            _isLoading.value = true
            try {
                val property = realPropertyCallerService.getPropertyWithDetails()
                _property.value = property
            } catch (e: Exception) {
                println("Error loading property: ${e.message}")
                e.printStackTrace()
            } finally {
                _isLoading.value = false
            }

        }
    }

}