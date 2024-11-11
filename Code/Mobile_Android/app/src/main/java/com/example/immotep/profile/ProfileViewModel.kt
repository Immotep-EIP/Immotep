package com.example.immotep.profile

import androidx.lifecycle.ViewModel
import androidx.lifecycle.ViewModelProvider
import androidx.lifecycle.viewModelScope
import androidx.navigation.NavController
import com.example.immotep.ApiClient.ApiClient
import com.example.immotep.AuthService.AuthService
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

class ProfileViewModel(private val navController: NavController) : ViewModel() {
    private val _infos = MutableStateFlow(ProfileState())
    val infos: StateFlow<ProfileState> = _infos.asStateFlow()

    init {
        viewModelScope.launch {
            try {
                val authServ = AuthService(navController.context.dataStore)
                val profile = ApiClient.apiService.getProfile(authServ.getBearerToken())
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

class ProfileViewModelFactory(private val navController: NavController) :
    ViewModelProvider.Factory {
    override fun <T : ViewModel> create(modelClass: Class<T>): T {
        if (modelClass.isAssignableFrom(ProfileViewModel::class.java)) {
            @Suppress("UNCHECKED_CAST")
            return ProfileViewModel(navController) as T
        }
        throw IllegalArgumentException("Unknown ViewModel class")
    }
}
