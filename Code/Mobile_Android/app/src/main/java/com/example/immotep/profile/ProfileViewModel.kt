package com.example.immotep.profile

import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import androidx.navigation.NavController
import com.example.immotep.apiCallerServices.ProfileCallerService
import com.example.immotep.apiCallerServices.ProfileUpdateInput
import com.example.immotep.apiClient.ApiService
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.launch

data class ProfileState(
    val email: String = "",
    val firstname: String = "",
    val lastname: String = "",
    val role: String = "",
) {
    fun toProfileUpdateInput(): ProfileUpdateInput {
        return ProfileUpdateInput(
            email = email,
            firstname = firstname,
            lastname = lastname,
        )
    }
}

class ProfileViewModel(
    navController: NavController,
    apiService: ApiService
) : ViewModel() {
    private val apiCaller = ProfileCallerService(apiService, navController)
    private val _infos = MutableStateFlow(ProfileState())
    private val _apiError = MutableStateFlow(false)
    private val _isLoading = MutableStateFlow(false)

    val infos: StateFlow<ProfileState> = _infos.asStateFlow()
    val apiError: StateFlow<Boolean> = _apiError.asStateFlow()
    val isLoading: StateFlow<Boolean> = _isLoading.asStateFlow()


    fun initProfile() {
        viewModelScope.launch {
            _apiError.value = false
            try {
                val profile = apiCaller.getProfile()
                _infos.value = _infos.value.copy(
                    email = profile.email,
                    firstname = profile.firstname,
                    lastname = profile.lastname,
                    role = profile.role,
                )
            } catch (e: Exception) {
                _apiError.value = true
                println(e)
            }
        }
    }

    fun setEmail(email: String) {
        _infos.value = _infos.value.copy(email = email)
    }

    fun setFirstName(firstName: String) {
        _infos.value = _infos.value.copy(firstname = firstName)
    }

    fun setLastName(lastName: String) {
        _infos.value = _infos.value.copy(lastname = lastName)
    }

    fun updateProfile() {
        viewModelScope.launch {
            _apiError.value = false
            _isLoading.value = true
            try {
                apiCaller.updateProfile(_infos.value.toProfileUpdateInput())
            } catch (e: Exception) {
                _apiError.value = true
                println(e)
            } finally {
                _isLoading.value = false
            }
        }
    }

}

