package fr.keyz.realProperty.tenant

import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import androidx.navigation.NavController
import fr.keyz.apiCallerServices.ApiCallerServiceException
import fr.keyz.apiCallerServices.DetailedProperty
import fr.keyz.apiCallerServices.RealPropertyCallerService
import fr.keyz.apiClient.ApiService
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.launch

class RealPropertyTenantViewModel(apiService: ApiService, navController: NavController) : ViewModel() {
    private val apiCaller = RealPropertyCallerService(
        apiService = apiService,
        navController = navController
    )
    private val _property = MutableStateFlow<DetailedProperty?>(null)
    private val _isLoading = MutableStateFlow(false)
    private val _loadingError = MutableStateFlow<Int?>(null)

    val property = _property.asStateFlow()
    val isLoading = _isLoading.asStateFlow()
    val loadingError = _loadingError.asStateFlow()

    fun loadProperty() {
        _loadingError.value = null
        viewModelScope.launch {
            _isLoading.value = true
            try {
                val property = apiCaller.getPropertyWithDetails()
                _property.value = property
            } catch (e : ApiCallerServiceException) {
                _loadingError.value = e.getCode()
            } catch (e: Exception) {
                println("Error loading property: ${e.message}")
                e.printStackTrace()
            } finally {
                _isLoading.value = false
            }

        }
    }

}