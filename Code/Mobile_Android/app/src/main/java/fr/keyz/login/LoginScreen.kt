package fr.keyz.login

import androidx.activity.compose.BackHandler
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
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.text.KeyboardOptions
import androidx.compose.foundation.verticalScroll
import androidx.compose.material3.Button
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.collectAsState
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
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
import fr.keyz.LocalApiService
import fr.keyz.LocalIsOwner
import fr.keyz.R
import fr.keyz.components.CheckBoxWithLabel
import fr.keyz.components.ErrorAlert
import fr.keyz.components.Header
import fr.keyz.components.LoadingDialog
import fr.keyz.components.OpenBrowserAnnotedString
import fr.keyz.components.OpenBrowserButton
import fr.keyz.components.TopText
import fr.keyz.ui.components.OutlinedTextField
import fr.keyz.ui.components.PasswordInput

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
    val isLoading = viewModel.isLoading.collectAsState()
    val columnPaddingApiError = if (errors.value.apiError == null) 40.dp else 20.dp

    BackHandler {}
    LoadingDialog(isLoading.value)
    Column(
        modifier = Modifier
            .background(MaterialTheme.colorScheme.background)
            .fillMaxSize()
            .padding(10.dp)
            .testTag("loginScreen")
            .verticalScroll(rememberScrollState())
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
                OpenBrowserAnnotedString("https://dev.space.keyz-app.fr/forgot-password", stringResource(R.string.forgot_password))
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
