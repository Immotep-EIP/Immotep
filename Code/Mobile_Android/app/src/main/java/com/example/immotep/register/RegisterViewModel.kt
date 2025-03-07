package com.example.immotep.register

import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.example.immotep.apiClient.ApiClient
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.launch
import androidx.navigation.NavController
import com.example.immotep.authService.AuthService
import com.example.immotep.authService.RegistrationInput
import com.example.immotep.components.decodeRetroFitMessagesToHttpCodes
import com.example.immotep.login.dataStore


data class RegisterConfirm(
    val password: String = "",
    val agreeToTerms: Boolean = false,
)

data class RegisterFormError(
    var lastName: Boolean = false,
    var firstName: Boolean = false,
    var email: Boolean = false,
    var password: Boolean = false,
    var confirmPassword: Boolean = false,
    var agreeToTerms: Boolean = false,
    var apiError: Int? = null,
)

class RegisterViewModel(navController: NavController, apiClient: ApiClient) : ViewModel() {
    private val authService = AuthService(navController.context.dataStore, apiClient.apiService)
    private val _registerForm = MutableStateFlow(RegistrationInput(
        email = "",
        password = "",
        firstName = "",
        lastName = ""
    ))
    private val _registerConfirm = MutableStateFlow(RegisterConfirm())
    private val _registerFormError = MutableStateFlow(RegisterFormError())
    val regForm = _registerForm.asStateFlow()
    val regConfirm: StateFlow<RegisterConfirm> = _registerConfirm.asStateFlow()
    val regFormError: StateFlow<RegisterFormError> = _registerFormError.asStateFlow()

    fun setLastName(lastName: String) {
        _registerForm.value = _registerForm.value.copy(lastName = lastName)
    }

    fun setFirstName(firstName: String) {
        _registerForm.value = _registerForm.value.copy(firstName = firstName)
    }

    fun setEmail(email: String) {
        _registerForm.value = _registerForm.value.copy(email = email)
    }

    fun setPassword(password: String) {
        _registerForm.value = _registerForm.value.copy(password = password)
    }

    fun setConfirmPassword(password: String) {
        _registerConfirm.value = _registerConfirm.value.copy(password = password)
    }

    fun setAgreeToTerms(agreeToTerms: Boolean) {
        _registerConfirm.value = _registerConfirm.value.copy(agreeToTerms = agreeToTerms)
    }

    private fun confirmedRegister(): Boolean =
        _registerConfirm.value.password.length > 3 && _registerForm.value.password == _registerConfirm.value.password


    fun onSubmit(navController: NavController) {
        val error = RegisterFormError()
        var noError = true
        if (_registerForm.value.lastName.length <= 2 || _registerForm.value.lastName.length >= 30) {
            error.lastName = true
            noError = false
        }
        if (_registerForm.value.firstName.length <= 2 || _registerForm.value.firstName.length >= 30) {
            error.firstName = true
            noError = false
        }
        if (!android.util.Patterns.EMAIL_ADDRESS
                .matcher(_registerForm.value.email)
                .matches()
        ) {
            error.email = true
            noError = false
        }
        if (_registerForm.value.password.length < 8) {
            error.password = true
            noError = false
        }
        if (!confirmedRegister()) {
            error.confirmPassword = true
            noError = false
        }
        if (!_registerConfirm.value.agreeToTerms) {
            error.agreeToTerms = true
            noError = false
        }
        _registerFormError.value = error
        if (!noError) {
            return
        }
        this.registerToApi(navController)
    }

    private fun registerToApi(navController: NavController) {
        _registerFormError.value = RegisterFormError()
        viewModelScope.launch {
            try {
                authService.register(_registerForm.value)
                navController.navigate("login")
                _registerForm.value = _registerForm.value.copy(
                    email = "",
                    password = "",
                    firstName = "",
                    lastName = ""
                )
                _registerConfirm.value = RegisterConfirm()
                return@launch
            } catch (err: Exception) {
                _registerFormError.value = _registerFormError.value.copy(apiError = decodeRetroFitMessagesToHttpCodes(err))
            }
        }
    }
}
