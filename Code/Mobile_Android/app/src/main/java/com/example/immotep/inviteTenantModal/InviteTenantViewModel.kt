package com.example.immotep.inviteTenantModal

import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import androidx.navigation.NavController
import com.example.immotep.apiClient.ApiClient
import com.example.immotep.apiClient.ApiService
import com.example.immotep.apiClient.InviteInput
import com.example.immotep.authService.AuthService
import com.example.immotep.login.dataStore
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.launch
import java.text.SimpleDateFormat
import java.util.Date
import java.util.Locale

data class InviteTenantInputForm(
    val email: String = "",
    val startDate: Long = Date().time,
    val endDate: Long = Date().time
) {
    fun toInviteInput(): InviteInput {
        val formatter = SimpleDateFormat("yyyy-MM-dd'T'HH:mm:ss.SSSXXX", Locale.getDefault())
        return InviteInput(
            tenant_email = email,
            start_date = formatter.format(Date(startDate)),
            end_date = formatter.format(Date(endDate)))
    }
}

data class InviteTenantInputFormError(
    var email: Boolean = false,
    var date: Boolean = false
)

class InviteTenantViewModel : ViewModel() {
    private val _invitationForm = MutableStateFlow(InviteTenantInputForm())
    private val _invitationFormError = MutableStateFlow(InviteTenantInputFormError())

    val invitationForm = _invitationForm.asStateFlow()
    val invitationFormError = _invitationFormError.asStateFlow()

    fun setStartDate(startDate: Long) {
        println(startDate)
        _invitationForm.value = _invitationForm.value.copy(startDate = startDate)
    }

    fun setEndDate(endDate: Long) {
        _invitationForm.value = _invitationForm.value.copy(endDate = endDate)
    }

    fun setEmail(email: String) {
        _invitationForm.value = _invitationForm.value.copy(email = email)
    }

    private fun inviteTenantValidator() : Boolean {
        val newFormError = InviteTenantInputFormError()
        if (!android.util.Patterns.EMAIL_ADDRESS
                .matcher(_invitationForm.value.email)
                .matches()) {
            newFormError.email = true
        }
        if (Date(_invitationForm.value.startDate).after(Date(_invitationForm.value.endDate))) {
            newFormError.date = true
        }
        if (newFormError.email || newFormError.date) {
            _invitationFormError.value = newFormError
            return false
        }
        return true
    }

    fun inviteTenant(navController: NavController, close : () -> Unit, propertyId : String) {
        if (!inviteTenantValidator()) {
            return
        }
        viewModelScope.launch {
            val authService = AuthService(navController.context.dataStore)
            val bearerToken = try {
                authService.getBearerToken()
            } catch (e: Exception) {
                authService.onLogout(navController)
                return@launch
            }
            try {
                close()
                ApiClient.apiService.inviteTenant(
                    bearerToken,
                    propertyId ,
                    _invitationForm.value.toInviteInput()
                )
            } catch(e: Exception) {
                println(e)
            }
        }
    }

}