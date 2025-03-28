package com.example.immotep.profile

import androidx.compose.foundation.border
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.width
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.foundation.text.KeyboardOptions
import androidx.compose.material.MaterialTheme
import androidx.compose.material3.Button
import androidx.compose.material3.ButtonDefaults
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.collectAsState
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.platform.testTag
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.text.input.KeyboardType
import androidx.compose.ui.unit.dp
import androidx.lifecycle.viewmodel.compose.viewModel
import androidx.navigation.NavController
import com.example.immotep.LocalApiService
import com.example.immotep.R
import com.example.immotep.components.ErrorAlert
import com.example.immotep.components.LoadingDialog
import com.example.immotep.dashboard.DashBoardLayout
import com.example.immotep.ui.components.OutlinedTextField

@Composable
fun ProfileScreen(
    navController: NavController,
) {
    val apiService = LocalApiService.current
    val viewModel: ProfileViewModel = viewModel {
        ProfileViewModel(navController, apiService)
    }
    val infos = viewModel.infos.collectAsState()
    val isLoading = viewModel.isLoading.collectAsState()

    LaunchedEffect(Unit) {
        viewModel.initProfile()
    }
    DashBoardLayout(navController, "profile") {
        LoadingDialog(isLoading.value)
        Column(
            verticalArrangement = Arrangement.Center,
            horizontalAlignment = Alignment.CenterHorizontally,
            modifier = Modifier.fillMaxSize()
        ) {
            ErrorAlert(null, null, if (viewModel.apiError.collectAsState().value) stringResource(R.string.profile_api_error) else null)
            Column(
                horizontalAlignment = Alignment.CenterHorizontally,
                modifier = Modifier
                    .width(300.dp)
                    .height(450.dp)
                    .border(1.dp, color = MaterialTheme.colors.onBackground, shape = RoundedCornerShape(10.dp))
                    .padding(10.dp)
            ) {
                OutlinedTextField(
                    label = stringResource(R.string.last_name),
                    value = infos.value.lastname,
                    onValueChange = { viewModel.setLastName(it) },
                    modifier = Modifier.fillMaxWidth().testTag("profileLastName"),
                    errorMessage = if (false) stringResource(R.string.last_name_error) else null,
                )
                OutlinedTextField(
                    label = stringResource(R.string.first_name),
                    value = infos.value.firstname,
                    onValueChange = { viewModel.setFirstName(it) },
                    modifier = Modifier.fillMaxWidth().testTag("profileFirstName"),
                    errorMessage = if (false) stringResource(R.string.first_name_error) else null,
                )
                OutlinedTextField(
                    label = stringResource(R.string.your_email),
                    value = infos.value.email,
                    keyboardOptions = KeyboardOptions(keyboardType = KeyboardType.Email),
                    onValueChange = { viewModel.setEmail(it) },
                    modifier = Modifier.fillMaxWidth().testTag("profileEmail"),
                    errorMessage = if (false) stringResource(R.string.email_error) else null,
                )
                Button(
                    onClick = { viewModel.updateProfile() },
                    colors = ButtonDefaults.buttonColors(containerColor = androidx.compose.material3.MaterialTheme.colorScheme.tertiary),
                    modifier = Modifier
                        .clip(RoundedCornerShape(5.dp))
                        .padding(5.dp)
                        .fillMaxWidth()
                        .testTag("updateProfile")
                ) {
                    Text(
                        stringResource(R.string.update_profile),
                        color = androidx.compose.material3.MaterialTheme.colorScheme.onTertiary
                    )
                }
            }
        }
    }
}
