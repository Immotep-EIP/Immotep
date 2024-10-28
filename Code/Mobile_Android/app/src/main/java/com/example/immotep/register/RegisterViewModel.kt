package com.example.immotep.register

import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.example.immotep.ApiClient.ApiClient
import com.example.immotep.ApiClient.RegistrationInput
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.launch
import androidx.navigation.NavController
import com.example.immotep.components.decodeRetroFitMessagesToHttpCodes

data class RegisterForm(
    val lastName: String = "",
    val firstName: String = "",
    val email: String = "",
    val password: String = "",
)

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

class RegisterViewModel : ViewModel() {
    private val _registerForm = MutableStateFlow(RegisterForm())
    private val _registerConfirm = MutableStateFlow(RegisterConfirm())
    private val _registerFormError = MutableStateFlow(RegisterFormError())
    val regForm: StateFlow<RegisterForm> = _registerForm.asStateFlow()
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
        viewModelScope.launch {
            try {
                ApiClient.apiService.register(RegistrationInput(
                    email = _registerForm.value.email,
                    password = _registerForm.value.password,
                    firstName = _registerForm.value.firstName,
                    lastName = _registerForm.value.lastName)
                )
                navController.navigate("login")
                return@launch
            } catch (err: Exception) {
                _registerFormError.value = _registerFormError.value.copy(apiError = decodeRetroFitMessagesToHttpCodes(err))
            }
        }
    }
}
