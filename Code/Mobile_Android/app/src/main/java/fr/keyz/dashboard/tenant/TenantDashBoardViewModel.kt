package fr.keyz.dashboard.tenant

import androidx.compose.runtime.mutableStateListOf
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import androidx.navigation.NavController
import fr.keyz.apiCallerServices.ApiCallerServiceException
import fr.keyz.apiCallerServices.Damage
import fr.keyz.apiCallerServices.DamageCallerService
import fr.keyz.apiCallerServices.DetailedProperty
import fr.keyz.apiCallerServices.ProfileCallerService
import fr.keyz.apiCallerServices.RealPropertyCallerService
import fr.keyz.apiClient.ApiService
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.launch

class TenantDashBoardViewModel(navController: NavController, apiService: ApiService) : ViewModel() {
    private val profileApiCaller = ProfileCallerService(apiService, navController)
    private val realPropertyApiCaller = RealPropertyCallerService(
        apiService = apiService,
        navController = navController
    )
    private val damageApiCaller = DamageCallerService(apiService, navController)
    private val _userName = MutableStateFlow("")
    private val _property = MutableStateFlow<DetailedProperty?>(null)
    private val _isLoading = MutableStateFlow(false)
    private val _apiError = MutableStateFlow<Int?>(null)

    val userName = _userName.asStateFlow()
    val property = _property.asStateFlow()
    val isLoading = _isLoading.asStateFlow()
    val apiError = _apiError.asStateFlow()
    val damages = mutableStateListOf<Damage>()

    fun loadDashBoard() {
        viewModelScope.launch {
            try {
                _apiError.value = null
                _isLoading.value = true
                damages.clear()
                val profile = profileApiCaller.getProfile()
                _userName.value = profile.firstname
                val property = realPropertyApiCaller.getPropertyWithDetails()
                _property.value = property
                if (property.lease != null) {
                    damages.addAll(damageApiCaller.getPropertyDamages(property.id, property.lease.id))
                }
            } catch(e : ApiCallerServiceException) {
                _apiError.value = e.getCode()
            } catch(e : Exception) {
                e.printStackTrace()
            } finally {
                _isLoading.value = false
            }
        }
    }
}