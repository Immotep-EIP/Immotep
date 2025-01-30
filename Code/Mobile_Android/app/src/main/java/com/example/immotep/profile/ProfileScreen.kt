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
import androidx.compose.runtime.Composable
import androidx.compose.runtime.collectAsState
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.testTag
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.text.input.KeyboardType
import androidx.compose.ui.unit.dp
import androidx.lifecycle.viewmodel.compose.viewModel
import androidx.navigation.NavController
import com.example.immotep.R
import com.example.immotep.dashboard.DashBoardLayout
import com.example.immotep.ui.components.OutlinedTextField

@Composable
fun ProfileScreen(
    navController: NavController,
) {
    val viewModel: ProfileViewModel = viewModel(factory = ProfileViewModelFactory(navController))
    val infos = viewModel.infos.collectAsState()
    DashBoardLayout(navController, "profile") {
        Column(
            verticalArrangement = Arrangement.Center,
            horizontalAlignment = Alignment.CenterHorizontally,
            modifier = Modifier.fillMaxSize()
        ) {
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
                    onValueChange = { value -> },
                    modifier = Modifier.fillMaxWidth().testTag("profileLastName"),
                    errorMessage = if (false) stringResource(R.string.last_name_error) else null,
                )
                OutlinedTextField(
                    label = stringResource(R.string.first_name),
                    value = infos.value.firstname,
                    onValueChange = { value -> },
                    modifier = Modifier.fillMaxWidth().testTag("profileFirstName"),
                    errorMessage = if (false) stringResource(R.string.first_name_error) else null,
                )
                OutlinedTextField(
                    label = stringResource(R.string.your_email),
                    value = infos.value.email,
                    keyboardOptions = KeyboardOptions(keyboardType = KeyboardType.Email),
                    onValueChange = { value -> },
                    modifier = Modifier.fillMaxWidth().testTag("profileEmail"),
                    errorMessage = if (false) stringResource(R.string.email_error) else null,
                )
            }
        }
    }
}
