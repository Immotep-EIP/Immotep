package com.example.immotep.register

import androidx.lifecycle.ViewModel
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.asStateFlow

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
        _registerForm.value.password == _registerConfirm.value.password && _registerConfirm.value.agreeToTerms

    fun onSubmit(): Boolean {
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
        if (!error.agreeToTerms) {
            error.agreeToTerms = true
            noError = false
        }
        _registerFormError.value = error
        return noError
    }
}
