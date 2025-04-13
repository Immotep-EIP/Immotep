package com.example.immotep.inviteTenantModal

import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import androidx.navigation.NavController
import com.example.immotep.apiCallerServices.InviteInput
import com.example.immotep.apiCallerServices.InviteTenantCallerService
import com.example.immotep.apiClient.ApiService
import com.example.immotep.utils.RegexUtils
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

class InviteTenantViewModel(
    apiService: ApiService,
    navController: NavController
) : ViewModel() {
    private val callerService = InviteTenantCallerService(apiService, navController)
    private val _invitationForm = MutableStateFlow(InviteTenantInputForm())
    private val _invitationFormError = MutableStateFlow(InviteTenantInputFormError())

    val invitationForm = _invitationForm.asStateFlow()
    val invitationFormError = _invitationFormError.asStateFlow()

    fun reset() {
        _invitationForm.value = InviteTenantInputForm()
        _invitationFormError.value = InviteTenantInputFormError()
    }

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
        if (!RegexUtils.isValidEmail(_invitationForm.value.email)) {
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

    fun inviteTenant(
        close : () -> Unit,
        propertyId : String, onError : () -> Unit,
        onSubmit: (email: String, startDate: Long, endDate: Long) -> Unit,
        setIsLoading: (Boolean) -> Unit
    ) {
        if (!inviteTenantValidator()) {
            return
        }
        viewModelScope.launch {
            try {
                setIsLoading(true)
                close()
                callerService.invite(propertyId, _invitationForm.value.toInviteInput(), onError)
                onSubmit(_invitationForm.value.email, _invitationForm.value.startDate, _invitationForm.value.endDate)
                reset()
            } catch(e: Exception) {
                println(e)
            } finally {
                setIsLoading(false)
            }
        }
    }
}
