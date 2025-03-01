package com.example.immotep.inviteTenantModal

import androidx.compose.runtime.Composable
import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material3.Button
import androidx.compose.material3.ButtonDefaults
import androidx.compose.material3.DateRangePicker
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.ModalBottomSheet
import androidx.compose.material3.Text
import androidx.compose.material3.rememberModalBottomSheetState
import androidx.compose.runtime.collectAsState
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.platform.LocalConfiguration
import androidx.compose.ui.platform.testTag
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import androidx.lifecycle.viewmodel.compose.viewModel
import androidx.navigation.NavController
import com.example.immotep.layouts.BigModalLayout
import com.example.immotep.ui.components.DateRangeInput
import kotlinx.coroutines.launch
import com.example.immotep.R
import com.example.immotep.ui.components.OutlinedTextField

@Composable
fun InviteTenantModal(open: Boolean, close: () -> Unit, navController: NavController, propertyId : String) {
    val viewModel: InviteTenantViewModel = viewModel()
    val form = viewModel.invitationForm.collectAsState()
    val formError = viewModel.invitationFormError.collectAsState()

    BigModalLayout(height = 0.8f, open = open, close = close) {
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
            errorMessage = if (formError.value.date) stringResource(R.string.not_end_date_before_start) else null
        )
        Spacer(modifier = Modifier.height(16.dp))
        DateRangeInput(
            currentDate = form.value.endDate,
            onDateSelected = { date -> if (date != null) viewModel.setEndDate(date) },
            label = stringResource(R.string.end_date),
            errorMessage = if (formError.value.date) stringResource(R.string.not_end_date_before_start) else null
        )
        Spacer(modifier = Modifier.height(16.dp))
        Button(
            onClick = { viewModel.inviteTenant(
                navController = navController,
                close = close,
                propertyId = propertyId
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

