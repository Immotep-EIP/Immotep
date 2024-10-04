package com.example.immotep.login

import androidx.compose.foundation.Image
import androidx.compose.foundation.background
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.text.KeyboardOptions
import androidx.compose.material3.Button
import androidx.compose.material3.Checkbox
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.OutlinedTextField
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.collectAsState
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.text.AnnotatedString
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.text.input.KeyboardType
import androidx.compose.ui.text.input.PasswordVisualTransformation
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import androidx.lifecycle.viewmodel.compose.viewModel
import androidx.navigation.NavController
import com.example.immotep.R

@Composable
fun LoginScreen(
    navController: NavController,
    viewModel: LoginViewModel = viewModel(),
) {
    val emailAndPassword = viewModel.emailAndPassword.collectAsState()
    Column(modifier = Modifier.background(MaterialTheme.colorScheme.background).fillMaxSize().padding(10.dp)) {
        Row(verticalAlignment = Alignment.CenterVertically) {
            Image(
                painter = painterResource(id = R.drawable.immotep_png_logo),
                contentDescription = stringResource(id = R.string.immotep_logo_desc),
                modifier = Modifier.size(50.dp).padding(end = 10.dp)
            )
            Text(stringResource(R.string.app_name), fontSize = 30.sp, color = MaterialTheme.colorScheme.primary)
        }
        Column(modifier = Modifier.fillMaxWidth().padding(top = 90.dp), horizontalAlignment = Alignment.CenterHorizontally,) {
            Text(stringResource(R.string.login_hello), fontSize = 30.sp, fontWeight = FontWeight.SemiBold, color = MaterialTheme.colorScheme.primary)
            Text(stringResource(R.string.login_details), fontSize = 15.sp, color = MaterialTheme.colorScheme.primary)
        }
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
                modifier = Modifier.fillMaxWidth()
            )
            OutlinedTextField(
                label = { Text(stringResource(R.string.your_password)) },
                value = emailAndPassword.value.password,
                onValueChange = { value ->
                    viewModel.updateEmailAndPassword(null, value, null)
                },
                keyboardOptions = KeyboardOptions(keyboardType = KeyboardType.Password),
                visualTransformation = PasswordVisualTransformation(),
                modifier = Modifier.padding(top = 10.dp).fillMaxWidth()
            )
            Row(verticalAlignment = Alignment.CenterVertically, horizontalArrangement = Arrangement.SpaceBetween, modifier = Modifier.fillMaxWidth()) {
                Row(verticalAlignment = Alignment.CenterVertically) {
                    Checkbox(
                        checked = emailAndPassword.value.keepSigned,
                        onCheckedChange = { value ->
                            viewModel.updateEmailAndPassword(
                                null,
                                null,
                                value
                            )
                        },
                    )
                    Text(stringResource(R.string.keep_signed), color = MaterialTheme.colorScheme.primary, fontSize = 12.sp)
                }
                Text(AnnotatedString(stringResource(R.string.forgot_password)), fontSize = 12.sp, color = MaterialTheme.colorScheme.tertiary, modifier = Modifier.clickable { navController.navigate("forgotPassword") })
            }
            Button(onClick = { navController.navigate("dashboard") }) { Text(stringResource(R.string.login_button)) }
            Row(verticalAlignment = Alignment.CenterVertically, modifier = Modifier.fillMaxWidth()) {
                Text(stringResource(R.string.no_account), color = MaterialTheme.colorScheme.primary, fontSize = 12.sp)
                Text(AnnotatedString(stringResource(R.string.sign_up)), fontSize = 12.sp, color = MaterialTheme.colorScheme.tertiary, modifier = Modifier.padding(start = 3.dp).clickable { navController.navigate("register") })
            }
        }
    }
}
