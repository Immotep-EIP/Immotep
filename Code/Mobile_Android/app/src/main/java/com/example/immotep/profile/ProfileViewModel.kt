package com.example.immotep.profile

import androidx.lifecycle.ViewModel
import androidx.lifecycle.ViewModelProvider
import androidx.lifecycle.viewModelScope
import androidx.navigation.NavController
import com.example.immotep.apiCallerServices.callers.ProfileCallerService
import com.example.immotep.apiClient.ApiClient
import com.example.immotep.apiClient.ApiService
import com.example.immotep.authService.AuthService
import com.example.immotep.login.dataStore
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.launch

data class ProfileState(
    val email: String = "",
    val firstname: String = "",
    val lastname: String = "",
    val role: String = "",
)

class ProfileViewModel(
    navController: NavController,
    apiService: ApiService
) : ViewModel() {
    private val apiCaller = ProfileCallerService(apiService, navController)
    private val _infos = MutableStateFlow(ProfileState())
    private val _apiError = MutableStateFlow(false)

    val infos: StateFlow<ProfileState> = _infos.asStateFlow()
    val apiError: StateFlow<Boolean> = _apiError.asStateFlow()

    fun initProfile() {
        viewModelScope.launch {
            _apiError.value = false
            try {
                val profile = apiCaller.getProfile({ _apiError.value = true })
                _infos.value = _infos.value.copy(
                    email = profile.email,
                    firstname = profile.firstname,
                    lastname = profile.lastname,
                    role = profile.role,
                )
            } catch (e: Exception) {
                println(e)
            }
        }
    }
}

