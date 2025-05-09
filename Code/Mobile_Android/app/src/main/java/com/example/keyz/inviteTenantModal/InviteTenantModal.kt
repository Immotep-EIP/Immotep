package com.example.keyz.inviteTenantModal

import android.widget.Toast
import androidx.compose.runtime.Composable
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material3.Button
import androidx.compose.material3.ButtonDefaults
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Text
import androidx.compose.runtime.collectAsState
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.platform.testTag
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import androidx.lifecycle.viewmodel.compose.viewModel
import androidx.navigation.NavController
import com.example.keyz.LocalApiService
import com.example.keyz.layouts.BigModalLayout
import com.example.keyz.ui.components.DateRangeInput
import com.example.keyz.R
import com.example.keyz.ui.components.OutlinedTextField

@Composable
fun InviteTenantModal(
    open: Boolean,
    close: () -> Unit,
    navController: NavController,
    propertyId : String,
    onSubmit: (email: String, startDate: Long, endDate: Long) -> Unit,
    setIsLoading: (Boolean) -> Unit
) {
    val apiService = LocalApiService.current
    val viewModel: InviteTenantViewModel = viewModel {
        InviteTenantViewModel(apiService = apiService, navController = navController)
    }
    val inviteTenantApiError = stringResource(R.string.invite_tenant_api_error)
    val form = viewModel.invitationForm.collectAsState()
    val formError = viewModel.invitationFormError.collectAsState()

    BigModalLayout(height = 0.8f, open = open, close = close, testTag = "inviteTenantModal") {
        Text(
            text = stringResource(R.string.invite_tenant),
            style = MaterialTheme.typography.headlineMedium
        )
        Spacer(modifier = Modifier.height(16.dp))
        OutlinedTextField(
            value = form.value.email,
            onValueChange = { viewModel.setEmail(it) },
            label = stringResource(R.string.tenant_email),
            errorMessage = if (formError.value.email) stringResource(R.string.invalid_email) else null,
            modifier = Modifier.testTag("tenantEmail").fillMaxWidth()
        )
        Spacer(modifier = Modifier.height(16.dp))
        DateRangeInput(
            currentDate = form.value.startDate,
            onDateSelected = { date -> if (date != null) viewModel.setStartDate(date) },
            label = stringResource(R.string.start_date),
            errorMessage = if (formError.value.date) stringResource(R.string.not_end_date_before_start) else null,
            globalTestTag = "startDateInput"
        )
        Spacer(modifier = Modifier.height(16.dp))
        DateRangeInput(
            currentDate = form.value.endDate,
            onDateSelected = { date -> if (date != null) viewModel.setEndDate(date) },
            label = stringResource(R.string.end_date),
            errorMessage = if (formError.value.date) stringResource(R.string.not_end_date_before_start) else null,
            globalTestTag = "endDateInput"
        )
        Spacer(modifier = Modifier.height(16.dp))
        Button(
            onClick = { viewModel.inviteTenant(
                close = close,
                propertyId = propertyId,
                onError = { Toast.makeText(navController.context, inviteTenantApiError, Toast.LENGTH_LONG).show() },
                onSubmit = onSubmit,
                setIsLoading = setIsLoading
            ) },
            colors = ButtonDefaults.buttonColors(containerColor = MaterialTheme.colorScheme.tertiary),
            modifier = Modifier
                .clip(RoundedCornerShape(5.dp))
                .padding(5.dp)
                .fillMaxWidth()
                .testTag("sendInvitation")
        ) {
            Text(
                stringResource(R.string.send_invitation),
                color = MaterialTheme.colorScheme.onTertiary
            )
        }
    }
}

