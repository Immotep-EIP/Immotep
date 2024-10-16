package com.example.immotep.login

import android.content.Context
import androidx.datastore.core.DataStore
import androidx.datastore.preferences.core.Preferences
import androidx.datastore.preferences.preferencesDataStore
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.example.immotep.AuthService.AuthService
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

class LoginViewModel : ViewModel() {
    private val _emailAndPassword = MutableStateFlow(LoginState())
    private val _errors = MutableStateFlow(LoginErrorState())
    val emailAndPassword: StateFlow<LoginState> = _emailAndPassword.asStateFlow()
    val errors: StateFlow<LoginErrorState> = _errors.asStateFlow()

    fun updateEmailAndPassword(
        email: String?,
        password: String?,
        keepSigned: Boolean?,
    ) {
        _emailAndPassword.value =
            _emailAndPassword.value.copy(
                email = email ?: _emailAndPassword.value.email,
                password = password ?: _emailAndPassword.value.password,
                keepSigned = keepSigned ?: _emailAndPassword.value.keepSigned,
            )
    }

    fun login(context: Context) {
        var noError = true
        _errors.value = _errors.value.copy(email = false, password = false, apiError = null)
        if (!android.util.Patterns.EMAIL_ADDRESS
                .matcher(_emailAndPassword.value.email)
                .matches()
        ) {
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
                AuthService(context.dataStore).onLogin(_emailAndPassword.value.email, _emailAndPassword.value.password)
                return@launch
            } catch (e: Exception) {
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
