package com.example.immotep.register

import androidx.compose.foundation.background
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.text.KeyboardOptions
import androidx.compose.material3.Button
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.collectAsState
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.text.AnnotatedString
import androidx.compose.ui.text.input.KeyboardType
import androidx.compose.ui.text.input.PasswordVisualTransformation
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import androidx.lifecycle.viewmodel.compose.viewModel
import androidx.navigation.NavController
import com.example.immotep.R
import com.example.immotep.components.CheckBoxWithLabel
import com.example.immotep.components.Header
import com.example.immotep.components.TopText
import com.example.immotep.ui.components.OutlinedTextField

@Composable
fun RegisterScreen(
    navController: NavController,
    viewModel: RegisterViewModel = viewModel(),
) {
    val registerForm = viewModel.regForm.collectAsState()
    val registerConfirm = viewModel.regConfirm.collectAsState()
    val errors = viewModel.regFormError.collectAsState()
    Column(modifier = Modifier.background(MaterialTheme.colorScheme.background).fillMaxSize().padding(10.dp)) {
        Header()
        TopText(stringResource(R.string.create_account), stringResource(R.string.create_account_subtitle), limitMarginTop = true)
        Column(
            modifier = Modifier.fillMaxSize().padding(top = 50.dp, start = 20.dp, end = 20.dp),
            horizontalAlignment = Alignment.CenterHorizontally,
        ) {
            OutlinedTextField(
                label = { Text(stringResource(R.string.last_name)) },
                value = registerForm.value.lastName,
                onValueChange = { value -> viewModel.setLastName(value) },
                modifier = Modifier.fillMaxWidth(),
                errorMessage = if (errors.value.lastName) stringResource(R.string.last_name_error) else null,
            )
            OutlinedTextField(
                label = { Text(stringResource(R.string.first_name)) },
                value = registerForm.value.firstName,
                onValueChange = { value -> viewModel.setFirstName(value) },
                modifier = Modifier.fillMaxWidth(),
                errorMessage = if (errors.value.firstName) stringResource(R.string.first_name_error) else null,
            )
            OutlinedTextField(
                label = { Text(stringResource(R.string.your_email)) },
                value = registerForm.value.email,
                keyboardOptions = KeyboardOptions(keyboardType = KeyboardType.Email),
                onValueChange = { value -> viewModel.setEmail(value) },
                modifier = Modifier.fillMaxWidth(),
                errorMessage = if (errors.value.lastName) stringResource(R.string.last_name_error) else null,
            )
            OutlinedTextField(
                label = { Text(stringResource(R.string.your_password)) },
                value = registerForm.value.password,
                onValueChange = { value -> viewModel.setPassword(value) },
                modifier = Modifier.fillMaxWidth(),
                keyboardOptions = KeyboardOptions(keyboardType = KeyboardType.Password),
                visualTransformation = PasswordVisualTransformation(),
                errorMessage = if (errors.value.password) stringResource(R.string.password_error) else null,
            )
            OutlinedTextField(
                label = { Text(stringResource(R.string.password_confirm)) },
                value = registerConfirm.value.password,
                onValueChange = { value -> viewModel.setConfirmPassword(value) },
                modifier = Modifier.fillMaxWidth(),
                keyboardOptions = KeyboardOptions(keyboardType = KeyboardType.Password),
                visualTransformation = PasswordVisualTransformation(),
                errorMessage = if (errors.value.confirmPassword) stringResource(R.string.password_confirm_error) else null,
            )
            CheckBoxWithLabel(
                label = stringResource(R.string.agree_terms),
                isChecked = registerConfirm.value.agreeToTerms,
                onCheckedChange = { value -> viewModel.setAgreeToTerms(value) },
                errorMessage = if (errors.value.agreeToTerms) stringResource(R.string.agree_terms_error) else null,
            )
            Button(onClick = {
                if (viewModel.onSubmit()) {
                    navController.navigate("login")
                }
            }) {
                Text(stringResource(R.string.sign_up))
            }
            Row(verticalAlignment = Alignment.CenterVertically, modifier = Modifier.fillMaxWidth()) {
                Text(stringResource(R.string.already_account), color = MaterialTheme.colorScheme.primary, fontSize = 12.sp)
                Text(
                    AnnotatedString(stringResource(R.string.sign_up)),
                    fontSize = 12.sp,
                    color = MaterialTheme.colorScheme.tertiary,
                    modifier = Modifier.padding(start = 3.dp).clickable { navController.navigate("login") },
                )
            }
        }
    }
}
