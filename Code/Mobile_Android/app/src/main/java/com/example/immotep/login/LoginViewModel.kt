package com.example.immotep.login

import android.content.Context
import androidx.datastore.core.DataStore
import androidx.datastore.preferences.core.Preferences
import androidx.datastore.preferences.preferencesDataStore
import androidx.lifecycle.ViewModel
import androidx.lifecycle.ViewModelProvider
import androidx.lifecycle.viewModelScope
import androidx.navigation.NavController
import com.example.immotep.apiClient.ApiService
import com.example.immotep.authService.AuthService
import com.example.immotep.utils.RegexUtils
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.launch

data class LoginState(
    val email: String = "",
    val password: String = "",
    val keepSigned: Boolean = false,
)

data class LoginErrorState(
    val email: Boolean = false,
    val password: Boolean = false,
    val apiError: Int? = null,
)

val Context.dataStore: DataStore<Preferences> by preferencesDataStore(name = "tokens")

class LoginViewModel(
    private val navController: NavController,
    apiService: ApiService
) : ViewModel() {
    private val _emailAndPassword = MutableStateFlow(LoginState())
    private val _errors = MutableStateFlow(LoginErrorState())
    private val authService = AuthService(navController.context.dataStore, apiService)
    val emailAndPassword: StateFlow<LoginState> = _emailAndPassword.asStateFlow()
    val errors: StateFlow<LoginErrorState> = _errors.asStateFlow()

    fun updateEmailAndPassword(email: String?, password: String?, keepSigned: Boolean?) {
        _emailAndPassword.value =
            _emailAndPassword.value.copy(
                email = email ?: _emailAndPassword.value.email,
                password = password ?: _emailAndPassword.value.password,
                keepSigned = keepSigned ?: _emailAndPassword.value.keepSigned,
            )
    }
    fun login(setIsOwner: (Boolean) -> Unit) {
        var noError = true
        _errors.value = _errors.value.copy(email = false, password = false, apiError = null)
        if (!RegexUtils.isValidEmail(_emailAndPassword.value.email)) {
            _errors.value = _errors.value.copy(email = true)
            noError = false
        }
        if (_emailAndPassword.value.password.length < 3) {
            _errors.value = _errors.value.copy(password = true)
            noError = false
        }
        if (!noError) {
            return
        }
        viewModelScope.launch {
            try {
                authService.onLogin(username = _emailAndPassword.value.email, password = _emailAndPassword.value.password)
                setIsOwner(authService.isUserOwner())
                navController.navigate("dashboard")
                return@launch
            } catch (e: Exception) {
                println("error: ${e.message}")
                val messageAndCode = e.message?.split(",")
                if (messageAndCode != null && messageAndCode.size == 2) {
                    val code = messageAndCode[1].toInt()
                    _errors.value = _errors.value.copy(apiError = code)
                }
                return@launch
            }
        }
    }
}

