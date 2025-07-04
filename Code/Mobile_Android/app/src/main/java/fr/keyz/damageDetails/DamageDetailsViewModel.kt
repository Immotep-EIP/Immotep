package fr.keyz.damageDetails

import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import androidx.navigation.NavController
import fr.keyz.apiCallerServices.Damage
import fr.keyz.apiCallerServices.DamageCallerService
import fr.keyz.apiClient.ApiService
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.launch

class DamageDetailsViewModel(apiService: ApiService, navController: NavController) : ViewModel() {
    private val apiCallerService = DamageCallerService(apiService, navController)
    private val _currentDamage = MutableStateFlow<Damage?>(null)
    private val _isLoading = MutableStateFlow(false)
    private val _apiError = MutableStateFlow(false)

    val currentDamage = _currentDamage.asStateFlow()
    val isLoading = _isLoading.asStateFlow()
    val apiError = _apiError.asStateFlow()

    fun getDamage(propertyId: String?, leaseId: String, damageId: String) {
        viewModelScope.launch {
            try {
                _isLoading.value = true
                _currentDamage.value = apiCallerService.getDamage(propertyId ?: "", leaseId, damageId)
            } catch (e: Exception) {
                _apiError.value = true
            }
        }
    }
}