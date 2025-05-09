package com.example.keyz.profile

import android.app.Activity
import android.content.Context
import android.content.Intent
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import androidx.navigation.NavController
import com.example.keyz.apiCallerServices.ProfileCallerService
import com.example.keyz.apiCallerServices.ProfileUpdateInput
import com.example.keyz.apiClient.ApiService
import com.example.keyz.authService.AuthService
import com.example.keyz.login.dataStore
import com.example.keyz.utils.LanguageSetter
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.launch
import kotlinx.coroutines.runBlocking

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
    private val navController: NavController,
    apiService: ApiService
) : ViewModel() {
    private val apiCaller = ProfileCallerService(apiService, navController)
    private val authApiCaller = AuthService(navController.context.dataStore, apiService)
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

    fun logout() {
        viewModelScope.launch {
            authApiCaller.onLogout(navController)
        }
    }

    fun changeLanguageAndRestart(context: Context, language: String) {
        val localeSetter = LanguageSetter(context.dataStore)
        runBlocking {
            localeSetter.setLanguage(language)
        }
        val intent = context.packageManager.getLaunchIntentForPackage(context.packageName)
        intent?.addFlags(Intent.FLAG_ACTIVITY_CLEAR_TOP or Intent.FLAG_ACTIVITY_NEW_TASK)
        context.startActivity(intent)
        if (context is Activity) {
            context.finish()
        }
        Runtime.getRuntime().exit(0) // To ensure full restart
    }
}

