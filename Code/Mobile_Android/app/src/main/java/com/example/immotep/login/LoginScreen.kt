package com.example.immotep.login

import androidx.compose.foundation.background
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.text.KeyboardOptions
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.outlined.Lock
import androidx.compose.material3.Button
import androidx.compose.material3.Icon
import androidx.compose.material3.IconButton
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.OutlinedTextField
import androidx.compose.material3.Text
/* ktlint-disable no-wildcard-imports */
import androidx.compose.runtime.*
import androidx.compose.runtime.Composable
import androidx.compose.runtime.collectAsState
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.saveable.rememberSaveable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.testTag
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.text.AnnotatedString
import androidx.compose.ui.text.input.KeyboardType
import androidx.compose.ui.text.input.PasswordVisualTransformation
import androidx.compose.ui.text.input.VisualTransformation
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import androidx.lifecycle.viewmodel.compose.viewModel
import androidx.navigation.NavController
import com.example.immotep.R
import com.example.immotep.components.CheckBoxWithLabel
import com.example.immotep.components.Header
import com.example.immotep.components.TopText

@Composable
fun LoginScreen(
    navController: NavController,
    viewModel: LoginViewModel = viewModel(),
) {
    val emailAndPassword = viewModel.emailAndPassword.collectAsState()
    var showPassword by rememberSaveable { mutableStateOf(false) }

    Column(modifier = Modifier.background(MaterialTheme.colorScheme.background).fillMaxSize().padding(10.dp)) {
        Header()
        TopText(stringResource(R.string.login_hello), stringResource(R.string.login_details))
        Column(
            modifier = Modifier.fillMaxSize().padding(top = 50.dp, start = 20.dp, end = 20.dp),
            horizontalAlignment = Alignment.CenterHorizontally,
        ) {
            OutlinedTextField(
                label = { Text(stringResource(R.string.your_email)) },
                value = emailAndPassword.value.email,
                keyboardOptions = KeyboardOptions(keyboardType = KeyboardType.Email),
                onValueChange = { value ->
                    viewModel.updateEmailAndPassword(value, null, null)
                },
                modifier = Modifier.fillMaxWidth().testTag("loginEmailInput"),
            )
            OutlinedTextField(
                label = { Text(stringResource(R.string.your_password)) },
                value = emailAndPassword.value.password,
                onValueChange = { value ->
                    viewModel.updateEmailAndPassword(null, value, null)
                },
                keyboardOptions = KeyboardOptions(keyboardType = if (showPassword) KeyboardType.Text else KeyboardType.Password),
                visualTransformation = if (showPassword) VisualTransformation.None else PasswordVisualTransformation(),
                modifier = Modifier.padding(top = 10.dp).fillMaxWidth().testTag("loginPasswordInput"),
                trailingIcon = {
                    IconButton(
                        onClick = {
                            showPassword = !showPassword
                        },
                        modifier =
                        Modifier.testTag("togglePasswordVisibility")
                    ) {
                        Icon(Icons.Outlined.Lock, contentDescription = "Toggle password visibility")
                    }
                },
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
                    color = MaterialTheme.colorScheme.tertiary,
                    modifier =
                    Modifier.clickable { navController.navigate("forgotPassword") },
                )
            }
            Button(
                onClick = { navController.navigate("dashboard") },
                modifier = Modifier.testTag("loginButton"),
            ) { Text(stringResource(R.string.login_button)) }
            Row(verticalAlignment = Alignment.CenterVertically, modifier = Modifier.fillMaxWidth()) {
                Text(stringResource(R.string.no_account), color = MaterialTheme.colorScheme.primary, fontSize = 12.sp)
                Text(
                    stringResource(R.string.sign_up),
                    fontSize = 12.sp,
                    color = MaterialTheme.colorScheme.tertiary,
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
