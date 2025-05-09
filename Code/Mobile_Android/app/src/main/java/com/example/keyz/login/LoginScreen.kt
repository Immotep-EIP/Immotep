package com.example.keyz.login

import androidx.compose.foundation.background
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.text.KeyboardOptions
import androidx.compose.material3.Button
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.collectAsState
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.testTag
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.text.AnnotatedString
import androidx.compose.ui.text.input.KeyboardType
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import androidx.lifecycle.viewmodel.compose.viewModel
import androidx.navigation.NavController
import com.example.keyz.LocalApiService
import com.example.keyz.LocalIsOwner
import com.example.keyz.R
import com.example.keyz.components.CheckBoxWithLabel
import com.example.keyz.components.ErrorAlert
import com.example.keyz.components.Header
import com.example.keyz.components.TopText
import com.example.keyz.ui.components.OutlinedTextField
import com.example.keyz.ui.components.PasswordInput

@Composable
fun LoginScreen(
    navController: NavController,
) {
    val isOwner = LocalIsOwner.current
    val apiService = LocalApiService.current
    val viewModel: LoginViewModel = viewModel {
        LoginViewModel(navController, apiService)
    }
    val emailAndPassword = viewModel.emailAndPassword.collectAsState()
    val errors = viewModel.errors.collectAsState()
    val columnPaddingApiError = if (errors.value.apiError == null) 40.dp else 20.dp

    Column(
        modifier = Modifier
            .background(MaterialTheme.colorScheme.background)
            .fillMaxSize()
            .padding(10.dp)
            .testTag("loginScreen")
    ) {
        Header()
        TopText(stringResource(R.string.login_hello), stringResource(R.string.login_details))
        Column(
            modifier =
            Modifier.fillMaxSize().padding(
                top = columnPaddingApiError,
                start = 20.dp,
                end = 20.dp,
            ),
            horizontalAlignment = Alignment.CenterHorizontally,
        ) {
            ErrorAlert(errors.value.apiError, true)
            Spacer(modifier = Modifier.height(10.dp))
            OutlinedTextField(
                label = stringResource(R.string.your_email),
                value = emailAndPassword.value.email,
                keyboardOptions = KeyboardOptions(keyboardType = KeyboardType.Email),
                onValueChange = { value ->
                    viewModel.updateEmailAndPassword(value, null, null)
                },
                modifier = Modifier.fillMaxWidth().testTag("loginEmailInput"),
                errorMessage = if (errors.value.email) stringResource(R.string.email_error) else null,
            )
            PasswordInput(
                label = stringResource(R.string.your_password),
                value = emailAndPassword.value.password,
                onValueChange = { value ->
                    viewModel.updateEmailAndPassword(null, value, null)
                },
                modifier = Modifier.padding(top = 10.dp).fillMaxWidth().testTag("loginPasswordInput"),
                errorMessage = if (errors.value.password) stringResource(R.string.password_error) else null,
                iconButtonTestId = "togglePasswordVisibility",
            )
            Row(
                verticalAlignment = Alignment.CenterVertically,
                horizontalArrangement = Arrangement.SpaceBetween,
                modifier = Modifier.fillMaxWidth(),
            ) {
                CheckBoxWithLabel(
                    modifier = Modifier.testTag("keepSignedCheckbox"),
                    label =
                    stringResource(R.string.keep_signed),
                    isChecked = emailAndPassword.value.keepSigned,
                    onCheckedChange = { value ->
                        viewModel.updateEmailAndPassword(null, null, value)
                    },
                )
                Text(
                    AnnotatedString(
                        stringResource(R.string.forgot_password),
                    ),
                    fontSize = 12.sp,
                    color = MaterialTheme.colorScheme.secondary,
                    modifier =
                    Modifier.clickable { navController.navigate("forgotPassword") },
                )
            }
            Button(
                onClick = { viewModel.login({ isOwner.value = it }) },
                modifier = Modifier.testTag("loginButton"),
            ) { Text(stringResource(R.string.login_button)) }
            Row(verticalAlignment = Alignment.CenterVertically, modifier = Modifier.fillMaxWidth()) {
                Text(stringResource(R.string.no_account), color = MaterialTheme.colorScheme.primary, fontSize = 12.sp)
                Text(
                    stringResource(R.string.sign_up),
                    fontSize = 12.sp,
                    color = MaterialTheme.colorScheme.secondary,
                    modifier =
                    Modifier
                        .padding(start = 3.dp)
                        .clickable { navController.navigate("register") }
                        .testTag("loginScreenToRegisterButton"),
                )
            }
        }
    }
}
